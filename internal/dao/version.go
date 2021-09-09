package dao

import "time"

type Version struct {
	Path           string
	Locale         string
	LangName       string `db:"lang_name"`
	LocalizationID int64  `db:"localization_id"`
	Platform       int64
	UpdatedAt      time.Time
}
