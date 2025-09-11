# OptimizationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | Defines the maximum duration [s] the optimization can use to reduce the cost of the scheduled routes. Please note that the upper bound on optimization duration is a technical limit. Check your individual price plan or contract to see which limits apply. The optimization will automatically stop when this duration is spent in the &#x60;RUNNING&#x60; status but can also be stopped manually using the *stopOptimization* endpoint. | 

## Methods

### NewOptimizationSettings

`func NewOptimizationSettings(duration int32, ) *OptimizationSettings`

NewOptimizationSettings instantiates a new OptimizationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationSettingsWithDefaults

`func NewOptimizationSettingsWithDefaults() *OptimizationSettings`

NewOptimizationSettingsWithDefaults instantiates a new OptimizationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *OptimizationSettings) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OptimizationSettings) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OptimizationSettings) SetDuration(v int32)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


