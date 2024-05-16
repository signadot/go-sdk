// Code generated by go-swagger; DO NOT EDIT.

package sandboxes

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

// DeleteSandboxReader is a Reader for the DeleteSandbox structure.
type DeleteSandboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSandboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSandboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSandboxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSandboxUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewDeleteSandboxBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}] delete-sandbox", response, response.Code())
	}
}

// NewDeleteSandboxOK creates a DeleteSandboxOK with default headers values
func NewDeleteSandboxOK() *DeleteSandboxOK {
	return &DeleteSandboxOK{}
}

/*
DeleteSandboxOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSandboxOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this delete sandbox o k response has a 2xx status code
func (o *DeleteSandboxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete sandbox o k response has a 3xx status code
func (o *DeleteSandboxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sandbox o k response has a 4xx status code
func (o *DeleteSandboxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete sandbox o k response has a 5xx status code
func (o *DeleteSandboxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sandbox o k response a status code equal to that given
func (o *DeleteSandboxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete sandbox o k response
func (o *DeleteSandboxOK) Code() int {
	return 200
}

func (o *DeleteSandboxOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}][%d] deleteSandboxOK %s", 200, payload)
}

func (o *DeleteSandboxOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}][%d] deleteSandboxOK %s", 200, payload)
}

func (o *DeleteSandboxOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteSandboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSandboxBadRequest creates a DeleteSandboxBadRequest with default headers values
func NewDeleteSandboxBadRequest() *DeleteSandboxBadRequest {
	return &DeleteSandboxBadRequest{}
}

/*
DeleteSandboxBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteSandboxBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete sandbox bad request response has a 2xx status code
func (o *DeleteSandboxBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sandbox bad request response has a 3xx status code
func (o *DeleteSandboxBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sandbox bad request response has a 4xx status code
func (o *DeleteSandboxBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sandbox bad request response has a 5xx status code
func (o *DeleteSandboxBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sandbox bad request response a status code equal to that given
func (o *DeleteSandboxBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete sandbox bad request response
func (o *DeleteSandboxBadRequest) Code() int {
	return 400
}

func (o *DeleteSandboxBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}][%d] deleteSandboxBadRequest %s", 400, payload)
}

func (o *DeleteSandboxBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}][%d] deleteSandboxBadRequest %s", 400, payload)
}

func (o *DeleteSandboxBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteSandboxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSandboxUnauthorized creates a DeleteSandboxUnauthorized with default headers values
func NewDeleteSandboxUnauthorized() *DeleteSandboxUnauthorized {
	return &DeleteSandboxUnauthorized{}
}

/*
DeleteSandboxUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteSandboxUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete sandbox unauthorized response has a 2xx status code
func (o *DeleteSandboxUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sandbox unauthorized response has a 3xx status code
func (o *DeleteSandboxUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sandbox unauthorized response has a 4xx status code
func (o *DeleteSandboxUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sandbox unauthorized response has a 5xx status code
func (o *DeleteSandboxUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sandbox unauthorized response a status code equal to that given
func (o *DeleteSandboxUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete sandbox unauthorized response
func (o *DeleteSandboxUnauthorized) Code() int {
	return 401
}

func (o *DeleteSandboxUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}][%d] deleteSandboxUnauthorized %s", 401, payload)
}

func (o *DeleteSandboxUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}][%d] deleteSandboxUnauthorized %s", 401, payload)
}

func (o *DeleteSandboxUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteSandboxUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSandboxBadGateway creates a DeleteSandboxBadGateway with default headers values
func NewDeleteSandboxBadGateway() *DeleteSandboxBadGateway {
	return &DeleteSandboxBadGateway{}
}

/*
DeleteSandboxBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type DeleteSandboxBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete sandbox bad gateway response has a 2xx status code
func (o *DeleteSandboxBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sandbox bad gateway response has a 3xx status code
func (o *DeleteSandboxBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sandbox bad gateway response has a 4xx status code
func (o *DeleteSandboxBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete sandbox bad gateway response has a 5xx status code
func (o *DeleteSandboxBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this delete sandbox bad gateway response a status code equal to that given
func (o *DeleteSandboxBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the delete sandbox bad gateway response
func (o *DeleteSandboxBadGateway) Code() int {
	return 502
}

func (o *DeleteSandboxBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}][%d] deleteSandboxBadGateway %s", 502, payload)
}

func (o *DeleteSandboxBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/sandboxes/{sandboxName}][%d] deleteSandboxBadGateway %s", 502, payload)
}

func (o *DeleteSandboxBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteSandboxBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
