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

// SettingStorageNoContentCode is the HTTP code returned for type SettingStorageNoContent
const SettingStorageNoContentCode int = 204

/*SettingStorageNoContent No content in answer

swagger:response settingStorageNoContent
*/
type SettingStorageNoContent struct {
}

// NewSettingStorageNoContent creates SettingStorageNoContent with default headers values
func NewSettingStorageNoContent() *SettingStorageNoContent {

	return &SettingStorageNoContent{}
}

// WriteResponse to the client
func (o *SettingStorageNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *SettingStorageNoContent) SettingStorageResponder() {}

/*SettingStorageDefault General errors using same model as used by go-swagger for validation errors.

swagger:response settingStorageDefault
*/
type SettingStorageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewSettingStorageDefault creates SettingStorageDefault with default headers values
func NewSettingStorageDefault(code int) *SettingStorageDefault {
	if code <= 0 {
		code = 500
	}

	return &SettingStorageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the setting storage default response
func (o *SettingStorageDefault) WithStatusCode(code int) *SettingStorageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the setting storage default response
func (o *SettingStorageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the setting storage default response
func (o *SettingStorageDefault) WithPayload(payload *model.Error) *SettingStorageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the setting storage default response
func (o *SettingStorageDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SettingStorageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *SettingStorageDefault) SettingStorageResponder() {}

type SettingStorageNotImplementedResponder struct {
	middleware.Responder
}

func (*SettingStorageNotImplementedResponder) SettingStorageResponder() {}

func SettingStorageNotImplemented() SettingStorageResponder {
	return &SettingStorageNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.SettingStorage has not yet been implemented",
		),
	}
}

type SettingStorageResponder interface {
	middleware.Responder
	SettingStorageResponder()
}
