# PreviewEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaselinePreviewURL** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ClusterID** | Pointer to **string** |  | [optional] 
**ForkOf** | Pointer to [**ForkOf**](ForkOf.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PreviewURL** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**RouteType** | Pointer to **string** |  | [optional] 

## Methods

### NewPreviewEndpoint

`func NewPreviewEndpoint() *PreviewEndpoint`

NewPreviewEndpoint instantiates a new PreviewEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewEndpointWithDefaults

`func NewPreviewEndpointWithDefaults() *PreviewEndpoint`

NewPreviewEndpointWithDefaults instantiates a new PreviewEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaselinePreviewURL

`func (o *PreviewEndpoint) GetBaselinePreviewURL() string`

GetBaselinePreviewURL returns the BaselinePreviewURL field if non-nil, zero value otherwise.

### GetBaselinePreviewURLOk

`func (o *PreviewEndpoint) GetBaselinePreviewURLOk() (*string, bool)`

GetBaselinePreviewURLOk returns a tuple with the BaselinePreviewURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselinePreviewURL

`func (o *PreviewEndpoint) SetBaselinePreviewURL(v string)`

SetBaselinePreviewURL sets BaselinePreviewURL field to given value.

### HasBaselinePreviewURL

`func (o *PreviewEndpoint) HasBaselinePreviewURL() bool`

HasBaselinePreviewURL returns a boolean if a field has been set.

### GetCluster

`func (o *PreviewEndpoint) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PreviewEndpoint) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PreviewEndpoint) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PreviewEndpoint) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterID

`func (o *PreviewEndpoint) GetClusterID() string`

GetClusterID returns the ClusterID field if non-nil, zero value otherwise.

### GetClusterIDOk

`func (o *PreviewEndpoint) GetClusterIDOk() (*string, bool)`

GetClusterIDOk returns a tuple with the ClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterID

`func (o *PreviewEndpoint) SetClusterID(v string)`

SetClusterID sets ClusterID field to given value.

### HasClusterID

`func (o *PreviewEndpoint) HasClusterID() bool`

HasClusterID returns a boolean if a field has been set.

### GetForkOf

`func (o *PreviewEndpoint) GetForkOf() ForkOf`

GetForkOf returns the ForkOf field if non-nil, zero value otherwise.

### GetForkOfOk

`func (o *PreviewEndpoint) GetForkOfOk() (*ForkOf, bool)`

GetForkOfOk returns a tuple with the ForkOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkOf

`func (o *PreviewEndpoint) SetForkOf(v ForkOf)`

SetForkOf sets ForkOf field to given value.

### HasForkOf

`func (o *PreviewEndpoint) HasForkOf() bool`

HasForkOf returns a boolean if a field has been set.

### GetHost

`func (o *PreviewEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PreviewEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PreviewEndpoint) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PreviewEndpoint) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *PreviewEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PreviewEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PreviewEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PreviewEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PreviewEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PreviewEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PreviewEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PreviewEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *PreviewEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PreviewEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PreviewEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PreviewEndpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPreviewURL

`func (o *PreviewEndpoint) GetPreviewURL() string`

GetPreviewURL returns the PreviewURL field if non-nil, zero value otherwise.

### GetPreviewURLOk

`func (o *PreviewEndpoint) GetPreviewURLOk() (*string, bool)`

GetPreviewURLOk returns a tuple with the PreviewURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewURL

`func (o *PreviewEndpoint) SetPreviewURL(v string)`

SetPreviewURL sets PreviewURL field to given value.

### HasPreviewURL

`func (o *PreviewEndpoint) HasPreviewURL() bool`

HasPreviewURL returns a boolean if a field has been set.

### GetProtocol

`func (o *PreviewEndpoint) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PreviewEndpoint) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PreviewEndpoint) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PreviewEndpoint) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRouteType

`func (o *PreviewEndpoint) GetRouteType() string`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *PreviewEndpoint) GetRouteTypeOk() (*string, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *PreviewEndpoint) SetRouteType(v string)`

SetRouteType sets RouteType field to given value.

### HasRouteType

`func (o *PreviewEndpoint) HasRouteType() bool`

HasRouteType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


