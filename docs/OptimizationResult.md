# OptimizationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the optimization. | 
**Metadata** | Pointer to [**OptimizationResultMetadata**](OptimizationResultMetadata.md) |  | [optional] 
**Status** | [**OptimizationStatus**](OptimizationStatus.md) |  | 
**Routes** | Pointer to [**[]Route**](Route.md) | The routes scheduled by the optimization. As long as the optimization is not yet &#x60;SUCCEEDED&#x60;, scheduled routes may be only an intermediate result. | [optional] [default to {}]
**Metrics** | Pointer to [**Metrics**](Metrics.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**Warnings** | Pointer to [**[]Warning**](Warning.md) | A list of warnings concerning the optimization. | [optional] [default to {}]

## Methods

### NewOptimizationResult

`func NewOptimizationResult(id string, status OptimizationStatus, ) *OptimizationResult`

NewOptimizationResult instantiates a new OptimizationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationResultWithDefaults

`func NewOptimizationResultWithDefaults() *OptimizationResult`

NewOptimizationResultWithDefaults instantiates a new OptimizationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptimizationResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptimizationResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptimizationResult) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *OptimizationResult) GetMetadata() OptimizationResultMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OptimizationResult) GetMetadataOk() (*OptimizationResultMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OptimizationResult) SetMetadata(v OptimizationResultMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OptimizationResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *OptimizationResult) GetStatus() OptimizationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptimizationResult) GetStatusOk() (*OptimizationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptimizationResult) SetStatus(v OptimizationStatus)`

SetStatus sets Status field to given value.


### GetRoutes

`func (o *OptimizationResult) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *OptimizationResult) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *OptimizationResult) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *OptimizationResult) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetMetrics

`func (o *OptimizationResult) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *OptimizationResult) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *OptimizationResult) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *OptimizationResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetError

`func (o *OptimizationResult) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OptimizationResult) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OptimizationResult) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *OptimizationResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWarnings

`func (o *OptimizationResult) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *OptimizationResult) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *OptimizationResult) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *OptimizationResult) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


