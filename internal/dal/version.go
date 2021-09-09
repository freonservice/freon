package dal

import (
	"database/sql"

	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/filter"
)

func (r *Repo) GetVersionFromTranslationFiles(ctx Ctx, f filter.VersionTranslationFilesFilter) ([]*dao.Version, error) {
	rows, err := f.CreateRows(ctx, r.ReformDB)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var (
		entities []*dao.Version
	)
	for rows.Next() {
		entity := new(dao.Version)
		err = rows.Scan(
			&entity.Path, &entity.Platform,
			&entity.LocalizationID, &entity.UpdatedAt,
			&entity.Locale, &entity.LangName,
		)
		if err != nil {
			break
		}
		entities = append(entities, entity)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

func (r *Repo) GetVersionFromTranslations(ctx Ctx, f filter.VersionTranslationsFilter) ([]*dao.Version, error) {
	rows, err := f.CreateRows(ctx, r.ReformDB)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var entities []*dao.Version
	for rows.Next() {
		entity := new(dao.Version)
		err = rows.Scan(
			&entity.LocalizationID, &entity.UpdatedAt,
			&entity.Locale, &entity.LangName,
		)
		if err != nil {
			break
		}
		entities = append(entities, entity)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}
