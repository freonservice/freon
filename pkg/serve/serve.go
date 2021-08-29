package serve

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/freonservice/freon/pkg/def"
	"github.com/freonservice/freon/pkg/netx"

	"github.com/powerman/must"
	"github.com/powerman/structlog"
	"google.golang.org/grpc"
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

func ServerGRPC(ctx Ctx, addr netx.Addr, srv *grpc.Server) error {
	log := structlog.FromContext(ctx, nil).New(def.LogServer, addr.String())

	listen, err := net.Listen("tcp", addr.String())
	if err != nil {
		return err
	}

	log.Info("serve", "service", "grpc", "addr", addr.String())
	errc := make(chan error, 1)
	go func() { errc <- srv.Serve(listen) }()

	select {
	case err = <-errc:
	case <-ctx.Done():
		log.Info("Stopping GRPC server")
		_ = srv.Stop
	}
	if err != nil {
		return log.Err("failed to serve grpc", "err", err)
	}
	log.Info("shutdown service name grpc")
	return nil
}

func ServerStatic(ctx Ctx, addr netx.Addr) error {
	log := structlog.FromContext(ctx, nil).New(def.LogServer, addr.String())
	const FSPATH = "./client/prod/"

	listen, err := net.Listen("tcp", addr.String())
	if err != nil {
		return err
	}
	log.Info("serve", "service", "static", "addr", addr.String())

	errc := make(chan error, 1)
	go func() { errc <- http.Serve(listen, nil) }()

	fs := http.FileServer(http.Dir(FSPATH))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			fullPath := FSPATH + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
			_, err = os.Stat(fullPath)
			if err != nil {
				if !os.IsNotExist(err) {
					log.Fatal(err)
				}
				r.URL.Path = "/"
			}
		}
		fs.ServeHTTP(w, r)
	})

	select {
	case err = <-errc:
	case <-ctx.Done():
		log.Info("Stopping static server")
	}
	if err != nil {
		return log.Err("failed to serve static", "err", err)
	}
	log.Info("shutdown service name static")
	return nil
}
