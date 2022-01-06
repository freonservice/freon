package domain

type TranslationConfiguration struct {
	Auto bool
	Use  int32
}

type SettingConfiguration struct {
	Translation TranslationConfiguration
}
