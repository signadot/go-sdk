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

// CreateJobReader is a Reader for the CreateJob structure.
type CreateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewCreateJobBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /orgs/{orgName}/jobs] create-job", response, response.Code())
	}
}

// NewCreateJobOK creates a CreateJobOK with default headers values
func NewCreateJobOK() *CreateJobOK {
	return &CreateJobOK{}
}

/*
CreateJobOK describes a response with status code 200, with default header values.

OK
*/
type CreateJobOK struct {
	Payload *models.Job
}

// IsSuccess returns true when this create job o k response has a 2xx status code
func (o *CreateJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create job o k response has a 3xx status code
func (o *CreateJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job o k response has a 4xx status code
func (o *CreateJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create job o k response has a 5xx status code
func (o *CreateJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create job o k response a status code equal to that given
func (o *CreateJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create job o k response
func (o *CreateJobOK) Code() int {
	return 200
}

func (o *CreateJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/jobs][%d] createJobOK %s", 200, payload)
}

func (o *CreateJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/jobs][%d] createJobOK %s", 200, payload)
}

func (o *CreateJobOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *CreateJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobBadRequest creates a CreateJobBadRequest with default headers values
func NewCreateJobBadRequest() *CreateJobBadRequest {
	return &CreateJobBadRequest{}
}

/*
CreateJobBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateJobBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create job bad request response has a 2xx status code
func (o *CreateJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create job bad request response has a 3xx status code
func (o *CreateJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job bad request response has a 4xx status code
func (o *CreateJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create job bad request response has a 5xx status code
func (o *CreateJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create job bad request response a status code equal to that given
func (o *CreateJobBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create job bad request response
func (o *CreateJobBadRequest) Code() int {
	return 400
}

func (o *CreateJobBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/jobs][%d] createJobBadRequest %s", 400, payload)
}

func (o *CreateJobBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/jobs][%d] createJobBadRequest %s", 400, payload)
}

func (o *CreateJobBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobUnauthorized creates a CreateJobUnauthorized with default headers values
func NewCreateJobUnauthorized() *CreateJobUnauthorized {
	return &CreateJobUnauthorized{}
}

/*
CreateJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateJobUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create job unauthorized response has a 2xx status code
func (o *CreateJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create job unauthorized response has a 3xx status code
func (o *CreateJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job unauthorized response has a 4xx status code
func (o *CreateJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create job unauthorized response has a 5xx status code
func (o *CreateJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create job unauthorized response a status code equal to that given
func (o *CreateJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create job unauthorized response
func (o *CreateJobUnauthorized) Code() int {
	return 401
}

func (o *CreateJobUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/jobs][%d] createJobUnauthorized %s", 401, payload)
}

func (o *CreateJobUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/jobs][%d] createJobUnauthorized %s", 401, payload)
}

func (o *CreateJobUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobBadGateway creates a CreateJobBadGateway with default headers values
func NewCreateJobBadGateway() *CreateJobBadGateway {
	return &CreateJobBadGateway{}
}

/*
CreateJobBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type CreateJobBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create job bad gateway response has a 2xx status code
func (o *CreateJobBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create job bad gateway response has a 3xx status code
func (o *CreateJobBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job bad gateway response has a 4xx status code
func (o *CreateJobBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this create job bad gateway response has a 5xx status code
func (o *CreateJobBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this create job bad gateway response a status code equal to that given
func (o *CreateJobBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the create job bad gateway response
func (o *CreateJobBadGateway) Code() int {
	return 502
}

func (o *CreateJobBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/jobs][%d] createJobBadGateway %s", 502, payload)
}

func (o *CreateJobBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/jobs][%d] createJobBadGateway %s", 502, payload)
}

func (o *CreateJobBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateJobBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
