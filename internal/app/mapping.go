package app

import (
	"strings"

	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/internal/domain"

	"github.com/AlekSi/pointer"
)

func mappingUser(user *dao.User) *domain.User {
	return &domain.User{
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

func mappingArrayUser(users []*dao.User) []*domain.User {
	var v = make([]*domain.User, len(users))
	for i, l := range users {
		v[i] = mappingUser(l)
	}
	return v
}

func mappingLocalization(localization *dao.Localization) *domain.Localization {
	return &domain.Localization{
		ID:           localization.ID,
		Locale:       localization.Locale,
		LanguageName: localization.LanguageName,
		Status:       localization.Status,
		CreatedAt:    localization.CreatedAt,
	}
}

func mappingArrayLocalization(localizations []*dao.Localization) []*domain.Localization {
	var v = make([]*domain.Localization, len(localizations))
	for i, l := range localizations {
		v[i] = mappingLocalization(l)
	}
	return v
}

func mappingIdentifier(identifier *dao.Identifier) *domain.Identifier {
	i := &domain.Identifier{
		ID:           identifier.ID,
		Name:         identifier.Name,
		Description:  identifier.Description.String,
		TextSingular: identifier.TextSingular.String,
		TextPlural:   identifier.TextPlural.String,
		Status:       identifier.Status,
		Platforms:    []string{},
		NamedList:    []string{},
		CreatedAt:    identifier.CreatedAt,
	}
	if len(identifier.Platforms) > 0 {
		i.Platforms = strings.Split(identifier.Platforms, ",")
	}
	if identifier.Category != nil && identifier.Category.ID > 0 {
		i.Category = mappingCategory(identifier.Category)
	}
	return i
}

func mappingArrayIdentifier(identifiers []*dao.Identifier) []*domain.Identifier {
	var v = make([]*domain.Identifier, len(identifiers))
	for i, l := range identifiers {
		v[i] = mappingIdentifier(l)
	}
	return v
}

func mappingCategory(category *dao.Category) *domain.Category {
	return &domain.Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}

func mappingArrayCategory(categories []*dao.Category) []*domain.Category {
	var v = make([]*domain.Category, len(categories))
	for i, c := range categories {
		v[i] = mappingCategory(c)
	}
	return v
}

func mappingTranslation(translation *dao.Translation) *domain.Translation {
	entity := &domain.Translation{
		ID:        translation.ID,
		Singular:  translation.Singular.String,
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

func mappingArrayTranslation(translations []*dao.Translation) []*domain.Translation {
	var v = make([]*domain.Translation, len(translations))
	for i, l := range translations {
		v[i] = mappingTranslation(l)
	}
	return v
}

func mappingGroupedTranslations(locale string, trxs []*dao.Translation) *domain.GroupedTranslations {
	entity := &domain.GroupedTranslations{
		Locale: locale,
	}
	var v = make([]*domain.Translation, len(trxs))
	for i, l := range trxs {
		v[i] = mappingTranslation(l)
	}
	entity.Translations = v

	return entity
}

func mappingArrayGroupedTranslations(gts map[string][]*dao.Translation) []*domain.GroupedTranslations {
	var e = make([]*domain.GroupedTranslations, len(gts))
	index := 0
	for k, v := range gts {
		e[index] = mappingGroupedTranslations(k, v)
		index++
	}
	return e
}

func mappingStatistic(statistic *dao.Statistic) *domain.Statistic {
	return &domain.Statistic{
		CountUsers:         statistic.CountUsers,
		CountCategories:    statistic.CountCategories,
		CountIdentifiers:   statistic.CountIdentifiers,
		CountLocalizations: statistic.CountLocalizations,
	}
}

func mappingTranslationFile(translationFile *dao.TranslationFile) *domain.TranslationFile {
	entity := &domain.TranslationFile{
		ID:          translationFile.ID,
		Name:        translationFile.Name,
		Path:        translationFile.Path,
		Platform:    translationFile.Platform,
		StorageType: translationFile.StorageType,
		Status:      translationFile.Status,
		CreatedAt:   translationFile.CreatedAt,
		UpdatedAt:   pointer.GetTime(translationFile.UpdatedAt),
	}
	if translationFile.Localization != nil {
		entity.Localization = mappingLocalization(translationFile.Localization)
	}
	return entity
}

func mappingArrayTranslationFile(translationFiles []*dao.TranslationFile) []*domain.TranslationFile {
	var v = make([]*domain.TranslationFile, len(translationFiles))
	for i, l := range translationFiles {
		v[i] = mappingTranslationFile(l)
	}
	return v
}

func apiVersion(v *dao.Version) *domain.Version {
	return &domain.Version{
		Path:           v.Path,
		Platform:       v.Platform,
		Locale:         v.Locale,
		LangName:       v.LangName,
		LocalizationID: v.LocalizationID,
		UpdatedAt:      v.UpdatedAt.UTC().Unix(),
	}
}

func apiArrayVersion(v []*dao.Version) []*domain.Version {
	var d = make([]*domain.Version, len(v))
	for i, e := range v {
		d[i] = apiVersion(e)
	}
	return d
}
