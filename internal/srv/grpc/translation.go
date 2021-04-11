package grpc

import (
	"context"

	"github.com/freonservice/freon/internal/app"

	"github.com/freonservice/freon/internal/filter"
	"github.com/freonservice/freon/pkg/api"

	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetListTranslations(ctx context.Context, req *api.GetListTranslationsReq) (*api.GetListTranslationsRes, error) {
	logger := structlog.FromContext(ctx, nil)

	trxs, err := s.app.GetTranslations(ctx, filter.TranslationFilter{Locale: req.Locale})
	switch errors.Cause(err) {
	default:
		return nil, status.Error(codes.Internal, "failed -> "+logger.Err(err).Error())
	case nil:
	}

	return &api.GetListTranslationsRes{
		Translations: mappingTranslations(trxs),
	}, nil
}

func (s *service) GetTranslation(ctx context.Context, req *api.GetTranslationReq) (*api.GetTranslationRes, error) {
	logger := structlog.FromContext(ctx, nil)

	trx, err := s.app.GetTranslation(ctx, req.Locale, req.IdentifierName)
	switch errors.Cause(err) {
	default:
		return nil, status.Error(codes.Internal, "failed -> "+logger.Err(err).Error())
	case app.ErrNotFound:
		return nil, status.Error(codes.NotFound, "entity not found")
	case nil:
	}

	return &api.GetTranslationRes{
		Translation: mappingTranslation(trx),
	}, nil
}
