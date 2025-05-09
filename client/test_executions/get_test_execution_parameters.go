// Code generated by go-swagger; DO NOT EDIT.

package test_executions

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

// NewGetTestExecutionParams creates a new GetTestExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTestExecutionParams() *GetTestExecutionParams {
	return &GetTestExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestExecutionParamsWithTimeout creates a new GetTestExecutionParams object
// with the ability to set a timeout on a request.
func NewGetTestExecutionParamsWithTimeout(timeout time.Duration) *GetTestExecutionParams {
	return &GetTestExecutionParams{
		timeout: timeout,
	}
}

// NewGetTestExecutionParamsWithContext creates a new GetTestExecutionParams object
// with the ability to set a context for a request.
func NewGetTestExecutionParamsWithContext(ctx context.Context) *GetTestExecutionParams {
	return &GetTestExecutionParams{
		Context: ctx,
	}
}

// NewGetTestExecutionParamsWithHTTPClient creates a new GetTestExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTestExecutionParamsWithHTTPClient(client *http.Client) *GetTestExecutionParams {
	return &GetTestExecutionParams{
		HTTPClient: client,
	}
}

/*
GetTestExecutionParams contains all the parameters to send to the API endpoint

	for the get test execution operation.

	Typically these are written to a http.Request.
*/
type GetTestExecutionParams struct {

	/* ExecutionID.

	   Test Execution ID
	*/
	ExecutionID string

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get test execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTestExecutionParams) WithDefaults() *GetTestExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get test execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTestExecutionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get test execution params
func (o *GetTestExecutionParams) WithTimeout(timeout time.Duration) *GetTestExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test execution params
func (o *GetTestExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test execution params
func (o *GetTestExecutionParams) WithContext(ctx context.Context) *GetTestExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test execution params
func (o *GetTestExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test execution params
func (o *GetTestExecutionParams) WithHTTPClient(client *http.Client) *GetTestExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test execution params
func (o *GetTestExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get test execution params
func (o *GetTestExecutionParams) WithExecutionID(executionID string) *GetTestExecutionParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get test execution params
func (o *GetTestExecutionParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithOrgName adds the orgName to the get test execution params
func (o *GetTestExecutionParams) WithOrgName(orgName string) *GetTestExecutionParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the get test execution params
func (o *GetTestExecutionParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param executionID
	if err := r.SetPathParam("executionID", o.ExecutionID); err != nil {
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
