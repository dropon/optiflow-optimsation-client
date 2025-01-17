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

// checks if the RoadAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoadAccess{}

// RoadAccess Use these coordinates for matching to the nearest road. Implies **includeLastMeters**, i.e. the air-line connection between the location coordinates and the matched coordinates is included in the relation distance and travel time. This is useful if the location should not be matched to the nearest possible road but to some road further away, e.g. garage exit at a different road.
type RoadAccess struct {
	// The latitude value in degrees (WGS84/EPSG:4326) from south to north.
	Latitude float64 `json:"latitude"`
	// The longitude value in degrees (WGS84/EPSG:4326) from west to east.
	Longitude float64 `json:"longitude"`
}

type _RoadAccess RoadAccess

// NewRoadAccess instantiates a new RoadAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoadAccess(latitude float64, longitude float64) *RoadAccess {
	this := RoadAccess{}
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewRoadAccessWithDefaults instantiates a new RoadAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoadAccessWithDefaults() *RoadAccess {
	this := RoadAccess{}
	return &this
}

// GetLatitude returns the Latitude field value
func (o *RoadAccess) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *RoadAccess) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *RoadAccess) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *RoadAccess) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *RoadAccess) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *RoadAccess) SetLongitude(v float64) {
	o.Longitude = v
}

func (o RoadAccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoadAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	return toSerialize, nil
}

func (o *RoadAccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"latitude",
		"longitude",
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

	varRoadAccess := _RoadAccess{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoadAccess)

	if err != nil {
		return err
	}

	*o = RoadAccess(varRoadAccess)

	return err
}

type NullableRoadAccess struct {
	value *RoadAccess
	isSet bool
}

func (v NullableRoadAccess) Get() *RoadAccess {
	return v.value
}

func (v *NullableRoadAccess) Set(val *RoadAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableRoadAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableRoadAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoadAccess(val *RoadAccess) *NullableRoadAccess {
	return &NullableRoadAccess{value: val, isSet: true}
}

func (v NullableRoadAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoadAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

