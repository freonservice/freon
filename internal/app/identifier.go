package app

import (
	"context"
	"time"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/filter"

	"golang.org/x/text/language"
)

func (a *appl) CreateIdentifier(
	ctx Ctx, creatorID, categoryID, parentID int64, name, description, textSingular, textPlural string, platforms []string,
) error {
	identifierID, err := a.svc.repo.CreateIdentifier(
		ctx, creatorID, categoryID, parentID, name,
		description, textSingular, textPlural, createConcatenatedString(platforms),
	)
	if err != nil {
		return err
	}

	err = a.updateTranslationsForIdentifier(ctx, identifierID, textSingular, textPlural)
	if err != nil {
		return err
	}

	return nil
}

func (a *appl) GetIdentifiers(ctx Ctx, f filter.IdentifierFilter) ([]*domain.Identifier, error) {
	l, err := a.svc.repo.GetIdentifiers(ctx, f)
	if err != nil {
		return nil, err
	}
	return mappingArrayIdentifier(l), nil
}

func (a *appl) DeleteIdentifier(ctx Ctx, id int64) error {
	return a.svc.repo.DeleteIdentifier(ctx, id)
}

func (a *appl) UpdateIdentifier(
	ctx Ctx, id, categoryID, parentID int64, name, description, textSingular, textPlural string, platforms []string,
) error {
	err := a.svc.repo.UpdateIdentifier(
		ctx, id, categoryID, parentID, name, description, textSingular, textPlural, createConcatenatedString(platforms),
	)
	if err != nil {
		return err
	}

	err = a.updateTranslationsForIdentifier(ctx, id, textSingular, textPlural)
	if err != nil {
		return err
	}

	return err
}

func (a *appl) updateTranslationsForIdentifier(ctx context.Context, identifierID int64, singular, plural string) error {
	var hasExampleText = len(singular) > 0 || len(plural) > 0
	if !a.svc.setting.GetCurrentSettingState().Translation.UseAutoTranslation() || !hasExampleText {
		return nil
	}

	localizations, err := a.GetLocalizations(ctx)
	if err != nil {
		return err
	}

	sourceLanguage, err := language.Parse(a.svc.setting.GetCurrentSettingState().Translation.MainLanguage)
	if err != nil {
		return err
	}

	for i := range localizations {
		go func(localization *domain.Localization) {
			var err error
			targetLanguage, err := language.Parse(localization.Locale)
			if err != nil {
				a.logger.PrintErr("language parsing error", "err", err, "locale", localization.Locale)
				return
			}

			if sourceLanguage.String() == targetLanguage.String() {
				return
			}

			var (
				iTextSingular, iTextPlural string
			)

			ctxThread, cancel := context.WithTimeout(context.Background(), 30*time.Second) //nolint:gomnd
			defer cancel()

			if len(singular) > 0 {
				iTextSingular, err = a.svc.translation.Translate(
					ctxThread,
					singular,
					sourceLanguage,
					targetLanguage,
				)
				if err != nil {
					a.logger.PrintErr("libra translation error", "err", err)
					return
				}
			}

			if len(plural) > 0 {
				iTextPlural, err = a.svc.translation.Translate(
					ctxThread,
					plural,
					sourceLanguage,
					targetLanguage,
				)
				if err != nil {
					a.logger.PrintErr("libra translation error", "err", err)
					return
				}
			}

			err = a.svc.repo.UpdateTranslationWithMeta(ctxThread, localization.ID, identifierID, iTextSingular, iTextPlural)
			if err != nil {
				a.logger.PrintErr("update translation error", "err", err)
				return
			}
		}(localizations[i])
	}

	return nil
}
