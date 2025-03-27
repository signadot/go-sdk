// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssistantsMessage assistants message
//
// swagger:model assistants.Message
type AssistantsMessage struct {

	// content
	Content string `json:"content,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// role
	Role string `json:"role,omitempty"`
}

// Validate validates this assistants message
func (m *AssistantsMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assistants message based on context it is used
func (m *AssistantsMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssistantsMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssistantsMessage) UnmarshalBinary(b []byte) error {
	var res AssistantsMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
