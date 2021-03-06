// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/freonservice/freon/internal/app"
)

// InfoHandlerFunc turns a function with the right signature into a info handler
type InfoHandlerFunc func(InfoParams, *app.UserSession) InfoResponder

// Handle executing the request and returning a response
func (fn InfoHandlerFunc) Handle(params InfoParams, principal *app.UserSession) InfoResponder {
	return fn(params, principal)
}

// InfoHandler interface for that can handle valid info params
type InfoHandler interface {
	Handle(InfoParams, *app.UserSession) InfoResponder
}

// NewInfo creates a new http.Handler for the info operation
func NewInfo(ctx *middleware.Context, handler InfoHandler) *Info {
	return &Info{Context: ctx, Handler: handler}
}

/* Info swagger:route GET /info info

getting actual user session and system configurations

*/
type Info struct {
	Context *middleware.Context
	Handler InfoHandler
}

func (o *Info) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewInfoParams()
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
