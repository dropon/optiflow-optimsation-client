# RouteMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfStops** | Pointer to **int32** | The number of stops in the route.A stop is a visit at a location where one or more tasks are executed. | [optional] 
**Cost** | Pointer to **float64** | The cost of the route. This includes the cost per hour, cost per kilometer and fixed cost of the vehicle. | [optional] 
**Distance** | Pointer to **int32** | The distance [m] travelled on the route. | [optional] 
**Duration** | Pointer to **int32** | The duration [s] of the route from start to end. | [optional] 

## Methods

### NewRouteMetrics

`func NewRouteMetrics() *RouteMetrics`

NewRouteMetrics instantiates a new RouteMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteMetricsWithDefaults

`func NewRouteMetricsWithDefaults() *RouteMetrics`

NewRouteMetricsWithDefaults instantiates a new RouteMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfStops

`func (o *RouteMetrics) GetNumberOfStops() int32`

GetNumberOfStops returns the NumberOfStops field if non-nil, zero value otherwise.

### GetNumberOfStopsOk

`func (o *RouteMetrics) GetNumberOfStopsOk() (*int32, bool)`

GetNumberOfStopsOk returns a tuple with the NumberOfStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfStops

`func (o *RouteMetrics) SetNumberOfStops(v int32)`

SetNumberOfStops sets NumberOfStops field to given value.

### HasNumberOfStops

`func (o *RouteMetrics) HasNumberOfStops() bool`

HasNumberOfStops returns a boolean if a field has been set.

### GetCost

`func (o *RouteMetrics) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *RouteMetrics) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *RouteMetrics) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *RouteMetrics) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDistance

`func (o *RouteMetrics) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RouteMetrics) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RouteMetrics) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *RouteMetrics) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetDuration

`func (o *RouteMetrics) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RouteMetrics) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RouteMetrics) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RouteMetrics) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


