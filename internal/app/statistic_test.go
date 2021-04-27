package app

import (
	"context"
	"testing"

	"github.com/freonservice/freon/internal/dao"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetStatistic(t *testing.T) {
	finish, appl, mockRepo, _, _ := testNew(t)
	defer finish()

	{
		var (
			percentages = []float64{25, 50, 100, 0}
		)
		mockRepo.EXPECT().GetStatistic(gomock.Any()).Return(&dao.Statistic{
			CountIdentifiers: 20,
			StatTranslations: []*dao.StatTranslation{
				{
					LangName:  "ka-GE",
					Fulfilled: 5,
				}, {
					LangName:  "ru-RU",
					Fulfilled: 10,
				}, {
					LangName:  "de-DE",
					Fulfilled: 20,
				}, {
					LangName:  "fr-FR",
					Fulfilled: 0,
				},
			},
		}, nil)
		statistic, err := appl.GetStatistic(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 4, len(statistic.StatTranslations))

		for i, s := range statistic.StatTranslations {
			assert.Equal(t, percentages[i], s.Percentage)
		}
	}
}
