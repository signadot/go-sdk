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

// SandboxCustomizations sandbox customizations
//
// swagger:model sandbox.Customizations
type SandboxCustomizations struct {

	// Environment variable modifications that will be applied to the forked workload
	Env []*SandboxEnvVar `json:"env"`

	// One or more docker images that will be applied to the forked workload
	Images []*SandboxImage `json:"images"`

	// patch
	Patch *SandboxCustomPatch `json:"patch,omitempty"`
}

// Validate validates this sandbox customizations
func (m *SandboxCustomizations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SandboxCustomizations) validateEnv(formats strfmt.Registry) error {
	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {
		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {
			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SandboxCustomizations) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SandboxCustomizations) validatePatch(formats strfmt.Registry) error {
	if swag.IsZero(m.Patch) { // not required
		return nil
	}

	if m.Patch != nil {
		if err := m.Patch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patch")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sandbox customizations based on the context it is used
func (m *SandboxCustomizations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePatch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SandboxCustomizations) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if m.Env[i] != nil {

			if swag.IsZero(m.Env[i]) { // not required
				return nil
			}

			if err := m.Env[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SandboxCustomizations) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {

			if swag.IsZero(m.Images[i]) { // not required
				return nil
			}

			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SandboxCustomizations) contextValidatePatch(ctx context.Context, formats strfmt.Registry) error {

	if m.Patch != nil {

		if swag.IsZero(m.Patch) { // not required
			return nil
		}

		if err := m.Patch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patch")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SandboxCustomizations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxCustomizations) UnmarshalBinary(b []byte) error {
	var res SandboxCustomizations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
