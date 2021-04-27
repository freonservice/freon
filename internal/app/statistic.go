package app

func (a *appl) GetStatistic(ctx Ctx) (*Statistic, error) {
	s, err := a.repo.GetStatistic(ctx)
	if err != nil {
		return nil, err
	}

	stat := mappingStatistic(s)
	sizeStatTranslations := len(s.StatTranslations)
	stat.StatTranslations = make([]*StatTranslation, sizeStatTranslations)
	for i := 0; i < sizeStatTranslations; i++ {
		stat.StatTranslations[i] = &StatTranslation{
			LangName:   s.StatTranslations[i].LangName,
			Percentage: float64(s.StatTranslations[i].Fulfilled) / float64(s.CountIdentifiers) * 100,
		}
	}

	return stat, nil
}
