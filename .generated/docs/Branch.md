# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeadCommit** | Pointer to **string** |  | [optional] 
**PullRequest** | Pointer to **int32** |  | [optional] 
**VcsName** | Pointer to **string** |  | [optional] 
**VcsRepo** | Pointer to **string** |  | [optional] 
**VcsType** | Pointer to **string** |  | [optional] 

## Methods

### NewBranch

`func NewBranch() *Branch`

NewBranch instantiates a new Branch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithDefaults

`func NewBranchWithDefaults() *Branch`

NewBranchWithDefaults instantiates a new Branch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeadCommit

`func (o *Branch) GetHeadCommit() string`

GetHeadCommit returns the HeadCommit field if non-nil, zero value otherwise.

### GetHeadCommitOk

`func (o *Branch) GetHeadCommitOk() (*string, bool)`

GetHeadCommitOk returns a tuple with the HeadCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommit

`func (o *Branch) SetHeadCommit(v string)`

SetHeadCommit sets HeadCommit field to given value.

### HasHeadCommit

`func (o *Branch) HasHeadCommit() bool`

HasHeadCommit returns a boolean if a field has been set.

### GetPullRequest

`func (o *Branch) GetPullRequest() int32`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *Branch) GetPullRequestOk() (*int32, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *Branch) SetPullRequest(v int32)`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *Branch) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetVcsName

`func (o *Branch) GetVcsName() string`

GetVcsName returns the VcsName field if non-nil, zero value otherwise.

### GetVcsNameOk

`func (o *Branch) GetVcsNameOk() (*string, bool)`

GetVcsNameOk returns a tuple with the VcsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsName

`func (o *Branch) SetVcsName(v string)`

SetVcsName sets VcsName field to given value.

### HasVcsName

`func (o *Branch) HasVcsName() bool`

HasVcsName returns a boolean if a field has been set.

### GetVcsRepo

`func (o *Branch) GetVcsRepo() string`

GetVcsRepo returns the VcsRepo field if non-nil, zero value otherwise.

### GetVcsRepoOk

`func (o *Branch) GetVcsRepoOk() (*string, bool)`

GetVcsRepoOk returns a tuple with the VcsRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsRepo

`func (o *Branch) SetVcsRepo(v string)`

SetVcsRepo sets VcsRepo field to given value.

### HasVcsRepo

`func (o *Branch) HasVcsRepo() bool`

HasVcsRepo returns a boolean if a field has been set.

### GetVcsType

`func (o *Branch) GetVcsType() string`

GetVcsType returns the VcsType field if non-nil, zero value otherwise.

### GetVcsTypeOk

`func (o *Branch) GetVcsTypeOk() (*string, bool)`

GetVcsTypeOk returns a tuple with the VcsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsType

`func (o *Branch) SetVcsType(v string)`

SetVcsType sets VcsType field to given value.

### HasVcsType

`func (o *Branch) HasVcsType() bool`

HasVcsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


