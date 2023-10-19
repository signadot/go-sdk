// Code generated by go-swagger; DO NOT EDIT.

package route_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/signadot/go-sdk/models"
)

// GetRoutegroupReader is a Reader for the GetRoutegroup structure.
type GetRoutegroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutegroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutegroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutegroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutegroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetRoutegroupBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{orgName}/routegroups/{routegroupName}] get-routegroup", response, response.Code())
	}
}

// NewGetRoutegroupOK creates a GetRoutegroupOK with default headers values
func NewGetRoutegroupOK() *GetRoutegroupOK {
	return &GetRoutegroupOK{}
}

/*
GetRoutegroupOK describes a response with status code 200, with default header values.

OK
*/
type GetRoutegroupOK struct {
	Payload *models.RouteGroup
}

// IsSuccess returns true when this get routegroup o k response has a 2xx status code
func (o *GetRoutegroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routegroup o k response has a 3xx status code
func (o *GetRoutegroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routegroup o k response has a 4xx status code
func (o *GetRoutegroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routegroup o k response has a 5xx status code
func (o *GetRoutegroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routegroup o k response a status code equal to that given
func (o *GetRoutegroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get routegroup o k response
func (o *GetRoutegroupOK) Code() int {
	return 200
}

func (o *GetRoutegroupOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/routegroups/{routegroupName}][%d] getRoutegroupOK  %+v", 200, o.Payload)
}

func (o *GetRoutegroupOK) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/routegroups/{routegroupName}][%d] getRoutegroupOK  %+v", 200, o.Payload)
}

func (o *GetRoutegroupOK) GetPayload() *models.RouteGroup {
	return o.Payload
}

func (o *GetRoutegroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutegroupBadRequest creates a GetRoutegroupBadRequest with default headers values
func NewGetRoutegroupBadRequest() *GetRoutegroupBadRequest {
	return &GetRoutegroupBadRequest{}
}

/*
GetRoutegroupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRoutegroupBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get routegroup bad request response has a 2xx status code
func (o *GetRoutegroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routegroup bad request response has a 3xx status code
func (o *GetRoutegroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routegroup bad request response has a 4xx status code
func (o *GetRoutegroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routegroup bad request response has a 5xx status code
func (o *GetRoutegroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routegroup bad request response a status code equal to that given
func (o *GetRoutegroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get routegroup bad request response
func (o *GetRoutegroupBadRequest) Code() int {
	return 400
}

func (o *GetRoutegroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/routegroups/{routegroupName}][%d] getRoutegroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutegroupBadRequest) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/routegroups/{routegroupName}][%d] getRoutegroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutegroupBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRoutegroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutegroupUnauthorized creates a GetRoutegroupUnauthorized with default headers values
func NewGetRoutegroupUnauthorized() *GetRoutegroupUnauthorized {
	return &GetRoutegroupUnauthorized{}
}

/*
GetRoutegroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRoutegroupUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get routegroup unauthorized response has a 2xx status code
func (o *GetRoutegroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routegroup unauthorized response has a 3xx status code
func (o *GetRoutegroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routegroup unauthorized response has a 4xx status code
func (o *GetRoutegroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routegroup unauthorized response has a 5xx status code
func (o *GetRoutegroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routegroup unauthorized response a status code equal to that given
func (o *GetRoutegroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get routegroup unauthorized response
func (o *GetRoutegroupUnauthorized) Code() int {
	return 401
}

func (o *GetRoutegroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/routegroups/{routegroupName}][%d] getRoutegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutegroupUnauthorized) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/routegroups/{routegroupName}][%d] getRoutegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutegroupUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRoutegroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutegroupBadGateway creates a GetRoutegroupBadGateway with default headers values
func NewGetRoutegroupBadGateway() *GetRoutegroupBadGateway {
	return &GetRoutegroupBadGateway{}
}

/*
GetRoutegroupBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type GetRoutegroupBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get routegroup bad gateway response has a 2xx status code
func (o *GetRoutegroupBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routegroup bad gateway response has a 3xx status code
func (o *GetRoutegroupBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routegroup bad gateway response has a 4xx status code
func (o *GetRoutegroupBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routegroup bad gateway response has a 5xx status code
func (o *GetRoutegroupBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this get routegroup bad gateway response a status code equal to that given
func (o *GetRoutegroupBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the get routegroup bad gateway response
func (o *GetRoutegroupBadGateway) Code() int {
	return 502
}

func (o *GetRoutegroupBadGateway) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/routegroups/{routegroupName}][%d] getRoutegroupBadGateway  %+v", 502, o.Payload)
}

func (o *GetRoutegroupBadGateway) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/routegroups/{routegroupName}][%d] getRoutegroupBadGateway  %+v", 502, o.Payload)
}

func (o *GetRoutegroupBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRoutegroupBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
