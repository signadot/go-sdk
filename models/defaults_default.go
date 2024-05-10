// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DefaultsDefault defaults default
//
// swagger:model defaults.Default
type DefaultsDefault struct {

	// class
	Class string `json:"class,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// resource kind
	ResourceKind string `json:"resourceKind,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this defaults default
func (m *DefaultsDefault) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this defaults default based on context it is used
func (m *DefaultsDefault) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DefaultsDefault) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefaultsDefault) UnmarshalBinary(b []byte) error {
	var res DefaultsDefault
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
