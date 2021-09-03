package dal

import (
	"database/sql"
	"time"

	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/filter"
	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/AlekSi/pointer"
)

func (r *Repo) CreateTranslationFile(ctx Ctx, name, path string, platform, storageType, creatorID, localizationID int64) error {
	entity := &dao.TranslationFile{
		LocalizationID: localizationID,
		CreatorID:      creatorID,
		Name:           name,
		Path:           path,
		Platform:       platform,
		Status:         int64(api.Status_ACTIVE),
		StorageType:    storageType,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      pointer.ToTime(time.Now().UTC()),
	}
	if err := r.ReformDB.Save(entity); err != nil {
		if isDuplicateKeyValue(err) {
			return ErrDuplicateKeyValue
		}
		return err
	}
	return nil
}

func (r *Repo) GetTranslationFile(ctx Ctx, id int64) (*dao.TranslationFile, error) {
	var entity dao.TranslationFile
	err := r.ReformDB.FindOneTo(&entity, "id", id)
	if err != nil {
		return nil, err
	}
	return &entity, err
}

func (r *Repo) GetTranslationFiles(ctx Ctx, f filter.TranslationFileFilter) ([]*dao.TranslationFile, error) {
	rows, err := f.CreateRows(ctx, r.ReformDB)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var entities []*dao.TranslationFile
	for rows.Next() {
		entity := new(dao.TranslationFile)
		entity.Localization = new(dao.Localization)
		err = rows.Scan(
			&entity.ID, &entity.Name, &entity.Path, &entity.Platform, &entity.StorageType, &entity.CreatedAt,
			&entity.Localization.ID, &entity.Localization.Locale, &entity.Localization.LanguageName,
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

func (r *Repo) DeleteTranslationFile(ctx Ctx, id int64) error {
	_, err := r.ReformDB.ExecContext(ctx, sqlDeleteTranslationFile, id)
	return err
}
