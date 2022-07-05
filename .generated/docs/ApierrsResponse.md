# ApierrsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**RequestID** | Pointer to **string** |  | [optional] 

## Methods

### NewApierrsResponse

`func NewApierrsResponse() *ApierrsResponse`

NewApierrsResponse instantiates a new ApierrsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApierrsResponseWithDefaults

`func NewApierrsResponseWithDefaults() *ApierrsResponse`

NewApierrsResponseWithDefaults instantiates a new ApierrsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApierrsResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApierrsResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApierrsResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApierrsResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestID

`func (o *ApierrsResponse) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *ApierrsResponse) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *ApierrsResponse) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *ApierrsResponse) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


