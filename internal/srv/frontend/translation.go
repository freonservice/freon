package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/pkg/api"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) createTranslation(params op.CreateTranslationParams, session *app.UserSession) op.CreateTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.CreateTranslation(
		ctx,
		session.UserID,
		swag.Int64Value(params.Args.LocalizationID),
		swag.Int64Value(params.Args.IdentifierID),
		swag.StringValue(params.Args.Text),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errCreateTranslation(log, err, codeInternal)
	case nil:
	}

	return op.NewCreateTranslationNoContent()
}

func (srv *server) listTranslations(params op.ListTranslationsParams, session *app.UserSession) op.ListTranslationsResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	entities, err := srv.app.GetTranslations(
		ctx,
		swag.Int64Value(params.LocalizationID),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errListTranslations(log, err, codeInternal)
	case nil:
	}

	return op.NewListTranslationsOK().WithPayload(apiArrayTranslation(entities))
}

func (srv *server) deleteTranslation(params op.DeleteTranslationParams, session *app.UserSession) op.DeleteTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.DeleteTranslation(ctx, params.ID)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errDeleteTranslation(log, err, codeInternal)
	case nil:
	}

	return op.NewDeleteTranslationNoContent()
}

func (srv *server) updateTranslation(params op.UpdateTranslationParams, session *app.UserSession) op.UpdateTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.UpdateTranslation(
		ctx,
		params.ID,
		swag.StringValue(params.Args.Text),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errUpdateTranslation(log, err, codeInternal)
	case nil:
	}

	return op.NewUpdateTranslationNoContent()
}

func (srv *server) hideTranslation(params op.HideTranslationParams, session *app.UserSession) op.HideTranslationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.HideTranslation(ctx, params.ID, params.Hide)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errHideTranslation(log, err, codeInternal)
	case nil:
	}

	return op.NewHideTranslationNoContent()
}

func getTranslationStatus(status api.TranslationStatus) string {
	switch status { //nolint:exhaustive
	case api.TranslationStatus_TRANSLATION_ACTIVE:
		return "Active"
	default:
		return "Hidden"
	}
}
