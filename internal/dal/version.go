package dal

import (
	"github.com/freonservice/freon/internal/dao"
)

func (r *Repo) GetVersionFromTranslationFiles(ctx Ctx, localizationID, typeVersion int64) ([]*dao.Version, error) {
	panic("implement me")
}

func (r *Repo) GetVersionFromTranslations(ctx Ctx, localizationID int64) ([]*dao.Version, error) {
	panic("implement me")
}
