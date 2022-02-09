package domain

import "github.com/pkg/errors"

type StorageConfiguration struct {
	Use             int32
	S3Configuration *S3Configuration
}

type S3Configuration struct {
	SecretAccessKey string
	AccessKeyID     string
	Region          string
	Endpoint        string
	DisableSSL      bool
	ForcePathStyle  bool
}

func (s *S3Configuration) IsValid() error {
	if len(s.SecretAccessKey) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty SecretAccessKey")
	}
	if len(s.AccessKeyID) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty AccessKeyID")
	}
	if len(s.Region) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty Region")
	}
	if len(s.Endpoint) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty Endpoint")
	}
	return nil
}

type TranslationConfiguration struct {
	Auto         bool
	Use          int32
	MainLanguage string // main language using like source for translations
}

type SettingConfiguration struct {
	Storage     StorageConfiguration
	Translation TranslationConfiguration
	FirstLaunch bool
}

func (t TranslationConfiguration) UseAutoTranslation() bool {
	return t.Use > 0 && t.Auto
}
