package dal

import (
	"database/sql"
	"time"

	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/pkg/api"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"gopkg.in/reform.v1"
)

func (r *r) CreateUser(ctx Ctx, email, password, firstName, secondName string, role int64) (*dao.User, error) {
	var err error
	user := new(dao.User)
	err = r.ReformDB.InTransactionContext(ctx, &sql.TxOptions{}, func(tx *reform.TX) error {
		user = &dao.User{
			UUIDID:     uuid.New().String(),
			Email:      email,
			Password:   password,
			FirstName:  sql.NullString{String: firstName, Valid: true},
			SecondName: sql.NullString{String: secondName, Valid: true},
			Role:       role,
			Status:     int64(api.UserStatus_USER_ACTIVE),
			CreatedAt:  time.Now().UTC(),
		}
		err = tx.Save(user)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

func (r *r) UpdatePassword(ctx app.Ctx, userID int64, passwordHash string) error {
	return r.ReformDB.InTransactionContext(ctx, &sql.TxOptions{}, func(tx *reform.TX) error {
		user, err := r.GetUserByID(userID)
		if err != nil {
			return err
		}
		user.Password = passwordHash
		return tx.Save(user)
	})
}

func (r *r) UpdateProfile(ctx app.Ctx, userID int64, email, firstName, secondName string, role, status int64) error {
	user, err := r.GetUserByID(userID)
	if err != nil {
		return err
	}
	user.Email = email
	user.FirstName = sql.NullString{String: firstName, Valid: true}
	user.SecondName = sql.NullString{String: secondName, Valid: true}
	user.Role = role
	user.Status = status
	user.UpdatedAt = pointer.ToTime(time.Now().UTC())
	return r.ReformDB.Save(user)
}

func (r *r) GetUserByUserUUID(userUUID string) (*dao.User, error) {
	var user dao.User
	err := r.ReformDB.FindOneTo(&user, "uuid_id", userUUID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *r) GetUserByEmail(email string) (*dao.User, error) {
	var user dao.User
	err := r.ReformDB.FindOneTo(&user, "email", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *r) GetUserByID(id int64) (*dao.User, error) {
	var user dao.User
	err := r.ReformDB.FindOneTo(&user, "id", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *r) GetUsers(ctx Ctx) ([]*dao.User, error) {
	rows, err := r.DB.QueryContext(ctx, sqlSelectUsers)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*dao.User
	for rows.Next() {
		entity := new(dao.User)
		err = rows.Scan(
			&entity.ID, &entity.UUIDID, &entity.Email,
			&entity.FirstName, &entity.SecondName, &entity.Status,
			&entity.Role, &entity.CreatedAt, &entity.UpdatedAt,
		)
		if err != nil {
			break
		}
		entities = append(entities, entity)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

func (r *r) UpdateStatus(ctx app.Ctx, userID, status int64) error {
	user, err := r.GetUserByID(userID)
	if err != nil {
		return err
	}
	user.Status = status
	user.UpdatedAt = pointer.ToTime(time.Now().UTC())
	return r.ReformDB.Save(user)
}
