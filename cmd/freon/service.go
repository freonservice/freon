package main

import (
	"context"
	"time"

	"github.com/freonservice/freon/api/openapi/frontend/restapi"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/auth"
	"github.com/freonservice/freon/internal/dal"
	"github.com/freonservice/freon/internal/password"
	"github.com/freonservice/freon/internal/srv/frontend"
	grpcServer "github.com/freonservice/freon/internal/srv/grpc"
	"github.com/freonservice/freon/pkg/api"
	"github.com/freonservice/freon/pkg/concurrent"
	"github.com/freonservice/freon/pkg/netx"
	"github.com/freonservice/freon/pkg/serve"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Ctx = context.Context

type service struct {
	frontendSrv *restapi.Server
	grpcSrv     *grpc.Server
}

func runServe(repo *dal.Repo, ctxShutdown Ctx, shutdown func()) error {
	authorization := auth.NewAuth(cfg.jwtSecretPath, repo, log)
	appl := app.New(repo, authorization, password.New())

	err := createFirstAdmin(appl)
	if err != nil {
		return errors.Wrap(err, "failed create admin")
	}

	srv := service{}
	srv.frontendSrv, err = frontend.NewServer(authorization, appl, frontend.Config{
		Addr: netx.NewAddr(cfg.serviceHost, cfg.apiPort),
	})
	if err != nil {
		return errors.Wrap(err, "failed to frontend.NewServer")
	}

	srv.grpcSrv = grpcServer.NewServer(appl)

	// nolint:gocritic
	// go func() {
	//	err = srv.reactStatic(ctxShutdown)
	//	if err != nil {
	//		_ = log.Err("failed to reactStatic", "err", err)
	//	}
	// }()

	err = concurrent.Serve(ctxShutdown, shutdown,
		srv.serveFrontendOpenAPI,
		srv.serveGRPC,
	)
	if err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return nil
}

func (srv *service) serveFrontendOpenAPI(ctx Ctx) error {
	return serve.OpenAPI(ctx, srv.frontendSrv, "frontendApi")
}

func (srv *service) serveGRPC(ctx Ctx) error {
	addr := netx.NewAddr(cfg.serviceHost, cfg.grpcPort)
	return serve.ServerGRPC(ctx, addr, srv.grpcSrv)
}

// func (srv *service) reactStatic(ctx Ctx) error {
//	const FSPATH = "./client/dist/"
//	fs := http.FileServer(http.Dir(FSPATH))
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		if r.URL.Path != "/" {
//			fullPath := FSPATH + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
//			_, err := os.Stat(fullPath)
//			if err != nil {
//				if !os.IsNotExist(err) {
//					log.Fatal(err)
//				}
//				r.URL.Path = "/"
//			}
//		}
//		fs.ServeHTTP(w, r)
//	})
//	return http.ListenAndServe(":4200", nil)
// }

func createFirstAdmin(appl app.Appl) error {
	if adminCred.email == "" || adminCred.password == "" {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //nolint:gomnd
	defer cancel()
	_, err := appl.RegisterUser(
		ctx,
		adminCred.email,
		adminCred.password,
		"Freon",
		"Administrator",
		int64(api.UserRole_USER_ROLE_ADMIN),
	)
	if err != nil && err != app.ErrEmailIsUsed {
		return err
	}
	return nil
}
