package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/model"
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
)

func (srv *server) supportedLanguages(params op.SupportedLanguagesParams, session *app.UserSession) op.SupportedLanguagesResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	languages, err := srv.app.GetSupportedLanguages(ctx)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errSupportedLanguages(log, err, codeInternal)
	case app.ErrAutoTranslation:
		return errSupportedLanguages(log, err, codeAutoTranslationNotSupported)
	case nil:
	}

	l := make([]*model.Language, len(languages))
	for i := range languages {
		l[i] = &model.Language{
			Name: pointer.ToString(languages[i].Name),
			Code: pointer.ToString(languages[i].Code),
		}
	}

	return op.NewSupportedLanguagesOK().WithPayload(l)
}
