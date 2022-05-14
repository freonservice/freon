package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/model"
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/domain"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
)

func (srv *server) settings(params op.SettingsParams, session *app.UserSession) op.SettingsResponder {
	state := srv.app.GetCurrentSettingState()
	return op.NewSettingsOK().WithPayload(&op.SettingsOKBody{
		Translation: &model.TranslationConfiguration{
			Auto:         pointer.ToBool(state.Translation.Auto),
			Use:          pointer.ToInt32(state.Translation.Use),
			MainLanguage: pointer.ToString(state.Translation.MainLanguage),
		},
		Storage: &model.StorageConfiguration{
			Use: pointer.ToInt32(state.Storage.Use),
		},
		FirstLaunch: state.FirstLaunch,
	})
}

func (srv *server) settingTranslation(params op.SettingTranslationParams, session *app.UserSession) op.SettingTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)
	err := srv.app.SetTranslationConfiguration(ctx, domain.TranslationConfiguration{
		Auto:         params.Args.Auto,
		Use:          params.Args.Use,
		MainLanguage: params.Args.MainLanguage,
	})
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errSettingTranslation(log, err, codeInternal)
	case nil:
	}
	return op.NewSettingTranslationNoContent()
}

func (srv *server) settingStorage(params op.SettingStorageParams, session *app.UserSession) op.SettingStorageResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)
	err := srv.app.SetStorageConfiguration(ctx, domain.StorageConfiguration{
		Use: params.Args.Use,
	})
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errSettingStorage(log, err, codeInternal)
	case nil:
	}
	return op.NewSettingStorageNoContent()
}

func (srv *server) settingFirstLaunch(params op.SettingFirstLaunchParams, session *app.UserSession) op.SettingFirstLaunchResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)
	err := srv.app.DisableSettingFirstLaunch(ctx)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errSettingFirstLaunch(log, err, codeInternal)
	case nil:
	}
	return op.NewSettingFirstLaunchNoContent()
}
