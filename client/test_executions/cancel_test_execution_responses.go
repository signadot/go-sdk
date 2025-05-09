// Code generated by go-swagger; DO NOT EDIT.

package test_executions

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

// CancelTestExecutionReader is a Reader for the CancelTestExecution structure.
type CancelTestExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelTestExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelTestExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelTestExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelTestExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewCancelTestExecutionBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel] cancel-test-execution", response, response.Code())
	}
}

// NewCancelTestExecutionOK creates a CancelTestExecutionOK with default headers values
func NewCancelTestExecutionOK() *CancelTestExecutionOK {
	return &CancelTestExecutionOK{}
}

/*
CancelTestExecutionOK describes a response with status code 200, with default header values.

OK
*/
type CancelTestExecutionOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this cancel test execution o k response has a 2xx status code
func (o *CancelTestExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel test execution o k response has a 3xx status code
func (o *CancelTestExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel test execution o k response has a 4xx status code
func (o *CancelTestExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel test execution o k response has a 5xx status code
func (o *CancelTestExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel test execution o k response a status code equal to that given
func (o *CancelTestExecutionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cancel test execution o k response
func (o *CancelTestExecutionOK) Code() int {
	return 200
}

func (o *CancelTestExecutionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel][%d] cancelTestExecutionOK %s", 200, payload)
}

func (o *CancelTestExecutionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel][%d] cancelTestExecutionOK %s", 200, payload)
}

func (o *CancelTestExecutionOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *CancelTestExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelTestExecutionBadRequest creates a CancelTestExecutionBadRequest with default headers values
func NewCancelTestExecutionBadRequest() *CancelTestExecutionBadRequest {
	return &CancelTestExecutionBadRequest{}
}

/*
CancelTestExecutionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CancelTestExecutionBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cancel test execution bad request response has a 2xx status code
func (o *CancelTestExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel test execution bad request response has a 3xx status code
func (o *CancelTestExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel test execution bad request response has a 4xx status code
func (o *CancelTestExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel test execution bad request response has a 5xx status code
func (o *CancelTestExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel test execution bad request response a status code equal to that given
func (o *CancelTestExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cancel test execution bad request response
func (o *CancelTestExecutionBadRequest) Code() int {
	return 400
}

func (o *CancelTestExecutionBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel][%d] cancelTestExecutionBadRequest %s", 400, payload)
}

func (o *CancelTestExecutionBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel][%d] cancelTestExecutionBadRequest %s", 400, payload)
}

func (o *CancelTestExecutionBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CancelTestExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelTestExecutionUnauthorized creates a CancelTestExecutionUnauthorized with default headers values
func NewCancelTestExecutionUnauthorized() *CancelTestExecutionUnauthorized {
	return &CancelTestExecutionUnauthorized{}
}

/*
CancelTestExecutionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CancelTestExecutionUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cancel test execution unauthorized response has a 2xx status code
func (o *CancelTestExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel test execution unauthorized response has a 3xx status code
func (o *CancelTestExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel test execution unauthorized response has a 4xx status code
func (o *CancelTestExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel test execution unauthorized response has a 5xx status code
func (o *CancelTestExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel test execution unauthorized response a status code equal to that given
func (o *CancelTestExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cancel test execution unauthorized response
func (o *CancelTestExecutionUnauthorized) Code() int {
	return 401
}

func (o *CancelTestExecutionUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel][%d] cancelTestExecutionUnauthorized %s", 401, payload)
}

func (o *CancelTestExecutionUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel][%d] cancelTestExecutionUnauthorized %s", 401, payload)
}

func (o *CancelTestExecutionUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CancelTestExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelTestExecutionBadGateway creates a CancelTestExecutionBadGateway with default headers values
func NewCancelTestExecutionBadGateway() *CancelTestExecutionBadGateway {
	return &CancelTestExecutionBadGateway{}
}

/*
CancelTestExecutionBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type CancelTestExecutionBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cancel test execution bad gateway response has a 2xx status code
func (o *CancelTestExecutionBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel test execution bad gateway response has a 3xx status code
func (o *CancelTestExecutionBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel test execution bad gateway response has a 4xx status code
func (o *CancelTestExecutionBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel test execution bad gateway response has a 5xx status code
func (o *CancelTestExecutionBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel test execution bad gateway response a status code equal to that given
func (o *CancelTestExecutionBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the cancel test execution bad gateway response
func (o *CancelTestExecutionBadGateway) Code() int {
	return 502
}

func (o *CancelTestExecutionBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel][%d] cancelTestExecutionBadGateway %s", 502, payload)
}

func (o *CancelTestExecutionBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /orgs/{orgName}/tests/executions/{executionID}/cancel][%d] cancelTestExecutionBadGateway %s", 502, payload)
}

func (o *CancelTestExecutionBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CancelTestExecutionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
