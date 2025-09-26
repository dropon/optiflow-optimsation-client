# OverdistanceCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | **int32** | The threshold for the route&#39;s total distance [m] above which the extra costs apply. | 
**ExtraFixed** | Pointer to **float64** | Specifies the extra fixed cost when exceeding the threshold. | [optional] [default to 0]
**ExtraPerKilometer** | Pointer to **float64** | Specifies the extra cost for every kilometer above the threshold. | [optional] [default to 0]

## Methods

### NewOverdistanceCost

`func NewOverdistanceCost(threshold int32, ) *OverdistanceCost`

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


### GetExtraFixed

`func (o *OverdistanceCost) GetExtraFixed() float64`

GetExtraFixed returns the ExtraFixed field if non-nil, zero value otherwise.

### GetExtraFixedOk

`func (o *OverdistanceCost) GetExtraFixedOk() (*float64, bool)`

GetExtraFixedOk returns a tuple with the ExtraFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFixed

`func (o *OverdistanceCost) SetExtraFixed(v float64)`

SetExtraFixed sets ExtraFixed field to given value.

### HasExtraFixed

`func (o *OverdistanceCost) HasExtraFixed() bool`

HasExtraFixed returns a boolean if a field has been set.

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

### HasExtraPerKilometer

`func (o *OverdistanceCost) HasExtraPerKilometer() bool`

HasExtraPerKilometer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


