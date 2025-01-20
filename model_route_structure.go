/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the RouteStructure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteStructure{}

// RouteStructure A description of the route structure that can be used to reconstruct the route.
type RouteStructure struct {
	// The unique identifier of the vehicle that is used to execute the route. Only a single route can be provided for each vehicle.
	VehicleId string `json:"vehicleId" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	// The point in time when the route should start. The start time will be respected as closely as possible. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before `1970-01-01T00:00:00+00:00` nor after `2037-12-31T23:59:59+00:00`. The date must provide an offset to UTC.
	Start time.Time `json:"start"`
	// A sequence of tasks scheduled on the route.
	Tasks []TaskStructure `json:"tasks,omitempty"`
	// A list of breaks scheduled on the route. When omitted, reconstruction will make a best effort to schedule breaks in order to satisfy the break settings.
	Breaks []BreakStructure `json:"breaks,omitempty"`
}

type _RouteStructure RouteStructure

// NewRouteStructure instantiates a new RouteStructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteStructure(vehicleId string, start time.Time) *RouteStructure {
	this := RouteStructure{}
	this.VehicleId = vehicleId
	this.Start = start
	return &this
}

// NewRouteStructureWithDefaults instantiates a new RouteStructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteStructureWithDefaults() *RouteStructure {
	this := RouteStructure{}
	return &this
}

// GetVehicleId returns the VehicleId field value
func (o *RouteStructure) GetVehicleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VehicleId
}

// GetVehicleIdOk returns a tuple with the VehicleId field value
// and a boolean to check if the value has been set.
func (o *RouteStructure) GetVehicleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VehicleId, true
}

// SetVehicleId sets field value
func (o *RouteStructure) SetVehicleId(v string) {
	o.VehicleId = v
}

// GetStart returns the Start field value
func (o *RouteStructure) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *RouteStructure) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *RouteStructure) SetStart(v time.Time) {
	o.Start = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *RouteStructure) GetTasks() []TaskStructure {
	if o == nil || IsNil(o.Tasks) {
		var ret []TaskStructure
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteStructure) GetTasksOk() ([]TaskStructure, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *RouteStructure) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []TaskStructure and assigns it to the Tasks field.
func (o *RouteStructure) SetTasks(v []TaskStructure) {
	o.Tasks = v
}

// GetBreaks returns the Breaks field value if set, zero value otherwise.
func (o *RouteStructure) GetBreaks() []BreakStructure {
	if o == nil || IsNil(o.Breaks) {
		var ret []BreakStructure
		return ret
	}
	return o.Breaks
}

// GetBreaksOk returns a tuple with the Breaks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteStructure) GetBreaksOk() ([]BreakStructure, bool) {
	if o == nil || IsNil(o.Breaks) {
		return nil, false
	}
	return o.Breaks, true
}

// HasBreaks returns a boolean if a field has been set.
func (o *RouteStructure) HasBreaks() bool {
	if o != nil && !IsNil(o.Breaks) {
		return true
	}

	return false
}

// SetBreaks gets a reference to the given []BreakStructure and assigns it to the Breaks field.
func (o *RouteStructure) SetBreaks(v []BreakStructure) {
	o.Breaks = v
}

func (o RouteStructure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteStructure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vehicleId"] = o.VehicleId
	toSerialize["start"] = o.Start
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Breaks) {
		toSerialize["breaks"] = o.Breaks
	}
	return toSerialize, nil
}

func (o *RouteStructure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vehicleId",
		"start",
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

	varRouteStructure := _RouteStructure{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRouteStructure)

	if err != nil {
		return err
	}

	*o = RouteStructure(varRouteStructure)

	return err
}

type NullableRouteStructure struct {
	value *RouteStructure
	isSet bool
}

func (v NullableRouteStructure) Get() *RouteStructure {
	return v.value
}

func (v *NullableRouteStructure) Set(val *RouteStructure) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteStructure) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteStructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteStructure(val *RouteStructure) *NullableRouteStructure {
	return &NullableRouteStructure{value: val, isSet: true}
}

func (v NullableRouteStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteStructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


