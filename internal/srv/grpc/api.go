package grpc

import (
	"context"
	"time"

	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/pkg/api"
	"github.com/freonservice/freon/pkg/middleware"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type (
	service struct {
		app app.Appl
	}
)

func (s *service) GetLatestTranslationFiles(
	ctx context.Context, req *api.GetLatestTranslationFilesReq) (*api.GetLatestTranslationFilesRes, error) {
	panic("implement me")
}

func NewServer(appl app.Appl) *grpc.Server {
	srv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    50 * time.Second,
			Timeout: 10 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             30 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.UnaryServerLogger,
			middleware.UnaryServerRecover,
			middleware.UnaryServerAccessLog,
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			middleware.StreamServerLogger,
			middleware.StreamServerRecover,
			middleware.StreamServerAccessLog,
		)),
	)

	api.RegisterFreonServiceServer(srv, &service{
		app: appl,
	})

	return srv
}
