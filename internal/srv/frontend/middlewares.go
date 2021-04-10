package frontend

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/freonservice/freon/pkg/def"

	"github.com/felixge/httpsnoop"
	"github.com/powerman/structlog"
	corspkg "github.com/rs/cors"
	"github.com/sebest/xff"
)

type middlewareFunc func(http.Handler) http.Handler

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Expires", "0")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Strict-Transport-Security", "max-age=15552000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}

func makeLogger(basePath string) middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			remote := xff.GetRemoteAddr(r)
			log := structlog.New(
				structlog.KeyTime, structlog.Auto,
				def.LogRemote, remote,
				def.LogHTTPStatus, "",
				def.LogHTTPMethod, r.Method,
				def.LogFunc, strings.TrimPrefix(r.URL.Path, basePath),
			).SetTimeFormat(time.StampMicro)

			r = r.WithContext(structlog.NewContext(r.Context(), log))

			next.ServeHTTP(w, r)
		})
	}
}
func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			switch err := recover(); err := err.(type) {
			default:
				log := structlog.FromContext(r.Context(), nil)
				log.PrintErr(err, structlog.KeyStack, structlog.Auto)
				w.WriteHeader(http.StatusInternalServerError)
			case nil:
			case net.Error:
				log := structlog.FromContext(r.Context(), nil)
				log.PrintErr(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func makeAccessLog() middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m := httpsnoop.CaptureMetrics(next, w, r)
			log := structlog.FromContext(r.Context(), nil)
			if m.Code < http.StatusInternalServerError {
				log.Info("handled", def.LogHTTPStatus, m.Code, "Method", r.Method, "URL", r.URL.Path)
			} else {
				log.PrintErr("failed to handle", def.LogHTTPStatus, m.Code, "Method", r.Method, "URL", r.URL.Path)
			}
		})
	}
}

func cors(next http.Handler) http.Handler {
	return corspkg.AllowAll().Handler(next)
}
