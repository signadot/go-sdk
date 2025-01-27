// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SandboxTestExecutionsSummaryChecks sandbox test executions summary checks
//
// swagger:model sandbox.TestExecutionsSummaryChecks
type SandboxTestExecutionsSummaryChecks struct {

	// failed
	Failed int64 `json:"failed,omitempty"`

	// passed
	Passed int64 `json:"passed,omitempty"`
}

// Validate validates this sandbox test executions summary checks
func (m *SandboxTestExecutionsSummaryChecks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sandbox test executions summary checks based on context it is used
func (m *SandboxTestExecutionsSummaryChecks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SandboxTestExecutionsSummaryChecks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxTestExecutionsSummaryChecks) UnmarshalBinary(b []byte) error {
	var res SandboxTestExecutionsSummaryChecks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
