package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/freonservice/freon/api/openapi/frontend/restapi"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/auth"
	"github.com/freonservice/freon/internal/dal"
	"github.com/freonservice/freon/internal/password"
	"github.com/freonservice/freon/internal/srv/frontend"
	grpcServer "github.com/freonservice/freon/internal/srv/grpc"
	"github.com/freonservice/freon/internal/utils"
	"github.com/freonservice/freon/pkg/concurrent"
	"github.com/freonservice/freon/pkg/freonApi"
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
	err := utils.CreateOrCheckTranslationFilesFolder(cfg.translationFilesFolder)
	if err != nil {
		return errors.Wrap(err, "generate doc folders")
	}

	settingRepo, err := dal.NewSettingRepo(cfg.badgerPath)
	if err != nil {
		return errors.Wrap(err, "badger error")
	}

	authorization := auth.NewAuth(cfg.jwtSecretPath, repo, log)
	config := app.Config{TranslationFilesPath: cfg.translationFilesFolder}
	appl := app.New(repo, authorization, password.New(), settingRepo, config)

	err = createFirstAdmin(appl)
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

	state := appl.GetCurrentSettingState()
	if state.Storage.Use == int32(freonApi.StorageType_STORAGE_TYPE_LOCAL) {
		go srv.serverDocsStatic()
	}

	srv.grpcSrv = grpcServer.NewServer(appl)
	err = concurrent.Serve(ctxShutdown, shutdown,
		srv.serveFrontendOpenAPI,
		srv.serveGRPC,
		srv.serveWebStatic,
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

func (srv *service) serveWebStatic(ctx Ctx) error {
	addr := netx.NewAddr(cfg.serviceHost, cfg.webStaticPort)
	return serve.ServerWebStatic(ctx, addr)
}

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
		int64(freonApi.UserRole_USER_ROLE_ADMIN),
	)
	if err != nil && err != app.ErrEmailIsUsed {
		return err
	}
	return nil
}

func (srv *service) serverDocsStatic() {
	fs := http.FileServer(http.Dir("./docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	log.Println("Starting service docs static with port", cfg.docsStaticPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.docsStaticPort), nil)
	if err != nil {
		log.Fatal(err, "docs static serving")
	}
}
