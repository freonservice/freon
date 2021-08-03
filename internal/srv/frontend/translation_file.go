package frontend

import (
	"sync"

	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/internal/filter"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) createTranslationFile(
	params op.CreateTranslationFilesParams, session *app.UserSession) op.CreateTranslationFilesResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	size := len(params.Args.Platforms)
	wg := sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		i := i
		go func() {
			defer wg.Done()

			err := srv.app.CreateTranslationFile(
				ctx,
				params.Args.Platforms[i],
				swag.StringValue(params.Args.StorageType),
				session.UserID,
				swag.Int64Value(params.Args.LocalizationID),
			)
			switch errors.Cause(err) {
			default:
				log.PrintErr(errors.WithStack(err))
			case nil:
			}
		}()
	}
	wg.Wait()

	return op.NewCreateTranslationFilesNoContent()
}

func (srv *server) listTranslationFiles(params op.ListTranslationFilesParams, session *app.UserSession) op.ListTranslationFilesResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	entities, err := srv.app.GetTranslationFiles(
		ctx,
		filter.TranslationFileFilter{
			LocalizationID: swag.Int64Value(params.LocalizationID),
			PlatformType:   getPlatformByString(swag.StringValue(params.Platform)),
		},
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errListTranslationFiles(log, err, codeInternal)
	case nil:
	}

	return op.NewListTranslationFilesOK().WithPayload(apiArrayTranslationFiles(entities))
}

func (srv *server) deleteTranslationFile(
	params op.DeleteTranslationFileParams, session *app.UserSession) op.DeleteTranslationFileResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)

	err := srv.app.DeleteTranslationFile(ctx, params.ID)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errDeleteTranslationFile(log, err, codeInternal)
	case nil:
	}

	return op.NewDeleteTranslationFileNoContent()
}
