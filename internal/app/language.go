package app

import (
	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/pkg/freonApi"
)

func (a *appl) GetSupportedLanguages(ctx Ctx) ([]*domain.Language, error) {
	switch a.svc.setting.GetCurrentSettingState().Translation.Use {
	case int32(freonApi.TranslationSource_TRANSLATION_LIBRA):
		return a.svc.translation.Languages(ctx)
	default:
	}
	languages, err := a.svc.repo.GetLanguages(ctx)
	if err != nil {
		return nil, err
	}
	return mappingArrayLanguages(languages), nil
}
