/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// OrderCompartmentCombinationConstraintType Describes which combination of orders and compartments have to be respected or excluded. * `ORDER_REQUIRES_COMPARTMENT` - An order from the given order category can only be loaded into a compartment from the given compartment category. * `COMPARTMENT_REQUIRES_ORDER` - A compartment from the given compartment category can only be used to load orders from the given order category. * `FORBIDDEN_COMBINATION` - An order from the given order category cannot be loaded into a compartment from the given compartment category.
type OrderCompartmentCombinationConstraintType string

// List of OrderCompartmentCombinationConstraintType
const (
	ORDER_REQUIRES_COMPARTMENT OrderCompartmentCombinationConstraintType = "ORDER_REQUIRES_COMPARTMENT"
	COMPARTMENT_REQUIRES_ORDER OrderCompartmentCombinationConstraintType = "COMPARTMENT_REQUIRES_ORDER"
	FORBIDDEN_COMBINATION OrderCompartmentCombinationConstraintType = "FORBIDDEN_COMBINATION"
)

// All allowed values of OrderCompartmentCombinationConstraintType enum
var AllowedOrderCompartmentCombinationConstraintTypeEnumValues = []OrderCompartmentCombinationConstraintType{
	"ORDER_REQUIRES_COMPARTMENT",
	"COMPARTMENT_REQUIRES_ORDER",
	"FORBIDDEN_COMBINATION",
}

func (v *OrderCompartmentCombinationConstraintType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCompartmentCombinationConstraintType(value)
	for _, existing := range AllowedOrderCompartmentCombinationConstraintTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCompartmentCombinationConstraintType", value)
}

// NewOrderCompartmentCombinationConstraintTypeFromValue returns a pointer to a valid OrderCompartmentCombinationConstraintType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCompartmentCombinationConstraintTypeFromValue(v string) (*OrderCompartmentCombinationConstraintType, error) {
	ev := OrderCompartmentCombinationConstraintType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCompartmentCombinationConstraintType: valid values are %v", v, AllowedOrderCompartmentCombinationConstraintTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCompartmentCombinationConstraintType) IsValid() bool {
	for _, existing := range AllowedOrderCompartmentCombinationConstraintTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderCompartmentCombinationConstraintType value
func (v OrderCompartmentCombinationConstraintType) Ptr() *OrderCompartmentCombinationConstraintType {
	return &v
}

type NullableOrderCompartmentCombinationConstraintType struct {
	value *OrderCompartmentCombinationConstraintType
	isSet bool
}

func (v NullableOrderCompartmentCombinationConstraintType) Get() *OrderCompartmentCombinationConstraintType {
	return v.value
}

func (v *NullableOrderCompartmentCombinationConstraintType) Set(val *OrderCompartmentCombinationConstraintType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCompartmentCombinationConstraintType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCompartmentCombinationConstraintType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCompartmentCombinationConstraintType(val *OrderCompartmentCombinationConstraintType) *NullableOrderCompartmentCombinationConstraintType {
	return &NullableOrderCompartmentCombinationConstraintType{value: val, isSet: true}
}

func (v NullableOrderCompartmentCombinationConstraintType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCompartmentCombinationConstraintType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

