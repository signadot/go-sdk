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

// EnvValueFromResource struct for EnvValueFromResource
type EnvValueFromResource struct {
	Name *string `json:"name,omitempty"`
	OutputKey *string `json:"outputKey,omitempty"`
}

// NewEnvValueFromResource instantiates a new EnvValueFromResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvValueFromResource() *EnvValueFromResource {
	this := EnvValueFromResource{}
	return &this
}

// NewEnvValueFromResourceWithDefaults instantiates a new EnvValueFromResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvValueFromResourceWithDefaults() *EnvValueFromResource {
	this := EnvValueFromResource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvValueFromResource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvValueFromResource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvValueFromResource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvValueFromResource) SetName(v string) {
	o.Name = &v
}

// GetOutputKey returns the OutputKey field value if set, zero value otherwise.
func (o *EnvValueFromResource) GetOutputKey() string {
	if o == nil || o.OutputKey == nil {
		var ret string
		return ret
	}
	return *o.OutputKey
}

// GetOutputKeyOk returns a tuple with the OutputKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvValueFromResource) GetOutputKeyOk() (*string, bool) {
	if o == nil || o.OutputKey == nil {
		return nil, false
	}
	return o.OutputKey, true
}

// HasOutputKey returns a boolean if a field has been set.
func (o *EnvValueFromResource) HasOutputKey() bool {
	if o != nil && o.OutputKey != nil {
		return true
	}

	return false
}

// SetOutputKey gets a reference to the given string and assigns it to the OutputKey field.
func (o *EnvValueFromResource) SetOutputKey(v string) {
	o.OutputKey = &v
}

func (o EnvValueFromResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OutputKey != nil {
		toSerialize["outputKey"] = o.OutputKey
	}
	return json.Marshal(toSerialize)
}

type NullableEnvValueFromResource struct {
	value *EnvValueFromResource
	isSet bool
}

func (v NullableEnvValueFromResource) Get() *EnvValueFromResource {
	return v.value
}

func (v *NullableEnvValueFromResource) Set(val *EnvValueFromResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvValueFromResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvValueFromResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvValueFromResource(val *EnvValueFromResource) *NullableEnvValueFromResource {
	return &NullableEnvValueFromResource{value: val, isSet: true}
}

func (v NullableEnvValueFromResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvValueFromResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


