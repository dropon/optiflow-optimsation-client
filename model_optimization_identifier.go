/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OptimizationIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationIdentifier{}

// OptimizationIdentifier A reference to the optimization.
type OptimizationIdentifier struct {
	// The unique identifier of the optimization.
	Id *string `json:"id,omitempty"`
}

// NewOptimizationIdentifier instantiates a new OptimizationIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationIdentifier() *OptimizationIdentifier {
	this := OptimizationIdentifier{}
	return &this
}

// NewOptimizationIdentifierWithDefaults instantiates a new OptimizationIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationIdentifierWithDefaults() *OptimizationIdentifier {
	this := OptimizationIdentifier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptimizationIdentifier) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationIdentifier) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptimizationIdentifier) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OptimizationIdentifier) SetId(v string) {
	o.Id = &v
}

func (o OptimizationIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptimizationIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOptimizationIdentifier struct {
	value *OptimizationIdentifier
	isSet bool
}

func (v NullableOptimizationIdentifier) Get() *OptimizationIdentifier {
	return v.value
}

func (v *NullableOptimizationIdentifier) Set(val *OptimizationIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationIdentifier(val *OptimizationIdentifier) *NullableOptimizationIdentifier {
	return &NullableOptimizationIdentifier{value: val, isSet: true}
}

func (v NullableOptimizationIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptimizationIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


