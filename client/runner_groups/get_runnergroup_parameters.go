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

// NewGetRunnergroupParams creates a new GetRunnergroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunnergroupParams() *GetRunnergroupParams {
	return &GetRunnergroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunnergroupParamsWithTimeout creates a new GetRunnergroupParams object
// with the ability to set a timeout on a request.
func NewGetRunnergroupParamsWithTimeout(timeout time.Duration) *GetRunnergroupParams {
	return &GetRunnergroupParams{
		timeout: timeout,
	}
}

// NewGetRunnergroupParamsWithContext creates a new GetRunnergroupParams object
// with the ability to set a context for a request.
func NewGetRunnergroupParamsWithContext(ctx context.Context) *GetRunnergroupParams {
	return &GetRunnergroupParams{
		Context: ctx,
	}
}

// NewGetRunnergroupParamsWithHTTPClient creates a new GetRunnergroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunnergroupParamsWithHTTPClient(client *http.Client) *GetRunnergroupParams {
	return &GetRunnergroupParams{
		HTTPClient: client,
	}
}

/*
GetRunnergroupParams contains all the parameters to send to the API endpoint

	for the get runnergroup operation.

	Typically these are written to a http.Request.
*/
type GetRunnergroupParams struct {

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	/* RunnergroupName.

	   RunnerGroup Name
	*/
	RunnergroupName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get runnergroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunnergroupParams) WithDefaults() *GetRunnergroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get runnergroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunnergroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get runnergroup params
func (o *GetRunnergroupParams) WithTimeout(timeout time.Duration) *GetRunnergroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runnergroup params
func (o *GetRunnergroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runnergroup params
func (o *GetRunnergroupParams) WithContext(ctx context.Context) *GetRunnergroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runnergroup params
func (o *GetRunnergroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runnergroup params
func (o *GetRunnergroupParams) WithHTTPClient(client *http.Client) *GetRunnergroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runnergroup params
func (o *GetRunnergroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the get runnergroup params
func (o *GetRunnergroupParams) WithOrgName(orgName string) *GetRunnergroupParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the get runnergroup params
func (o *GetRunnergroupParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithRunnergroupName adds the runnergroupName to the get runnergroup params
func (o *GetRunnergroupParams) WithRunnergroupName(runnergroupName string) *GetRunnergroupParams {
	o.SetRunnergroupName(runnergroupName)
	return o
}

// SetRunnergroupName adds the runnergroupName to the get runnergroup params
func (o *GetRunnergroupParams) SetRunnergroupName(runnergroupName string) {
	o.RunnergroupName = runnergroupName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunnergroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	// path param runnergroupName
	if err := r.SetPathParam("runnergroupName", o.RunnergroupName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
