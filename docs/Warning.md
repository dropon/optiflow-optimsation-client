# Warning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable message that describes the warning. | 
**WarningCode** | **string** | A constant string that can be used to identify this warning class programmatically. A warningCode can have **details** to provide information in additional properties which are described with the code they apply to. They are of type string unless otherwise specified. Note that additional warningCodes as well as the **details** of existing warningCodes may be added at any time. Furthermore, the **description** may change at any time. * &#x60;GENERAL_PARAMETER_IGNORED&#x60; - A parameter was ignored.   * &#x60;parameter&#x60; - The ignored parameter.   * &#x60;value&#x60; - The value of the ignored parameter.   * &#x60;relatedParameter&#x60; - The parameter which caused the parameter in question to be ignored.   * &#x60;relatedValue&#x60; - The value which caused the parameter in question to be ignored. Not present if the conflict is independent of the value. * &#x60;ROUTEOPTIMIZATION_LOCATIONS_DONT_MATCH_PROFILE_REGION&#x60; - The region of most locations does not match the region of the vehicle profile specified in the request.  * &#x60;profileRegion&#x60; - The region of the profile.  * &#x60;locationsRegion&#x60; - The region of the locations. | 
**Details** | Pointer to **map[string]interface{}** | Additional properties specific to this class of warnings. | [optional] 

## Methods

### NewWarning

`func NewWarning(description string, warningCode string, ) *Warning`

NewWarning instantiates a new Warning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningWithDefaults

`func NewWarningWithDefaults() *Warning`

NewWarningWithDefaults instantiates a new Warning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Warning) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Warning) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Warning) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetWarningCode

`func (o *Warning) GetWarningCode() string`

GetWarningCode returns the WarningCode field if non-nil, zero value otherwise.

### GetWarningCodeOk

`func (o *Warning) GetWarningCodeOk() (*string, bool)`

GetWarningCodeOk returns a tuple with the WarningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningCode

`func (o *Warning) SetWarningCode(v string)`

SetWarningCode sets WarningCode field to given value.


### GetDetails

`func (o *Warning) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Warning) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Warning) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Warning) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


