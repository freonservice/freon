package dal

import (
	"database/sql"

	"github.com/freonservice/freon/internal/dao"
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
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entities, nil
}
