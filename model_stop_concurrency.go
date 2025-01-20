/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the StopConcurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopConcurrency{}

// StopConcurrency If stop concurrency is specified, each stop at the location must be assigned to a vehicle slot. Two stops may be assigned to the same vehicle slot if the duration between the end of the last appointment of the first stop and the start of the first appointment of the other stop is at least the specified minimum buffer duration, in which case they are not considered concurrent stops. Assigning at least one stop to a slot incurs a cost, encouraging reduction of the maximum number of concurrent stops at the location. If there is no minimum buffer duration, stops where the first appointment starts at the same time as the last ends do not need to be assigned to a vehicle slot. When omitted, stops at this location will not be assigned to a vehicle slot.
type StopConcurrency struct {
	// A list of available vehicle slots.
	VehicleSlots []VehicleSlot `json:"vehicleSlots,omitempty"`
	// The additional cost incurred for using an extra vehicle slot beyond the available ones. This must be greater than or equal to the cost of each vehicle slot.
	ViolationCostPerExtraSlot float64 `json:"violationCostPerExtraSlot"`
	// The minimum duration [s] between the end of the last appointment of a stop and the start of the first appointment of another stop assigned to the same vehicle slot.
	MinimumBufferDuration *int32 `json:"minimumBufferDuration,omitempty"`
}

type _StopConcurrency StopConcurrency

// NewStopConcurrency instantiates a new StopConcurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopConcurrency(violationCostPerExtraSlot float64) *StopConcurrency {
	this := StopConcurrency{}
	this.ViolationCostPerExtraSlot = violationCostPerExtraSlot
	var minimumBufferDuration int32 = 0
	this.MinimumBufferDuration = &minimumBufferDuration
	return &this
}

// NewStopConcurrencyWithDefaults instantiates a new StopConcurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopConcurrencyWithDefaults() *StopConcurrency {
	this := StopConcurrency{}
	var minimumBufferDuration int32 = 0
	this.MinimumBufferDuration = &minimumBufferDuration
	return &this
}

// GetVehicleSlots returns the VehicleSlots field value if set, zero value otherwise.
func (o *StopConcurrency) GetVehicleSlots() []VehicleSlot {
	if o == nil || IsNil(o.VehicleSlots) {
		var ret []VehicleSlot
		return ret
	}
	return o.VehicleSlots
}

// GetVehicleSlotsOk returns a tuple with the VehicleSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopConcurrency) GetVehicleSlotsOk() ([]VehicleSlot, bool) {
	if o == nil || IsNil(o.VehicleSlots) {
		return nil, false
	}
	return o.VehicleSlots, true
}

// HasVehicleSlots returns a boolean if a field has been set.
func (o *StopConcurrency) HasVehicleSlots() bool {
	if o != nil && !IsNil(o.VehicleSlots) {
		return true
	}

	return false
}

// SetVehicleSlots gets a reference to the given []VehicleSlot and assigns it to the VehicleSlots field.
func (o *StopConcurrency) SetVehicleSlots(v []VehicleSlot) {
	o.VehicleSlots = v
}

// GetViolationCostPerExtraSlot returns the ViolationCostPerExtraSlot field value
func (o *StopConcurrency) GetViolationCostPerExtraSlot() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ViolationCostPerExtraSlot
}

// GetViolationCostPerExtraSlotOk returns a tuple with the ViolationCostPerExtraSlot field value
// and a boolean to check if the value has been set.
func (o *StopConcurrency) GetViolationCostPerExtraSlotOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViolationCostPerExtraSlot, true
}

// SetViolationCostPerExtraSlot sets field value
func (o *StopConcurrency) SetViolationCostPerExtraSlot(v float64) {
	o.ViolationCostPerExtraSlot = v
}

// GetMinimumBufferDuration returns the MinimumBufferDuration field value if set, zero value otherwise.
func (o *StopConcurrency) GetMinimumBufferDuration() int32 {
	if o == nil || IsNil(o.MinimumBufferDuration) {
		var ret int32
		return ret
	}
	return *o.MinimumBufferDuration
}

// GetMinimumBufferDurationOk returns a tuple with the MinimumBufferDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopConcurrency) GetMinimumBufferDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumBufferDuration) {
		return nil, false
	}
	return o.MinimumBufferDuration, true
}

// HasMinimumBufferDuration returns a boolean if a field has been set.
func (o *StopConcurrency) HasMinimumBufferDuration() bool {
	if o != nil && !IsNil(o.MinimumBufferDuration) {
		return true
	}

	return false
}

// SetMinimumBufferDuration gets a reference to the given int32 and assigns it to the MinimumBufferDuration field.
func (o *StopConcurrency) SetMinimumBufferDuration(v int32) {
	o.MinimumBufferDuration = &v
}

func (o StopConcurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopConcurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VehicleSlots) {
		toSerialize["vehicleSlots"] = o.VehicleSlots
	}
	toSerialize["violationCostPerExtraSlot"] = o.ViolationCostPerExtraSlot
	if !IsNil(o.MinimumBufferDuration) {
		toSerialize["minimumBufferDuration"] = o.MinimumBufferDuration
	}
	return toSerialize, nil
}

func (o *StopConcurrency) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"violationCostPerExtraSlot",
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

	varStopConcurrency := _StopConcurrency{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStopConcurrency)

	if err != nil {
		return err
	}

	*o = StopConcurrency(varStopConcurrency)

	return err
}

type NullableStopConcurrency struct {
	value *StopConcurrency
	isSet bool
}

func (v NullableStopConcurrency) Get() *StopConcurrency {
	return v.value
}

func (v *NullableStopConcurrency) Set(val *StopConcurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableStopConcurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableStopConcurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopConcurrency(val *StopConcurrency) *NullableStopConcurrency {
	return &NullableStopConcurrency{value: val, isSet: true}
}

func (v NullableStopConcurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopConcurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

