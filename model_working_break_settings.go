/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WorkingBreakSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkingBreakSettings{}

// WorkingBreakSettings Restricts how long the driver of the vehicle may work without taking a break of at least the specified duration. All the time spent during a route is considered working time except for waiting time and break time.
type WorkingBreakSettings struct {
	// Describes how long [s] the driver may work without taking a break of at least the specified duration.
	MaximumWorkingDuration int32 `json:"maximumWorkingDuration"`
	// Specifies the duration [s] of a break a driver must take if they exceed the maximum working duration.
	MinimumBreakDuration int32 `json:"minimumBreakDuration"`
}

type _WorkingBreakSettings WorkingBreakSettings

// NewWorkingBreakSettings instantiates a new WorkingBreakSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkingBreakSettings(maximumWorkingDuration int32, minimumBreakDuration int32) *WorkingBreakSettings {
	this := WorkingBreakSettings{}
	this.MaximumWorkingDuration = maximumWorkingDuration
	this.MinimumBreakDuration = minimumBreakDuration
	return &this
}

// NewWorkingBreakSettingsWithDefaults instantiates a new WorkingBreakSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkingBreakSettingsWithDefaults() *WorkingBreakSettings {
	this := WorkingBreakSettings{}
	return &this
}

// GetMaximumWorkingDuration returns the MaximumWorkingDuration field value
func (o *WorkingBreakSettings) GetMaximumWorkingDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumWorkingDuration
}

// GetMaximumWorkingDurationOk returns a tuple with the MaximumWorkingDuration field value
// and a boolean to check if the value has been set.
func (o *WorkingBreakSettings) GetMaximumWorkingDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumWorkingDuration, true
}

// SetMaximumWorkingDuration sets field value
func (o *WorkingBreakSettings) SetMaximumWorkingDuration(v int32) {
	o.MaximumWorkingDuration = v
}

// GetMinimumBreakDuration returns the MinimumBreakDuration field value
func (o *WorkingBreakSettings) GetMinimumBreakDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinimumBreakDuration
}

// GetMinimumBreakDurationOk returns a tuple with the MinimumBreakDuration field value
// and a boolean to check if the value has been set.
func (o *WorkingBreakSettings) GetMinimumBreakDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumBreakDuration, true
}

// SetMinimumBreakDuration sets field value
func (o *WorkingBreakSettings) SetMinimumBreakDuration(v int32) {
	o.MinimumBreakDuration = v
}

func (o WorkingBreakSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkingBreakSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maximumWorkingDuration"] = o.MaximumWorkingDuration
	toSerialize["minimumBreakDuration"] = o.MinimumBreakDuration
	return toSerialize, nil
}

func (o *WorkingBreakSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"maximumWorkingDuration",
		"minimumBreakDuration",
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

	varWorkingBreakSettings := _WorkingBreakSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkingBreakSettings)

	if err != nil {
		return err
	}

	*o = WorkingBreakSettings(varWorkingBreakSettings)

	return err
}

type NullableWorkingBreakSettings struct {
	value *WorkingBreakSettings
	isSet bool
}

func (v NullableWorkingBreakSettings) Get() *WorkingBreakSettings {
	return v.value
}

func (v *NullableWorkingBreakSettings) Set(val *WorkingBreakSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkingBreakSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkingBreakSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkingBreakSettings(val *WorkingBreakSettings) *NullableWorkingBreakSettings {
	return &NullableWorkingBreakSettings{value: val, isSet: true}
}

func (v NullableWorkingBreakSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkingBreakSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


