package dao

type StatTranslation struct {
	LangName  string `db:"lang_name"`
	Fulfilled int64  `db:"f"`
}

type Statistic struct {
	CountUsers         int64
	CountCategories    int64
	CountIdentifiers   int64
	CountLocalizations int64
	StatTranslations   []*StatTranslation
}
