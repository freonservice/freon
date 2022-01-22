package storage

import (
	"github.com/freonservice/freon/internal/generator"
	api "github.com/freonservice/freon/pkg/freonApi"
)

const (
	DefaultAppleFile   = "Localizable" // hidden format .strings/.stringsdict
	DefaultAndroidFile = "strings.xml"
)

var (
	IosFormat = []string{".strings", ".stringsdict"}
)

type Storage interface {
	Create(parameter FileParameter) (*File, error)
}

type FileParameter struct {
	LocalizationLocale string
	Platform           string
	TranslatedText     generator.Document
}

type File struct {
	Name     string
	WebPath  string
	S3FileID string
	S3Bucket string
}

func GetPlatformByString(platform string) int64 {
	switch platform {
	case "ios":
		return int64(api.PlatformType_PLATFORM_TYPE_IOS)
	case "android":
		return int64(api.PlatformType_PLATFORM_TYPE_ANDROID)
	default:
		return int64(api.PlatformType_PLATFORM_TYPE_WEB)
	}
}
