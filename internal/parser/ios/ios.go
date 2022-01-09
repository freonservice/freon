package ios

import (
	"fmt"
	"strings"

	"github.com/freonservice/freon/internal/domain"
	"github.com/freonservice/freon/internal/parser"
)

type generator struct {
	v []*domain.Translation

	localizationStrings     strings.Builder
	localizationStringsDict strings.Builder
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

func (p *generator) Generate() ([]string, error) {
	p.localizationStrings.Grow(50)      //nolint:gomnd
	p.localizationStringsDict.Grow(150) //nolint:gomnd

	p.startPluralDict()
	for _, v := range p.v {
		if len(v.Plural) > 0 {
			p.pluralDict(v.Identifier.Name, v.Singular, v.Plural)
		} else {
			p.localizationStrings.WriteString(fmt.Sprintf("%q = %q;\n", v.Identifier.Name, v.Singular))
		}
	}
	p.endPluralDict()

	return []string{p.localizationStrings.String(), p.localizationStringsDict.String()}, nil
}

func (p *generator) startPluralDict() {
	p.localizationStringsDict.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n<plist version=\"1.0\">\n  <dict>\n") //nolint:lll
}

func (p *generator) endPluralDict() {
	p.localizationStringsDict.WriteString("</dict>\n</plist>\n")
}

func (p *generator) pluralDict(key, one, other string) {
	p.localizationStringsDict.WriteString(fmt.Sprintf(`
	<key>%s</key>
    <dict>
        <key>NSStringLocalizedFormatKey</key>
        <string>%%#@value@</string>
        <key>value</key>
        <dict>
            <key>NSStringFormatSpecTypeKey</key>
            <string>NSStringPluralRuleType</string>
            <key>NSStringFormatValueTypeKey</key>
            <string>d</string>
            <key>one</key>
            <string>%s</string>
            <key>other</key>
            <string>%s</string>
        </dict>
    </dict>
`, key, one, other))
}
