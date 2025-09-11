# RespectedTaskSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskCategories** | **[]string** | The sequence of task categories that needs to be respected within a route. The index of the category in the list determines the sequence. Categories that do not correspond to any task are ignored. | 
**VehicleCategory** | Pointer to **string** | The category of vehicles to which this constraint applies. When omitted the constraint applies to all vehicles. The constraint will be ignored when no vehicle belongs to this category. | [optional] 

## Methods

### NewRespectedTaskSequence

`func NewRespectedTaskSequence(taskCategories []string, ) *RespectedTaskSequence`

NewRespectedTaskSequence instantiates a new RespectedTaskSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespectedTaskSequenceWithDefaults

`func NewRespectedTaskSequenceWithDefaults() *RespectedTaskSequence`

NewRespectedTaskSequenceWithDefaults instantiates a new RespectedTaskSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskCategories

`func (o *RespectedTaskSequence) GetTaskCategories() []string`

GetTaskCategories returns the TaskCategories field if non-nil, zero value otherwise.

### GetTaskCategoriesOk

`func (o *RespectedTaskSequence) GetTaskCategoriesOk() (*[]string, bool)`

GetTaskCategoriesOk returns a tuple with the TaskCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCategories

`func (o *RespectedTaskSequence) SetTaskCategories(v []string)`

SetTaskCategories sets TaskCategories field to given value.


### GetVehicleCategory

`func (o *RespectedTaskSequence) GetVehicleCategory() string`

GetVehicleCategory returns the VehicleCategory field if non-nil, zero value otherwise.

### GetVehicleCategoryOk

`func (o *RespectedTaskSequence) GetVehicleCategoryOk() (*string, bool)`

GetVehicleCategoryOk returns a tuple with the VehicleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategory

`func (o *RespectedTaskSequence) SetVehicleCategory(v string)`

SetVehicleCategory sets VehicleCategory field to given value.

### HasVehicleCategory

`func (o *RespectedTaskSequence) HasVehicleCategory() bool`

HasVehicleCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


