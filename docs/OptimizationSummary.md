# OptimizationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the optimization. | [optional] 
**Status** | Pointer to [**OptimizationStatus**](OptimizationStatus.md) |  | [optional] 
**Metrics** | Pointer to [**Metrics**](Metrics.md) |  | [optional] 
**Request** | Pointer to [**RequestSummary**](RequestSummary.md) |  | [optional] 
**Metadata** | Pointer to [**OptimizationResultMetadata**](OptimizationResultMetadata.md) |  | [optional] 

## Methods

### NewOptimizationSummary

`func NewOptimizationSummary() *OptimizationSummary`

NewOptimizationSummary instantiates a new OptimizationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationSummaryWithDefaults

`func NewOptimizationSummaryWithDefaults() *OptimizationSummary`

NewOptimizationSummaryWithDefaults instantiates a new OptimizationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptimizationSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptimizationSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptimizationSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OptimizationSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OptimizationSummary) GetStatus() OptimizationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptimizationSummary) GetStatusOk() (*OptimizationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptimizationSummary) SetStatus(v OptimizationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OptimizationSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetrics

`func (o *OptimizationSummary) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *OptimizationSummary) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *OptimizationSummary) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *OptimizationSummary) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetRequest

`func (o *OptimizationSummary) GetRequest() RequestSummary`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *OptimizationSummary) GetRequestOk() (*RequestSummary, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *OptimizationSummary) SetRequest(v RequestSummary)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *OptimizationSummary) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetMetadata

`func (o *OptimizationSummary) GetMetadata() OptimizationResultMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OptimizationSummary) GetMetadataOk() (*OptimizationResultMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OptimizationSummary) SetMetadata(v OptimizationResultMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OptimizationSummary) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


