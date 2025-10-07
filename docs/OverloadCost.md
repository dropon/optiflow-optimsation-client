# OverloadCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | **string** | Indicates the dimension of the load, to which the threshold applies. | 
**Threshold** | **float64** | The threshold for the specified load dimension.  When an order is loaded, only the part of that order which contributes to exceeding the threshold is considered for the extra cost. | 
**ExtraPerUnit** | **float64** | Specifies the extra cost per unit when loading an order which exceeds the threshold. | 

## Methods

### NewOverloadCost

`func NewOverloadCost(dimension string, threshold float64, extraPerUnit float64, ) *OverloadCost`

NewOverloadCost instantiates a new OverloadCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverloadCostWithDefaults

`func NewOverloadCostWithDefaults() *OverloadCost`

NewOverloadCostWithDefaults instantiates a new OverloadCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *OverloadCost) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *OverloadCost) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *OverloadCost) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetThreshold

`func (o *OverloadCost) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *OverloadCost) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *OverloadCost) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetExtraPerUnit

`func (o *OverloadCost) GetExtraPerUnit() float64`

GetExtraPerUnit returns the ExtraPerUnit field if non-nil, zero value otherwise.

### GetExtraPerUnitOk

`func (o *OverloadCost) GetExtraPerUnitOk() (*float64, bool)`

GetExtraPerUnitOk returns a tuple with the ExtraPerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPerUnit

`func (o *OverloadCost) SetExtraPerUnit(v float64)`

SetExtraPerUnit sets ExtraPerUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


