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

// RemoveClusterReader is a Reader for the RemoveCluster structure.
type RemoveClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewRemoveClusterBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveClusterOK creates a RemoveClusterOK with default headers values
func NewRemoveClusterOK() *RemoveClusterOK {
	return &RemoveClusterOK{}
}

/* RemoveClusterOK describes a response with status code 200, with default header values.

OK
*/
type RemoveClusterOK struct {
	Payload models.EmptyResponse
}

func (o *RemoveClusterOK) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/][%d] removeClusterOK  %+v", 200, o.Payload)
}
func (o *RemoveClusterOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *RemoveClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveClusterBadRequest creates a RemoveClusterBadRequest with default headers values
func NewRemoveClusterBadRequest() *RemoveClusterBadRequest {
	return &RemoveClusterBadRequest{}
}

/* RemoveClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RemoveClusterBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *RemoveClusterBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/][%d] removeClusterBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveClusterBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RemoveClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveClusterUnauthorized creates a RemoveClusterUnauthorized with default headers values
func NewRemoveClusterUnauthorized() *RemoveClusterUnauthorized {
	return &RemoveClusterUnauthorized{}
}

/* RemoveClusterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RemoveClusterUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *RemoveClusterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/][%d] removeClusterUnauthorized  %+v", 401, o.Payload)
}
func (o *RemoveClusterUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RemoveClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveClusterBadGateway creates a RemoveClusterBadGateway with default headers values
func NewRemoveClusterBadGateway() *RemoveClusterBadGateway {
	return &RemoveClusterBadGateway{}
}

/* RemoveClusterBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type RemoveClusterBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *RemoveClusterBadGateway) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/][%d] removeClusterBadGateway  %+v", 502, o.Payload)
}
func (o *RemoveClusterBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RemoveClusterBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
