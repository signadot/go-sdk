// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunnergroupsAutoScaling runnergroups auto scaling
//
// swagger:model runnergroups.AutoScaling
type RunnergroupsAutoScaling struct {

	// max pods
	MaxPods int64 `json:"maxPods,omitempty"`

	// min pods
	MinPods int64 `json:"minPods,omitempty"`
}

// Validate validates this runnergroups auto scaling
func (m *RunnergroupsAutoScaling) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this runnergroups auto scaling based on context it is used
func (m *RunnergroupsAutoScaling) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunnergroupsAutoScaling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunnergroupsAutoScaling) UnmarshalBinary(b []byte) error {
	var res RunnergroupsAutoScaling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
