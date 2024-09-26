// Code generated by go-swagger; DO NOT EDIT.

package jobs

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

// NewCancelJobParams creates a new CancelJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelJobParams() *CancelJobParams {
	return &CancelJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelJobParamsWithTimeout creates a new CancelJobParams object
// with the ability to set a timeout on a request.
func NewCancelJobParamsWithTimeout(timeout time.Duration) *CancelJobParams {
	return &CancelJobParams{
		timeout: timeout,
	}
}

// NewCancelJobParamsWithContext creates a new CancelJobParams object
// with the ability to set a context for a request.
func NewCancelJobParamsWithContext(ctx context.Context) *CancelJobParams {
	return &CancelJobParams{
		Context: ctx,
	}
}

// NewCancelJobParamsWithHTTPClient creates a new CancelJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelJobParamsWithHTTPClient(client *http.Client) *CancelJobParams {
	return &CancelJobParams{
		HTTPClient: client,
	}
}

/*
CancelJobParams contains all the parameters to send to the API endpoint

	for the cancel job operation.

	Typically these are written to a http.Request.
*/
type CancelJobParams struct {

	/* JobName.

	   Job Name
	*/
	JobName string

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelJobParams) WithDefaults() *CancelJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel job params
func (o *CancelJobParams) WithTimeout(timeout time.Duration) *CancelJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel job params
func (o *CancelJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel job params
func (o *CancelJobParams) WithContext(ctx context.Context) *CancelJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel job params
func (o *CancelJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel job params
func (o *CancelJobParams) WithHTTPClient(client *http.Client) *CancelJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel job params
func (o *CancelJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobName adds the jobName to the cancel job params
func (o *CancelJobParams) WithJobName(jobName string) *CancelJobParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the cancel job params
func (o *CancelJobParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WithOrgName adds the orgName to the cancel job params
func (o *CancelJobParams) WithOrgName(orgName string) *CancelJobParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the cancel job params
func (o *CancelJobParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WriteToRequest writes these params to a swagger request
func (o *CancelJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param jobName
	if err := r.SetPathParam("jobName", o.JobName); err != nil {
		return err
	}

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
