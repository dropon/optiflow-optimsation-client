# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the operation, possible values: \&quot;optimization\&quot;, \&quot;evaluation\&quot;. | 
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**StartTime** | **time.Time** | The start time of the operation formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339), by means the client can decide to delete the operation. The value is always returned in UTC time. | 
**ElapsedTime** | **int32** | The elapsed time of the operation [s], by means the client can decide to delete the operation. It represents the duration from the start time until now when the operation is still active or from the start time until the end of operation when this operation is already completed. | 
**Error** | Pointer to [**ErrorResponse**](ErrorResponse.md) |  | [optional] 

## Methods

### NewOperation

`func NewOperation(name string, status OperationStatus, startTime time.Time, elapsedTime int32, ) *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Operation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Operation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Operation) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Operation) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Operation) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Operation) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetStartTime

`func (o *Operation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Operation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Operation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetElapsedTime

`func (o *Operation) GetElapsedTime() int32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *Operation) GetElapsedTimeOk() (*int32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *Operation) SetElapsedTime(v int32)`

SetElapsedTime sets ElapsedTime field to given value.


### GetError

`func (o *Operation) GetError() ErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Operation) GetErrorOk() (*ErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Operation) SetError(v ErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *Operation) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


