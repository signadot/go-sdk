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

// JobsState jobs state
//
// swagger:model jobs.State
type JobsState struct {

	// canceled
	Canceled *JobsCanceledState `json:"canceled,omitempty"`

	// completed
	Completed *JobsCompletedState `json:"completed,omitempty"`

	// running
	Running *JobsRunningState `json:"running,omitempty"`
}

// Validate validates this jobs state
func (m *JobsState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanceled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsState) validateCanceled(formats strfmt.Registry) error {
	if swag.IsZero(m.Canceled) { // not required
		return nil
	}

	if m.Canceled != nil {
		if err := m.Canceled.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("canceled")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("canceled")
			}
			return err
		}
	}

	return nil
}

func (m *JobsState) validateCompleted(formats strfmt.Registry) error {
	if swag.IsZero(m.Completed) { // not required
		return nil
	}

	if m.Completed != nil {
		if err := m.Completed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("completed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("completed")
			}
			return err
		}
	}

	return nil
}

func (m *JobsState) validateRunning(formats strfmt.Registry) error {
	if swag.IsZero(m.Running) { // not required
		return nil
	}

	if m.Running != nil {
		if err := m.Running.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("running")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("running")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jobs state based on the context it is used
func (m *JobsState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCanceled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCompleted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsState) contextValidateCanceled(ctx context.Context, formats strfmt.Registry) error {

	if m.Canceled != nil {

		if swag.IsZero(m.Canceled) { // not required
			return nil
		}

		if err := m.Canceled.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("canceled")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("canceled")
			}
			return err
		}
	}

	return nil
}

func (m *JobsState) contextValidateCompleted(ctx context.Context, formats strfmt.Registry) error {

	if m.Completed != nil {

		if swag.IsZero(m.Completed) { // not required
			return nil
		}

		if err := m.Completed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("completed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("completed")
			}
			return err
		}
	}

	return nil
}

func (m *JobsState) contextValidateRunning(ctx context.Context, formats strfmt.Registry) error {

	if m.Running != nil {

		if swag.IsZero(m.Running) { // not required
			return nil
		}

		if err := m.Running.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("running")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("running")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsState) UnmarshalBinary(b []byte) error {
	var res JobsState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
