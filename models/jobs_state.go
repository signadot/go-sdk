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

	// failed
	Failed *JobsFailedState `json:"failed,omitempty"`

	// queued
	Queued *JobsQueuedState `json:"queued,omitempty"`

	// running
	Running *JobsRunningState `json:"running,omitempty"`

	// succeeded
	Succeeded *JobsSucceededState `json:"succeeded,omitempty"`
}

// Validate validates this jobs state
func (m *JobsState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanceled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSucceeded(formats); err != nil {
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

func (m *JobsState) validateFailed(formats strfmt.Registry) error {
	if swag.IsZero(m.Failed) { // not required
		return nil
	}

	if m.Failed != nil {
		if err := m.Failed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("failed")
			}
			return err
		}
	}

	return nil
}

func (m *JobsState) validateQueued(formats strfmt.Registry) error {
	if swag.IsZero(m.Queued) { // not required
		return nil
	}

	if m.Queued != nil {
		if err := m.Queued.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queued")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queued")
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

func (m *JobsState) validateSucceeded(formats strfmt.Registry) error {
	if swag.IsZero(m.Succeeded) { // not required
		return nil
	}

	if m.Succeeded != nil {
		if err := m.Succeeded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("succeeded")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("succeeded")
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

	if err := m.contextValidateFailed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueued(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSucceeded(ctx, formats); err != nil {
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

func (m *JobsState) contextValidateFailed(ctx context.Context, formats strfmt.Registry) error {

	if m.Failed != nil {

		if swag.IsZero(m.Failed) { // not required
			return nil
		}

		if err := m.Failed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("failed")
			}
			return err
		}
	}

	return nil
}

func (m *JobsState) contextValidateQueued(ctx context.Context, formats strfmt.Registry) error {

	if m.Queued != nil {

		if swag.IsZero(m.Queued) { // not required
			return nil
		}

		if err := m.Queued.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queued")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queued")
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

func (m *JobsState) contextValidateSucceeded(ctx context.Context, formats strfmt.Registry) error {

	if m.Succeeded != nil {

		if swag.IsZero(m.Succeeded) { // not required
			return nil
		}

		if err := m.Succeeded.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("succeeded")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("succeeded")
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
