// Code generated by go-swagger; DO NOT EDIT.

package jobs

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

// ListJobsReader is a Reader for the ListJobs structure.
type ListJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewListJobsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{orgName}/jobs] list-jobs", response, response.Code())
	}
}

// NewListJobsOK creates a ListJobsOK with default headers values
func NewListJobsOK() *ListJobsOK {
	return &ListJobsOK{}
}

/*
ListJobsOK describes a response with status code 200, with default header values.

OK
*/
type ListJobsOK struct {
	Payload []*models.JobsJob
}

// IsSuccess returns true when this list jobs o k response has a 2xx status code
func (o *ListJobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list jobs o k response has a 3xx status code
func (o *ListJobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list jobs o k response has a 4xx status code
func (o *ListJobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list jobs o k response has a 5xx status code
func (o *ListJobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list jobs o k response a status code equal to that given
func (o *ListJobsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list jobs o k response
func (o *ListJobsOK) Code() int {
	return 200
}

func (o *ListJobsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/jobs][%d] listJobsOK %s", 200, payload)
}

func (o *ListJobsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/jobs][%d] listJobsOK %s", 200, payload)
}

func (o *ListJobsOK) GetPayload() []*models.JobsJob {
	return o.Payload
}

func (o *ListJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobsBadRequest creates a ListJobsBadRequest with default headers values
func NewListJobsBadRequest() *ListJobsBadRequest {
	return &ListJobsBadRequest{}
}

/*
ListJobsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListJobsBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list jobs bad request response has a 2xx status code
func (o *ListJobsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list jobs bad request response has a 3xx status code
func (o *ListJobsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list jobs bad request response has a 4xx status code
func (o *ListJobsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list jobs bad request response has a 5xx status code
func (o *ListJobsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list jobs bad request response a status code equal to that given
func (o *ListJobsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list jobs bad request response
func (o *ListJobsBadRequest) Code() int {
	return 400
}

func (o *ListJobsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/jobs][%d] listJobsBadRequest %s", 400, payload)
}

func (o *ListJobsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/jobs][%d] listJobsBadRequest %s", 400, payload)
}

func (o *ListJobsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobsUnauthorized creates a ListJobsUnauthorized with default headers values
func NewListJobsUnauthorized() *ListJobsUnauthorized {
	return &ListJobsUnauthorized{}
}

/*
ListJobsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListJobsUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list jobs unauthorized response has a 2xx status code
func (o *ListJobsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list jobs unauthorized response has a 3xx status code
func (o *ListJobsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list jobs unauthorized response has a 4xx status code
func (o *ListJobsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list jobs unauthorized response has a 5xx status code
func (o *ListJobsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list jobs unauthorized response a status code equal to that given
func (o *ListJobsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list jobs unauthorized response
func (o *ListJobsUnauthorized) Code() int {
	return 401
}

func (o *ListJobsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/jobs][%d] listJobsUnauthorized %s", 401, payload)
}

func (o *ListJobsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/jobs][%d] listJobsUnauthorized %s", 401, payload)
}

func (o *ListJobsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobsBadGateway creates a ListJobsBadGateway with default headers values
func NewListJobsBadGateway() *ListJobsBadGateway {
	return &ListJobsBadGateway{}
}

/*
ListJobsBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type ListJobsBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list jobs bad gateway response has a 2xx status code
func (o *ListJobsBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list jobs bad gateway response has a 3xx status code
func (o *ListJobsBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list jobs bad gateway response has a 4xx status code
func (o *ListJobsBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this list jobs bad gateway response has a 5xx status code
func (o *ListJobsBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this list jobs bad gateway response a status code equal to that given
func (o *ListJobsBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the list jobs bad gateway response
func (o *ListJobsBadGateway) Code() int {
	return 502
}

func (o *ListJobsBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/jobs][%d] listJobsBadGateway %s", 502, payload)
}

func (o *ListJobsBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/jobs][%d] listJobsBadGateway %s", 502, payload)
}

func (o *ListJobsBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListJobsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
