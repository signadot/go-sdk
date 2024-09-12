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

// NewTestExecutionTrafficDiffParams creates a new TestExecutionTrafficDiffParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestExecutionTrafficDiffParams() *TestExecutionTrafficDiffParams {
	return &TestExecutionTrafficDiffParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestExecutionTrafficDiffParamsWithTimeout creates a new TestExecutionTrafficDiffParams object
// with the ability to set a timeout on a request.
func NewTestExecutionTrafficDiffParamsWithTimeout(timeout time.Duration) *TestExecutionTrafficDiffParams {
	return &TestExecutionTrafficDiffParams{
		timeout: timeout,
	}
}

// NewTestExecutionTrafficDiffParamsWithContext creates a new TestExecutionTrafficDiffParams object
// with the ability to set a context for a request.
func NewTestExecutionTrafficDiffParamsWithContext(ctx context.Context) *TestExecutionTrafficDiffParams {
	return &TestExecutionTrafficDiffParams{
		Context: ctx,
	}
}

// NewTestExecutionTrafficDiffParamsWithHTTPClient creates a new TestExecutionTrafficDiffParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestExecutionTrafficDiffParamsWithHTTPClient(client *http.Client) *TestExecutionTrafficDiffParams {
	return &TestExecutionTrafficDiffParams{
		HTTPClient: client,
	}
}

/*
TestExecutionTrafficDiffParams contains all the parameters to send to the API endpoint

	for the test execution traffic diff operation.

	Typically these are written to a http.Request.
*/
type TestExecutionTrafficDiffParams struct {

	/* ExecutionName.

	   Test Execution Name
	*/
	ExecutionName string

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	/* TestName.

	   Test Name
	*/
	TestName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test execution traffic diff params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestExecutionTrafficDiffParams) WithDefaults() *TestExecutionTrafficDiffParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test execution traffic diff params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestExecutionTrafficDiffParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) WithTimeout(timeout time.Duration) *TestExecutionTrafficDiffParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) WithContext(ctx context.Context) *TestExecutionTrafficDiffParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) WithHTTPClient(client *http.Client) *TestExecutionTrafficDiffParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionName adds the executionName to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) WithExecutionName(executionName string) *TestExecutionTrafficDiffParams {
	o.SetExecutionName(executionName)
	return o
}

// SetExecutionName adds the executionName to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) SetExecutionName(executionName string) {
	o.ExecutionName = executionName
}

// WithOrgName adds the orgName to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) WithOrgName(orgName string) *TestExecutionTrafficDiffParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithTestName adds the testName to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) WithTestName(testName string) *TestExecutionTrafficDiffParams {
	o.SetTestName(testName)
	return o
}

// SetTestName adds the testName to the test execution traffic diff params
func (o *TestExecutionTrafficDiffParams) SetTestName(testName string) {
	o.TestName = testName
}

// WriteToRequest writes these params to a swagger request
func (o *TestExecutionTrafficDiffParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param executionName
	if err := r.SetPathParam("executionName", o.ExecutionName); err != nil {
		return err
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