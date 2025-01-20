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

// checks if the OrderVehicleCombinationConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderVehicleCombinationConstraint{}

// OrderVehicleCombinationConstraint A constraint on the combination of orders and vehicles belonging to a certain category.
type OrderVehicleCombinationConstraint struct {
	Type OrderVehicleCombinationConstraintType `json:"type"`
	// The category of orders to which the constraint applies. The constraint will be ignored if no order belongs to this category.
	OrderCategory string `json:"orderCategory" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	// The category of vehicles to which the constraint applies. The constraint will be ignored if no vehicle belongs to this category.
	VehicleCategory string `json:"vehicleCategory" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	// The cost incurred when an order-vehicle combination does not meet this constraint. When omitted, all order-vehicle combinations must satisfy this constraint.
	ViolationCost *float64 `json:"violationCost,omitempty"`
}

type _OrderVehicleCombinationConstraint OrderVehicleCombinationConstraint

// NewOrderVehicleCombinationConstraint instantiates a new OrderVehicleCombinationConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderVehicleCombinationConstraint(type_ OrderVehicleCombinationConstraintType, orderCategory string, vehicleCategory string) *OrderVehicleCombinationConstraint {
	this := OrderVehicleCombinationConstraint{}
	this.Type = type_
	this.OrderCategory = orderCategory
	this.VehicleCategory = vehicleCategory
	return &this
}

// NewOrderVehicleCombinationConstraintWithDefaults instantiates a new OrderVehicleCombinationConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderVehicleCombinationConstraintWithDefaults() *OrderVehicleCombinationConstraint {
	this := OrderVehicleCombinationConstraint{}
	return &this
}

// GetType returns the Type field value
func (o *OrderVehicleCombinationConstraint) GetType() OrderVehicleCombinationConstraintType {
	if o == nil {
		var ret OrderVehicleCombinationConstraintType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderVehicleCombinationConstraint) GetTypeOk() (*OrderVehicleCombinationConstraintType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderVehicleCombinationConstraint) SetType(v OrderVehicleCombinationConstraintType) {
	o.Type = v
}

// GetOrderCategory returns the OrderCategory field value
func (o *OrderVehicleCombinationConstraint) GetOrderCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderCategory
}

// GetOrderCategoryOk returns a tuple with the OrderCategory field value
// and a boolean to check if the value has been set.
func (o *OrderVehicleCombinationConstraint) GetOrderCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderCategory, true
}

// SetOrderCategory sets field value
func (o *OrderVehicleCombinationConstraint) SetOrderCategory(v string) {
	o.OrderCategory = v
}

// GetVehicleCategory returns the VehicleCategory field value
func (o *OrderVehicleCombinationConstraint) GetVehicleCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VehicleCategory
}

// GetVehicleCategoryOk returns a tuple with the VehicleCategory field value
// and a boolean to check if the value has been set.
func (o *OrderVehicleCombinationConstraint) GetVehicleCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VehicleCategory, true
}

// SetVehicleCategory sets field value
func (o *OrderVehicleCombinationConstraint) SetVehicleCategory(v string) {
	o.VehicleCategory = v
}

// GetViolationCost returns the ViolationCost field value if set, zero value otherwise.
func (o *OrderVehicleCombinationConstraint) GetViolationCost() float64 {
	if o == nil || IsNil(o.ViolationCost) {
		var ret float64
		return ret
	}
	return *o.ViolationCost
}

// GetViolationCostOk returns a tuple with the ViolationCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVehicleCombinationConstraint) GetViolationCostOk() (*float64, bool) {
	if o == nil || IsNil(o.ViolationCost) {
		return nil, false
	}
	return o.ViolationCost, true
}

// HasViolationCost returns a boolean if a field has been set.
func (o *OrderVehicleCombinationConstraint) HasViolationCost() bool {
	if o != nil && !IsNil(o.ViolationCost) {
		return true
	}

	return false
}

// SetViolationCost gets a reference to the given float64 and assigns it to the ViolationCost field.
func (o *OrderVehicleCombinationConstraint) SetViolationCost(v float64) {
	o.ViolationCost = &v
}

func (o OrderVehicleCombinationConstraint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderVehicleCombinationConstraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["orderCategory"] = o.OrderCategory
	toSerialize["vehicleCategory"] = o.VehicleCategory
	if !IsNil(o.ViolationCost) {
		toSerialize["violationCost"] = o.ViolationCost
	}
	return toSerialize, nil
}

func (o *OrderVehicleCombinationConstraint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"orderCategory",
		"vehicleCategory",
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

	varOrderVehicleCombinationConstraint := _OrderVehicleCombinationConstraint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderVehicleCombinationConstraint)

	if err != nil {
		return err
	}

	*o = OrderVehicleCombinationConstraint(varOrderVehicleCombinationConstraint)

	return err
}

type NullableOrderVehicleCombinationConstraint struct {
	value *OrderVehicleCombinationConstraint
	isSet bool
}

func (v NullableOrderVehicleCombinationConstraint) Get() *OrderVehicleCombinationConstraint {
	return v.value
}

func (v *NullableOrderVehicleCombinationConstraint) Set(val *OrderVehicleCombinationConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderVehicleCombinationConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderVehicleCombinationConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderVehicleCombinationConstraint(val *OrderVehicleCombinationConstraint) *NullableOrderVehicleCombinationConstraint {
	return &NullableOrderVehicleCombinationConstraint{value: val, isSet: true}
}

func (v NullableOrderVehicleCombinationConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderVehicleCombinationConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


