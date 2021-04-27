package dal

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/filter"
	"github.com/freonservice/freon/pkg/api"

	"github.com/AlekSi/pointer"
	"github.com/jmoiron/sqlx"
	"gopkg.in/reform.v1"
)

func (r *Repo) CreateIdentifier(
	ctx Ctx, creatorID, categoryID, parentID int64,
	name, description, exampleText, platforms, namedList string,
) error {
	var err error
	return r.ReformDB.InTransactionContext(ctx, &sql.TxOptions{}, func(tx *reform.TX) error {
		identifier := &dao.Identifier{
			CreatorID:   creatorID,
			Name:        name,
			Description: sql.NullString{String: description, Valid: true},
			ExampleText: sql.NullString{String: exampleText, Valid: true},
			Platforms:   platforms,
			NamedList:   sql.NullString{String: namedList, Valid: true},
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   pointer.ToTime(time.Now().UTC()),
		}

		if categoryID > 0 {
			identifier.CategoryID = sql.NullInt64{Int64: categoryID, Valid: true}
		}

		if err = tx.Save(identifier); err != nil {
			if isDuplicateKeyValue(err) {
				return ErrDuplicateKeyValue
			}
			return err
		}

		if parentID > 0 {
			parent, err := r.GetIdentifierByID(tx, parentID) // nolint:govet
			if err != nil {
				return err
			}
			identifier.ParentPath = fmt.Sprintf("%s.%d", parent.ParentPath, identifier.ID)
		} else {
			identifier.ParentPath = fmt.Sprintf("%d", identifier.ID)
		}

		err = tx.Save(identifier)
		if err != nil {
			return err
		}

		localizationIds, err := r.SelectLocalizationListID(ctx, tx)
		if err != nil {
			return err
		}

		for _, id := range localizationIds {
			localizationIdentifier := &dao.LocalizationIdentifier{
				LocalizationID: id,
				IdentifierID:   identifier.ID,
				Status:         int64(api.Status_ACTIVE),
				CreatedAt:      time.Now().UTC(),
			}
			err = tx.Save(localizationIdentifier)
			if err != nil {
				return err
			}

			translation := &dao.Translation{
				LocalizationID: id,
				IdentifierID:   identifier.ID,
				CreatorID:      creatorID,
				Status:         int64(api.TranslationStatus_TRANSLATION_EMPTY),
				CreatedAt:      time.Now().UTC(),
			}
			err = tx.Save(translation)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *Repo) GetIdentifiers(ctx Ctx, f filter.IdentifierFilter) ([]*dao.Identifier, error) {
	rows, err := f.CreateRows(ctx, r.ReformDB)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var entities []*dao.Identifier
	for rows.Next() {
		entity := new(dao.Identifier)
		err = rows.Scan(
			&entity.ID, &entity.Name, &entity.Description,
			&entity.ExampleText, &entity.Platforms, &entity.NamedList, &entity.CategoryID,
		)
		if err != nil {
			break
		}
		if entity.CategoryID.Valid {
			c, err := r.GetCategory(entity.CategoryID.Int64) //nolint:govet
			if err == nil {
				entity.Category = c
			}
		}
		entities = append(entities, entity)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

func (r *Repo) DeleteIdentifier(ctx Ctx, id int64) error {
	return r.Tx(ctx, &sql.TxOptions{}, func(tx *sqlx.Tx) error {
		var err error
		_, err = tx.ExecContext(ctx, sqlDeleteIdentifier, id)
		return err
	})
}

func (r *Repo) SelectLocalizationListID(ctx Ctx, tx *reform.TX) ([]int64, error) {
	rows, err := tx.QueryContext(ctx, sqlSelectLocalizationListID)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
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

func (r *Repo) GetIdentifierByID(tx *reform.TX, id int64) (*dao.Identifier, error) {
	var i dao.Identifier
	err := tx.FindOneTo(&i, "id", id)
	if err != nil {
		return nil, err
	}
	return &i, err
}

func (r *Repo) UpdateIdentifier(
	ctx app.Ctx, id, categoryID, parentID int64,
	name, description, exampleText, platforms, namedList string,
) error {
	var i dao.Identifier
	err := r.ReformDB.FindOneTo(&i, "id", id)
	if err != nil {
		return err
	}

	i.Name = name
	i.Description = sql.NullString{String: description, Valid: true}
	i.ExampleText = sql.NullString{String: exampleText, Valid: true}
	i.UpdatedAt = pointer.ToTime(time.Now().UTC())
	i.Platforms = platforms
	i.NamedList = sql.NullString{String: namedList, Valid: true}
	if categoryID > 0 {
		i.CategoryID = sql.NullInt64{Int64: categoryID, Valid: true}
	}

	if err := r.ReformDB.Save(&i); err != nil {
		if isDuplicateKeyValue(err) {
			return ErrDuplicateKeyValue
		}
		return err
	}
	return nil
}
