// Code generated by go-swagger; DO NOT EDIT.

package tests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new tests API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new tests API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new tests API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for tests API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ApplyTest(params *ApplyTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ApplyTestOK, error)

	DeleteTest(params *DeleteTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTestOK, error)

	GetTest(params *GetTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTestOK, error)

	ListTests(params *ListTestsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTestsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ApplyTest creates or update a test

Creates or updates a test with the provided parameters.
*/
func (a *Client) ApplyTest(params *ApplyTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ApplyTestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplyTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "apply-test",
		Method:             "PUT",
		PathPattern:        "/orgs/{orgName}/tests/{testName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplyTestReader{formats: a.formats},
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
	success, ok := result.(*ApplyTestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for apply-test: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTest deletes a test

Delete a given test.
*/
func (a *Client) DeleteTest(params *DeleteTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-test",
		Method:             "DELETE",
		PathPattern:        "/orgs/{orgName}/tests/{testName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTestReader{formats: a.formats},
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
	success, ok := result.(*DeleteTestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-test: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTest gets a test

Fetch the details about a given test.
*/
func (a *Client) GetTest(params *GetTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-test",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/tests/{testName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTestReader{formats: a.formats},
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
	success, ok := result.(*GetTestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-test: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListTests lists tests

List Tests
*/
func (a *Client) ListTests(params *ListTestsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-tests",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/tests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTestsReader{formats: a.formats},
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
	success, ok := result.(*ListTestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-tests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}