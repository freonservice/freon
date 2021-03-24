package dao

type StatTranslation struct {
	LangName string `db:"lang_name"`
	Count    int64  `db:"c"`
}

type Statistic struct {
	CountUsers         int64
	CountCategories    int64
	CountIdentifiers   int64
	CountLocalizations int64
	StatTranslations   []*StatTranslation
}
