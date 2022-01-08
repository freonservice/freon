package android

import (
	"fmt"
	"strings"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/parser"
)

type generator struct {
	v []*domain.Translation
}

func NewGenerator() parser.Generator {
	return &generator{}
}

func (p *generator) SetTranslations(v []*domain.Translation) parser.Generator {
	p.v = v
	return p
}

func (p *generator) SetFormat(format parser.Format) parser.Generator {
	return p
}

func (p *generator) SetPluralFormat(format parser.PluralFormat) parser.Generator {
	return p
}

func (p *generator) Generate() (string, error) {
	var f strings.Builder
	f.Grow(50) //nolint:gomnd

	f.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<resources>")
	f.WriteString("\n")
	for _, v := range p.v {
		if len(v.Plural) > 0 {
			f.WriteString(p.wrapPluralString(v.Identifier.Name, v.Singular, v.Plural))
			f.WriteString("\n")
		} else {
			f.WriteString(p.wrapSingleString(v.Identifier.Name, v.Singular))
		}
	}
	f.WriteString("</resources>")
	return f.String(), nil
}

func (p *generator) wrapSingleString(key, value string) string {
	return fmt.Sprintf("<string name=%q>%s</string>\n", key, value)
}

func (p *generator) wrapPluralString(key, singular, plural string) string {
	return fmt.Sprintf(
		`<plurals name="%s">
	<item quantity="one">%s</item>
	<item quantity="other">%s</item> 
</plurals>`, key, singular, plural)
}
