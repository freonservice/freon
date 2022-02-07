package dal

import (
	"github.com/freonservice/freon/internal/dao"

	"gopkg.in/reform.v1"
)

func (r *Repo) GetLanguages(ctx Ctx) ([]*dao.Language, error) {
	rows, err := r.ReformDB.SelectRows(dao.LanguageTable, "")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*dao.Language
	for {
		var entity dao.Language
		if err = r.ReformDB.NextRow(&entity, rows); err != nil {
			break
		}
		entities = append(entities, &entity)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	if err != nil && err != reform.ErrNoRows {
		return nil, err
	}
	return entities, nil
}
