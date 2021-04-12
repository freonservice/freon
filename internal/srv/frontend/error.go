package frontend

import (
	"net/http"

	"github.com/freonservice/freon/api/openapi/frontend/model"
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/pkg/def"

	"github.com/go-openapi/swag"
)

func errHealthCheck(log Log, err error, code errCode) op.HealthCheckResponder { //nolint:dupl
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = "internal error" //nolint:goconst
	}

	return op.NewHealthCheckDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}
