package main

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/freonservice/freon/internal/config"
	"github.com/freonservice/freon/internal/dal"
	"github.com/freonservice/freon/pkg/def"
	"github.com/freonservice/freon/pkg/repo"
	"github.com/freonservice/freon/pkg/version"

	_ "github.com/lib/pq"
	"github.com/powerman/structlog"
)

type (
	dbConfig struct {
		host         string
		port         int
		user         string
		pass         string
		name         string
		maxIdleConns int
		maxOpenConns int
	}
)

var (
	log    = structlog.New(structlog.KeyUnit, "main")
	dbConf dbConfig
	cfg    struct {
		version                bool
		logLevel               string
		apiPort                int
		grpcPort               int
		webStaticPort          int
		docsStaticPort         int
		cpuLimit               int
		serviceHost            string
		migrationPath          string
		badgerPath             string
		jwtSecretPath          string
		jwtExpiration          time.Duration
		translationFilesFolder string
		libraURL               string
		defaultLanguage        string
	}

	s3Storage struct {
		secretAccessKey string
		accessKeyID     string
		region          string
		appleBucket     string
		androidBucket   string
		webBucket       string
		url             string
		disableSSL      bool
		forcePathStyle  bool
	}

	adminCred struct {
		email    string
		password string
	}
)

func Init() {
	def.Init()
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&cfg.logLevel, "log.level", "debug", "log `level` (debug|info|warn|err)")
	flag.IntVar(&dbConf.port, "db.port", config.DBPort, "psql db port is not specified")
	flag.StringVar(&dbConf.host, "db.host", config.DBHost, "psql db host is not specified")
	flag.StringVar(&dbConf.user, "db.user", config.DBUser, "psql db user is not specified")
	flag.StringVar(&dbConf.name, "db.name", config.DBName, "psql db name is not specified")
	flag.StringVar(&dbConf.pass, "db.pass", config.DBPass, "psql db pass is not specified")
	flag.IntVar(&dbConf.maxIdleConns, "db.maxIdleConns", config.DBMaxIdleConns, "db max idle conns must be set")
	flag.IntVar(&dbConf.maxOpenConns, "db.maxOpenConns", config.DBMaxOpenConns, "db max open conns must be set")
	flag.IntVar(&cfg.apiPort, "api.port", config.APIServicePort, "listen `api port` must be >0")
	flag.IntVar(&cfg.grpcPort, "grpc.port", config.GrpcServicePort, "listen `grpc port` must be >0")
	flag.IntVar(&cfg.webStaticPort, "web.static.port", config.WebStaticServicePort, "listen `web static port` must be >0")
	flag.IntVar(&cfg.docsStaticPort, "docs.static.port", config.DocsStaticServicePort, "listen `docs static port` must be >0")
	flag.StringVar(&cfg.serviceHost, "service.host", config.ServiceHost, "listen `service host`")
	flag.StringVar(&cfg.migrationPath, "migration.path", config.MigrationPath, "migration path cant be empty")
	flag.StringVar(&cfg.badgerPath, "badger.path", config.BadgerPath, "badger path cant be empty")
	flag.StringVar(&cfg.jwtSecretPath, "jwt_secret_path", config.JwtSecretKey, "jwt secret path cant be empty")
	flag.DurationVar(&cfg.jwtExpiration, "jwt_expiration_time", config.JwtExpirationTime, "jwt expiration time")
	flag.StringVar(&adminCred.email, "admin.email", config.DefaultAdminEmail, "admin email cant be empty")
	flag.StringVar(&adminCred.password, "admin.password", config.DefaultAdminPass, "admin password cant be empty")
	flag.StringVar(&cfg.translationFilesFolder, "translation.folders", config.TranslationFilesPath, "translation files folder")
	flag.IntVar(&cfg.cpuLimit, "cpu-limit", config.CPULimit, "maximum usage cpu")
	flag.StringVar(&cfg.libraURL, "libra.url", config.LibraURL, "libra api url")
	flag.StringVar(&cfg.defaultLanguage, "defaultLanguage", config.DefaultLanguage, "default target language")

	flag.StringVar(&s3Storage.secretAccessKey, "s3.secret-access-key", config.S3SecretAccessKey, "s3.secret-access-key cant be empty")
	flag.StringVar(&s3Storage.accessKeyID, "s3.access-key-id", config.S3AccessKeyID, "s3.access-key-id cant be empty")
	flag.StringVar(&s3Storage.region, "s3.region", config.S3Region, "s3.region cant be empty")
	flag.StringVar(&s3Storage.url, "s3.url", config.S3URL, "s3.url cant be empty")
	flag.StringVar(&s3Storage.appleBucket, "s3.apple-bucket", config.S3AppleBucket, "s3.apple-bucket cant be empty")
	flag.StringVar(&s3Storage.androidBucket, "s3.android-bucket", config.S3AndroidBucket, "s3.android-bucket cant be empty")
	flag.StringVar(&s3Storage.webBucket, "s3.web-bucket", config.S3WebBucket, "s3.web-bucket cant be empty")
	flag.BoolVar(&s3Storage.disableSSL, "s3.disable-ssl", config.S3DisableSSL, "s3.disable-ssl cant be empty")
	flag.BoolVar(&s3Storage.forcePathStyle, "s3.force-path-style", config.S3ForcePathStyle, "s3.force-path-style cant be empty")

	log.SetDefaultKeyvals(structlog.KeyUnit, "main")
}

func main() {
	Init()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(cfg.logLevel))
	log.Info(version.Get(), "Maximum Usage CPU", cfg.cpuLimit)

	dbRepo, err := dal.New(&repo.Config{
		Host:          dbConf.host,
		Port:          dbConf.port,
		Name:          dbConf.name,
		User:          dbConf.user,
		Pass:          dbConf.pass,
		MaxIdleConns:  dbConf.maxIdleConns,
		MaxOpenConns:  dbConf.maxOpenConns,
		MigrationPath: cfg.migrationPath,
	}, log)
	if err != nil {
		log.Fatal(err)
	}

	settingRepo, err := dal.NewSettingRepo(cfg.badgerPath, cfg.defaultLanguage)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current Setting State %+v\n", settingRepo.GetCurrentSettingState())

	ctxShutdown, shutdown := context.WithCancel(context.Background())
	err = runServe(dbRepo, settingRepo, ctxShutdown, shutdown)
	if err != nil {
		log.Fatal(err)
	}

	<-sigc
	log.Println("Graceful stop server")
	shutdown()
	dbRepo.Close()
	_ = settingRepo.Close()
	os.Exit(0)
}
