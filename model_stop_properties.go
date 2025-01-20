/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StopProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopProperties{}

// StopProperties Specifies the necessary information needed to schedule a stop at this location. This information is not relevant when the location is the start or end location of a vehicle.
type StopProperties struct {
	// Defines a duration [s] that is needed before an appointment (one or more tasks) can start at this location. This duration is needed once per stop whenever tasks are executed.
	PreparationDuration *int32 `json:"preparationDuration,omitempty"`
	// A list of time intervals that describe when tasks can be executed at this location. Consecutive tasks with the same time slot are grouped to an appointment. The timings of the appointment must satisfy the restrictions of the time slot. When omitted or empty, all tasks within a stop at this location will be grouped into one appointment and the timings of this appointment are unrestricted.
	TimeSlots []TimeSlot `json:"timeSlots,omitempty"`
	Concurrency *StopConcurrency `json:"concurrency,omitempty"`
}

// NewStopProperties instantiates a new StopProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopProperties() *StopProperties {
	this := StopProperties{}
	var preparationDuration int32 = 0
	this.PreparationDuration = &preparationDuration
	return &this
}

// NewStopPropertiesWithDefaults instantiates a new StopProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopPropertiesWithDefaults() *StopProperties {
	this := StopProperties{}
	var preparationDuration int32 = 0
	this.PreparationDuration = &preparationDuration
	return &this
}

// GetPreparationDuration returns the PreparationDuration field value if set, zero value otherwise.
func (o *StopProperties) GetPreparationDuration() int32 {
	if o == nil || IsNil(o.PreparationDuration) {
		var ret int32
		return ret
	}
	return *o.PreparationDuration
}

// GetPreparationDurationOk returns a tuple with the PreparationDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopProperties) GetPreparationDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.PreparationDuration) {
		return nil, false
	}
	return o.PreparationDuration, true
}

// HasPreparationDuration returns a boolean if a field has been set.
func (o *StopProperties) HasPreparationDuration() bool {
	if o != nil && !IsNil(o.PreparationDuration) {
		return true
	}

	return false
}

// SetPreparationDuration gets a reference to the given int32 and assigns it to the PreparationDuration field.
func (o *StopProperties) SetPreparationDuration(v int32) {
	o.PreparationDuration = &v
}

// GetTimeSlots returns the TimeSlots field value if set, zero value otherwise.
func (o *StopProperties) GetTimeSlots() []TimeSlot {
	if o == nil || IsNil(o.TimeSlots) {
		var ret []TimeSlot
		return ret
	}
	return o.TimeSlots
}

// GetTimeSlotsOk returns a tuple with the TimeSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopProperties) GetTimeSlotsOk() ([]TimeSlot, bool) {
	if o == nil || IsNil(o.TimeSlots) {
		return nil, false
	}
	return o.TimeSlots, true
}

// HasTimeSlots returns a boolean if a field has been set.
func (o *StopProperties) HasTimeSlots() bool {
	if o != nil && !IsNil(o.TimeSlots) {
		return true
	}

	return false
}

// SetTimeSlots gets a reference to the given []TimeSlot and assigns it to the TimeSlots field.
func (o *StopProperties) SetTimeSlots(v []TimeSlot) {
	o.TimeSlots = v
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise.
func (o *StopProperties) GetConcurrency() StopConcurrency {
	if o == nil || IsNil(o.Concurrency) {
		var ret StopConcurrency
		return ret
	}
	return *o.Concurrency
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopProperties) GetConcurrencyOk() (*StopConcurrency, bool) {
	if o == nil || IsNil(o.Concurrency) {
		return nil, false
	}
	return o.Concurrency, true
}

// HasConcurrency returns a boolean if a field has been set.
func (o *StopProperties) HasConcurrency() bool {
	if o != nil && !IsNil(o.Concurrency) {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given StopConcurrency and assigns it to the Concurrency field.
func (o *StopProperties) SetConcurrency(v StopConcurrency) {
	o.Concurrency = &v
}

func (o StopProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreparationDuration) {
		toSerialize["preparationDuration"] = o.PreparationDuration
	}
	if !IsNil(o.TimeSlots) {
		toSerialize["timeSlots"] = o.TimeSlots
	}
	if !IsNil(o.Concurrency) {
		toSerialize["concurrency"] = o.Concurrency
	}
	return toSerialize, nil
}

type NullableStopProperties struct {
	value *StopProperties
	isSet bool
}

func (v NullableStopProperties) Get() *StopProperties {
	return v.value
}

func (v *NullableStopProperties) Set(val *StopProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableStopProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableStopProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopProperties(val *StopProperties) *NullableStopProperties {
	return &NullableStopProperties{value: val, isSet: true}
}

func (v NullableStopProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


