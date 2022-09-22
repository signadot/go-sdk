// Code generated by go-swagger; DO NOT EDIT.

package sandboxes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sandboxes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sandboxes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ApplySandbox(params *ApplySandboxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ApplySandboxOK, error)

	DeleteSandbox(params *DeleteSandboxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSandboxOK, error)

	GetSandbox(params *GetSandboxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSandboxOK, error)

	ListSandboxes(params *ListSandboxesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSandboxesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ApplySandbox creates or update a sandbox

  Creates or updates a sandbox with the provided parameters.
*/
func (a *Client) ApplySandbox(params *ApplySandboxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ApplySandboxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplySandboxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "apply-sandbox",
		Method:             "PUT",
		PathPattern:        "/orgs/{orgName}/sandboxes/{sandboxName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplySandboxReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplySandboxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for apply-sandbox: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSandbox deletes a sandbox

  Delete a given sandbox.
*/
func (a *Client) DeleteSandbox(params *DeleteSandboxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSandboxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSandboxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-sandbox",
		Method:             "DELETE",
		PathPattern:        "/orgs/{orgName}/sandboxes/{sandboxName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSandboxReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSandboxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-sandbox: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSandbox gets a sandbox

  Fetch the details about a given sandbox.
*/
func (a *Client) GetSandbox(params *GetSandboxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSandboxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSandboxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-sandbox",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/sandboxes/{sandboxName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSandboxReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSandboxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-sandbox: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListSandboxes lists sandboxes

  List all sandboxes under the specified Signadot org.
*/
func (a *Client) ListSandboxes(params *ListSandboxesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSandboxesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSandboxesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-sandboxes",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/sandboxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSandboxesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSandboxesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-sandboxes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
