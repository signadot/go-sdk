// Code generated by go-swagger; DO NOT EDIT.

package tests

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

// NewApplyTestParams creates a new ApplyTestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplyTestParams() *ApplyTestParams {
	return &ApplyTestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplyTestParamsWithTimeout creates a new ApplyTestParams object
// with the ability to set a timeout on a request.
func NewApplyTestParamsWithTimeout(timeout time.Duration) *ApplyTestParams {
	return &ApplyTestParams{
		timeout: timeout,
	}
}

// NewApplyTestParamsWithContext creates a new ApplyTestParams object
// with the ability to set a context for a request.
func NewApplyTestParamsWithContext(ctx context.Context) *ApplyTestParams {
	return &ApplyTestParams{
		Context: ctx,
	}
}

// NewApplyTestParamsWithHTTPClient creates a new ApplyTestParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplyTestParamsWithHTTPClient(client *http.Client) *ApplyTestParams {
	return &ApplyTestParams{
		HTTPClient: client,
	}
}

/*
ApplyTestParams contains all the parameters to send to the API endpoint

	for the apply test operation.

	Typically these are written to a http.Request.
*/
type ApplyTestParams struct {

	/* Data.

	   Test Spec
	*/
	Data *models.TestSpec

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

// WithDefaults hydrates default values in the apply test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplyTestParams) WithDefaults() *ApplyTestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the apply test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplyTestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the apply test params
func (o *ApplyTestParams) WithTimeout(timeout time.Duration) *ApplyTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the apply test params
func (o *ApplyTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the apply test params
func (o *ApplyTestParams) WithContext(ctx context.Context) *ApplyTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the apply test params
func (o *ApplyTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the apply test params
func (o *ApplyTestParams) WithHTTPClient(client *http.Client) *ApplyTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the apply test params
func (o *ApplyTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the apply test params
func (o *ApplyTestParams) WithData(data *models.TestSpec) *ApplyTestParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the apply test params
func (o *ApplyTestParams) SetData(data *models.TestSpec) {
	o.Data = data
}

// WithOrgName adds the orgName to the apply test params
func (o *ApplyTestParams) WithOrgName(orgName string) *ApplyTestParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the apply test params
func (o *ApplyTestParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithTestName adds the testName to the apply test params
func (o *ApplyTestParams) WithTestName(testName string) *ApplyTestParams {
	o.SetTestName(testName)
	return o
}

// SetTestName adds the testName to the apply test params
func (o *ApplyTestParams) SetTestName(testName string) {
	o.TestName = testName
}

// WriteToRequest writes these params to a swagger request
func (o *ApplyTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
