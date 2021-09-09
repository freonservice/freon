package app

import (
	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/entities"
	"github.com/freonservice/freon/internal/filter"
	"github.com/freonservice/freon/pkg/freonApi"
)

func (a *appl) GetVersion(ctx Ctx, localizationID, platform int64) ([]*entities.Version, error) {
	var (
		err  error
		data []*dao.Version
	)
	if platform >= int64(freonApi.PlatformType_PLATFORM_TYPE_WEB) && platform <= int64(freonApi.PlatformType_PLATFORM_TYPE_ANDROID) {
		data, err = a.repo.GetVersionFromTranslationFiles(
			ctx,
			filter.VersionTranslationFilesFilter{
				LocalizationID: localizationID,
				PlatformType:   platform,
			},
		)
	} else {
		data, err = a.repo.GetVersionFromTranslations(
			ctx,
			filter.VersionTranslationsFilter{LocalizationID: localizationID},
		)
	}

	if err != nil {
		return nil, err
	}

	return apiArrayVersion(data), nil
}
