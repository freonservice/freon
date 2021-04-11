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

func (i *IdentifierFilter) CreateRows(ctx context.Context, r *reform.DB) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	if i.CategoryID > 0 {
		rows, err = r.QueryContext(
			ctx,
			"SELECT "+
				"id, name, description, example_text, platforms, named_list, category_id "+
				"FROM identifiers WHERE status=$1 AND category_id=$2 ORDER BY created_at DESC",
			i.Status, i.CategoryID,
		)
	} else {
		rows, err = r.QueryContext(
			ctx,
			"SELECT "+
				"id, name, description, example_text, platforms, named_list, category_id  "+
				"FROM identifiers WHERE status=$1 ORDER BY created_at DESC",
			i.Status,
		)
	}

	return rows, err
}
