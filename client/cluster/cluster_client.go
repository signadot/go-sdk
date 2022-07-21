// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddCluster(params *AddClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddClusterOK, error)

	CreateClusterToken(params *CreateClusterTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClusterTokenOK, error)

	DeleteClusterToken(params *DeleteClusterTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClusterTokenOK, error)

	GetCluster(params *GetClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterOK, error)

	GetClusterToken(params *GetClusterTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterTokenOK, error)

	ListClusterTokens(params *ListClusterTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClusterTokensOK, error)

	ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClustersOK, error)

	RemoveCluster(params *RemoveClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddCluster adds a cluster

  Add a Kubernetes cluster to Signadot.
*/
func (a *Client) AddCluster(params *AddClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "add-cluster",
		Method:             "PUT",
		PathPattern:        "/orgs/{orgName}/clusters/{clusterName}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddClusterReader{formats: a.formats},
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
	success, ok := result.(*AddClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for add-cluster: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateClusterToken creates cluster token

  Create a new token for connecting a cluster.
*/
func (a *Client) CreateClusterToken(params *CreateClusterTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClusterTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-cluster-token",
		Method:             "POST",
		PathPattern:        "/orgs/{orgName}/clusters/{clusterName}/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterTokenReader{formats: a.formats},
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
	success, ok := result.(*CreateClusterTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-cluster-token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteClusterToken deletes cluster token

  Delete a cluster token associated with a cluster.
*/
func (a *Client) DeleteClusterToken(params *DeleteClusterTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClusterTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-cluster-token",
		Method:             "DELETE",
		PathPattern:        "/orgs/{orgName}/clusters/{clusterName}/tokens/{tokenId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClusterTokenReader{formats: a.formats},
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
	success, ok := result.(*DeleteClusterTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-cluster-token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCluster gets a cluster

  Get a cluster.
*/
func (a *Client) GetCluster(params *GetClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-cluster",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/clusters/{clusterName}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
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
	success, ok := result.(*GetClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-cluster: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClusterToken gets a cluster token

  Get a cluster token associated with a cluster.
*/
func (a *Client) GetClusterToken(params *GetClusterTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-cluster-token",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/clusters/{clusterName}/tokens/{tokenId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterTokenReader{formats: a.formats},
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
	success, ok := result.(*GetClusterTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-cluster-token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListClusterTokens lists cluster tokens

  List the cluster tokens associated with a cluster.
*/
func (a *Client) ListClusterTokens(params *ListClusterTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClusterTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClusterTokensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-cluster-tokens",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/clusters/{clusterName}/tokens/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListClusterTokensReader{formats: a.formats},
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
	success, ok := result.(*ListClusterTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-cluster-tokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListClusters lists clusters

  List clusters.
*/
func (a *Client) ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-clusters",
		Method:             "GET",
		PathPattern:        "/orgs/{orgName}/clusters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
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
	success, ok := result.(*ListClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-clusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveCluster removes a cluster

  Remove a Kubernetes cluster from Signadot.
*/
func (a *Client) RemoveCluster(params *RemoveClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "remove-cluster",
		Method:             "DELETE",
		PathPattern:        "/orgs/{orgName}/clusters/{clusterName}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveClusterReader{formats: a.formats},
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
	success, ok := result.(*RemoveClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for remove-cluster: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
