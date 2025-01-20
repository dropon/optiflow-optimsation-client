# OvertimeCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | **int32** | The threshold on the route duration [s] that defines when this overtime starts. | 
**ExtraPerHour** | **float64** | Specifies the extra cost for every hour above the threshold. If the threshold of multiple overtimes is exceeded, the extra costs per hour are added. | 

## Methods

### NewOvertimeCost

`func NewOvertimeCost(threshold int32, extraPerHour float64, ) *OvertimeCost`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


