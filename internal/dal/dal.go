package dal

import (
	"context"

	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/pkg/repo"

	"github.com/jmoiron/sqlx"
)

type Ctx = context.Context

type r struct {
	*repo.Repo
}

func (r *r) GetDB() *sqlx.DB {
	return r.DB
}

func New(ctx Ctx, cfg *repo.Config) (app.Repo, error) {
	r := &r{}
	var err error

	r.Repo, err = repo.New(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return r, nil
}
