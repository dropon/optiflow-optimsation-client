# BreakRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakTime** | **int32** | The minimum duration of a break [s]. Minimum is 15 minutes, maximum is 1 hour 30 minutes. | 
**MaximumDrivingTimeBetweenBreaks** | Pointer to **NullableInt32** | Maximum duration that the driver is allowed to drive [s] before taking a break. The maximum driving time is considered as infinite if it is not set. | [optional] 
**MaximumWorkingTimeBetweenBreaks** | Pointer to **NullableInt32** | Maximum duration that the driver is allowed to work [s] before taking a break. The maximum working time is considered as infinite if it is not set. | [optional] 
**WorkingTimeThreshold** | Pointer to **int32** | Idle time of the driver counts as working time if it is shorter than this value [s]. May not be higher than **breakTime**. | [optional] [default to 0]

## Methods

### NewBreakRule

`func NewBreakRule(breakTime int32, ) *BreakRule`

NewBreakRule instantiates a new BreakRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakRuleWithDefaults

`func NewBreakRuleWithDefaults() *BreakRule`

NewBreakRuleWithDefaults instantiates a new BreakRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakTime

`func (o *BreakRule) GetBreakTime() int32`

GetBreakTime returns the BreakTime field if non-nil, zero value otherwise.

### GetBreakTimeOk

`func (o *BreakRule) GetBreakTimeOk() (*int32, bool)`

GetBreakTimeOk returns a tuple with the BreakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakTime

`func (o *BreakRule) SetBreakTime(v int32)`

SetBreakTime sets BreakTime field to given value.


### GetMaximumDrivingTimeBetweenBreaks

`func (o *BreakRule) GetMaximumDrivingTimeBetweenBreaks() int32`

GetMaximumDrivingTimeBetweenBreaks returns the MaximumDrivingTimeBetweenBreaks field if non-nil, zero value otherwise.

### GetMaximumDrivingTimeBetweenBreaksOk

`func (o *BreakRule) GetMaximumDrivingTimeBetweenBreaksOk() (*int32, bool)`

GetMaximumDrivingTimeBetweenBreaksOk returns a tuple with the MaximumDrivingTimeBetweenBreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDrivingTimeBetweenBreaks

`func (o *BreakRule) SetMaximumDrivingTimeBetweenBreaks(v int32)`

SetMaximumDrivingTimeBetweenBreaks sets MaximumDrivingTimeBetweenBreaks field to given value.

### HasMaximumDrivingTimeBetweenBreaks

`func (o *BreakRule) HasMaximumDrivingTimeBetweenBreaks() bool`

HasMaximumDrivingTimeBetweenBreaks returns a boolean if a field has been set.

### SetMaximumDrivingTimeBetweenBreaksNil

`func (o *BreakRule) SetMaximumDrivingTimeBetweenBreaksNil(b bool)`

 SetMaximumDrivingTimeBetweenBreaksNil sets the value for MaximumDrivingTimeBetweenBreaks to be an explicit nil

### UnsetMaximumDrivingTimeBetweenBreaks
`func (o *BreakRule) UnsetMaximumDrivingTimeBetweenBreaks()`

UnsetMaximumDrivingTimeBetweenBreaks ensures that no value is present for MaximumDrivingTimeBetweenBreaks, not even an explicit nil
### GetMaximumWorkingTimeBetweenBreaks

`func (o *BreakRule) GetMaximumWorkingTimeBetweenBreaks() int32`

GetMaximumWorkingTimeBetweenBreaks returns the MaximumWorkingTimeBetweenBreaks field if non-nil, zero value otherwise.

### GetMaximumWorkingTimeBetweenBreaksOk

`func (o *BreakRule) GetMaximumWorkingTimeBetweenBreaksOk() (*int32, bool)`

GetMaximumWorkingTimeBetweenBreaksOk returns a tuple with the MaximumWorkingTimeBetweenBreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumWorkingTimeBetweenBreaks

`func (o *BreakRule) SetMaximumWorkingTimeBetweenBreaks(v int32)`

SetMaximumWorkingTimeBetweenBreaks sets MaximumWorkingTimeBetweenBreaks field to given value.

### HasMaximumWorkingTimeBetweenBreaks

`func (o *BreakRule) HasMaximumWorkingTimeBetweenBreaks() bool`

HasMaximumWorkingTimeBetweenBreaks returns a boolean if a field has been set.

### SetMaximumWorkingTimeBetweenBreaksNil

`func (o *BreakRule) SetMaximumWorkingTimeBetweenBreaksNil(b bool)`

 SetMaximumWorkingTimeBetweenBreaksNil sets the value for MaximumWorkingTimeBetweenBreaks to be an explicit nil

### UnsetMaximumWorkingTimeBetweenBreaks
`func (o *BreakRule) UnsetMaximumWorkingTimeBetweenBreaks()`

UnsetMaximumWorkingTimeBetweenBreaks ensures that no value is present for MaximumWorkingTimeBetweenBreaks, not even an explicit nil
### GetWorkingTimeThreshold

`func (o *BreakRule) GetWorkingTimeThreshold() int32`

GetWorkingTimeThreshold returns the WorkingTimeThreshold field if non-nil, zero value otherwise.

### GetWorkingTimeThresholdOk

`func (o *BreakRule) GetWorkingTimeThresholdOk() (*int32, bool)`

GetWorkingTimeThresholdOk returns a tuple with the WorkingTimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeThreshold

`func (o *BreakRule) SetWorkingTimeThreshold(v int32)`

SetWorkingTimeThreshold sets WorkingTimeThreshold field to given value.

### HasWorkingTimeThreshold

`func (o *BreakRule) HasWorkingTimeThreshold() bool`

HasWorkingTimeThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


