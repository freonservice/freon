package parser

import (
	"github.com/freonservice/freon/internal/domain"
)

type Parser interface {
	SetTranslations(t []*domain.Translation)
	Generate() (string, error)
}
