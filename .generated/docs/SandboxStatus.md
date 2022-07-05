# SandboxStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the sandbox. | [optional] 
**Message** | Pointer to **string** | Message is a human readable explanation of why the sandbox is healthy or not. | [optional] 
**Ready** | Pointer to **bool** | Ready indicates whether the sandbox is ready, meaning that it can be used for testing. | [optional] 
**Reason** | Pointer to **string** | Reason is a machine readable explanation of why the sandbox is healthy or not. | [optional] 

## Methods

### NewSandboxStatus

`func NewSandboxStatus() *SandboxStatus`

NewSandboxStatus instantiates a new SandboxStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxStatusWithDefaults

`func NewSandboxStatusWithDefaults() *SandboxStatus`

NewSandboxStatusWithDefaults instantiates a new SandboxStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SandboxStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *SandboxStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SandboxStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SandboxStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SandboxStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReady

`func (o *SandboxStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *SandboxStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *SandboxStatus) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *SandboxStatus) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetReason

`func (o *SandboxStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SandboxStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SandboxStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SandboxStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


