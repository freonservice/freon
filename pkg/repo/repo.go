package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

func New(ctx Ctx, cfg Config) (*Repo, error) {
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

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	for err != nil {
		nextErr := db.PingContext(ctx)
		if errors.Is(nextErr, context.DeadlineExceeded) || errors.Is(nextErr, context.Canceled) {
			log.WarnIfFail(db.Close)
			return nil, fmt.Errorf("db.Ping: %w", err)
		}
		err = nextErr
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

// Turn sqlx errors like `missing destination …` into panics
// https://github.com/jmoiron/sqlx/issues/529. As we can't distinguish
// between sqlx and other errors except driver ones, let's hope filtering
// driver errors is enough and there are no other non-driver regular errors.
func (r *Repo) strict(err error) error {
	switch {
	case err == nil:
	case errors.Is(err, ErrSchemaVer):
	case errors.Is(err, sql.ErrNoRows):
	case errors.Is(err, context.Canceled):
	case errors.Is(err, context.DeadlineExceeded):
	default:
		panic(err)
	}
	return err
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
		} else if err := tx.Rollback(); err != nil {
			log := structlog.FromContext(ctx, nil)
			log.Warn("failed to tx.Rollback", "method", "err", err)
		}
	}
	return err
}
