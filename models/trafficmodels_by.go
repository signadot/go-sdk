// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrafficmodelsBy trafficmodels by
//
// swagger:model trafficmodels.By
type TrafficmodelsBy struct {

	// direction
	Direction string `json:"direction,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this trafficmodels by
func (m *TrafficmodelsBy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this trafficmodels by based on context it is used
func (m *TrafficmodelsBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrafficmodelsBy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrafficmodelsBy) UnmarshalBinary(b []byte) error {
	var res TrafficmodelsBy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}