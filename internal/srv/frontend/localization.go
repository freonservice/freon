package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/dal"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) createLocalization(params op.CreateLocalizationParams, session *app.UserSession) op.CreateLocalizationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.CreateLocalization(
		ctx,
		session.UserID,
		swag.StringValue(params.Args.Locale),
		swag.StringValue(params.Args.LangName),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errCreateLocalization(log, err, codeInternal)
	case dal.ErrDuplicateKeyValue:
		return errCreateLocalization(log, err, codeLocalizationIsExist)
	case nil:
	}

	return op.NewCreateLocalizationNoContent()
}

func (srv *server) listLocalization(params op.ListLocalizationParams, session *app.UserSession) op.ListLocalizationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	entities, err := srv.app.GetLocalizations(
		ctx,
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errListLocalization(log, err, codeInternal)
	case nil:
	}

	return op.NewListLocalizationOK().WithPayload(apiArrayLocalization(entities))
}

func (srv *server) deleteLocalization(params op.DeleteLocalizationParams, session *app.UserSession) op.DeleteLocalizationResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.DeleteLocalization(
		ctx,
		params.ID,
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errDeleteLocalization(log, err, codeInternal)
	case nil:
	}

	return op.NewDeleteLocalizationNoContent()
}
