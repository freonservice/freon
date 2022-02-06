package app

import "github.com/freonservice/freon/internal/domain"

func (a *appl) CreateLocalization(ctx Ctx, creatorID int64, name, code string) error {
	_, err := a.repo.CreateLocalization(ctx, creatorID, name, code)
	return err
}

func (a *appl) GetLocalizations(ctx Ctx) ([]*domain.Localization, error) {
	res, err := a.repo.GetLocalizations(ctx)
	if err != nil {
		return nil, err
	}
	return mappingArrayLocalization(res), nil
}

func (a *appl) DeleteLocalization(ctx Ctx, id int64) error {
	return a.repo.DeleteLocalization(ctx, id)
}
