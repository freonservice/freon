package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/model"
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/pkg/api"

	"github.com/AlekSi/pointer"
)

func apiUser(v *app.User) *model.User {
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

func apiArrayUser(v []*app.User) []*model.User {
	var entities = make([]*model.User, len(v))
	for i, e := range v {
		entities[i] = apiUser(e)
	}
	return entities
}

func apiLocalization(v *app.Localization) *model.Localization {
	t := v.CreatedAt.UTC().Unix()
	return &model.Localization{
		ID:        &v.ID,
		Locale:    &v.Locale,
		LangName:  &v.LanguageName,
		Icon:      v.Icon,
		CreatedAt: &t,
	}
}

func apiArrayLocalization(v []*app.Localization) []*model.Localization {
	var entities = make([]*model.Localization, len(v))
	for i, e := range v {
		entities[i] = apiLocalization(e)
	}
	return entities
}

func apiIdentifier(v *app.Identifier) *model.Identifier {
	i := &model.Identifier{
		ID:          &v.ID,
		Name:        &v.Name,
		Description: v.Description,
		ExampleText: v.ExampleText,
		Platforms:   v.Platforms,
		NamedList:   v.NamedList,
	}
	if v.Category != nil && v.Category.ID > 0 {
		i.Category = apiCategory(v.Category)
	}
	return i
}

func apiArrayIdentifier(v []*app.Identifier) []*model.Identifier {
	var entities = make([]*model.Identifier, len(v))
	for i, e := range v {
		entities[i] = apiIdentifier(e)
	}
	return entities
}

func apiCategory(v *app.Category) *model.Category {
	return &model.Category{
		ID:   &v.ID,
		Name: &v.Name,
	}
}

func apiArrayCategory(v []*app.Category) []*model.Category {
	var entities = make([]*model.Category, len(v))
	for i, e := range v {
		entities[i] = apiCategory(e)
	}
	return entities
}

func apiTranslation(v *app.Translation) *model.Translation {
	i := &model.Translation{
		ID:           &v.ID,
		Text:         &v.Text,
		Localization: apiLocalization(v.Localization),
		Identifier:   apiIdentifier(v.Identifier),
		Status:       getTranslationStatus(api.TranslationStatus(v.Status)),
		CreatedAt:    pointer.ToInt64(v.CreatedAt.Unix()),
	}
	return i
}

func apiArrayTranslation(v []*app.Translation) []*model.Translation {
	var entities = make([]*model.Translation, len(v))
	for i, e := range v {
		entities[i] = apiTranslation(e)
	}
	return entities
}

func apiStatistic(v *app.Statistic) *op.StatisticOKBody {
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

func apiTranslationFile(v *app.TranslationFile) *model.TranslationFile {
	i := &model.TranslationFile{
		ID:          &v.ID,
		Name:        &v.Name,
		Path:        &v.Path,
		Platform:    pointer.ToString(getPlatformByInteger(v.Platform)),
		Status:      getStatusByInteger(api.Status(v.Status)),
		StorageType: pointer.ToString(getStorageTypeByInteger(v.StorageType)),
	}
	return i
}

func apiArrayTranslationFiles(v []*app.TranslationFile) []*model.TranslationFile {
	var entities = make([]*model.TranslationFile, len(v))
	for i, e := range v {
		entities[i] = apiTranslationFile(e)
	}
	return entities
}
