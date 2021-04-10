package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/AlekSi/pointer"
	"gopkg.in/reform.v1"
)

// go:generate reform
// reform:identifiers
type Identifier struct {
	ID          int64          `reform:"id,pk"`
	ParentPath  string         `reform:"parent_path"`
	Name        string         `reform:"name"`
	Description sql.NullString `reform:"description"`
	ExampleText sql.NullString `reform:"example_text"`
	CreatorID   int64          `reform:"creator_id"`
	CategoryID  sql.NullInt64  `reform:"category_id"`
	Status      int64          `reform:"status"`
	Platforms   string         `reform:"platforms"` // example: web,ios,android
	NamedList   sql.NullString `reform:"named_list"`
	CreatedAt   time.Time      `reform:"created_at"`
	UpdatedAt   *time.Time     `reform:"updated_at"`

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
			"SELECT id, name, description, example_text, platforms, named_list, category_id "+
				"FROM identifiers WHERE status=$1 AND category_id=$2 ORDER BY created_at DESC",
			i.Status, i.CategoryID,
		)
	} else {
		rows, err = r.QueryContext(
			ctx,
			"SELECT id, name, description, example_text, platforms, named_list, category_id "+
				"FROM identifiers WHERE status=$1 ORDER BY created_at DESC",
			i.Status,
		)
	}

	return rows, err
}
