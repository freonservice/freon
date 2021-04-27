package app

import (
	"strings"

	"github.com/freonservice/freon/internal/dao"

	"github.com/AlekSi/pointer"
)

func mappingUser(user *dao.User) *User {
	return &User{
		ID:         user.ID,
		Email:      user.Email,
		Password:   user.Password,
		FirstName:  user.FirstName.String,
		SecondName: user.SecondName.String,
		UUIDID:     user.UUIDID,
		Role:       user.Role,
		Status:     user.Status,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  pointer.GetTime(user.UpdatedAt),
	}
}

func mappingArrayUser(users []*dao.User) []*User {
	var entities = make([]*User, len(users))
	for i, l := range users {
		entities[i] = mappingUser(l)
	}
	return entities
}

func mappingLocalization(localization *dao.Localization) *Localization {
	return &Localization{
		ID:           localization.ID,
		Locale:       localization.Locale,
		LanguageName: localization.LanguageName,
		Icon:         localization.Icon,
		Status:       localization.Status,
		CreatedAt:    localization.CreatedAt,
	}
}

func mappingArrayLocalization(localizations []*dao.Localization) []*Localization {
	var entities = make([]*Localization, len(localizations))
	for i, l := range localizations {
		entities[i] = mappingLocalization(l)
	}
	return entities
}

func mappingIdentifier(identifier *dao.Identifier) *Identifier {
	i := &Identifier{
		ID:          identifier.ID,
		Name:        identifier.Name,
		Description: identifier.Description.String,
		ExampleText: identifier.ExampleText.String,
		Status:      identifier.Status,
		Platforms:   []string{},
		NamedList:   []string{},
		CreatedAt:   identifier.CreatedAt,
	}
	if len(identifier.Platforms) > 0 {
		i.Platforms = strings.Split(identifier.Platforms, ",")
	}
	if identifier.NamedList.Valid && len(identifier.NamedList.String) > 0 {
		i.NamedList = strings.Split(identifier.NamedList.String, ",")
	}
	if identifier.Category != nil && identifier.Category.ID > 0 {
		i.Category = mappingCategory(identifier.Category)
	}
	return i
}

func mappingArrayIdentifier(identifiers []*dao.Identifier) []*Identifier {
	var entities = make([]*Identifier, len(identifiers))
	for i, l := range identifiers {
		entities[i] = mappingIdentifier(l)
	}
	return entities
}

func mappingCategory(category *dao.Category) *Category {
	return &Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}

func mappingArrayCategory(categories []*dao.Category) []*Category {
	var entities = make([]*Category, len(categories))
	for i, c := range categories {
		entities[i] = mappingCategory(c)
	}
	return entities
}

func mappingTranslation(translation *dao.Translation) *Translation {
	entity := &Translation{
		ID:        translation.ID,
		Text:      translation.Text,
		Status:    translation.Status,
		CreatedAt: translation.CreatedAt,
	}
	if translation.Localization != nil {
		entity.Localization = mappingLocalization(translation.Localization)
	}
	if translation.Identifier != nil {
		entity.Identifier = mappingIdentifier(translation.Identifier)
	}
	return entity
}

func mappingArrayTranslation(translations []*dao.Translation) []*Translation {
	var entities = make([]*Translation, len(translations))
	for i, l := range translations {
		entities[i] = mappingTranslation(l)
	}
	return entities
}

func mappingGroupedTranslations(locale string, trxs []*dao.Translation) *GroupedTranslations {
	entity := &GroupedTranslations{
		Locale: locale,
	}
	var entities = make([]*Translation, len(trxs))
	for i, l := range trxs {
		entities[i] = mappingTranslation(l)
	}
	entity.Translations = entities

	return entity
}

func mappingArrayGroupedTranslations(gts map[string][]*dao.Translation) []*GroupedTranslations {
	var entities = make([]*GroupedTranslations, len(gts))
	index := 0
	for k, v := range gts {
		entities[index] = mappingGroupedTranslations(k, v)
		index++
	}
	return entities
}

func mappingStatistic(statistic *dao.Statistic) *Statistic {
	return &Statistic{
		CountUsers:         statistic.CountUsers,
		CountCategories:    statistic.CountCategories,
		CountIdentifiers:   statistic.CountIdentifiers,
		CountLocalizations: statistic.CountLocalizations,
	}
}

func mappingTranslationFile(translationFile *dao.TranslationFile) *TranslationFile {
	entity := &TranslationFile{
		ID:          translationFile.ID,
		Name:        translationFile.Name,
		Path:        translationFile.Path,
		Platform:    translationFile.Platform,
		StorageType: translationFile.StorageType,
		Status:      translationFile.Status,
		CreatedAt:   translationFile.CreatedAt,
	}
	if translationFile.Localization != nil {
		entity.Localization = mappingLocalization(translationFile.Localization)
	}
	return entity
}

func mappingArrayTranslationFile(translationFiles []*dao.TranslationFile) []*TranslationFile {
	var entities = make([]*TranslationFile, len(translationFiles))
	for i, l := range translationFiles {
		entities[i] = mappingTranslationFile(l)
	}
	return entities
}
