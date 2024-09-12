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

// TestsTrigger tests trigger
//
// swagger:model tests.Trigger
type TestsTrigger struct {

	// execution template
	ExecutionTemplate *TestexecutionsExecutionTemplate `json:"executionTemplate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sandbox of
	SandboxOf *SandboxForkOf `json:"sandboxOf,omitempty"`
}

// Validate validates this tests trigger
func (m *TestsTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSandboxOf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestsTrigger) validateExecutionTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecutionTemplate) { // not required
		return nil
	}

	if m.ExecutionTemplate != nil {
		if err := m.ExecutionTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("executionTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("executionTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *TestsTrigger) validateSandboxOf(formats strfmt.Registry) error {
	if swag.IsZero(m.SandboxOf) { // not required
		return nil
	}

	if m.SandboxOf != nil {
		if err := m.SandboxOf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sandboxOf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sandboxOf")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tests trigger based on the context it is used
func (m *TestsTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExecutionTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSandboxOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestsTrigger) contextValidateExecutionTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ExecutionTemplate != nil {

		if swag.IsZero(m.ExecutionTemplate) { // not required
			return nil
		}

		if err := m.ExecutionTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("executionTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("executionTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *TestsTrigger) contextValidateSandboxOf(ctx context.Context, formats strfmt.Registry) error {

	if m.SandboxOf != nil {

		if swag.IsZero(m.SandboxOf) { // not required
			return nil
		}

		if err := m.SandboxOf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sandboxOf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sandboxOf")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestsTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestsTrigger) UnmarshalBinary(b []byte) error {
	var res TestsTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}