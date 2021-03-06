// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/freonservice/freon/api/openapi/frontend/model"
)

// CreateCategoryNoContentCode is the HTTP code returned for type CreateCategoryNoContent
const CreateCategoryNoContentCode int = 204

/*CreateCategoryNoContent No content in answer

swagger:response createCategoryNoContent
*/
type CreateCategoryNoContent struct {
}

// NewCreateCategoryNoContent creates CreateCategoryNoContent with default headers values
func NewCreateCategoryNoContent() *CreateCategoryNoContent {

	return &CreateCategoryNoContent{}
}

// WriteResponse to the client
func (o *CreateCategoryNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *CreateCategoryNoContent) CreateCategoryResponder() {}

/*CreateCategoryDefault General errors using same model as used by go-swagger for validation errors.

swagger:response createCategoryDefault
*/
type CreateCategoryDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewCreateCategoryDefault creates CreateCategoryDefault with default headers values
func NewCreateCategoryDefault(code int) *CreateCategoryDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateCategoryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create category default response
func (o *CreateCategoryDefault) WithStatusCode(code int) *CreateCategoryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create category default response
func (o *CreateCategoryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create category default response
func (o *CreateCategoryDefault) WithPayload(payload *model.Error) *CreateCategoryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create category default response
func (o *CreateCategoryDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCategoryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateCategoryDefault) CreateCategoryResponder() {}

type CreateCategoryNotImplementedResponder struct {
	middleware.Responder
}

func (*CreateCategoryNotImplementedResponder) CreateCategoryResponder() {}

func CreateCategoryNotImplemented() CreateCategoryResponder {
	return &CreateCategoryNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CreateCategory has not yet been implemented",
		),
	}
}

type CreateCategoryResponder interface {
	middleware.Responder
	CreateCategoryResponder()
}
