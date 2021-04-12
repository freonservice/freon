package dao

import (
	"time"

	"github.com/AlekSi/pointer"
)

// go:generate reform
// reform:categories
type Category struct {
	ID        int64      `reform:"id,pk"`
	Name      string     `reform:"name"`
	CreatedAt time.Time  `reform:"created_at"`
	UpdatedAt *time.Time `reform:"updated_at"`
}

func (l *Category) BeforeInsert() error {
	if l.UpdatedAt != nil {
		l.UpdatedAt = pointer.ToTime(l.UpdatedAt.UTC().Truncate(time.Second))
	}
	return nil
}

func (l *Category) BeforeUpdate() error {
	now := time.Now().UTC()
	l.UpdatedAt = &now
	l.UpdatedAt = pointer.ToTime(l.UpdatedAt.UTC().Truncate(time.Second))
	return nil
}
