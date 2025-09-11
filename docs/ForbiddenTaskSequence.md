# ForbiddenTaskSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstTaskCategory** | **string** | The category of tasks that cannot be scheduled on the route before a task with the second task category. The constraint will be ignored when no task belongs to this category. | 
**Type** | [**ForbiddenTaskSequenceType**](ForbiddenTaskSequenceType.md) |  | 
**SecondTaskCategory** | **string** | The category of the tasks that must not be preceded in the route by tasks belonging to the first task category. The constraint will be ignored when no task belongs to this category. | 
**VehicleCategory** | Pointer to **string** | The category of vehicles to which this constraint applies. When omitted the constraint applies to all vehicles. The constraint will be ignored when no vehicle belongs to this category. | [optional] 

## Methods

### NewForbiddenTaskSequence

`func NewForbiddenTaskSequence(firstTaskCategory string, type_ ForbiddenTaskSequenceType, secondTaskCategory string, ) *ForbiddenTaskSequence`

NewForbiddenTaskSequence instantiates a new ForbiddenTaskSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenTaskSequenceWithDefaults

`func NewForbiddenTaskSequenceWithDefaults() *ForbiddenTaskSequence`

NewForbiddenTaskSequenceWithDefaults instantiates a new ForbiddenTaskSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstTaskCategory

`func (o *ForbiddenTaskSequence) GetFirstTaskCategory() string`

GetFirstTaskCategory returns the FirstTaskCategory field if non-nil, zero value otherwise.

### GetFirstTaskCategoryOk

`func (o *ForbiddenTaskSequence) GetFirstTaskCategoryOk() (*string, bool)`

GetFirstTaskCategoryOk returns a tuple with the FirstTaskCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTaskCategory

`func (o *ForbiddenTaskSequence) SetFirstTaskCategory(v string)`

SetFirstTaskCategory sets FirstTaskCategory field to given value.


### GetType

`func (o *ForbiddenTaskSequence) GetType() ForbiddenTaskSequenceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ForbiddenTaskSequence) GetTypeOk() (*ForbiddenTaskSequenceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ForbiddenTaskSequence) SetType(v ForbiddenTaskSequenceType)`

SetType sets Type field to given value.


### GetSecondTaskCategory

`func (o *ForbiddenTaskSequence) GetSecondTaskCategory() string`

GetSecondTaskCategory returns the SecondTaskCategory field if non-nil, zero value otherwise.

### GetSecondTaskCategoryOk

`func (o *ForbiddenTaskSequence) GetSecondTaskCategoryOk() (*string, bool)`

GetSecondTaskCategoryOk returns a tuple with the SecondTaskCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondTaskCategory

`func (o *ForbiddenTaskSequence) SetSecondTaskCategory(v string)`

SetSecondTaskCategory sets SecondTaskCategory field to given value.


### GetVehicleCategory

`func (o *ForbiddenTaskSequence) GetVehicleCategory() string`

GetVehicleCategory returns the VehicleCategory field if non-nil, zero value otherwise.

### GetVehicleCategoryOk

`func (o *ForbiddenTaskSequence) GetVehicleCategoryOk() (*string, bool)`

GetVehicleCategoryOk returns a tuple with the VehicleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategory

`func (o *ForbiddenTaskSequence) SetVehicleCategory(v string)`

SetVehicleCategory sets VehicleCategory field to given value.

### HasVehicleCategory

`func (o *ForbiddenTaskSequence) HasVehicleCategory() bool`

HasVehicleCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


