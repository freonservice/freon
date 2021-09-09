package app

import (
	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/entities"
)

func (a *appl) GetVersion(ctx Ctx, localizationID, typeVersion int64) ([]*entities.Version, error) {
	var (
		err  error
		data []*dao.Version
	)
	if typeVersion >= 0 && typeVersion < 3 {
		data, err = a.repo.GetVersionFromTranslationFiles(ctx, localizationID, typeVersion)
	} else {
		data, err = a.repo.GetVersionFromTranslations(ctx, localizationID)
	}

	if err != nil {
		return nil, err
	}

	return apiArrayVersion(data), nil
}
