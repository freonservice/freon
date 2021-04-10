package serve

import (
	"context"
	"fmt"
	"net"

	"github.com/freonservice/freon/pkg/def"

	"github.com/powerman/must"
	"github.com/powerman/structlog"
)

type Ctx = context.Context

type OpenAPIServer interface {
	HTTPListener() (net.Listener, error)
	TLSListener() (net.Listener, error)
	Serve() error
	Shutdown() error
}

func OpenAPI(ctx Ctx, srv OpenAPIServer, service string) error {
	log := structlog.FromContext(ctx, nil).New(def.LogServer, service)

	for _, f := range []func() (net.Listener, error){srv.HTTPListener, srv.TLSListener} {
		ln, err := f()
		if err != nil {
			return fmt.Errorf("listen: %w", err)
		}
		if ln != nil {
			host, port, err := net.SplitHostPort(ln.Addr().String())
			must.NoErr(err)
			log.Info("serve", def.LogHost, host, def.LogPort, port)
		}
	}

	go func() { <-ctx.Done(); _ = srv.Shutdown() }()
	err := srv.Serve()
	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	log.Info("shutdown", "service name", service)
	return nil
}
