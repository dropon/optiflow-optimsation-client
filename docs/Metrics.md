# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfUnscheduledOrders** | **int32** | The number of orders that are not planned by the optimization. | 
**NumberOfRoutes** | **int32** | The number of routes that where scheduled by the optimization. | 
**TotalCost** | **float64** | The total cost of the scheduled routes. This includes the cost of the routes and the outsourcing cost of the unplanned orders. | 
**TotalDistance** | **int64** | The sum of the distances [m] of the scheduled routes. | 
**TotalDuration** | **int64** | The sum of the durations [s] of the scheduled routes. | 

## Methods

### NewMetrics

`func NewMetrics(numberOfUnscheduledOrders int32, numberOfRoutes int32, totalCost float64, totalDistance int64, totalDuration int64, ) *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfUnscheduledOrders

`func (o *Metrics) GetNumberOfUnscheduledOrders() int32`

GetNumberOfUnscheduledOrders returns the NumberOfUnscheduledOrders field if non-nil, zero value otherwise.

### GetNumberOfUnscheduledOrdersOk

`func (o *Metrics) GetNumberOfUnscheduledOrdersOk() (*int32, bool)`

GetNumberOfUnscheduledOrdersOk returns a tuple with the NumberOfUnscheduledOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUnscheduledOrders

`func (o *Metrics) SetNumberOfUnscheduledOrders(v int32)`

SetNumberOfUnscheduledOrders sets NumberOfUnscheduledOrders field to given value.


### GetNumberOfRoutes

`func (o *Metrics) GetNumberOfRoutes() int32`

GetNumberOfRoutes returns the NumberOfRoutes field if non-nil, zero value otherwise.

### GetNumberOfRoutesOk

`func (o *Metrics) GetNumberOfRoutesOk() (*int32, bool)`

GetNumberOfRoutesOk returns a tuple with the NumberOfRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRoutes

`func (o *Metrics) SetNumberOfRoutes(v int32)`

SetNumberOfRoutes sets NumberOfRoutes field to given value.


### GetTotalCost

`func (o *Metrics) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *Metrics) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *Metrics) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.


### GetTotalDistance

`func (o *Metrics) GetTotalDistance() int64`

GetTotalDistance returns the TotalDistance field if non-nil, zero value otherwise.

### GetTotalDistanceOk

`func (o *Metrics) GetTotalDistanceOk() (*int64, bool)`

GetTotalDistanceOk returns a tuple with the TotalDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDistance

`func (o *Metrics) SetTotalDistance(v int64)`

SetTotalDistance sets TotalDistance field to given value.


### GetTotalDuration

`func (o *Metrics) GetTotalDuration() int64`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *Metrics) GetTotalDurationOk() (*int64, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *Metrics) SetTotalDuration(v int64)`

SetTotalDuration sets TotalDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


