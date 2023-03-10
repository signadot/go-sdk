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

// ResourcepluginSpec resourceplugin spec
//
// swagger:model resourceplugin.Spec
type ResourcepluginSpec struct {

	// Create refers to the `create` steps for spinning up the resource
	Create []*ResourcepluginStep `json:"create"`

	// Delete refers to the `delete` steps for spinning up the resource
	Delete []*ResourcepluginStep `json:"delete"`

	// Description for the resource
	Description string `json:"description,omitempty"`

	// runner
	Runner *ResourcepluginRunner `json:"runner,omitempty"`
}

// Validate validates this resourceplugin spec
func (m *ResourcepluginSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcepluginSpec) validateCreate(formats strfmt.Registry) error {
	if swag.IsZero(m.Create) { // not required
		return nil
	}

	for i := 0; i < len(m.Create); i++ {
		if swag.IsZero(m.Create[i]) { // not required
			continue
		}

		if m.Create[i] != nil {
			if err := m.Create[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("create" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("create" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourcepluginSpec) validateDelete(formats strfmt.Registry) error {
	if swag.IsZero(m.Delete) { // not required
		return nil
	}

	for i := 0; i < len(m.Delete); i++ {
		if swag.IsZero(m.Delete[i]) { // not required
			continue
		}

		if m.Delete[i] != nil {
			if err := m.Delete[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delete" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delete" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourcepluginSpec) validateRunner(formats strfmt.Registry) error {
	if swag.IsZero(m.Runner) { // not required
		return nil
	}

	if m.Runner != nil {
		if err := m.Runner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this resourceplugin spec based on the context it is used
func (m *ResourcepluginSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDelete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcepluginSpec) contextValidateCreate(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Create); i++ {

		if m.Create[i] != nil {
			if err := m.Create[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("create" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("create" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourcepluginSpec) contextValidateDelete(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Delete); i++ {

		if m.Delete[i] != nil {
			if err := m.Delete[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delete" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delete" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourcepluginSpec) contextValidateRunner(ctx context.Context, formats strfmt.Registry) error {

	if m.Runner != nil {
		if err := m.Runner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcepluginSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcepluginSpec) UnmarshalBinary(b []byte) error {
	var res ResourcepluginSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
