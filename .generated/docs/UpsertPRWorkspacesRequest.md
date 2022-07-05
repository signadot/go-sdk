# UpsertPRWorkspacesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Cluster within which this workspace should be created | 
**DefaultService** | Pointer to **string** | Deprecated. use endpoints instead. | [optional] 
**DefaultServicePort** | Pointer to **int32** | Deprecated. use endpoints instead. | [optional] 
**Description** | Pointer to **string** | Description of the purpose of this workspace. | [optional] 
**Endpoints** | Pointer to [**[]CreatePreviewEndpointRequest**](CreatePreviewEndpointRequest.md) | Each endpoint specifies a target service or workload corresponding to which a preview URL will be generated. | [optional] 
**HeadCommit** | Pointer to **string** | HeadCommit is the commit hash of the current HEAD of the PR branch. It is automatically computed if not specified but it is recommended that you specify this. | [optional] 
**Images** | Pointer to [**[]V1ImageReplacement**](V1ImageReplacement.md) | Image replacement rules that are used to create the workspace. | [optional] 
**Name** | Pointer to **string** | Human-readable name of this workspace | [optional] 
**Namespace** | **string** | Namespace within which this workspace should be created | 

## Methods

### NewUpsertPRWorkspacesRequest

`func NewUpsertPRWorkspacesRequest(cluster string, namespace string, ) *UpsertPRWorkspacesRequest`

NewUpsertPRWorkspacesRequest instantiates a new UpsertPRWorkspacesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertPRWorkspacesRequestWithDefaults

`func NewUpsertPRWorkspacesRequestWithDefaults() *UpsertPRWorkspacesRequest`

NewUpsertPRWorkspacesRequestWithDefaults instantiates a new UpsertPRWorkspacesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *UpsertPRWorkspacesRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpsertPRWorkspacesRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpsertPRWorkspacesRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetDefaultService

`func (o *UpsertPRWorkspacesRequest) GetDefaultService() string`

GetDefaultService returns the DefaultService field if non-nil, zero value otherwise.

### GetDefaultServiceOk

`func (o *UpsertPRWorkspacesRequest) GetDefaultServiceOk() (*string, bool)`

GetDefaultServiceOk returns a tuple with the DefaultService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultService

`func (o *UpsertPRWorkspacesRequest) SetDefaultService(v string)`

SetDefaultService sets DefaultService field to given value.

### HasDefaultService

`func (o *UpsertPRWorkspacesRequest) HasDefaultService() bool`

HasDefaultService returns a boolean if a field has been set.

### GetDefaultServicePort

`func (o *UpsertPRWorkspacesRequest) GetDefaultServicePort() int32`

GetDefaultServicePort returns the DefaultServicePort field if non-nil, zero value otherwise.

### GetDefaultServicePortOk

`func (o *UpsertPRWorkspacesRequest) GetDefaultServicePortOk() (*int32, bool)`

GetDefaultServicePortOk returns a tuple with the DefaultServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServicePort

`func (o *UpsertPRWorkspacesRequest) SetDefaultServicePort(v int32)`

SetDefaultServicePort sets DefaultServicePort field to given value.

### HasDefaultServicePort

`func (o *UpsertPRWorkspacesRequest) HasDefaultServicePort() bool`

HasDefaultServicePort returns a boolean if a field has been set.

### GetDescription

`func (o *UpsertPRWorkspacesRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsertPRWorkspacesRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsertPRWorkspacesRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpsertPRWorkspacesRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoints

`func (o *UpsertPRWorkspacesRequest) GetEndpoints() []CreatePreviewEndpointRequest`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *UpsertPRWorkspacesRequest) GetEndpointsOk() (*[]CreatePreviewEndpointRequest, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *UpsertPRWorkspacesRequest) SetEndpoints(v []CreatePreviewEndpointRequest)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *UpsertPRWorkspacesRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetHeadCommit

`func (o *UpsertPRWorkspacesRequest) GetHeadCommit() string`

GetHeadCommit returns the HeadCommit field if non-nil, zero value otherwise.

### GetHeadCommitOk

`func (o *UpsertPRWorkspacesRequest) GetHeadCommitOk() (*string, bool)`

GetHeadCommitOk returns a tuple with the HeadCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommit

`func (o *UpsertPRWorkspacesRequest) SetHeadCommit(v string)`

SetHeadCommit sets HeadCommit field to given value.

### HasHeadCommit

`func (o *UpsertPRWorkspacesRequest) HasHeadCommit() bool`

HasHeadCommit returns a boolean if a field has been set.

### GetImages

`func (o *UpsertPRWorkspacesRequest) GetImages() []V1ImageReplacement`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *UpsertPRWorkspacesRequest) GetImagesOk() (*[]V1ImageReplacement, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *UpsertPRWorkspacesRequest) SetImages(v []V1ImageReplacement)`

SetImages sets Images field to given value.

### HasImages

`func (o *UpsertPRWorkspacesRequest) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetName

`func (o *UpsertPRWorkspacesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertPRWorkspacesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertPRWorkspacesRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpsertPRWorkspacesRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *UpsertPRWorkspacesRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpsertPRWorkspacesRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpsertPRWorkspacesRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


