# GetSandboxByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sandbox** | Pointer to [**SandboxDetails**](SandboxDetails.md) |  | [optional] 

## Methods

### NewGetSandboxByIdResponse

`func NewGetSandboxByIdResponse() *GetSandboxByIdResponse`

NewGetSandboxByIdResponse instantiates a new GetSandboxByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSandboxByIdResponseWithDefaults

`func NewGetSandboxByIdResponseWithDefaults() *GetSandboxByIdResponse`

NewGetSandboxByIdResponseWithDefaults instantiates a new GetSandboxByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSandbox

`func (o *GetSandboxByIdResponse) GetSandbox() SandboxDetails`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *GetSandboxByIdResponse) GetSandboxOk() (*SandboxDetails, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *GetSandboxByIdResponse) SetSandbox(v SandboxDetails)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *GetSandboxByIdResponse) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


