// Code generated by go-swagger; DO NOT EDIT.

package runner_groups

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

// NewListRunnergroupParams creates a new ListRunnergroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRunnergroupParams() *ListRunnergroupParams {
	return &ListRunnergroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRunnergroupParamsWithTimeout creates a new ListRunnergroupParams object
// with the ability to set a timeout on a request.
func NewListRunnergroupParamsWithTimeout(timeout time.Duration) *ListRunnergroupParams {
	return &ListRunnergroupParams{
		timeout: timeout,
	}
}

// NewListRunnergroupParamsWithContext creates a new ListRunnergroupParams object
// with the ability to set a context for a request.
func NewListRunnergroupParamsWithContext(ctx context.Context) *ListRunnergroupParams {
	return &ListRunnergroupParams{
		Context: ctx,
	}
}

// NewListRunnergroupParamsWithHTTPClient creates a new ListRunnergroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRunnergroupParamsWithHTTPClient(client *http.Client) *ListRunnergroupParams {
	return &ListRunnergroupParams{
		HTTPClient: client,
	}
}

/*
ListRunnergroupParams contains all the parameters to send to the API endpoint

	for the list runnergroup operation.

	Typically these are written to a http.Request.
*/
type ListRunnergroupParams struct {

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list runnergroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRunnergroupParams) WithDefaults() *ListRunnergroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list runnergroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRunnergroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list runnergroup params
func (o *ListRunnergroupParams) WithTimeout(timeout time.Duration) *ListRunnergroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list runnergroup params
func (o *ListRunnergroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list runnergroup params
func (o *ListRunnergroupParams) WithContext(ctx context.Context) *ListRunnergroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list runnergroup params
func (o *ListRunnergroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list runnergroup params
func (o *ListRunnergroupParams) WithHTTPClient(client *http.Client) *ListRunnergroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list runnergroup params
func (o *ListRunnergroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the list runnergroup params
func (o *ListRunnergroupParams) WithOrgName(orgName string) *ListRunnergroupParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the list runnergroup params
func (o *ListRunnergroupParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WriteToRequest writes these params to a swagger request
func (o *ListRunnergroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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