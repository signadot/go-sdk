// Code generated by go-swagger; DO NOT EDIT.

package test_executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new test executions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new test executions API client with basic auth credentials.
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

// New creates a new test executions API client with a bearer token for authentication.
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
Client for test executions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelTestExecution(params *CancelTestExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelTestExecutionOK, error)

	CreateTestExecution(params *CreateTestExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTestExecutionOK, error)

	GetTestExecution(params *GetTestExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTestExecutionOK, error)

	ListTestExecutions(params *ListTestExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTestExecutionsOK, error)

	QueryTestExecutions(params *QueryTestExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryTestExecutionsOK, error)

	TestExecutionTrafficDiff(params *TestExecutionTrafficDiffParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TestExecutionTrafficDiffOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CancelTestExecution cancels a test execution

Cancel a given test execution.
*/
func (a *Client) CancelTestExecution(params *CancelTestExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelTestExecutionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelTestExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancel-test-execution",
		Method:             "POST",
		PathPattern:        "/orgs/{orgName}/tests/{testName}/executions/{executionName}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelTestExecutionReader{formats: a.formats},
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
	success, ok := result.(*CancelTestExecutionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cancel-test-execution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateTestExecution creates a test execution

Creates a test with the provided parameters.
*/
func (a *Client) CreateTestExecution(params *CreateTestExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTestExecutionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTestExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-test-execution",
		Method:             "POST",
		PathPattern:        "/orgs/{orgName}/tests/{testName}/executions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTestExecutionReader{formats: a.formats},
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
	success, ok := result.(*CreateTestExecutionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-test-execution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTestExecution gets a test execution

Fetch the details about a given test execution.
*/
func (a *Client) GetTestExecution(params *GetTestExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTestExecutionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTestExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-test-execution",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/tests/{testName}/executions/{executionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTestExecutionReader{formats: a.formats},
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
	success, ok := result.(*GetTestExecutionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-test-execution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListTestExecutions lists test executions

List test executions for a given test
*/
func (a *Client) ListTestExecutions(params *ListTestExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTestExecutionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTestExecutionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-test-executions",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/tests/{testName}/executions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTestExecutionsReader{formats: a.formats},
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
	success, ok := result.(*ListTestExecutionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-test-executions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryTestExecutions queries test executions

Query test executions based on different criteria
*/
func (a *Client) QueryTestExecutions(params *QueryTestExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QueryTestExecutionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryTestExecutionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "query-test-executions",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/tests/executions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryTestExecutionsReader{formats: a.formats},
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
	success, ok := result.(*QueryTestExecutionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for query-test-executions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TestExecutionTrafficDiff gets the full traffic diff of a test execution

Get the full traffic diff of a test execution
*/
func (a *Client) TestExecutionTrafficDiff(params *TestExecutionTrafficDiffParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TestExecutionTrafficDiffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestExecutionTrafficDiffParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "test-execution-traffic-diff",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/tests/{testName}/executions/{executionName}/traffic-diff",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestExecutionTrafficDiffReader{formats: a.formats},
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
	success, ok := result.(*TestExecutionTrafficDiffOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for test-execution-traffic-diff: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
