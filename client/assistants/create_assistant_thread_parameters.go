// Code generated by go-swagger; DO NOT EDIT.

package assistants

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

// NewCreateAssistantThreadParams creates a new CreateAssistantThreadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAssistantThreadParams() *CreateAssistantThreadParams {
	return &CreateAssistantThreadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAssistantThreadParamsWithTimeout creates a new CreateAssistantThreadParams object
// with the ability to set a timeout on a request.
func NewCreateAssistantThreadParamsWithTimeout(timeout time.Duration) *CreateAssistantThreadParams {
	return &CreateAssistantThreadParams{
		timeout: timeout,
	}
}

// NewCreateAssistantThreadParamsWithContext creates a new CreateAssistantThreadParams object
// with the ability to set a context for a request.
func NewCreateAssistantThreadParamsWithContext(ctx context.Context) *CreateAssistantThreadParams {
	return &CreateAssistantThreadParams{
		Context: ctx,
	}
}

// NewCreateAssistantThreadParamsWithHTTPClient creates a new CreateAssistantThreadParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAssistantThreadParamsWithHTTPClient(client *http.Client) *CreateAssistantThreadParams {
	return &CreateAssistantThreadParams{
		HTTPClient: client,
	}
}

/*
CreateAssistantThreadParams contains all the parameters to send to the API endpoint

	for the create assistant thread operation.

	Typically these are written to a http.Request.
*/
type CreateAssistantThreadParams struct {

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create assistant thread params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAssistantThreadParams) WithDefaults() *CreateAssistantThreadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create assistant thread params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAssistantThreadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create assistant thread params
func (o *CreateAssistantThreadParams) WithTimeout(timeout time.Duration) *CreateAssistantThreadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create assistant thread params
func (o *CreateAssistantThreadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create assistant thread params
func (o *CreateAssistantThreadParams) WithContext(ctx context.Context) *CreateAssistantThreadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create assistant thread params
func (o *CreateAssistantThreadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create assistant thread params
func (o *CreateAssistantThreadParams) WithHTTPClient(client *http.Client) *CreateAssistantThreadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create assistant thread params
func (o *CreateAssistantThreadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the create assistant thread params
func (o *CreateAssistantThreadParams) WithOrgName(orgName string) *CreateAssistantThreadParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the create assistant thread params
func (o *CreateAssistantThreadParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAssistantThreadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
