// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/MarcSky/freon/internal/app"
)

// UpdateTranslationHandlerFunc turns a function with the right signature into a update translation handler
type UpdateTranslationHandlerFunc func(UpdateTranslationParams, *app.UserSession) UpdateTranslationResponder

// Handle executing the request and returning a response
func (fn UpdateTranslationHandlerFunc) Handle(params UpdateTranslationParams, principal *app.UserSession) UpdateTranslationResponder {
	return fn(params, principal)
}

// UpdateTranslationHandler interface for that can handle valid update translation params
type UpdateTranslationHandler interface {
	Handle(UpdateTranslationParams, *app.UserSession) UpdateTranslationResponder
}

// NewUpdateTranslation creates a new http.Handler for the update translation operation
func NewUpdateTranslation(ctx *middleware.Context, handler UpdateTranslationHandler) *UpdateTranslation {
	return &UpdateTranslation{Context: ctx, Handler: handler}
}

/*UpdateTranslation swagger:route PUT /translation/{id} updateTranslation

update translation

*/
type UpdateTranslation struct {
	Context *middleware.Context
	Handler UpdateTranslationHandler
}

func (o *UpdateTranslation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateTranslationParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
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

// UpdateTranslationBody update translation body
//
// swagger:model UpdateTranslationBody
type UpdateTranslationBody struct {

	// text
	// Required: true
	Text *string `json:"text"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *UpdateTranslationBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// text
		// Required: true
		Text *string `json:"text"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Text = props.Text
	return nil
}

// Validate validates this update translation body
func (o *UpdateTranslationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateTranslationBody) validateText(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"text", "body", o.Text); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateTranslationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateTranslationBody) UnmarshalBinary(b []byte) error {
	var res UpdateTranslationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
