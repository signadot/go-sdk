# CreatePreviewEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForkOf** | Pointer to [**ForkOf**](ForkOf.md) |  | [optional] 
**Host** | Pointer to **string** | Host is the host-name of a service within Kubernetes of the form &#x60;service-name.namespace.svc&#x60;. Only required when RouteType is &#39;static&#39; | [optional] 
**Name** | Pointer to **string** | Name of this endpoint. If not specified, a name will be generated automatically. | [optional] 
**Port** | Pointer to **int32** | Port is the port on the above host that the preview will point to. Only required when RouteType is &#39;static&#39; | [optional] 
**Protocol** | Pointer to **string** | Protocol of the endpoint that we will be connecting to for this preview URL. One of &#39;http&#39;, &#39;https&#39;, or &#39;grpc&#39;. If not specified, the default is &#39;http&#39;. | [optional] 
**RouteType** | **string** | RouteType is one of &#39;static&#39; or &#39;fork&#39;. If you choose a route of type &#39;static&#39;, this preview endpoint will route traffic to the Host / Port specified. If you choose a route of type &#39;fork&#39;, this preview endpoint will route traffic to a forked entity as specified in forkOf. | 

## Methods

### NewCreatePreviewEndpointRequest

`func NewCreatePreviewEndpointRequest(routeType string, ) *CreatePreviewEndpointRequest`

NewCreatePreviewEndpointRequest instantiates a new CreatePreviewEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePreviewEndpointRequestWithDefaults

`func NewCreatePreviewEndpointRequestWithDefaults() *CreatePreviewEndpointRequest`

NewCreatePreviewEndpointRequestWithDefaults instantiates a new CreatePreviewEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForkOf

`func (o *CreatePreviewEndpointRequest) GetForkOf() ForkOf`

GetForkOf returns the ForkOf field if non-nil, zero value otherwise.

### GetForkOfOk

`func (o *CreatePreviewEndpointRequest) GetForkOfOk() (*ForkOf, bool)`

GetForkOfOk returns a tuple with the ForkOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkOf

`func (o *CreatePreviewEndpointRequest) SetForkOf(v ForkOf)`

SetForkOf sets ForkOf field to given value.

### HasForkOf

`func (o *CreatePreviewEndpointRequest) HasForkOf() bool`

HasForkOf returns a boolean if a field has been set.

### GetHost

`func (o *CreatePreviewEndpointRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreatePreviewEndpointRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreatePreviewEndpointRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreatePreviewEndpointRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *CreatePreviewEndpointRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePreviewEndpointRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePreviewEndpointRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePreviewEndpointRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *CreatePreviewEndpointRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreatePreviewEndpointRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreatePreviewEndpointRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreatePreviewEndpointRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *CreatePreviewEndpointRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreatePreviewEndpointRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreatePreviewEndpointRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CreatePreviewEndpointRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRouteType

`func (o *CreatePreviewEndpointRequest) GetRouteType() string`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *CreatePreviewEndpointRequest) GetRouteTypeOk() (*string, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *CreatePreviewEndpointRequest) SetRouteType(v string)`

SetRouteType sets RouteType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


