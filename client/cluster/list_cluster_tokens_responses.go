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

// ListClusterTokensReader is a Reader for the ListClusterTokens structure.
type ListClusterTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{orgName}/clusters/{clusterName}/tokens/] list-cluster-tokens", response, response.Code())
	}
}

// NewListClusterTokensOK creates a ListClusterTokensOK with default headers values
func NewListClusterTokensOK() *ListClusterTokensOK {
	return &ListClusterTokensOK{}
}

/*
ListClusterTokensOK describes a response with status code 200, with default header values.

OK
*/
type ListClusterTokensOK struct {
	Payload []*models.ClusterToken
}

// IsSuccess returns true when this list cluster tokens o k response has a 2xx status code
func (o *ListClusterTokensOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list cluster tokens o k response has a 3xx status code
func (o *ListClusterTokensOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster tokens o k response has a 4xx status code
func (o *ListClusterTokensOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list cluster tokens o k response has a 5xx status code
func (o *ListClusterTokensOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster tokens o k response a status code equal to that given
func (o *ListClusterTokensOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list cluster tokens o k response
func (o *ListClusterTokensOK) Code() int {
	return 200
}

func (o *ListClusterTokensOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/{clusterName}/tokens/][%d] listClusterTokensOK %s", 200, payload)
}

func (o *ListClusterTokensOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{orgName}/clusters/{clusterName}/tokens/][%d] listClusterTokensOK %s", 200, payload)
}

func (o *ListClusterTokensOK) GetPayload() []*models.ClusterToken {
	return o.Payload
}

func (o *ListClusterTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
