# Constraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Combinations** | Pointer to [**CombinationConstraints**](CombinationConstraints.md) |  | [optional] 
**Orders** | Pointer to [**OrderConstraints**](OrderConstraints.md) |  | [optional] 
**Tasks** | Pointer to [**TaskConstraints**](TaskConstraints.md) |  | [optional] 

## Methods

### NewConstraints

`func NewConstraints() *Constraints`

NewConstraints instantiates a new Constraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintsWithDefaults

`func NewConstraintsWithDefaults() *Constraints`

NewConstraintsWithDefaults instantiates a new Constraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinations

`func (o *Constraints) GetCombinations() CombinationConstraints`

GetCombinations returns the Combinations field if non-nil, zero value otherwise.

### GetCombinationsOk

`func (o *Constraints) GetCombinationsOk() (*CombinationConstraints, bool)`

GetCombinationsOk returns a tuple with the Combinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinations

`func (o *Constraints) SetCombinations(v CombinationConstraints)`

SetCombinations sets Combinations field to given value.

### HasCombinations

`func (o *Constraints) HasCombinations() bool`

HasCombinations returns a boolean if a field has been set.

### GetOrders

`func (o *Constraints) GetOrders() OrderConstraints`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *Constraints) GetOrdersOk() (*OrderConstraints, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *Constraints) SetOrders(v OrderConstraints)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *Constraints) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetTasks

`func (o *Constraints) GetTasks() TaskConstraints`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Constraints) GetTasksOk() (*TaskConstraints, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Constraints) SetTasks(v TaskConstraints)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Constraints) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


