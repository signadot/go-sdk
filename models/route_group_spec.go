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

// RouteGroupSpec route group spec
//
// swagger:model routeGroup.Spec
type RouteGroupSpec struct {

	// Cluster gives the cluster to which the route group applies.
	Cluster string `json:"cluster,omitempty"`

	// Description provides a short description of the route group.
	Description string `json:"description,omitempty"`

	// Endpoints define endpoints which target different in-cluster
	// services.
	Endpoints []*RouteGroupSpecEndpoint `json:"endpoints"`

	// match
	Match *RouteGroupMatch `json:"match,omitempty"`
}

// Validate validates this route group spec
func (m *RouteGroupSpec) Validate(formats strfmt.Registry) error {
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

func (m *RouteGroupSpec) validateEndpoints(formats strfmt.Registry) error {
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

func (m *RouteGroupSpec) validateMatch(formats strfmt.Registry) error {
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

// ContextValidate validate this route group spec based on the context it is used
func (m *RouteGroupSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *RouteGroupSpec) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RouteGroupSpec) contextValidateMatch(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RouteGroupSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteGroupSpec) UnmarshalBinary(b []byte) error {
	var res RouteGroupSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
