package web

import (
	"fmt"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/parser"
)

type generator struct {
	v []*domain.Translation

	format       parser.Format
	pluralFormat parser.PluralFormat
}

func NewGenerator() parser.Generator {
	return &generator{}
}

func (p *generator) SetTranslations(v []*domain.Translation) parser.Generator {
	p.v = v
	return p
}

func (p *generator) SetFormat(format parser.Format) parser.Generator {
	p.format = format
	return p
}

func (p *generator) SetPluralFormat(format parser.PluralFormat) parser.Generator {
	p.pluralFormat = format
	return p
}

func (p *generator) Generate() ([]string, error) {
	var text string
	switch p.pluralFormat { //nolint:exhaustive
	case parser.PluralFormat18N:
		text = p.generatePlural18N()
	default:
	}
	return []string{text}, nil
}

func (p *generator) generatePlural18N() string {
	var f string
	f += "{"
	for _, v := range p.v {
		if len(v.Plural) > 0 {
			f += fmt.Sprintf("%q:%q,\n", v.Identifier.Name, v.Singular)
			f += fmt.Sprintf("\"%s_plural\":%q,\n", v.Identifier.Name, v.Plural)
		} else {
			f += fmt.Sprintf("%q:%q,\n", v.Identifier.Name, v.Singular)
		}
	}
	f = f[:len(f)-2]
	f += "}"
	return f
}
