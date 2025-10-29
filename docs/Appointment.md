# Appointment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeSlotId** | Pointer to **string** | The unique identifier of the time slot this appointment is assigned to. | [optional] 
**Breaks** | Pointer to [**[]Break**](Break.md) | A list of breaks that are scheduled to be taken before the appointment. | [optional] [default to {}]
**WaitingDuration** | Pointer to **int32** | The duration [s] of the waiting period before the appointment can start. | [optional] 
**Start** | Pointer to **time.Time** | The point in time when the appointment starts. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**PreparationDuration** | Pointer to **int32** | The duration [s] of the preparation period before the first task can start. This is determined by the allocated time slot. | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) | The list of tasks that are scheduled to be executed within this appointment. Each task can either be a pickup or a delivery. For every order scheduled on the route, the route will contain a task describing the pickup of the order and a task describing the delivery of the order. When the order is a pickup order (resp. delivery order), its delivery task (resp. pickup task) will be scheduled at a depot. | [optional] 
**End** | Pointer to **time.Time** | The point in time when the appointment ends. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 

## Methods

### NewAppointment

`func NewAppointment() *Appointment`

NewAppointment instantiates a new Appointment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppointmentWithDefaults

`func NewAppointmentWithDefaults() *Appointment`

NewAppointmentWithDefaults instantiates a new Appointment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeSlotId

`func (o *Appointment) GetTimeSlotId() string`

GetTimeSlotId returns the TimeSlotId field if non-nil, zero value otherwise.

### GetTimeSlotIdOk

`func (o *Appointment) GetTimeSlotIdOk() (*string, bool)`

GetTimeSlotIdOk returns a tuple with the TimeSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotId

`func (o *Appointment) SetTimeSlotId(v string)`

SetTimeSlotId sets TimeSlotId field to given value.

### HasTimeSlotId

`func (o *Appointment) HasTimeSlotId() bool`

HasTimeSlotId returns a boolean if a field has been set.

### GetBreaks

`func (o *Appointment) GetBreaks() []Break`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *Appointment) GetBreaksOk() (*[]Break, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *Appointment) SetBreaks(v []Break)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *Appointment) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.

### GetWaitingDuration

`func (o *Appointment) GetWaitingDuration() int32`

GetWaitingDuration returns the WaitingDuration field if non-nil, zero value otherwise.

### GetWaitingDurationOk

`func (o *Appointment) GetWaitingDurationOk() (*int32, bool)`

GetWaitingDurationOk returns a tuple with the WaitingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingDuration

`func (o *Appointment) SetWaitingDuration(v int32)`

SetWaitingDuration sets WaitingDuration field to given value.

### HasWaitingDuration

`func (o *Appointment) HasWaitingDuration() bool`

HasWaitingDuration returns a boolean if a field has been set.

### GetStart

`func (o *Appointment) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Appointment) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Appointment) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Appointment) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetPreparationDuration

`func (o *Appointment) GetPreparationDuration() int32`

GetPreparationDuration returns the PreparationDuration field if non-nil, zero value otherwise.

### GetPreparationDurationOk

`func (o *Appointment) GetPreparationDurationOk() (*int32, bool)`

GetPreparationDurationOk returns a tuple with the PreparationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationDuration

`func (o *Appointment) SetPreparationDuration(v int32)`

SetPreparationDuration sets PreparationDuration field to given value.

### HasPreparationDuration

`func (o *Appointment) HasPreparationDuration() bool`

HasPreparationDuration returns a boolean if a field has been set.

### GetTasks

`func (o *Appointment) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Appointment) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Appointment) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Appointment) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetEnd

`func (o *Appointment) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Appointment) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Appointment) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Appointment) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


