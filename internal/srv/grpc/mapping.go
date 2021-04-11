package grpc

import (
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/pkg/api"
)

func mappingLocalizations(ts []*app.Localization) []*api.Localization {
	txs := make([]*api.Localization, len(ts))

	for i, t := range ts {
		txs[i] = mappingLocalization(t)
	}

	return txs
}

func mappingLocalization(entity *app.Localization) *api.Localization {
	return &api.Localization{
		Id:       entity.ID,
		Locale:   entity.Locale,
		LangName: entity.LanguageName,
	}
}

func mappingTranslations(ts []*app.Translation) []*api.Translation {
	txs := make([]*api.Translation, len(ts))

	for i, t := range ts {
		txs[i] = mappingTranslation(t)
	}

	return txs
}

func mappingTranslation(entity *app.Translation) *api.Translation {
	trx := &api.Translation{
		Text: entity.Text,
	}
	if entity.Identifier != nil {
		trx.IdentifierName = entity.Identifier.Name
	}
	return trx
}
