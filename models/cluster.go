// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cluster cluster
//
// swagger:model Cluster
type Cluster struct {

	// The time when this cluster was registered with Signadot.
	CreatedAt string `json:"createdAt,omitempty"`

	// The name of the cluster.
	Name string `json:"name,omitempty"`

	// operator
	Operator *ClusterOperator `json:"operator,omitempty"`

	// sync status
	SyncStatus *ClustersSyncStatus `json:"syncStatus,omitempty"`
}

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	if m.Operator != nil {
		if err := m.Operator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operator")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) validateSyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncStatus) { // not required
		return nil
	}

	if m.SyncStatus != nil {
		if err := m.SyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("syncStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster based on the context it is used
func (m *Cluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) contextValidateOperator(ctx context.Context, formats strfmt.Registry) error {

	if m.Operator != nil {

		if swag.IsZero(m.Operator) { // not required
			return nil
		}

		if err := m.Operator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operator")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) contextValidateSyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.SyncStatus != nil {

		if swag.IsZero(m.SyncStatus) { // not required
			return nil
		}

		if err := m.SyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("syncStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
