package dao

import (
	"database/sql"
	"time"

	"github.com/AlekSi/pointer"
)

// go:generate reform
// reform:users
type User struct {
	ID         int64          `reform:"id,pk"`
	UUIDID     string         `reform:"uuid_id"`
	Email      string         `reform:"email"`
	Password   string         `reform:"password"`
	FirstName  sql.NullString `reform:"first_name"`
	SecondName sql.NullString `reform:"second_name"`
	Status     int64          `reform:"status"`
	Role       int64          `reform:"role"`
	CreatedAt  time.Time      `reform:"created_at"`
	UpdatedAt  *time.Time     `reform:"updated_at"`
}

func (u *User) BeforeInsert() error {
	if u.UpdatedAt != nil {
		u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second))
	}
	return nil
}

func (u *User) BeforeUpdate() error {
	now := time.Now()
	u.UpdatedAt = &now
	u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second))
	return nil
}

// func (u *User) IsBanned() bool {
//	return u.Status == status.UserBanned
// }
//
// func (u *User) IsActive() bool {
//	return u.Status == status.UserActive
// }
//
// func (u *User) IsNotConfirmed() bool {
//	return u.Status == status.UserNotConfirmed
// }
