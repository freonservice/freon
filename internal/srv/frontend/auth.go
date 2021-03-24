package frontend

import (
	"github.com/MarcSky/freon/api/openapi/frontend/restapi/op"
	"github.com/MarcSky/freon/internal/app"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) authorize(params op.LoginParams) op.LoginResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	token, user, err := srv.app.AuthorizeUser(
		ctx,
		swag.StringValue(params.Args.Email),
		params.Args.Password.String(),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errLogin(log, err, codeInternal)
	case app.ErrNotFound:
		return errLogin(log, err, codeNotFound)
	case app.ErrWrongPassword:
		return errLogin(log, err, codeWrongPassword)
	case nil:
	}

	return op.NewLoginOK().WithPayload(
		&op.LoginOKBody{
			Token: swag.String(token),
			User: &op.LoginOKBodyUser{
				Email:      &user.Email,
				FirstName:  &user.FirstName,
				Role:       &user.Role,
				SecondName: &user.SecondName,
				UUIDID:     &user.UuidID,
			},
		},
	)
}

func (srv *server) logout(params op.LogoutUserParams, session *app.UserSession) op.LogoutUserResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	err := srv.app.LogoutUser(
		ctx,
		params.HTTPRequest.Header.Get("Authorization"),
	)
	switch errors.Cause(err) {
	default:
		log.Println(errors.WithStack(err))
		return errLogoutUser(log, err, codeInternal)
	case app.ErrNotFound:
		return errLogoutUser(log, err, codeNotFound)
	case nil:
	}

	return op.NewLogoutUserNoContent()
}
