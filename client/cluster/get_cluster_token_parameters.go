// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetClusterTokenParams creates a new GetClusterTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterTokenParams() *GetClusterTokenParams {
	return &GetClusterTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTokenParamsWithTimeout creates a new GetClusterTokenParams object
// with the ability to set a timeout on a request.
func NewGetClusterTokenParamsWithTimeout(timeout time.Duration) *GetClusterTokenParams {
	return &GetClusterTokenParams{
		timeout: timeout,
	}
}

// NewGetClusterTokenParamsWithContext creates a new GetClusterTokenParams object
// with the ability to set a context for a request.
func NewGetClusterTokenParamsWithContext(ctx context.Context) *GetClusterTokenParams {
	return &GetClusterTokenParams{
		Context: ctx,
	}
}

// NewGetClusterTokenParamsWithHTTPClient creates a new GetClusterTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterTokenParamsWithHTTPClient(client *http.Client) *GetClusterTokenParams {
	return &GetClusterTokenParams{
		HTTPClient: client,
	}
}

/* GetClusterTokenParams contains all the parameters to send to the API endpoint
   for the get cluster token operation.

   Typically these are written to a http.Request.
*/
type GetClusterTokenParams struct {

	/* ClusterName.

	   Cluster Name
	*/
	ClusterName string

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	/* TokenID.

	   Token Id
	*/
	TokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTokenParams) WithDefaults() *GetClusterTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster token params
func (o *GetClusterTokenParams) WithTimeout(timeout time.Duration) *GetClusterTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster token params
func (o *GetClusterTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster token params
func (o *GetClusterTokenParams) WithContext(ctx context.Context) *GetClusterTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster token params
func (o *GetClusterTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster token params
func (o *GetClusterTokenParams) WithHTTPClient(client *http.Client) *GetClusterTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster token params
func (o *GetClusterTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterName adds the clusterName to the get cluster token params
func (o *GetClusterTokenParams) WithClusterName(clusterName string) *GetClusterTokenParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get cluster token params
func (o *GetClusterTokenParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithOrgName adds the orgName to the get cluster token params
func (o *GetClusterTokenParams) WithOrgName(orgName string) *GetClusterTokenParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the get cluster token params
func (o *GetClusterTokenParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithTokenID adds the tokenID to the get cluster token params
func (o *GetClusterTokenParams) WithTokenID(tokenID string) *GetClusterTokenParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the get cluster token params
func (o *GetClusterTokenParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	// path param tokenId
	if err := r.SetPathParam("tokenId", o.TokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
