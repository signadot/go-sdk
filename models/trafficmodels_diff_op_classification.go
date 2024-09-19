// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrafficmodelsDiffOpClassification trafficmodels diff op classification
//
// swagger:model trafficmodels.DiffOpClassification
type TrafficmodelsDiffOpClassification struct {

	// noise
	Noise float64 `json:"noise,omitempty"`
}

// Validate validates this trafficmodels diff op classification
func (m *TrafficmodelsDiffOpClassification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this trafficmodels diff op classification based on context it is used
func (m *TrafficmodelsDiffOpClassification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrafficmodelsDiffOpClassification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrafficmodelsDiffOpClassification) UnmarshalBinary(b []byte) error {
	var res TrafficmodelsDiffOpClassification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
