package dao

import (
	"database/sql"
	"time"

	"github.com/AlekSi/pointer"
)

// go:generate reform
// reform:identifiers
type Identifier struct {
	ID           int64          `reform:"id,pk"`
	Name         string         `reform:"name"`
	ParentPath   string         `reform:"parent_path"`
	Description  sql.NullString `reform:"description"`
	TextSingular sql.NullString `reform:"text_singular"`
	TextPlural   sql.NullString `reform:"text_plural"`
	CreatorID    int64          `reform:"creator_id"`
	CategoryID   sql.NullInt64  `reform:"category_id"`
	Status       int64          `reform:"status"`
	Platforms    string         `reform:"platforms"` // example: web,ios,android
	CreatedAt    time.Time      `reform:"created_at"`
	UpdatedAt    *time.Time     `reform:"updated_at"`

	Creator  *User
	Category *Category
}

func (l *Identifier) BeforeInsert() error {
	if l.UpdatedAt != nil {
		l.UpdatedAt = pointer.ToTime(l.UpdatedAt.UTC().Truncate(time.Second))
	}
	return nil
}

func (l *Identifier) BeforeUpdate() error {
	now := time.Now().UTC()
	l.UpdatedAt = &now
	l.UpdatedAt = pointer.ToTime(l.UpdatedAt.UTC().Truncate(time.Second))
	return nil
}
