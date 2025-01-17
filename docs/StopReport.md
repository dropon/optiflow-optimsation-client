# StopReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrivalTime** | **time.Time** | The planned time of arrival at the stop formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). | 
**DepartureTime** | **time.Time** | The planned time of departure from the stop formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). | 
**ServiceTime** | **int32** | The service time at the stop [s]. | 
**WaitingTime** | **int32** | The waiting time at the stop [s]. | 
**BreakTime** | **int32** | The break time at the stop [s]. | 
**RestTime** | **int32** | The rest time at the stop [s]. | 
**Quantities** | Pointer to **[]int32** | The quantities loaded on the vehicle when leaving the stop. | [optional] 
**AlternativeCapacitiesIndex** | Pointer to **NullableInt32** | This field is only returned if the vehicle can transport the **quantities** loaded on the vehicle when leaving the stop with the given **capacities**. In this case this field contains the index of the chosen **alternativeCapacities**. If the route has at least one &#x60;VEHICLE_CAPACITY&#x60; violation this field will not be set.  See [here](./concepts/capacities-and-alternative-capacities) for more information. | [optional] 

## Methods

### NewStopReport

`func NewStopReport(arrivalTime time.Time, departureTime time.Time, serviceTime int32, waitingTime int32, breakTime int32, restTime int32, ) *StopReport`

NewStopReport instantiates a new StopReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopReportWithDefaults

`func NewStopReportWithDefaults() *StopReport`

NewStopReportWithDefaults instantiates a new StopReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrivalTime

`func (o *StopReport) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *StopReport) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *StopReport) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.


### GetDepartureTime

`func (o *StopReport) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *StopReport) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *StopReport) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.


### GetServiceTime

`func (o *StopReport) GetServiceTime() int32`

GetServiceTime returns the ServiceTime field if non-nil, zero value otherwise.

### GetServiceTimeOk

`func (o *StopReport) GetServiceTimeOk() (*int32, bool)`

GetServiceTimeOk returns a tuple with the ServiceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTime

`func (o *StopReport) SetServiceTime(v int32)`

SetServiceTime sets ServiceTime field to given value.


### GetWaitingTime

`func (o *StopReport) GetWaitingTime() int32`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *StopReport) GetWaitingTimeOk() (*int32, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *StopReport) SetWaitingTime(v int32)`

SetWaitingTime sets WaitingTime field to given value.


### GetBreakTime

`func (o *StopReport) GetBreakTime() int32`

GetBreakTime returns the BreakTime field if non-nil, zero value otherwise.

### GetBreakTimeOk

`func (o *StopReport) GetBreakTimeOk() (*int32, bool)`

GetBreakTimeOk returns a tuple with the BreakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakTime

`func (o *StopReport) SetBreakTime(v int32)`

SetBreakTime sets BreakTime field to given value.


### GetRestTime

`func (o *StopReport) GetRestTime() int32`

GetRestTime returns the RestTime field if non-nil, zero value otherwise.

### GetRestTimeOk

`func (o *StopReport) GetRestTimeOk() (*int32, bool)`

GetRestTimeOk returns a tuple with the RestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTime

`func (o *StopReport) SetRestTime(v int32)`

SetRestTime sets RestTime field to given value.


### GetQuantities

`func (o *StopReport) GetQuantities() []int32`

GetQuantities returns the Quantities field if non-nil, zero value otherwise.

### GetQuantitiesOk

`func (o *StopReport) GetQuantitiesOk() (*[]int32, bool)`

GetQuantitiesOk returns a tuple with the Quantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantities

`func (o *StopReport) SetQuantities(v []int32)`

SetQuantities sets Quantities field to given value.

### HasQuantities

`func (o *StopReport) HasQuantities() bool`

HasQuantities returns a boolean if a field has been set.

### GetAlternativeCapacitiesIndex

`func (o *StopReport) GetAlternativeCapacitiesIndex() int32`

GetAlternativeCapacitiesIndex returns the AlternativeCapacitiesIndex field if non-nil, zero value otherwise.

### GetAlternativeCapacitiesIndexOk

`func (o *StopReport) GetAlternativeCapacitiesIndexOk() (*int32, bool)`

GetAlternativeCapacitiesIndexOk returns a tuple with the AlternativeCapacitiesIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeCapacitiesIndex

`func (o *StopReport) SetAlternativeCapacitiesIndex(v int32)`

SetAlternativeCapacitiesIndex sets AlternativeCapacitiesIndex field to given value.

### HasAlternativeCapacitiesIndex

`func (o *StopReport) HasAlternativeCapacitiesIndex() bool`

HasAlternativeCapacitiesIndex returns a boolean if a field has been set.

### SetAlternativeCapacitiesIndexNil

`func (o *StopReport) SetAlternativeCapacitiesIndexNil(b bool)`

 SetAlternativeCapacitiesIndexNil sets the value for AlternativeCapacitiesIndex to be an explicit nil

### UnsetAlternativeCapacitiesIndex
`func (o *StopReport) UnsetAlternativeCapacitiesIndex()`

UnsetAlternativeCapacitiesIndex ensures that no value is present for AlternativeCapacitiesIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


