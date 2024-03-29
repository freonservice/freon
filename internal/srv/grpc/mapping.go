package grpc

import (
	"github.com/freonservice/freon/internal/domain"
	api "github.com/freonservice/freon/pkg/freonApi"
)

func mappingLocalizations(ts []*domain.Localization) []*api.Localization {
	txs := make([]*api.Localization, len(ts))

	for i, t := range ts {
		txs[i] = mappingLocalization(t)
	}

	return txs
}

func mappingLocalization(entity *domain.Localization) *api.Localization {
	return &api.Localization{
		Id:       entity.ID,
		Locale:   entity.Locale,
		LangName: entity.LanguageName,
	}
}

func mappingTranslations(ts []*domain.Translation) []*api.Translation {
	txs := make([]*api.Translation, len(ts))

	for i, t := range ts {
		txs[i] = mappingTranslation(t)
	}

	return txs
}

func mappingTranslation(entity *domain.Translation) *api.Translation {
	trx := &api.Translation{
		Id:       entity.ID,
		Singular: entity.Singular,
		Plural:   entity.Plural,
	}
	if entity.Identifier != nil {
		trx.Identifier = entity.Identifier.Name
	}
	if entity.Localization != nil {
		trx.Localization = entity.Localization.Locale
	}
	return trx
}
