# OptimizationResultMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the optimization. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags of the optimization. | [optional] 
**Priority** | Pointer to **int32** | The priority level of the optimization. Higher values indicate higher priority. In the case of queuing optimizations, the higher priority optimizations are taken off the queue first. | [optional] 
**Created** | Pointer to **time.Time** | The creation time of the optimization. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 

## Methods

### NewOptimizationResultMetadata

`func NewOptimizationResultMetadata() *OptimizationResultMetadata`

NewOptimizationResultMetadata instantiates a new OptimizationResultMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationResultMetadataWithDefaults

`func NewOptimizationResultMetadataWithDefaults() *OptimizationResultMetadata`

NewOptimizationResultMetadataWithDefaults instantiates a new OptimizationResultMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OptimizationResultMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptimizationResultMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptimizationResultMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptimizationResultMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *OptimizationResultMetadata) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OptimizationResultMetadata) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OptimizationResultMetadata) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OptimizationResultMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPriority

`func (o *OptimizationResultMetadata) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OptimizationResultMetadata) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OptimizationResultMetadata) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OptimizationResultMetadata) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCreated

`func (o *OptimizationResultMetadata) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OptimizationResultMetadata) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OptimizationResultMetadata) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OptimizationResultMetadata) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


