package dal

import (
	"database/sql"
	"time"

	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/filter"
	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/AlekSi/pointer"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

func (r *Repo) CreateTranslation(ctx Ctx, creatorID, localizationID, identifierID int64, singular, plural string) error {
	entity := new(dao.Translation)
	return r.ReformDB.InTransactionContext(ctx, &sql.TxOptions{}, func(tx *reform.TX) error {
		entity = &dao.Translation{
			CreatorID:      creatorID,
			LocalizationID: localizationID,
			IdentifierID:   identifierID,
			Singular:       singular,
			Plural:         sql.NullString{String: plural, Valid: true},
			CreatedAt:      time.Now().UTC(),
			UpdatedAt:      pointer.ToTime(time.Now().UTC()),
			Status:         int64(api.Status_NOT_ACTIVE),
		}
		if err := tx.Save(entity); err != nil {
			return err
		}
		return nil
	})
}

func (r *Repo) GetTranslations(ctx Ctx, f filter.TranslationFilter) ([]*dao.Translation, error) {
	rows, err := f.CreateRows(ctx, r.ReformDB)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var entities []*dao.Translation
	for rows.Next() {
		entity := new(dao.Translation)
		entity.Localization = new(dao.Localization)
		entity.Identifier = new(dao.Identifier)
		err = rows.Scan(
			&entity.ID, &entity.Singular, &entity.Plural, &entity.Status, &entity.CreatedAt,
			&entity.Localization.ID, &entity.Localization.Locale, &entity.Localization.LanguageName,
			&entity.Identifier.ID, &entity.Identifier.Name, &entity.Identifier.Description,
			&entity.Identifier.ExampleText, &entity.Identifier.Platforms,
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

func (r *Repo) UpdateTranslation(ctx Ctx, id int64, singular, plural string) error {
	var t dao.Translation
	err := r.ReformDB.FindOneTo(&t, "id", id)
	if err != nil {
		return err
	}

	if singular == "" {
		t.Status = int64(api.StatusTranslation_HIDDEN)
	} else {
		t.Singular = singular
		t.Status = int64(api.StatusTranslation_DRAFT)
	}

	if plural != "" {
		t.Plural = sql.NullString{String: plural, Valid: true}
	}
	t.UpdatedAt = pointer.ToTime(time.Now().UTC())

	return r.ReformDB.Save(&t)
}

func (r *Repo) DeleteTranslation(ctx Ctx, id int64) error {
	return r.Tx(ctx, &sql.TxOptions{}, func(tx *sqlx.Tx) error {
		var err error
		_, err = tx.ExecContext(ctx, sqlDeleteTranslation, id)
		return err
	})
}

func (r *Repo) UpdateStatusTranslation(ctx Ctx, id, status int64) error {
	return r.Tx(ctx, &sql.TxOptions{}, func(tx *sqlx.Tx) error {
		var err error
		_, err = tx.ExecContext(ctx, sqlUpdateStatusTranslation, status, id)
		return err
	})
}

func (r *Repo) GetTranslation(ctx app.Ctx, locale, identifierName string) (*dao.Translation, error) {
	entity := new(dao.Translation)

	row := r.DB.QueryRowContext(ctx, sqlSelectTranslation, locale, identifierName)
	err := row.Scan(&entity.Singular, &entity.Plural)
	switch errors.Cause(err) {
	default:
		return nil, err
	case sql.ErrNoRows:
		return nil, app.ErrNotFound
	case nil:
	}

	return entity, nil
}

func (r *Repo) GetGroupedTranslations(ctx Ctx, f filter.GroupedTranslationFilter) (map[string][]*dao.Translation, error) {
	rows, err := f.CreateRows(ctx, r.DB)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	var gts = make(map[string][]*dao.Translation)
	for rows.Next() {
		entity := new(dao.Translation)
		entity.Localization = new(dao.Localization)
		entity.Identifier = new(dao.Identifier)
		err = rows.Scan(
			&entity.ID, &entity.Singular, &entity.Plural, &entity.Status, &entity.CreatedAt,
			&entity.Localization.ID, &entity.Localization.Locale, &entity.Localization.LanguageName,
			&entity.Identifier.ID, &entity.Identifier.Name, &entity.Identifier.Description,
			&entity.Identifier.ExampleText, &entity.Identifier.Platforms,
		)
		if err != nil {
			break
		}
		gts[entity.Localization.Locale] = append(gts[entity.Localization.Locale], entity)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return gts, nil
}
