package app

import (
	"context"
	"time"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/filter"

	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

func (a *appl) CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, singular, plural string) error {
	return a.svc.repo.CreateTranslation(ctx, creatorID, localizationID, identifierID, singular, plural)
}

func (a *appl) GetTranslations(ctx Ctx, f filter.TranslationFilter) ([]*domain.Translation, error) {
	translations, err := a.svc.repo.GetTranslations(ctx, f)
	return mappingArrayTranslation(translations), err
}

func (a *appl) DeleteTranslation(ctx Ctx, id int64) error {
	return a.svc.repo.DeleteTranslation(ctx, id)
}

func (a *appl) UpdateTranslation(ctx Ctx, id int64, singular, plural string) error {
	return a.svc.repo.UpdateTranslation(ctx, id, singular, plural)
}

func (a *appl) UpdateStatusTranslation(ctx Ctx, id, status int64) error {
	return a.svc.repo.UpdateStatusTranslation(ctx, id, status)
}

func (a *appl) GetTranslation(ctx Ctx, locale, identifierName string) (*domain.Translation, error) {
	t, err := a.svc.repo.GetTranslation(ctx, locale, identifierName)
	return mappingTranslation(t), err
}

func (a *appl) GetGroupedTranslations(ctx Ctx, f filter.GroupedTranslationFilter) ([]*domain.GroupedTranslations, error) {
	gts, err := a.svc.repo.GetGroupedTranslations(ctx, f)
	return mappingArrayGroupedTranslations(gts), err
}

func (a *appl) Translate(ctx Ctx, text string, source, target language.Tag) (string, error) {
	if a.svc.setting.GetCurrentSettingState().Translation.Use == 0 {
		return "", ErrAutoTranslationDisable
	}
	return a.svc.translation.Translate(ctx, text, source, target)
}

func (a *appl) CreateAutoTranslationByID(ctx Ctx, id int64) error {
	ctxThread, cancel := context.WithTimeout(context.Background(), 30*time.Second) //nolint:gomnd
	defer cancel()

	t, err := a.svc.repo.GetTranslationByID(ctxThread, id)
	if err != nil {
		return err
	}

	var hasExampleText = len(t.Singular.String) > 0 || len(t.Plural.String) > 0
	if !a.svc.setting.GetCurrentSettingState().Translation.UseAutoTranslation() || !hasExampleText {
		return err
	}

	sourceLanguage, err := language.Parse(a.svc.setting.GetCurrentSettingState().Translation.MainLanguage)
	if err != nil {
		return err
	}

	targetLanguage, err := language.Parse(t.Localization.Locale)
	if err != nil {
		return err
	}

	if sourceLanguage.String() == targetLanguage.String() {
		return errors.New("cant translate equally source and target languages")
	}

	var (
		iTextSingular, iTextPlural string
	)

	if len(t.Singular.String) > 0 {
		iTextSingular, err = a.svc.translation.Translate(
			ctxThread,
			t.Singular.String,
			sourceLanguage,
			targetLanguage,
		)
		if err != nil {
			return err
		}
	}

	if len(t.Plural.String) > 0 {
		iTextPlural, err = a.svc.translation.Translate(
			ctxThread,
			t.Plural.String,
			sourceLanguage,
			targetLanguage,
		)
		if err != nil {
			return err
		}
	}

	err = a.svc.repo.UpdateTranslationWithMeta(ctxThread, t.LocalizationID, t.IdentifierID, iTextSingular, iTextPlural)
	if err != nil {
		return err
	}

	return nil
}
