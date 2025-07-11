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

// checks if the PickupOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PickupOrder{}

// PickupOrder A request to pickup an order and transport it to a depot.
type PickupOrder struct {
	// A unique identifier of the order. This must be unique across all orders.
	Id string `json:"id" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	Pickup TaskProperties `json:"pickup"`
	Properties *OrderProperties `json:"properties,omitempty"`
}

type _PickupOrder PickupOrder

// NewPickupOrder instantiates a new PickupOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPickupOrder(id string, pickup TaskProperties) *PickupOrder {
	this := PickupOrder{}
	this.Id = id
	this.Pickup = pickup
	return &this
}

// NewPickupOrderWithDefaults instantiates a new PickupOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPickupOrderWithDefaults() *PickupOrder {
	this := PickupOrder{}
	return &this
}

// GetId returns the Id field value
func (o *PickupOrder) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PickupOrder) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PickupOrder) SetId(v string) {
	o.Id = v
}

// GetPickup returns the Pickup field value
func (o *PickupOrder) GetPickup() TaskProperties {
	if o == nil {
		var ret TaskProperties
		return ret
	}

	return o.Pickup
}

// GetPickupOk returns a tuple with the Pickup field value
// and a boolean to check if the value has been set.
func (o *PickupOrder) GetPickupOk() (*TaskProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pickup, true
}

// SetPickup sets field value
func (o *PickupOrder) SetPickup(v TaskProperties) {
	o.Pickup = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PickupOrder) GetProperties() OrderProperties {
	if o == nil || IsNil(o.Properties) {
		var ret OrderProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickupOrder) GetPropertiesOk() (*OrderProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PickupOrder) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given OrderProperties and assigns it to the Properties field.
func (o *PickupOrder) SetProperties(v OrderProperties) {
	o.Properties = &v
}

func (o PickupOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PickupOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["pickup"] = o.Pickup
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

func (o *PickupOrder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"pickup",
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

	varPickupOrder := _PickupOrder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPickupOrder)

	if err != nil {
		return err
	}

	*o = PickupOrder(varPickupOrder)

	return err
}

type NullablePickupOrder struct {
	value *PickupOrder
	isSet bool
}

func (v NullablePickupOrder) Get() *PickupOrder {
	return v.value
}

func (v *NullablePickupOrder) Set(val *PickupOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePickupOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePickupOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePickupOrder(val *PickupOrder) *NullablePickupOrder {
	return &NullablePickupOrder{value: val, isSet: true}
}

func (v NullablePickupOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePickupOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


