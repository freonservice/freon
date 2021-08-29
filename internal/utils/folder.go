package utils

import (
	"os"

	"github.com/pkg/errors"
)

// CreateOrCheckTranslationFilesFolder generating working folders for ios, android and web localization files
func CreateOrCheckTranslationFilesFolder(currentPath string) error {
	var err error

	_, err = os.Stat(currentPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(currentPath, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "internal.utils.GenerateDocFolders error with creating folder /docs")
		}
	}

	var platforms = []string{currentPath + "/ios", currentPath + "/android", currentPath + "/web"}
	for _, platform := range platforms {
		_, err = os.Stat(platform)
		if os.IsNotExist(err) {
			err = os.Mkdir(platform, os.ModePerm)
			if err != nil {
				return errors.Wrapf(err, "internal.utils.GenerateDocFolders error with creating folder %s", platform)
			}
		}
	}

	return nil
}
