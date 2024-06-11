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

// JobsAttempt jobs attempt
//
// swagger:model jobs.Attempt
type JobsAttempt struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// execution count
	ExecutionCount int64 `json:"executionCount,omitempty"`

	// finished at
	FinishedAt string `json:"finishedAt,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`

	// started at
	StartedAt string `json:"startedAt,omitempty"`

	// state
	State *JobsState `json:"state,omitempty"`

	// tries
	Tries []*JobsTry `json:"tries"`
}

// Validate validates this jobs attempt
func (m *JobsAttempt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsAttempt) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAttempt) validateTries(formats strfmt.Registry) error {
	if swag.IsZero(m.Tries) { // not required
		return nil
	}

	for i := 0; i < len(m.Tries); i++ {
		if swag.IsZero(m.Tries[i]) { // not required
			continue
		}

		if m.Tries[i] != nil {
			if err := m.Tries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this jobs attempt based on the context it is used
func (m *JobsAttempt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsAttempt) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAttempt) contextValidateTries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tries); i++ {

		if m.Tries[i] != nil {

			if swag.IsZero(m.Tries[i]) { // not required
				return nil
			}

			if err := m.Tries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsAttempt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsAttempt) UnmarshalBinary(b []byte) error {
	var res JobsAttempt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
