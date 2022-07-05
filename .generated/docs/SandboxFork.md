# SandboxFork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customizations** | Pointer to [**SandboxCustomizations**](SandboxCustomizations.md) |  | [optional] 
**Endpoints** | Pointer to [**[]ForkEndpoint**](ForkEndpoint.md) | Endpoints that correspond to this forked workload | [optional] 
**ForkOf** | Pointer to [**ForkOf**](ForkOf.md) |  | [optional] 

## Methods

### NewSandboxFork

`func NewSandboxFork() *SandboxFork`

NewSandboxFork instantiates a new SandboxFork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxForkWithDefaults

`func NewSandboxForkWithDefaults() *SandboxFork`

NewSandboxForkWithDefaults instantiates a new SandboxFork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomizations

`func (o *SandboxFork) GetCustomizations() SandboxCustomizations`

GetCustomizations returns the Customizations field if non-nil, zero value otherwise.

### GetCustomizationsOk

`func (o *SandboxFork) GetCustomizationsOk() (*SandboxCustomizations, bool)`

GetCustomizationsOk returns a tuple with the Customizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizations

`func (o *SandboxFork) SetCustomizations(v SandboxCustomizations)`

SetCustomizations sets Customizations field to given value.

### HasCustomizations

`func (o *SandboxFork) HasCustomizations() bool`

HasCustomizations returns a boolean if a field has been set.

### GetEndpoints

`func (o *SandboxFork) GetEndpoints() []ForkEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *SandboxFork) GetEndpointsOk() (*[]ForkEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *SandboxFork) SetEndpoints(v []ForkEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *SandboxFork) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetForkOf

`func (o *SandboxFork) GetForkOf() ForkOf`

GetForkOf returns the ForkOf field if non-nil, zero value otherwise.

### GetForkOfOk

`func (o *SandboxFork) GetForkOfOk() (*ForkOf, bool)`

GetForkOfOk returns a tuple with the ForkOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkOf

`func (o *SandboxFork) SetForkOf(v ForkOf)`

SetForkOf sets ForkOf field to given value.

### HasForkOf

`func (o *SandboxFork) HasForkOf() bool`

HasForkOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


