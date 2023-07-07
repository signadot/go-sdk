// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourcepluginStep resourceplugin step
//
// swagger:model resourceplugin.Step
type ResourcepluginStep struct {

	// Inputs for the step
	Inputs []*ResourcepluginStepInput `json:"inputs"`

	// Name for the step
	Name string `json:"name,omitempty"`

	// Outputs for the step
	Outputs []*ResourcepluginStepOut `json:"outputs"`

	// Script to execute in the step
	Script string `json:"script,omitempty"`
}

// Validate validates this resourceplugin step
func (m *ResourcepluginStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcepluginStep) validateInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourcepluginStep) validateOutputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Outputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Outputs); i++ {
		if swag.IsZero(m.Outputs[i]) { // not required
			continue
		}

		if m.Outputs[i] != nil {
			if err := m.Outputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this resourceplugin step based on the context it is used
func (m *ResourcepluginStep) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcepluginStep) contextValidateInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Inputs); i++ {

		if m.Inputs[i] != nil {

			if swag.IsZero(m.Inputs[i]) { // not required
				return nil
			}

			if err := m.Inputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourcepluginStep) contextValidateOutputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Outputs); i++ {

		if m.Outputs[i] != nil {

			if swag.IsZero(m.Outputs[i]) { // not required
				return nil
			}

			if err := m.Outputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcepluginStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcepluginStep) UnmarshalBinary(b []byte) error {
	var res ResourcepluginStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
