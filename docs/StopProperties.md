# StopProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreparationDuration** | Pointer to **int32** | Defines a duration [s] that is needed before an appointment (one or more tasks) can start at this location. This duration is needed once per stop whenever tasks are executed. | [optional] [default to 0]
**TimeSlots** | Pointer to [**[]TimeSlot**](TimeSlot.md) | A list of time intervals that describe when tasks can be executed at this location. Consecutive tasks with the same time slot are grouped to an appointment. The timings of the appointment must satisfy the restrictions of the time slot. When omitted or empty, all tasks within a stop at this location will be grouped into one appointment and the timings of this appointment are unrestricted. | [optional] [default to {}]
**Concurrency** | Pointer to [**StopConcurrency**](StopConcurrency.md) |  | [optional] 

## Methods

### NewStopProperties

`func NewStopProperties() *StopProperties`

NewStopProperties instantiates a new StopProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopPropertiesWithDefaults

`func NewStopPropertiesWithDefaults() *StopProperties`

NewStopPropertiesWithDefaults instantiates a new StopProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreparationDuration

`func (o *StopProperties) GetPreparationDuration() int32`

GetPreparationDuration returns the PreparationDuration field if non-nil, zero value otherwise.

### GetPreparationDurationOk

`func (o *StopProperties) GetPreparationDurationOk() (*int32, bool)`

GetPreparationDurationOk returns a tuple with the PreparationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationDuration

`func (o *StopProperties) SetPreparationDuration(v int32)`

SetPreparationDuration sets PreparationDuration field to given value.

### HasPreparationDuration

`func (o *StopProperties) HasPreparationDuration() bool`

HasPreparationDuration returns a boolean if a field has been set.

### GetTimeSlots

`func (o *StopProperties) GetTimeSlots() []TimeSlot`

GetTimeSlots returns the TimeSlots field if non-nil, zero value otherwise.

### GetTimeSlotsOk

`func (o *StopProperties) GetTimeSlotsOk() (*[]TimeSlot, bool)`

GetTimeSlotsOk returns a tuple with the TimeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlots

`func (o *StopProperties) SetTimeSlots(v []TimeSlot)`

SetTimeSlots sets TimeSlots field to given value.

### HasTimeSlots

`func (o *StopProperties) HasTimeSlots() bool`

HasTimeSlots returns a boolean if a field has been set.

### GetConcurrency

`func (o *StopProperties) GetConcurrency() StopConcurrency`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *StopProperties) GetConcurrencyOk() (*StopConcurrency, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *StopProperties) SetConcurrency(v StopConcurrency)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *StopProperties) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


