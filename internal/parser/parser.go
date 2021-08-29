package parser

import "github.com/freonservice/freon/internal/entities"

type Parser interface {
	SetTranslations(t []*entities.Translation)
	Generate() (string, error)
}
