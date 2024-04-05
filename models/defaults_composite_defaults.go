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

// DefaultsCompositeDefaults defaults composite defaults
//
// swagger:model defaults.CompositeDefaults
type DefaultsCompositeDefaults struct {

	// cluster
	Cluster []*DefaultsDefault `json:"cluster"`
}

// Validate validates this defaults composite defaults
func (m *DefaultsCompositeDefaults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DefaultsCompositeDefaults) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	for i := 0; i < len(m.Cluster); i++ {
		if swag.IsZero(m.Cluster[i]) { // not required
			continue
		}

		if m.Cluster[i] != nil {
			if err := m.Cluster[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this defaults composite defaults based on the context it is used
func (m *DefaultsCompositeDefaults) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DefaultsCompositeDefaults) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cluster); i++ {

		if m.Cluster[i] != nil {

			if swag.IsZero(m.Cluster[i]) { // not required
				return nil
			}

			if err := m.Cluster[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DefaultsCompositeDefaults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefaultsCompositeDefaults) UnmarshalBinary(b []byte) error {
	var res DefaultsCompositeDefaults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
