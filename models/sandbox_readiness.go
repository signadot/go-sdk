// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SandboxReadiness sandbox readiness
//
// swagger:model SandboxReadiness
type SandboxReadiness struct {

	// Message is a human readable explanation of why
	// the sandbox is healthy or not.
	Message string `json:"message,omitempty"`

	// Ready indicates whether the sandbox is ready,
	// meaning that it can be used for testing.
	Ready bool `json:"ready,omitempty"`

	// Reason is a machine readable explanation of why
	// the sandbox is healthy or not.
	Reason string `json:"reason,omitempty"`
}

// Validate validates this sandbox readiness
func (m *SandboxReadiness) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sandbox readiness based on context it is used
func (m *SandboxReadiness) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SandboxReadiness) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxReadiness) UnmarshalBinary(b []byte) error {
	var res SandboxReadiness
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
