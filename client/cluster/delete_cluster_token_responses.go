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

// DeleteClusterTokenReader is a Reader for the DeleteClusterToken structure.
type DeleteClusterTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteClusterTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteClusterTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteClusterTokenOK creates a DeleteClusterTokenOK with default headers values
func NewDeleteClusterTokenOK() *DeleteClusterTokenOK {
	return &DeleteClusterTokenOK{}
}

/* DeleteClusterTokenOK describes a response with status code 200, with default header values.

OK
*/
type DeleteClusterTokenOK struct {
	Payload models.HandlerEmptyResponse
}

func (o *DeleteClusterTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/tokens/{tokenId}][%d] deleteClusterTokenOK  %+v", 200, o.Payload)
}
func (o *DeleteClusterTokenOK) GetPayload() models.HandlerEmptyResponse {
	return o.Payload
}

func (o *DeleteClusterTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterTokenUnauthorized creates a DeleteClusterTokenUnauthorized with default headers values
func NewDeleteClusterTokenUnauthorized() *DeleteClusterTokenUnauthorized {
	return &DeleteClusterTokenUnauthorized{}
}

/* DeleteClusterTokenUnauthorized describes a response with status code 401, with default header values.

Authorization failure
*/
type DeleteClusterTokenUnauthorized struct {
}

func (o *DeleteClusterTokenUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/tokens/{tokenId}][%d] deleteClusterTokenUnauthorized ", 401)
}

func (o *DeleteClusterTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterTokenInternalServerError creates a DeleteClusterTokenInternalServerError with default headers values
func NewDeleteClusterTokenInternalServerError() *DeleteClusterTokenInternalServerError {
	return &DeleteClusterTokenInternalServerError{}
}

/* DeleteClusterTokenInternalServerError describes a response with status code 500, with default header values.

Internal server failure.
*/
type DeleteClusterTokenInternalServerError struct {
}

func (o *DeleteClusterTokenInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/tokens/{tokenId}][%d] deleteClusterTokenInternalServerError ", 500)
}

func (o *DeleteClusterTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
