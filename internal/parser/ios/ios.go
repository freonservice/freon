package ios

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

	for _, v := range p.v {
		if len(v.Plural) > 0 {
			f.WriteString(fmt.Sprintf("%q = %q;\n", v.Identifier.Name, v.Singular))
			f.WriteString(fmt.Sprintf("%q = %q;\n", v.Identifier.Name, v.Singular))
		} else {
			f.WriteString(fmt.Sprintf("%q = %q;\n", v.Identifier.Name, v.Singular))
		}
	}

	return f.String(), nil
}
