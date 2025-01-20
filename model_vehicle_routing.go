/*
Route Optimization OptiFlow

With the Route Optimization OptiFlow service you can schedule and optimize the routes of your fleet.

API version: 1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VehicleRouting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleRouting{}

// VehicleRouting Specifies how time-distance data must be calculated for routes driven by the vehicle.
type VehicleRouting struct {
	// A profile defines a vehicle by a set of attributes, matching typical transport situations. It must be the name of one of the predefined profiles `EUR_PEDESTRIAN`, `EUR_BICYCLE`, `EUR_CAR`, `EUR_VAN`, `EUR_TRUCK_7_49T`, `EUR_TRUCK_11_99T`, `EUR_TRUCK_40T`, `EUR_TRAILER_TRUCK`, `AUS_LCV_LIGHT_COMMERCIAL`, `AUS_MR_MEDIUM_RIGID`, `AUS_HR_HEAVY_RIGID`, `IMEA_CAR`, `IMEA_VAN`, `IMEA_TRUCK_7_49T`, `IMEA_TRUCK_40T`, `USA_1_PICKUP`, `USA_5_DELIVERY`, `USA_8_SEMITRAILER_5AXLE`. At most ten profiles may be used within a single optimization. Please note that the upper bound on number of different routing profiles is a technical limit. Check your individual price plan or contract to see which limits apply.
	Profile string `json:"profile" validate:"regexp=^[a-zA-Z0-9_-]{1,100}$"`
	// An additional factor to apply to the speed of the vehicle. When lower than one, the driving durations of the vehicle will increase, when greater than one, the driving durations of the vehicle will decrease.
	SpeedFactor *float64 `json:"speedFactor,omitempty"`
	Violations *RoutingViolationStrategy `json:"violations,omitempty"`
	TrafficMode *TrafficMode `json:"trafficMode,omitempty"`
}

type _VehicleRouting VehicleRouting

// NewVehicleRouting instantiates a new VehicleRouting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleRouting(profile string) *VehicleRouting {
	this := VehicleRouting{}
	this.Profile = profile
	var speedFactor float64 = 1
	this.SpeedFactor = &speedFactor
	var violations RoutingViolationStrategy = ALLOW
	this.Violations = &violations
	var trafficMode TrafficMode = CONSTANT
	this.TrafficMode = &trafficMode
	return &this
}

// NewVehicleRoutingWithDefaults instantiates a new VehicleRouting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleRoutingWithDefaults() *VehicleRouting {
	this := VehicleRouting{}
	var speedFactor float64 = 1
	this.SpeedFactor = &speedFactor
	var violations RoutingViolationStrategy = ALLOW
	this.Violations = &violations
	var trafficMode TrafficMode = CONSTANT
	this.TrafficMode = &trafficMode
	return &this
}

// GetProfile returns the Profile field value
func (o *VehicleRouting) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *VehicleRouting) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *VehicleRouting) SetProfile(v string) {
	o.Profile = v
}

// GetSpeedFactor returns the SpeedFactor field value if set, zero value otherwise.
func (o *VehicleRouting) GetSpeedFactor() float64 {
	if o == nil || IsNil(o.SpeedFactor) {
		var ret float64
		return ret
	}
	return *o.SpeedFactor
}

// GetSpeedFactorOk returns a tuple with the SpeedFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleRouting) GetSpeedFactorOk() (*float64, bool) {
	if o == nil || IsNil(o.SpeedFactor) {
		return nil, false
	}
	return o.SpeedFactor, true
}

// HasSpeedFactor returns a boolean if a field has been set.
func (o *VehicleRouting) HasSpeedFactor() bool {
	if o != nil && !IsNil(o.SpeedFactor) {
		return true
	}

	return false
}

// SetSpeedFactor gets a reference to the given float64 and assigns it to the SpeedFactor field.
func (o *VehicleRouting) SetSpeedFactor(v float64) {
	o.SpeedFactor = &v
}

// GetViolations returns the Violations field value if set, zero value otherwise.
func (o *VehicleRouting) GetViolations() RoutingViolationStrategy {
	if o == nil || IsNil(o.Violations) {
		var ret RoutingViolationStrategy
		return ret
	}
	return *o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleRouting) GetViolationsOk() (*RoutingViolationStrategy, bool) {
	if o == nil || IsNil(o.Violations) {
		return nil, false
	}
	return o.Violations, true
}

// HasViolations returns a boolean if a field has been set.
func (o *VehicleRouting) HasViolations() bool {
	if o != nil && !IsNil(o.Violations) {
		return true
	}

	return false
}

// SetViolations gets a reference to the given RoutingViolationStrategy and assigns it to the Violations field.
func (o *VehicleRouting) SetViolations(v RoutingViolationStrategy) {
	o.Violations = &v
}

// GetTrafficMode returns the TrafficMode field value if set, zero value otherwise.
func (o *VehicleRouting) GetTrafficMode() TrafficMode {
	if o == nil || IsNil(o.TrafficMode) {
		var ret TrafficMode
		return ret
	}
	return *o.TrafficMode
}

// GetTrafficModeOk returns a tuple with the TrafficMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleRouting) GetTrafficModeOk() (*TrafficMode, bool) {
	if o == nil || IsNil(o.TrafficMode) {
		return nil, false
	}
	return o.TrafficMode, true
}

// HasTrafficMode returns a boolean if a field has been set.
func (o *VehicleRouting) HasTrafficMode() bool {
	if o != nil && !IsNil(o.TrafficMode) {
		return true
	}

	return false
}

// SetTrafficMode gets a reference to the given TrafficMode and assigns it to the TrafficMode field.
func (o *VehicleRouting) SetTrafficMode(v TrafficMode) {
	o.TrafficMode = &v
}

func (o VehicleRouting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleRouting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profile"] = o.Profile
	if !IsNil(o.SpeedFactor) {
		toSerialize["speedFactor"] = o.SpeedFactor
	}
	if !IsNil(o.Violations) {
		toSerialize["violations"] = o.Violations
	}
	if !IsNil(o.TrafficMode) {
		toSerialize["trafficMode"] = o.TrafficMode
	}
	return toSerialize, nil
}

func (o *VehicleRouting) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
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

	varVehicleRouting := _VehicleRouting{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVehicleRouting)

	if err != nil {
		return err
	}

	*o = VehicleRouting(varVehicleRouting)

	return err
}

type NullableVehicleRouting struct {
	value *VehicleRouting
	isSet bool
}

func (v NullableVehicleRouting) Get() *VehicleRouting {
	return v.value
}

func (v *NullableVehicleRouting) Set(val *VehicleRouting) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleRouting) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleRouting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleRouting(val *VehicleRouting) *NullableVehicleRouting {
	return &NullableVehicleRouting{value: val, isSet: true}
}

func (v NullableVehicleRouting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleRouting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


