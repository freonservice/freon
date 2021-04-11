package config

import (
	"github.com/MarcSky/environ"
)

var (
	ServiceHost       = environ.GetStrEnv("SERVICE_HOST", "")          // frontend service host
	APIServicePort    = environ.GetIntEnv("API_SERVICE_PORT", 4000)    // frontend service port
	GrpcServicePort   = environ.GetIntEnv("GRPC_SERVICE_PORT", 4001)   // grpc service port
	StatikServicePort = environ.GetIntEnv("STATIK_SERVICE_PORT", 4002) // statik service port
	DBHost            = environ.MustGetString("DB_HOST")
	DBPort            = environ.GetIntEnv("DB_PORT", 5432)
	DBUser            = environ.MustGetString("DB_USER")
	DBName            = environ.MustGetString("DB_NAME")
	DBPass            = environ.MustGetString("DB_PASS")
	DefaultAdminEmail = environ.GetStrEnv("DEFAULT_ADMIN_EMAIL", "admin@admin.com")
	DefaultAdminPass  = environ.GetStrEnv("DEFAULT_ADMIN_PASS", "adminpassword")
	JwtSecretKey      = environ.MustGetString("JWT_SECRET_KEY")
	JwtTokenLifetime  = environ.GetTimeDurationEnv("JWT_TOKEN_LIFETIME", "24h")
	MigrationPath     = environ.MustGetString("MIGRATION_PATH")
	DBMaxOpenConn     = environ.GetIntEnv("DB_MAX_OPEN_CONN", 10)
	DBMaxIdleConn     = environ.GetIntEnv("DB_MAX_IDLE_CONN", 10)
)
