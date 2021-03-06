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

// AddClusterReader is a Reader for the AddCluster structure.
type AddClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewAddClusterBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddClusterOK creates a AddClusterOK with default headers values
func NewAddClusterOK() *AddClusterOK {
	return &AddClusterOK{}
}

/* AddClusterOK describes a response with status code 200, with default header values.

OK
*/
type AddClusterOK struct {
	Payload *models.Cluster
}

func (o *AddClusterOK) Error() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/clusters/{clusterName}/][%d] addClusterOK  %+v", 200, o.Payload)
}
func (o *AddClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *AddClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterBadRequest creates a AddClusterBadRequest with default headers values
func NewAddClusterBadRequest() *AddClusterBadRequest {
	return &AddClusterBadRequest{}
}

/* AddClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddClusterBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *AddClusterBadRequest) Error() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/clusters/{clusterName}/][%d] addClusterBadRequest  %+v", 400, o.Payload)
}
func (o *AddClusterBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterUnauthorized creates a AddClusterUnauthorized with default headers values
func NewAddClusterUnauthorized() *AddClusterUnauthorized {
	return &AddClusterUnauthorized{}
}

/* AddClusterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddClusterUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *AddClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/clusters/{clusterName}/][%d] addClusterUnauthorized  %+v", 401, o.Payload)
}
func (o *AddClusterUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterBadGateway creates a AddClusterBadGateway with default headers values
func NewAddClusterBadGateway() *AddClusterBadGateway {
	return &AddClusterBadGateway{}
}

/* AddClusterBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type AddClusterBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *AddClusterBadGateway) Error() string {
	return fmt.Sprintf("[PUT /orgs/{orgName}/clusters/{clusterName}/][%d] addClusterBadGateway  %+v", 502, o.Payload)
}
func (o *AddClusterBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddClusterBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
