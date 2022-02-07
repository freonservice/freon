package libra

import (
	"context"
	"time"

	"github.com/freonservice/freon/internal/domain"
	iface "github.com/freonservice/freon/internal/translation"
	libra "github.com/freonservice/libretranslate-sdk"

	"golang.org/x/text/language"
)

type translation struct {
	client libra.Client
}

func NewLibraTranslation(apiURL string, timeout time.Duration) iface.Translation {
	return &translation{
		client: libra.NewLibreTranslate(apiURL).SetConnTimeout(timeout),
	}
}

func (t translation) Languages(ctx context.Context) ([]*domain.Language, error) {
	libraLanguages, err := t.client.GetLanguages(ctx)
	if err != nil {
		return nil, err
	}
	return mappingArrayLanguages(libraLanguages), nil
}

func (t translation) Translate(ctx context.Context, text string, source, target language.Tag) (string, error) {
	return t.client.Translate(ctx, text, source.String(), target.String())
}

func mappingArrayLanguages(libraLanguages []libra.Language) []*domain.Language {
	languages := make([]*domain.Language, len(libraLanguages))
	for i := range libraLanguages {
		languages[i] = &domain.Language{
			Name: libraLanguages[i].Name,
			Code: libraLanguages[i].Code,
		}
	}
	return languages
}
