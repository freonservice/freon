package dal

import (
	"database/sql"
	"time"

	"github.com/MarcSky/freon/internal/app"
	"github.com/MarcSky/freon/internal/dao"
)

func (r r) CreateCategory(ctx Ctx, name string) error {
	entity := &dao.Category{
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
	if err := r.ReformDB.Save(entity); err != nil {
		if isDuplicateKeyValue(err) {
			return ErrDuplicateKeyValue
		}
		return err
	}
	return nil
}

func (r r) GetCategories(ctx Ctx) ([]*dao.Category, error) {
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
	return entities, nil
}

func (r r) DeleteCategory(ctx Ctx, id int64) error {
	_, err := r.ReformDB.ExecContext(ctx, sqlDeleteCategory, id)
	return err
}

func (r r) UpdateCategory(ctx app.Ctx, id int64, name string) error {
	_, err := r.ReformDB.ExecContext(ctx, sqlUpdateNameCategory, name, id)
	return err
}

func (r r) GetCategory(id int64) (*dao.Category, error) {
	var entity dao.Category
	err := r.ReformDB.FindOneTo(&entity, "id", id)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
