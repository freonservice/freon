package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) version(params op.VersionParams, session *app.UserSession) op.VersionResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	entities, err := srv.app.GetVersion(
		ctx,
		swag.Int64Value(params.LocalizationID),
		swag.Int64Value(params.Platform),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errVersion(log, err, codeInternal)
	case nil:
	}

	return op.NewVersionOK().WithPayload(apiArrayVersion(entities))
}
