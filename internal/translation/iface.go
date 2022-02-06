package translation

import (
	"context"

	"golang.org/x/text/language"
)

type Language struct {
	Name string
	Code language.Tag
}

type Translation interface {
	Languages(ctx context.Context) ([]Language, error)
	Translate(ctx context.Context, text string, source, target language.Tag) (string, error)
}
