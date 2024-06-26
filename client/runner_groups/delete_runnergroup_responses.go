// Code generated by go-swagger; DO NOT EDIT.

package runner_groups

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

// DeleteRunnergroupReader is a Reader for the DeleteRunnergroup structure.
type DeleteRunnergroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunnergroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRunnergroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRunnergroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRunnergroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewDeleteRunnergroupBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}] delete-runnergroup", response, response.Code())
	}
}

// NewDeleteRunnergroupOK creates a DeleteRunnergroupOK with default headers values
func NewDeleteRunnergroupOK() *DeleteRunnergroupOK {
	return &DeleteRunnergroupOK{}
}

/*
DeleteRunnergroupOK describes a response with status code 200, with default header values.

OK
*/
type DeleteRunnergroupOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this delete runnergroup o k response has a 2xx status code
func (o *DeleteRunnergroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete runnergroup o k response has a 3xx status code
func (o *DeleteRunnergroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runnergroup o k response has a 4xx status code
func (o *DeleteRunnergroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete runnergroup o k response has a 5xx status code
func (o *DeleteRunnergroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runnergroup o k response a status code equal to that given
func (o *DeleteRunnergroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete runnergroup o k response
func (o *DeleteRunnergroupOK) Code() int {
	return 200
}

func (o *DeleteRunnergroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}][%d] deleteRunnergroupOK %s", 200, payload)
}

func (o *DeleteRunnergroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}][%d] deleteRunnergroupOK %s", 200, payload)
}

func (o *DeleteRunnergroupOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteRunnergroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunnergroupBadRequest creates a DeleteRunnergroupBadRequest with default headers values
func NewDeleteRunnergroupBadRequest() *DeleteRunnergroupBadRequest {
	return &DeleteRunnergroupBadRequest{}
}

/*
DeleteRunnergroupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteRunnergroupBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete runnergroup bad request response has a 2xx status code
func (o *DeleteRunnergroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete runnergroup bad request response has a 3xx status code
func (o *DeleteRunnergroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runnergroup bad request response has a 4xx status code
func (o *DeleteRunnergroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete runnergroup bad request response has a 5xx status code
func (o *DeleteRunnergroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runnergroup bad request response a status code equal to that given
func (o *DeleteRunnergroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete runnergroup bad request response
func (o *DeleteRunnergroupBadRequest) Code() int {
	return 400
}

func (o *DeleteRunnergroupBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}][%d] deleteRunnergroupBadRequest %s", 400, payload)
}

func (o *DeleteRunnergroupBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}][%d] deleteRunnergroupBadRequest %s", 400, payload)
}

func (o *DeleteRunnergroupBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRunnergroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunnergroupUnauthorized creates a DeleteRunnergroupUnauthorized with default headers values
func NewDeleteRunnergroupUnauthorized() *DeleteRunnergroupUnauthorized {
	return &DeleteRunnergroupUnauthorized{}
}

/*
DeleteRunnergroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRunnergroupUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete runnergroup unauthorized response has a 2xx status code
func (o *DeleteRunnergroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete runnergroup unauthorized response has a 3xx status code
func (o *DeleteRunnergroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runnergroup unauthorized response has a 4xx status code
func (o *DeleteRunnergroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete runnergroup unauthorized response has a 5xx status code
func (o *DeleteRunnergroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runnergroup unauthorized response a status code equal to that given
func (o *DeleteRunnergroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete runnergroup unauthorized response
func (o *DeleteRunnergroupUnauthorized) Code() int {
	return 401
}

func (o *DeleteRunnergroupUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}][%d] deleteRunnergroupUnauthorized %s", 401, payload)
}

func (o *DeleteRunnergroupUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}][%d] deleteRunnergroupUnauthorized %s", 401, payload)
}

func (o *DeleteRunnergroupUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRunnergroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunnergroupBadGateway creates a DeleteRunnergroupBadGateway with default headers values
func NewDeleteRunnergroupBadGateway() *DeleteRunnergroupBadGateway {
	return &DeleteRunnergroupBadGateway{}
}

/*
DeleteRunnergroupBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type DeleteRunnergroupBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete runnergroup bad gateway response has a 2xx status code
func (o *DeleteRunnergroupBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete runnergroup bad gateway response has a 3xx status code
func (o *DeleteRunnergroupBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runnergroup bad gateway response has a 4xx status code
func (o *DeleteRunnergroupBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete runnergroup bad gateway response has a 5xx status code
func (o *DeleteRunnergroupBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this delete runnergroup bad gateway response a status code equal to that given
func (o *DeleteRunnergroupBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the delete runnergroup bad gateway response
func (o *DeleteRunnergroupBadGateway) Code() int {
	return 502
}

func (o *DeleteRunnergroupBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}][%d] deleteRunnergroupBadGateway %s", 502, payload)
}

func (o *DeleteRunnergroupBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{orgName}/runnergroups/{runnergroupName}][%d] deleteRunnergroupBadGateway %s", 502, payload)
}

func (o *DeleteRunnergroupBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRunnergroupBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
