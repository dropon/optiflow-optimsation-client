# LocationRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationCategory** | Pointer to **string** | The rule applies only if the location belongs to this category. When omitted, it applies to all locations. | [optional] 
**PreviousLocationCategory** | Pointer to **string** | The rule applies only if the previous location in the route belongs to this category. When omitted, it applies independently of the previous location. | [optional] 
**VehicleCategory** | Pointer to **string** | The rule applies only if the vehicle visiting the location belongs to this category. When omitted, it applies independently of the vehicle visiting the location. | [optional] 

## Methods

### NewLocationRuleCondition

`func NewLocationRuleCondition() *LocationRuleCondition`

NewLocationRuleCondition instantiates a new LocationRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationRuleConditionWithDefaults

`func NewLocationRuleConditionWithDefaults() *LocationRuleCondition`

NewLocationRuleConditionWithDefaults instantiates a new LocationRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationCategory

`func (o *LocationRuleCondition) GetLocationCategory() string`

GetLocationCategory returns the LocationCategory field if non-nil, zero value otherwise.

### GetLocationCategoryOk

`func (o *LocationRuleCondition) GetLocationCategoryOk() (*string, bool)`

GetLocationCategoryOk returns a tuple with the LocationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCategory

`func (o *LocationRuleCondition) SetLocationCategory(v string)`

SetLocationCategory sets LocationCategory field to given value.

### HasLocationCategory

`func (o *LocationRuleCondition) HasLocationCategory() bool`

HasLocationCategory returns a boolean if a field has been set.

### GetPreviousLocationCategory

`func (o *LocationRuleCondition) GetPreviousLocationCategory() string`

GetPreviousLocationCategory returns the PreviousLocationCategory field if non-nil, zero value otherwise.

### GetPreviousLocationCategoryOk

`func (o *LocationRuleCondition) GetPreviousLocationCategoryOk() (*string, bool)`

GetPreviousLocationCategoryOk returns a tuple with the PreviousLocationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLocationCategory

`func (o *LocationRuleCondition) SetPreviousLocationCategory(v string)`

SetPreviousLocationCategory sets PreviousLocationCategory field to given value.

### HasPreviousLocationCategory

`func (o *LocationRuleCondition) HasPreviousLocationCategory() bool`

HasPreviousLocationCategory returns a boolean if a field has been set.

### GetVehicleCategory

`func (o *LocationRuleCondition) GetVehicleCategory() string`

GetVehicleCategory returns the VehicleCategory field if non-nil, zero value otherwise.

### GetVehicleCategoryOk

`func (o *LocationRuleCondition) GetVehicleCategoryOk() (*string, bool)`

GetVehicleCategoryOk returns a tuple with the VehicleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategory

`func (o *LocationRuleCondition) SetVehicleCategory(v string)`

SetVehicleCategory sets VehicleCategory field to given value.

### HasVehicleCategory

`func (o *LocationRuleCondition) HasVehicleCategory() bool`

HasVehicleCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


