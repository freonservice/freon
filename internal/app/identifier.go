package app

import (
	"fmt"
	"sort"

	"github.com/freonservice/freon/internal/filter"
)

func (a *appl) CreateIdentifier(
	ctx Ctx, creatorID, categoryID, parentID int64, name, description, exampleText string,
	platforms, namedList []string,
) error {
	return a.repo.CreateIdentifier(
		ctx, creatorID, categoryID, parentID, name, description, exampleText,
		createConcatenatedString(platforms), createConcatenatedString(namedList),
	)
}

func (a *appl) GetIdentifiers(ctx Ctx, f filter.IdentifierFilter) ([]*Identifier, error) {
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
	ctx Ctx, id, categoryID, parentID int64, name, description, exampleText string,
	platforms, namedList []string,
) error {
	return a.repo.UpdateIdentifier(
		ctx, id, categoryID, parentID, name, description, exampleText,
		createConcatenatedString(platforms), createConcatenatedString(namedList),
	)
}

func uniqueStringSlice(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func createConcatenatedString(data []string) string {
	var result string
	if len(data) > 0 {
		data = uniqueStringSlice(data)
		sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
		for i, v := range data {
			if i > 0 {
				result = fmt.Sprintf("%s,%s", result, v)
			} else {
				result = v
			}
		}
	}
	return result
}
