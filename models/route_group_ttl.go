// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RouteGroupTTL route group TTL
//
// swagger:model routeGroup.TTL
type RouteGroupTTL struct {

	// Duration represents the duration until routegroup's end of life.
	// It should be an unsigned integer not exceeding 32 bits followed by
	// a units character, which can be one of the following.
	//   - 'm' for minutes
	//   - 'h' for hours
	//   - 'd' for days
	//   - 'w' for weeks
	Duration string `json:"duration,omitempty"`

	// OffsetFrom indicates what the Duration is relative to.  It
	// may be the empty string, "noMatchedSandboxes", "createdAt" or "updatedAt". The empty string
	// defaults to meaning "noMatchedSandboxes".
	OffsetFrom string `json:"offsetFrom,omitempty"`
}

// Validate validates this route group TTL
func (m *RouteGroupTTL) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this route group TTL based on context it is used
func (m *RouteGroupTTL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouteGroupTTL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteGroupTTL) UnmarshalBinary(b []byte) error {
	var res RouteGroupTTL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
