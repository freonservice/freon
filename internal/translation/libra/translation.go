package libra

import (
	"context"
	"log"
	"time"

	iface "github.com/freonservice/freon/internal/translation"
	libra "github.com/freonservice/libretranslate-sdk"

	"github.com/pkg/errors"
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

func (t translation) Languages(ctx context.Context) ([]iface.Language, error) {
	libraLanguages, err := t.client.GetLanguages(ctx)
	if err != nil {
		return nil, err
	}
	return mappingArrayLanguages(libraLanguages), nil
}

func (t translation) Translate(ctx context.Context, text string, source, target language.Tag) (string, error) {
	return t.client.Translate(ctx, text, source.String(), target.String())
}

func mappingArrayLanguages(libraLanguages []libra.Language) []iface.Language {
	languages := make([]iface.Language, 0, len(libraLanguages))
	for i := range libraLanguages {
		code, err := language.Parse(libraLanguages[i].Code)
		if err != nil {
			log.Println("mappingArrayLanguages", errors.WithStack(err))
			continue
		}
		languages = append(languages, iface.Language{
			Name: libraLanguages[i].Name,
			Code: code,
		})
	}
	return languages
}
