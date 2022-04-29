// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SandboxResource sandbox resource
//
// swagger:model SandboxResource
type SandboxResource struct {

	// name
	Name string `json:"name,omitempty"`

	// params
	Params map[string]string `json:"params,omitempty"`

	// plugin
	Plugin string `json:"plugin,omitempty"`
}

// Validate validates this sandbox resource
func (m *SandboxResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sandbox resource based on context it is used
func (m *SandboxResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SandboxResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxResource) UnmarshalBinary(b []byte) error {
	var res SandboxResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
