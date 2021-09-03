package grpc

import (
	"context"

	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetListLocalizations(ctx context.Context, req *api.GetListLocalizationsReq) (*api.GetListLocalizationsRes, error) {
	logger := structlog.FromContext(ctx, nil)

	localizations, err := s.app.GetLocalizations(ctx)
	switch errors.Cause(err) {
	default:
		return nil, status.Error(codes.Internal, "failed -> "+logger.Err(err).Error())
	case nil:
	}

	return &api.GetListLocalizationsRes{
		Localizations: mappingLocalizations(localizations),
	}, nil
}
