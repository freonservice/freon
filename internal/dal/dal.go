package dal

import (
	"context"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/pkg/repo"
	"github.com/freonservice/freon/pkg/setting"

	"github.com/jmoiron/sqlx"
	"github.com/powerman/structlog"
)

type (
	Ctx = context.Context

	Repo struct {
		*repo.Repo
	}
)

func (r *Repo) GetDB() *sqlx.DB {
	return r.DB
}

func New(cfg *repo.Config, logger *structlog.Logger) (*Repo, error) {
	r := &Repo{}
	var err error

	r.Repo, err = repo.New(cfg, logger)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type SettingRepo struct {
	*setting.Storage

	state domain.SettingConfiguration
}

func NewSettingRepo(path string) (*SettingRepo, error) {
	r := &SettingRepo{}
	var err error

	r.Storage, err = setting.NewSetting(path)
	if err != nil {
		return nil, err
	}

	err = r.getSettingTranslateState()
	if err != nil {
		return nil, err
	}

	err = r.getSettingStorageState()
	if err != nil {
		return nil, err
	}

	return r, nil
}
