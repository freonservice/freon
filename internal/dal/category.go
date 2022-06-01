package dal

import (
	"database/sql"
	"time"

	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/dao"

	"github.com/AlekSi/pointer"
)

func (r *Repo) CreateCategory(ctx Ctx, name string) error {
	entity := &dao.Category{
		Name:      name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: pointer.ToTime(time.Now().UTC()),
	}
	if err := r.ReformDB.Save(entity); err != nil {
		if isDuplicateKeyValue(err) {
			return ErrDuplicateKeyValue
		}
		return err
	}
	return nil
}

func (r *Repo) GetCategories(ctx Ctx) ([]*dao.Category, error) {
	rows, err := r.ReformDB.QueryContext(ctx, sqlSelectCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*dao.Category
	for rows.Next() {
		var entity dao.Category
		err = rows.Scan(&entity.ID, &entity.Name)
		if err != nil {
			break
		}
		entities = append(entities, &entity)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *Repo) DeleteCategory(ctx Ctx, id int64) error {
	_, err := r.ReformDB.ExecContext(ctx, sqlDeleteCategory, id)
	return err
}

func (r *Repo) UpdateCategory(ctx app.Ctx, id int64, name string) error {
	_, err := r.ReformDB.ExecContext(ctx, sqlUpdateNameCategory, name, id)
	return err
}

func (r *Repo) GetCategory(id int64) (*dao.Category, error) {
	var entity dao.Category
	err := r.ReformDB.FindOneTo(&entity, "id", id)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
