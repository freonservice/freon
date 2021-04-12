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
		version       bool
		logLevel      string
		apiPort       int
		grpcPort      int
		statikPort    int
		serviceHost   string
		migrationPath string
		jwtSecretPath string
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
	flag.IntVar(&dbConf.maxIdleConns, "db.port", config.DBMaxIdleConns, "db max idle conns must be set")
	flag.IntVar(&dbConf.maxOpenConns, "db.port", config.DBMaxOpenConns, "db max open conns must be set")
	flag.IntVar(&cfg.apiPort, "api.port", config.APIServicePort, "listen `api port` must be >0")
	flag.IntVar(&cfg.grpcPort, "grpc.port", config.GrpcServicePort, "listen `grpc port` must be >0")
	flag.IntVar(&cfg.statikPort, "statik.port", config.StatikServicePort, "listen `statik port` must be >0")
	flag.StringVar(&cfg.serviceHost, "service.host", config.ServiceHost, "listen `service host`")
	flag.StringVar(&cfg.migrationPath, "migration_path", config.MigrationPath, "migration path cant be empty")
	flag.StringVar(&cfg.jwtSecretPath, "jwt_secret_path", config.JwtSecretKey, "jwt secret path cant be empty")
	flag.StringVar(&adminCred.email, "admin.email", config.DefaultAdminEmail, "admin email cant be empty")
	flag.StringVar(&adminCred.password, "admin.password", config.DefaultAdminPass, "admin password cant be empty")

	log.SetDefaultKeyvals(structlog.KeyUnit, "main")
}

func main() {
	Init()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(cfg.logLevel))
	log.Info(version.Get())

	r, err := dal.New(&repo.Config{
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

	ctxShutdown, shutdown := context.WithCancel(context.Background())
	err = runServe(r, ctxShutdown, shutdown)
	if err != nil {
		log.Fatal(err)
	}

	<-sigc
	log.Println("Graceful stop server")
	shutdown()
	r.Close()
	os.Exit(0)
}
