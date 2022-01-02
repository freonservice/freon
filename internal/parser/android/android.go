package android

import (
	"github.com/freonservice/freon/internal/domain"
	iParser "github.com/freonservice/freon/internal/parser"
)

type parser struct {
	v []*domain.Translation
}

func NewParser() iParser.Parser {
	return &parser{}
}

func (p *parser) SetTranslations(v []*domain.Translation) {
	p.v = v
}

func (p *parser) Generate() (string, error) {
	panic("implement me")
}
