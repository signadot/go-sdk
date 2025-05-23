// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

// GetOrgNameReader is a Reader for the GetOrgName structure.
type GetOrgNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/] get-org-name", response, response.Code())
	}
}

// NewGetOrgNameOK creates a GetOrgNameOK with default headers values
func NewGetOrgNameOK() *GetOrgNameOK {
	return &GetOrgNameOK{}
}

/*
GetOrgNameOK describes a response with status code 200, with default header values.

OK
*/
type GetOrgNameOK struct {
	Payload *models.OrgsGetOrgNameResponse
}

// IsSuccess returns true when this get org name o k response has a 2xx status code
func (o *GetOrgNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org name o k response has a 3xx status code
func (o *GetOrgNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org name o k response has a 4xx status code
func (o *GetOrgNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org name o k response has a 5xx status code
func (o *GetOrgNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org name o k response a status code equal to that given
func (o *GetOrgNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org name o k response
func (o *GetOrgNameOK) Code() int {
	return 200
}

func (o *GetOrgNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/][%d] getOrgNameOK %s", 200, payload)
}

func (o *GetOrgNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/][%d] getOrgNameOK %s", 200, payload)
}

func (o *GetOrgNameOK) GetPayload() *models.OrgsGetOrgNameResponse {
	return o.Payload
}

func (o *GetOrgNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrgsGetOrgNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgNameBadRequest creates a GetOrgNameBadRequest with default headers values
func NewGetOrgNameBadRequest() *GetOrgNameBadRequest {
	return &GetOrgNameBadRequest{}
}

/*
GetOrgNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrgNameBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get org name bad request response has a 2xx status code
func (o *GetOrgNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org name bad request response has a 3xx status code
func (o *GetOrgNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org name bad request response has a 4xx status code
func (o *GetOrgNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org name bad request response has a 5xx status code
func (o *GetOrgNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get org name bad request response a status code equal to that given
func (o *GetOrgNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get org name bad request response
func (o *GetOrgNameBadRequest) Code() int {
	return 400
}

func (o *GetOrgNameBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/][%d] getOrgNameBadRequest %s", 400, payload)
}

func (o *GetOrgNameBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/][%d] getOrgNameBadRequest %s", 400, payload)
}

func (o *GetOrgNameBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOrgNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgNameForbidden creates a GetOrgNameForbidden with default headers values
func NewGetOrgNameForbidden() *GetOrgNameForbidden {
	return &GetOrgNameForbidden{}
}

/*
GetOrgNameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrgNameForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get org name forbidden response has a 2xx status code
func (o *GetOrgNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org name forbidden response has a 3xx status code
func (o *GetOrgNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org name forbidden response has a 4xx status code
func (o *GetOrgNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org name forbidden response has a 5xx status code
func (o *GetOrgNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org name forbidden response a status code equal to that given
func (o *GetOrgNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org name forbidden response
func (o *GetOrgNameForbidden) Code() int {
	return 403
}

func (o *GetOrgNameForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/][%d] getOrgNameForbidden %s", 403, payload)
}

func (o *GetOrgNameForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/][%d] getOrgNameForbidden %s", 403, payload)
}

func (o *GetOrgNameForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOrgNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgNameInternalServerError creates a GetOrgNameInternalServerError with default headers values
func NewGetOrgNameInternalServerError() *GetOrgNameInternalServerError {
	return &GetOrgNameInternalServerError{}
}

/*
GetOrgNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetOrgNameInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get org name internal server error response has a 2xx status code
func (o *GetOrgNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org name internal server error response has a 3xx status code
func (o *GetOrgNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org name internal server error response has a 4xx status code
func (o *GetOrgNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org name internal server error response has a 5xx status code
func (o *GetOrgNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org name internal server error response a status code equal to that given
func (o *GetOrgNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org name internal server error response
func (o *GetOrgNameInternalServerError) Code() int {
	return 500
}

func (o *GetOrgNameInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/][%d] getOrgNameInternalServerError %s", 500, payload)
}

func (o *GetOrgNameInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/][%d] getOrgNameInternalServerError %s", 500, payload)
}

func (o *GetOrgNameInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOrgNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
