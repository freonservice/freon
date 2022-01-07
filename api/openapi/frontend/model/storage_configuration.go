// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StorageConfiguration storage configuration
//
// swagger:model StorageConfiguration
type StorageConfiguration struct {

	// use
	// Required: true
	Use *int32 `json:"use"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *StorageConfiguration) UnmarshalJSON(data []byte) error {
	var props struct {

		// use
		// Required: true
		Use *int32 `json:"use"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Use = props.Use
	return nil
}

// Validate validates this storage configuration
func (m *StorageConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageConfiguration) validateUse(formats strfmt.Registry) error {

	if err := validate.Required("use", "body", m.Use); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this storage configuration based on context it is used
func (m *StorageConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageConfiguration) UnmarshalBinary(b []byte) error {
	var res StorageConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
