// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Branch branch
//
// swagger:model Branch
type Branch struct {

	// head commit
	HeadCommit string `json:"headCommit,omitempty"`

	// pull request
	PullRequest int64 `json:"pullRequest,omitempty"`

	// vcs name
	VcsName string `json:"vcsName,omitempty"`

	// vcs repo
	VcsRepo string `json:"vcsRepo,omitempty"`

	// vcs type
	VcsType string `json:"vcsType,omitempty"`
}

// Validate validates this branch
func (m *Branch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this branch based on context it is used
func (m *Branch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Branch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Branch) UnmarshalBinary(b []byte) error {
	var res Branch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
