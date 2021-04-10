package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlekSi/pointer"
	"gopkg.in/reform.v1"
)

// go:generate reform
// reform:translations
type Translation struct {
	ID             int64      `reform:"id,pk"`
	LocalizationID int64      `reform:"localization_id"`
	IdentifierID   int64      `reform:"identifier_id"`
	CreatorID      int64      `reform:"creator_id"`
	Text           string     `reform:"text"`
	Status         int64      `reform:"status"`
	CreatedAt      time.Time  `reform:"created_at"`
	UpdatedAt      *time.Time `reform:"updated_at"`

	Localization *Localization
	Identifier   *Identifier
	Creator      *User
}

func (u *Translation) BeforeInsert() error {
	if u.UpdatedAt != nil {
		u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second).AddDate(0, 0, 0))
	}
	return nil
}

func (u *Translation) BeforeUpdate() error {
	now := time.Now().UTC()
	u.UpdatedAt = &now
	u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second).AddDate(0, 0, 0))
	return nil
}

type TranslationFilter struct {
	Status         int64
	LocalizationID int64
}

func (t *TranslationFilter) CreateRows(ctx context.Context, r *reform.DB) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	if t.LocalizationID > 0 {
		rows, err = r.QueryContext(
			ctx,
			`SELECT 
			t.id, t.text, t.status, t.created_at,
			l.id, l.locale, l.lang_name, 
			i.id, i.name, i.description, i.example_text, i.platforms, i.named_list 
			FROM translations AS t 
			LEFT JOIN localizations l ON t.localization_id=l.id 
			LEFT JOIN identifiers i ON t.identifier_id=i.id 
			WHERE t.localization_id = $1
			ORDER BY t.id DESC`,
			t.LocalizationID,
		)
	} else {
		rows, err = r.QueryContext(
			ctx,
			`SELECT 
			t.id, t.text, t.status, t.created_at,
			l.id, l.locale, l.lang_name, 
			i.id, i.name, i.description, i.example_text, i.platforms, i.named_list 
			FROM translations AS t
			LEFT JOIN localizations l ON l.id=t.localization_id 
			LEFT JOIN identifiers i ON i.id=t.identifier_id 
			ORDER BY t.id DESC`,
		)
	}

	return rows, err
}
