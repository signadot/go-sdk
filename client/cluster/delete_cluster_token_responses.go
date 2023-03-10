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
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteClusterTokenOK creates a DeleteClusterTokenOK with default headers values
func NewDeleteClusterTokenOK() *DeleteClusterTokenOK {
	return &DeleteClusterTokenOK{}
}

/*
DeleteClusterTokenOK describes a response with status code 200, with default header values.

OK
*/
type DeleteClusterTokenOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this delete cluster token o k response has a 2xx status code
func (o *DeleteClusterTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cluster token o k response has a 3xx status code
func (o *DeleteClusterTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster token o k response has a 4xx status code
func (o *DeleteClusterTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cluster token o k response has a 5xx status code
func (o *DeleteClusterTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster token o k response a status code equal to that given
func (o *DeleteClusterTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete cluster token o k response
func (o *DeleteClusterTokenOK) Code() int {
	return 200
}

func (o *DeleteClusterTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/tokens/{tokenId}][%d] deleteClusterTokenOK  %+v", 200, o.Payload)
}

func (o *DeleteClusterTokenOK) String() string {
	return fmt.Sprintf("[DELETE /orgs/{orgName}/clusters/{clusterName}/tokens/{tokenId}][%d] deleteClusterTokenOK  %+v", 200, o.Payload)
}

func (o *DeleteClusterTokenOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteClusterTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
