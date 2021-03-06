// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SandboxEnvValueFromResource sandbox env value from resource
//
// swagger:model sandbox.EnvValueFromResource
type SandboxEnvValueFromResource struct {

	// name
	Name string `json:"name,omitempty"`

	// output key
	OutputKey string `json:"outputKey,omitempty"`
}

// Validate validates this sandbox env value from resource
func (m *SandboxEnvValueFromResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sandbox env value from resource based on context it is used
func (m *SandboxEnvValueFromResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SandboxEnvValueFromResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxEnvValueFromResource) UnmarshalBinary(b []byte) error {
	var res SandboxEnvValueFromResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
