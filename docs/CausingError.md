# CausingError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable message that describes the error. | 
**ErrorCode** | **string** | A constant string that can be used to identify this error class programmatically.  If additional information is available for an errorCode, it will be provided as key-value pairs with the parameter **details**. The keys available for a specific errorCode are documented directly with the errorCode. Unless stated otherwise, the values are of type string.  As an example, the following errorCode provides one key-value pair in the **details**. The key is called **value**. * &#x60;GENERAL_INVALID_VALUE&#x60; - A parameter is set to an invalid value.   * &#x60;value&#x60; - The invalid value.  Note that additional errorCodes as well as the **details** of existing errorCodes may be added at any time. Furthermore, the **description** may change at any time.  **Error codes for** &#x60;GENERAL_VALIDATION_ERROR&#x60; * &#x60;GENERAL_UNRECOGNIZED_PARAMETER&#x60; - A parameter is unknown. * &#x60;GENERAL_MISSING_PARAMETER&#x60; - A required parameter is missing. * &#x60;GENERAL_TYPE_VIOLATED&#x60; - The value of a parameter has an invalid type.   * &#x60;type&#x60; - The type. * &#x60;GENERAL_FORMAT_VIOLATED&#x60; - The value of a parameter has an invalid format.   * &#x60;format&#x60; - The format. * &#x60;GENERAL_PATTERN_VIOLATED&#x60; - The value of a string parameter does not satisfy the required pattern.   * &#x60;pattern&#x60; - The pattern. * &#x60;GENERAL_MINIMUM_LENGTH_VIOLATED&#x60; - The minimum length of a string is violated.   * &#x60;minimumLength&#x60; - The minimum length (integer). * &#x60;GENERAL_MAXIMUM_LENGTH_VIOLATED&#x60; - The maximum length of a string is violated.   * &#x60;maximumLength&#x60; - The maximum length (integer). * &#x60;GENERAL_MINIMUM_ITEMS_VIOLATED&#x60; - The minimum number of items of an array is violated.   * &#x60;minimumItems&#x60; - The minimum number of items (integer). * &#x60;GENERAL_MAXIMUM_ITEMS_VIOLATED&#x60; - The maximum number of items of an array is violated.   * &#x60;maximumItems&#x60; - The maximum number of items (integer). * &#x60;GENERAL_UNIQUE_ITEMS_VIOLATED&#x60; - The items of an array are not unique.   * &#x60;indexes&#x60; - The list of indexes that have the same value. * &#x60;GENERAL_MINIMUM_VALUE_VIOLATED&#x60; - The minimum value of a parameter is violated.   * &#x60;minimumValue&#x60; - The minimum value (integer or double). * &#x60;GENERAL_MAXIMUM_VALUE_VIOLATED&#x60; - The maximum value of a parameter is violated.   * &#x60;maximumValue&#x60; - The maximum value (integer or double). * &#x60;GENERAL_MULTIPLE_OF_VIOLATED&#x60; - The value of a parameter is not a multiple of a specific value.   * &#x60;multipleOf&#x60; - The number the value should be a multiple of (integer or double). * &#x60;GENERAL_ENUM_VIOLATED&#x60; - The value of a parameter is not one of the specified enum values.   * &#x60;enum&#x60; - The allowed enum values. * &#x60;GENERAL_INVALID_VALUE&#x60; - A parameter is set to an invalid value.   * &#x60;value&#x60; - The invalid value.   * &#x60;message&#x60; - A message explaining the invalid value. * &#x60;GENERAL_CONFLICTING_PARAMETER&#x60; - The value of a parameter conflicts with the value of one or more other parameters.   * &#x60;conflictingParameters&#x60; - The parameters whose values cause the conflict.   * &#x60;message&#x60; - A message explaining the conflict. * &#x60;GENERAL_DUPLICATE_PARAMETER&#x60; - A parameter is duplicated.  **Error codes for** &#x60;ROUTE_OPTIMIZATION_RESTRICTION_EXCEEDED&#x60; * &#x60;ROUTE_OPTIMIZATION_TOO_MANY_LOCATIONS&#x60; - The request contains too many locations.   * &#x60;locations&#x60; - The number of locations.   * &#x60;limit&#x60; - The limit. * &#x60;ROUTE_OPTIMIZATION_TOO_MANY_ORDERS&#x60; - The request contains too many orders.   * &#x60;orders&#x60; - The number of orders.   * &#x60;limit&#x60; - The limit. * &#x60;ROUTE_OPTIMIZATION_TOO_MANY_VEHICLES&#x60; - The request contains too many vehicles.   * &#x60;vehicles&#x60; - The number of vehicles.   * &#x60;limit&#x60; - The limit. * &#x60;ROUTE_OPTIMIZATION_TOO_MANY_RESOURCES&#x60; - The request contains too many resources.   * &#x60;resources&#x60; - The number of resources.   * &#x60;limit&#x60; - The limit. * &#x60;ROUTE_OPTIMIZATION_TOO_LONG_DURATION&#x60; - The duration of the request is too long.   * &#x60;duration&#x60; - The duration of the request [s].   * &#x60;limit&#x60; - The limit. * &#x60;ROUTE_OPTIMIZATION_TOO_MANY_PROFILES&#x60; - The request contains too many different combinations of profiles and traffic modes.   * &#x60;profiles&#x60; - The number of combinations of profiles and traffic modes.   * &#x60;limit&#x60; - The limit. * &#x60;ROUTE_OPTIMIZATION_ROUTES_FORBIDDEN&#x60; - The request is not allowed to contain routes. * &#x60;ROUTE_OPTIMIZATION_STOP_CONCURRENCY_FORBIDDEN&#x60; - The request is not allowed to contain locations with stop concurrency.  **Error codes for** &#x60;GENERAL_RESOURCE_NOT_FOUND&#x60; * &#x60;GENERAL_INVALID_ID&#x60; - No resource exists for the provided ID.   * &#x60;value&#x60; - The ID for which no resource exists.  **Error codes for** &#x60;GENERAL_UNSUPPORTED_MEDIA_TYPE&#x60; * &#x60;GENERAL_INVALID_CONTENT_ENCODING&#x60; - The content encoding of the request is not supported.   * &#x60;contentEncodings&#x60; - A list of supported content encodings. | 
**Parameter** | Pointer to **string** | The name of the affected query or path parameter or a JSONPath to the affected property of the request. | [optional] 
**Details** | Pointer to **map[string]interface{}** | Additional properties specific to this error class. | [optional] 

## Methods

### NewCausingError

`func NewCausingError(description string, errorCode string, ) *CausingError`

NewCausingError instantiates a new CausingError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCausingErrorWithDefaults

`func NewCausingErrorWithDefaults() *CausingError`

NewCausingErrorWithDefaults instantiates a new CausingError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CausingError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CausingError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CausingError) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetErrorCode

`func (o *CausingError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CausingError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CausingError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetParameter

`func (o *CausingError) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *CausingError) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *CausingError) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *CausingError) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetDetails

`func (o *CausingError) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CausingError) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CausingError) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CausingError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


