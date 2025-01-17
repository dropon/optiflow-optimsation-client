/*
Route Optimization

With the Route Optimization service you can schedule and optimize the routes of your fleet.

API version: 1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MixedLoadingProhibition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MixedLoadingProhibition{}

// MixedLoadingProhibition There may be some transports that shall not be mixed with some other transports on one trip. For instance, it may be prohibited to load certain dangerous goods together on the same trip, such as flammable solids on the one hand and explosive substances on the other hand. A mixed loading prohibition is a pair of two conflicting load categories that prohibits transports with these load categories to be mixed on the same trip. The load category can be specified for every transport. For a vehicle, there is a flag that lets the vehicle ignore this restriction.  See [here](./concepts/mixed-loading-prohibition) for more information.
type MixedLoadingProhibition struct {
	// A transport with this load category is not allowed to be on the same trip as a transport with load category conflictingLoadCategory2. The load category can be any string but it must not be empty and not the same as conflictingLoadCategory2.
	ConflictingLoadCategory1 string `json:"conflictingLoadCategory1" validate:"regexp=.*[^ ].*"`
	// A transport with this load category is not allowed to be on the same trip as a transport with load category conflictingLoadCategory1. The load category can be any string but it must not be empty and not the same as conflictingLoadCategory1.
	ConflictingLoadCategory2 string `json:"conflictingLoadCategory2" validate:"regexp=.*[^ ].*"`
}

type _MixedLoadingProhibition MixedLoadingProhibition

// NewMixedLoadingProhibition instantiates a new MixedLoadingProhibition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMixedLoadingProhibition(conflictingLoadCategory1 string, conflictingLoadCategory2 string) *MixedLoadingProhibition {
	this := MixedLoadingProhibition{}
	this.ConflictingLoadCategory1 = conflictingLoadCategory1
	this.ConflictingLoadCategory2 = conflictingLoadCategory2
	return &this
}

// NewMixedLoadingProhibitionWithDefaults instantiates a new MixedLoadingProhibition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMixedLoadingProhibitionWithDefaults() *MixedLoadingProhibition {
	this := MixedLoadingProhibition{}
	return &this
}

// GetConflictingLoadCategory1 returns the ConflictingLoadCategory1 field value
func (o *MixedLoadingProhibition) GetConflictingLoadCategory1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConflictingLoadCategory1
}

// GetConflictingLoadCategory1Ok returns a tuple with the ConflictingLoadCategory1 field value
// and a boolean to check if the value has been set.
func (o *MixedLoadingProhibition) GetConflictingLoadCategory1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConflictingLoadCategory1, true
}

// SetConflictingLoadCategory1 sets field value
func (o *MixedLoadingProhibition) SetConflictingLoadCategory1(v string) {
	o.ConflictingLoadCategory1 = v
}

// GetConflictingLoadCategory2 returns the ConflictingLoadCategory2 field value
func (o *MixedLoadingProhibition) GetConflictingLoadCategory2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConflictingLoadCategory2
}

// GetConflictingLoadCategory2Ok returns a tuple with the ConflictingLoadCategory2 field value
// and a boolean to check if the value has been set.
func (o *MixedLoadingProhibition) GetConflictingLoadCategory2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConflictingLoadCategory2, true
}

// SetConflictingLoadCategory2 sets field value
func (o *MixedLoadingProhibition) SetConflictingLoadCategory2(v string) {
	o.ConflictingLoadCategory2 = v
}

func (o MixedLoadingProhibition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MixedLoadingProhibition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conflictingLoadCategory1"] = o.ConflictingLoadCategory1
	toSerialize["conflictingLoadCategory2"] = o.ConflictingLoadCategory2
	return toSerialize, nil
}

func (o *MixedLoadingProhibition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conflictingLoadCategory1",
		"conflictingLoadCategory2",
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

	varMixedLoadingProhibition := _MixedLoadingProhibition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMixedLoadingProhibition)

	if err != nil {
		return err
	}

	*o = MixedLoadingProhibition(varMixedLoadingProhibition)

	return err
}

type NullableMixedLoadingProhibition struct {
	value *MixedLoadingProhibition
	isSet bool
}

func (v NullableMixedLoadingProhibition) Get() *MixedLoadingProhibition {
	return v.value
}

func (v *NullableMixedLoadingProhibition) Set(val *MixedLoadingProhibition) {
	v.value = val
	v.isSet = true
}

func (v NullableMixedLoadingProhibition) IsSet() bool {
	return v.isSet
}

func (v *NullableMixedLoadingProhibition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMixedLoadingProhibition(val *MixedLoadingProhibition) *NullableMixedLoadingProhibition {
	return &NullableMixedLoadingProhibition{value: val, isSet: true}
}

func (v NullableMixedLoadingProhibition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMixedLoadingProhibition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


