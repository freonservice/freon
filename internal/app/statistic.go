package app

import "github.com/freonservice/freon/internal/entities"

func (a *appl) GetStatistic(ctx Ctx) (*entities.Statistic, error) {
	s, err := a.repo.GetStatistic(ctx)
	if err != nil {
		return nil, err
	}

	stat := mappingStatistic(s)
	sizeStatTranslations := len(s.StatTranslations)
	stat.StatTranslations = make([]*entities.StatTranslation, sizeStatTranslations)
	for i := 0; i < sizeStatTranslations; i++ {
		stat.StatTranslations[i] = &entities.StatTranslation{
			LangName:   s.StatTranslations[i].LangName,
			Percentage: float64(s.StatTranslations[i].Fulfilled) / float64(s.CountIdentifiers) * 100,
		}
	}

	return stat, nil
}
