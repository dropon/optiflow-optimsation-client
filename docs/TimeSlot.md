# TimeSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the time slot. Must be unique within each location. | 
**EarliestStart** | Pointer to **time.Time** | The earliest point in time an appointment may start in this time slot. When omitted the appointment may start as early as desired. When used in conjunction with a latest start time, the earliest start time of a time slot must not be later than its latest start time. When used in conjunction with a latest end time, the earliest start time of a time slot must not be later than its latest end time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before &#x60;1970-01-01T00:00:00+00:00&#x60; nor after &#x60;2037-12-31T23:59:59+00:00&#x60;. The date must provide an offset to UTC. | [optional] 
**LatestStart** | Pointer to **time.Time** | The latest point in time an appointment may start in this time slot. When omitted the appointment may start as late as desired. When used in conjunction with an earliest start time, the latest start time of a time slot must not be earlier than its earliest start time. When used in conjunction with a latest end time, the latest start time of a time slot must not be later than its latest end time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before &#x60;1970-01-01T00:00:00+00:00&#x60; nor after &#x60;2037-12-31T23:59:59+00:00&#x60;. The date must provide an offset to UTC. | [optional] 
**LatestEnd** | Pointer to **time.Time** | The latest point in time an appointment may end in this time slot. When omitted the appointment may end as late as desired. When used in conjunction with an earliest start time, the latest end time of a time slot must not be earlier than its earliest start time. When used in conjunction with an latest start time, the latest end time of a time slot must not be earlier than its latest start time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before &#x60;1970-01-01T00:00:00+00:00&#x60; nor after &#x60;2037-12-31T23:59:59+00:00&#x60;. The date must provide an offset to UTC. | [optional] 
**MaximumAppointments** | Pointer to **int32** | Describes how many appointments may be assigned to this time slot. When omitted, an unlimited number of appointments may be assigned to this time slot. | [optional] 
**CostPerAppointment** | Pointer to **float64** | Describes the cost for assigning an appointment to this time slot. | [optional] [default to 0]
**PreparationDuration** | Pointer to **int32** | Describes how long [s] it takes before the first task can be executed after starting the appointment. | [optional] [default to 0]

## Methods

### NewTimeSlot

`func NewTimeSlot(id string, ) *TimeSlot`

NewTimeSlot instantiates a new TimeSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSlotWithDefaults

`func NewTimeSlotWithDefaults() *TimeSlot`

NewTimeSlotWithDefaults instantiates a new TimeSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeSlot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeSlot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeSlot) SetId(v string)`

SetId sets Id field to given value.


### GetEarliestStart

`func (o *TimeSlot) GetEarliestStart() time.Time`

GetEarliestStart returns the EarliestStart field if non-nil, zero value otherwise.

### GetEarliestStartOk

`func (o *TimeSlot) GetEarliestStartOk() (*time.Time, bool)`

GetEarliestStartOk returns a tuple with the EarliestStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestStart

`func (o *TimeSlot) SetEarliestStart(v time.Time)`

SetEarliestStart sets EarliestStart field to given value.

### HasEarliestStart

`func (o *TimeSlot) HasEarliestStart() bool`

HasEarliestStart returns a boolean if a field has been set.

### GetLatestStart

`func (o *TimeSlot) GetLatestStart() time.Time`

GetLatestStart returns the LatestStart field if non-nil, zero value otherwise.

### GetLatestStartOk

`func (o *TimeSlot) GetLatestStartOk() (*time.Time, bool)`

GetLatestStartOk returns a tuple with the LatestStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestStart

`func (o *TimeSlot) SetLatestStart(v time.Time)`

SetLatestStart sets LatestStart field to given value.

### HasLatestStart

`func (o *TimeSlot) HasLatestStart() bool`

HasLatestStart returns a boolean if a field has been set.

### GetLatestEnd

`func (o *TimeSlot) GetLatestEnd() time.Time`

GetLatestEnd returns the LatestEnd field if non-nil, zero value otherwise.

### GetLatestEndOk

`func (o *TimeSlot) GetLatestEndOk() (*time.Time, bool)`

GetLatestEndOk returns a tuple with the LatestEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEnd

`func (o *TimeSlot) SetLatestEnd(v time.Time)`

SetLatestEnd sets LatestEnd field to given value.

### HasLatestEnd

`func (o *TimeSlot) HasLatestEnd() bool`

HasLatestEnd returns a boolean if a field has been set.

### GetMaximumAppointments

`func (o *TimeSlot) GetMaximumAppointments() int32`

GetMaximumAppointments returns the MaximumAppointments field if non-nil, zero value otherwise.

### GetMaximumAppointmentsOk

`func (o *TimeSlot) GetMaximumAppointmentsOk() (*int32, bool)`

GetMaximumAppointmentsOk returns a tuple with the MaximumAppointments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAppointments

`func (o *TimeSlot) SetMaximumAppointments(v int32)`

SetMaximumAppointments sets MaximumAppointments field to given value.

### HasMaximumAppointments

`func (o *TimeSlot) HasMaximumAppointments() bool`

HasMaximumAppointments returns a boolean if a field has been set.

### GetCostPerAppointment

`func (o *TimeSlot) GetCostPerAppointment() float64`

GetCostPerAppointment returns the CostPerAppointment field if non-nil, zero value otherwise.

### GetCostPerAppointmentOk

`func (o *TimeSlot) GetCostPerAppointmentOk() (*float64, bool)`

GetCostPerAppointmentOk returns a tuple with the CostPerAppointment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerAppointment

`func (o *TimeSlot) SetCostPerAppointment(v float64)`

SetCostPerAppointment sets CostPerAppointment field to given value.

### HasCostPerAppointment

`func (o *TimeSlot) HasCostPerAppointment() bool`

HasCostPerAppointment returns a boolean if a field has been set.

### GetPreparationDuration

`func (o *TimeSlot) GetPreparationDuration() int32`

GetPreparationDuration returns the PreparationDuration field if non-nil, zero value otherwise.

### GetPreparationDurationOk

`func (o *TimeSlot) GetPreparationDurationOk() (*int32, bool)`

GetPreparationDurationOk returns a tuple with the PreparationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationDuration

`func (o *TimeSlot) SetPreparationDuration(v int32)`

SetPreparationDuration sets PreparationDuration field to given value.

### HasPreparationDuration

`func (o *TimeSlot) HasPreparationDuration() bool`

HasPreparationDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


