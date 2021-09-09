package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/model"
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/entities"
	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/AlekSi/pointer"
)

func apiUser(v *entities.User) *model.User {
	return &model.User{
		ID:         &v.ID,
		Email:      &v.Email,
		FirstName:  v.FirstName,
		SecondName: v.SecondName,
		UUIDID:     &v.UUIDID,
		Role:       pointer.ToString(getUserRoleByInteger(api.UserRole(v.Role))),
		Status:     pointer.ToString(getUserStatusByInteger(api.UserStatus(v.Status))),
		CreatedAt:  pointer.ToInt64(v.CreatedAt.Unix()),
	}
}

func apiArrayUser(v []*entities.User) []*model.User {
	var d = make([]*model.User, len(v))
	for i, e := range v {
		d[i] = apiUser(e)
	}
	return d
}

func apiLocalization(v *entities.Localization) *model.Localization {
	t := v.CreatedAt.UTC().Unix()
	return &model.Localization{
		ID:        &v.ID,
		Locale:    &v.Locale,
		LangName:  &v.LanguageName,
		Icon:      v.Icon,
		CreatedAt: &t,
	}
}

func apiArrayLocalization(v []*entities.Localization) []*model.Localization {
	var d = make([]*model.Localization, len(v))
	for i, e := range v {
		d[i] = apiLocalization(e)
	}
	return d
}

func apiIdentifier(v *entities.Identifier) *model.Identifier {
	i := &model.Identifier{
		ID:          &v.ID,
		Name:        &v.Name,
		Description: v.Description,
		ExampleText: v.ExampleText,
		Platforms:   v.Platforms,
	}
	if v.Category != nil && v.Category.ID > 0 {
		i.Category = apiCategory(v.Category)
	}
	return i
}

func apiArrayIdentifier(v []*entities.Identifier) []*model.Identifier {
	var d = make([]*model.Identifier, len(v))
	for i, e := range v {
		d[i] = apiIdentifier(e)
	}
	return d
}

func apiCategory(v *entities.Category) *model.Category {
	return &model.Category{
		ID:   &v.ID,
		Name: &v.Name,
	}
}

func apiArrayCategory(v []*entities.Category) []*model.Category {
	var d = make([]*model.Category, len(v))
	for i, e := range v {
		d[i] = apiCategory(e)
	}
	return d
}

func apiTranslation(v *entities.Translation) *model.Translation {
	i := &model.Translation{
		ID:           &v.ID,
		Singular:     &v.Singular,
		Plural:       v.Plural,
		Localization: apiLocalization(v.Localization),
		Identifier:   apiIdentifier(v.Identifier),
		Status:       getTranslationStatus(api.StatusTranslation(v.Status)),
		CreatedAt:    pointer.ToInt64(v.CreatedAt.Unix()),
	}
	return i
}

func apiArrayTranslation(v []*entities.Translation) []*model.Translation {
	var d = make([]*model.Translation, len(v))
	for i, e := range v {
		d[i] = apiTranslation(e)
	}
	return d
}

func apiStatistic(v *entities.Statistic) *op.StatisticOKBody {
	entity := &op.StatisticOKBody{
		CountCategories:    &v.CountCategories,
		CountIdentifiers:   &v.CountIdentifiers,
		CountLocalizations: &v.CountLocalizations,
		CountUsers:         &v.CountUsers,
	}

	var statCompletedTranslations = make([]*op.StatisticOKBodyStatCompletedTranslationsItems0, len(v.StatTranslations))
	for i := 0; i < len(v.StatTranslations); i++ {
		statCompletedTranslations[i] = &op.StatisticOKBodyStatCompletedTranslationsItems0{
			LangName:   &v.StatTranslations[i].LangName,
			Percentage: &v.StatTranslations[i].Percentage,
		}
	}
	entity.StatCompletedTranslations = statCompletedTranslations

	return entity
}

func apiTranslationFile(v *entities.TranslationFile) *model.TranslationFile {
	i := &model.TranslationFile{
		ID:          &v.ID,
		Name:        &v.Name,
		Path:        &v.Path,
		Platform:    pointer.ToString(getPlatformByInteger(v.Platform)),
		Status:      getStatusByInteger(api.Status(v.Status)),
		StorageType: pointer.ToString(getStorageTypeByInteger(v.StorageType)),
		CreatedAt:   pointer.ToInt64(v.CreatedAt.UTC().Unix()),
		UpdatedAt:   pointer.ToInt64(v.UpdatedAt.UTC().Unix()),
	}
	return i
}

func apiArrayTranslationFiles(v []*entities.TranslationFile) []*model.TranslationFile {
	var d = make([]*model.TranslationFile, len(v))
	for i, e := range v {
		d[i] = apiTranslationFile(e)
	}
	return d
}

func apiVersion(v *entities.Version) *model.Version {
	return &model.Version{
		Path:           v.Path,
		Platform:       v.Platform,
		Locale:         &v.Locale,
		LangName:       v.LangName,
		LocalizationID: v.LocalizationID,
		UpdatedAt:      &v.UpdatedAt,
	}
}

func apiArrayVersion(v []*entities.Version) []*model.Version {
	var d = make([]*model.Version, len(v))
	for i, e := range v {
		d[i] = apiVersion(e)
	}
	return d
}
