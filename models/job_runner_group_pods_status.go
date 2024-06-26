// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobRunnerGroupPodsStatus job runner group pods status
//
// swagger:model jobRunnerGroup.PodsStatus
type JobRunnerGroupPodsStatus struct {

	// idle
	Idle int64 `json:"idle,omitempty"`

	// not ready
	NotReady int64 `json:"notReady,omitempty"`

	// ready
	Ready int64 `json:"ready,omitempty"`
}

// Validate validates this job runner group pods status
func (m *JobRunnerGroupPodsStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this job runner group pods status based on context it is used
func (m *JobRunnerGroupPodsStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobRunnerGroupPodsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobRunnerGroupPodsStatus) UnmarshalBinary(b []byte) error {
	var res JobRunnerGroupPodsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
