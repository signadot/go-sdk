// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrgsOrg orgs org
//
// swagger:model orgs.Org
type OrgsOrg struct {

	// billing tier
	BillingTier int64 `json:"billingTier,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this orgs org
func (m *OrgsOrg) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this orgs org based on context it is used
func (m *OrgsOrg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrgsOrg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrgsOrg) UnmarshalBinary(b []byte) error {
	var res OrgsOrg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
