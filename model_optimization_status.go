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

// OptimizationStatus Describes the current status of the optimization. * `PREPARING` - The optimization has been accepted. Time-distance data is being calculated and the data is being prepared for optimization. * `RUNNING` - The routes are being optimized. Metrics will provide insights into the current result of the optimization. * `STOPPING` - A request has been received to stop the optimization. The optimization will stop as soon as possible. * `FAILED` - The optimization has failed. An error will be provided to clarify what went wrong. * `SUCCEEDED` - The optimization has completed successfully.
type OptimizationStatus string

// List of OptimizationStatus
const (
	PREPARING OptimizationStatus = "PREPARING"
	RUNNING OptimizationStatus = "RUNNING"
	STOPPING OptimizationStatus = "STOPPING"
	FAILED OptimizationStatus = "FAILED"
	SUCCEEDED OptimizationStatus = "SUCCEEDED"
)

// All allowed values of OptimizationStatus enum
var AllowedOptimizationStatusEnumValues = []OptimizationStatus{
	"PREPARING",
	"RUNNING",
	"STOPPING",
	"FAILED",
	"SUCCEEDED",
}

func (v *OptimizationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OptimizationStatus(value)
	for _, existing := range AllowedOptimizationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OptimizationStatus", value)
}

// NewOptimizationStatusFromValue returns a pointer to a valid OptimizationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOptimizationStatusFromValue(v string) (*OptimizationStatus, error) {
	ev := OptimizationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OptimizationStatus: valid values are %v", v, AllowedOptimizationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OptimizationStatus) IsValid() bool {
	for _, existing := range AllowedOptimizationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OptimizationStatus value
func (v OptimizationStatus) Ptr() *OptimizationStatus {
	return &v
}

type NullableOptimizationStatus struct {
	value *OptimizationStatus
	isSet bool
}

func (v NullableOptimizationStatus) Get() *OptimizationStatus {
	return v.value
}

func (v *NullableOptimizationStatus) Set(val *OptimizationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationStatus(val *OptimizationStatus) *NullableOptimizationStatus {
	return &NullableOptimizationStatus{value: val, isSet: true}
}

func (v NullableOptimizationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptimizationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

