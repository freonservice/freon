// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/freonservice/freon/internal/app"
)

// ListLocalizationHandlerFunc turns a function with the right signature into a list localization handler
type ListLocalizationHandlerFunc func(ListLocalizationParams, *app.UserSession) ListLocalizationResponder

// Handle executing the request and returning a response
func (fn ListLocalizationHandlerFunc) Handle(params ListLocalizationParams, principal *app.UserSession) ListLocalizationResponder {
	return fn(params, principal)
}

// ListLocalizationHandler interface for that can handle valid list localization params
type ListLocalizationHandler interface {
	Handle(ListLocalizationParams, *app.UserSession) ListLocalizationResponder
}

// NewListLocalization creates a new http.Handler for the list localization operation
func NewListLocalization(ctx *middleware.Context, handler ListLocalizationHandler) *ListLocalization {
	return &ListLocalization{Context: ctx, Handler: handler}
}

/* ListLocalization swagger:route GET /localizations listLocalization

get list localization sorted by user permission

*/
type ListLocalization struct {
	Context *middleware.Context
	Handler ListLocalizationHandler
}

func (o *ListLocalization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListLocalizationParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *app.UserSession
	if uprinc != nil {
		principal = uprinc.(*app.UserSession) // this is really a app.UserSession, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
