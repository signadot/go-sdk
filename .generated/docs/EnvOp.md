# EnvOp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | name of container to which it applies | [optional] 
**Name** | Pointer to **string** | environmental variable name | [optional] 
**Operation** | Pointer to **string** | upsert or delete | [optional] 
**Value** | Pointer to **string** | environmental variable value | [optional] 
**ValueFrom** | Pointer to [**EnvValueFrom**](EnvValueFrom.md) |  | [optional] 

## Methods

### NewEnvOp

`func NewEnvOp() *EnvOp`

NewEnvOp instantiates a new EnvOp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvOpWithDefaults

`func NewEnvOpWithDefaults() *EnvOp`

NewEnvOpWithDefaults instantiates a new EnvOp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *EnvOp) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *EnvOp) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *EnvOp) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *EnvOp) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetName

`func (o *EnvOp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvOp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvOp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvOp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *EnvOp) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *EnvOp) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *EnvOp) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *EnvOp) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetValue

`func (o *EnvOp) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvOp) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvOp) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvOp) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueFrom

`func (o *EnvOp) GetValueFrom() EnvValueFrom`

GetValueFrom returns the ValueFrom field if non-nil, zero value otherwise.

### GetValueFromOk

`func (o *EnvOp) GetValueFromOk() (*EnvValueFrom, bool)`

GetValueFromOk returns a tuple with the ValueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFrom

`func (o *EnvOp) SetValueFrom(v EnvValueFrom)`

SetValueFrom sets ValueFrom field to given value.

### HasValueFrom

`func (o *EnvOp) HasValueFrom() bool`

HasValueFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


