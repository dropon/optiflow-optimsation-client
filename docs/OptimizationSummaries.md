# OptimizationSummaries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Optimizations** | Pointer to [**[]OptimizationSummary**](OptimizationSummary.md) | A list of optimization summaries. | [optional] 
**NextCursor** | Pointer to **string** | Pagination cursor for the next page of results. Present when additional results exist beyond the current page. Pass this value as the &#x60;cursor&#x60; parameter in the next request to continue pagination. Null indicates the final page. | [optional] 

## Methods

### NewOptimizationSummaries

`func NewOptimizationSummaries() *OptimizationSummaries`

NewOptimizationSummaries instantiates a new OptimizationSummaries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationSummariesWithDefaults

`func NewOptimizationSummariesWithDefaults() *OptimizationSummaries`

NewOptimizationSummariesWithDefaults instantiates a new OptimizationSummaries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptimizations

`func (o *OptimizationSummaries) GetOptimizations() []OptimizationSummary`

GetOptimizations returns the Optimizations field if non-nil, zero value otherwise.

### GetOptimizationsOk

`func (o *OptimizationSummaries) GetOptimizationsOk() (*[]OptimizationSummary, bool)`

GetOptimizationsOk returns a tuple with the Optimizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizations

`func (o *OptimizationSummaries) SetOptimizations(v []OptimizationSummary)`

SetOptimizations sets Optimizations field to given value.

### HasOptimizations

`func (o *OptimizationSummaries) HasOptimizations() bool`

HasOptimizations returns a boolean if a field has been set.

### GetNextCursor

`func (o *OptimizationSummaries) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OptimizationSummaries) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OptimizationSummaries) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *OptimizationSummaries) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


