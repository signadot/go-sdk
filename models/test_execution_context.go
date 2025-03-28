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

// TestExecutionContext test execution context
//
// swagger:model TestExecutionContext
type TestExecutionContext struct {

	// auto diff
	AutoDiff *TestExecutionAutoDiff `json:"autoDiff,omitempty"`

	// cluster
	Cluster string `json:"cluster,omitempty"`

	// Publish represents whether the test execution will show up by default
	// in the UI, i.e. to filter out ad-hoc and test the test usage.
	Publish bool `json:"publish,omitempty"`

	// routing
	Routing *JobRoutingContext `json:"routing,omitempty"`

	// Timeout represents an optional timeout for the test execution.
	// If not supplied, it defaults to the DefaultTimeout of the associated
	// test, if that is present.  If that is not present, it defaults to "5m".
	// It should be a string representing an unsigned integer not exceeding 32 bits followed by
	// a units character, which can be one of the following.
	//   - 'm' for minutes
	//   - 'h' for hours
	//   - 'd' for days
	//   - 'w' for weeks
	Timeout string `json:"timeout,omitempty"`
}

// Validate validates this test execution context
func (m *TestExecutionContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoDiff(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestExecutionContext) validateAutoDiff(formats strfmt.Registry) error {
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

func (m *TestExecutionContext) validateRouting(formats strfmt.Registry) error {
	if swag.IsZero(m.Routing) { // not required
		return nil
	}

	if m.Routing != nil {
		if err := m.Routing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("routing")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this test execution context based on the context it is used
func (m *TestExecutionContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoDiff(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestExecutionContext) contextValidateAutoDiff(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TestExecutionContext) contextValidateRouting(ctx context.Context, formats strfmt.Registry) error {

	if m.Routing != nil {

		if swag.IsZero(m.Routing) { // not required
			return nil
		}

		if err := m.Routing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("routing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestExecutionContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestExecutionContext) UnmarshalBinary(b []byte) error {
	var res TestExecutionContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
