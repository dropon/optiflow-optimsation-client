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

// checks if the OrderConstraints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderConstraints{}

// OrderConstraints Describes constraints on the way orders can be scheduled together on routes.
type OrderConstraints struct {
	// A list of sequences that must be respected when scheduling routes. Orders belonging to a category that occurs earlier in the sequence must be delivered in the route before an order belonging to a category later in the sequence can be picked up.
	RespectedSequences []RespectedOrderSequence `json:"respectedSequences,omitempty"`
	// A list of constraints that prevent orders to be loaded or unloaded while other orders are loaded in the vehicle.
	LoadingIncompatibilities []LoadingIncompatibilityConstraint `json:"loadingIncompatibilities,omitempty"`
}

// NewOrderConstraints instantiates a new OrderConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderConstraints() *OrderConstraints {
	this := OrderConstraints{}
	return &this
}

// NewOrderConstraintsWithDefaults instantiates a new OrderConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderConstraintsWithDefaults() *OrderConstraints {
	this := OrderConstraints{}
	return &this
}

// GetRespectedSequences returns the RespectedSequences field value if set, zero value otherwise.
func (o *OrderConstraints) GetRespectedSequences() []RespectedOrderSequence {
	if o == nil || IsNil(o.RespectedSequences) {
		var ret []RespectedOrderSequence
		return ret
	}
	return o.RespectedSequences
}

// GetRespectedSequencesOk returns a tuple with the RespectedSequences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderConstraints) GetRespectedSequencesOk() ([]RespectedOrderSequence, bool) {
	if o == nil || IsNil(o.RespectedSequences) {
		return nil, false
	}
	return o.RespectedSequences, true
}

// HasRespectedSequences returns a boolean if a field has been set.
func (o *OrderConstraints) HasRespectedSequences() bool {
	if o != nil && !IsNil(o.RespectedSequences) {
		return true
	}

	return false
}

// SetRespectedSequences gets a reference to the given []RespectedOrderSequence and assigns it to the RespectedSequences field.
func (o *OrderConstraints) SetRespectedSequences(v []RespectedOrderSequence) {
	o.RespectedSequences = v
}

// GetLoadingIncompatibilities returns the LoadingIncompatibilities field value if set, zero value otherwise.
func (o *OrderConstraints) GetLoadingIncompatibilities() []LoadingIncompatibilityConstraint {
	if o == nil || IsNil(o.LoadingIncompatibilities) {
		var ret []LoadingIncompatibilityConstraint
		return ret
	}
	return o.LoadingIncompatibilities
}

// GetLoadingIncompatibilitiesOk returns a tuple with the LoadingIncompatibilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderConstraints) GetLoadingIncompatibilitiesOk() ([]LoadingIncompatibilityConstraint, bool) {
	if o == nil || IsNil(o.LoadingIncompatibilities) {
		return nil, false
	}
	return o.LoadingIncompatibilities, true
}

// HasLoadingIncompatibilities returns a boolean if a field has been set.
func (o *OrderConstraints) HasLoadingIncompatibilities() bool {
	if o != nil && !IsNil(o.LoadingIncompatibilities) {
		return true
	}

	return false
}

// SetLoadingIncompatibilities gets a reference to the given []LoadingIncompatibilityConstraint and assigns it to the LoadingIncompatibilities field.
func (o *OrderConstraints) SetLoadingIncompatibilities(v []LoadingIncompatibilityConstraint) {
	o.LoadingIncompatibilities = v
}

func (o OrderConstraints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderConstraints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RespectedSequences) {
		toSerialize["respectedSequences"] = o.RespectedSequences
	}
	if !IsNil(o.LoadingIncompatibilities) {
		toSerialize["loadingIncompatibilities"] = o.LoadingIncompatibilities
	}
	return toSerialize, nil
}

type NullableOrderConstraints struct {
	value *OrderConstraints
	isSet bool
}

func (v NullableOrderConstraints) Get() *OrderConstraints {
	return v.value
}

func (v *NullableOrderConstraints) Set(val *OrderConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderConstraints(val *OrderConstraints) *NullableOrderConstraints {
	return &NullableOrderConstraints{value: val, isSet: true}
}

func (v NullableOrderConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

