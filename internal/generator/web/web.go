package web

import (
	"fmt"

	"github.com/freonservice/freon/internal/domain"
	gen "github.com/freonservice/freon/internal/generator"
)

type generator struct {
	v []*domain.Translation

	format       gen.Format
	pluralFormat gen.PluralFormat
}

func NewGenerator() gen.Generator {
	return &generator{}
}

func (p *generator) SetTranslations(v []*domain.Translation) gen.Generator {
	p.v = v
	return p
}

func (p *generator) SetFormat(format gen.Format) gen.Generator {
	p.format = format
	return p
}

func (p *generator) SetPluralFormat(format gen.PluralFormat) gen.Generator {
	p.pluralFormat = format
	return p
}

func (p *generator) Generate() (gen.Document, error) {
	var text string
	switch p.pluralFormat { //nolint:exhaustive
	case gen.PluralFormat18N:
		text = p.generatePlural18N()
	default:
	}
	return gen.Document{TextFirst: text}, nil
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
