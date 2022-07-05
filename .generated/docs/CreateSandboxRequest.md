# CreateSandboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Cluster within which this sandbox should be created | 
**Description** | Pointer to **string** | Description of the purpose of this sandbox | [optional] 
**Endpoints** | Pointer to [**[]SandboxEndpoint**](SandboxEndpoint.md) | Endpoints that can be used to point to external DNS names or ingress gateways | [optional] 
**Forks** | [**[]SandboxFork**](SandboxFork.md) | Forks is the specification of each forked entity | 
**Name** | Pointer to **string** | Human-readable name of this sandbox | [optional] 
**Resources** | Pointer to [**[]SandboxResource**](SandboxResource.md) | Resources specifies each required resource to spin up the sandbox | [optional] 

## Methods

### NewCreateSandboxRequest

`func NewCreateSandboxRequest(cluster string, forks []SandboxFork, ) *CreateSandboxRequest`

NewCreateSandboxRequest instantiates a new CreateSandboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSandboxRequestWithDefaults

`func NewCreateSandboxRequestWithDefaults() *CreateSandboxRequest`

NewCreateSandboxRequestWithDefaults instantiates a new CreateSandboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *CreateSandboxRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CreateSandboxRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CreateSandboxRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetDescription

`func (o *CreateSandboxRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSandboxRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSandboxRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSandboxRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoints

`func (o *CreateSandboxRequest) GetEndpoints() []SandboxEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CreateSandboxRequest) GetEndpointsOk() (*[]SandboxEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CreateSandboxRequest) SetEndpoints(v []SandboxEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *CreateSandboxRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetForks

`func (o *CreateSandboxRequest) GetForks() []SandboxFork`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *CreateSandboxRequest) GetForksOk() (*[]SandboxFork, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *CreateSandboxRequest) SetForks(v []SandboxFork)`

SetForks sets Forks field to given value.


### GetName

`func (o *CreateSandboxRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSandboxRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSandboxRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSandboxRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *CreateSandboxRequest) GetResources() []SandboxResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CreateSandboxRequest) GetResourcesOk() (*[]SandboxResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CreateSandboxRequest) SetResources(v []SandboxResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CreateSandboxRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


