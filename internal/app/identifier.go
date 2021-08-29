package app

import (
	"github.com/freonservice/freon/internal/entities"
	"github.com/freonservice/freon/internal/filter"
)

func (a *appl) CreateIdentifier(
	ctx Ctx, creatorID, categoryID, parentID int64, name, description, exampleText string, platforms []string,
) error {
	return a.repo.CreateIdentifier(
		ctx, creatorID, categoryID, parentID, name, description, exampleText, createConcatenatedString(platforms),
	)
}

func (a *appl) GetIdentifiers(ctx Ctx, f filter.IdentifierFilter) ([]*entities.Identifier, error) {
	l, err := a.repo.GetIdentifiers(ctx, f)
	if err != nil {
		return nil, err
	}
	return mappingArrayIdentifier(l), err
}

func (a *appl) DeleteIdentifier(ctx Ctx, id int64) error {
	return a.repo.DeleteIdentifier(ctx, id)
}

func (a *appl) UpdateIdentifier(
	ctx Ctx, id, categoryID, parentID int64, name, description, exampleText string, platforms []string,
) error {
	return a.repo.UpdateIdentifier(
		ctx, id, categoryID, parentID, name, description, exampleText, createConcatenatedString(platforms),
	)
}
