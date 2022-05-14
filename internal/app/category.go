package app

import "github.com/freonservice/freon/internal/domain"

func (a *appl) CreateCategory(ctx Ctx, name string) error {
	return a.svc.repo.CreateCategory(ctx, name)
}

func (a *appl) GetCategories(ctx Ctx) ([]*domain.Category, error) {
	c, err := a.svc.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return mappingArrayCategory(c), err
}

func (a *appl) DeleteCategory(ctx Ctx, id int64) error {
	return a.svc.repo.DeleteCategory(ctx, id)
}

func (a *appl) UpdateCategory(ctx Ctx, id int64, name string) error {
	return a.svc.repo.UpdateCategory(ctx, id, name)
}
