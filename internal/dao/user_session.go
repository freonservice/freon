package dao

import (
	"time"
)

// go:generate reform
// reform:user_sessions
type UserSession struct {
	ID        int64     `reform:"id,pk"`
	UserID    int64     `reform:"user_id"`
	Token     string    `reform:"token"`
	Active    bool      `reform:"active"`
	CreatedAt time.Time `reform:"created_at"`

	User *User
}
