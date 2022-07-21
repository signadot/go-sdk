// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SandboxEnvVar sandbox env var
//
// swagger:model sandbox.EnvVar
type SandboxEnvVar struct {

	// name of container to which it applies
	Container string `json:"container,omitempty"`

	// environmental variable name
	Name string `json:"name,omitempty"`

	// upsert or delete
	Operation string `json:"operation,omitempty"`

	// environmental variable value
	Value string `json:"value,omitempty"`

	// value from
	ValueFrom *SandboxEnvValueFrom `json:"valueFrom,omitempty"`
}

// Validate validates this sandbox env var
func (m *SandboxEnvVar) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValueFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SandboxEnvVar) validateValueFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.ValueFrom) { // not required
		return nil
	}

	if m.ValueFrom != nil {
		if err := m.ValueFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueFrom")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sandbox env var based on the context it is used
func (m *SandboxEnvVar) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValueFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SandboxEnvVar) contextValidateValueFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.ValueFrom != nil {
		if err := m.ValueFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueFrom")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SandboxEnvVar) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxEnvVar) UnmarshalBinary(b []byte) error {
	var res SandboxEnvVar
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}