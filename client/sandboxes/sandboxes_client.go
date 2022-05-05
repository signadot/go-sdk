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
	CreateNewSandbox(params *CreateNewSandboxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNewSandboxOK, error)

	DeleteSandboxByID(params *DeleteSandboxByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSandboxByIDOK, error)

	GetSandboxByID(params *GetSandboxByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSandboxByIDOK, error)

	GetSandboxReady(params *GetSandboxReadyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSandboxReadyOK, error)

	GetSandboxes(params *GetSandboxesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSandboxesOK, error)

	UpsertPrWorkspace(params *UpsertPrWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpsertPrWorkspaceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNewSandbox creates a new sandbox

  Creates a new sandbox with the provided parameters
*/
func (a *Client) CreateNewSandbox(params *CreateNewSandboxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNewSandboxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNewSandboxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-new-sandbox",
		Method:             "POST",
		PathPattern:        "/orgs/{orgName}/sandboxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNewSandboxReader{formats: a.formats},
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
	success, ok := result.(*CreateNewSandboxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-new-sandbox: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSandboxByID deletes a sandbox by ID

  Delete the sandbox when its ID is specified
*/
func (a *Client) DeleteSandboxByID(params *DeleteSandboxByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSandboxByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSandboxByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-sandbox-by-id",
		Method:             "DELETE",
		PathPattern:        "/orgs/{orgName}/sandboxes/{sandboxID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSandboxByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteSandboxByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-sandbox-by-id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSandboxByID gets a sandbox by ID

  Fetch the details about a sandbox when its ID is specified
*/
func (a *Client) GetSandboxByID(params *GetSandboxByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSandboxByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSandboxByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-sandbox-by-id",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/sandboxes/{sandboxID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSandboxByIDReader{formats: a.formats},
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
	success, ok := result.(*GetSandboxByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-sandbox-by-id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSandboxReady checks sandbox readiness

  Checks readiness of a sandbox with rate limiting enforced by polling
*/
func (a *Client) GetSandboxReady(params *GetSandboxReadyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSandboxReadyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSandboxReadyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-sandbox-ready",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/sandboxes/{sandboxID}/ready",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSandboxReadyReader{formats: a.formats},
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
	success, ok := result.(*GetSandboxReadyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-sandbox-ready: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSandboxes lists sandboxes

  List all sandboxes under the specified Signadot org.
*/
func (a *Client) GetSandboxes(params *GetSandboxesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSandboxesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSandboxesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-sandboxes",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/sandboxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSandboxesReader{formats: a.formats},
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
	success, ok := result.(*GetSandboxesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-sandboxes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpsertPrWorkspace creates or update workspace from pull request

  Create a workspace that is associated with a specified pull request.
If no workspaces already exist for the given pull request, this creates a new workspace. Otherwise, the new changes are applied to the existing workspaces.
This endpoint uses the old terminology "workspace" instead of "sandbox" for backward compatability.
### Example

Here's an example to create/update a workspace from a pull request in the [HotROD](https://github.com/signadot/hotrod) application. To simplify the use case, consider that only the [Route Service](https://github.com/signadot/hotrod/tree/main/services/route) had changes, and that the changes were published to an image file named `signadot/hotrod-route` and tag `4e75b0b822ecbbbb4c917b0fffeb337589d82456`. This could be tagged as anything e.g. `latest`, `e2e-test` etc.

As a part of workspace creation, Signadot creates forked Kubernetes workloads (such as deployments) from existing workloads using the provided images. In order to access the forked workloads(s), `endpoints` are specified and corresponding to each endpoint, a preview URL is generated. These preview URLs are returned as part of the response. In the example below, we're dealing with a deployment named `route` in the `hotrod` namespace that is running the docker image `signadot/hotrod-route`.

In the below request to create a workspace, the deployment named `route` is forked and the fork will be created with the new docker image tag as specified by `newTag`. The new deployment that is created will be running the a new docker image tag as specified by `newTag`. There will be a single preview URL associated with this workspace as specified in `endpoints` and it will point to the fork of the `route`
deployment that was created.

```json
{
  "cluster": "signadot",
  "namespace": "hotrod",
  "headCommit": "5e35abfa94626c4853eca51ecd435a779ded4123",
  "images": [
    {
      "name": "signadot/hotrod-route",
      "newTag": "4e75b0b822ecbbbb4c917b0fffeb337589d82456"
    }
  ],
  "endpoints": [
    {
      "routeType": "fork",
      "protocol": "http",
      "forkOf": {
        "kind": "Deployment",
        "name": "route",
        "namespace": "hotrod"
      }
    }
  ]
}
```

In response the API call returns us a confirmation of the workspace created / updated along with the information on preview endpoints.

```json
{
  "workspaceIDs": [
    "xtc54uh8p2rhs"
  ],
  "previewEndpoints": [
    {
      "id": "fbm4983sxb111",
      "routeType": "fork",
      "name": "route",
      "protocol": "http",
      "clusterID": "55rjfjf3rn222",
      "cluster": "signadot",
      "forkOf": {
        "kind": "Deployment",
        "namespace": "hotrod",
        "name": "route"
      },
      "previewURL": "https://route--hotrod-131.preview.signadot.com"
    }
  ]
}

```

You can parse the `previewEndpoints` section in the response to find the previewURL associated with the endpoint created as part of the workspace.

*/
func (a *Client) UpsertPrWorkspace(params *UpsertPrWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpsertPrWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertPrWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "upsert-pr-workspace",
		Method:             "POST",
		PathPattern:        "/repos/{githubOrg}/{githubRepo}/pulls/{prNumber}/workspaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpsertPrWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*UpsertPrWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upsert-pr-workspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
