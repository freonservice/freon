package dao

import (
	"time"

	"github.com/AlekSi/pointer"
)

// go:generate reform
// reform:localization_identifiers
type LocalizationIdentifier struct {
	ID             int64      `reform:"id,pk"`
	LocalizationID int64      `reform:"localization_id"`
	IdentifierID   int64      `reform:"identifier_id"`
	Status         int64      `reform:"status"`
	CreatedAt      time.Time  `reform:"created_at"`
	UpdatedAt      *time.Time `reform:"updated_at"`

	Localization *Localization
	Identifier   *Identifier
}

func (l *LocalizationIdentifier) BeforeInsert() error {
	if l.UpdatedAt != nil {
		l.UpdatedAt = pointer.ToTime(l.UpdatedAt.UTC().Truncate(time.Second))
	}
	return nil
}

func (l *LocalizationIdentifier) BeforeUpdate() error {
	now := time.Now()
	l.UpdatedAt = &now
	l.UpdatedAt = pointer.ToTime(l.UpdatedAt.UTC().Truncate(time.Second))
	return nil
}
