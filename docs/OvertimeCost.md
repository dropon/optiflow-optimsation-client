# OvertimeCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | **int32** | The threshold for the route duration [s] above which the extra cost per hour applies. | 
**ExtraFixed** | Pointer to **float64** | Specifies the extra fixed cost when exceeding the threshold. | [optional] [default to 0]
**ExtraPerHour** | Pointer to **float64** | Specifies the extra cost for every hour above the threshold. | [optional] [default to 0]

## Methods

### NewOvertimeCost

`func NewOvertimeCost(threshold int32, ) *OvertimeCost`

NewOvertimeCost instantiates a new OvertimeCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOvertimeCostWithDefaults

`func NewOvertimeCostWithDefaults() *OvertimeCost`

NewOvertimeCostWithDefaults instantiates a new OvertimeCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *OvertimeCost) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *OvertimeCost) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *OvertimeCost) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetExtraFixed

`func (o *OvertimeCost) GetExtraFixed() float64`

GetExtraFixed returns the ExtraFixed field if non-nil, zero value otherwise.

### GetExtraFixedOk

`func (o *OvertimeCost) GetExtraFixedOk() (*float64, bool)`

GetExtraFixedOk returns a tuple with the ExtraFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFixed

`func (o *OvertimeCost) SetExtraFixed(v float64)`

SetExtraFixed sets ExtraFixed field to given value.

### HasExtraFixed

`func (o *OvertimeCost) HasExtraFixed() bool`

HasExtraFixed returns a boolean if a field has been set.

### GetExtraPerHour

`func (o *OvertimeCost) GetExtraPerHour() float64`

GetExtraPerHour returns the ExtraPerHour field if non-nil, zero value otherwise.

### GetExtraPerHourOk

`func (o *OvertimeCost) GetExtraPerHourOk() (*float64, bool)`

GetExtraPerHourOk returns a tuple with the ExtraPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPerHour

`func (o *OvertimeCost) SetExtraPerHour(v float64)`

SetExtraPerHour sets ExtraPerHour field to given value.

### HasExtraPerHour

`func (o *OvertimeCost) HasExtraPerHour() bool`

HasExtraPerHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


