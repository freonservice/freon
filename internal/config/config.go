package config

import (
	"time"

	"github.com/MarcSky/environ"

	"github.com/powerman/must"
	"github.com/powerman/structlog"
)

const (
	LogRemote   = "remote"
	LogFunc     = "func"
	LogGRPCCode = "grpcCode"
)

var (
	FServiceHost      = environ.MustGetString("F_SERVICE_HOST")   // frontend service host
	FServicePort      = environ.GetIntEnv("F_SERVICE_PORT", 4000) // frontend service port
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

func Init() {
	must.AbortIf = must.PanicIf

	structlog.DefaultLogger.
		SetPrefixKeys(
			structlog.KeyTime, LogGRPCCode,
			structlog.KeyLevel, structlog.KeyUnit,
		).
		SetDefaultKeyvals(
			structlog.KeyTime, structlog.Auto,
			structlog.KeySource, structlog.Auto,
		).
		SetSuffixKeys(
			structlog.KeyStack,
			structlog.KeySource,
		).
		SetKeysFormat(map[string]string{
			structlog.KeyTime:   " %[2]s",
			structlog.KeySource: " %6[2]s",
			structlog.KeyUnit:   " %6[2]s",
			LogGRPCCode:         " %-16.16[2]s",
			"config":            " %+[2]v",
			"duration":          " %[2]q",
			"request":           " %[1]s=`% [2]X`",
			"response":          " %[1]s=`% [2]X`",
		}).
		SetTimeFormat(time.StampMicro)
}
