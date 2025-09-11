# RespectedOrderSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderCategories** | **[]string** | The sequence of order categories that has to be respected within a route. The index of the category in the list determines the sequence. Categories that do not correspond to any order will be ignored. | 
**VehicleCategory** | Pointer to **string** | The category of vehicles to which this constraint applies. When omitted the constraint applies to all vehicles. The constraint will be ignored if no vehicle belongs to this category. | [optional] 

## Methods

### NewRespectedOrderSequence

`func NewRespectedOrderSequence(orderCategories []string, ) *RespectedOrderSequence`

NewRespectedOrderSequence instantiates a new RespectedOrderSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespectedOrderSequenceWithDefaults

`func NewRespectedOrderSequenceWithDefaults() *RespectedOrderSequence`

NewRespectedOrderSequenceWithDefaults instantiates a new RespectedOrderSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderCategories

`func (o *RespectedOrderSequence) GetOrderCategories() []string`

GetOrderCategories returns the OrderCategories field if non-nil, zero value otherwise.

### GetOrderCategoriesOk

`func (o *RespectedOrderSequence) GetOrderCategoriesOk() (*[]string, bool)`

GetOrderCategoriesOk returns a tuple with the OrderCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCategories

`func (o *RespectedOrderSequence) SetOrderCategories(v []string)`

SetOrderCategories sets OrderCategories field to given value.


### GetVehicleCategory

`func (o *RespectedOrderSequence) GetVehicleCategory() string`

GetVehicleCategory returns the VehicleCategory field if non-nil, zero value otherwise.

### GetVehicleCategoryOk

`func (o *RespectedOrderSequence) GetVehicleCategoryOk() (*string, bool)`

GetVehicleCategoryOk returns a tuple with the VehicleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategory

`func (o *RespectedOrderSequence) SetVehicleCategory(v string)`

SetVehicleCategory sets VehicleCategory field to given value.

### HasVehicleCategory

`func (o *RespectedOrderSequence) HasVehicleCategory() bool`

HasVehicleCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


