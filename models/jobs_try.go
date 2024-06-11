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

// JobsTry jobs try
//
// swagger:model jobs.Try
type JobsTry struct {

	// failed
	Failed *JobsFailedState `json:"failed,omitempty"`
}

// Validate validates this jobs try
func (m *JobsTry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsTry) validateFailed(formats strfmt.Registry) error {
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

// ContextValidate validate this jobs try based on the context it is used
func (m *JobsTry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFailed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsTry) contextValidateFailed(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *JobsTry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsTry) UnmarshalBinary(b []byte) error {
	var res JobsTry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
