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

// TestExecutionTriggeredBy test execution triggered by
//
// swagger:model TestExecutionTriggeredBy
type TestExecutionTriggeredBy struct {

	// sandbox
	Sandbox string `json:"sandbox,omitempty"`

	// trigger
	Trigger *TestTrigger `json:"trigger,omitempty"`
}

// Validate validates this test execution triggered by
func (m *TestExecutionTriggeredBy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestExecutionTriggeredBy) validateTrigger(formats strfmt.Registry) error {
	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if m.Trigger != nil {
		if err := m.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this test execution triggered by based on the context it is used
func (m *TestExecutionTriggeredBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrigger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestExecutionTriggeredBy) contextValidateTrigger(ctx context.Context, formats strfmt.Registry) error {

	if m.Trigger != nil {

		if swag.IsZero(m.Trigger) { // not required
			return nil
		}

		if err := m.Trigger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestExecutionTriggeredBy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestExecutionTriggeredBy) UnmarshalBinary(b []byte) error {
	var res TestExecutionTriggeredBy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
