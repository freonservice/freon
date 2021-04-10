package dal

import (
	"database/sql"
	"time"

	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/dao"

	"github.com/jmoiron/sqlx"
	"gopkg.in/reform.v1"
)

func (r *r) SaveSession(ctx Ctx, userID int64, token string) error {
	return r.ReformDB.InTransactionContext(ctx, &sql.TxOptions{}, func(tx *reform.TX) error {
		session := &dao.UserSession{
			UserID:    userID,
			Token:     token,
			CreatedAt: time.Now().UTC(),
			Active:    true,
		}
		if err := tx.Save(session); err != nil {
			return err
		}
		return nil
	})
}

func (r *r) SessionByAccessToken(ctx Ctx, token string) (*dao.UserSession, error) {
	session := &dao.UserSession{}
	err := r.Tx(ctx, &sql.TxOptions{}, func(tx *sqlx.Tx) error {
		session.User = &dao.User{}
		err := r.DB.QueryRowContext(ctx, sqlSelectActiveUserSession, token).Scan(
			&session.ID,
			&session.UserID,
			&session.User.Status,
		)
		switch {
		case err == sql.ErrNoRows:
			return app.ErrNotFound
		case err != nil:
			return err
		}
		return nil
	})
	return session, err
}

func (r *r) DeleteSession(ctx Ctx, token string) error {
	return r.Tx(ctx, &sql.TxOptions{}, func(tx *sqlx.Tx) error {
		var err error
		_, err = tx.ExecContext(ctx, sqlUpdateUserSession, token)
		return err
	})
}
