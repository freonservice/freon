package dal

import (
	"context"

	"github.com/freonservice/freon/pkg/repo"

	"github.com/jmoiron/sqlx"
	"github.com/powerman/structlog"
)

type Ctx = context.Context

type Repo struct {
	*repo.Repo
}

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
