package utils

import (
	"os"

	"github.com/pkg/errors"
)

// GenerateDocFolders generating working folders for ios, android and web localization files
func GenerateDocFolders(currentPath string) error {
	var err error

	if currentPath != "/" {
		currentPath += "/"
	}

	_, err = os.Stat(currentPath + "docs")
	if os.IsNotExist(err) {
		err = os.Mkdir(currentPath+"docs", os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "internal.utils.GenerateDocFolders error with creating folder /docs")
		}
	}

	var platforms = []string{currentPath + "docs/ios", currentPath + "docs/android", currentPath + "docs/web"}
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
