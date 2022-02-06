package domain

type StorageConfiguration struct {
	Use int32
}

type TranslationConfiguration struct {
	Auto         bool
	Use          int32
	CodeLanguage string
}

type SettingConfiguration struct {
	Storage     StorageConfiguration
	Translation TranslationConfiguration
}

func (t TranslationConfiguration) UseAutoTranslation() bool {
	return t.Use > 0 && t.Auto
}
