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

// CaptureRequest capture request
//
// swagger:model capture.Request
type CaptureRequest struct {

	// host
	Host string `json:"host,omitempty"`

	// message
	Message *CaptureMessage `json:"message,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// proto
	Proto string `json:"proto,omitempty"`

	// query
	Query CaptureValues `json:"query,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`
}

// Validate validates this capture request
func (m *CaptureRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CaptureRequest) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {
		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *CaptureRequest) validateQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this capture request based on the context it is used
func (m *CaptureRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CaptureRequest) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.Message != nil {

		if swag.IsZero(m.Message) { // not required
			return nil
		}

		if err := m.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *CaptureRequest) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if err := m.Query.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("query")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("query")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CaptureRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CaptureRequest) UnmarshalBinary(b []byte) error {
	var res CaptureRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}