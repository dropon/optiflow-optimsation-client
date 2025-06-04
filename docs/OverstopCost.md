# OverstopCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | **int32** | The threshold for the route&#39;s number of stops above which the extra cost per stop applies. | 
**ExtraPerStop** | **float64** | Specifies the extra cost for every stop above the threshold. | 

## Methods

### NewOverstopCost

`func NewOverstopCost(threshold int32, extraPerStop float64, ) *OverstopCost`

NewOverstopCost instantiates a new OverstopCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverstopCostWithDefaults

`func NewOverstopCostWithDefaults() *OverstopCost`

NewOverstopCostWithDefaults instantiates a new OverstopCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *OverstopCost) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *OverstopCost) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *OverstopCost) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetExtraPerStop

`func (o *OverstopCost) GetExtraPerStop() float64`

GetExtraPerStop returns the ExtraPerStop field if non-nil, zero value otherwise.

### GetExtraPerStopOk

`func (o *OverstopCost) GetExtraPerStopOk() (*float64, bool)`

GetExtraPerStopOk returns a tuple with the ExtraPerStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPerStop

`func (o *OverstopCost) SetExtraPerStop(v float64)`

SetExtraPerStop sets ExtraPerStop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


