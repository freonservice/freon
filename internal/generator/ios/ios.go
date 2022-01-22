package ios

import (
	"strings"

	"github.com/freonservice/freon/internal/domain"
	gen "github.com/freonservice/freon/internal/generator"
)

type generator struct {
	v      []*domain.Translation
	format gen.Format

	localizationStrings     strings.Builder
	localizationStringsDict strings.Builder
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
	return p
}

func (p *generator) Generate() (gen.Document, error) {
	p.localizationStrings.Grow(50) //nolint:gomnd
	if p.format == gen.AppleStrings {
		p.localizationStringsDict.Grow(150) //nolint:gomnd
	}

	var source, target string
	p.startPlural(source, target)
	for _, v := range p.v {
		if len(v.Plural) > 0 {
			p.pluralStringsDictRow(v.Identifier.Name, v.Singular, v.Plural)
		} else {
			p.singleStringsRow(v.Identifier.Name, v.Singular)
		}
	}
	p.endPlural()

	return gen.Document{
		TextFirst:  p.localizationStrings.String(),
		TextSecond: p.localizationStringsDict.String(),
	}, nil
}

func (p *generator) startPlural(source, target string) {
	switch p.format { //nolint:exhaustive
	case gen.AppleStrings:
		p.startPluralStringsDict()
	case gen.AppleXliff:
		p.startPluralXliff(source, target)
	default:
	}
}

func (p *generator) endPlural() {
	switch p.format { //nolint:exhaustive
	case gen.AppleStrings:
		p.endPluralStringsDict()
	case gen.AppleXliff:
		p.endPluralXliff()
	default:
	}
}

func (p *generator) singleString(key, value, source string) { //nolint:unused
	switch p.format { //nolint:exhaustive
	case gen.AppleStrings:
		p.singleStringsRow(key, value)
	case gen.AppleXliff:
		p.xliffRow(key, source, value)
	default:
	}
}
