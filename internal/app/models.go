package app

import "time"

type User struct {
	ID         int64
	Email      string
	Password   string
	FirstName  string
	SecondName string
	UUIDID     string
	Role       int64
	Status     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ChangePassword struct {
	PreviousPassword string
	NewPassword      string
}

type Localization struct {
	ID           int64
	CreatorID    int64
	Locale       string
	LanguageName string
	Icon         string
	Status       int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type LocalizationIdentifier struct {
	ID             int64
	LocalizationID int64
	IdentifierID   int64
	Status         int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Category struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Identifier struct {
	ID          int64
	CreatorID   int64
	CategoryID  int64
	Name        string
	Description string
	ExampleText string
	Status      int64
	Platforms   []string
	NamedList   []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Category    *Category
}

type Translation struct {
	ID           int64
	CreatorID    int64
	Text         string
	Status       int64
	Localization *Localization
	Identifier   *Identifier
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type GroupedTranslations struct {
	Locale       string
	Translations []*Translation
}

type StatTranslation struct {
	LangName   string
	Percentage float64
}

type Statistic struct {
	CountUsers         int64
	CountCategories    int64
	CountIdentifiers   int64
	CountLocalizations int64
	StatTranslations   []*StatTranslation
}

type TranslationFile struct {
	ID             int64
	CreatorID      int64
	LocalizationID int64
	Name           string
	Path           string
	Platform       int64
	StorageType    int64
	Status         int64
	Localization   *Localization
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
