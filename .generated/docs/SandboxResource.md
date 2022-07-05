# SandboxResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**Plugin** | Pointer to **string** |  | [optional] 

## Methods

### NewSandboxResource

`func NewSandboxResource() *SandboxResource`

NewSandboxResource instantiates a new SandboxResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxResourceWithDefaults

`func NewSandboxResourceWithDefaults() *SandboxResource`

NewSandboxResourceWithDefaults instantiates a new SandboxResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SandboxResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SandboxResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SandboxResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SandboxResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParams

`func (o *SandboxResource) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SandboxResource) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SandboxResource) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *SandboxResource) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetPlugin

`func (o *SandboxResource) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *SandboxResource) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *SandboxResource) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *SandboxResource) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


