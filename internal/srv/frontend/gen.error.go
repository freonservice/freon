// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package frontend

import (
	"net/http"

	"github.com/freonservice/freon/api/openapi/frontend/model"
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/pkg/def"
	"github.com/go-openapi/swag"
)

func errLogin(log Log, err error, code errCode) op.LoginResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewLoginDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errLogoutUser(log Log, err error, code errCode) op.LogoutUserResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewLogoutUserDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errRegUser(log Log, err error, code errCode) op.RegUserResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewRegUserDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errCreateLocalization(log Log, err error, code errCode) op.CreateLocalizationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewCreateLocalizationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errUserMe(log Log, err error, code errCode) op.UserMeResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewUserMeDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errInfo(log Log, err error, code errCode) op.InfoResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewInfoDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errListLocalization(log Log, err error, code errCode) op.ListLocalizationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewListLocalizationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errDeleteLocalization(log Log, err error, code errCode) op.DeleteLocalizationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewDeleteLocalizationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errCreateIdentifier(log Log, err error, code errCode) op.CreateIdentifierResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewCreateIdentifierDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errListIdentifiers(log Log, err error, code errCode) op.ListIdentifiersResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewListIdentifiersDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errDeleteIdentifier(log Log, err error, code errCode) op.DeleteIdentifierResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewDeleteIdentifierDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errCreateCategory(log Log, err error, code errCode) op.CreateCategoryResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewCreateCategoryDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errListCategories(log Log, err error, code errCode) op.ListCategoriesResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewListCategoriesDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errDeleteCategory(log Log, err error, code errCode) op.DeleteCategoryResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewDeleteCategoryDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errUpdateCategory(log Log, err error, code errCode) op.UpdateCategoryResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewUpdateCategoryDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errUpdateIdentifier(log Log, err error, code errCode) op.UpdateIdentifierResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewUpdateIdentifierDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errCreateTranslation(log Log, err error, code errCode) op.CreateTranslationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewCreateTranslationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errListTranslations(log Log, err error, code errCode) op.ListTranslationsResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewListTranslationsDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errDeleteTranslation(log Log, err error, code errCode) op.DeleteTranslationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewDeleteTranslationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errUpdateTranslation(log Log, err error, code errCode) op.UpdateTranslationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewUpdateTranslationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errStatusTranslation(log Log, err error, code errCode) op.StatusTranslationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewStatusTranslationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errUserChangePassword(log Log, err error, code errCode) op.UserChangePasswordResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewUserChangePasswordDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errUserChangeProfile(log Log, err error, code errCode) op.UserChangeProfileResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewUserChangeProfileDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errListUser(log Log, err error, code errCode) op.ListUserResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewListUserDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errStatistic(log Log, err error, code errCode) op.StatisticResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewStatisticDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errListTranslationFiles(log Log, err error, code errCode) op.ListTranslationFilesResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewListTranslationFilesDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errDeleteTranslationFile(log Log, err error, code errCode) op.DeleteTranslationFileResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewDeleteTranslationFileDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errVersion(log Log, err error, code errCode) op.VersionResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewVersionDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errSettingTranslation(log Log, err error, code errCode) op.SettingTranslationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewSettingTranslationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errSettingStorage(log Log, err error, code errCode) op.SettingStorageResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewSettingStorageDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errAutoTranslation(log Log, err error, code errCode) op.AutoTranslationResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewAutoTranslationDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errSupportedLanguages(log Log, err error, code errCode) op.SupportedLanguagesResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewSupportedLanguagesDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errSettingFirstLaunch(log Log, err error, code errCode) op.SettingFirstLaunchResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewSettingFirstLaunchDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}

func errAutoTranslationByID(log Log, err error, code errCode) op.AutoTranslationByIDResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = internalError
	}

	return op.NewAutoTranslationByIDDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}
