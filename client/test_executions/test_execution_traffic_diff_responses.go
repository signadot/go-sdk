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

// TestExecutionTrafficDiffReader is a Reader for the TestExecutionTrafficDiff structure.
type TestExecutionTrafficDiffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestExecutionTrafficDiffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestExecutionTrafficDiffOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestExecutionTrafficDiffBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTestExecutionTrafficDiffUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewTestExecutionTrafficDiffBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff] test-execution-traffic-diff", response, response.Code())
	}
}

// NewTestExecutionTrafficDiffOK creates a TestExecutionTrafficDiffOK with default headers values
func NewTestExecutionTrafficDiffOK() *TestExecutionTrafficDiffOK {
	return &TestExecutionTrafficDiffOK{}
}

/*
TestExecutionTrafficDiffOK describes a response with status code 200, with default header values.

OK
*/
type TestExecutionTrafficDiffOK struct {
	Payload *models.TrafficDiffResult
}

// IsSuccess returns true when this test execution traffic diff o k response has a 2xx status code
func (o *TestExecutionTrafficDiffOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this test execution traffic diff o k response has a 3xx status code
func (o *TestExecutionTrafficDiffOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test execution traffic diff o k response has a 4xx status code
func (o *TestExecutionTrafficDiffOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this test execution traffic diff o k response has a 5xx status code
func (o *TestExecutionTrafficDiffOK) IsServerError() bool {
	return false
}

// IsCode returns true when this test execution traffic diff o k response a status code equal to that given
func (o *TestExecutionTrafficDiffOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the test execution traffic diff o k response
func (o *TestExecutionTrafficDiffOK) Code() int {
	return 200
}

func (o *TestExecutionTrafficDiffOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff][%d] testExecutionTrafficDiffOK %s", 200, payload)
}

func (o *TestExecutionTrafficDiffOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff][%d] testExecutionTrafficDiffOK %s", 200, payload)
}

func (o *TestExecutionTrafficDiffOK) GetPayload() *models.TrafficDiffResult {
	return o.Payload
}

func (o *TestExecutionTrafficDiffOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrafficDiffResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestExecutionTrafficDiffBadRequest creates a TestExecutionTrafficDiffBadRequest with default headers values
func NewTestExecutionTrafficDiffBadRequest() *TestExecutionTrafficDiffBadRequest {
	return &TestExecutionTrafficDiffBadRequest{}
}

/*
TestExecutionTrafficDiffBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TestExecutionTrafficDiffBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this test execution traffic diff bad request response has a 2xx status code
func (o *TestExecutionTrafficDiffBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test execution traffic diff bad request response has a 3xx status code
func (o *TestExecutionTrafficDiffBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test execution traffic diff bad request response has a 4xx status code
func (o *TestExecutionTrafficDiffBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this test execution traffic diff bad request response has a 5xx status code
func (o *TestExecutionTrafficDiffBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this test execution traffic diff bad request response a status code equal to that given
func (o *TestExecutionTrafficDiffBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the test execution traffic diff bad request response
func (o *TestExecutionTrafficDiffBadRequest) Code() int {
	return 400
}

func (o *TestExecutionTrafficDiffBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff][%d] testExecutionTrafficDiffBadRequest %s", 400, payload)
}

func (o *TestExecutionTrafficDiffBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff][%d] testExecutionTrafficDiffBadRequest %s", 400, payload)
}

func (o *TestExecutionTrafficDiffBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TestExecutionTrafficDiffBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestExecutionTrafficDiffUnauthorized creates a TestExecutionTrafficDiffUnauthorized with default headers values
func NewTestExecutionTrafficDiffUnauthorized() *TestExecutionTrafficDiffUnauthorized {
	return &TestExecutionTrafficDiffUnauthorized{}
}

/*
TestExecutionTrafficDiffUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TestExecutionTrafficDiffUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this test execution traffic diff unauthorized response has a 2xx status code
func (o *TestExecutionTrafficDiffUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test execution traffic diff unauthorized response has a 3xx status code
func (o *TestExecutionTrafficDiffUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test execution traffic diff unauthorized response has a 4xx status code
func (o *TestExecutionTrafficDiffUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this test execution traffic diff unauthorized response has a 5xx status code
func (o *TestExecutionTrafficDiffUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this test execution traffic diff unauthorized response a status code equal to that given
func (o *TestExecutionTrafficDiffUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the test execution traffic diff unauthorized response
func (o *TestExecutionTrafficDiffUnauthorized) Code() int {
	return 401
}

func (o *TestExecutionTrafficDiffUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff][%d] testExecutionTrafficDiffUnauthorized %s", 401, payload)
}

func (o *TestExecutionTrafficDiffUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff][%d] testExecutionTrafficDiffUnauthorized %s", 401, payload)
}

func (o *TestExecutionTrafficDiffUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TestExecutionTrafficDiffUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestExecutionTrafficDiffBadGateway creates a TestExecutionTrafficDiffBadGateway with default headers values
func NewTestExecutionTrafficDiffBadGateway() *TestExecutionTrafficDiffBadGateway {
	return &TestExecutionTrafficDiffBadGateway{}
}

/*
TestExecutionTrafficDiffBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type TestExecutionTrafficDiffBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this test execution traffic diff bad gateway response has a 2xx status code
func (o *TestExecutionTrafficDiffBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test execution traffic diff bad gateway response has a 3xx status code
func (o *TestExecutionTrafficDiffBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test execution traffic diff bad gateway response has a 4xx status code
func (o *TestExecutionTrafficDiffBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this test execution traffic diff bad gateway response has a 5xx status code
func (o *TestExecutionTrafficDiffBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this test execution traffic diff bad gateway response a status code equal to that given
func (o *TestExecutionTrafficDiffBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the test execution traffic diff bad gateway response
func (o *TestExecutionTrafficDiffBadGateway) Code() int {
	return 502
}

func (o *TestExecutionTrafficDiffBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff][%d] testExecutionTrafficDiffBadGateway %s", 502, payload)
}

func (o *TestExecutionTrafficDiffBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/tests/executions/{executionName}/traffic-diff][%d] testExecutionTrafficDiffBadGateway %s", 502, payload)
}

func (o *TestExecutionTrafficDiffBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TestExecutionTrafficDiffBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
