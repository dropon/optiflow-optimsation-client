# Driver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the driver. | 
**VehicleId** | **string** | The ID of the driver&#39;s vehicle. This vehicle must not be referenced by another driver. | 
**Availabilities** | Pointer to [**[]TimeInterval**](TimeInterval.md) | Intervals during which the driver is available, each specified by two points in time - the beginning and the end of the interval. Each trip must lie completely within one of the intervals. The route start and trip start events must lie within one of the intervals. The intervals must have a gap of more than 1 second. Leaving this parameter empty means that the driver is always available. | [optional] 
**WorkingHoursPreset** | Pointer to [**NullableWorkingHoursPreset**](WorkingHoursPreset.md) |  | [optional] 
**BreakRule** | Pointer to [**BreakRule**](BreakRule.md) |  | [optional] 
**DailyRestRule** | Pointer to [**DailyRestRule**](DailyRestRule.md) |  | [optional] 
**WorkLogbook** | Pointer to [**WorkLogbook**](WorkLogbook.md) |  | [optional] 
**MaximumDrivingTime** | Pointer to **NullableInt32** | The maximum driving time of the driver [s].   This includes the driving time before the start of the route (see **accumulatedDrivingTimeSinceLastDailyRest** in **workLogbook**). The maximum driving time is considered as infinite if it is not set. Currently, each driver must have the same value specified. | [optional] 
**MaximumTravelTime** | Pointer to **NullableInt32** | The maximum travel time of the driver [s]. The travel time contains all waiting, service and driving times.  This includes the travel time before the start of the route (see **accumulatedTravelTimeSinceLastDailyRest** in **workLogbook**). The maximum travel time is considered as infinite if it is not set. Currently, each driver must have the same value specified. | [optional] 

## Methods

### NewDriver

`func NewDriver(id string, vehicleId string, ) *Driver`

NewDriver instantiates a new Driver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverWithDefaults

`func NewDriverWithDefaults() *Driver`

NewDriverWithDefaults instantiates a new Driver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Driver) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Driver) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Driver) SetId(v string)`

SetId sets Id field to given value.


### GetVehicleId

`func (o *Driver) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *Driver) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *Driver) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.


### GetAvailabilities

`func (o *Driver) GetAvailabilities() []TimeInterval`

GetAvailabilities returns the Availabilities field if non-nil, zero value otherwise.

### GetAvailabilitiesOk

`func (o *Driver) GetAvailabilitiesOk() (*[]TimeInterval, bool)`

GetAvailabilitiesOk returns a tuple with the Availabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilities

`func (o *Driver) SetAvailabilities(v []TimeInterval)`

SetAvailabilities sets Availabilities field to given value.

### HasAvailabilities

`func (o *Driver) HasAvailabilities() bool`

HasAvailabilities returns a boolean if a field has been set.

### GetWorkingHoursPreset

`func (o *Driver) GetWorkingHoursPreset() WorkingHoursPreset`

GetWorkingHoursPreset returns the WorkingHoursPreset field if non-nil, zero value otherwise.

### GetWorkingHoursPresetOk

`func (o *Driver) GetWorkingHoursPresetOk() (*WorkingHoursPreset, bool)`

GetWorkingHoursPresetOk returns a tuple with the WorkingHoursPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHoursPreset

`func (o *Driver) SetWorkingHoursPreset(v WorkingHoursPreset)`

SetWorkingHoursPreset sets WorkingHoursPreset field to given value.

### HasWorkingHoursPreset

`func (o *Driver) HasWorkingHoursPreset() bool`

HasWorkingHoursPreset returns a boolean if a field has been set.

### SetWorkingHoursPresetNil

`func (o *Driver) SetWorkingHoursPresetNil(b bool)`

 SetWorkingHoursPresetNil sets the value for WorkingHoursPreset to be an explicit nil

### UnsetWorkingHoursPreset
`func (o *Driver) UnsetWorkingHoursPreset()`

UnsetWorkingHoursPreset ensures that no value is present for WorkingHoursPreset, not even an explicit nil
### GetBreakRule

`func (o *Driver) GetBreakRule() BreakRule`

GetBreakRule returns the BreakRule field if non-nil, zero value otherwise.

### GetBreakRuleOk

`func (o *Driver) GetBreakRuleOk() (*BreakRule, bool)`

GetBreakRuleOk returns a tuple with the BreakRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakRule

`func (o *Driver) SetBreakRule(v BreakRule)`

SetBreakRule sets BreakRule field to given value.

### HasBreakRule

`func (o *Driver) HasBreakRule() bool`

HasBreakRule returns a boolean if a field has been set.

### GetDailyRestRule

`func (o *Driver) GetDailyRestRule() DailyRestRule`

GetDailyRestRule returns the DailyRestRule field if non-nil, zero value otherwise.

### GetDailyRestRuleOk

`func (o *Driver) GetDailyRestRuleOk() (*DailyRestRule, bool)`

GetDailyRestRuleOk returns a tuple with the DailyRestRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyRestRule

`func (o *Driver) SetDailyRestRule(v DailyRestRule)`

SetDailyRestRule sets DailyRestRule field to given value.

### HasDailyRestRule

`func (o *Driver) HasDailyRestRule() bool`

HasDailyRestRule returns a boolean if a field has been set.

### GetWorkLogbook

`func (o *Driver) GetWorkLogbook() WorkLogbook`

GetWorkLogbook returns the WorkLogbook field if non-nil, zero value otherwise.

### GetWorkLogbookOk

`func (o *Driver) GetWorkLogbookOk() (*WorkLogbook, bool)`

GetWorkLogbookOk returns a tuple with the WorkLogbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLogbook

`func (o *Driver) SetWorkLogbook(v WorkLogbook)`

SetWorkLogbook sets WorkLogbook field to given value.

### HasWorkLogbook

`func (o *Driver) HasWorkLogbook() bool`

HasWorkLogbook returns a boolean if a field has been set.

### GetMaximumDrivingTime

`func (o *Driver) GetMaximumDrivingTime() int32`

GetMaximumDrivingTime returns the MaximumDrivingTime field if non-nil, zero value otherwise.

### GetMaximumDrivingTimeOk

`func (o *Driver) GetMaximumDrivingTimeOk() (*int32, bool)`

GetMaximumDrivingTimeOk returns a tuple with the MaximumDrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDrivingTime

`func (o *Driver) SetMaximumDrivingTime(v int32)`

SetMaximumDrivingTime sets MaximumDrivingTime field to given value.

### HasMaximumDrivingTime

`func (o *Driver) HasMaximumDrivingTime() bool`

HasMaximumDrivingTime returns a boolean if a field has been set.

### SetMaximumDrivingTimeNil

`func (o *Driver) SetMaximumDrivingTimeNil(b bool)`

 SetMaximumDrivingTimeNil sets the value for MaximumDrivingTime to be an explicit nil

### UnsetMaximumDrivingTime
`func (o *Driver) UnsetMaximumDrivingTime()`

UnsetMaximumDrivingTime ensures that no value is present for MaximumDrivingTime, not even an explicit nil
### GetMaximumTravelTime

`func (o *Driver) GetMaximumTravelTime() int32`

GetMaximumTravelTime returns the MaximumTravelTime field if non-nil, zero value otherwise.

### GetMaximumTravelTimeOk

`func (o *Driver) GetMaximumTravelTimeOk() (*int32, bool)`

GetMaximumTravelTimeOk returns a tuple with the MaximumTravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumTravelTime

`func (o *Driver) SetMaximumTravelTime(v int32)`

SetMaximumTravelTime sets MaximumTravelTime field to given value.

### HasMaximumTravelTime

`func (o *Driver) HasMaximumTravelTime() bool`

HasMaximumTravelTime returns a boolean if a field has been set.

### SetMaximumTravelTimeNil

`func (o *Driver) SetMaximumTravelTimeNil(b bool)`

 SetMaximumTravelTimeNil sets the value for MaximumTravelTime to be an explicit nil

### UnsetMaximumTravelTime
`func (o *Driver) UnsetMaximumTravelTime()`

UnsetMaximumTravelTime ensures that no value is present for MaximumTravelTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


