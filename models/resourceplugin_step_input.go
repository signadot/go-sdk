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

	// as
	As *ResourcepluginStepInputTo `json:"as,omitempty"`

	// Name for the input
	Name string `json:"name,omitempty"`

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

	if err := m.validateAs(formats); err != nil {
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

func (m *ResourcepluginStepInput) validateAs(formats strfmt.Registry) error {
	if swag.IsZero(m.As) { // not required
		return nil
	}

	if m.As != nil {
		if err := m.As.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("as")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("as")
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

	if err := m.contextValidateAs(ctx, formats); err != nil {
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

func (m *ResourcepluginStepInput) contextValidateAs(ctx context.Context, formats strfmt.Registry) error {

	if m.As != nil {

		if swag.IsZero(m.As) { // not required
			return nil
		}

		if err := m.As.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("as")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("as")
			}
			return err
		}
	}

	return nil
}

func (m *ResourcepluginStepInput) contextValidateValueFromStep(ctx context.Context, formats strfmt.Registry) error {

	if m.ValueFromStep != nil {

		if swag.IsZero(m.ValueFromStep) { // not required
			return nil
		}

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
