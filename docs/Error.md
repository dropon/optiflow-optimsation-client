# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable message that describes the error. | 
**ErrorCode** | **string** | A constant string that can be used to identify this error class programmatically. An errorCode can have **details** to provide information in additional properties which are described with the code they apply to. They are of type string unless otherwise specified. Note that additional errorCodes as well as the **details** of existing errorCodes may be added at any time. Furthermore, the **description** may change at any time. * &#x60;ROUTE_OPTIMIZATION_UNEXPECTED_ERROR&#x60; - An unexpected error occurred. * &#x60;ROUTE_OPTIMIZATION_NOTHING_TO_OPTIMIZE&#x60; - No order can be scheduled on any route of the provided vehicles. The optimization stopped because there is nothing to optimize. * &#x60;ROUTE_OPTIMIZATION_STOPPED_UNEXPECTEDLY&#x60; - The optimization stopped unexpectedly. The provided routes are from the last known optimization result. * &#x60;ROUTE_OPTIMIZATION_STOPPED_WHILE_PREPARING&#x60; - The optimization has been stopped while it was preparing. No routes will be provided. * &#x60;ROUTE_OPTIMIZATION_START_LOCATION_UNREACHABLE&#x60; - Travel times and distances cannot be calculated to or from the start location of a vehicle.   * &#x60;vehicleId&#x60; - The vehicle ID for which time-distance information could not be retrieved. * &#x60;ROUTE_OPTIMIZATION_END_LOCATION_UNREACHABLE&#x60; - Travel times and distances cannot be calculated to or from the end location of a vehicle.   * &#x60;vehicleId&#x60; - The vehicle ID for which time-distance information could not be retrieved. | 
**Details** | Pointer to **map[string]interface{}** | Additional properties specific to this class of errors. | [optional] 

## Methods

### NewError

`func NewError(description string, errorCode string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Error) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Error) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Error) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetErrorCode

`func (o *Error) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Error) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Error) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetDetails

`func (o *Error) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Error) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Error) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Error) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


