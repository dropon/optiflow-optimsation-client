# OptimizationRequestMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the optimization. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags of the optimization. | [optional] 
**Priority** | Pointer to **int32** | The priority level of the optimization. Higher values indicate higher priority. In the case of queuing optimizations, the higher priority optimizations are taken off the queue first. | [optional] [default to 0]

## Methods

### NewOptimizationRequestMetadata

`func NewOptimizationRequestMetadata() *OptimizationRequestMetadata`

NewOptimizationRequestMetadata instantiates a new OptimizationRequestMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationRequestMetadataWithDefaults

`func NewOptimizationRequestMetadataWithDefaults() *OptimizationRequestMetadata`

NewOptimizationRequestMetadataWithDefaults instantiates a new OptimizationRequestMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OptimizationRequestMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptimizationRequestMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptimizationRequestMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptimizationRequestMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *OptimizationRequestMetadata) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OptimizationRequestMetadata) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OptimizationRequestMetadata) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OptimizationRequestMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPriority

`func (o *OptimizationRequestMetadata) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OptimizationRequestMetadata) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OptimizationRequestMetadata) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OptimizationRequestMetadata) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


