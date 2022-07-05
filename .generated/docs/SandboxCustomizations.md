# SandboxCustomizations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**[]EnvOp**](EnvOp.md) | Env var modifications that will be applied to the forked workload | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) | One or more docker images that will be applied to the forked workload | [optional] 
**Patch** | Pointer to [**CustomPatch**](CustomPatch.md) |  | [optional] 

## Methods

### NewSandboxCustomizations

`func NewSandboxCustomizations() *SandboxCustomizations`

NewSandboxCustomizations instantiates a new SandboxCustomizations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCustomizationsWithDefaults

`func NewSandboxCustomizationsWithDefaults() *SandboxCustomizations`

NewSandboxCustomizationsWithDefaults instantiates a new SandboxCustomizations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *SandboxCustomizations) GetEnv() []EnvOp`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *SandboxCustomizations) GetEnvOk() (*[]EnvOp, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *SandboxCustomizations) SetEnv(v []EnvOp)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *SandboxCustomizations) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetImages

`func (o *SandboxCustomizations) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SandboxCustomizations) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SandboxCustomizations) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *SandboxCustomizations) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetPatch

`func (o *SandboxCustomizations) GetPatch() CustomPatch`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *SandboxCustomizations) GetPatchOk() (*CustomPatch, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *SandboxCustomizations) SetPatch(v CustomPatch)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *SandboxCustomizations) HasPatch() bool`

HasPatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


