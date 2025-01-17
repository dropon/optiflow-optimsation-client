# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleId** | **string** | The ID of the vehicle that is assigned to this route. A vehicle can be assigned to one route at most. | 
**Stops** | [**[]Stop**](Stop.md) | A sequence of stops along this route. Each stop is at a specific location, either a customer location or a depot location. At a customer location, transports are scheduled in order of their type: deliveries before pickups. The sequence of deliveries and pickups is always sorted by Last In - First Out (LIFO). At a depot location a stop groups either deliveries or pickups.  See [here](./concepts/locations-transports-and-stops) for more information. | 
**Report** | [**RouteReport**](RouteReport.md) | Returns a summary of all events and all reports that belong to this route, including the start time and the end time of the route. | [readonly] 

## Methods

### NewRoute

`func NewRoute(vehicleId string, stops []Stop, report RouteReport, ) *Route`

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


### GetReport

`func (o *Route) GetReport() RouteReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *Route) GetReportOk() (*RouteReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *Route) SetReport(v RouteReport)`

SetReport sets Report field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


