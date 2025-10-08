# VehicleRouting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | **string** | A profile defines a vehicle by a set of attributes, matching typical transport situations. It must be the name of one of the predefined profiles &#x60;EUR_PEDESTRIAN&#x60;, &#x60;EUR_BICYCLE&#x60;, &#x60;EUR_CAR&#x60;, &#x60;EUR_VAN&#x60;, &#x60;EUR_TRUCK_7_49T&#x60;, &#x60;EUR_TRUCK_11_99T&#x60;, &#x60;EUR_TRUCK_40T&#x60;, &#x60;EUR_TRAILER_TRUCK&#x60;, &#x60;EUR_TLN_VAN&#x60;, &#x60;EUR_TLN_TRUCK_11_99T&#x60;, &#x60;EUR_TLN_TRUCK_20T&#x60;, &#x60;EUR_TLN_TRUCK_40T&#x60;, &#x60;AUS_LCV_LIGHT_COMMERCIAL&#x60;, &#x60;AUS_MR_MEDIUM_RIGID&#x60;, &#x60;AUS_HR_HEAVY_RIGID&#x60;, &#x60;IMEA_CAR&#x60;, &#x60;IMEA_VAN&#x60;, &#x60;IMEA_TRUCK_7_49T&#x60;, &#x60;IMEA_TRUCK_40T&#x60;, &#x60;USA_1_PICKUP&#x60;, &#x60;USA_5_DELIVERY&#x60;, &#x60;USA_8_SEMITRAILER_5AXLE&#x60;. At most ten profiles may be used within a single optimization. Please note that the upper bound on the number of different routing profiles is a technical limit. Check your individual price plan or contract to see which limits apply. | 
**SpeedFactor** | Pointer to **float64** | An additional factor to apply to the speed of the vehicle. When lower than one, the driving durations of the vehicle will increase, when greater than one, the driving durations of the vehicle will decrease. | [optional] [default to 1]
**Violations** | Pointer to [**RoutingViolationStrategy**](RoutingViolationStrategy.md) |  | [optional] [default to ALLOW]
**TrafficMode** | Pointer to [**TrafficMode**](TrafficMode.md) |  | [optional] [default to CONSTANT]

## Methods

### NewVehicleRouting

`func NewVehicleRouting(profile string, ) *VehicleRouting`

NewVehicleRouting instantiates a new VehicleRouting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleRoutingWithDefaults

`func NewVehicleRoutingWithDefaults() *VehicleRouting`

NewVehicleRoutingWithDefaults instantiates a new VehicleRouting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *VehicleRouting) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VehicleRouting) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VehicleRouting) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetSpeedFactor

`func (o *VehicleRouting) GetSpeedFactor() float64`

GetSpeedFactor returns the SpeedFactor field if non-nil, zero value otherwise.

### GetSpeedFactorOk

`func (o *VehicleRouting) GetSpeedFactorOk() (*float64, bool)`

GetSpeedFactorOk returns a tuple with the SpeedFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedFactor

`func (o *VehicleRouting) SetSpeedFactor(v float64)`

SetSpeedFactor sets SpeedFactor field to given value.

### HasSpeedFactor

`func (o *VehicleRouting) HasSpeedFactor() bool`

HasSpeedFactor returns a boolean if a field has been set.

### GetViolations

`func (o *VehicleRouting) GetViolations() RoutingViolationStrategy`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *VehicleRouting) GetViolationsOk() (*RoutingViolationStrategy, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *VehicleRouting) SetViolations(v RoutingViolationStrategy)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *VehicleRouting) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### GetTrafficMode

`func (o *VehicleRouting) GetTrafficMode() TrafficMode`

GetTrafficMode returns the TrafficMode field if non-nil, zero value otherwise.

### GetTrafficModeOk

`func (o *VehicleRouting) GetTrafficModeOk() (*TrafficMode, bool)`

GetTrafficModeOk returns a tuple with the TrafficMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficMode

`func (o *VehicleRouting) SetTrafficMode(v TrafficMode)`

SetTrafficMode sets TrafficMode field to given value.

### HasTrafficMode

`func (o *VehicleRouting) HasTrafficMode() bool`

HasTrafficMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


