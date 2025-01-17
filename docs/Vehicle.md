# Vehicle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID across all vehicles. The ID does not influence the result except the sorting of the routes. | 
**Capacities** | Pointer to **[]int32** | List of capacities for the different quantity dimensions of goods that can be transported. The maximum length of this list is 100. That is, up to 100 different quantity dimensions (e.g. number of pallets, weight, volume, etc.) can be distinguished. Transports can only be executed by a vehicle with a higher (or an equal) maximum capacity in every quantity dimension. The length of this list has to be the same for all transports and all vehicles. If and only if this list of capacities is empty for all vehicles, the quantities of each transport must be empty. | [optional] 
**AlternativeCapacities** | Pointer to **[][]int32** | List of alternative capacities. Each entry in this array must be valid capacities, more details are described in **capacities**. If a route cannot be driven with the **capacities**, alternative capacities may be chosen by the optimization considering the chosen **capacitiesChangePosition**.  See [here](./concepts/capacities-and-alternative-capacities) for more information. | [optional] 
**CapacitiesChangePosition** | Pointer to [**CapacitiesChangePosition**](CapacitiesChangePosition.md) |  | [optional] [default to BETWEEN_TRIPS]
**Equipment** | Pointer to **[]string** | List of vehicle equipment. A transport can only be served by the vehicle if this list is a superset of (or equal to) the transport&#39;s required vehicle equipment. | [optional] 
**Profile** | Pointer to **string** | The profile defines attributes of the vehicle relevant to determine travel times and distances between any two locations. See [here](./concepts/profiles-and-countries) for a complete list of allowed values. If the majority of locations are in the Americas, _USA_8_SEMITRAILER_5AXLE_ is used as the default. Otherwise, _EUR_TRAILER_TRUCK_ is used as the default.  If most locations are located in the Americas but a non-American vehicle profile is specified or vice-versa, a warning is returned. Always use a vehicle profile which matches the region of the locations to obtain best results. | [optional] 
**StartLocationId** | Pointer to **string** | ID of the vehicle&#39;s start location. If the vehicle start location does not coincide with the location of the first pickup, only the coordinates of the vehicle start location are considered and all other attributes such as opening intervals, service time or type are ignored. If no start location is specified, it is assumed that the vehicle is available at the first stop. | [optional] 
**EndLocationId** | Pointer to **string** | ID of the vehicle&#39;s end location. If the vehicle end location does not coincide with the location of the last delivery, only the coordinates of the vehicle end location are considered and all other attributes such as opening intervals, service time or type are ignored. If no end location is specified, it is assumed that the vehicle remains at the last stop. | [optional] 
**ServiceTimePerTransportStop** | Pointer to **int32** | Vehicle-dependent service time [s], for example, for maneuvering. This service time is taken into account for each stop served by this vehicle to pick up or deliver goods. Besides a vehicle-dependent service time, the user may specify location- and transport-dependent service times at the locations and the transports respectively. | [optional] [default to 0]
**ServiceTimeFactor** | Pointer to **float64** | A factor that scales transport-dependent service times of all transport-related service actions, for example, if loading and unloading is more or less complicated than for other vehicles. A factor less than one means that the vehicle speeds up the service, a factor greater than one means that it slows the service down. | [optional] [default to 1]
**IgnoreMixedLoadingProhibitions** | Pointer to **bool** | Indicates whether the mixed loading prohibitions are relevant for this vehicle. If set to false, the mixed loading prohibitions -- if there are any -- must be respected on every trip of the vehicle. If set to true, the mixed loading prohibitions are ignored by the vehicle. | [optional] [default to false]
**RouteStartInterval** | Pointer to [**TimeInterval**](TimeInterval.md) | Interval in which the vehicle has to start its route. Start and end of the interval may be the same. Leaving this parameter empty means that the route start is unrestricted. | [optional] 
**MaximumDistance** | Pointer to **NullableInt32** | Restricts the maximum allowed total driving distance [m] for the route of this vehicle. | [optional] 
**MaximumNumberOfCustomerStops** | Pointer to **NullableInt32** | Restricts the maximum allowed total number of customer stops, i.e. stops at customer locations, for the route of this vehicle. | [optional] 

## Methods

### NewVehicle

`func NewVehicle(id string, ) *Vehicle`

NewVehicle instantiates a new Vehicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleWithDefaults

`func NewVehicleWithDefaults() *Vehicle`

NewVehicleWithDefaults instantiates a new Vehicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vehicle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vehicle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vehicle) SetId(v string)`

SetId sets Id field to given value.


### GetCapacities

`func (o *Vehicle) GetCapacities() []int32`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *Vehicle) GetCapacitiesOk() (*[]int32, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *Vehicle) SetCapacities(v []int32)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *Vehicle) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetAlternativeCapacities

`func (o *Vehicle) GetAlternativeCapacities() [][]int32`

GetAlternativeCapacities returns the AlternativeCapacities field if non-nil, zero value otherwise.

### GetAlternativeCapacitiesOk

`func (o *Vehicle) GetAlternativeCapacitiesOk() (*[][]int32, bool)`

GetAlternativeCapacitiesOk returns a tuple with the AlternativeCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeCapacities

`func (o *Vehicle) SetAlternativeCapacities(v [][]int32)`

SetAlternativeCapacities sets AlternativeCapacities field to given value.

### HasAlternativeCapacities

`func (o *Vehicle) HasAlternativeCapacities() bool`

HasAlternativeCapacities returns a boolean if a field has been set.

### GetCapacitiesChangePosition

`func (o *Vehicle) GetCapacitiesChangePosition() CapacitiesChangePosition`

GetCapacitiesChangePosition returns the CapacitiesChangePosition field if non-nil, zero value otherwise.

### GetCapacitiesChangePositionOk

`func (o *Vehicle) GetCapacitiesChangePositionOk() (*CapacitiesChangePosition, bool)`

GetCapacitiesChangePositionOk returns a tuple with the CapacitiesChangePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacitiesChangePosition

`func (o *Vehicle) SetCapacitiesChangePosition(v CapacitiesChangePosition)`

SetCapacitiesChangePosition sets CapacitiesChangePosition field to given value.

### HasCapacitiesChangePosition

`func (o *Vehicle) HasCapacitiesChangePosition() bool`

HasCapacitiesChangePosition returns a boolean if a field has been set.

### GetEquipment

`func (o *Vehicle) GetEquipment() []string`

GetEquipment returns the Equipment field if non-nil, zero value otherwise.

### GetEquipmentOk

`func (o *Vehicle) GetEquipmentOk() (*[]string, bool)`

GetEquipmentOk returns a tuple with the Equipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipment

`func (o *Vehicle) SetEquipment(v []string)`

SetEquipment sets Equipment field to given value.

### HasEquipment

`func (o *Vehicle) HasEquipment() bool`

HasEquipment returns a boolean if a field has been set.

### GetProfile

`func (o *Vehicle) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Vehicle) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Vehicle) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Vehicle) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStartLocationId

`func (o *Vehicle) GetStartLocationId() string`

GetStartLocationId returns the StartLocationId field if non-nil, zero value otherwise.

### GetStartLocationIdOk

`func (o *Vehicle) GetStartLocationIdOk() (*string, bool)`

GetStartLocationIdOk returns a tuple with the StartLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocationId

`func (o *Vehicle) SetStartLocationId(v string)`

SetStartLocationId sets StartLocationId field to given value.

### HasStartLocationId

`func (o *Vehicle) HasStartLocationId() bool`

HasStartLocationId returns a boolean if a field has been set.

### GetEndLocationId

`func (o *Vehicle) GetEndLocationId() string`

GetEndLocationId returns the EndLocationId field if non-nil, zero value otherwise.

### GetEndLocationIdOk

`func (o *Vehicle) GetEndLocationIdOk() (*string, bool)`

GetEndLocationIdOk returns a tuple with the EndLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocationId

`func (o *Vehicle) SetEndLocationId(v string)`

SetEndLocationId sets EndLocationId field to given value.

### HasEndLocationId

`func (o *Vehicle) HasEndLocationId() bool`

HasEndLocationId returns a boolean if a field has been set.

### GetServiceTimePerTransportStop

`func (o *Vehicle) GetServiceTimePerTransportStop() int32`

GetServiceTimePerTransportStop returns the ServiceTimePerTransportStop field if non-nil, zero value otherwise.

### GetServiceTimePerTransportStopOk

`func (o *Vehicle) GetServiceTimePerTransportStopOk() (*int32, bool)`

GetServiceTimePerTransportStopOk returns a tuple with the ServiceTimePerTransportStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTimePerTransportStop

`func (o *Vehicle) SetServiceTimePerTransportStop(v int32)`

SetServiceTimePerTransportStop sets ServiceTimePerTransportStop field to given value.

### HasServiceTimePerTransportStop

`func (o *Vehicle) HasServiceTimePerTransportStop() bool`

HasServiceTimePerTransportStop returns a boolean if a field has been set.

### GetServiceTimeFactor

`func (o *Vehicle) GetServiceTimeFactor() float64`

GetServiceTimeFactor returns the ServiceTimeFactor field if non-nil, zero value otherwise.

### GetServiceTimeFactorOk

`func (o *Vehicle) GetServiceTimeFactorOk() (*float64, bool)`

GetServiceTimeFactorOk returns a tuple with the ServiceTimeFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTimeFactor

`func (o *Vehicle) SetServiceTimeFactor(v float64)`

SetServiceTimeFactor sets ServiceTimeFactor field to given value.

### HasServiceTimeFactor

`func (o *Vehicle) HasServiceTimeFactor() bool`

HasServiceTimeFactor returns a boolean if a field has been set.

### GetIgnoreMixedLoadingProhibitions

`func (o *Vehicle) GetIgnoreMixedLoadingProhibitions() bool`

GetIgnoreMixedLoadingProhibitions returns the IgnoreMixedLoadingProhibitions field if non-nil, zero value otherwise.

### GetIgnoreMixedLoadingProhibitionsOk

`func (o *Vehicle) GetIgnoreMixedLoadingProhibitionsOk() (*bool, bool)`

GetIgnoreMixedLoadingProhibitionsOk returns a tuple with the IgnoreMixedLoadingProhibitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMixedLoadingProhibitions

`func (o *Vehicle) SetIgnoreMixedLoadingProhibitions(v bool)`

SetIgnoreMixedLoadingProhibitions sets IgnoreMixedLoadingProhibitions field to given value.

### HasIgnoreMixedLoadingProhibitions

`func (o *Vehicle) HasIgnoreMixedLoadingProhibitions() bool`

HasIgnoreMixedLoadingProhibitions returns a boolean if a field has been set.

### GetRouteStartInterval

`func (o *Vehicle) GetRouteStartInterval() TimeInterval`

GetRouteStartInterval returns the RouteStartInterval field if non-nil, zero value otherwise.

### GetRouteStartIntervalOk

`func (o *Vehicle) GetRouteStartIntervalOk() (*TimeInterval, bool)`

GetRouteStartIntervalOk returns a tuple with the RouteStartInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteStartInterval

`func (o *Vehicle) SetRouteStartInterval(v TimeInterval)`

SetRouteStartInterval sets RouteStartInterval field to given value.

### HasRouteStartInterval

`func (o *Vehicle) HasRouteStartInterval() bool`

HasRouteStartInterval returns a boolean if a field has been set.

### GetMaximumDistance

`func (o *Vehicle) GetMaximumDistance() int32`

GetMaximumDistance returns the MaximumDistance field if non-nil, zero value otherwise.

### GetMaximumDistanceOk

`func (o *Vehicle) GetMaximumDistanceOk() (*int32, bool)`

GetMaximumDistanceOk returns a tuple with the MaximumDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDistance

`func (o *Vehicle) SetMaximumDistance(v int32)`

SetMaximumDistance sets MaximumDistance field to given value.

### HasMaximumDistance

`func (o *Vehicle) HasMaximumDistance() bool`

HasMaximumDistance returns a boolean if a field has been set.

### SetMaximumDistanceNil

`func (o *Vehicle) SetMaximumDistanceNil(b bool)`

 SetMaximumDistanceNil sets the value for MaximumDistance to be an explicit nil

### UnsetMaximumDistance
`func (o *Vehicle) UnsetMaximumDistance()`

UnsetMaximumDistance ensures that no value is present for MaximumDistance, not even an explicit nil
### GetMaximumNumberOfCustomerStops

`func (o *Vehicle) GetMaximumNumberOfCustomerStops() int32`

GetMaximumNumberOfCustomerStops returns the MaximumNumberOfCustomerStops field if non-nil, zero value otherwise.

### GetMaximumNumberOfCustomerStopsOk

`func (o *Vehicle) GetMaximumNumberOfCustomerStopsOk() (*int32, bool)`

GetMaximumNumberOfCustomerStopsOk returns a tuple with the MaximumNumberOfCustomerStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumberOfCustomerStops

`func (o *Vehicle) SetMaximumNumberOfCustomerStops(v int32)`

SetMaximumNumberOfCustomerStops sets MaximumNumberOfCustomerStops field to given value.

### HasMaximumNumberOfCustomerStops

`func (o *Vehicle) HasMaximumNumberOfCustomerStops() bool`

HasMaximumNumberOfCustomerStops returns a boolean if a field has been set.

### SetMaximumNumberOfCustomerStopsNil

`func (o *Vehicle) SetMaximumNumberOfCustomerStopsNil(b bool)`

 SetMaximumNumberOfCustomerStopsNil sets the value for MaximumNumberOfCustomerStops to be an explicit nil

### UnsetMaximumNumberOfCustomerStops
`func (o *Vehicle) UnsetMaximumNumberOfCustomerStops()`

UnsetMaximumNumberOfCustomerStops ensures that no value is present for MaximumNumberOfCustomerStops, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


