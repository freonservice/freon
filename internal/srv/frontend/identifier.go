package frontend

import (
	"github.com/MarcSky/freon/api/openapi/frontend/restapi/op"
	"github.com/MarcSky/freon/internal/app"
	"github.com/MarcSky/freon/internal/dal"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) createIdentifier(params op.CreateIdentifierParams, session *app.UserSession) op.CreateIdentifierResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	err := srv.app.CreateIdentifier(
		ctx,
		session.UserID,
		params.Args.CategoryID,
		params.Args.ParentID,
		swag.StringValue(params.Args.Name),
		params.Args.Description,
		params.Args.ExampleText,
		params.Args.Platforms,
		params.Args.NamedList,
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errCreateIdentifier(log, err, codeInternal)
	case dal.ErrDuplicateKeyValue:
		return errCreateIdentifier(log, err, codeIdentifierIsExist)
	case nil:
	}

	return op.NewCreateIdentifierNoContent()
}

func (srv *server) listIdentifiers(params op.ListIdentifiersParams, session *app.UserSession) op.ListIdentifiersResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	entities, err := srv.app.GetIdentifiers(
		ctx,
		swag.Int64Value(params.CategoryID),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errListIdentifiers(log, err, codeInternal)
	case nil:
	}

	return op.NewListIdentifiersOK().WithPayload(apiArrayIdentifier(entities))
}

func (srv *server) deleteIdentifier(params op.DeleteIdentifierParams, session *app.UserSession) op.DeleteIdentifierResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	err := srv.app.DeleteIdentifier(ctx, params.ID)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errDeleteIdentifier(log, err, codeInternal)
	case nil:
	}

	return op.NewDeleteIdentifierNoContent()
}

func (srv *server) updateIdentifier(params op.UpdateIdentifierParams, session *app.UserSession) op.UpdateIdentifierResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	err := srv.app.UpdateIdentifier(
		ctx,
		params.ID,
		params.Args.CategoryID,
		params.Args.ParentID,
		swag.StringValue(params.Args.Name),
		params.Args.Description,
		params.Args.ExampleText,
		params.Args.Platforms,
		params.Args.NamedList,
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errUpdateIdentifier(log, err, codeInternal)
	case dal.ErrDuplicateKeyValue:
		return errUpdateIdentifier(log, err, codeIdentifierIsExist)
	case nil:
	}

	return op.NewUpdateIdentifierNoContent()
}
