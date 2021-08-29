package app

import "github.com/freonservice/freon/internal/entities"

func (a *appl) CreateCategory(ctx Ctx, name string) error {
	return a.repo.CreateCategory(ctx, name)
}

func (a *appl) GetCategories(ctx Ctx) ([]*entities.Category, error) {
	c, err := a.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return mappingArrayCategory(c), err
}

func (a *appl) DeleteCategory(ctx Ctx, id int64) error {
	return a.repo.DeleteCategory(ctx, id)
}

func (a *appl) UpdateCategory(ctx Ctx, id int64, name string) error {
	return a.repo.UpdateCategory(ctx, id, name)
}
