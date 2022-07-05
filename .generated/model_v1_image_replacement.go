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

// V1ImageReplacement struct for V1ImageReplacement
type V1ImageReplacement struct {
	// Name specifies which image name in live workloads will be replaced.  Example: us.gcr.io/my-staging-registry/widget
	Name *string `json:"name,omitempty"`
	// Namespace optionally specifies which namespace will be searched.
	Namespace *string `json:"namespace,omitempty"`
	// NewName provides a replacement for the image name (the part before the tag). If this is left unset, the image name will not be changed.  Example: us.gcr.io/my-dev-registry/username/widget
	NewName *string `json:"newName,omitempty"`
	// NewTag provides a replacement tag for the image. If this is left unset, the image tag will not be changed.  Example: v1.0.0-snapshot-abc123
	NewTag *string `json:"newTag,omitempty"`
}

// NewV1ImageReplacement instantiates a new V1ImageReplacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageReplacement() *V1ImageReplacement {
	this := V1ImageReplacement{}
	return &this
}

// NewV1ImageReplacementWithDefaults instantiates a new V1ImageReplacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageReplacementWithDefaults() *V1ImageReplacement {
	this := V1ImageReplacement{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ImageReplacement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageReplacement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ImageReplacement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ImageReplacement) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1ImageReplacement) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageReplacement) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1ImageReplacement) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1ImageReplacement) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *V1ImageReplacement) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageReplacement) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *V1ImageReplacement) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *V1ImageReplacement) SetNewName(v string) {
	o.NewName = &v
}

// GetNewTag returns the NewTag field value if set, zero value otherwise.
func (o *V1ImageReplacement) GetNewTag() string {
	if o == nil || o.NewTag == nil {
		var ret string
		return ret
	}
	return *o.NewTag
}

// GetNewTagOk returns a tuple with the NewTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageReplacement) GetNewTagOk() (*string, bool) {
	if o == nil || o.NewTag == nil {
		return nil, false
	}
	return o.NewTag, true
}

// HasNewTag returns a boolean if a field has been set.
func (o *V1ImageReplacement) HasNewTag() bool {
	if o != nil && o.NewTag != nil {
		return true
	}

	return false
}

// SetNewTag gets a reference to the given string and assigns it to the NewTag field.
func (o *V1ImageReplacement) SetNewTag(v string) {
	o.NewTag = &v
}

func (o V1ImageReplacement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NewName != nil {
		toSerialize["newName"] = o.NewName
	}
	if o.NewTag != nil {
		toSerialize["newTag"] = o.NewTag
	}
	return json.Marshal(toSerialize)
}

type NullableV1ImageReplacement struct {
	value *V1ImageReplacement
	isSet bool
}

func (v NullableV1ImageReplacement) Get() *V1ImageReplacement {
	return v.value
}

func (v *NullableV1ImageReplacement) Set(val *V1ImageReplacement) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageReplacement) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageReplacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageReplacement(val *V1ImageReplacement) *NullableV1ImageReplacement {
	return &NullableV1ImageReplacement{value: val, isSet: true}
}

func (v NullableV1ImageReplacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageReplacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


