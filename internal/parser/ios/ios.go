package ios

import (
	"fmt"

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
	var f string

	for _, v := range p.v {
		f += fmt.Sprintf("\"%q\" = \"%q\";\n", v.Localization.LanguageName, v.Singular)
	}

	return f, nil
}
