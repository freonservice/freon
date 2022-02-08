package filter

import (
	"context"
	"database/sql"

	"gopkg.in/reform.v1"
)

type IdentifierFilter struct {
	Status     int64
	CategoryID int64
}

func (i IdentifierFilter) CreateRows(ctx context.Context, r *reform.DB) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	if i.CategoryID > 0 {
		rows, err = r.QueryContext(
			ctx,
			`
				select id, name, description, text_singular, text_plural, platforms, category_id 
				from identifiers where status=$1 and category_id=$2 order by created_at DESC
			`,
			i.Status, i.CategoryID,
		)
	} else {
		rows, err = r.QueryContext(
			ctx,
			`
				select id, name, description, text_singular, text_plural, platforms, category_id 
				from identifiers where status=$1 order by created_at DESC
			`,
			i.Status,
		)
	}

	return rows, err
}
