package app

import "github.com/freonservice/freon/internal/filter"

func (a *appl) CreateTranslationFile(ctx Ctx, platform, storageType string, creatorID, localizationID int64) error {
	localization, err := a.repo.GetLocalization(ctx, localizationID)
	if err != nil {
		return err
	}
	return a.repo.CreateTranslationFile(
		ctx,
		"translation."+localization.Locale+".json", "translation.json",
		getPlatformByString(platform), getStorageTypeByString(storageType),
		creatorID, localizationID,
	)
}

func (a *appl) GetTranslationFiles(ctx Ctx, f filter.TranslationFileFilter) ([]*TranslationFile, error) {
	c, err := a.repo.GetTranslationFiles(ctx, f)
	if err != nil {
		return nil, err
	}
	return mappingArrayTranslationFile(c), err
}

func (a *appl) DeleteTranslationFile(ctx Ctx, id int64) error {
	return a.repo.DeleteTranslationFile(ctx, id)
}
