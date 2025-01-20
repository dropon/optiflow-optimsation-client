/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TimeSlot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSlot{}

// TimeSlot Describes a possibility to schedule an appointment to execute tasks at a location.
type TimeSlot struct {
	// A unique identifier of the time slot. Must be unique within each location.
	Id string `json:"id" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	// The earliest point in time an appointment may start in this time slot. When omitted the appointment may start as early as desired. When used in conjunction with a latest start time, the earliest start time of a time slot must not be later than its latest start time. When used in conjunction with a latest end time, the earliest start time of a time slot must not be later than its latest end time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before `1970-01-01T00:00:00+00:00` nor after `2037-12-31T23:59:59+00:00`. The date must provide an offset to UTC.
	EarliestStart *time.Time `json:"earliestStart,omitempty"`
	// The latest point in time an appointment may start in this time slot. When omitted the appointment may start as late as desired. When used in conjunction with an earliest start time, the latest start time of a time slot must not be earlier than its earliest start time. When used in conjunction with a latest end time, the latest start time of a time slot must not be later than its latest end time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before `1970-01-01T00:00:00+00:00` nor after `2037-12-31T23:59:59+00:00`. The date must provide an offset to UTC.
	LatestStart *time.Time `json:"latestStart,omitempty"`
	// The latest point in time an appointment may end in this time slot. When omitted the appointment may end as late as desired. When used in conjunction with an earliest start time, the latest end time of a time slot must not be earlier than its earliest start time. When used in conjunction with an latest start time, the latest end time of a time slot must not be earlier than its latest start time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before `1970-01-01T00:00:00+00:00` nor after `2037-12-31T23:59:59+00:00`. The date must provide an offset to UTC.
	LatestEnd *time.Time `json:"latestEnd,omitempty"`
	// Describes how many appointments may be assigned to this time slot. When omitted, an unlimited number of appointments may be assigned to this time slot.
	MaximumAppointments *int32 `json:"maximumAppointments,omitempty"`
	// Describes the cost for assigning an appointment to this time slot.
	CostPerAppointment *float64 `json:"costPerAppointment,omitempty"`
	// Describes how long [s] it takes before the first task can be executed after starting the appointment.
	PreparationDuration *int32 `json:"preparationDuration,omitempty"`
}

type _TimeSlot TimeSlot

// NewTimeSlot instantiates a new TimeSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSlot(id string) *TimeSlot {
	this := TimeSlot{}
	this.Id = id
	var costPerAppointment float64 = 0
	this.CostPerAppointment = &costPerAppointment
	var preparationDuration int32 = 0
	this.PreparationDuration = &preparationDuration
	return &this
}

// NewTimeSlotWithDefaults instantiates a new TimeSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSlotWithDefaults() *TimeSlot {
	this := TimeSlot{}
	var costPerAppointment float64 = 0
	this.CostPerAppointment = &costPerAppointment
	var preparationDuration int32 = 0
	this.PreparationDuration = &preparationDuration
	return &this
}

// GetId returns the Id field value
func (o *TimeSlot) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TimeSlot) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TimeSlot) SetId(v string) {
	o.Id = v
}

// GetEarliestStart returns the EarliestStart field value if set, zero value otherwise.
func (o *TimeSlot) GetEarliestStart() time.Time {
	if o == nil || IsNil(o.EarliestStart) {
		var ret time.Time
		return ret
	}
	return *o.EarliestStart
}

// GetEarliestStartOk returns a tuple with the EarliestStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlot) GetEarliestStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EarliestStart) {
		return nil, false
	}
	return o.EarliestStart, true
}

// HasEarliestStart returns a boolean if a field has been set.
func (o *TimeSlot) HasEarliestStart() bool {
	if o != nil && !IsNil(o.EarliestStart) {
		return true
	}

	return false
}

// SetEarliestStart gets a reference to the given time.Time and assigns it to the EarliestStart field.
func (o *TimeSlot) SetEarliestStart(v time.Time) {
	o.EarliestStart = &v
}

// GetLatestStart returns the LatestStart field value if set, zero value otherwise.
func (o *TimeSlot) GetLatestStart() time.Time {
	if o == nil || IsNil(o.LatestStart) {
		var ret time.Time
		return ret
	}
	return *o.LatestStart
}

// GetLatestStartOk returns a tuple with the LatestStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlot) GetLatestStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LatestStart) {
		return nil, false
	}
	return o.LatestStart, true
}

// HasLatestStart returns a boolean if a field has been set.
func (o *TimeSlot) HasLatestStart() bool {
	if o != nil && !IsNil(o.LatestStart) {
		return true
	}

	return false
}

// SetLatestStart gets a reference to the given time.Time and assigns it to the LatestStart field.
func (o *TimeSlot) SetLatestStart(v time.Time) {
	o.LatestStart = &v
}

// GetLatestEnd returns the LatestEnd field value if set, zero value otherwise.
func (o *TimeSlot) GetLatestEnd() time.Time {
	if o == nil || IsNil(o.LatestEnd) {
		var ret time.Time
		return ret
	}
	return *o.LatestEnd
}

// GetLatestEndOk returns a tuple with the LatestEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlot) GetLatestEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LatestEnd) {
		return nil, false
	}
	return o.LatestEnd, true
}

// HasLatestEnd returns a boolean if a field has been set.
func (o *TimeSlot) HasLatestEnd() bool {
	if o != nil && !IsNil(o.LatestEnd) {
		return true
	}

	return false
}

// SetLatestEnd gets a reference to the given time.Time and assigns it to the LatestEnd field.
func (o *TimeSlot) SetLatestEnd(v time.Time) {
	o.LatestEnd = &v
}

// GetMaximumAppointments returns the MaximumAppointments field value if set, zero value otherwise.
func (o *TimeSlot) GetMaximumAppointments() int32 {
	if o == nil || IsNil(o.MaximumAppointments) {
		var ret int32
		return ret
	}
	return *o.MaximumAppointments
}

// GetMaximumAppointmentsOk returns a tuple with the MaximumAppointments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlot) GetMaximumAppointmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumAppointments) {
		return nil, false
	}
	return o.MaximumAppointments, true
}

// HasMaximumAppointments returns a boolean if a field has been set.
func (o *TimeSlot) HasMaximumAppointments() bool {
	if o != nil && !IsNil(o.MaximumAppointments) {
		return true
	}

	return false
}

// SetMaximumAppointments gets a reference to the given int32 and assigns it to the MaximumAppointments field.
func (o *TimeSlot) SetMaximumAppointments(v int32) {
	o.MaximumAppointments = &v
}

// GetCostPerAppointment returns the CostPerAppointment field value if set, zero value otherwise.
func (o *TimeSlot) GetCostPerAppointment() float64 {
	if o == nil || IsNil(o.CostPerAppointment) {
		var ret float64
		return ret
	}
	return *o.CostPerAppointment
}

// GetCostPerAppointmentOk returns a tuple with the CostPerAppointment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlot) GetCostPerAppointmentOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerAppointment) {
		return nil, false
	}
	return o.CostPerAppointment, true
}

// HasCostPerAppointment returns a boolean if a field has been set.
func (o *TimeSlot) HasCostPerAppointment() bool {
	if o != nil && !IsNil(o.CostPerAppointment) {
		return true
	}

	return false
}

// SetCostPerAppointment gets a reference to the given float64 and assigns it to the CostPerAppointment field.
func (o *TimeSlot) SetCostPerAppointment(v float64) {
	o.CostPerAppointment = &v
}

// GetPreparationDuration returns the PreparationDuration field value if set, zero value otherwise.
func (o *TimeSlot) GetPreparationDuration() int32 {
	if o == nil || IsNil(o.PreparationDuration) {
		var ret int32
		return ret
	}
	return *o.PreparationDuration
}

// GetPreparationDurationOk returns a tuple with the PreparationDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlot) GetPreparationDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.PreparationDuration) {
		return nil, false
	}
	return o.PreparationDuration, true
}

// HasPreparationDuration returns a boolean if a field has been set.
func (o *TimeSlot) HasPreparationDuration() bool {
	if o != nil && !IsNil(o.PreparationDuration) {
		return true
	}

	return false
}

// SetPreparationDuration gets a reference to the given int32 and assigns it to the PreparationDuration field.
func (o *TimeSlot) SetPreparationDuration(v int32) {
	o.PreparationDuration = &v
}

func (o TimeSlot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSlot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.EarliestStart) {
		toSerialize["earliestStart"] = o.EarliestStart
	}
	if !IsNil(o.LatestStart) {
		toSerialize["latestStart"] = o.LatestStart
	}
	if !IsNil(o.LatestEnd) {
		toSerialize["latestEnd"] = o.LatestEnd
	}
	if !IsNil(o.MaximumAppointments) {
		toSerialize["maximumAppointments"] = o.MaximumAppointments
	}
	if !IsNil(o.CostPerAppointment) {
		toSerialize["costPerAppointment"] = o.CostPerAppointment
	}
	if !IsNil(o.PreparationDuration) {
		toSerialize["preparationDuration"] = o.PreparationDuration
	}
	return toSerialize, nil
}

func (o *TimeSlot) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTimeSlot := _TimeSlot{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimeSlot)

	if err != nil {
		return err
	}

	*o = TimeSlot(varTimeSlot)

	return err
}

type NullableTimeSlot struct {
	value *TimeSlot
	isSet bool
}

func (v NullableTimeSlot) Get() *TimeSlot {
	return v.value
}

func (v *NullableTimeSlot) Set(val *TimeSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSlot(val *TimeSlot) *NullableTimeSlot {
	return &NullableTimeSlot{value: val, isSet: true}
}

func (v NullableTimeSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

