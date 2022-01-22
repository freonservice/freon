package ios

import "fmt"

func (p *generator) startPluralXliff(source, target string) {
	p.localizationStrings.WriteString(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<xliff xmlns="urn:oasis:names:tc:xliff:document:1.2" 
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" version="1.2" 
xsi:schemaLocation="urn:oasis:names:tc:xliff:document:1.2 
http://docs.oasis-open.org/xliff/v1.2/os/xliff-core-1.2-strict.xsd">
	<file original="" datatype="plaintext" xml:space="preserve" source-language="%s" target-language="%s">
		<header>
			<tool tool-id="lokalise.com" tool-name="Lokalise"/>
		</header>
		<body>`, source, target),
	)
}

func (p *generator) endPluralXliff() {
	p.localizationStrings.WriteString("\t\t</body>\n\t</file>\n</xliff>\n")
}

func (p *generator) xliffRow(key, source, target string) { //nolint:unused
	p.localizationStrings.WriteString(fmt.Sprintf(`			<trans-unit id=%q>
				<source>%s</source>
				<target>%s</target>
			</trans-unit>
`, key, source, target))
}
