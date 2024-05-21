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

// JobsStatus jobs status
//
// swagger:model jobs.Status
type JobsStatus struct {

	// attempts
	Attempts []*JobsAttempt `json:"attempts"`

	// finished at
	FinishedAt string `json:"finishedAt,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`

	// started at
	StartedAt string `json:"startedAt,omitempty"`
}

// Validate validates this jobs status
func (m *JobsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttempts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsStatus) validateAttempts(formats strfmt.Registry) error {
	if swag.IsZero(m.Attempts) { // not required
		return nil
	}

	for i := 0; i < len(m.Attempts); i++ {
		if swag.IsZero(m.Attempts[i]) { // not required
			continue
		}

		if m.Attempts[i] != nil {
			if err := m.Attempts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attempts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attempts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this jobs status based on the context it is used
func (m *JobsStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttempts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsStatus) contextValidateAttempts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attempts); i++ {

		if m.Attempts[i] != nil {

			if swag.IsZero(m.Attempts[i]) { // not required
				return nil
			}

			if err := m.Attempts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attempts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attempts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsStatus) UnmarshalBinary(b []byte) error {
	var res JobsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
