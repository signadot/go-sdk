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

// V1PersistentVolumeClaimSpec v1 persistent volume claim spec
//
// swagger:model v1.PersistentVolumeClaimSpec
type V1PersistentVolumeClaimSpec struct {

	// accessModes contains the desired access modes the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	// +optional
	AccessModes []string `json:"accessModes"`

	// data source
	DataSource *V1TypedLocalObjectReference `json:"dataSource,omitempty"`

	// data source ref
	DataSourceRef *V1TypedLocalObjectReference `json:"dataSourceRef,omitempty"`

	// resources
	Resources *V1ResourceRequirements `json:"resources,omitempty"`

	// selector
	Selector *V1LabelSelector `json:"selector,omitempty"`

	// storageClassName is the name of the StorageClass required by the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	// +optional
	StorageClassName string `json:"storageClassName,omitempty"`

	// volumeMode defines what type of volume is required by the claim.
	// Value of Filesystem is implied when not included in claim spec.
	// +optional
	VolumeMode string `json:"volumeMode,omitempty"`

	// volumeName is the binding reference to the PersistentVolume backing this claim.
	// +optional
	VolumeName string `json:"volumeName,omitempty"`
}

// Validate validates this v1 persistent volume claim spec
func (m *V1PersistentVolumeClaimSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSourceRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PersistentVolumeClaimSpec) validateDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSource) { // not required
		return nil
	}

	if m.DataSource != nil {
		if err := m.DataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSource")
			}
			return err
		}
	}

	return nil
}

func (m *V1PersistentVolumeClaimSpec) validateDataSourceRef(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSourceRef) { // not required
		return nil
	}

	if m.DataSourceRef != nil {
		if err := m.DataSourceRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSourceRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSourceRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1PersistentVolumeClaimSpec) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *V1PersistentVolumeClaimSpec) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 persistent volume claim spec based on the context it is used
func (m *V1PersistentVolumeClaimSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PersistentVolumeClaimSpec) contextValidateDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.DataSource != nil {
		if err := m.DataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSource")
			}
			return err
		}
	}

	return nil
}

func (m *V1PersistentVolumeClaimSpec) contextValidateDataSourceRef(ctx context.Context, formats strfmt.Registry) error {

	if m.DataSourceRef != nil {
		if err := m.DataSourceRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSourceRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSourceRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1PersistentVolumeClaimSpec) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {
		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *V1PersistentVolumeClaimSpec) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.Selector != nil {
		if err := m.Selector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PersistentVolumeClaimSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PersistentVolumeClaimSpec) UnmarshalBinary(b []byte) error {
	var res V1PersistentVolumeClaimSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
