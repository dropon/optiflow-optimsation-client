# Stop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** | A reference to a location. Usually, goods are picked up or delivered here. If no goods are picked up or delivered, it is the location where a route or a trip starts or ends. | 
**TripId** | Pointer to **string** | The trip ID groups several subsequent stops. The first stop and the last stop of a route may not have a trip ID only if the vehicle that is assigned to this route has a start or an end location, respectively.  See [here](./concepts/routes-and-trips) for more information. | [optional] 
**DeliveryIds** | Pointer to **[]string** | A list of transport IDs that are delivered at this stop. | [optional] 
**PickupIds** | Pointer to **[]string** | A list of transport IDs that are picked up at this stop. | [optional] 
**ReportForWayToStop** | [**WayReport**](WayReport.md) | Contains summary values of the events on the way from the previous stop to the current stop, including the driving time. The very first stop of a route has a way report with empty values. | [readonly] 
**ReportForStop** | [**StopReport**](StopReport.md) | Contains summary values of the events at a stop, including the arrival time, the departure time and the service time It also contains information about the loaded quantities when leaving the current stop. | [readonly] 
**EventsOnWayToStop** | [**[]Event**](Event.md) | A list of events that occur on the way from the previous stop to this stop. An event specifies what happens on a route at a certain point in time. It can describe the driver&#39;s activity such as driving or break. | [readonly] 
**EventsAtStop** | [**[]Event**](Event.md) | A list of events that occur at this stop. An event specifies what happens on a route at a certain point in time. It can describe the driver&#39;s activity such as performing service or waiting, or it can denote the start/end of a route/trip. | [readonly] 
**ViolationsOnWayToStop** | [**[]Violation**](Violation.md) | A list of violations that occur on the way from the previous stop to this stop. The following violation types can occur here: _PLANNING_HORIZON_, _DRIVER_AVAILABILITY_, _MAXIMUM_TRAVEL_TIME_OF_DRIVER_, _MAXIMUM_DRIVING_TIME_OF_DRIVER_, _MAXIMUM_DISTANCE_, _REST_POSITION_.    _REST_POSITION_ is reported if a daily rest has to be taken on the way to this stop to respect the daily rest rule. The expected position for a daily rest is at the end of a trip. For all other violations, the maximum time or distance exceedance on the way to this stop is reported.    Violations of type _PLANNING_HORIZON_ and _MAXIMUM_TRAVEL_TIME_OF_DRIVER_ will reoccur (with increasing value for TimeExceedence) for all following ways and stops of the current route.    Violations of type _MAXIMUM_DRIVING_TIME_OF_DRIVER_ and _MAXIMUM_DISTANCE_ will reoccur (with increasing value for TimeExceedence resp. DistanceExceedance) for all following ways of the current route.    Violations of type _DRIVER_AVAILABILITY_ will reoccur (with increasing value for TimeExceedence) for all following ways and stops of the current trip. | [readonly] 
**ViolationsAtStop** | [**[]Violation**](Violation.md) | A list of violations that occur at this stop. The following violation types can occur here: _PLANNING_HORIZON_, _DRIVER_AVAILABILITY_, _OPENING_INTERVAL_, _VEHICLE_CAPACITY_, _VEHICLE_EQUIPMENT_, _ROUTE_START_INTERVAL_, _MAXIMUM_TRAVEL_TIME_OF_DRIVER_, _MAXIMUM_NUMBER_OF_CUSTOMER_STOPS_, _REST_POSITION_, _STOP_POSITION_IN_TRIP_, _TRIP_SECTION_, _MIXED_LOADING_PROHIBITION_.    _ROUTE_START_INTERVAL_ is reported at the first stop in a route.    For _VEHICLE_CAPACITY_ and _VEHICLE_EQUIPMENT_, the violation when leaving the stop is reported (i.e., at the last stop of the route, when everything is unloaded, there will be no such violation).    _MAXIMUM_NUMBER_OF_CUSTOMER_STOPS_ is reported when the current stop reaches the limit of the number of customer stops and is repeated for all following customer stops (with increasing value for NumberOfStopsExceedance).    _REST_POSITION_ is reported if a daily rest has to be taken at this stop to respect the daily rest rule. The expected position for a daily rest is at the end of a trip.    For all other violations, the maximum time exceedance at this stop is reported.    Violations of type _PLANNING_HORIZON_ and _MAXIMUM_TRAVEL_TIME_OF_DRIVER_ will reoccur (with increasing value for TimeExceedence) for all following ways and stops of the current route.    Violations of type _DRIVER_AVAILABILITY_ will reoccur (with increasing value for TimeExceedence) for all following ways and stops of the current trip.    Violations of type _VEHICLE_CAPACITY_ will reoccur at all following stops until enough goods are unloaded and the vehicle capacity is not exceeded anymore.    Violations of type _VEHICLE_EQUIPMENT_ will reoccur at all following stops until goods which require the missing equipment are unloaded and no equipment is missing anymore.    Violations of type _STOP_POSITION_IN_TRIP_ will occur at all stops where a position in trip is set for this location ant the current customer stop is not the first/last in the trip.    Violations of type _TRIP_SECTION_ will occur at all stops where the previous trip section number is higher than the one at this customer location.    Violations of type _MIXED_LOADING_PROHIBITION_ will occur at all stops where a load whose category must not be loaded with another load category in the trip is picked up. | [readonly] 

## Methods

### NewStop

`func NewStop(locationId string, reportForWayToStop WayReport, reportForStop StopReport, eventsOnWayToStop []Event, eventsAtStop []Event, violationsOnWayToStop []Violation, violationsAtStop []Violation, ) *Stop`

NewStop instantiates a new Stop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopWithDefaults

`func NewStopWithDefaults() *Stop`

NewStopWithDefaults instantiates a new Stop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *Stop) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Stop) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Stop) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetTripId

`func (o *Stop) GetTripId() string`

GetTripId returns the TripId field if non-nil, zero value otherwise.

### GetTripIdOk

`func (o *Stop) GetTripIdOk() (*string, bool)`

GetTripIdOk returns a tuple with the TripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripId

`func (o *Stop) SetTripId(v string)`

SetTripId sets TripId field to given value.

### HasTripId

`func (o *Stop) HasTripId() bool`

HasTripId returns a boolean if a field has been set.

### GetDeliveryIds

`func (o *Stop) GetDeliveryIds() []string`

GetDeliveryIds returns the DeliveryIds field if non-nil, zero value otherwise.

### GetDeliveryIdsOk

`func (o *Stop) GetDeliveryIdsOk() (*[]string, bool)`

GetDeliveryIdsOk returns a tuple with the DeliveryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIds

`func (o *Stop) SetDeliveryIds(v []string)`

SetDeliveryIds sets DeliveryIds field to given value.

### HasDeliveryIds

`func (o *Stop) HasDeliveryIds() bool`

HasDeliveryIds returns a boolean if a field has been set.

### GetPickupIds

`func (o *Stop) GetPickupIds() []string`

GetPickupIds returns the PickupIds field if non-nil, zero value otherwise.

### GetPickupIdsOk

`func (o *Stop) GetPickupIdsOk() (*[]string, bool)`

GetPickupIdsOk returns a tuple with the PickupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupIds

`func (o *Stop) SetPickupIds(v []string)`

SetPickupIds sets PickupIds field to given value.

### HasPickupIds

`func (o *Stop) HasPickupIds() bool`

HasPickupIds returns a boolean if a field has been set.

### GetReportForWayToStop

`func (o *Stop) GetReportForWayToStop() WayReport`

GetReportForWayToStop returns the ReportForWayToStop field if non-nil, zero value otherwise.

### GetReportForWayToStopOk

`func (o *Stop) GetReportForWayToStopOk() (*WayReport, bool)`

GetReportForWayToStopOk returns a tuple with the ReportForWayToStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportForWayToStop

`func (o *Stop) SetReportForWayToStop(v WayReport)`

SetReportForWayToStop sets ReportForWayToStop field to given value.


### GetReportForStop

`func (o *Stop) GetReportForStop() StopReport`

GetReportForStop returns the ReportForStop field if non-nil, zero value otherwise.

### GetReportForStopOk

`func (o *Stop) GetReportForStopOk() (*StopReport, bool)`

GetReportForStopOk returns a tuple with the ReportForStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportForStop

`func (o *Stop) SetReportForStop(v StopReport)`

SetReportForStop sets ReportForStop field to given value.


### GetEventsOnWayToStop

`func (o *Stop) GetEventsOnWayToStop() []Event`

GetEventsOnWayToStop returns the EventsOnWayToStop field if non-nil, zero value otherwise.

### GetEventsOnWayToStopOk

`func (o *Stop) GetEventsOnWayToStopOk() (*[]Event, bool)`

GetEventsOnWayToStopOk returns a tuple with the EventsOnWayToStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsOnWayToStop

`func (o *Stop) SetEventsOnWayToStop(v []Event)`

SetEventsOnWayToStop sets EventsOnWayToStop field to given value.


### GetEventsAtStop

`func (o *Stop) GetEventsAtStop() []Event`

GetEventsAtStop returns the EventsAtStop field if non-nil, zero value otherwise.

### GetEventsAtStopOk

`func (o *Stop) GetEventsAtStopOk() (*[]Event, bool)`

GetEventsAtStopOk returns a tuple with the EventsAtStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsAtStop

`func (o *Stop) SetEventsAtStop(v []Event)`

SetEventsAtStop sets EventsAtStop field to given value.


### GetViolationsOnWayToStop

`func (o *Stop) GetViolationsOnWayToStop() []Violation`

GetViolationsOnWayToStop returns the ViolationsOnWayToStop field if non-nil, zero value otherwise.

### GetViolationsOnWayToStopOk

`func (o *Stop) GetViolationsOnWayToStopOk() (*[]Violation, bool)`

GetViolationsOnWayToStopOk returns a tuple with the ViolationsOnWayToStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationsOnWayToStop

`func (o *Stop) SetViolationsOnWayToStop(v []Violation)`

SetViolationsOnWayToStop sets ViolationsOnWayToStop field to given value.


### GetViolationsAtStop

`func (o *Stop) GetViolationsAtStop() []Violation`

GetViolationsAtStop returns the ViolationsAtStop field if non-nil, zero value otherwise.

### GetViolationsAtStopOk

`func (o *Stop) GetViolationsAtStopOk() (*[]Violation, bool)`

GetViolationsAtStopOk returns a tuple with the ViolationsAtStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationsAtStop

`func (o *Stop) SetViolationsAtStop(v []Violation)`

SetViolationsAtStop sets ViolationsAtStop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


