package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cenkalti/backoff/v4"
	"github.com/jmoiron/sqlx"
	"github.com/powerman/structlog"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type Ctx = context.Context

var (
	ErrSchemaVer = errors.New("unsupported DB schema version")
)

type Config struct {
	Host        string
	Port        int
	User        string
	Pass        string
	Name        string
	MaxIdleConn int
	MaxOpenConn int
}

type Repo struct {
	DB       *sqlx.DB
	ReformDB *reform.DB
	log      *structlog.Logger
}

func New(ctx Ctx, cfg *Config) (*Repo, error) {
	log := structlog.FromContext(ctx, nil)

	dbDsnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Name,
	)
	db, err := sqlx.Connect("postgres", dbDsnString)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetMaxOpenConns(cfg.MaxOpenConn)

	var pingDB backoff.Operation = func() error {
		err = db.Ping()
		if err != nil {
			log.Println("DB is not ready...backing off...")
			return err
		}
		log.Println("DB is ready!")
		return nil
	}

	err = backoff.Retry(pingDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	return &Repo{
		DB:       db,
		ReformDB: reform.NewDB(db.DB, postgresql.Dialect, nil),
		log:      log,
	}, nil
}

// Close closes connection to DB.
func (r *Repo) Close() {
	r.log.WarnIfFail(r.DB.Close)
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
