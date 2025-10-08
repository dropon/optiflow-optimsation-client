# OptimizationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | Defines the maximum duration [s] the optimization can use to reduce the cost of the scheduled routes. Please note that the upper bound on optimization duration is a technical limit. Check your individual price plan or contract to see which limits apply. The optimization will automatically stop when this duration is spent in the &#x60;RUNNING&#x60; status but can also be stopped manually using the *stopOptimization* endpoint. | 
**MaximumStagnationDuration** | Pointer to **int32** | Defines the maximum duration [s] the optimization may continue without improving the best-known solution. If no better solution is found within this period, the optimization will stop early - even if the overall duration limit has not yet been reached. This mechanism helps reduce unnecessary computation time when progress stalls. The stagnation timer starts when the optimization enters the &#x60;RUNNING&#x60; status and resets each time a better solution is found. *Use with caution:* This feature may prematurely terminate optimizations that could still yield better solutions given more time. Since the timing of improvements is inherently unpredictable, relying on this setting can lead to suboptimal results. It is recommended to disable this feature or set a conservative value unless you have strong reasons to limit runtime based on stagnation. | [optional] 

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


### GetMaximumStagnationDuration

`func (o *OptimizationSettings) GetMaximumStagnationDuration() int32`

GetMaximumStagnationDuration returns the MaximumStagnationDuration field if non-nil, zero value otherwise.

### GetMaximumStagnationDurationOk

`func (o *OptimizationSettings) GetMaximumStagnationDurationOk() (*int32, bool)`

GetMaximumStagnationDurationOk returns a tuple with the MaximumStagnationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumStagnationDuration

`func (o *OptimizationSettings) SetMaximumStagnationDuration(v int32)`

SetMaximumStagnationDuration sets MaximumStagnationDuration field to given value.

### HasMaximumStagnationDuration

`func (o *OptimizationSettings) HasMaximumStagnationDuration() bool`

HasMaximumStagnationDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


