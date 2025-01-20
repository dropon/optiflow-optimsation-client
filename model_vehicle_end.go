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

// checks if the VehicleEnd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleEnd{}

// VehicleEnd Describes the end of a route assigned to the vehicle.
type VehicleEnd struct {
	// The unique identifier of the location where a route assigned to the vehicle must end.
	LocationId string `json:"locationId" validate:"regexp=^[a-zA-Z0-9_-]{1,36}$"`
	// The latest point in time a route assigned to the vehicle may end. This must not be earlier than the vehicle's earliest start time. When used in conjunction with a latest start time, the latest end time must not be earlier than the latest start time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before `1970-01-01T00:00:00+00:00` nor after `2037-12-31T23:59:59+00:00`. The date must provide an offset to UTC.
	LatestEndTime time.Time `json:"latestEndTime"`
	// Describes how long [s] it takes for the vehicle to wrap up at its end location before the route ends.
	Duration *int32 `json:"duration,omitempty"`
}

type _VehicleEnd VehicleEnd

// NewVehicleEnd instantiates a new VehicleEnd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleEnd(locationId string, latestEndTime time.Time) *VehicleEnd {
	this := VehicleEnd{}
	this.LocationId = locationId
	this.LatestEndTime = latestEndTime
	var duration int32 = 0
	this.Duration = &duration
	return &this
}

// NewVehicleEndWithDefaults instantiates a new VehicleEnd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleEndWithDefaults() *VehicleEnd {
	this := VehicleEnd{}
	var duration int32 = 0
	this.Duration = &duration
	return &this
}

// GetLocationId returns the LocationId field value
func (o *VehicleEnd) GetLocationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value
// and a boolean to check if the value has been set.
func (o *VehicleEnd) GetLocationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationId, true
}

// SetLocationId sets field value
func (o *VehicleEnd) SetLocationId(v string) {
	o.LocationId = v
}

// GetLatestEndTime returns the LatestEndTime field value
func (o *VehicleEnd) GetLatestEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LatestEndTime
}

// GetLatestEndTimeOk returns a tuple with the LatestEndTime field value
// and a boolean to check if the value has been set.
func (o *VehicleEnd) GetLatestEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestEndTime, true
}

// SetLatestEndTime sets field value
func (o *VehicleEnd) SetLatestEndTime(v time.Time) {
	o.LatestEndTime = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *VehicleEnd) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleEnd) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *VehicleEnd) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *VehicleEnd) SetDuration(v int32) {
	o.Duration = &v
}

func (o VehicleEnd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleEnd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locationId"] = o.LocationId
	toSerialize["latestEndTime"] = o.LatestEndTime
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

func (o *VehicleEnd) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locationId",
		"latestEndTime",
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

	varVehicleEnd := _VehicleEnd{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVehicleEnd)

	if err != nil {
		return err
	}

	*o = VehicleEnd(varVehicleEnd)

	return err
}

type NullableVehicleEnd struct {
	value *VehicleEnd
	isSet bool
}

func (v NullableVehicleEnd) Get() *VehicleEnd {
	return v.value
}

func (v *NullableVehicleEnd) Set(val *VehicleEnd) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleEnd) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleEnd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleEnd(val *VehicleEnd) *NullableVehicleEnd {
	return &NullableVehicleEnd{value: val, isSet: true}
}

func (v NullableVehicleEnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleEnd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

