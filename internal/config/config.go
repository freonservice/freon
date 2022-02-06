package config

import (
	"github.com/MarcSky/environ"
)

var (
	ServiceHost           = environ.GetStrEnv("SERVICE_HOST", "")               // frontend service host
	APIServicePort        = environ.GetIntEnv("API_SERVICE_PORT", 4000)         //nolint:gomnd
	GrpcServicePort       = environ.GetIntEnv("GRPC_SERVICE_PORT", 4001)        //nolint:gomnd
	WebStaticServicePort  = environ.GetIntEnv("WEB_STATIC_SERVICE_PORT", 4002)  //nolint:gomnd
	DocsStaticServicePort = environ.GetIntEnv("DOCS_STATIC_SERVICE_PORT", 4003) //nolint:gomnd
	MigrationPath         = environ.GetStrEnv("MIGRATION_PATH", "/migrations")
	BadgerPath            = environ.GetStrEnv("BADGER_PATH", "/badger")
	TranslationFilesPath  = environ.GetStrEnv("TRANSLATION_FILES_PATH", "")
	CPULimit              = environ.GetIntEnv("CPU_LIMIT", -1)
	LibraURL              = environ.GetStrEnv("LIBRA_API_URL", "http://127.0.0.1:5000")

	// Default Admin User
	DefaultAdminEmail = environ.GetStrEnv("DEFAULT_ADMIN_EMAIL", "admin@admin.com")
	DefaultAdminPass  = environ.GetStrEnv("DEFAULT_ADMIN_PASS", "adminpassword")

	// Auth Configuration
	JwtSecretKey     = environ.MustGetString("JWT_SECRET_KEY")
	JwtTokenLifetime = environ.GetTimeDurationEnv("JWT_TOKEN_LIFETIME", "24h")

	// Postgres DB Configuration
	DBHost         = environ.MustGetString("DB_HOST")
	DBPort         = environ.GetIntEnv("DB_PORT", 5432) //nolint:gomnd
	DBUser         = environ.MustGetString("DB_USER")
	DBName         = environ.MustGetString("DB_NAME")
	DBPass         = environ.MustGetString("DB_PASS")
	DBMaxOpenConns = environ.GetIntEnv("DB_MAX_OPEN_CONNS", 10) //nolint:gomnd
	DBMaxIdleConns = environ.GetIntEnv("DB_MAX_IDLE_CONNS", 10) //nolint:gomnd

	// S3 Storage Configuration
	S3SecretAccessKey = environ.GetStrEnv("S3_SECRET_ACCESS_KEY", "")
	S3AccessKeyID     = environ.GetStrEnv("S3_ACCESS_KEY_ID", "")
	S3Region          = environ.GetStrEnv("S3_REGION", "eu-west-1")
	S3URL             = environ.GetStrEnv("S3_URL", "http://localhost:9000")
	S3AppleBucket     = environ.GetStrEnv("S3_APPLE_BUCKET", "apple")
	S3AndroidBucket   = environ.GetStrEnv("S3_ANDROID_BUCKET", "android")
	S3WebBucket       = environ.GetStrEnv("S3_WEB_BUCKET", "web")
	S3DisableSSL      = environ.GetBoolEnv("S3_DISABLE_SSL", true)
	S3ForcePathStyle  = environ.GetBoolEnv("S3_FORCE_PATH_STYLE", true)
)
