# WorkingBreakSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumWorkingDuration** | **int32** | Describes how long [s] the driver may work without taking a break of at least the specified duration. | 
**MinimumBreakDuration** | Pointer to **int32** | **Deprecated, instead use &#x60;minimumBreakDurations&#x60; by specifying a list containing only this minimum break duration.** Specifies the duration [s] of a break a driver must take if they exceed the maximum working duration. | [optional] 
**MinimumBreakDurations** | Pointer to **[]int32** | Specifies the durations [s] of the breaks a driver has to take before exceeding the maximum working duration. The breaks must be taken in the order provided in this list but consecutive elements in this list may be combined into single breaks. For example, when specifying &#x60;[300, 900, 1200]&#x60;, the possible break configurations are &#x60;[300, 900, 1200]&#x60;, &#x60;[300 + 900 &#x3D; 1200, 1200]&#x60;, &#x60;[300, 900 + 1200 &#x3D; 2100]&#x60; and a single break &#x60;[300 + 900 + 1200 &#x3D; 2400]&#x60;. | [optional] 

## Methods

### NewWorkingBreakSettings

`func NewWorkingBreakSettings(maximumWorkingDuration int32, ) *WorkingBreakSettings`

NewWorkingBreakSettings instantiates a new WorkingBreakSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkingBreakSettingsWithDefaults

`func NewWorkingBreakSettingsWithDefaults() *WorkingBreakSettings`

NewWorkingBreakSettingsWithDefaults instantiates a new WorkingBreakSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumWorkingDuration

`func (o *WorkingBreakSettings) GetMaximumWorkingDuration() int32`

GetMaximumWorkingDuration returns the MaximumWorkingDuration field if non-nil, zero value otherwise.

### GetMaximumWorkingDurationOk

`func (o *WorkingBreakSettings) GetMaximumWorkingDurationOk() (*int32, bool)`

GetMaximumWorkingDurationOk returns a tuple with the MaximumWorkingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumWorkingDuration

`func (o *WorkingBreakSettings) SetMaximumWorkingDuration(v int32)`

SetMaximumWorkingDuration sets MaximumWorkingDuration field to given value.


### GetMinimumBreakDuration

`func (o *WorkingBreakSettings) GetMinimumBreakDuration() int32`

GetMinimumBreakDuration returns the MinimumBreakDuration field if non-nil, zero value otherwise.

### GetMinimumBreakDurationOk

`func (o *WorkingBreakSettings) GetMinimumBreakDurationOk() (*int32, bool)`

GetMinimumBreakDurationOk returns a tuple with the MinimumBreakDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBreakDuration

`func (o *WorkingBreakSettings) SetMinimumBreakDuration(v int32)`

SetMinimumBreakDuration sets MinimumBreakDuration field to given value.

### HasMinimumBreakDuration

`func (o *WorkingBreakSettings) HasMinimumBreakDuration() bool`

HasMinimumBreakDuration returns a boolean if a field has been set.

### GetMinimumBreakDurations

`func (o *WorkingBreakSettings) GetMinimumBreakDurations() []int32`

GetMinimumBreakDurations returns the MinimumBreakDurations field if non-nil, zero value otherwise.

### GetMinimumBreakDurationsOk

`func (o *WorkingBreakSettings) GetMinimumBreakDurationsOk() (*[]int32, bool)`

GetMinimumBreakDurationsOk returns a tuple with the MinimumBreakDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBreakDurations

`func (o *WorkingBreakSettings) SetMinimumBreakDurations(v []int32)`

SetMinimumBreakDurations sets MinimumBreakDurations field to given value.

### HasMinimumBreakDurations

`func (o *WorkingBreakSettings) HasMinimumBreakDurations() bool`

HasMinimumBreakDurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


