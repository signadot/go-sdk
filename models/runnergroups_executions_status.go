// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunnergroupsExecutionsStatus runnergroups executions status
//
// swagger:model runnergroups.ExecutionsStatus
type RunnergroupsExecutionsStatus struct {

	// completed
	Completed int64 `json:"completed,omitempty"`

	// queued
	Queued int64 `json:"queued,omitempty"`

	// running
	Running int64 `json:"running,omitempty"`
}

// Validate validates this runnergroups executions status
func (m *RunnergroupsExecutionsStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this runnergroups executions status based on context it is used
func (m *RunnergroupsExecutionsStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunnergroupsExecutionsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunnergroupsExecutionsStatus) UnmarshalBinary(b []byte) error {
	var res RunnergroupsExecutionsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}