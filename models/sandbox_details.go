// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SandboxDetails sandbox details
//
// swagger:model SandboxDetails
type SandboxDetails struct {

	// branches
	Branches []*Branch `json:"branches"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// default service
	DefaultService string `json:"defaultService,omitempty"`

	// default service port
	DefaultServicePort int64 `json:"defaultServicePort,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// preview endpoints
	PreviewEndpoints []*PreviewEndpoint `json:"previewEndpoints"`

	// preview URL
	PreviewURL string `json:"previewURL,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this sandbox details
func (m *SandboxDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviewEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SandboxDetails) validateBranches(formats strfmt.Registry) error {
	if swag.IsZero(m.Branches) { // not required
		return nil
	}

	for i := 0; i < len(m.Branches); i++ {
		if swag.IsZero(m.Branches[i]) { // not required
			continue
		}

		if m.Branches[i] != nil {
			if err := m.Branches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("branches" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("branches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SandboxDetails) validatePreviewEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.PreviewEndpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.PreviewEndpoints); i++ {
		if swag.IsZero(m.PreviewEndpoints[i]) { // not required
			continue
		}

		if m.PreviewEndpoints[i] != nil {
			if err := m.PreviewEndpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("previewEndpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("previewEndpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sandbox details based on the context it is used
func (m *SandboxDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBranches(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreviewEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SandboxDetails) contextValidateBranches(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Branches); i++ {

		if m.Branches[i] != nil {
			if err := m.Branches[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("branches" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("branches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SandboxDetails) contextValidatePreviewEndpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PreviewEndpoints); i++ {

		if m.PreviewEndpoints[i] != nil {
			if err := m.PreviewEndpoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("previewEndpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("previewEndpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SandboxDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxDetails) UnmarshalBinary(b []byte) error {
	var res SandboxDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
