package repo

import (
	"context"
	"database/sql"
	"os"

	"github.com/cenkalti/backoff/v4"
	"github.com/jmoiron/sqlx"
	"github.com/powerman/structlog"
	"github.com/pressly/goose"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type Ctx = context.Context

type Repo struct {
	DB       *sqlx.DB
	ReformDB *reform.DB
	logger   *structlog.Logger
}

func New(cfg *Config, logger *structlog.Logger) (*Repo, error) {
	db, err := sqlx.Connect("postgres", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	var pingDB backoff.Operation = func() error {
		err = db.Ping()
		if err != nil {
			logger.Println("DB is not ready...backing off...")
			return err
		}
		logger.Println("DB is ready!")
		return nil
	}

	err = backoff.Retry(pingDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	err = migrationDB(db, cfg.MigrationPath)
	if err != nil {
		return nil, err
	}

	return &Repo{
		DB:       db,
		ReformDB: reform.NewDB(db.DB, postgresql.Dialect, nil),
		logger:   logger,
	}, nil
}

// Close closes connection to DB.
func (r *Repo) Close() {
	r.logger.WarnIfFail(r.DB.Close)
}

func migrationDB(db *sqlx.DB, migrationPath string) error {
	var err error

	err = goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	current, err := goose.EnsureDBVersion(db.DB)
	if err != nil {
		return err
	}
	files, err := os.ReadDir(migrationPath)
	if err != nil {
		return err
	}

	migrations, err := goose.CollectMigrations(migrationPath, current, int64(len(files)))
	if err != nil {
		return err
	}

	for i := range migrations {
		if err := migrations[i].Up(db.DB); err != nil {
			return err
		}
	}

	return nil
}

func (r *Repo) Tx(ctx Ctx, opts *sql.TxOptions, f func(*sqlx.Tx) error) (err error) {
	tx, err := r.DB.BeginTxx(ctx, opts)
	if err == nil { //nolint:nestif // No idea how to simplify.
		defer func() {
			if err := recover(); err != nil {
				if err := tx.Rollback(); err != nil {
					log := structlog.FromContext(ctx, nil)
					log.Warn("failed to tx.Rollback", "err", err)
				}
				panic(err)
			}
		}()
		err = f(tx)
		if err == nil {
			err = tx.Commit()
		} else if err := tx.Rollback(); err != nil { //nolint:govet
			log := structlog.FromContext(ctx, nil)
			log.Warn("failed to tx.Rollback", "method", "err", err)
		}
	}
	return err
}
