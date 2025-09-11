# DurationModifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Factor** | Pointer to **float64** | The factor that should be applied to the original duration before any extra durations are added. If multiple rules apply, their factors are multiplied. | [optional] [default to 1]
**Extra** | Pointer to **int32** | The extra duration [s] that should be added. If multiple rules apply, their extra durations are added. | [optional] [default to 0]

## Methods

### NewDurationModifier

`func NewDurationModifier() *DurationModifier`

NewDurationModifier instantiates a new DurationModifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDurationModifierWithDefaults

`func NewDurationModifierWithDefaults() *DurationModifier`

NewDurationModifierWithDefaults instantiates a new DurationModifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactor

`func (o *DurationModifier) GetFactor() float64`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *DurationModifier) GetFactorOk() (*float64, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *DurationModifier) SetFactor(v float64)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *DurationModifier) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetExtra

`func (o *DurationModifier) GetExtra() int32`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *DurationModifier) GetExtraOk() (*int32, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *DurationModifier) SetExtra(v int32)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *DurationModifier) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


