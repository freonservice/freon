package main

import (
	"context"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/MarcSky/freon/api/openapi/frontend/restapi"
	"github.com/MarcSky/freon/internal/app"
	"github.com/MarcSky/freon/internal/auth"
	"github.com/MarcSky/freon/internal/config"
	"github.com/MarcSky/freon/internal/dal"
	"github.com/MarcSky/freon/internal/password"
	"github.com/MarcSky/freon/internal/srv/frontend"
	"github.com/MarcSky/freon/pkg/concurrent"
	"github.com/MarcSky/freon/pkg/netx"
	"github.com/MarcSky/freon/pkg/repo"
	"github.com/MarcSky/freon/pkg/serve"

	"github.com/powerman/structlog"
)

type Ctx = context.Context

type service struct {
	repo        app.Repo
	appl        app.Appl
	auth        app.Auth
	frontendSrv *restapi.Server
}

func (srv *service) runServe(ctxShutdown Ctx, shutdown func()) (err error) {
	log := structlog.FromContext(ctxShutdown, nil)

	srv.repo, err = srv.connectRepo(context.Background())
	if err != nil {
		return log.Err("failed to connect repo", "err", err)
	}

	err = migrationDB(srv.repo.GetDB())
	if err != nil {
		return log.Err("failed to migrate", "err", err)
	}

	srv.auth = auth.NewAuth(config.JwtSecretKey, srv.repo, log)
	srv.appl = app.New(srv.repo, srv.auth, password.New())

	err = srv.createFirstAdmin()
	if err != nil {
		return log.Err("failed create admin", "err", err)
	}

	srv.frontendSrv, err = frontend.NewServer(srv.auth, srv.appl, frontend.Config{
		Addr: netx.NewAddr(config.FServiceHost, config.FServicePort),
	})
	if err != nil {
		return log.Err("failed to frontend.NewServer", "err", err)
	}

	//go func() {
	//	err = srv.reactStatic(ctxShutdown)
	//	if err != nil {
	//		_ = log.Err("failed to reactStatic", "err", err)
	//	}
	//}()

	err = concurrent.Serve(ctxShutdown, shutdown,
		srv.serveFrontendOpenAPI,
	)
	if err != nil {
		return log.Err("failed to serve", "err", err)
	}

	return nil
}

func (srv *service) connectRepo(ctx Ctx) (app.Repo, error) {
	return dal.New(ctx, repo.Config{
		Host:        config.DBHost,
		Port:        config.DBPort,
		User:        config.DBUser,
		Pass:        config.DBPass,
		Name:        config.DBName,
		MaxIdleConn: config.DBMaxIdleConn,
		MaxOpenConn: config.DBMaxOpenConn,
	})
}

func (srv *service) serveFrontendOpenAPI(ctx Ctx) error {
	return serve.OpenAPI(ctx, srv.frontendSrv, "frontendApi")
}

func (srv *service) reactStatic(ctx Ctx) error {
	const FSPATH = "./client/dist/"
	fs := http.FileServer(http.Dir(FSPATH))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			fullPath := FSPATH + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
			_, err := os.Stat(fullPath)
			if err != nil {
				if !os.IsNotExist(err) {
					log.Fatal(err)
				}
				r.URL.Path = "/"
			}
		}
		fs.ServeHTTP(w, r)
	})
	return http.ListenAndServe(":4200", nil)
}
