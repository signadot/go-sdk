// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SandboxForkOf sandbox fork of
//
// swagger:model sandbox.ForkOf
type SandboxForkOf struct {

	// Kind of entity that we want to route to. One of (Service or Deployment or Rollout).
	// Example: Deployment
	// Required: true
	Kind *string `json:"kind"`

	// Name of the entity within the Kubernetes cluster.
	// Example: my-frontend
	// Required: true
	Name *string `json:"name"`

	// Namespace within which the entity lives in the Kubernetes cluster.
	// Example: default
	// Required: true
	Namespace *string `json:"namespace"`
}

// Validate validates this sandbox fork of
func (m *SandboxForkOf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SandboxForkOf) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *SandboxForkOf) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SandboxForkOf) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sandbox fork of based on context it is used
func (m *SandboxForkOf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SandboxForkOf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SandboxForkOf) UnmarshalBinary(b []byte) error {
	var res SandboxForkOf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}