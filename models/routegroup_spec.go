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

// RoutegroupSpec routegroup spec
//
// swagger:model routegroup.Spec
type RoutegroupSpec struct {

	// Cluster gives the cluster to which the route group applies.
	Cluster string `json:"cluster,omitempty"`

	// Endpoints define endpoints which target different in-cluster
	// services.
	Endpoints []*SandboxRouteGroupEndpoint `json:"endpoints"`

	// match
	Match *RoutegroupMatch `json:"match,omitempty"`
}

// Validate validates this routegroup spec
func (m *RoutegroupSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutegroupSpec) validateEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Endpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.Endpoints); i++ {
		if swag.IsZero(m.Endpoints[i]) { // not required
			continue
		}

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RoutegroupSpec) validateMatch(formats strfmt.Registry) error {
	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if m.Match != nil {
		if err := m.Match.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("match")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("match")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this routegroup spec based on the context it is used
func (m *RoutegroupSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMatch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutegroupSpec) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Endpoints); i++ {

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RoutegroupSpec) contextValidateMatch(ctx context.Context, formats strfmt.Registry) error {

	if m.Match != nil {
		if err := m.Match.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("match")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("match")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoutegroupSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutegroupSpec) UnmarshalBinary(b []byte) error {
	var res RoutegroupSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
