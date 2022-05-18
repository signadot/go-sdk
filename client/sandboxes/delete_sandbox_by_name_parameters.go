// Code generated by go-swagger; DO NOT EDIT.

package sandboxes

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

// NewDeleteSandboxByNameParams creates a new DeleteSandboxByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSandboxByNameParams() *DeleteSandboxByNameParams {
	return &DeleteSandboxByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSandboxByNameParamsWithTimeout creates a new DeleteSandboxByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteSandboxByNameParamsWithTimeout(timeout time.Duration) *DeleteSandboxByNameParams {
	return &DeleteSandboxByNameParams{
		timeout: timeout,
	}
}

// NewDeleteSandboxByNameParamsWithContext creates a new DeleteSandboxByNameParams object
// with the ability to set a context for a request.
func NewDeleteSandboxByNameParamsWithContext(ctx context.Context) *DeleteSandboxByNameParams {
	return &DeleteSandboxByNameParams{
		Context: ctx,
	}
}

// NewDeleteSandboxByNameParamsWithHTTPClient creates a new DeleteSandboxByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSandboxByNameParamsWithHTTPClient(client *http.Client) *DeleteSandboxByNameParams {
	return &DeleteSandboxByNameParams{
		HTTPClient: client,
	}
}

/* DeleteSandboxByNameParams contains all the parameters to send to the API endpoint
   for the delete sandbox by name operation.

   Typically these are written to a http.Request.
*/
type DeleteSandboxByNameParams struct {

	/* Name.

	   Sandbox Name to search for
	*/
	Name string

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete sandbox by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSandboxByNameParams) WithDefaults() *DeleteSandboxByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete sandbox by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSandboxByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) WithTimeout(timeout time.Duration) *DeleteSandboxByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) WithContext(ctx context.Context) *DeleteSandboxByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) WithHTTPClient(client *http.Client) *DeleteSandboxByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) WithName(name string) *DeleteSandboxByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) SetName(name string) {
	o.Name = name
}

// WithOrgName adds the orgName to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) WithOrgName(orgName string) *DeleteSandboxByNameParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the delete sandbox by name params
func (o *DeleteSandboxByNameParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSandboxByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
