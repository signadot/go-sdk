// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClustersReadinessStats clusters readiness stats
//
// swagger:model clusters.ReadinessStats
type ClustersReadinessStats struct {

	// count
	Count int64 `json:"count,omitempty"`

	// ready condition
	ReadyCondition string `json:"readyCondition,omitempty"`

	// ready reason
	ReadyReason string `json:"readyReason,omitempty"`
}

// Validate validates this clusters readiness stats
func (m *ClustersReadinessStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clusters readiness stats based on context it is used
func (m *ClustersReadinessStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClustersReadinessStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersReadinessStats) UnmarshalBinary(b []byte) error {
	var res ClustersReadinessStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
