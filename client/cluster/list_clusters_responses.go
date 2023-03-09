// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/signadot/go-sdk/models"
)

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewListClustersBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/*
ListClustersOK describes a response with status code 200, with default header values.

OK
*/
type ListClustersOK struct {
	Payload []*models.Cluster
}

// IsSuccess returns true when this list clusters o k response has a 2xx status code
func (o *ListClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list clusters o k response has a 3xx status code
func (o *ListClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list clusters o k response has a 4xx status code
func (o *ListClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list clusters o k response has a 5xx status code
func (o *ListClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list clusters o k response a status code equal to that given
func (o *ListClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list clusters o k response
func (o *ListClustersOK) Code() int {
	return 200
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) GetPayload() []*models.Cluster {
	return o.Payload
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersBadRequest creates a ListClustersBadRequest with default headers values
func NewListClustersBadRequest() *ListClustersBadRequest {
	return &ListClustersBadRequest{}
}

/*
ListClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListClustersBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list clusters bad request response has a 2xx status code
func (o *ListClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list clusters bad request response has a 3xx status code
func (o *ListClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list clusters bad request response has a 4xx status code
func (o *ListClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list clusters bad request response has a 5xx status code
func (o *ListClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list clusters bad request response a status code equal to that given
func (o *ListClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list clusters bad request response
func (o *ListClustersBadRequest) Code() int {
	return 400
}

func (o *ListClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/][%d] listClustersBadRequest  %+v", 400, o.Payload)
}

func (o *ListClustersBadRequest) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/][%d] listClustersBadRequest  %+v", 400, o.Payload)
}

func (o *ListClustersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersUnauthorized creates a ListClustersUnauthorized with default headers values
func NewListClustersUnauthorized() *ListClustersUnauthorized {
	return &ListClustersUnauthorized{}
}

/*
ListClustersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListClustersUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list clusters unauthorized response has a 2xx status code
func (o *ListClustersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list clusters unauthorized response has a 3xx status code
func (o *ListClustersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list clusters unauthorized response has a 4xx status code
func (o *ListClustersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list clusters unauthorized response has a 5xx status code
func (o *ListClustersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list clusters unauthorized response a status code equal to that given
func (o *ListClustersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list clusters unauthorized response
func (o *ListClustersUnauthorized) Code() int {
	return 401
}

func (o *ListClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/][%d] listClustersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListClustersUnauthorized) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/][%d] listClustersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListClustersUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersBadGateway creates a ListClustersBadGateway with default headers values
func NewListClustersBadGateway() *ListClustersBadGateway {
	return &ListClustersBadGateway{}
}

/*
ListClustersBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type ListClustersBadGateway struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list clusters bad gateway response has a 2xx status code
func (o *ListClustersBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list clusters bad gateway response has a 3xx status code
func (o *ListClustersBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list clusters bad gateway response has a 4xx status code
func (o *ListClustersBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this list clusters bad gateway response has a 5xx status code
func (o *ListClustersBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this list clusters bad gateway response a status code equal to that given
func (o *ListClustersBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the list clusters bad gateway response
func (o *ListClustersBadGateway) Code() int {
	return 502
}

func (o *ListClustersBadGateway) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/][%d] listClustersBadGateway  %+v", 502, o.Payload)
}

func (o *ListClustersBadGateway) String() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/][%d] listClustersBadGateway  %+v", 502, o.Payload)
}

func (o *ListClustersBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClustersBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
