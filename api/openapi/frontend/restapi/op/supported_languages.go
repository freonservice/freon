// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/freonservice/freon/internal/app"
)

// SupportedLanguagesHandlerFunc turns a function with the right signature into a supported languages handler
type SupportedLanguagesHandlerFunc func(SupportedLanguagesParams, *app.UserSession) SupportedLanguagesResponder

// Handle executing the request and returning a response
func (fn SupportedLanguagesHandlerFunc) Handle(params SupportedLanguagesParams, principal *app.UserSession) SupportedLanguagesResponder {
	return fn(params, principal)
}

// SupportedLanguagesHandler interface for that can handle valid supported languages params
type SupportedLanguagesHandler interface {
	Handle(SupportedLanguagesParams, *app.UserSession) SupportedLanguagesResponder
}

// NewSupportedLanguages creates a new http.Handler for the supported languages operation
func NewSupportedLanguages(ctx *middleware.Context, handler SupportedLanguagesHandler) *SupportedLanguages {
	return &SupportedLanguages{Context: ctx, Handler: handler}
}

/* SupportedLanguages swagger:route GET /supported-languages supportedLanguages

supported-languages

*/
type SupportedLanguages struct {
	Context *middleware.Context
	Handler SupportedLanguagesHandler
}

func (o *SupportedLanguages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSupportedLanguagesParams()
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
