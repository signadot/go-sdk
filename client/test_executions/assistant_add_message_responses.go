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

// AssistantAddMessageReader is a Reader for the AssistantAddMessage structure.
type AssistantAddMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssistantAddMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssistantAddMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssistantAddMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAssistantAddMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /orgs/{orgName}/assistants/threads/{threadID}/messages] assistant-add-message", response, response.Code())
	}
}

// NewAssistantAddMessageOK creates a AssistantAddMessageOK with default headers values
func NewAssistantAddMessageOK() *AssistantAddMessageOK {
	return &AssistantAddMessageOK{}
}

/*
AssistantAddMessageOK describes a response with status code 200, with default header values.

OK
*/
type AssistantAddMessageOK struct {
	Payload *models.AssistantsMessage
}

// IsSuccess returns true when this assistant add message o k response has a 2xx status code
func (o *AssistantAddMessageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assistant add message o k response has a 3xx status code
func (o *AssistantAddMessageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assistant add message o k response has a 4xx status code
func (o *AssistantAddMessageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assistant add message o k response has a 5xx status code
func (o *AssistantAddMessageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assistant add message o k response a status code equal to that given
func (o *AssistantAddMessageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assistant add message o k response
func (o *AssistantAddMessageOK) Code() int {
	return 200
}

func (o *AssistantAddMessageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/assistants/threads/{threadID}/messages][%d] assistantAddMessageOK %s", 200, payload)
}

func (o *AssistantAddMessageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/assistants/threads/{threadID}/messages][%d] assistantAddMessageOK %s", 200, payload)
}

func (o *AssistantAddMessageOK) GetPayload() *models.AssistantsMessage {
	return o.Payload
}

func (o *AssistantAddMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssistantsMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssistantAddMessageBadRequest creates a AssistantAddMessageBadRequest with default headers values
func NewAssistantAddMessageBadRequest() *AssistantAddMessageBadRequest {
	return &AssistantAddMessageBadRequest{}
}

/*
AssistantAddMessageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AssistantAddMessageBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assistant add message bad request response has a 2xx status code
func (o *AssistantAddMessageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assistant add message bad request response has a 3xx status code
func (o *AssistantAddMessageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assistant add message bad request response has a 4xx status code
func (o *AssistantAddMessageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this assistant add message bad request response has a 5xx status code
func (o *AssistantAddMessageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this assistant add message bad request response a status code equal to that given
func (o *AssistantAddMessageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the assistant add message bad request response
func (o *AssistantAddMessageBadRequest) Code() int {
	return 400
}

func (o *AssistantAddMessageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/assistants/threads/{threadID}/messages][%d] assistantAddMessageBadRequest %s", 400, payload)
}

func (o *AssistantAddMessageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/assistants/threads/{threadID}/messages][%d] assistantAddMessageBadRequest %s", 400, payload)
}

func (o *AssistantAddMessageBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssistantAddMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssistantAddMessageUnauthorized creates a AssistantAddMessageUnauthorized with default headers values
func NewAssistantAddMessageUnauthorized() *AssistantAddMessageUnauthorized {
	return &AssistantAddMessageUnauthorized{}
}

/*
AssistantAddMessageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AssistantAddMessageUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assistant add message unauthorized response has a 2xx status code
func (o *AssistantAddMessageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assistant add message unauthorized response has a 3xx status code
func (o *AssistantAddMessageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assistant add message unauthorized response has a 4xx status code
func (o *AssistantAddMessageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this assistant add message unauthorized response has a 5xx status code
func (o *AssistantAddMessageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this assistant add message unauthorized response a status code equal to that given
func (o *AssistantAddMessageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the assistant add message unauthorized response
func (o *AssistantAddMessageUnauthorized) Code() int {
	return 401
}

func (o *AssistantAddMessageUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/assistants/threads/{threadID}/messages][%d] assistantAddMessageUnauthorized %s", 401, payload)
}

func (o *AssistantAddMessageUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/assistants/threads/{threadID}/messages][%d] assistantAddMessageUnauthorized %s", 401, payload)
}

func (o *AssistantAddMessageUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssistantAddMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
