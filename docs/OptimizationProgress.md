# OptimizationProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to [**[]OptimizationProgressSample**](OptimizationProgressSample.md) | A list of samples that describe the optimization metrics at various points in time. | [optional] 

## Methods

### NewOptimizationProgress

`func NewOptimizationProgress() *OptimizationProgress`

NewOptimizationProgress instantiates a new OptimizationProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationProgressWithDefaults

`func NewOptimizationProgressWithDefaults() *OptimizationProgress`

NewOptimizationProgressWithDefaults instantiates a new OptimizationProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *OptimizationProgress) GetSamples() []OptimizationProgressSample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OptimizationProgress) GetSamplesOk() (*[]OptimizationProgressSample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OptimizationProgress) SetSamples(v []OptimizationProgressSample)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OptimizationProgress) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


