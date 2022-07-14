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

// GetClusterReader is a Reader for the GetCluster structure.
type GetClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetClusterBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterOK creates a GetClusterOK with default headers values
func NewGetClusterOK() *GetClusterOK {
	return &GetClusterOK{}
}

/* GetClusterOK describes a response with status code 200, with default header values.

OK
*/
type GetClusterOK struct {
	Payload *models.ClusterRegistration
}

func (o *GetClusterOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/{clusterName}/][%d] getClusterOK  %+v", 200, o.Payload)
}
func (o *GetClusterOK) GetPayload() *models.ClusterRegistration {
	return o.Payload
}

func (o *GetClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterRegistration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterBadRequest creates a GetClusterBadRequest with default headers values
func NewGetClusterBadRequest() *GetClusterBadRequest {
	return &GetClusterBadRequest{}
}

/* GetClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetClusterBadRequest struct {
	Payload *models.ApierrsResponse
}

func (o *GetClusterBadRequest) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/{clusterName}/][%d] getClusterBadRequest  %+v", 400, o.Payload)
}
func (o *GetClusterBadRequest) GetPayload() *models.ApierrsResponse {
	return o.Payload
}

func (o *GetClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApierrsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterUnauthorized creates a GetClusterUnauthorized with default headers values
func NewGetClusterUnauthorized() *GetClusterUnauthorized {
	return &GetClusterUnauthorized{}
}

/* GetClusterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetClusterUnauthorized struct {
	Payload *models.ApierrsResponse
}

func (o *GetClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/{clusterName}/][%d] getClusterUnauthorized  %+v", 401, o.Payload)
}
func (o *GetClusterUnauthorized) GetPayload() *models.ApierrsResponse {
	return o.Payload
}

func (o *GetClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApierrsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterBadGateway creates a GetClusterBadGateway with default headers values
func NewGetClusterBadGateway() *GetClusterBadGateway {
	return &GetClusterBadGateway{}
}

/* GetClusterBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type GetClusterBadGateway struct {
	Payload *models.ApierrsResponse
}

func (o *GetClusterBadGateway) Error() string {
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/{clusterName}/][%d] getClusterBadGateway  %+v", 502, o.Payload)
}
func (o *GetClusterBadGateway) GetPayload() *models.ApierrsResponse {
	return o.Payload
}

func (o *GetClusterBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApierrsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
