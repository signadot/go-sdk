// Code generated by go-swagger; DO NOT EDIT.

package resource_plugins

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

// NewDeleteResourcePluginParams creates a new DeleteResourcePluginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteResourcePluginParams() *DeleteResourcePluginParams {
	return &DeleteResourcePluginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResourcePluginParamsWithTimeout creates a new DeleteResourcePluginParams object
// with the ability to set a timeout on a request.
func NewDeleteResourcePluginParamsWithTimeout(timeout time.Duration) *DeleteResourcePluginParams {
	return &DeleteResourcePluginParams{
		timeout: timeout,
	}
}

// NewDeleteResourcePluginParamsWithContext creates a new DeleteResourcePluginParams object
// with the ability to set a context for a request.
func NewDeleteResourcePluginParamsWithContext(ctx context.Context) *DeleteResourcePluginParams {
	return &DeleteResourcePluginParams{
		Context: ctx,
	}
}

// NewDeleteResourcePluginParamsWithHTTPClient creates a new DeleteResourcePluginParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteResourcePluginParamsWithHTTPClient(client *http.Client) *DeleteResourcePluginParams {
	return &DeleteResourcePluginParams{
		HTTPClient: client,
	}
}

/*
DeleteResourcePluginParams contains all the parameters to send to the API endpoint

	for the delete resource plugin operation.

	Typically these are written to a http.Request.
*/
type DeleteResourcePluginParams struct {

	/* OrgName.

	   Signadot Org Name
	*/
	OrgName string

	/* PluginName.

	   Resource plugin name
	*/
	PluginName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete resource plugin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourcePluginParams) WithDefaults() *DeleteResourcePluginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete resource plugin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourcePluginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete resource plugin params
func (o *DeleteResourcePluginParams) WithTimeout(timeout time.Duration) *DeleteResourcePluginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete resource plugin params
func (o *DeleteResourcePluginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete resource plugin params
func (o *DeleteResourcePluginParams) WithContext(ctx context.Context) *DeleteResourcePluginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete resource plugin params
func (o *DeleteResourcePluginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete resource plugin params
func (o *DeleteResourcePluginParams) WithHTTPClient(client *http.Client) *DeleteResourcePluginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete resource plugin params
func (o *DeleteResourcePluginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the delete resource plugin params
func (o *DeleteResourcePluginParams) WithOrgName(orgName string) *DeleteResourcePluginParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the delete resource plugin params
func (o *DeleteResourcePluginParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithPluginName adds the pluginName to the delete resource plugin params
func (o *DeleteResourcePluginParams) WithPluginName(pluginName string) *DeleteResourcePluginParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the delete resource plugin params
func (o *DeleteResourcePluginParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResourcePluginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
