package app

import (
	"github.com/freonservice/freon/internal/entities"
	"github.com/freonservice/freon/internal/filter"
)

func (a *appl) CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, text string) error {
	return a.repo.CreateTranslation(ctx, creatorID, localizationID, identifierID, text)
}

func (a *appl) GetTranslations(ctx Ctx, f filter.TranslationFilter) ([]*entities.Translation, error) {
	translations, err := a.repo.GetTranslations(ctx, f)
	return mappingArrayTranslation(translations), err
}

func (a *appl) DeleteTranslation(ctx Ctx, id int64) error {
	return a.repo.DeleteTranslation(ctx, id)
}

func (a *appl) UpdateTranslation(ctx Ctx, id int64, text string) error {
	return a.repo.UpdateTranslation(ctx, id, text)
}

func (a *appl) HideTranslation(ctx Ctx, id int64, hide bool) error {
	return a.repo.UpdateHideStatusTranslation(ctx, id, hide)
}

func (a *appl) GetTranslation(ctx Ctx, locale, identifierName string) (*entities.Translation, error) {
	t, err := a.repo.GetTranslation(ctx, locale, identifierName)
	return mappingTranslation(t), err
}

func (a *appl) GetGroupedTranslations(ctx Ctx, f filter.GroupedTranslationFilter) ([]*entities.GroupedTranslations, error) {
	gts, err := a.repo.GetGroupedTranslations(ctx, f)
	return mappingArrayGroupedTranslations(gts), err
}
