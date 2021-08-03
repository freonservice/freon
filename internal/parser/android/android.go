package android

import (
	"github.com/freonservice/freon/internal/entities"
	iParser "github.com/freonservice/freon/internal/parser"
)

type parser struct {
	v []*entities.Translation
}

func NewParser() iParser.Parser {
	return &parser{}
}

func (p *parser) SetTranslations(v []*entities.Translation) {
	p.v = v
}

func (p *parser) Generate() (string, error) {
	panic("implement me")
}
