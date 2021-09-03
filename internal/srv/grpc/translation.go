package grpc

import (
	"context"

	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/filter"
	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetListTranslations(ctx context.Context, req *api.GetListTranslationsReq) (*api.GetListTranslationsRes, error) {
	logger := structlog.FromContext(ctx, nil)

	entities, err := s.app.GetGroupedTranslations(ctx, filter.GroupedTranslationFilter{Locales: req.Locales})
	switch errors.Cause(err) {
	default:
		return nil, status.Error(codes.Internal, "failed -> "+logger.Err(err).Error())
	case nil:
	}

	gts := make([]*api.GroupedTranslations, len(entities))
	for i, v := range entities {
		gts[i] = &api.GroupedTranslations{
			Locale:       v.Locale,
			Translations: mappingTranslations(v.Translations),
		}
	}

	return &api.GetListTranslationsRes{
		GroupedTranslations: gts,
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
