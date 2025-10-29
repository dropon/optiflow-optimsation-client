# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleId** | Pointer to **string** | The unique identifier of the vehicle assigned to the route. | [optional] 
**ResourceIds** | Pointer to **[]string** | The unique identifiers of the resources assigned to the vehicle executing the route. | [optional] 
**Start** | Pointer to [**RouteStart**](RouteStart.md) |  | [optional] 
**Stops** | Pointer to [**[]Stop**](Stop.md) | The list of stops scheduled on the route. A stop describes the visit of a location on a route where one or more tasks are scheduled. Its approach describes how to reach the location from the previous location visited on the route. Consecutive tasks are grouped to an appointment if they are assigned to the same time slot. Consecutive appointments are grouped as a stop if they are scheduled at the same location. | [optional] [default to {}]
**End** | Pointer to [**RouteEnd**](RouteEnd.md) |  | [optional] 
**Metrics** | Pointer to [**RouteMetrics**](RouteMetrics.md) |  | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleId

`func (o *Route) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *Route) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *Route) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *Route) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### GetResourceIds

`func (o *Route) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *Route) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *Route) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *Route) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### GetStart

`func (o *Route) GetStart() RouteStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Route) GetStartOk() (*RouteStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Route) SetStart(v RouteStart)`

SetStart sets Start field to given value.

### HasStart

`func (o *Route) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStops

`func (o *Route) GetStops() []Stop`

GetStops returns the Stops field if non-nil, zero value otherwise.

### GetStopsOk

`func (o *Route) GetStopsOk() (*[]Stop, bool)`

GetStopsOk returns a tuple with the Stops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStops

`func (o *Route) SetStops(v []Stop)`

SetStops sets Stops field to given value.

### HasStops

`func (o *Route) HasStops() bool`

HasStops returns a boolean if a field has been set.

### GetEnd

`func (o *Route) GetEnd() RouteEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Route) GetEndOk() (*RouteEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Route) SetEnd(v RouteEnd)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Route) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetMetrics

`func (o *Route) GetMetrics() RouteMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Route) GetMetricsOk() (*RouteMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Route) SetMetrics(v RouteMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Route) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


