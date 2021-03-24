package main

import (
	"context"
	"fmt"
	"github.com/MarcSky/freon/internal/app"
	"io/ioutil"
	"time"

	"github.com/MarcSky/freon/internal/config"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

func migrationDB(db *sqlx.DB) error {
	_ = goose.SetDialect("postgres")

	current, err := goose.EnsureDBVersion(db.DB)
	if err != nil {
		return fmt.Errorf("failed to EnsureDBVersion: %v", errors.WithStack(err))
	}

	files, err := ioutil.ReadDir(config.MigrationPath)
	if err != nil {
		return err
	}

	migrations, err := goose.CollectMigrations(config.MigrationPath, current, int64(len(files)))
	if err != nil {
		return err
	}

	for _, m := range migrations {
		if err := m.Up(db.DB); err != nil {
			return err
		}
	}

	return nil
}

func (srv *service) createFirstAdmin() error {
	if len(config.DefaultAdminPass) == 0 || len(config.DefaultAdminPass) == 0 {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := srv.appl.RegisterUser(ctx, config.DefaultAdminEmail, config.DefaultAdminPass, "Freon", "Administrator", 0)
	if err != nil && err != app.ErrEmailIsUsed {
		return err
	}
	return nil
}
