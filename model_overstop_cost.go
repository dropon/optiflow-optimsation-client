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

// checks if the OverstopCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverstopCost{}

// OverstopCost Specifies an increased cost per stop if the number of stops of a route exceeds a threshold.
type OverstopCost struct {
	// The threshold for the route's number of stops above which the extra cost per stop applies.
	Threshold int32 `json:"threshold"`
	// Specifies the extra cost for every stop above the threshold.
	ExtraPerStop float64 `json:"extraPerStop"`
}

type _OverstopCost OverstopCost

// NewOverstopCost instantiates a new OverstopCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverstopCost(threshold int32, extraPerStop float64) *OverstopCost {
	this := OverstopCost{}
	this.Threshold = threshold
	this.ExtraPerStop = extraPerStop
	return &this
}

// NewOverstopCostWithDefaults instantiates a new OverstopCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverstopCostWithDefaults() *OverstopCost {
	this := OverstopCost{}
	return &this
}

// GetThreshold returns the Threshold field value
func (o *OverstopCost) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *OverstopCost) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *OverstopCost) SetThreshold(v int32) {
	o.Threshold = v
}

// GetExtraPerStop returns the ExtraPerStop field value
func (o *OverstopCost) GetExtraPerStop() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ExtraPerStop
}

// GetExtraPerStopOk returns a tuple with the ExtraPerStop field value
// and a boolean to check if the value has been set.
func (o *OverstopCost) GetExtraPerStopOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraPerStop, true
}

// SetExtraPerStop sets field value
func (o *OverstopCost) SetExtraPerStop(v float64) {
	o.ExtraPerStop = v
}

func (o OverstopCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverstopCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threshold"] = o.Threshold
	toSerialize["extraPerStop"] = o.ExtraPerStop
	return toSerialize, nil
}

func (o *OverstopCost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threshold",
		"extraPerStop",
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

	varOverstopCost := _OverstopCost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverstopCost)

	if err != nil {
		return err
	}

	*o = OverstopCost(varOverstopCost)

	return err
}

type NullableOverstopCost struct {
	value *OverstopCost
	isSet bool
}

func (v NullableOverstopCost) Get() *OverstopCost {
	return v.value
}

func (v *NullableOverstopCost) Set(val *OverstopCost) {
	v.value = val
	v.isSet = true
}

func (v NullableOverstopCost) IsSet() bool {
	return v.isSet
}

func (v *NullableOverstopCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverstopCost(val *OverstopCost) *NullableOverstopCost {
	return &NullableOverstopCost{value: val, isSet: true}
}

func (v NullableOverstopCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverstopCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


