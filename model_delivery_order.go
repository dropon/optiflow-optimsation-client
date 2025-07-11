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

// checks if the DeliveryOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryOrder{}

// DeliveryOrder An order to deliver goods, which are picked up from a depot.
type DeliveryOrder struct {
	// A unique identifier of the order. This must be unique across all orders.
	Id string `json:"id" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	Delivery TaskProperties `json:"delivery"`
	Properties *OrderProperties `json:"properties,omitempty"`
}

type _DeliveryOrder DeliveryOrder

// NewDeliveryOrder instantiates a new DeliveryOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryOrder(id string, delivery TaskProperties) *DeliveryOrder {
	this := DeliveryOrder{}
	this.Id = id
	this.Delivery = delivery
	return &this
}

// NewDeliveryOrderWithDefaults instantiates a new DeliveryOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryOrderWithDefaults() *DeliveryOrder {
	this := DeliveryOrder{}
	return &this
}

// GetId returns the Id field value
func (o *DeliveryOrder) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeliveryOrder) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeliveryOrder) SetId(v string) {
	o.Id = v
}

// GetDelivery returns the Delivery field value
func (o *DeliveryOrder) GetDelivery() TaskProperties {
	if o == nil {
		var ret TaskProperties
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *DeliveryOrder) GetDeliveryOk() (*TaskProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *DeliveryOrder) SetDelivery(v TaskProperties) {
	o.Delivery = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *DeliveryOrder) GetProperties() OrderProperties {
	if o == nil || IsNil(o.Properties) {
		var ret OrderProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOrder) GetPropertiesOk() (*OrderProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *DeliveryOrder) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given OrderProperties and assigns it to the Properties field.
func (o *DeliveryOrder) SetProperties(v OrderProperties) {
	o.Properties = &v
}

func (o DeliveryOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["delivery"] = o.Delivery
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

func (o *DeliveryOrder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"delivery",
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

	varDeliveryOrder := _DeliveryOrder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeliveryOrder)

	if err != nil {
		return err
	}

	*o = DeliveryOrder(varDeliveryOrder)

	return err
}

type NullableDeliveryOrder struct {
	value *DeliveryOrder
	isSet bool
}

func (v NullableDeliveryOrder) Get() *DeliveryOrder {
	return v.value
}

func (v *NullableDeliveryOrder) Set(val *DeliveryOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryOrder(val *DeliveryOrder) *NullableDeliveryOrder {
	return &NullableDeliveryOrder{value: val, isSet: true}
}

func (v NullableDeliveryOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


