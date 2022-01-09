package utils

import (
	"os"

	"github.com/pkg/errors"
)

// CreateOrCheckTranslationFilesFolder generating working folders for ios, android and web localization files
func CreateOrCheckTranslationFilesFolder(currentPath string) error {
	var err error

	err = CheckFolderExist(currentPath)
	if err != nil {
		return errors.Wrap(err, "internal.utils.GenerateDocFolders error with creating folder /docs")
	}

	var platforms = []string{currentPath + "/ios", currentPath + "/android", currentPath + "/web"}
	for i := range platforms {
		err = CheckFolderExist(platforms[i])
		if err != nil {
			return errors.Wrapf(err, "internal.utils.GenerateDocFolders error with creating folder %s", platforms[i])
		}
	}

	return nil
}

func CheckFolderExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
