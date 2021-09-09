package dao

import "time"

type Version struct {
	Path           string
	LocalizationID int64 `db:"localization_id"`
	Platform       int64
	UpdatedAt      time.Time

	Localization *Localization
}
