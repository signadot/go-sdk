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

// GetSandboxesResponse struct for GetSandboxesResponse
type GetSandboxesResponse struct {
	Sandboxes []SandboxInfo `json:"sandboxes,omitempty"`
}

// NewGetSandboxesResponse instantiates a new GetSandboxesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSandboxesResponse() *GetSandboxesResponse {
	this := GetSandboxesResponse{}
	return &this
}

// NewGetSandboxesResponseWithDefaults instantiates a new GetSandboxesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSandboxesResponseWithDefaults() *GetSandboxesResponse {
	this := GetSandboxesResponse{}
	return &this
}

// GetSandboxes returns the Sandboxes field value if set, zero value otherwise.
func (o *GetSandboxesResponse) GetSandboxes() []SandboxInfo {
	if o == nil || o.Sandboxes == nil {
		var ret []SandboxInfo
		return ret
	}
	return o.Sandboxes
}

// GetSandboxesOk returns a tuple with the Sandboxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSandboxesResponse) GetSandboxesOk() ([]SandboxInfo, bool) {
	if o == nil || o.Sandboxes == nil {
		return nil, false
	}
	return o.Sandboxes, true
}

// HasSandboxes returns a boolean if a field has been set.
func (o *GetSandboxesResponse) HasSandboxes() bool {
	if o != nil && o.Sandboxes != nil {
		return true
	}

	return false
}

// SetSandboxes gets a reference to the given []SandboxInfo and assigns it to the Sandboxes field.
func (o *GetSandboxesResponse) SetSandboxes(v []SandboxInfo) {
	o.Sandboxes = v
}

func (o GetSandboxesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sandboxes != nil {
		toSerialize["sandboxes"] = o.Sandboxes
	}
	return json.Marshal(toSerialize)
}

type NullableGetSandboxesResponse struct {
	value *GetSandboxesResponse
	isSet bool
}

func (v NullableGetSandboxesResponse) Get() *GetSandboxesResponse {
	return v.value
}

func (v *NullableGetSandboxesResponse) Set(val *GetSandboxesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSandboxesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSandboxesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSandboxesResponse(val *GetSandboxesResponse) *NullableGetSandboxesResponse {
	return &NullableGetSandboxesResponse{value: val, isSet: true}
}

func (v NullableGetSandboxesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSandboxesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


