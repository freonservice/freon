package local

import (
	"fmt"
	"os"
	"strings"

	"github.com/freonservice/freon/internal/storage"
	"github.com/freonservice/freon/internal/utils"
	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/pkg/errors"
	"github.com/powerman/structlog"
)

type store struct {
	docsPath string

	logger *structlog.Logger
}

func NewStorage(logger *structlog.Logger, docsPath string) storage.Storage {
	return &store{
		docsPath: docsPath,
		logger:   logger,
	}
}

func (s *store) Create(parameter storage.FileParameter) (*storage.File, error) {
	var (
		err                error
		fileName           string
		localizationFolder string

		storageFullPath = s.docsPath + "/" + parameter.Platform
		webFullPath     = "/docs/" + parameter.Platform
		platformType    = storage.GetPlatformByString(parameter.Platform)
	)
	switch api.PlatformType(platformType) { //nolint:exhaustive
	case api.PlatformType_PLATFORM_TYPE_IOS:
		localizationFolder = "/" + parameter.LocalizationLocale + ".lproj"
		storageFullPath += localizationFolder
		webFullPath += localizationFolder
		fileName = storage.DefaultAppleFile
	case api.PlatformType_PLATFORM_TYPE_ANDROID:
		localizationFolder = "/values-" + parameter.LocalizationLocale
		storageFullPath += localizationFolder
		webFullPath += localizationFolder
		fileName = storage.DefaultAndroidFile
	case api.PlatformType_PLATFORM_TYPE_WEB:
		fileName = fmt.Sprintf("%s.json", parameter.LocalizationLocale)
	}

	if api.PlatformType(platformType) != api.PlatformType_PLATFORM_TYPE_WEB {
		err = utils.CheckOrCreateFolder(storageFullPath)
		if err != nil {
			return nil, err
		}
	}

	if api.PlatformType(platformType) == api.PlatformType_PLATFORM_TYPE_IOS {
		text := []string{parameter.TranslatedText.TextFirst, parameter.TranslatedText.TextSecond}
		var webPaths []string
		for i := range text {
			err = s.saveFileLocalStorage(storageFullPath+"/"+fileName+storage.IosFormat[i], text[i])
			webPaths = append(webPaths, webFullPath+"/"+fileName+storage.IosFormat[i])
		}
		webFullPath = strings.Join(webPaths, ",")
	} else {
		storageFullPath += "/" + fileName
		webFullPath += "/" + fileName
		err = s.saveFileLocalStorage(storageFullPath, parameter.TranslatedText.TextFirst)
	}
	if err != nil {
		return nil, err
	}

	return &storage.File{
		Name:    fileName,
		WebPath: webFullPath,
	}, nil
}

func (s *store) saveFileLocalStorage(path, text string) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "saveFileLocalStorage os.Create file")
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return errors.Wrap(err, "saveFileLocalStorage WriteString")
	}

	return nil
}
