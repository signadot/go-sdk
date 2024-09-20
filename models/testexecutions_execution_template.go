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

// TestexecutionsExecutionTemplate testexecutions execution template
//
// swagger:model testexecutions.ExecutionTemplate
type TestexecutionsExecutionTemplate struct {

	// auto diff
	AutoDiff *TestexecutionsAutoDiff `json:"autoDiff,omitempty"`

	// cluster
	Cluster string `json:"cluster,omitempty"`
}

// Validate validates this testexecutions execution template
func (m *TestexecutionsExecutionTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoDiff(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestexecutionsExecutionTemplate) validateAutoDiff(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoDiff) { // not required
		return nil
	}

	if m.AutoDiff != nil {
		if err := m.AutoDiff.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDiff")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoDiff")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this testexecutions execution template based on the context it is used
func (m *TestexecutionsExecutionTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoDiff(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestexecutionsExecutionTemplate) contextValidateAutoDiff(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoDiff != nil {

		if swag.IsZero(m.AutoDiff) { // not required
			return nil
		}

		if err := m.AutoDiff.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDiff")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoDiff")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestexecutionsExecutionTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestexecutionsExecutionTemplate) UnmarshalBinary(b []byte) error {
	var res TestexecutionsExecutionTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
