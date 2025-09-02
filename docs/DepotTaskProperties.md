# DepotTaskProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | The duration [s] it takes to execute this task. | [optional] [default to 0]

## Methods

### NewDepotTaskProperties

`func NewDepotTaskProperties() *DepotTaskProperties`

NewDepotTaskProperties instantiates a new DepotTaskProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotTaskPropertiesWithDefaults

`func NewDepotTaskPropertiesWithDefaults() *DepotTaskProperties`

NewDepotTaskPropertiesWithDefaults instantiates a new DepotTaskProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DepotTaskProperties) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DepotTaskProperties) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DepotTaskProperties) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DepotTaskProperties) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


