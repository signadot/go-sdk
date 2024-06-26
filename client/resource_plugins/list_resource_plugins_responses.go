// Code generated by go-swagger; DO NOT EDIT.

package resource_plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/signadot/go-sdk/models"
)

// ListResourcePluginsReader is a Reader for the ListResourcePlugins structure.
type ListResourcePluginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListResourcePluginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListResourcePluginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListResourcePluginsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListResourcePluginsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewListResourcePluginsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{orgName}/resource-plugins] list-resource-plugins", response, response.Code())
	}
}

// NewListResourcePluginsOK creates a ListResourcePluginsOK with default headers values
func NewListResourcePluginsOK() *ListResourcePluginsOK {
	return &ListResourcePluginsOK{}
}

/*
ListResourcePluginsOK describes a response with status code 200, with default header values.

OK
*/
type ListResourcePluginsOK struct {
	Payload []*models.ResourcePlugin
}

// IsSuccess returns true when this list resource plugins o k response has a 2xx status code
func (o *ListResourcePluginsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list resource plugins o k response has a 3xx status code
func (o *ListResourcePluginsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list resource plugins o k response has a 4xx status code
func (o *ListResourcePluginsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list resource plugins o k response has a 5xx status code
func (o *ListResourcePluginsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list resource plugins o k response a status code equal to that given
func (o *ListResourcePluginsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list resource plugins o k response
func (o *ListResourcePluginsOK) Code() int {
	return 200
}

func (o *ListResourcePluginsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins][%d] listResourcePluginsOK %s", 200, payload)
}

func (o *ListResourcePluginsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins][%d] listResourcePluginsOK %s", 200, payload)
}

func (o *ListResourcePluginsOK) GetPayload() []*models.ResourcePlugin {
	return o.Payload
}

func (o *ListResourcePluginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourcePluginsBadRequest creates a ListResourcePluginsBadRequest with default headers values
func NewListResourcePluginsBadRequest() *ListResourcePluginsBadRequest {
	return &ListResourcePluginsBadRequest{}
}

/*
ListResourcePluginsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListResourcePluginsBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list resource plugins bad request response has a 2xx status code
func (o *ListResourcePluginsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list resource plugins bad request response has a 3xx status code
func (o *ListResourcePluginsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list resource plugins bad request response has a 4xx status code
func (o *ListResourcePluginsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list resource plugins bad request response has a 5xx status code
func (o *ListResourcePluginsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list resource plugins bad request response a status code equal to that given
func (o *ListResourcePluginsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list resource plugins bad request response
func (o *ListResourcePluginsBadRequest) Code() int {
	return 400
}

func (o *ListResourcePluginsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins][%d] listResourcePluginsBadRequest %s", 400, payload)
}

func (o *ListResourcePluginsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins][%d] listResourcePluginsBadRequest %s", 400, payload)
}

func (o *ListResourcePluginsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListResourcePluginsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourcePluginsUnauthorized creates a ListResourcePluginsUnauthorized with default headers values
func NewListResourcePluginsUnauthorized() *ListResourcePluginsUnauthorized {
	return &ListResourcePluginsUnauthorized{}
}

/*
ListResourcePluginsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListResourcePluginsUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list resource plugins unauthorized response has a 2xx status code
func (o *ListResourcePluginsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list resource plugins unauthorized response has a 3xx status code
func (o *ListResourcePluginsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list resource plugins unauthorized response has a 4xx status code
func (o *ListResourcePluginsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list resource plugins unauthorized response has a 5xx status code
func (o *ListResourcePluginsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list resource plugins unauthorized response a status code equal to that given
func (o *ListResourcePluginsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list resource plugins unauthorized response
func (o *ListResourcePluginsUnauthorized) Code() int {
	return 401
}

func (o *ListResourcePluginsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins][%d] listResourcePluginsUnauthorized %s", 401, payload)
}

func (o *ListResourcePluginsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins][%d] listResourcePluginsUnauthorized %s", 401, payload)
}

func (o *ListResourcePluginsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListResourcePluginsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourcePluginsBadGateway creates a ListResourcePluginsBadGateway with default headers values
func NewListResourcePluginsBadGateway() *ListResourcePluginsBadGateway {
	return &ListResourcePluginsBadGateway{}
}

/*
ListResourcePluginsBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type ListResourcePluginsBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list resource plugins bad gateway response has a 2xx status code
func (o *ListResourcePluginsBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list resource plugins bad gateway response has a 3xx status code
func (o *ListResourcePluginsBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list resource plugins bad gateway response has a 4xx status code
func (o *ListResourcePluginsBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this list resource plugins bad gateway response has a 5xx status code
func (o *ListResourcePluginsBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this list resource plugins bad gateway response a status code equal to that given
func (o *ListResourcePluginsBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the list resource plugins bad gateway response
func (o *ListResourcePluginsBadGateway) Code() int {
	return 502
}

func (o *ListResourcePluginsBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins][%d] listResourcePluginsBadGateway %s", 502, payload)
}

func (o *ListResourcePluginsBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins][%d] listResourcePluginsBadGateway %s", 502, payload)
}

func (o *ListResourcePluginsBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListResourcePluginsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
