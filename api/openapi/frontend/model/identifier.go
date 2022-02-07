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

// Identifier identifier
//
// swagger:model Identifier
type Identifier struct {

	// category
	Category *Category `json:"category,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// name
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Name *string `json:"name"`

	// platforms
	// Required: true
	Platforms []string `json:"platforms"`

	// text plural
	TextPlural string `json:"text_plural,omitempty"`

	// text singular
	TextSingular string `json:"text_singular,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Identifier) UnmarshalJSON(data []byte) error {
	var props struct {

		// category
		Category *Category `json:"category,omitempty"`

		// description
		Description string `json:"description,omitempty"`

		// id
		// Required: true
		ID *int64 `json:"id"`

		// name
		// Required: true
		// Max Length: 255
		// Min Length: 1
		Name *string `json:"name"`

		// platforms
		// Required: true
		Platforms []string `json:"platforms"`

		// text plural
		TextPlural string `json:"text_plural,omitempty"`

		// text singular
		TextSingular string `json:"text_singular,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Category = props.Category
	m.Description = props.Description
	m.ID = props.ID
	m.Name = props.Name
	m.Platforms = props.Platforms
	m.TextPlural = props.TextPlural
	m.TextSingular = props.TextSingular
	return nil
}

// Validate validates this identifier
func (m *Identifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Identifier) validateCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *Identifier) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Identifier) validateName(formats strfmt.Registry) error {

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

func (m *Identifier) validatePlatforms(formats strfmt.Registry) error {

	if err := validate.Required("platforms", "body", m.Platforms); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this identifier based on the context it is used
func (m *Identifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Identifier) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.Category != nil {
		if err := m.Category.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Identifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identifier) UnmarshalBinary(b []byte) error {
	var res Identifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
