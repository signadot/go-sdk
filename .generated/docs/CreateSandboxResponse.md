# CreateSandboxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviewEndpoints** | Pointer to [**[]PreviewEndpoint**](PreviewEndpoint.md) |  | [optional] 
**SandboxID** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSandboxResponse

`func NewCreateSandboxResponse() *CreateSandboxResponse`

NewCreateSandboxResponse instantiates a new CreateSandboxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSandboxResponseWithDefaults

`func NewCreateSandboxResponseWithDefaults() *CreateSandboxResponse`

NewCreateSandboxResponseWithDefaults instantiates a new CreateSandboxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviewEndpoints

`func (o *CreateSandboxResponse) GetPreviewEndpoints() []PreviewEndpoint`

GetPreviewEndpoints returns the PreviewEndpoints field if non-nil, zero value otherwise.

### GetPreviewEndpointsOk

`func (o *CreateSandboxResponse) GetPreviewEndpointsOk() (*[]PreviewEndpoint, bool)`

GetPreviewEndpointsOk returns a tuple with the PreviewEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewEndpoints

`func (o *CreateSandboxResponse) SetPreviewEndpoints(v []PreviewEndpoint)`

SetPreviewEndpoints sets PreviewEndpoints field to given value.

### HasPreviewEndpoints

`func (o *CreateSandboxResponse) HasPreviewEndpoints() bool`

HasPreviewEndpoints returns a boolean if a field has been set.

### GetSandboxID

`func (o *CreateSandboxResponse) GetSandboxID() string`

GetSandboxID returns the SandboxID field if non-nil, zero value otherwise.

### GetSandboxIDOk

`func (o *CreateSandboxResponse) GetSandboxIDOk() (*string, bool)`

GetSandboxIDOk returns a tuple with the SandboxID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxID

`func (o *CreateSandboxResponse) SetSandboxID(v string)`

SetSandboxID sets SandboxID field to given value.

### HasSandboxID

`func (o *CreateSandboxResponse) HasSandboxID() bool`

HasSandboxID returns a boolean if a field has been set.

### GetWarnings

`func (o *CreateSandboxResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CreateSandboxResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CreateSandboxResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CreateSandboxResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


