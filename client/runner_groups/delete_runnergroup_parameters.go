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

// NewDeleteRunnergroupParams creates a new DeleteRunnergroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRunnergroupParams() *DeleteRunnergroupParams {
	return &DeleteRunnergroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunnergroupParamsWithTimeout creates a new DeleteRunnergroupParams object
// with the ability to set a timeout on a request.
func NewDeleteRunnergroupParamsWithTimeout(timeout time.Duration) *DeleteRunnergroupParams {
	return &DeleteRunnergroupParams{
		timeout: timeout,
	}
}

// NewDeleteRunnergroupParamsWithContext creates a new DeleteRunnergroupParams object
// with the ability to set a context for a request.
func NewDeleteRunnergroupParamsWithContext(ctx context.Context) *DeleteRunnergroupParams {
	return &DeleteRunnergroupParams{
		Context: ctx,
	}
}

// NewDeleteRunnergroupParamsWithHTTPClient creates a new DeleteRunnergroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRunnergroupParamsWithHTTPClient(client *http.Client) *DeleteRunnergroupParams {
	return &DeleteRunnergroupParams{
		HTTPClient: client,
	}
}

/*
DeleteRunnergroupParams contains all the parameters to send to the API endpoint

	for the delete runnergroup operation.

	Typically these are written to a http.Request.
*/
type DeleteRunnergroupParams struct {

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

// WithDefaults hydrates default values in the delete runnergroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRunnergroupParams) WithDefaults() *DeleteRunnergroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete runnergroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRunnergroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete runnergroup params
func (o *DeleteRunnergroupParams) WithTimeout(timeout time.Duration) *DeleteRunnergroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runnergroup params
func (o *DeleteRunnergroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runnergroup params
func (o *DeleteRunnergroupParams) WithContext(ctx context.Context) *DeleteRunnergroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runnergroup params
func (o *DeleteRunnergroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runnergroup params
func (o *DeleteRunnergroupParams) WithHTTPClient(client *http.Client) *DeleteRunnergroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runnergroup params
func (o *DeleteRunnergroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the delete runnergroup params
func (o *DeleteRunnergroupParams) WithOrgName(orgName string) *DeleteRunnergroupParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the delete runnergroup params
func (o *DeleteRunnergroupParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithRunnergroupName adds the runnergroupName to the delete runnergroup params
func (o *DeleteRunnergroupParams) WithRunnergroupName(runnergroupName string) *DeleteRunnergroupParams {
	o.SetRunnergroupName(runnergroupName)
	return o
}

// SetRunnergroupName adds the runnergroupName to the delete runnergroup params
func (o *DeleteRunnergroupParams) SetRunnergroupName(runnergroupName string) {
	o.RunnergroupName = runnergroupName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunnergroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
