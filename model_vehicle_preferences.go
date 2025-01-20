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

// checks if the VehiclePreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehiclePreferences{}

// VehiclePreferences Defines preferences for routes assigned to this vehicle that are not reflected in the cost. These preferences guide the optimization process, which **may result in increased costs** as a tradeoff for meeting them.
type VehiclePreferences struct {
	// A scale between 0 and 1 resembling a tradeoff between minimizing distance cost and maximizing compactness, where higher values indicate a stronger preference for compact routes. As the cost per kilometer increases, the influence of this tradeoff becomes more significant. A route is considered compact if all stops for executing non-depot tasks are close to each other.
	Compactness *float64 `json:"compactness,omitempty"`
}

// NewVehiclePreferences instantiates a new VehiclePreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehiclePreferences() *VehiclePreferences {
	this := VehiclePreferences{}
	var compactness float64 = 0
	this.Compactness = &compactness
	return &this
}

// NewVehiclePreferencesWithDefaults instantiates a new VehiclePreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehiclePreferencesWithDefaults() *VehiclePreferences {
	this := VehiclePreferences{}
	var compactness float64 = 0
	this.Compactness = &compactness
	return &this
}

// GetCompactness returns the Compactness field value if set, zero value otherwise.
func (o *VehiclePreferences) GetCompactness() float64 {
	if o == nil || IsNil(o.Compactness) {
		var ret float64
		return ret
	}
	return *o.Compactness
}

// GetCompactnessOk returns a tuple with the Compactness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehiclePreferences) GetCompactnessOk() (*float64, bool) {
	if o == nil || IsNil(o.Compactness) {
		return nil, false
	}
	return o.Compactness, true
}

// HasCompactness returns a boolean if a field has been set.
func (o *VehiclePreferences) HasCompactness() bool {
	if o != nil && !IsNil(o.Compactness) {
		return true
	}

	return false
}

// SetCompactness gets a reference to the given float64 and assigns it to the Compactness field.
func (o *VehiclePreferences) SetCompactness(v float64) {
	o.Compactness = &v
}

func (o VehiclePreferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehiclePreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compactness) {
		toSerialize["compactness"] = o.Compactness
	}
	return toSerialize, nil
}

type NullableVehiclePreferences struct {
	value *VehiclePreferences
	isSet bool
}

func (v NullableVehiclePreferences) Get() *VehiclePreferences {
	return v.value
}

func (v *NullableVehiclePreferences) Set(val *VehiclePreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableVehiclePreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableVehiclePreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehiclePreferences(val *VehiclePreferences) *NullableVehiclePreferences {
	return &NullableVehiclePreferences{value: val, isSet: true}
}

func (v NullableVehiclePreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehiclePreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

