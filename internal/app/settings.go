package app

import (
	"errors"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/storage/local"
	"github.com/freonservice/freon/internal/storage/s3"
	"github.com/freonservice/freon/pkg/freonApi"
)

func (a *appl) GetCurrentSettingState() domain.SettingConfiguration {
	return a.svc.setting.GetCurrentSettingState()
}

func (a *appl) SetTranslationConfiguration(ctx Ctx, data domain.TranslationConfiguration) error {
	if !a.svc.setting.GetCurrentSettingState().FirstLaunch {
		data.MainLanguage = a.svc.setting.GetCurrentSettingState().Translation.MainLanguage
	}
	return a.svc.setting.SetTranslationConfiguration(ctx, data)
}

func (a *appl) SetStorageConfiguration(ctx Ctx, data domain.StorageConfiguration) error {
	if data.Use == a.svc.setting.GetCurrentSettingState().Storage.Use {
		return nil
	}

	if !a.svc.setting.GetCurrentSettingState().FirstLaunch {
		return nil
	}

	var err error
	switch data.Use {
	case int32(freonApi.StorageType_STORAGE_TYPE_LOCAL):
		a.svc.storage = local.NewStorage(a.config.TranslationFilesFolder, a.logger)
	case int32(freonApi.StorageType_STORAGE_TYPE_S3):
		if data.S3Configuration == nil {
			return errors.New("s3 storage configuration entity is nil")
		}
		a.svc.storage, err = s3.NewStorage(&domain.S3Configuration{
			SecretAccessKey: data.S3Configuration.SecretAccessKey,
			AccessKeyID:     data.S3Configuration.AccessKeyID,
			Region:          data.S3Configuration.Region,
			Endpoint:        data.S3Configuration.Endpoint,
			DisableSSL:      data.S3Configuration.DisableSSL,
			ForcePathStyle:  data.S3Configuration.ForcePathStyle,
		}, a.logger)
	}
	if err != nil {
		return err
	}

	return a.svc.setting.SetStorageConfiguration(ctx, data)
}

func (a *appl) DisableSettingFirstLaunch(ctx Ctx) error {
	if !a.svc.setting.GetCurrentSettingState().FirstLaunch {
		return nil
	}
	return a.svc.setting.DisableSettingFirstLaunch(ctx)
}
