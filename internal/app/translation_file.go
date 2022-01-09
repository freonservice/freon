package app

import (
	"fmt"
	"os"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/filter"
	"github.com/freonservice/freon/internal/parser"
	"github.com/freonservice/freon/internal/parser/android"
	"github.com/freonservice/freon/internal/parser/ios"
	"github.com/freonservice/freon/internal/parser/web"
	"github.com/freonservice/freon/internal/utils"
	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/pkg/errors"
)

const (
	defaultAppleFile   = "Localizable.strings"
	defaultAndroidFile = "strings.xml"
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
	var p parser.Generator
	var fileName string
	var localizationFolder string
	var storageFullPath = a.config.TranslationFilesPath + "/" + platform
	var webFullPath = "/docs/" + platform
	platformType := getPlatformByString(platform)
	switch api.PlatformType(platformType) { //nolint:exhaustive
	case api.PlatformType_PLATFORM_TYPE_IOS:
		p = ios.NewGenerator()
		localizationFolder = "/" + localization.Locale + ".lproj"
		storageFullPath += localizationFolder
		webFullPath += localizationFolder
		fileName = defaultAppleFile
		_ = utils.CheckAndCreateFolder(storageFullPath)
	case api.PlatformType_PLATFORM_TYPE_ANDROID:
		p = android.NewGenerator()
		localizationFolder = "/values-" + localization.Locale
		storageFullPath += localizationFolder
		webFullPath += localizationFolder
		fileName = defaultAndroidFile
		_ = utils.CheckAndCreateFolder(storageFullPath)
	default:
		p = web.NewGenerator().SetPluralFormat(parser.PluralFormat18N)
	}
	p.SetTranslations(translations)
	text, err := p.Generate()
	if err != nil {
		return err
	}

	if api.PlatformType(platformType) == api.PlatformType_PLATFORM_TYPE_WEB {
		fileName = fmt.Sprintf("%s.json", localization.Locale)
	}
	storageFullPath += "/" + fileName
	webFullPath += "/" + fileName
	f, err := os.Create(storageFullPath)
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
		webFullPath,
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
