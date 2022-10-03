// Code generated by go-swagger; DO NOT EDIT.

package sandboxes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/signadot/go-sdk/models"
)

// ApplySandboxReader is a Reader for the ApplySandbox structure.
type ApplySandboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplySandboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplySandboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApplySandboxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewApplySandboxUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewApplySandboxBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewApplySandboxOK creates a ApplySandboxOK with default headers values
func NewApplySandboxOK() *ApplySandboxOK {
	return &ApplySandboxOK{}
}

/* ApplySandboxOK describes a response with status code 200, with default header values.

OK
*/
type ApplySandboxOK struct {
	Payload *models.Sandbox
}

// IsSuccess returns true when this apply sandbox o k response has a 2xx status code
func (o *ApplySandboxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this apply sandbox o k response has a 3xx status code
func (o *ApplySandboxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply sandbox o k response has a 4xx status code
func (o *ApplySandboxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this apply sandbox o k response has a 5xx status code
func (o *ApplySandboxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this apply sandbox o k response a status code equal to that given
func (o *ApplySandboxOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApplySandboxOK) Error() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/sandboxes/{sandboxName}][%d] applySandboxOK  %+v", 200, o.Payload)
}

func (o *ApplySandboxOK) String() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/sandboxes/{sandboxName}][%d] applySandboxOK  %+v", 200, o.Payload)
}

func (o *ApplySandboxOK) GetPayload() *models.Sandbox {
	return o.Payload
}

func (o *ApplySandboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Sandbox)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplySandboxBadRequest creates a ApplySandboxBadRequest with default headers values
func NewApplySandboxBadRequest() *ApplySandboxBadRequest {
	return &ApplySandboxBadRequest{}
}

/* ApplySandboxBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ApplySandboxBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this apply sandbox bad request response has a 2xx status code
func (o *ApplySandboxBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this apply sandbox bad request response has a 3xx status code
func (o *ApplySandboxBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply sandbox bad request response has a 4xx status code
func (o *ApplySandboxBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this apply sandbox bad request response has a 5xx status code
func (o *ApplySandboxBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this apply sandbox bad request response a status code equal to that given
func (o *ApplySandboxBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ApplySandboxBadRequest) Error() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/sandboxes/{sandboxName}][%d] applySandboxBadRequest  %+v", 400, o.Payload)
}

func (o *ApplySandboxBadRequest) String() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/sandboxes/{sandboxName}][%d] applySandboxBadRequest  %+v", 400, o.Payload)
}

func (o *ApplySandboxBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplySandboxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplySandboxUnauthorized creates a ApplySandboxUnauthorized with default headers values
func NewApplySandboxUnauthorized() *ApplySandboxUnauthorized {
	return &ApplySandboxUnauthorized{}
}

/* ApplySandboxUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ApplySandboxUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this apply sandbox unauthorized response has a 2xx status code
func (o *ApplySandboxUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this apply sandbox unauthorized response has a 3xx status code
func (o *ApplySandboxUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply sandbox unauthorized response has a 4xx status code
func (o *ApplySandboxUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this apply sandbox unauthorized response has a 5xx status code
func (o *ApplySandboxUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this apply sandbox unauthorized response a status code equal to that given
func (o *ApplySandboxUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ApplySandboxUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/sandboxes/{sandboxName}][%d] applySandboxUnauthorized  %+v", 401, o.Payload)
}

func (o *ApplySandboxUnauthorized) String() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/sandboxes/{sandboxName}][%d] applySandboxUnauthorized  %+v", 401, o.Payload)
}

func (o *ApplySandboxUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplySandboxUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplySandboxBadGateway creates a ApplySandboxBadGateway with default headers values
func NewApplySandboxBadGateway() *ApplySandboxBadGateway {
	return &ApplySandboxBadGateway{}
}

/* ApplySandboxBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type ApplySandboxBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this apply sandbox bad gateway response has a 2xx status code
func (o *ApplySandboxBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this apply sandbox bad gateway response has a 3xx status code
func (o *ApplySandboxBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply sandbox bad gateway response has a 4xx status code
func (o *ApplySandboxBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this apply sandbox bad gateway response has a 5xx status code
func (o *ApplySandboxBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this apply sandbox bad gateway response a status code equal to that given
func (o *ApplySandboxBadGateway) IsCode(code int) bool {
	return code == 502
}

func (o *ApplySandboxBadGateway) Error() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/sandboxes/{sandboxName}][%d] applySandboxBadGateway  %+v", 502, o.Payload)
}

func (o *ApplySandboxBadGateway) String() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/sandboxes/{sandboxName}][%d] applySandboxBadGateway  %+v", 502, o.Payload)
}

func (o *ApplySandboxBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplySandboxBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
