// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Image image
//
// swagger:model Image
type Image struct {

	// Container is the container this image must apply to in the forked workload
	Container string `json:"container,omitempty"`

	// Name of the image (e.g. gcr.io/my-image/repo:my-tag)
	Image string `json:"image,omitempty"`
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image based on context it is used
func (m *Image) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
