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

// ResourcepluginStepInput resourceplugin step input
//
// swagger:model resourceplugin.StepInput
type ResourcepluginStepInput struct {

	// Name for the input
	Name string `json:"name,omitempty"`

	// to
	To *ResourcepluginStepInputTo `json:"to,omitempty"`

	// Type of input
	Type string `json:"type,omitempty"`

	// ValueFromSandbox defines whether or not to source value from the sandbox spec
	ValueFromSandbox bool `json:"valueFromSandbox,omitempty"`

	// value from step
	ValueFromStep *ResourcepluginValueFromStep `json:"valueFromStep,omitempty"`
}

// Validate validates this resourceplugin step input
func (m *ResourcepluginStepInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueFromStep(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcepluginStepInput) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if m.To != nil {
		if err := m.To.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

func (m *ResourcepluginStepInput) validateValueFromStep(formats strfmt.Registry) error {
	if swag.IsZero(m.ValueFromStep) { // not required
		return nil
	}

	if m.ValueFromStep != nil {
		if err := m.ValueFromStep.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueFromStep")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueFromStep")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this resourceplugin step input based on the context it is used
func (m *ResourcepluginStepInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValueFromStep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcepluginStepInput) contextValidateTo(ctx context.Context, formats strfmt.Registry) error {

	if m.To != nil {
		if err := m.To.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

func (m *ResourcepluginStepInput) contextValidateValueFromStep(ctx context.Context, formats strfmt.Registry) error {

	if m.ValueFromStep != nil {
		if err := m.ValueFromStep.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueFromStep")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueFromStep")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcepluginStepInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcepluginStepInput) UnmarshalBinary(b []byte) error {
	var res ResourcepluginStepInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
