package frontend

import (
	"github.com/MarcSky/freon/api/openapi/frontend/restapi/op"
	"github.com/MarcSky/freon/internal/app"
	"github.com/MarcSky/freon/internal/dal"

	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (srv *server) createCategory(params op.CreateCategoryParams, session *app.UserSession) op.CreateCategoryResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	err := srv.app.CreateCategory(
		ctx,
		swag.StringValue(params.Args.Name),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errCreateCategory(log, err, codeInternal)
	case dal.ErrDuplicateKeyValue:
		return errCreateCategory(log, err, codeCategoryIsExist)
	case nil:
	}

	return op.NewCreateCategoryNoContent()
}

func (srv *server) listCategories(params op.ListCategoriesParams, session *app.UserSession) op.ListCategoriesResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	entities, err := srv.app.GetCategories(ctx)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errListCategories(log, err, codeInternal)
	case nil:
	}

	return op.NewListCategoriesOK().WithPayload(apiArrayCategory(entities))
}

func (srv *server) deleteCategory(params op.DeleteCategoryParams, session *app.UserSession) op.DeleteCategoryResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	err := srv.app.DeleteCategory(
		ctx,
		params.ID,
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errDeleteCategory(log, err, codeInternal)
	case nil:
	}

	return op.NewDeleteCategoryNoContent()
}

func (srv *server) updateCategory(params op.UpdateCategoryParams, session *app.UserSession) op.UpdateCategoryResponder {
	ctx, log := fromRequest(params.HTTPRequest, nil)

	err := srv.app.UpdateCategory(
		ctx,
		params.ID,
		swag.StringValue(params.Args.Name),
	)
	switch errors.Cause(err) {
	default:
		log.PrintErr(errors.WithStack(err))
		return errUpdateCategory(log, err, codeInternal)
	case nil:
	}

	return op.NewUpdateCategoryNoContent()
}
