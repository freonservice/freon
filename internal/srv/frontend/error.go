//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "HealthCheck=Login,LogoutUser,RegUser,CreateLocalization,UserMe,ListLocalization,DeleteLocalization,CreateIdentifier,ListIdentifiers,DeleteIdentifier,CreateCategory,ListCategories,DeleteCategory,UpdateCategory,UpdateIdentifier,CreateTranslation,ListTranslations,DeleteTranslation,UpdateTranslation,HideTranslation,UserChangePassword,UserChangeProfile,ListUser,Statistic"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

package frontend

import (
	"net/http"

	"github.com/MarcSky/freon/api/openapi/frontend/model"
	"github.com/MarcSky/freon/api/openapi/frontend/restapi/op"
	"github.com/MarcSky/freon/pkg/def"

	"github.com/go-openapi/swag"
)

func errHealthCheck(log Log, err error, code errCode) op.HealthCheckResponder {
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.status, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = "internal error" //nolint:goconst // Duplicated by go:generate.
	}

	return op.NewHealthCheckDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(int32(code.status)),
		Message: swag.String(msg),
	})
}
