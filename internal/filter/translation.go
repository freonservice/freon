package filter

import (
	"context"
	"database/sql"

	"gopkg.in/reform.v1"
)

type TranslationFilter struct {
	LocalizationID int64
	Locale         string
}

func (t *TranslationFilter) CreateRows(ctx context.Context, r *reform.DB) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	if t.LocalizationID > 0 {
		rows, err = r.QueryContext(
			ctx,
			`SELECT  
			t.id, t.text, t.status, t.created_at, l.id, l.locale, l.lang_name, 
			i.id, i.name, i.description, i.example_text, i.platforms, i.named_list 
			FROM translations AS t 
			LEFT JOIN localizations l ON t.localization_id=l.id 
			LEFT JOIN identifiers i ON t.identifier_id=i.id 
			WHERE t.localization_id = $1
			ORDER BY t.id DESC`,
			t.LocalizationID,
		)
	} else if t.Locale != "" {
		rows, err = r.QueryContext(
			ctx,
			`SELECT 
			t.id, t.text, t.status, t.created_at, l.id, l.locale, l.lang_name, 
			i.id, i.name, i.description, i.example_text, i.platforms, i.named_list 
			FROM translations AS t 
			LEFT JOIN localizations l ON t.localization_id=l.id 
			LEFT JOIN identifiers i ON t.identifier_id=i.id 
			WHERE l.locale = $1
			ORDER BY t.id DESC`,
			t.Locale,
		)
	} else {
		rows, err = r.QueryContext(
			ctx,
			`SELECT  
			t.id, t.text, t.status, t.created_at, l.id, l.locale, l.lang_name, 
			i.id, i.name, i.description, i.example_text, i.platforms, i.named_list 
			FROM translations AS t
			LEFT JOIN localizations l ON l.id=t.localization_id 
			LEFT JOIN identifiers i ON i.id=t.identifier_id 
			ORDER BY t.id DESC`,
		)
	}

	return rows, err
}
