package domain

type TranslationConfiguration struct {
	Auto     bool
	UseLibra bool
}

type SettingConfiguration struct {
	Translation TranslationConfiguration
}
