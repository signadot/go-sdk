# SandboxInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSandboxInfo

`func NewSandboxInfo() *SandboxInfo`

NewSandboxInfo instantiates a new SandboxInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxInfoWithDefaults

`func NewSandboxInfoWithDefaults() *SandboxInfo`

NewSandboxInfoWithDefaults instantiates a new SandboxInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *SandboxInfo) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *SandboxInfo) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *SandboxInfo) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *SandboxInfo) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SandboxInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SandboxInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SandboxInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SandboxInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultService

`func (o *SandboxInfo) GetDefaultService() string`

GetDefaultService returns the DefaultService field if non-nil, zero value otherwise.

### GetDefaultServiceOk

`func (o *SandboxInfo) GetDefaultServiceOk() (*string, bool)`

GetDefaultServiceOk returns a tuple with the DefaultService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultService

`func (o *SandboxInfo) SetDefaultService(v string)`

SetDefaultService sets DefaultService field to given value.

### HasDefaultService

`func (o *SandboxInfo) HasDefaultService() bool`

HasDefaultService returns a boolean if a field has been set.

### GetDefaultServicePort

`func (o *SandboxInfo) GetDefaultServicePort() int32`

GetDefaultServicePort returns the DefaultServicePort field if non-nil, zero value otherwise.

### GetDefaultServicePortOk

`func (o *SandboxInfo) GetDefaultServicePortOk() (*int32, bool)`

GetDefaultServicePortOk returns a tuple with the DefaultServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServicePort

`func (o *SandboxInfo) SetDefaultServicePort(v int32)`

SetDefaultServicePort sets DefaultServicePort field to given value.

### HasDefaultServicePort

`func (o *SandboxInfo) HasDefaultServicePort() bool`

HasDefaultServicePort returns a boolean if a field has been set.

### GetDescription

`func (o *SandboxInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SandboxInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SandboxInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SandboxInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SandboxInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SandboxInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SandboxInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SandboxInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SandboxInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SandboxInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SandboxInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SandboxInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SandboxInfo) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPreviewEndpoints

`func (o *SandboxInfo) GetPreviewEndpoints() []PreviewEndpoint`

GetPreviewEndpoints returns the PreviewEndpoints field if non-nil, zero value otherwise.

### GetPreviewEndpointsOk

`func (o *SandboxInfo) GetPreviewEndpointsOk() (*[]PreviewEndpoint, bool)`

GetPreviewEndpointsOk returns a tuple with the PreviewEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewEndpoints

`func (o *SandboxInfo) SetPreviewEndpoints(v []PreviewEndpoint)`

SetPreviewEndpoints sets PreviewEndpoints field to given value.

### HasPreviewEndpoints

`func (o *SandboxInfo) HasPreviewEndpoints() bool`

HasPreviewEndpoints returns a boolean if a field has been set.

### GetPreviewURL

`func (o *SandboxInfo) GetPreviewURL() string`

GetPreviewURL returns the PreviewURL field if non-nil, zero value otherwise.

### GetPreviewURLOk

`func (o *SandboxInfo) GetPreviewURLOk() (*string, bool)`

GetPreviewURLOk returns a tuple with the PreviewURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewURL

`func (o *SandboxInfo) SetPreviewURL(v string)`

SetPreviewURL sets PreviewURL field to given value.

### HasPreviewURL

`func (o *SandboxInfo) HasPreviewURL() bool`

HasPreviewURL returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SandboxInfo) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SandboxInfo) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SandboxInfo) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SandboxInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


