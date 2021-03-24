package dao

import (
	"time"

	"github.com/AlekSi/pointer"
)

// go:generate reform
// reform:localizations
type Localization struct {
	ID           int64      `reform:"id,pk"`
	CreatorID    int64      `reform:"creator_id"`
	Locale       string     `reform:"locale"`
	Icon         string     `reform:"icon"`
	LanguageName string     `reform:"lang_name"`
	Status       int64      `reform:"status"`
	CreatedAt    time.Time  `reform:"created_at"`
	UpdatedAt    *time.Time `reform:"updated_at"`

	Creator *User
}

func (u *Localization) BeforeInsert() error {
	if u.UpdatedAt != nil {
		u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second).AddDate(0, 0, 0))
	}
	return nil
}

func (u *Localization) BeforeUpdate() error {
	now := time.Now()
	u.UpdatedAt = &now
	u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second).AddDate(0, 0, 0))
	return nil
}
