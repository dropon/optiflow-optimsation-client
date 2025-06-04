# OverdistanceCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | **int32** | The threshold for the route&#39;s total distance [m] above which the extra cost per kilometer applies. | 
**ExtraPerKilometer** | **float64** | Specifies the extra cost for every kilometer above the threshold. | 

## Methods

### NewOverdistanceCost

`func NewOverdistanceCost(threshold int32, extraPerKilometer float64, ) *OverdistanceCost`

NewOverdistanceCost instantiates a new OverdistanceCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverdistanceCostWithDefaults

`func NewOverdistanceCostWithDefaults() *OverdistanceCost`

NewOverdistanceCostWithDefaults instantiates a new OverdistanceCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *OverdistanceCost) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *OverdistanceCost) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *OverdistanceCost) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetExtraPerKilometer

`func (o *OverdistanceCost) GetExtraPerKilometer() float64`

GetExtraPerKilometer returns the ExtraPerKilometer field if non-nil, zero value otherwise.

### GetExtraPerKilometerOk

`func (o *OverdistanceCost) GetExtraPerKilometerOk() (*float64, bool)`

GetExtraPerKilometerOk returns a tuple with the ExtraPerKilometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPerKilometer

`func (o *OverdistanceCost) SetExtraPerKilometer(v float64)`

SetExtraPerKilometer sets ExtraPerKilometer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


