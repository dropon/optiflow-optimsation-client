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

// checks if the DepotVehicleCombinationConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepotVehicleCombinationConstraint{}

// DepotVehicleCombinationConstraint A constraint on the combination of depots and vehicles belonging to a certain category.
type DepotVehicleCombinationConstraint struct {
	Type DepotVehicleCombinationConstraintType `json:"type"`
	// The category of depots to which the constraint applies. The constraint will be ignored when no depot belongs to this category.
	DepotCategory string `json:"depotCategory" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	// The category of vehicles to which the constraint applies. The constraint will be ignored when no vehicle belongs to this category.
	VehicleCategory string `json:"vehicleCategory" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
}

type _DepotVehicleCombinationConstraint DepotVehicleCombinationConstraint

// NewDepotVehicleCombinationConstraint instantiates a new DepotVehicleCombinationConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepotVehicleCombinationConstraint(type_ DepotVehicleCombinationConstraintType, depotCategory string, vehicleCategory string) *DepotVehicleCombinationConstraint {
	this := DepotVehicleCombinationConstraint{}
	this.Type = type_
	this.DepotCategory = depotCategory
	this.VehicleCategory = vehicleCategory
	return &this
}

// NewDepotVehicleCombinationConstraintWithDefaults instantiates a new DepotVehicleCombinationConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepotVehicleCombinationConstraintWithDefaults() *DepotVehicleCombinationConstraint {
	this := DepotVehicleCombinationConstraint{}
	return &this
}

// GetType returns the Type field value
func (o *DepotVehicleCombinationConstraint) GetType() DepotVehicleCombinationConstraintType {
	if o == nil {
		var ret DepotVehicleCombinationConstraintType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DepotVehicleCombinationConstraint) GetTypeOk() (*DepotVehicleCombinationConstraintType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DepotVehicleCombinationConstraint) SetType(v DepotVehicleCombinationConstraintType) {
	o.Type = v
}

// GetDepotCategory returns the DepotCategory field value
func (o *DepotVehicleCombinationConstraint) GetDepotCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotCategory
}

// GetDepotCategoryOk returns a tuple with the DepotCategory field value
// and a boolean to check if the value has been set.
func (o *DepotVehicleCombinationConstraint) GetDepotCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotCategory, true
}

// SetDepotCategory sets field value
func (o *DepotVehicleCombinationConstraint) SetDepotCategory(v string) {
	o.DepotCategory = v
}

// GetVehicleCategory returns the VehicleCategory field value
func (o *DepotVehicleCombinationConstraint) GetVehicleCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VehicleCategory
}

// GetVehicleCategoryOk returns a tuple with the VehicleCategory field value
// and a boolean to check if the value has been set.
func (o *DepotVehicleCombinationConstraint) GetVehicleCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VehicleCategory, true
}

// SetVehicleCategory sets field value
func (o *DepotVehicleCombinationConstraint) SetVehicleCategory(v string) {
	o.VehicleCategory = v
}

func (o DepotVehicleCombinationConstraint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepotVehicleCombinationConstraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["depotCategory"] = o.DepotCategory
	toSerialize["vehicleCategory"] = o.VehicleCategory
	return toSerialize, nil
}

func (o *DepotVehicleCombinationConstraint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"depotCategory",
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

	varDepotVehicleCombinationConstraint := _DepotVehicleCombinationConstraint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDepotVehicleCombinationConstraint)

	if err != nil {
		return err
	}

	*o = DepotVehicleCombinationConstraint(varDepotVehicleCombinationConstraint)

	return err
}

type NullableDepotVehicleCombinationConstraint struct {
	value *DepotVehicleCombinationConstraint
	isSet bool
}

func (v NullableDepotVehicleCombinationConstraint) Get() *DepotVehicleCombinationConstraint {
	return v.value
}

func (v *NullableDepotVehicleCombinationConstraint) Set(val *DepotVehicleCombinationConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableDepotVehicleCombinationConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableDepotVehicleCombinationConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepotVehicleCombinationConstraint(val *DepotVehicleCombinationConstraint) *NullableDepotVehicleCombinationConstraint {
	return &NullableDepotVehicleCombinationConstraint{value: val, isSet: true}
}

func (v NullableDepotVehicleCombinationConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepotVehicleCombinationConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

