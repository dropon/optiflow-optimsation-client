# Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**[]LocationRule**](LocationRule.md) | A list of rules that conditionally modify location properties. | [optional] [default to {}]

## Methods

### NewRules

`func NewRules() *Rules`

NewRules instantiates a new Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesWithDefaults

`func NewRulesWithDefaults() *Rules`

NewRulesWithDefaults instantiates a new Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *Rules) GetLocations() []LocationRule`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Rules) GetLocationsOk() (*[]LocationRule, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Rules) SetLocations(v []LocationRule)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *Rules) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


