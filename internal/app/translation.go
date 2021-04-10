package app

import (
	"github.com/freonservice/freon/internal/dao"
)

func (a *appl) CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, text string) error {
	return a.repo.CreateTranslation(ctx, creatorID, localizationID, identifierID, text)
}

func (a *appl) GetTranslations(ctx Ctx, localizationID int64) ([]*Translation, error) {
	filter := dao.TranslationFilter{
		LocalizationID: localizationID,
	}
	translations, err := a.repo.GetTranslations(ctx, filter)
	if err != nil {
		return nil, err
	}
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
