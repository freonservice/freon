package app

import (
	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/filter"
	"github.com/freonservice/freon/internal/generator"
	"github.com/freonservice/freon/internal/generator/android"
	"github.com/freonservice/freon/internal/generator/ios"
	"github.com/freonservice/freon/internal/generator/web"
	"github.com/freonservice/freon/internal/storage"
	api "github.com/freonservice/freon/pkg/freonApi"
)

func (a *appl) CreateTranslationFile(ctx Ctx, platform, storageType string, creatorID, localizationID int64) error {
	localization, err := a.repo.GetLocalization(ctx, localizationID)
	if err != nil {
		return err
	}

	data, err := a.repo.GetTranslations(ctx, filter.TranslationFilter{LocalizationID: localizationID})
	if err != nil {
		return err
	}
	translations := mappingArrayTranslation(data)

	// Generate translation FILE
	var p generator.Generator
	switch api.PlatformType(getPlatformByString(platform)) { //nolint:exhaustive
	case api.PlatformType_PLATFORM_TYPE_IOS:
		p = ios.NewGenerator()
	case api.PlatformType_PLATFORM_TYPE_ANDROID:
		p = android.NewGenerator()
	default:
		p = web.NewGenerator().SetPluralFormat(generator.PluralFormat18N)
	}
	p.SetTranslations(translations)
	document, err := p.Generate()
	if err != nil {
		return err
	}

	file, err := a.storage.Create(storage.FileParameter{
		LocalizationLocale: localization.Locale,
		TranslatedText:     document,
		Platform:           platform,
	})
	if err != nil {
		return err
	}

	return a.repo.CreateTranslationFile(
		ctx,
		file.Name,
		file.WebPath,
		file.S3FileID,
		file.S3Bucket,
		getPlatformByString(platform),
		getStorageTypeByString(storageType),
		creatorID,
		localizationID,
	)
}

func (a *appl) GetTranslationFiles(ctx Ctx, f filter.TranslationFileFilter) ([]*domain.TranslationFile, error) {
	c, err := a.repo.GetTranslationFiles(ctx, f)
	if err != nil {
		return nil, err
	}
	return mappingArrayTranslationFile(c), err
}

func (a *appl) DeleteTranslationFile(ctx Ctx, id int64) error {
	return a.repo.DeleteTranslationFile(ctx, id)
}
