# EnvValueFromFork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** |  | [optional] 
**ForkOf** | Pointer to [**ForkOf**](ForkOf.md) |  | [optional] 

## Methods

### NewEnvValueFromFork

`func NewEnvValueFromFork() *EnvValueFromFork`

NewEnvValueFromFork instantiates a new EnvValueFromFork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvValueFromForkWithDefaults

`func NewEnvValueFromForkWithDefaults() *EnvValueFromFork`

NewEnvValueFromForkWithDefaults instantiates a new EnvValueFromFork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *EnvValueFromFork) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *EnvValueFromFork) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *EnvValueFromFork) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *EnvValueFromFork) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetForkOf

`func (o *EnvValueFromFork) GetForkOf() ForkOf`

GetForkOf returns the ForkOf field if non-nil, zero value otherwise.

### GetForkOfOk

`func (o *EnvValueFromFork) GetForkOfOk() (*ForkOf, bool)`

GetForkOfOk returns a tuple with the ForkOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkOf

`func (o *EnvValueFromFork) SetForkOf(v ForkOf)`

SetForkOf sets ForkOf field to given value.

### HasForkOf

`func (o *EnvValueFromFork) HasForkOf() bool`

HasForkOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


