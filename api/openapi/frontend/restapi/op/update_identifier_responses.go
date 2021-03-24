// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/MarcSky/freon/api/openapi/frontend/model"
)

// UpdateIdentifierNoContentCode is the HTTP code returned for type UpdateIdentifierNoContent
const UpdateIdentifierNoContentCode int = 204

/*UpdateIdentifierNoContent No content in answer

swagger:response updateIdentifierNoContent
*/
type UpdateIdentifierNoContent struct {
}

// NewUpdateIdentifierNoContent creates UpdateIdentifierNoContent with default headers values
func NewUpdateIdentifierNoContent() *UpdateIdentifierNoContent {

	return &UpdateIdentifierNoContent{}
}

// WriteResponse to the client
func (o *UpdateIdentifierNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *UpdateIdentifierNoContent) UpdateIdentifierResponder() {}

/*UpdateIdentifierDefault General errors using same model as used by go-swagger for validation errors.

swagger:response updateIdentifierDefault
*/
type UpdateIdentifierDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewUpdateIdentifierDefault creates UpdateIdentifierDefault with default headers values
func NewUpdateIdentifierDefault(code int) *UpdateIdentifierDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateIdentifierDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update identifier default response
func (o *UpdateIdentifierDefault) WithStatusCode(code int) *UpdateIdentifierDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update identifier default response
func (o *UpdateIdentifierDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update identifier default response
func (o *UpdateIdentifierDefault) WithPayload(payload *model.Error) *UpdateIdentifierDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update identifier default response
func (o *UpdateIdentifierDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateIdentifierDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateIdentifierDefault) UpdateIdentifierResponder() {}

type UpdateIdentifierNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateIdentifierNotImplementedResponder) UpdateIdentifierResponder() {}

func UpdateIdentifierNotImplemented() UpdateIdentifierResponder {
	return &UpdateIdentifierNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UpdateIdentifier has not yet been implemented",
		),
	}
}

type UpdateIdentifierResponder interface {
	middleware.Responder
	UpdateIdentifierResponder()
}
