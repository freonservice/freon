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

// Category category
//
// swagger:model Category
type Category struct {

	// id
	// Required: true
	ID *int64 `json:"id"`

	// name
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Name *string `json:"name"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Category) UnmarshalJSON(data []byte) error {
	var props struct {

		// id
		// Required: true
		ID *int64 `json:"id"`

		// name
		// Required: true
		// Max Length: 255
		// Min Length: 1
		Name *string `json:"name"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.ID = props.ID
	m.Name = props.Name
	return nil
}

// Validate validates this category
func (m *Category) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Category) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Category) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 255); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this category based on context it is used
func (m *Category) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Category) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Category) UnmarshalBinary(b []byte) error {
	var res Category
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
