# SandboxDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]Branch**](Branch.md) |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DefaultService** | Pointer to **string** |  | [optional] 
**DefaultServicePort** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**PreviewEndpoints** | Pointer to [**[]PreviewEndpoint**](PreviewEndpoint.md) |  | [optional] 
**PreviewURL** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewSandboxDetails

`func NewSandboxDetails() *SandboxDetails`

NewSandboxDetails instantiates a new SandboxDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxDetailsWithDefaults

`func NewSandboxDetailsWithDefaults() *SandboxDetails`

NewSandboxDetailsWithDefaults instantiates a new SandboxDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *SandboxDetails) GetBranches() []Branch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *SandboxDetails) GetBranchesOk() (*[]Branch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *SandboxDetails) SetBranches(v []Branch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *SandboxDetails) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetClusterName

`func (o *SandboxDetails) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *SandboxDetails) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *SandboxDetails) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *SandboxDetails) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SandboxDetails) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SandboxDetails) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SandboxDetails) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SandboxDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultService

`func (o *SandboxDetails) GetDefaultService() string`

GetDefaultService returns the DefaultService field if non-nil, zero value otherwise.

### GetDefaultServiceOk

`func (o *SandboxDetails) GetDefaultServiceOk() (*string, bool)`

GetDefaultServiceOk returns a tuple with the DefaultService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultService

`func (o *SandboxDetails) SetDefaultService(v string)`

SetDefaultService sets DefaultService field to given value.

### HasDefaultService

`func (o *SandboxDetails) HasDefaultService() bool`

HasDefaultService returns a boolean if a field has been set.

### GetDefaultServicePort

`func (o *SandboxDetails) GetDefaultServicePort() int32`

GetDefaultServicePort returns the DefaultServicePort field if non-nil, zero value otherwise.

### GetDefaultServicePortOk

`func (o *SandboxDetails) GetDefaultServicePortOk() (*int32, bool)`

GetDefaultServicePortOk returns a tuple with the DefaultServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServicePort

`func (o *SandboxDetails) SetDefaultServicePort(v int32)`

SetDefaultServicePort sets DefaultServicePort field to given value.

### HasDefaultServicePort

`func (o *SandboxDetails) HasDefaultServicePort() bool`

HasDefaultServicePort returns a boolean if a field has been set.

### GetDescription

`func (o *SandboxDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SandboxDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SandboxDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SandboxDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SandboxDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SandboxDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SandboxDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SandboxDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SandboxDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SandboxDetails) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SandboxDetails) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SandboxDetails) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SandboxDetails) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPreviewEndpoints

`func (o *SandboxDetails) GetPreviewEndpoints() []PreviewEndpoint`

GetPreviewEndpoints returns the PreviewEndpoints field if non-nil, zero value otherwise.

### GetPreviewEndpointsOk

`func (o *SandboxDetails) GetPreviewEndpointsOk() (*[]PreviewEndpoint, bool)`

GetPreviewEndpointsOk returns a tuple with the PreviewEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewEndpoints

`func (o *SandboxDetails) SetPreviewEndpoints(v []PreviewEndpoint)`

SetPreviewEndpoints sets PreviewEndpoints field to given value.

### HasPreviewEndpoints

`func (o *SandboxDetails) HasPreviewEndpoints() bool`

HasPreviewEndpoints returns a boolean if a field has been set.

### GetPreviewURL

`func (o *SandboxDetails) GetPreviewURL() string`

GetPreviewURL returns the PreviewURL field if non-nil, zero value otherwise.

### GetPreviewURLOk

`func (o *SandboxDetails) GetPreviewURLOk() (*string, bool)`

GetPreviewURLOk returns a tuple with the PreviewURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewURL

`func (o *SandboxDetails) SetPreviewURL(v string)`

SetPreviewURL sets PreviewURL field to given value.

### HasPreviewURL

`func (o *SandboxDetails) HasPreviewURL() bool`

HasPreviewURL returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SandboxDetails) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SandboxDetails) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SandboxDetails) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SandboxDetails) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


