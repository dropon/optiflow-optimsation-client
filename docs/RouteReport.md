# RouteReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **time.Time** | The start time of the route formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). | 
**EndTime** | **time.Time** | The end time of the route formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). | 
**TravelTime** | **int32** | The total travel time of the route [s]. Equals difference between end time and start time. | 
**Distance** | **int32** | The total driving distance of the route [m]. | 
**DrivingTime** | **int32** | The total driving time of the route [s]. | 
**ServiceTime** | **int32** | The total service time of the route [s]. | 
**WaitingTime** | **int32** | The total waiting time of the route [s]. | 
**BreakTime** | **int32** | The total break time of the route [s]. | 
**RestTime** | **int32** | The total rest time of the route [s]. | 

## Methods

### NewRouteReport

`func NewRouteReport(startTime time.Time, endTime time.Time, travelTime int32, distance int32, drivingTime int32, serviceTime int32, waitingTime int32, breakTime int32, restTime int32, ) *RouteReport`

NewRouteReport instantiates a new RouteReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteReportWithDefaults

`func NewRouteReportWithDefaults() *RouteReport`

NewRouteReportWithDefaults instantiates a new RouteReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *RouteReport) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RouteReport) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RouteReport) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *RouteReport) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RouteReport) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RouteReport) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetTravelTime

`func (o *RouteReport) GetTravelTime() int32`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *RouteReport) GetTravelTimeOk() (*int32, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *RouteReport) SetTravelTime(v int32)`

SetTravelTime sets TravelTime field to given value.


### GetDistance

`func (o *RouteReport) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RouteReport) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RouteReport) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetDrivingTime

`func (o *RouteReport) GetDrivingTime() int32`

GetDrivingTime returns the DrivingTime field if non-nil, zero value otherwise.

### GetDrivingTimeOk

`func (o *RouteReport) GetDrivingTimeOk() (*int32, bool)`

GetDrivingTimeOk returns a tuple with the DrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingTime

`func (o *RouteReport) SetDrivingTime(v int32)`

SetDrivingTime sets DrivingTime field to given value.


### GetServiceTime

`func (o *RouteReport) GetServiceTime() int32`

GetServiceTime returns the ServiceTime field if non-nil, zero value otherwise.

### GetServiceTimeOk

`func (o *RouteReport) GetServiceTimeOk() (*int32, bool)`

GetServiceTimeOk returns a tuple with the ServiceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTime

`func (o *RouteReport) SetServiceTime(v int32)`

SetServiceTime sets ServiceTime field to given value.


### GetWaitingTime

`func (o *RouteReport) GetWaitingTime() int32`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *RouteReport) GetWaitingTimeOk() (*int32, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *RouteReport) SetWaitingTime(v int32)`

SetWaitingTime sets WaitingTime field to given value.


### GetBreakTime

`func (o *RouteReport) GetBreakTime() int32`

GetBreakTime returns the BreakTime field if non-nil, zero value otherwise.

### GetBreakTimeOk

`func (o *RouteReport) GetBreakTimeOk() (*int32, bool)`

GetBreakTimeOk returns a tuple with the BreakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakTime

`func (o *RouteReport) SetBreakTime(v int32)`

SetBreakTime sets BreakTime field to given value.


### GetRestTime

`func (o *RouteReport) GetRestTime() int32`

GetRestTime returns the RestTime field if non-nil, zero value otherwise.

### GetRestTimeOk

`func (o *RouteReport) GetRestTimeOk() (*int32, bool)`

GetRestTimeOk returns a tuple with the RestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTime

`func (o *RouteReport) SetRestTime(v int32)`

SetRestTime sets RestTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


