package frontend

import (
	"github.com/MarcSky/freon/api/openapi/frontend/restapi/op"
)

func (srv *server) HealthCheck(params op.HealthCheckParams) op.HealthCheckResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)
	status, err := srv.app.HealthCheck(ctx)
	switch {
	default:
		return errHealthCheck(log, err, codeInternal)
	case err == nil:
		return op.NewHealthCheckOK().WithPayload(status)
	}
}
