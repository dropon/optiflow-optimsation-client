/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CompartmentLoadingStrategy Describes the strategy for loading and unloading orders in the compartment. * `NONE` - There is no restriction on the sequence of loading or unloading. * `LAST_IN_FIRST_OUT` - The last order loaded must be the first to be unloaded.
type CompartmentLoadingStrategy string

// List of CompartmentLoadingStrategy
const (
	NONE CompartmentLoadingStrategy = "NONE"
	LAST_IN_FIRST_OUT CompartmentLoadingStrategy = "LAST_IN_FIRST_OUT"
)

// All allowed values of CompartmentLoadingStrategy enum
var AllowedCompartmentLoadingStrategyEnumValues = []CompartmentLoadingStrategy{
	"NONE",
	"LAST_IN_FIRST_OUT",
}

func (v *CompartmentLoadingStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompartmentLoadingStrategy(value)
	for _, existing := range AllowedCompartmentLoadingStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompartmentLoadingStrategy", value)
}

// NewCompartmentLoadingStrategyFromValue returns a pointer to a valid CompartmentLoadingStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompartmentLoadingStrategyFromValue(v string) (*CompartmentLoadingStrategy, error) {
	ev := CompartmentLoadingStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompartmentLoadingStrategy: valid values are %v", v, AllowedCompartmentLoadingStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompartmentLoadingStrategy) IsValid() bool {
	for _, existing := range AllowedCompartmentLoadingStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompartmentLoadingStrategy value
func (v CompartmentLoadingStrategy) Ptr() *CompartmentLoadingStrategy {
	return &v
}

type NullableCompartmentLoadingStrategy struct {
	value *CompartmentLoadingStrategy
	isSet bool
}

func (v NullableCompartmentLoadingStrategy) Get() *CompartmentLoadingStrategy {
	return v.value
}

func (v *NullableCompartmentLoadingStrategy) Set(val *CompartmentLoadingStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableCompartmentLoadingStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableCompartmentLoadingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompartmentLoadingStrategy(val *CompartmentLoadingStrategy) *NullableCompartmentLoadingStrategy {
	return &NullableCompartmentLoadingStrategy{value: val, isSet: true}
}

func (v NullableCompartmentLoadingStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompartmentLoadingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

