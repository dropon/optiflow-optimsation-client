# TaskGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskCategory** | **string** | The category that defines which tasks belong to this task group. The constraint will be ignored when no task belongs to this category. | 
**Constraint** | [**TaskGroupConstraint**](TaskGroupConstraint.md) |  | 

## Methods

### NewTaskGroup

`func NewTaskGroup(taskCategory string, constraint TaskGroupConstraint, ) *TaskGroup`

NewTaskGroup instantiates a new TaskGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskGroupWithDefaults

`func NewTaskGroupWithDefaults() *TaskGroup`

NewTaskGroupWithDefaults instantiates a new TaskGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskCategory

`func (o *TaskGroup) GetTaskCategory() string`

GetTaskCategory returns the TaskCategory field if non-nil, zero value otherwise.

### GetTaskCategoryOk

`func (o *TaskGroup) GetTaskCategoryOk() (*string, bool)`

GetTaskCategoryOk returns a tuple with the TaskCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategory

`func (o *TaskGroup) SetTaskCategory(v string)`

SetTaskCategory sets TaskCategory field to given value.


### GetConstraint

`func (o *TaskGroup) GetConstraint() TaskGroupConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *TaskGroup) GetConstraintOk() (*TaskGroupConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *TaskGroup) SetConstraint(v TaskGroupConstraint)`

SetConstraint sets Constraint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


