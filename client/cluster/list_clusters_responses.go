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

/* ListClustersOK describes a response with status code 200, with default header values.

OK
*/
type ListClustersOK struct {
	Payload []*models.Cluster
}

func (o *ListClustersOK) Error() string {
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

/* ListClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListClustersBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ListClustersBadRequest) Error() string {
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

/* ListClustersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListClustersUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *ListClustersUnauthorized) Error() string {
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

/* ListClustersBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type ListClustersBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *ListClustersBadGateway) Error() string {
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
