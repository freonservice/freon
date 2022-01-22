package dao

import (
	"time"

	"github.com/AlekSi/pointer"
)

// go:generate reform
// reform:translation_files
type TranslationFile struct {
	ID             int64      `reform:"id,pk"`
	LocalizationID int64      `reform:"localization_id"`
	CreatorID      int64      `reform:"creator_id"`
	Name           string     `reform:"name"`
	Path           string     `reform:"path"`
	S3FileID       string     `reform:"s3_file_id"`
	S3Bucket       string     `reform:"s3_bucket"`
	Platform       int64      `reform:"platform"`
	Status         int64      `reform:"status"`
	StorageType    int64      `reform:"storage_type"`
	CreatedAt      time.Time  `reform:"created_at"`
	UpdatedAt      *time.Time `reform:"updated_at"`

	Localization *Localization
	Creator      *User
}

func (u *TranslationFile) BeforeInsert() error {
	if u.UpdatedAt != nil {
		u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second).AddDate(0, 0, 0))
	}
	return nil
}

func (u *TranslationFile) BeforeUpdate() error {
	now := time.Now().UTC()
	u.UpdatedAt = &now
	u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second).AddDate(0, 0, 0))
	return nil
}
