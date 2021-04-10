package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) regUser(params op.RegUserParams, session *app.UserSession) op.RegUserResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	if params.Args.Password.String() != params.Args.RepeatPassword.String() {
		return errRegUser(log, errors.New("passwords not equals"), codePasswordsNotEquals)
	}
	_, err := srv.app.RegisterUser(
		ctx,
		swag.StringValue(params.Args.Email),
		params.Args.Password.String(),
		swag.StringValue(params.Args.FirstName),
		swag.StringValue(params.Args.SecondName),
		int64(getUserRoleByString(swag.StringValue(params.Args.Role))),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errRegUser(log, err, codeInternal)
	case app.ErrEmailIsUsed:
		return errRegUser(log, err, codeEmailUsed)
	case nil:
	}

	return op.NewRegUserNoContent()
}

func (srv *server) userMe(params op.UserMeParams, session *app.UserSession) op.UserMeResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	user, err := srv.app.GetUserByID(ctx, session.UserID)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errUserMe(log, err, codeInternal)
	case app.ErrNotFound:
		return errUserMe(log, err, codeNotFound)
	case nil:
	}

	return op.NewUserMeOK().WithPayload(apiUser(user))
}

func (srv *server) userChangePassword(params op.UserChangePasswordParams, session *app.UserSession) op.UserChangePasswordResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	newPass := swag.StringValue(params.Args.NewPassword)
	repeatPass := swag.StringValue(params.Args.RepeatPassword)
	if newPass != repeatPass {
		return errUserChangePassword(log, errors.New("passwords not equals"), codePasswordsNotEquals)
	}

	err := srv.app.UpdatePassword(ctx, session.UserID, app.ChangePassword{
		PreviousPassword: swag.StringValue(params.Args.OldPassword),
		NewPassword:      swag.StringValue(params.Args.NewPassword),
	})
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errUserChangePassword(log, err, codeInternal)
	case app.ErrPasswordNotCorrect:
		return errUserChangePassword(log, err, codeWrongPassword)
	case nil:
	}

	return op.NewUserChangePasswordNoContent()
}

func (srv *server) userChangeProfile(params op.UserChangeProfileParams, session *app.UserSession) op.UserChangeProfileResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	id := params.Args.UserID
	if id <= 0 {
		id = session.UserID
	}
	err := srv.app.UpdateProfile(
		ctx,
		id,
		swag.StringValue(params.Args.Email),
		swag.StringValue(params.Args.FirstName),
		swag.StringValue(params.Args.SecondName),
		int64(getUserRoleByString(swag.StringValue(params.Args.Role))),
		int64(getUserStatusByString(swag.StringValue(params.Args.Status))),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errUserChangeProfile(log, err, codeInternal)
	case app.ErrNotFound:
		return errUserChangeProfile(log, err, codeNotFound)
	case nil:
	}

	return op.NewUserChangeProfileNoContent()
}

func (srv *server) listUser(params op.ListUserParams, session *app.UserSession) op.ListUserResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	entities, err := srv.app.GetUsers(ctx)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errListUser(log, err, codeInternal)
	case nil:
	}

	return op.NewListUserOK().WithPayload(apiArrayUser(entities))
}
