# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | The unique identifier of the order whose pickup or delivery is described by this task. | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 
**Breaks** | Pointer to [**[]Break**](Break.md) | A list of breaks that are scheduled to be taken before the execution of the task. | [optional] [default to []]
**Start** | Pointer to **time.Time** | The point in time when the execution of the task is scheduled to start. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**Duration** | Pointer to **int32** | The scheduled duration [s] for the task to be executed. This is the duration between the start and end of the task. | [optional] 
**End** | Pointer to **time.Time** | The point in time when the execution of the task is scheduled to end. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**CompartmentId** | Pointer to **string** | The unique identifier of the compartment that the order needs to be loaded in or unloaded from. | [optional] 
**DepotId** | Pointer to **string** | The unique identifier of the depot in case the task is a pickup or a delivery at a depot. | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *Task) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Task) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Task) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Task) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetType

`func (o *Task) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Task) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBreaks

`func (o *Task) GetBreaks() []Break`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *Task) GetBreaksOk() (*[]Break, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *Task) SetBreaks(v []Break)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *Task) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.

### GetStart

`func (o *Task) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Task) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Task) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Task) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetDuration

`func (o *Task) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Task) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Task) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Task) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnd

`func (o *Task) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Task) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Task) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Task) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetCompartmentId

`func (o *Task) GetCompartmentId() string`

GetCompartmentId returns the CompartmentId field if non-nil, zero value otherwise.

### GetCompartmentIdOk

`func (o *Task) GetCompartmentIdOk() (*string, bool)`

GetCompartmentIdOk returns a tuple with the CompartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentId

`func (o *Task) SetCompartmentId(v string)`

SetCompartmentId sets CompartmentId field to given value.

### HasCompartmentId

`func (o *Task) HasCompartmentId() bool`

HasCompartmentId returns a boolean if a field has been set.

### GetDepotId

`func (o *Task) GetDepotId() string`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *Task) GetDepotIdOk() (*string, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *Task) SetDepotId(v string)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *Task) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


