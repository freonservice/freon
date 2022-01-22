package ios

import "fmt"

func (p *generator) singleStringsRow(key, text string) {
	p.localizationStrings.WriteString(fmt.Sprintf("%q = %q;\n", key, text))
}

func (p *generator) startPluralStringsDict() {
	p.localizationStringsDict.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n<plist version=\"1.0\">\n  <dict>\n") //nolint:lll
}

func (p *generator) endPluralStringsDict() {
	p.localizationStringsDict.WriteString("</dict>\n</plist>\n")
}

func (p *generator) pluralStringsDictRow(key, one, other string) {
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
