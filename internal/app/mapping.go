package app

import (
	"strings"

	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/entities"

	"github.com/AlekSi/pointer"
)

func mappingUser(user *dao.User) *entities.User {
	return &entities.User{
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

func mappingArrayUser(users []*dao.User) []*entities.User {
	var v = make([]*entities.User, len(users))
	for i, l := range users {
		v[i] = mappingUser(l)
	}
	return v
}

func mappingLocalization(localization *dao.Localization) *entities.Localization {
	return &entities.Localization{
		ID:           localization.ID,
		Locale:       localization.Locale,
		LanguageName: localization.LanguageName,
		Icon:         localization.Icon,
		Status:       localization.Status,
		CreatedAt:    localization.CreatedAt,
	}
}

func mappingArrayLocalization(localizations []*dao.Localization) []*entities.Localization {
	var v = make([]*entities.Localization, len(localizations))
	for i, l := range localizations {
		v[i] = mappingLocalization(l)
	}
	return v
}

func mappingIdentifier(identifier *dao.Identifier) *entities.Identifier {
	i := &entities.Identifier{
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
	if identifier.Category != nil && identifier.Category.ID > 0 {
		i.Category = mappingCategory(identifier.Category)
	}
	return i
}

func mappingArrayIdentifier(identifiers []*dao.Identifier) []*entities.Identifier {
	var v = make([]*entities.Identifier, len(identifiers))
	for i, l := range identifiers {
		v[i] = mappingIdentifier(l)
	}
	return v
}

func mappingCategory(category *dao.Category) *entities.Category {
	return &entities.Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}

func mappingArrayCategory(categories []*dao.Category) []*entities.Category {
	var v = make([]*entities.Category, len(categories))
	for i, c := range categories {
		v[i] = mappingCategory(c)
	}
	return v
}

func mappingTranslation(translation *dao.Translation) *entities.Translation {
	entity := &entities.Translation{
		ID:        translation.ID,
		Singular:  translation.Singular,
		Plural:    translation.Plural.String,
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

func mappingArrayTranslation(translations []*dao.Translation) []*entities.Translation {
	var v = make([]*entities.Translation, len(translations))
	for i, l := range translations {
		v[i] = mappingTranslation(l)
	}
	return v
}

func mappingGroupedTranslations(locale string, trxs []*dao.Translation) *entities.GroupedTranslations {
	entity := &entities.GroupedTranslations{
		Locale: locale,
	}
	var v = make([]*entities.Translation, len(trxs))
	for i, l := range trxs {
		v[i] = mappingTranslation(l)
	}
	entity.Translations = v

	return entity
}

func mappingArrayGroupedTranslations(gts map[string][]*dao.Translation) []*entities.GroupedTranslations {
	var e = make([]*entities.GroupedTranslations, len(gts))
	index := 0
	for k, v := range gts {
		e[index] = mappingGroupedTranslations(k, v)
		index++
	}
	return e
}

func mappingStatistic(statistic *dao.Statistic) *entities.Statistic {
	return &entities.Statistic{
		CountUsers:         statistic.CountUsers,
		CountCategories:    statistic.CountCategories,
		CountIdentifiers:   statistic.CountIdentifiers,
		CountLocalizations: statistic.CountLocalizations,
	}
}

func mappingTranslationFile(translationFile *dao.TranslationFile) *entities.TranslationFile {
	entity := &entities.TranslationFile{
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

func mappingArrayTranslationFile(translationFiles []*dao.TranslationFile) []*entities.TranslationFile {
	var v = make([]*entities.TranslationFile, len(translationFiles))
	for i, l := range translationFiles {
		v[i] = mappingTranslationFile(l)
	}
	return v
}
