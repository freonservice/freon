package translation

import (
	"context"

	"github.com/freonservice/freon/internal/domain"

	"golang.org/x/text/language"
)

type Translation interface {
	Languages(ctx context.Context) ([]*domain.Language, error)
	Translate(ctx context.Context, text string, source, target language.Tag) (string, error)
}
