# ForkOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity that we want to route to. One of (Service or Deployment or Rollout). | 
**Name** | **string** | Name of the entity within the Kubernetes cluster. | 
**Namespace** | **string** | Namespace within which the entity lives in the Kubernetes cluster. | 

## Methods

### NewForkOf

`func NewForkOf(kind string, name string, namespace string, ) *ForkOf`

NewForkOf instantiates a new ForkOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForkOfWithDefaults

`func NewForkOfWithDefaults() *ForkOf`

NewForkOfWithDefaults instantiates a new ForkOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ForkOf) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ForkOf) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ForkOf) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ForkOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForkOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForkOf) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *ForkOf) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ForkOf) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ForkOf) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


