package config

import (
	"github.com/MarcSky/environ"
)

var (
	ServiceHost          = environ.GetStrEnv("SERVICE_HOST", "")          // frontend service host
	APIServicePort       = environ.GetIntEnv("API_SERVICE_PORT", 4000)    //nolint:gomnd
	GrpcServicePort      = environ.GetIntEnv("GRPC_SERVICE_PORT", 4001)   //nolint:gomnd
	StaticServicePort    = environ.GetIntEnv("STATIC_SERVICE_PORT", 4002) //nolint:gomnd
	DBHost               = environ.MustGetString("DB_HOST")
	DBPort               = environ.GetIntEnv("DB_PORT", 5432) //nolint:gomnd
	DBUser               = environ.MustGetString("DB_USER")
	DBName               = environ.MustGetString("DB_NAME")
	DBPass               = environ.MustGetString("DB_PASS")
	DefaultAdminEmail    = environ.GetStrEnv("DEFAULT_ADMIN_EMAIL", "admin@admin.com")
	DefaultAdminPass     = environ.GetStrEnv("DEFAULT_ADMIN_PASS", "adminpassword")
	JwtSecretKey         = environ.MustGetString("JWT_SECRET_KEY")
	JwtTokenLifetime     = environ.GetTimeDurationEnv("JWT_TOKEN_LIFETIME", "24h")
	MigrationPath        = environ.GetStrEnv("MIGRATION_PATH", "/migrations")
	DBMaxOpenConns       = environ.GetIntEnv("DB_MAX_OPEN_CONNS", 10) //nolint:gomnd
	DBMaxIdleConns       = environ.GetIntEnv("DB_MAX_IDLE_CONNS", 10) //nolint:gomnd
	TranslationFilesPath = environ.GetStrEnv("TRANSLATION_FILES_PATH", "")
	CPULimit             = environ.GetIntEnv("CPU_LIMIT", -1)
)
