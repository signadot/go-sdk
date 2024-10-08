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

// RequestDiff request diff
//
// swagger:model RequestDiff
type RequestDiff struct {

	// baseline
	Baseline *CaptureRequest `json:"baseline,omitempty"`

	// diff ops
	DiffOps []*DiffOp `json:"diffOps"`

	// target
	Target *CaptureRequest `json:"target,omitempty"`
}

// Validate validates this request diff
func (m *RequestDiff) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiffOps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestDiff) validateBaseline(formats strfmt.Registry) error {
	if swag.IsZero(m.Baseline) { // not required
		return nil
	}

	if m.Baseline != nil {
		if err := m.Baseline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseline")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseline")
			}
			return err
		}
	}

	return nil
}

func (m *RequestDiff) validateDiffOps(formats strfmt.Registry) error {
	if swag.IsZero(m.DiffOps) { // not required
		return nil
	}

	for i := 0; i < len(m.DiffOps); i++ {
		if swag.IsZero(m.DiffOps[i]) { // not required
			continue
		}

		if m.DiffOps[i] != nil {
			if err := m.DiffOps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("diffOps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("diffOps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RequestDiff) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this request diff based on the context it is used
func (m *RequestDiff) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiffOps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestDiff) contextValidateBaseline(ctx context.Context, formats strfmt.Registry) error {

	if m.Baseline != nil {

		if swag.IsZero(m.Baseline) { // not required
			return nil
		}

		if err := m.Baseline.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseline")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseline")
			}
			return err
		}
	}

	return nil
}

func (m *RequestDiff) contextValidateDiffOps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DiffOps); i++ {

		if m.DiffOps[i] != nil {

			if swag.IsZero(m.DiffOps[i]) { // not required
				return nil
			}

			if err := m.DiffOps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("diffOps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("diffOps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RequestDiff) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestDiff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestDiff) UnmarshalBinary(b []byte) error {
	var res RequestDiff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
