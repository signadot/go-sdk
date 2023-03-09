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

// NewListClustersParams creates a new ListClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListClustersParams() *ListClustersParams {
	return &ListClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListClustersParamsWithTimeout creates a new ListClustersParams object
// with the ability to set a timeout on a request.
func NewListClustersParamsWithTimeout(timeout time.Duration) *ListClustersParams {
	return &ListClustersParams{
		timeout: timeout,
	}
}

// NewListClustersParamsWithContext creates a new ListClustersParams object
// with the ability to set a context for a request.
func NewListClustersParamsWithContext(ctx context.Context) *ListClustersParams {
	return &ListClustersParams{
		Context: ctx,
	}
}

// NewListClustersParamsWithHTTPClient creates a new ListClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListClustersParamsWithHTTPClient(client *http.Client) *ListClustersParams {
	return &ListClustersParams{
		HTTPClient: client,
	}
}

/*
ListClustersParams contains all the parameters to send to the API endpoint

	for the list clusters operation.

	Typically these are written to a http.Request.
*/
type ListClustersParams struct {

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClustersParams) WithDefaults() *ListClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list clusters params
func (o *ListClustersParams) WithTimeout(timeout time.Duration) *ListClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list clusters params
func (o *ListClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list clusters params
func (o *ListClustersParams) WithContext(ctx context.Context) *ListClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list clusters params
func (o *ListClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list clusters params
func (o *ListClustersParams) WithHTTPClient(client *http.Client) *ListClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list clusters params
func (o *ListClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the list clusters params
func (o *ListClustersParams) WithOrgName(orgName string) *ListClustersParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the list clusters params
func (o *ListClustersParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WriteToRequest writes these params to a swagger request
func (o *ListClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
