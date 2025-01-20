# TaskConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]TaskGroup**](TaskGroup.md) | A list of task groups. Tasks belonging to the same task group must be planned on the same route, or consecutively if they are on the same route, depending on the constraint defined in the task group. | [optional] [default to []]
**RespectedSequences** | Pointer to [**[]RespectedTaskSequence**](RespectedTaskSequence.md) | A list of sequences that must be respected when scheduling routes. Tasks belonging to a category that occurs earlier in the sequence must be scheduled on the route before a task belonging to a category later in the sequence. | [optional] [default to []]
**ForbiddenSequences** | Pointer to [**[]ForbiddenTaskSequence**](ForbiddenTaskSequence.md) | A list of sequences that are forbidden to be scheduled on a route. Tasks belonging to certain categories must not be scheduled before, or immediately before, tasks belonging to another specific category. | [optional] [default to []]

## Methods

### NewTaskConstraints

`func NewTaskConstraints() *TaskConstraints`

NewTaskConstraints instantiates a new TaskConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskConstraintsWithDefaults

`func NewTaskConstraintsWithDefaults() *TaskConstraints`

NewTaskConstraintsWithDefaults instantiates a new TaskConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *TaskConstraints) GetGroups() []TaskGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TaskConstraints) GetGroupsOk() (*[]TaskGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TaskConstraints) SetGroups(v []TaskGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *TaskConstraints) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetRespectedSequences

`func (o *TaskConstraints) GetRespectedSequences() []RespectedTaskSequence`

GetRespectedSequences returns the RespectedSequences field if non-nil, zero value otherwise.

### GetRespectedSequencesOk

`func (o *TaskConstraints) GetRespectedSequencesOk() (*[]RespectedTaskSequence, bool)`

GetRespectedSequencesOk returns a tuple with the RespectedSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectedSequences

`func (o *TaskConstraints) SetRespectedSequences(v []RespectedTaskSequence)`

SetRespectedSequences sets RespectedSequences field to given value.

### HasRespectedSequences

`func (o *TaskConstraints) HasRespectedSequences() bool`

HasRespectedSequences returns a boolean if a field has been set.

### GetForbiddenSequences

`func (o *TaskConstraints) GetForbiddenSequences() []ForbiddenTaskSequence`

GetForbiddenSequences returns the ForbiddenSequences field if non-nil, zero value otherwise.

### GetForbiddenSequencesOk

`func (o *TaskConstraints) GetForbiddenSequencesOk() (*[]ForbiddenTaskSequence, bool)`

GetForbiddenSequencesOk returns a tuple with the ForbiddenSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenSequences

`func (o *TaskConstraints) SetForbiddenSequences(v []ForbiddenTaskSequence)`

SetForbiddenSequences sets ForbiddenSequences field to given value.

### HasForbiddenSequences

`func (o *TaskConstraints) HasForbiddenSequences() bool`

HasForbiddenSequences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


