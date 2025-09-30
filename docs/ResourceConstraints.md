# ResourceConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumVehicles** | Pointer to **int32** | Restricts the number of vehicles the resource can be assigned to. | [optional] 
**MaximumTotalRouteDuration** | Pointer to **int32** | Restricts the summed duration [s] of routes corresponding to vehicles assigned to the resource. | [optional] 

## Methods

### NewResourceConstraints

`func NewResourceConstraints() *ResourceConstraints`

NewResourceConstraints instantiates a new ResourceConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConstraintsWithDefaults

`func NewResourceConstraintsWithDefaults() *ResourceConstraints`

NewResourceConstraintsWithDefaults instantiates a new ResourceConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumVehicles

`func (o *ResourceConstraints) GetMaximumVehicles() int32`

GetMaximumVehicles returns the MaximumVehicles field if non-nil, zero value otherwise.

### GetMaximumVehiclesOk

`func (o *ResourceConstraints) GetMaximumVehiclesOk() (*int32, bool)`

GetMaximumVehiclesOk returns a tuple with the MaximumVehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVehicles

`func (o *ResourceConstraints) SetMaximumVehicles(v int32)`

SetMaximumVehicles sets MaximumVehicles field to given value.

### HasMaximumVehicles

`func (o *ResourceConstraints) HasMaximumVehicles() bool`

HasMaximumVehicles returns a boolean if a field has been set.

### GetMaximumTotalRouteDuration

`func (o *ResourceConstraints) GetMaximumTotalRouteDuration() int32`

GetMaximumTotalRouteDuration returns the MaximumTotalRouteDuration field if non-nil, zero value otherwise.

### GetMaximumTotalRouteDurationOk

`func (o *ResourceConstraints) GetMaximumTotalRouteDurationOk() (*int32, bool)`

GetMaximumTotalRouteDurationOk returns a tuple with the MaximumTotalRouteDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumTotalRouteDuration

`func (o *ResourceConstraints) SetMaximumTotalRouteDuration(v int32)`

SetMaximumTotalRouteDuration sets MaximumTotalRouteDuration field to given value.

### HasMaximumTotalRouteDuration

`func (o *ResourceConstraints) HasMaximumTotalRouteDuration() bool`

HasMaximumTotalRouteDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


