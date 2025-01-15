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

// TestExecutionStatus test execution status
//
// swagger:model TestExecutionStatus
type TestExecutionStatus struct {

	// baseline job
	BaselineJob string `json:"baselineJob,omitempty"`

	// final state
	FinalState *TestExecutionState `json:"finalState,omitempty"`

	// finished at
	FinishedAt string `json:"finishedAt,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`

	// started at
	StartedAt string `json:"startedAt,omitempty"`

	// target job
	TargetJob string `json:"targetJob,omitempty"`

	// test deleted at
	TestDeletedAt string `json:"testDeletedAt,omitempty"`

	// triggered by
	TriggeredBy *TestExecutionTriggeredBy `json:"triggeredBy,omitempty"`
}

// Validate validates this test execution status
func (m *TestExecutionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinalState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggeredBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestExecutionStatus) validateFinalState(formats strfmt.Registry) error {
	if swag.IsZero(m.FinalState) { // not required
		return nil
	}

	if m.FinalState != nil {
		if err := m.FinalState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("finalState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("finalState")
			}
			return err
		}
	}

	return nil
}

func (m *TestExecutionStatus) validateTriggeredBy(formats strfmt.Registry) error {
	if swag.IsZero(m.TriggeredBy) { // not required
		return nil
	}

	if m.TriggeredBy != nil {
		if err := m.TriggeredBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggeredBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("triggeredBy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this test execution status based on the context it is used
func (m *TestExecutionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFinalState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTriggeredBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestExecutionStatus) contextValidateFinalState(ctx context.Context, formats strfmt.Registry) error {

	if m.FinalState != nil {

		if swag.IsZero(m.FinalState) { // not required
			return nil
		}

		if err := m.FinalState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("finalState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("finalState")
			}
			return err
		}
	}

	return nil
}

func (m *TestExecutionStatus) contextValidateTriggeredBy(ctx context.Context, formats strfmt.Registry) error {

	if m.TriggeredBy != nil {

		if swag.IsZero(m.TriggeredBy) { // not required
			return nil
		}

		if err := m.TriggeredBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggeredBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("triggeredBy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestExecutionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestExecutionStatus) UnmarshalBinary(b []byte) error {
	var res TestExecutionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
