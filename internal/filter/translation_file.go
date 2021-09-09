package filter

import (
	"context"
	"database/sql"

	"gopkg.in/reform.v1"
)

type TranslationFileFilter struct {
	LocalizationID int64
	PlatformType   int64
}

func (t *TranslationFileFilter) CreateRows(ctx context.Context, r *reform.DB) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	if t.LocalizationID > 0 && t.PlatformType > 0 {
		rows, err = r.QueryContext(
			ctx,
			`SELECT  
			tf.id, tf.name, tf.path, tf.platform, tf.storage_type, tf.created_at, tf.updated_at, l.id, l.locale, l.lang_name  
			FROM translation_files AS tf 
			LEFT JOIN localizations l ON tf.localization_id=l.id 
			WHERE tf.platform = $1 AND tf.localization_id = $2
			ORDER BY tf.id DESC`,
			t.PlatformType, t.LocalizationID,
		)
	} else if t.LocalizationID > 0 {
		rows, err = r.QueryContext(
			ctx,
			`SELECT
			tf.id, tf.name, tf.path, tf.platform, tf.storage_type, tf.created_at, tf.updated_at, l.id, l.locale, l.lang_name  
			FROM translation_files AS tf
			LEFT JOIN localizations l ON tf.localization_id=l.id 
			WHERE tf.localization_id = $1
			ORDER BY tf.id DESC`,
			t.LocalizationID,
		)
	} else if t.PlatformType > 0 {
		rows, err = r.QueryContext(
			ctx,
			`SELECT  
			tf.id, tf.name, tf.path, tf.platform, tf.storage_type, tf.created_at, tf.updated_at, l.id, l.locale, l.lang_name  
			FROM translation_files AS tf 
			LEFT JOIN localizations l ON tf.localization_id=l.id 
			WHERE tf.platform = $1
			ORDER BY tf.id DESC`,
			t.PlatformType,
		)
	} else {
		rows, err = r.QueryContext(
			ctx,
			`SELECT  
			tf.id, tf.name, tf.path, tf.platform, tf.storage_type, tf.created_at, tf.updated_at, l.id, l.locale, l.lang_name  
			FROM translation_files AS tf 
			LEFT JOIN localizations l ON tf.localization_id=l.id 
			ORDER BY tf.id DESC`,
		)
	}

	return rows, err
}
