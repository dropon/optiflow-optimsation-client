# MixedLoadingProhibition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConflictingLoadCategory1** | **string** | A transport with this load category is not allowed to be on the same trip as a transport with load category conflictingLoadCategory2. The load category can be any string but it must not be empty and not the same as conflictingLoadCategory2. | 
**ConflictingLoadCategory2** | **string** | A transport with this load category is not allowed to be on the same trip as a transport with load category conflictingLoadCategory1. The load category can be any string but it must not be empty and not the same as conflictingLoadCategory1. | 

## Methods

### NewMixedLoadingProhibition

`func NewMixedLoadingProhibition(conflictingLoadCategory1 string, conflictingLoadCategory2 string, ) *MixedLoadingProhibition`

NewMixedLoadingProhibition instantiates a new MixedLoadingProhibition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixedLoadingProhibitionWithDefaults

`func NewMixedLoadingProhibitionWithDefaults() *MixedLoadingProhibition`

NewMixedLoadingProhibitionWithDefaults instantiates a new MixedLoadingProhibition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConflictingLoadCategory1

`func (o *MixedLoadingProhibition) GetConflictingLoadCategory1() string`

GetConflictingLoadCategory1 returns the ConflictingLoadCategory1 field if non-nil, zero value otherwise.

### GetConflictingLoadCategory1Ok

`func (o *MixedLoadingProhibition) GetConflictingLoadCategory1Ok() (*string, bool)`

GetConflictingLoadCategory1Ok returns a tuple with the ConflictingLoadCategory1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingLoadCategory1

`func (o *MixedLoadingProhibition) SetConflictingLoadCategory1(v string)`

SetConflictingLoadCategory1 sets ConflictingLoadCategory1 field to given value.


### GetConflictingLoadCategory2

`func (o *MixedLoadingProhibition) GetConflictingLoadCategory2() string`

GetConflictingLoadCategory2 returns the ConflictingLoadCategory2 field if non-nil, zero value otherwise.

### GetConflictingLoadCategory2Ok

`func (o *MixedLoadingProhibition) GetConflictingLoadCategory2Ok() (*string, bool)`

GetConflictingLoadCategory2Ok returns a tuple with the ConflictingLoadCategory2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingLoadCategory2

`func (o *MixedLoadingProhibition) SetConflictingLoadCategory2(v string)`

SetConflictingLoadCategory2 sets ConflictingLoadCategory2 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


