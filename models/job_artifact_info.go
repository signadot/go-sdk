// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobArtifactInfo job artifact info
//
// swagger:model JobArtifactInfo
type JobArtifactInfo struct {

	// checksum s h a256
	ChecksumSHA256 string `json:"checksumSHA256,omitempty"`

	// content encoding
	ContentEncoding string `json:"contentEncoding,omitempty"`

	// content length
	ContentLength int64 `json:"contentLength,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Validate validates this job artifact info
func (m *JobArtifactInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this job artifact info based on context it is used
func (m *JobArtifactInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobArtifactInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobArtifactInfo) UnmarshalBinary(b []byte) error {
	var res JobArtifactInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
