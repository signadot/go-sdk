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

// RouteGroupMatch route group match
//
// swagger:model routeGroup.Match
type RouteGroupMatch struct {

	// When All is non-nil, T matches a set of labels if and only if every element of All matches them.
	// Only one field may be non-nil.
	All []*RouteGroupMatch `json:"all,omitempty"`

	// When Any is non-nil, T matches a set of labels if and only if some element of Any matches them.
	// Only one field may be non-nil.
	Any []*RouteGroupMatch `json:"any,omitempty"`

	// label
	Label *RouteGroupMatchLabel `json:"label,omitempty"`
}

// Validate validates this route group match
func (m *RouteGroupMatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAny(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouteGroupMatch) validateAll(formats strfmt.Registry) error {
	if swag.IsZero(m.All) { // not required
		return nil
	}

	for i := 0; i < len(m.All); i++ {
		if swag.IsZero(m.All[i]) { // not required
			continue
		}

		if m.All[i] != nil {
			if err := m.All[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("all" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("all" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RouteGroupMatch) validateAny(formats strfmt.Registry) error {
	if swag.IsZero(m.Any) { // not required
		return nil
	}

	for i := 0; i < len(m.Any); i++ {
		if swag.IsZero(m.Any[i]) { // not required
			continue
		}

		if m.Any[i] != nil {
			if err := m.Any[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("any" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("any" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RouteGroupMatch) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if m.Label != nil {
		if err := m.Label.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this route group match based on the context it is used
func (m *RouteGroupMatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAll(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAny(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouteGroupMatch) contextValidateAll(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.All); i++ {

		if m.All[i] != nil {

			if swag.IsZero(m.All[i]) { // not required
				return nil
			}

			if err := m.All[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("all" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("all" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RouteGroupMatch) contextValidateAny(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Any); i++ {

		if m.Any[i] != nil {

			if swag.IsZero(m.Any[i]) { // not required
				return nil
			}

			if err := m.Any[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("any" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("any" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RouteGroupMatch) contextValidateLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.Label != nil {

		if swag.IsZero(m.Label) { // not required
			return nil
		}

		if err := m.Label.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RouteGroupMatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteGroupMatch) UnmarshalBinary(b []byte) error {
	var res RouteGroupMatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
