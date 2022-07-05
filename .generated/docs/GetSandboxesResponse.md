# GetSandboxesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sandboxes** | Pointer to [**[]SandboxInfo**](SandboxInfo.md) |  | [optional] 

## Methods

### NewGetSandboxesResponse

`func NewGetSandboxesResponse() *GetSandboxesResponse`

NewGetSandboxesResponse instantiates a new GetSandboxesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSandboxesResponseWithDefaults

`func NewGetSandboxesResponseWithDefaults() *GetSandboxesResponse`

NewGetSandboxesResponseWithDefaults instantiates a new GetSandboxesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSandboxes

`func (o *GetSandboxesResponse) GetSandboxes() []SandboxInfo`

GetSandboxes returns the Sandboxes field if non-nil, zero value otherwise.

### GetSandboxesOk

`func (o *GetSandboxesResponse) GetSandboxesOk() (*[]SandboxInfo, bool)`

GetSandboxesOk returns a tuple with the Sandboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxes

`func (o *GetSandboxesResponse) SetSandboxes(v []SandboxInfo)`

SetSandboxes sets Sandboxes field to given value.

### HasSandboxes

`func (o *GetSandboxesResponse) HasSandboxes() bool`

HasSandboxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


