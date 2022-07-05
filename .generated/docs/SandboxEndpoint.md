# SandboxEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Hostname that this endpoint points to | [optional] 
**Name** | Pointer to **string** | Name of the endpoint | [optional] 
**Port** | Pointer to **int32** | Port it will map to on the specified host | [optional] 
**Protocol** | Pointer to **string** | Protocol that this endpoint uses | [optional] 

## Methods

### NewSandboxEndpoint

`func NewSandboxEndpoint() *SandboxEndpoint`

NewSandboxEndpoint instantiates a new SandboxEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxEndpointWithDefaults

`func NewSandboxEndpointWithDefaults() *SandboxEndpoint`

NewSandboxEndpointWithDefaults instantiates a new SandboxEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SandboxEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SandboxEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SandboxEndpoint) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SandboxEndpoint) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *SandboxEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SandboxEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SandboxEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SandboxEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *SandboxEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SandboxEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SandboxEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SandboxEndpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *SandboxEndpoint) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SandboxEndpoint) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SandboxEndpoint) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SandboxEndpoint) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


