package def

import (
	"time"

	"github.com/powerman/must"
	"github.com/powerman/structlog"
)

const (
	LogServer     = "server"
	LogGRPCCode   = "grpcCode"
	LogRemote     = "remote"
	LogHost       = "host"
	LogPort       = "port"
	LogHTTPMethod = "httpMethod"
	LogHTTPStatus = "httpStatus"
	LogFunc       = "func"
	LogUserID     = "userID"
	LogHandler    = "handler"
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
