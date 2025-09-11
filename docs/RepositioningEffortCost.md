# RepositioningEffortCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | **int32** | The threshold for the repositioning effort from which the cost per effort applies. | 
**ExtraPerEffort** | **float64** | Specifies the extra cost for every unit of effort above the threshold. If the threshold of multiple repositioning effort costs is exceeded, the extra costs per effort are added. | 

## Methods

### NewRepositioningEffortCost

`func NewRepositioningEffortCost(threshold int32, extraPerEffort float64, ) *RepositioningEffortCost`

NewRepositioningEffortCost instantiates a new RepositioningEffortCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositioningEffortCostWithDefaults

`func NewRepositioningEffortCostWithDefaults() *RepositioningEffortCost`

NewRepositioningEffortCostWithDefaults instantiates a new RepositioningEffortCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *RepositioningEffortCost) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RepositioningEffortCost) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RepositioningEffortCost) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetExtraPerEffort

`func (o *RepositioningEffortCost) GetExtraPerEffort() float64`

GetExtraPerEffort returns the ExtraPerEffort field if non-nil, zero value otherwise.

### GetExtraPerEffortOk

`func (o *RepositioningEffortCost) GetExtraPerEffortOk() (*float64, bool)`

GetExtraPerEffortOk returns a tuple with the ExtraPerEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPerEffort

`func (o *RepositioningEffortCost) SetExtraPerEffort(v float64)`

SetExtraPerEffort sets ExtraPerEffort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


