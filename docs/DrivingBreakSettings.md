# DrivingBreakSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumDrivingDuration** | **int32** | Describes how long [s] the driver may drive without taking a break of at least the specified duration. | 
**MinimumBreakDuration** | Pointer to **int32** | **Deprecated, instead use &#x60;minimumBreakDurations&#x60; by specifying a list containing only this minimum break duration.** Specifies the duration [s] of a break a driver has to take if they exceed the maximum driving duration. | [optional] 
**MinimumBreakDurations** | Pointer to **[]int32** | Specifies the durations [s] of the breaks a driver has to take before exceeding the maximum driving duration. The breaks must be taken in the order provided in this list but consecutive elements in this list may be combined into single breaks. For example, when specifying &#x60;[300, 900, 1200]&#x60;, the possible break configurations are &#x60;[300, 900, 1200]&#x60;, &#x60;[300 + 900 &#x3D; 1200, 1200]&#x60;, &#x60;[300, 900 + 1200 &#x3D; 2100]&#x60; and a single break &#x60;[300 + 900 + 1200 &#x3D; 2400]&#x60;. | [optional] 

## Methods

### NewDrivingBreakSettings

`func NewDrivingBreakSettings(maximumDrivingDuration int32, ) *DrivingBreakSettings`

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

### HasMinimumBreakDuration

`func (o *DrivingBreakSettings) HasMinimumBreakDuration() bool`

HasMinimumBreakDuration returns a boolean if a field has been set.

### GetMinimumBreakDurations

`func (o *DrivingBreakSettings) GetMinimumBreakDurations() []int32`

GetMinimumBreakDurations returns the MinimumBreakDurations field if non-nil, zero value otherwise.

### GetMinimumBreakDurationsOk

`func (o *DrivingBreakSettings) GetMinimumBreakDurationsOk() (*[]int32, bool)`

GetMinimumBreakDurationsOk returns a tuple with the MinimumBreakDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBreakDurations

`func (o *DrivingBreakSettings) SetMinimumBreakDurations(v []int32)`

SetMinimumBreakDurations sets MinimumBreakDurations field to given value.

### HasMinimumBreakDurations

`func (o *DrivingBreakSettings) HasMinimumBreakDurations() bool`

HasMinimumBreakDurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


