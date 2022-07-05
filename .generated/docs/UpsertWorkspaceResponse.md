# UpsertWorkspaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviewEndpoints** | Pointer to [**[]PreviewEndpoint**](PreviewEndpoint.md) |  | [optional] 
**PreviewURLs** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**WorkspaceIDs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpsertWorkspaceResponse

`func NewUpsertWorkspaceResponse() *UpsertWorkspaceResponse`

NewUpsertWorkspaceResponse instantiates a new UpsertWorkspaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertWorkspaceResponseWithDefaults

`func NewUpsertWorkspaceResponseWithDefaults() *UpsertWorkspaceResponse`

NewUpsertWorkspaceResponseWithDefaults instantiates a new UpsertWorkspaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviewEndpoints

`func (o *UpsertWorkspaceResponse) GetPreviewEndpoints() []PreviewEndpoint`

GetPreviewEndpoints returns the PreviewEndpoints field if non-nil, zero value otherwise.

### GetPreviewEndpointsOk

`func (o *UpsertWorkspaceResponse) GetPreviewEndpointsOk() (*[]PreviewEndpoint, bool)`

GetPreviewEndpointsOk returns a tuple with the PreviewEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewEndpoints

`func (o *UpsertWorkspaceResponse) SetPreviewEndpoints(v []PreviewEndpoint)`

SetPreviewEndpoints sets PreviewEndpoints field to given value.

### HasPreviewEndpoints

`func (o *UpsertWorkspaceResponse) HasPreviewEndpoints() bool`

HasPreviewEndpoints returns a boolean if a field has been set.

### GetPreviewURLs

`func (o *UpsertWorkspaceResponse) GetPreviewURLs() []string`

GetPreviewURLs returns the PreviewURLs field if non-nil, zero value otherwise.

### GetPreviewURLsOk

`func (o *UpsertWorkspaceResponse) GetPreviewURLsOk() (*[]string, bool)`

GetPreviewURLsOk returns a tuple with the PreviewURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewURLs

`func (o *UpsertWorkspaceResponse) SetPreviewURLs(v []string)`

SetPreviewURLs sets PreviewURLs field to given value.

### HasPreviewURLs

`func (o *UpsertWorkspaceResponse) HasPreviewURLs() bool`

HasPreviewURLs returns a boolean if a field has been set.

### GetWarnings

`func (o *UpsertWorkspaceResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *UpsertWorkspaceResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *UpsertWorkspaceResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *UpsertWorkspaceResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetWorkspaceIDs

`func (o *UpsertWorkspaceResponse) GetWorkspaceIDs() []string`

GetWorkspaceIDs returns the WorkspaceIDs field if non-nil, zero value otherwise.

### GetWorkspaceIDsOk

`func (o *UpsertWorkspaceResponse) GetWorkspaceIDsOk() (*[]string, bool)`

GetWorkspaceIDsOk returns a tuple with the WorkspaceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceIDs

`func (o *UpsertWorkspaceResponse) SetWorkspaceIDs(v []string)`

SetWorkspaceIDs sets WorkspaceIDs field to given value.

### HasWorkspaceIDs

`func (o *UpsertWorkspaceResponse) HasWorkspaceIDs() bool`

HasWorkspaceIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


