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

// DeleteRoutegroupReader is a Reader for the DeleteRoutegroup structure.
type DeleteRoutegroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutegroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoutegroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutegroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutegroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewDeleteRoutegroupBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /orgs/{orgName}/routegroups/{routegroupName}] delete-routegroup", response, response.Code())
	}
}

// NewDeleteRoutegroupOK creates a DeleteRoutegroupOK with default headers values
func NewDeleteRoutegroupOK() *DeleteRoutegroupOK {
	return &DeleteRoutegroupOK{}
}

/* DeleteRoutegroupOK describes a response with status code 200, with default header values.

OK
*/
type DeleteRoutegroupOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this delete routegroup o k response has a 2xx status code
func (o *DeleteRoutegroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routegroup o k response has a 3xx status code
func (o *DeleteRoutegroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routegroup o k response has a 4xx status code
func (o *DeleteRoutegroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routegroup o k response has a 5xx status code
func (o *DeleteRoutegroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routegroup o k response a status code equal to that given
func (o *DeleteRoutegroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete routegroup o k response
func (o *DeleteRoutegroupOK) Code() int {
	return 200
}

func (o *DeleteRoutegroupOK) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/routegroups/{routegroupName}][%d] deleteRoutegroupOK  %+v", 200, o.Payload)
}

func (o *DeleteRoutegroupOK) String() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/routegroups/{routegroupName}][%d] deleteRoutegroupOK  %+v", 200, o.Payload)
}

func (o *DeleteRoutegroupOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteRoutegroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutegroupBadRequest creates a DeleteRoutegroupBadRequest with default headers values
func NewDeleteRoutegroupBadRequest() *DeleteRoutegroupBadRequest {
	return &DeleteRoutegroupBadRequest{}
}

/* DeleteRoutegroupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteRoutegroupBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete routegroup bad request response has a 2xx status code
func (o *DeleteRoutegroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routegroup bad request response has a 3xx status code
func (o *DeleteRoutegroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routegroup bad request response has a 4xx status code
func (o *DeleteRoutegroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routegroup bad request response has a 5xx status code
func (o *DeleteRoutegroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routegroup bad request response a status code equal to that given
func (o *DeleteRoutegroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete routegroup bad request response
func (o *DeleteRoutegroupBadRequest) Code() int {
	return 400
}

func (o *DeleteRoutegroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/routegroups/{routegroupName}][%d] deleteRoutegroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutegroupBadRequest) String() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/routegroups/{routegroupName}][%d] deleteRoutegroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutegroupBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRoutegroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutegroupUnauthorized creates a DeleteRoutegroupUnauthorized with default headers values
func NewDeleteRoutegroupUnauthorized() *DeleteRoutegroupUnauthorized {
	return &DeleteRoutegroupUnauthorized{}
}

/* DeleteRoutegroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRoutegroupUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete routegroup unauthorized response has a 2xx status code
func (o *DeleteRoutegroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routegroup unauthorized response has a 3xx status code
func (o *DeleteRoutegroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routegroup unauthorized response has a 4xx status code
func (o *DeleteRoutegroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routegroup unauthorized response has a 5xx status code
func (o *DeleteRoutegroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routegroup unauthorized response a status code equal to that given
func (o *DeleteRoutegroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete routegroup unauthorized response
func (o *DeleteRoutegroupUnauthorized) Code() int {
	return 401
}

func (o *DeleteRoutegroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/routegroups/{routegroupName}][%d] deleteRoutegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutegroupUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/routegroups/{routegroupName}][%d] deleteRoutegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutegroupUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRoutegroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutegroupBadGateway creates a DeleteRoutegroupBadGateway with default headers values
func NewDeleteRoutegroupBadGateway() *DeleteRoutegroupBadGateway {
	return &DeleteRoutegroupBadGateway{}
}

/* DeleteRoutegroupBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type DeleteRoutegroupBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete routegroup bad gateway response has a 2xx status code
func (o *DeleteRoutegroupBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routegroup bad gateway response has a 3xx status code
func (o *DeleteRoutegroupBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routegroup bad gateway response has a 4xx status code
func (o *DeleteRoutegroupBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routegroup bad gateway response has a 5xx status code
func (o *DeleteRoutegroupBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routegroup bad gateway response a status code equal to that given
func (o *DeleteRoutegroupBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the delete routegroup bad gateway response
func (o *DeleteRoutegroupBadGateway) Code() int {
	return 502
}

func (o *DeleteRoutegroupBadGateway) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/routegroups/{routegroupName}][%d] deleteRoutegroupBadGateway  %+v", 502, o.Payload)
}

func (o *DeleteRoutegroupBadGateway) String() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/routegroups/{routegroupName}][%d] deleteRoutegroupBadGateway  %+v", 502, o.Payload)
}

func (o *DeleteRoutegroupBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRoutegroupBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
