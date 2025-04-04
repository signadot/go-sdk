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

	"github.com/signadot/go-sdk/models"
)

// NewCreateHostedTestExecutionParams creates a new CreateHostedTestExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateHostedTestExecutionParams() *CreateHostedTestExecutionParams {
	return &CreateHostedTestExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHostedTestExecutionParamsWithTimeout creates a new CreateHostedTestExecutionParams object
// with the ability to set a timeout on a request.
func NewCreateHostedTestExecutionParamsWithTimeout(timeout time.Duration) *CreateHostedTestExecutionParams {
	return &CreateHostedTestExecutionParams{
		timeout: timeout,
	}
}

// NewCreateHostedTestExecutionParamsWithContext creates a new CreateHostedTestExecutionParams object
// with the ability to set a context for a request.
func NewCreateHostedTestExecutionParamsWithContext(ctx context.Context) *CreateHostedTestExecutionParams {
	return &CreateHostedTestExecutionParams{
		Context: ctx,
	}
}

// NewCreateHostedTestExecutionParamsWithHTTPClient creates a new CreateHostedTestExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateHostedTestExecutionParamsWithHTTPClient(client *http.Client) *CreateHostedTestExecutionParams {
	return &CreateHostedTestExecutionParams{
		HTTPClient: client,
	}
}

/*
CreateHostedTestExecutionParams contains all the parameters to send to the API endpoint

	for the create hosted test execution operation.

	Typically these are written to a http.Request.
*/
type CreateHostedTestExecutionParams struct {

	/* Data.

	   Request to create a test execution
	*/
	Data *models.TestExecution

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	/* TestName.

	   Test name
	*/
	TestName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create hosted test execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHostedTestExecutionParams) WithDefaults() *CreateHostedTestExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create hosted test execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHostedTestExecutionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) WithTimeout(timeout time.Duration) *CreateHostedTestExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) WithContext(ctx context.Context) *CreateHostedTestExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) WithHTTPClient(client *http.Client) *CreateHostedTestExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) WithData(data *models.TestExecution) *CreateHostedTestExecutionParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) SetData(data *models.TestExecution) {
	o.Data = data
}

// WithOrgName adds the orgName to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) WithOrgName(orgName string) *CreateHostedTestExecutionParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithTestName adds the testName to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) WithTestName(testName string) *CreateHostedTestExecutionParams {
	o.SetTestName(testName)
	return o
}

// SetTestName adds the testName to the create hosted test execution params
func (o *CreateHostedTestExecutionParams) SetTestName(testName string) {
	o.TestName = testName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHostedTestExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	// path param testName
	if err := r.SetPathParam("testName", o.TestName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
