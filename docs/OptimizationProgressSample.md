# OptimizationProgressSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | The time at which the metrics were captured. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**Metrics** | Pointer to [**Metrics**](Metrics.md) |  | [optional] 

## Methods

### NewOptimizationProgressSample

`func NewOptimizationProgressSample() *OptimizationProgressSample`

NewOptimizationProgressSample instantiates a new OptimizationProgressSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationProgressSampleWithDefaults

`func NewOptimizationProgressSampleWithDefaults() *OptimizationProgressSample`

NewOptimizationProgressSampleWithDefaults instantiates a new OptimizationProgressSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *OptimizationProgressSample) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OptimizationProgressSample) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OptimizationProgressSample) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *OptimizationProgressSample) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetMetrics

`func (o *OptimizationProgressSample) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *OptimizationProgressSample) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *OptimizationProgressSample) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *OptimizationProgressSample) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


