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

// checks if the Load type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Load{}

// Load For an order this describes how much capacity for a certain dimension is needed in the vehicle. For a vehicle, this describes its capacity for a certain dimension. For each dimension, the sum of the values of orders loaded into the vehicle must be lower than or equal to the value of the vehicle.
type Load struct {
	// Indicates the specific dimension of the load, such as its volume, weight, or size.
	Dimension string `json:"dimension" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	// Represents the numeric value associated with the load's dimension. This value could be the actual measurement or quantity of the load.
	Value float64 `json:"value"`
}

type _Load Load

// NewLoad instantiates a new Load object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoad(dimension string, value float64) *Load {
	this := Load{}
	this.Dimension = dimension
	this.Value = value
	return &this
}

// NewLoadWithDefaults instantiates a new Load object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadWithDefaults() *Load {
	this := Load{}
	return &this
}

// GetDimension returns the Dimension field value
func (o *Load) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *Load) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *Load) SetDimension(v string) {
	o.Dimension = v
}

// GetValue returns the Value field value
func (o *Load) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Load) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Load) SetValue(v float64) {
	o.Value = v
}

func (o Load) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Load) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dimension"] = o.Dimension
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *Load) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dimension",
		"value",
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

	varLoad := _Load{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoad)

	if err != nil {
		return err
	}

	*o = Load(varLoad)

	return err
}

type NullableLoad struct {
	value *Load
	isSet bool
}

func (v NullableLoad) Get() *Load {
	return v.value
}

func (v *NullableLoad) Set(val *Load) {
	v.value = val
	v.isSet = true
}

func (v NullableLoad) IsSet() bool {
	return v.isSet
}

func (v *NullableLoad) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoad(val *Load) *NullableLoad {
	return &NullableLoad{value: val, isSet: true}
}

func (v NullableLoad) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoad) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


