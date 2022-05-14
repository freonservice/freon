package app

import "github.com/freonservice/freon/internal/domain"

func (a *appl) GetStatistic(ctx Ctx) (*domain.Statistic, error) {
	s, err := a.svc.repo.GetStatistic(ctx)
	if err != nil {
		return nil, err
	}

	stat := mappingStatistic(s)
	sizeStatTranslations := len(s.StatTranslations)
	stat.StatTranslations = make([]*domain.StatTranslation, sizeStatTranslations)
	for i := 0; i < sizeStatTranslations; i++ {
		stat.StatTranslations[i] = &domain.StatTranslation{
			LangName:   s.StatTranslations[i].LangName,
			Percentage: float64(s.StatTranslations[i].Fulfilled) / float64(s.CountIdentifiers) * 100,
		}
	}

	return stat, nil
}
