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

// DeleteTranslationFileNoContentCode is the HTTP code returned for type DeleteTranslationFileNoContent
const DeleteTranslationFileNoContentCode int = 204

/*DeleteTranslationFileNoContent No content in answer

swagger:response deleteTranslationFileNoContent
*/
type DeleteTranslationFileNoContent struct {
}

// NewDeleteTranslationFileNoContent creates DeleteTranslationFileNoContent with default headers values
func NewDeleteTranslationFileNoContent() *DeleteTranslationFileNoContent {

	return &DeleteTranslationFileNoContent{}
}

// WriteResponse to the client
func (o *DeleteTranslationFileNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *DeleteTranslationFileNoContent) DeleteTranslationFileResponder() {}

/*DeleteTranslationFileDefault General errors using same model as used by go-swagger for validation errors.

swagger:response deleteTranslationFileDefault
*/
type DeleteTranslationFileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewDeleteTranslationFileDefault creates DeleteTranslationFileDefault with default headers values
func NewDeleteTranslationFileDefault(code int) *DeleteTranslationFileDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTranslationFileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete translation file default response
func (o *DeleteTranslationFileDefault) WithStatusCode(code int) *DeleteTranslationFileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete translation file default response
func (o *DeleteTranslationFileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete translation file default response
func (o *DeleteTranslationFileDefault) WithPayload(payload *model.Error) *DeleteTranslationFileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete translation file default response
func (o *DeleteTranslationFileDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTranslationFileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *DeleteTranslationFileDefault) DeleteTranslationFileResponder() {}

type DeleteTranslationFileNotImplementedResponder struct {
	middleware.Responder
}

func (*DeleteTranslationFileNotImplementedResponder) DeleteTranslationFileResponder() {}

func DeleteTranslationFileNotImplemented() DeleteTranslationFileResponder {
	return &DeleteTranslationFileNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.DeleteTranslationFile has not yet been implemented",
		),
	}
}

type DeleteTranslationFileResponder interface {
	middleware.Responder
	DeleteTranslationFileResponder()
}
