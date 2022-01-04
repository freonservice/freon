package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/model"
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"

	"github.com/AlekSi/pointer"
)

func (srv *server) settings(params op.SettingsParams, session *app.UserSession) op.SettingsResponder {
	state := srv.app.GetCurrentSettingState()
	return op.NewSettingsOK().WithPayload(&op.SettingsOKBody{
		Translation: &model.TranslationConfiguration{
			Auto:     pointer.ToBool(state.Translation.Auto),
			UseLibra: pointer.ToBool(state.Translation.UseLibra),
		},
	})
}

func (srv *server) settingTranslation(params op.SettingTranslationParams, session *app.UserSession) op.SettingTranslationResponder {
	return op.NewSettingTranslationNoContent()
}
