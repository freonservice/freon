package dal

import (
	"database/sql"
	"time"

	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/pkg/api"

	"github.com/jmoiron/sqlx"
	"gopkg.in/reform.v1"
)

func (r *Repo) CreateLocalization(ctx Ctx, creatorID int64, locale, languageName, icon string) (*dao.Localization, error) {
	var err error
	entity := new(dao.Localization)
	err = r.ReformDB.InTransactionContext(ctx, &sql.TxOptions{}, func(tx *reform.TX) error {
		entity = &dao.Localization{
			CreatorID:    creatorID,
			Locale:       locale,
			LanguageName: languageName,
			Icon:         icon,
			CreatedAt:    time.Now().UTC(),
		}
		err = tx.Save(entity)
		if err != nil {
			if isDuplicateKeyValue(err) {
				return ErrDuplicateKeyValue
			}
			return err
		}

		identifierIds, err := r.SelectIdentifierListID(ctx, tx) //nolint:govet
		if err != nil {
			return err
		}

		for _, id := range identifierIds {
			localizationIdentifier := &dao.LocalizationIdentifier{
				LocalizationID: entity.ID,
				IdentifierID:   id,
				Status:         int64(api.LocalizationIdentifierStatus_LOCALIZATION_IDENTIFIER_ACTIVE),
				CreatedAt:      time.Now().UTC(),
			}
			err = tx.Save(localizationIdentifier)
			if err != nil {
				return err
			}

			translation := &dao.Translation{
				LocalizationID: entity.ID,
				IdentifierID:   id,
				CreatorID:      creatorID,
				Status:         int64(api.TranslationStatus_TRANSLATION_ACTIVE),
				CreatedAt:      time.Now().UTC(),
			}
			err = tx.Save(translation)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return entity, err
}

func (r *Repo) GetLocalizations(ctx Ctx) ([]*dao.Localization, error) {
	rows, err := r.ReformDB.SelectRows(
		dao.LocalizationTable, "WHERE status = $1 ORDER BY id DESC", api.LocalizationStatus_LOCALIZATION_ACTIVE,
	)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*dao.Localization
	for {
		var entity dao.Localization
		if err = r.ReformDB.NextRow(&entity, rows); err != nil {
			break
		}
		entities = append(entities, &entity)
	}
	if err != reform.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

func (r *Repo) DeleteLocalization(ctx Ctx, id int64) error {
	return r.Tx(ctx, &sql.TxOptions{}, func(tx *sqlx.Tx) error {
		var err error
		_, err = tx.ExecContext(ctx, sqlDeleteLocalization, id)
		return err
	})
}

func (r *Repo) SelectIdentifierListID(ctx Ctx, tx *reform.TX) ([]int64, error) {
	rows, err := tx.QueryContext(ctx, sqlSelectIdentifierListID, api.IdentifierStatus_IDENTIFIER_ACTIVE)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err = rows.Scan(&id); err != nil {
			break
		}
		ids = append(ids, id)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return ids, nil
}
