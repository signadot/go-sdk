# V1ImageReplacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name specifies which image name in live workloads will be replaced.  Example: us.gcr.io/my-staging-registry/widget | [optional] 
**Namespace** | Pointer to **string** | Namespace optionally specifies which namespace will be searched. | [optional] 
**NewName** | Pointer to **string** | NewName provides a replacement for the image name (the part before the tag). If this is left unset, the image name will not be changed.  Example: us.gcr.io/my-dev-registry/username/widget | [optional] 
**NewTag** | Pointer to **string** | NewTag provides a replacement tag for the image. If this is left unset, the image tag will not be changed.  Example: v1.0.0-snapshot-abc123 | [optional] 

## Methods

### NewV1ImageReplacement

`func NewV1ImageReplacement() *V1ImageReplacement`

NewV1ImageReplacement instantiates a new V1ImageReplacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageReplacementWithDefaults

`func NewV1ImageReplacementWithDefaults() *V1ImageReplacement`

NewV1ImageReplacementWithDefaults instantiates a new V1ImageReplacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1ImageReplacement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ImageReplacement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ImageReplacement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ImageReplacement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1ImageReplacement) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1ImageReplacement) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1ImageReplacement) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1ImageReplacement) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNewName

`func (o *V1ImageReplacement) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *V1ImageReplacement) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *V1ImageReplacement) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *V1ImageReplacement) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNewTag

`func (o *V1ImageReplacement) GetNewTag() string`

GetNewTag returns the NewTag field if non-nil, zero value otherwise.

### GetNewTagOk

`func (o *V1ImageReplacement) GetNewTagOk() (*string, bool)`

GetNewTagOk returns a tuple with the NewTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTag

`func (o *V1ImageReplacement) SetNewTag(v string)`

SetNewTag sets NewTag field to given value.

### HasNewTag

`func (o *V1ImageReplacement) HasNewTag() bool`

HasNewTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


