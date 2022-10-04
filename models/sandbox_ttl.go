// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SandboxTTL sandbox TTL
//
// swagger:model sandbox.TTL
type SandboxTTL struct {

	// Duration represents the duration until sandbox end of life.
	// It should be an unsigned integer not exceeding 32 bits followed by
	// a units character, which can be one of the following.
	//  - 'm' for minutes
	//  - 'h' for hours
	//  - 'd' for days
	//  - 'w' for weeks
	Duration string `json:"duration,omitempty"`

	// OffsetFrom indicates what the Duration is relative to.  It
	// may be the empty string or "createdAt".  The empty string
	// defaults to meaning "createdAt".
	OffsetFrom string `json:"offsetFrom,omitempty"`
}

// Validate validates this sandbox TTL
func (m *SandboxTTL) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sandbox TTL based on context it is used
func (m *SandboxTTL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SandboxTTL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxTTL) UnmarshalBinary(b []byte) error {
	var res SandboxTTL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}