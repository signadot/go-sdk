// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SandboxEndpoint sandbox endpoint
//
// swagger:model sandbox.Endpoint
type SandboxEndpoint struct {

	// baseline Url
	BaselineURL string `json:"baselineUrl,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// route type
	RouteType string `json:"routeType,omitempty"`

	// target
	Target string `json:"target,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this sandbox endpoint
func (m *SandboxEndpoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sandbox endpoint based on context it is used
func (m *SandboxEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SandboxEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxEndpoint) UnmarshalBinary(b []byte) error {
	var res SandboxEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
