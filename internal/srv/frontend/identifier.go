package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/dal"
	"github.com/freonservice/freon/internal/filter"
	"github.com/freonservice/freon/pkg/api"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) createIdentifier( //nolint:dupl
	params op.CreateIdentifierParams, session *app.UserSession, //nolint:gocritic
) op.CreateIdentifierResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

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

func (srv *server) listIdentifiers(
	params op.ListIdentifiersParams, session *app.UserSession,
) op.ListIdentifiersResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	entities, err := srv.app.GetIdentifiers(
		ctx,
		filter.IdentifierFilter{
			CategoryID: swag.Int64Value(params.CategoryID),
			Status:     int64(api.Status_ACTIVE),
		},
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
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.DeleteIdentifier(ctx, params.ID)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errDeleteIdentifier(log, err, codeInternal)
	case nil:
	}

	return op.NewDeleteIdentifierNoContent()
}

func (srv *server) updateIdentifier( //nolint:dupl
	params op.UpdateIdentifierParams, session *app.UserSession, //nolint:gocritic
) op.UpdateIdentifierResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

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
