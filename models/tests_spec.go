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

// TestsSpec tests spec
//
// swagger:model tests.Spec
type TestsSpec struct {

	// name
	Name string `json:"name,omitempty"`

	// script
	Script string `json:"script,omitempty"`

	// triggers
	Triggers []*TestsTrigger `json:"triggers"`
}

// Validate validates this tests spec
func (m *TestsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTriggers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestsSpec) validateTriggers(formats strfmt.Registry) error {
	if swag.IsZero(m.Triggers) { // not required
		return nil
	}

	for i := 0; i < len(m.Triggers); i++ {
		if swag.IsZero(m.Triggers[i]) { // not required
			continue
		}

		if m.Triggers[i] != nil {
			if err := m.Triggers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("triggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("triggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tests spec based on the context it is used
func (m *TestsSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTriggers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestsSpec) contextValidateTriggers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Triggers); i++ {

		if m.Triggers[i] != nil {

			if swag.IsZero(m.Triggers[i]) { // not required
				return nil
			}

			if err := m.Triggers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("triggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("triggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestsSpec) UnmarshalBinary(b []byte) error {
	var res TestsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}