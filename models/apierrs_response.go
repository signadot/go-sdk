// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApierrsResponse apierrs response
//
// swagger:model apierrs.Response
type ApierrsResponse struct {

	// error
	Error string `json:"error,omitempty"`

	// request ID
	RequestID string `json:"requestID,omitempty"`
}

// Validate validates this apierrs response
func (m *ApierrsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this apierrs response based on context it is used
func (m *ApierrsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApierrsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApierrsResponse) UnmarshalBinary(b []byte) error {
	var res ApierrsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
