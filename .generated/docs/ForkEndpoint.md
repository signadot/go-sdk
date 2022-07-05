# ForkEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the endpoint | [optional] 
**Port** | Pointer to **int32** | Port it will map to on the forked workload | [optional] 
**Protocol** | Pointer to **string** | Protocol that this endpoint uses | [optional] 

## Methods

### NewForkEndpoint

`func NewForkEndpoint() *ForkEndpoint`

NewForkEndpoint instantiates a new ForkEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForkEndpointWithDefaults

`func NewForkEndpointWithDefaults() *ForkEndpoint`

NewForkEndpointWithDefaults instantiates a new ForkEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ForkEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForkEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForkEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ForkEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ForkEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ForkEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ForkEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ForkEndpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *ForkEndpoint) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ForkEndpoint) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ForkEndpoint) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ForkEndpoint) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


