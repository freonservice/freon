package dal

import (
	"database/sql"
	"time"

	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/pkg/api"

	"github.com/AlekSi/pointer"
	"github.com/jmoiron/sqlx"
	"gopkg.in/reform.v1"
)

func (r *r) CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, text string) error {
	entity := new(dao.Translation)
	return r.ReformDB.InTransactionContext(ctx, &sql.TxOptions{}, func(tx *reform.TX) error {
		entity = &dao.Translation{
			CreatorID:      creatorID,
			LocalizationID: localizationID,
			IdentifierID:   identifierID,
			Text:           text,
			CreatedAt:      time.Now().UTC(),
		}
		if err := tx.Save(entity); err != nil {
			return err
		}
		return nil
	})
}

func (r *r) GetTranslations(ctx Ctx, f dao.TranslationFilter) ([]*dao.Translation, error) {
	rows, err := f.CreateRows(ctx, r.ReformDB)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*dao.Translation
	for rows.Next() {
		entity := new(dao.Translation)
		entity.Localization = new(dao.Localization)
		entity.Identifier = new(dao.Identifier)
		err = rows.Scan(
			&entity.ID, &entity.Text, &entity.Status, &entity.CreatedAt,
			&entity.Localization.ID, &entity.Localization.Locale, &entity.Localization.LanguageName,
			&entity.Identifier.ID, &entity.Identifier.Name, &entity.Identifier.Description,
			&entity.Identifier.ExampleText, &entity.Identifier.Platforms, &entity.Identifier.NamedList,
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

func (r *r) UpdateTranslation(ctx Ctx, id int64, text string) error {
	var t dao.Translation
	err := r.ReformDB.FindOneTo(&t, "id", id)
	if err != nil {
		return err
	}

	t.Text = text
	t.UpdatedAt = pointer.ToTime(time.Now().UTC())

	return r.ReformDB.Save(&t)
}

func (r *r) DeleteTranslation(ctx Ctx, id int64) error {
	return r.Tx(ctx, &sql.TxOptions{}, func(tx *sqlx.Tx) error {
		var err error
		_, err = tx.ExecContext(ctx, sqlDeleteTranslation, id)
		return err
	})
}

func (r *r) UpdateHideStatusTranslation(ctx Ctx, id int64, hide bool) error {
	status := api.TranslationStatus_TRANSLATION_HIDDEN
	if !hide {
		status = api.TranslationStatus_TRANSLATION_ACTIVE
	}
	return r.Tx(ctx, &sql.TxOptions{}, func(tx *sqlx.Tx) error {
		var err error
		_, err = tx.ExecContext(ctx, sqlUpdateHideStatusTranslation, status, id)
		return err
	})
}
