# DrivingBreakSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumDrivingDuration** | **int32** | Describes how long [s] the driver may drive without taking a break of at least the specified duration. | 
**MinimumBreakDuration** | **int32** | Specifies the duration [s] of a break a driver has to take if they exceed the maximum driving duration. | 

## Methods

### NewDrivingBreakSettings

`func NewDrivingBreakSettings(maximumDrivingDuration int32, minimumBreakDuration int32, ) *DrivingBreakSettings`

NewDrivingBreakSettings instantiates a new DrivingBreakSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrivingBreakSettingsWithDefaults

`func NewDrivingBreakSettingsWithDefaults() *DrivingBreakSettings`

NewDrivingBreakSettingsWithDefaults instantiates a new DrivingBreakSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumDrivingDuration

`func (o *DrivingBreakSettings) GetMaximumDrivingDuration() int32`

GetMaximumDrivingDuration returns the MaximumDrivingDuration field if non-nil, zero value otherwise.

### GetMaximumDrivingDurationOk

`func (o *DrivingBreakSettings) GetMaximumDrivingDurationOk() (*int32, bool)`

GetMaximumDrivingDurationOk returns a tuple with the MaximumDrivingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDrivingDuration

`func (o *DrivingBreakSettings) SetMaximumDrivingDuration(v int32)`

SetMaximumDrivingDuration sets MaximumDrivingDuration field to given value.


### GetMinimumBreakDuration

`func (o *DrivingBreakSettings) GetMinimumBreakDuration() int32`

GetMinimumBreakDuration returns the MinimumBreakDuration field if non-nil, zero value otherwise.

### GetMinimumBreakDurationOk

`func (o *DrivingBreakSettings) GetMinimumBreakDurationOk() (*int32, bool)`

GetMinimumBreakDurationOk returns a tuple with the MinimumBreakDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBreakDuration

`func (o *DrivingBreakSettings) SetMinimumBreakDuration(v int32)`

SetMinimumBreakDuration sets MinimumBreakDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


