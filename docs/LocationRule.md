# LocationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]LocationRuleCondition**](LocationRuleCondition.md) | A list of conditions that describes when the rule applies. The rule applies if any of the conditions are met. A condition is met if all of its properties are matched. | 
**PreparationDuration** | [**DurationModifier**](DurationModifier.md) |  | 

## Methods

### NewLocationRule

`func NewLocationRule(conditions []LocationRuleCondition, preparationDuration DurationModifier, ) *LocationRule`

NewLocationRule instantiates a new LocationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationRuleWithDefaults

`func NewLocationRuleWithDefaults() *LocationRule`

NewLocationRuleWithDefaults instantiates a new LocationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *LocationRule) GetConditions() []LocationRuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *LocationRule) GetConditionsOk() (*[]LocationRuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *LocationRule) SetConditions(v []LocationRuleCondition)`

SetConditions sets Conditions field to given value.


### GetPreparationDuration

`func (o *LocationRule) GetPreparationDuration() DurationModifier`

GetPreparationDuration returns the PreparationDuration field if non-nil, zero value otherwise.

### GetPreparationDurationOk

`func (o *LocationRule) GetPreparationDurationOk() (*DurationModifier, bool)`

GetPreparationDurationOk returns a tuple with the PreparationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationDuration

`func (o *LocationRule) SetPreparationDuration(v DurationModifier)`

SetPreparationDuration sets PreparationDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


