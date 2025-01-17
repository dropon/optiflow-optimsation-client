# WayReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distance** | **int32** | The driving distance from the previous stop to this stop [m]. | 
**DrivingTime** | **int32** | The driving time from the previous stop to this stop [s]. | 
**WaitingTime** | **int32** | The sum of waiting times between the departure from the previous stop and the arrival at this stop [s]. | 
**BreakTime** | **int32** | The sum of break times between the departure from the previous stop and the arrival at this stop [s]. | 
**RestTime** | **int32** | The sum of rest times between the departure from the previous stop and the arrival at this stop [s]. | 

## Methods

### NewWayReport

`func NewWayReport(distance int32, drivingTime int32, waitingTime int32, breakTime int32, restTime int32, ) *WayReport`

NewWayReport instantiates a new WayReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWayReportWithDefaults

`func NewWayReportWithDefaults() *WayReport`

NewWayReportWithDefaults instantiates a new WayReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistance

`func (o *WayReport) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *WayReport) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *WayReport) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetDrivingTime

`func (o *WayReport) GetDrivingTime() int32`

GetDrivingTime returns the DrivingTime field if non-nil, zero value otherwise.

### GetDrivingTimeOk

`func (o *WayReport) GetDrivingTimeOk() (*int32, bool)`

GetDrivingTimeOk returns a tuple with the DrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingTime

`func (o *WayReport) SetDrivingTime(v int32)`

SetDrivingTime sets DrivingTime field to given value.


### GetWaitingTime

`func (o *WayReport) GetWaitingTime() int32`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *WayReport) GetWaitingTimeOk() (*int32, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *WayReport) SetWaitingTime(v int32)`

SetWaitingTime sets WaitingTime field to given value.


### GetBreakTime

`func (o *WayReport) GetBreakTime() int32`

GetBreakTime returns the BreakTime field if non-nil, zero value otherwise.

### GetBreakTimeOk

`func (o *WayReport) GetBreakTimeOk() (*int32, bool)`

GetBreakTimeOk returns a tuple with the BreakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakTime

`func (o *WayReport) SetBreakTime(v int32)`

SetBreakTime sets BreakTime field to given value.


### GetRestTime

`func (o *WayReport) GetRestTime() int32`

GetRestTime returns the RestTime field if non-nil, zero value otherwise.

### GetRestTimeOk

`func (o *WayReport) GetRestTimeOk() (*int32, bool)`

GetRestTimeOk returns a tuple with the RestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTime

`func (o *WayReport) SetRestTime(v int32)`

SetRestTime sets RestTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


