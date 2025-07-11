/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LocationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationRule{}

// LocationRule A rule that conditionally modifies location properties.
type LocationRule struct {
	// A list of conditions that describes when the rule applies. The rule applies if any of the conditions are met. A condition is met if all of its properties are matched.
	Conditions []LocationRuleCondition `json:"conditions"`
	PreparationDuration DurationModifier `json:"preparationDuration"`
}

type _LocationRule LocationRule

// NewLocationRule instantiates a new LocationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRule(conditions []LocationRuleCondition, preparationDuration DurationModifier) *LocationRule {
	this := LocationRule{}
	this.Conditions = conditions
	this.PreparationDuration = preparationDuration
	return &this
}

// NewLocationRuleWithDefaults instantiates a new LocationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRuleWithDefaults() *LocationRule {
	this := LocationRule{}
	return &this
}

// GetConditions returns the Conditions field value
func (o *LocationRule) GetConditions() []LocationRuleCondition {
	if o == nil {
		var ret []LocationRuleCondition
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *LocationRule) GetConditionsOk() ([]LocationRuleCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *LocationRule) SetConditions(v []LocationRuleCondition) {
	o.Conditions = v
}

// GetPreparationDuration returns the PreparationDuration field value
func (o *LocationRule) GetPreparationDuration() DurationModifier {
	if o == nil {
		var ret DurationModifier
		return ret
	}

	return o.PreparationDuration
}

// GetPreparationDurationOk returns a tuple with the PreparationDuration field value
// and a boolean to check if the value has been set.
func (o *LocationRule) GetPreparationDurationOk() (*DurationModifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreparationDuration, true
}

// SetPreparationDuration sets field value
func (o *LocationRule) SetPreparationDuration(v DurationModifier) {
	o.PreparationDuration = v
}

func (o LocationRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditions"] = o.Conditions
	toSerialize["preparationDuration"] = o.PreparationDuration
	return toSerialize, nil
}

func (o *LocationRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conditions",
		"preparationDuration",
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

	varLocationRule := _LocationRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocationRule)

	if err != nil {
		return err
	}

	*o = LocationRule(varLocationRule)

	return err
}

type NullableLocationRule struct {
	value *LocationRule
	isSet bool
}

func (v NullableLocationRule) Get() *LocationRule {
	return v.value
}

func (v *NullableLocationRule) Set(val *LocationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRule(val *LocationRule) *NullableLocationRule {
	return &NullableLocationRule{value: val, isSet: true}
}

func (v NullableLocationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


