// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// CreateClusterTokenReader is a Reader for the CreateClusterToken structure.
type CreateClusterTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateClusterTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /orgs/{orgName}/clusters/{clusterName}/tokens] create-cluster-token", response, response.Code())
	}
}

// NewCreateClusterTokenOK creates a CreateClusterTokenOK with default headers values
func NewCreateClusterTokenOK() *CreateClusterTokenOK {
	return &CreateClusterTokenOK{}
}

/*
CreateClusterTokenOK describes a response with status code 200, with default header values.

OK
*/
type CreateClusterTokenOK struct {
	Payload *models.ClusterToken
}

// IsSuccess returns true when this create cluster token o k response has a 2xx status code
func (o *CreateClusterTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cluster token o k response has a 3xx status code
func (o *CreateClusterTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster token o k response has a 4xx status code
func (o *CreateClusterTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster token o k response has a 5xx status code
func (o *CreateClusterTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster token o k response a status code equal to that given
func (o *CreateClusterTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create cluster token o k response
func (o *CreateClusterTokenOK) Code() int {
	return 200
}

func (o *CreateClusterTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/clusters/{clusterName}/tokens][%d] createClusterTokenOK %s", 200, payload)
}

func (o *CreateClusterTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /orgs/{orgName}/clusters/{clusterName}/tokens][%d] createClusterTokenOK %s", 200, payload)
}

func (o *CreateClusterTokenOK) GetPayload() *models.ClusterToken {
	return o.Payload
}

func (o *CreateClusterTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
