package parser

import (
	"github.com/freonservice/freon/internal/domain"
)

type Format int8

const (
	AppleStrings Format = iota
	AndroidStrings
	WebJSON
)

type PluralFormat int8

const (
	PluralFormatDefault PluralFormat = iota // using for android, apple
	PluralFormat18N                         // using for web
)

type Generator interface {
	SetTranslations(t []*domain.Translation) Generator
	SetFormat(format Format) Generator             // web-json, ios-strings, android-xml
	SetPluralFormat(format PluralFormat) Generator // json, i18n
	Generate() (string, error)
}
