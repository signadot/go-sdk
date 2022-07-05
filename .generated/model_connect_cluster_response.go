/*
Signadot API

API for Signadot Sandboxes

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConnectClusterResponse struct for ConnectClusterResponse
type ConnectClusterResponse struct {
	Id *string `json:"id,omitempty"`
}

// NewConnectClusterResponse instantiates a new ConnectClusterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectClusterResponse() *ConnectClusterResponse {
	this := ConnectClusterResponse{}
	return &this
}

// NewConnectClusterResponseWithDefaults instantiates a new ConnectClusterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectClusterResponseWithDefaults() *ConnectClusterResponse {
	this := ConnectClusterResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectClusterResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectClusterResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectClusterResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectClusterResponse) SetId(v string) {
	o.Id = &v
}

func (o ConnectClusterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableConnectClusterResponse struct {
	value *ConnectClusterResponse
	isSet bool
}

func (v NullableConnectClusterResponse) Get() *ConnectClusterResponse {
	return v.value
}

func (v *NullableConnectClusterResponse) Set(val *ConnectClusterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectClusterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectClusterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectClusterResponse(val *ConnectClusterResponse) *NullableConnectClusterResponse {
	return &NullableConnectClusterResponse{value: val, isSet: true}
}

func (v NullableConnectClusterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectClusterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


