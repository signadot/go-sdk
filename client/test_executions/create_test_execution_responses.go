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

// CreateTestExecutionReader is a Reader for the CreateTestExecution structure.
type CreateTestExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTestExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTestExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTestExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTestExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewCreateTestExecutionBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /orgs/{orgName}/tests/executions/] create-test-execution", response, response.Code())
	}
}

// NewCreateTestExecutionOK creates a CreateTestExecutionOK with default headers values
func NewCreateTestExecutionOK() *CreateTestExecutionOK {
	return &CreateTestExecutionOK{}
}

/*
CreateTestExecutionOK describes a response with status code 200, with default header values.

OK
*/
type CreateTestExecutionOK struct {
	Payload *models.TestExecution
}

// IsSuccess returns true when this create test execution o k response has a 2xx status code
func (o *CreateTestExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create test execution o k response has a 3xx status code
func (o *CreateTestExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create test execution o k response has a 4xx status code
func (o *CreateTestExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create test execution o k response has a 5xx status code
func (o *CreateTestExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create test execution o k response a status code equal to that given
func (o *CreateTestExecutionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create test execution o k response
func (o *CreateTestExecutionOK) Code() int {
	return 200
}

func (o *CreateTestExecutionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/tests/executions/][%d] createTestExecutionOK %s", 200, payload)
}

func (o *CreateTestExecutionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/tests/executions/][%d] createTestExecutionOK %s", 200, payload)
}

func (o *CreateTestExecutionOK) GetPayload() *models.TestExecution {
	return o.Payload
}

func (o *CreateTestExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTestExecutionBadRequest creates a CreateTestExecutionBadRequest with default headers values
func NewCreateTestExecutionBadRequest() *CreateTestExecutionBadRequest {
	return &CreateTestExecutionBadRequest{}
}

/*
CreateTestExecutionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateTestExecutionBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create test execution bad request response has a 2xx status code
func (o *CreateTestExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create test execution bad request response has a 3xx status code
func (o *CreateTestExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create test execution bad request response has a 4xx status code
func (o *CreateTestExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create test execution bad request response has a 5xx status code
func (o *CreateTestExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create test execution bad request response a status code equal to that given
func (o *CreateTestExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create test execution bad request response
func (o *CreateTestExecutionBadRequest) Code() int {
	return 400
}

func (o *CreateTestExecutionBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/tests/executions/][%d] createTestExecutionBadRequest %s", 400, payload)
}

func (o *CreateTestExecutionBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/tests/executions/][%d] createTestExecutionBadRequest %s", 400, payload)
}

func (o *CreateTestExecutionBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateTestExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTestExecutionUnauthorized creates a CreateTestExecutionUnauthorized with default headers values
func NewCreateTestExecutionUnauthorized() *CreateTestExecutionUnauthorized {
	return &CreateTestExecutionUnauthorized{}
}

/*
CreateTestExecutionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateTestExecutionUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create test execution unauthorized response has a 2xx status code
func (o *CreateTestExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create test execution unauthorized response has a 3xx status code
func (o *CreateTestExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create test execution unauthorized response has a 4xx status code
func (o *CreateTestExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create test execution unauthorized response has a 5xx status code
func (o *CreateTestExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create test execution unauthorized response a status code equal to that given
func (o *CreateTestExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create test execution unauthorized response
func (o *CreateTestExecutionUnauthorized) Code() int {
	return 401
}

func (o *CreateTestExecutionUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/tests/executions/][%d] createTestExecutionUnauthorized %s", 401, payload)
}

func (o *CreateTestExecutionUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/tests/executions/][%d] createTestExecutionUnauthorized %s", 401, payload)
}

func (o *CreateTestExecutionUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateTestExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTestExecutionBadGateway creates a CreateTestExecutionBadGateway with default headers values
func NewCreateTestExecutionBadGateway() *CreateTestExecutionBadGateway {
	return &CreateTestExecutionBadGateway{}
}

/*
CreateTestExecutionBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type CreateTestExecutionBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create test execution bad gateway response has a 2xx status code
func (o *CreateTestExecutionBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create test execution bad gateway response has a 3xx status code
func (o *CreateTestExecutionBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create test execution bad gateway response has a 4xx status code
func (o *CreateTestExecutionBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this create test execution bad gateway response has a 5xx status code
func (o *CreateTestExecutionBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this create test execution bad gateway response a status code equal to that given
func (o *CreateTestExecutionBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the create test execution bad gateway response
func (o *CreateTestExecutionBadGateway) Code() int {
	return 502
}

func (o *CreateTestExecutionBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/tests/executions/][%d] createTestExecutionBadGateway %s", 502, payload)
}

func (o *CreateTestExecutionBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/tests/executions/][%d] createTestExecutionBadGateway %s", 502, payload)
}

func (o *CreateTestExecutionBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateTestExecutionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
