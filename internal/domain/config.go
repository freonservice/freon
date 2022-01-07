package domain

type StorageConfiguration struct {
	Use int32
}

type TranslationConfiguration struct {
	Auto bool
	Use  int32
}

type SettingConfiguration struct {
	Storage     StorageConfiguration
	Translation TranslationConfiguration
}
