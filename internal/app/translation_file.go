package app

import (
	"fmt"
	"os"
	"time"

	"github.com/freonservice/freon/internal/parser/web"
	"github.com/pkg/errors"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/filter"
)

func (a *appl) CreateTranslationFile(ctx Ctx, platform, storageType string, creatorID, localizationID int64) error {
	localization, err := a.repo.GetLocalization(ctx, localizationID)
	if err != nil {
		return err
	}

	// Generate translation FILE
	parser := web.NewParser()
	parser.SetTranslations(nil)
	text, err := parser.Generate()
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s.%d.json", localization.LanguageName, time.Now().Unix())
	f, err := os.Create(fileName)
	if err != nil {
		return errors.Wrap(err, "CreateTranslationFile os.Create file")
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return errors.Wrap(err, "CreateTranslationFile WriteString")
	}

	return a.repo.CreateTranslationFile(
		ctx,
		fileName,
		"translation.json",
		getPlatformByString(platform), getStorageTypeByString(storageType),
		creatorID, localizationID,
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
