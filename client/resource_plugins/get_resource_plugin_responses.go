// Code generated by go-swagger; DO NOT EDIT.

package resource_plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/signadot/go-sdk/models"
)

// GetResourcePluginReader is a Reader for the GetResourcePlugin structure.
type GetResourcePluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcePluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcePluginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResourcePluginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResourcePluginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetResourcePluginBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourcePluginOK creates a GetResourcePluginOK with default headers values
func NewGetResourcePluginOK() *GetResourcePluginOK {
	return &GetResourcePluginOK{}
}

/*
GetResourcePluginOK describes a response with status code 200, with default header values.

OK
*/
type GetResourcePluginOK struct {
	Payload *models.ResourcePlugin
}

// IsSuccess returns true when this get resource plugin o k response has a 2xx status code
func (o *GetResourcePluginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource plugin o k response has a 3xx status code
func (o *GetResourcePluginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource plugin o k response has a 4xx status code
func (o *GetResourcePluginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource plugin o k response has a 5xx status code
func (o *GetResourcePluginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource plugin o k response a status code equal to that given
func (o *GetResourcePluginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get resource plugin o k response
func (o *GetResourcePluginOK) Code() int {
	return 200
}

func (o *GetResourcePluginOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins/{pluginName}][%d] getResourcePluginOK  %+v", 200, o.Payload)
}

func (o *GetResourcePluginOK) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins/{pluginName}][%d] getResourcePluginOK  %+v", 200, o.Payload)
}

func (o *GetResourcePluginOK) GetPayload() *models.ResourcePlugin {
	return o.Payload
}

func (o *GetResourcePluginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourcePlugin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcePluginBadRequest creates a GetResourcePluginBadRequest with default headers values
func NewGetResourcePluginBadRequest() *GetResourcePluginBadRequest {
	return &GetResourcePluginBadRequest{}
}

/*
GetResourcePluginBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetResourcePluginBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get resource plugin bad request response has a 2xx status code
func (o *GetResourcePluginBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource plugin bad request response has a 3xx status code
func (o *GetResourcePluginBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource plugin bad request response has a 4xx status code
func (o *GetResourcePluginBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource plugin bad request response has a 5xx status code
func (o *GetResourcePluginBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource plugin bad request response a status code equal to that given
func (o *GetResourcePluginBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get resource plugin bad request response
func (o *GetResourcePluginBadRequest) Code() int {
	return 400
}

func (o *GetResourcePluginBadRequest) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins/{pluginName}][%d] getResourcePluginBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourcePluginBadRequest) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins/{pluginName}][%d] getResourcePluginBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourcePluginBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourcePluginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcePluginUnauthorized creates a GetResourcePluginUnauthorized with default headers values
func NewGetResourcePluginUnauthorized() *GetResourcePluginUnauthorized {
	return &GetResourcePluginUnauthorized{}
}

/*
GetResourcePluginUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetResourcePluginUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get resource plugin unauthorized response has a 2xx status code
func (o *GetResourcePluginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource plugin unauthorized response has a 3xx status code
func (o *GetResourcePluginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource plugin unauthorized response has a 4xx status code
func (o *GetResourcePluginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource plugin unauthorized response has a 5xx status code
func (o *GetResourcePluginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource plugin unauthorized response a status code equal to that given
func (o *GetResourcePluginUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get resource plugin unauthorized response
func (o *GetResourcePluginUnauthorized) Code() int {
	return 401
}

func (o *GetResourcePluginUnauthorized) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins/{pluginName}][%d] getResourcePluginUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResourcePluginUnauthorized) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins/{pluginName}][%d] getResourcePluginUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResourcePluginUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourcePluginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcePluginBadGateway creates a GetResourcePluginBadGateway with default headers values
func NewGetResourcePluginBadGateway() *GetResourcePluginBadGateway {
	return &GetResourcePluginBadGateway{}
}

/*
GetResourcePluginBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type GetResourcePluginBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get resource plugin bad gateway response has a 2xx status code
func (o *GetResourcePluginBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource plugin bad gateway response has a 3xx status code
func (o *GetResourcePluginBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource plugin bad gateway response has a 4xx status code
func (o *GetResourcePluginBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource plugin bad gateway response has a 5xx status code
func (o *GetResourcePluginBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this get resource plugin bad gateway response a status code equal to that given
func (o *GetResourcePluginBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the get resource plugin bad gateway response
func (o *GetResourcePluginBadGateway) Code() int {
	return 502
}

func (o *GetResourcePluginBadGateway) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins/{pluginName}][%d] getResourcePluginBadGateway  %+v", 502, o.Payload)
}

func (o *GetResourcePluginBadGateway) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/resource-plugins/{pluginName}][%d] getResourcePluginBadGateway  %+v", 502, o.Payload)
}

func (o *GetResourcePluginBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourcePluginBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
