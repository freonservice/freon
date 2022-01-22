package android

import (
	"fmt"
	"strings"

	"github.com/freonservice/freon/internal/domain"
	gen "github.com/freonservice/freon/internal/generator"
)

type generator struct {
	v []*domain.Translation
}

func NewGenerator() gen.Generator {
	return &generator{}
}

func (p *generator) SetTranslations(v []*domain.Translation) gen.Generator {
	p.v = v
	return p
}

func (p *generator) SetFormat(format gen.Format) gen.Generator {
	return p
}

func (p *generator) SetPluralFormat(format gen.PluralFormat) gen.Generator {
	return p
}

func (p *generator) Generate() (gen.Document, error) {
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
	return gen.Document{
		TextFirst: f.String(),
	}, nil
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
