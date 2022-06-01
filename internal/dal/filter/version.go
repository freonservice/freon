package filter

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/freonservice/freon/pkg/freonApi"

	"gopkg.in/reform.v1"
)

type VersionTranslationFilesFilter struct {
	LocalizationID int64
	PlatformType   int64
}

type VersionTranslationsFilter struct {
	LocalizationID int64
}

func (t VersionTranslationFilesFilter) CreateRows(ctx context.Context, r *reform.DB) (*sql.Rows, error) {
	var (
		err      error
		rows     *sql.Rows
		buildSQL string
	)

	const (
		selectFromSQL = `
		select t1.path, t1.platform, t1.localization_id, t1.updated_at, l.locale, l.lang_name
			from translation_files t1
			left join localizations l on l.id = t1.localization_id`
		subSQL = `
		(select max(t2.updated_at)
		   from translation_files t2
		   where t2.platform = t1.platform
			and t2.localization_id = t1.localization_id
			and t2.status = t1.status)`
	)
	if t.LocalizationID > 0 && t.PlatformType > 0 {
		buildSQL = fmt.Sprintf(`%s where t1.localization_id = $1 and t1.platform = $2 and t1.status = $3 and t1.updated_at = %s`, selectFromSQL, subSQL) // nolint:lll
		rows, err = r.QueryContext(ctx, buildSQL, t.PlatformType, t.LocalizationID, freonApi.Status_ACTIVE)
	} else if t.LocalizationID > 0 {
		buildSQL = fmt.Sprintf("%s where t1.localization_id = $1 and t1.status = $2 and t1.updated_at = %s", selectFromSQL, subSQL)
		rows, err = r.QueryContext(ctx, buildSQL, t.LocalizationID, freonApi.Status_ACTIVE)
	} else if t.PlatformType > 0 {
		buildSQL = fmt.Sprintf(`%s where t1.platform = $1 and t1.status = $2 and t1.updated_at = %s`, selectFromSQL, subSQL)
		rows, err = r.QueryContext(ctx, buildSQL, t.PlatformType, freonApi.Status_ACTIVE)
	} else {
		buildSQL = fmt.Sprintf("%s where t1.status = $1 and t1.updated_at = %s", selectFromSQL, subSQL)
		rows, err = r.QueryContext(ctx, buildSQL, freonApi.Status_ACTIVE)
	}

	return rows, err
}

func (t VersionTranslationsFilter) CreateRows(ctx context.Context, r *reform.DB) (*sql.Rows, error) {
	var (
		err      error
		rows     *sql.Rows
		buildSQL string
	)

	const (
		selectFromSQL = `
		select t1.localization_id, t1.updated_at, l.locale, l.lang_name
			from translations t1
			left join localizations l on l.id = t1.localization_id`
		subSQL = `
		(select max(t2.updated_at)
		   from translations t2
		   where t2.localization_id = t1.localization_id
			and t2.status = t1.status)`
	)
	if t.LocalizationID > 0 {
		buildSQL = fmt.Sprintf("%s where t1.localization_id = $1 and t1.status = $2 and t1.updated_at = %s", selectFromSQL, subSQL)
		rows, err = r.QueryContext(ctx, buildSQL, t.LocalizationID, freonApi.Status_ACTIVE)
	} else {
		buildSQL = fmt.Sprintf("%s where t1.status = $1 and t1.updated_at = %s", selectFromSQL, subSQL)
		rows, err = r.QueryContext(ctx, buildSQL, freonApi.Status_ACTIVE)
	}

	return rows, err
}
