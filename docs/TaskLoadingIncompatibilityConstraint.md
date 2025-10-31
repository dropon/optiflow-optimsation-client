# TaskLoadingIncompatibilityConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadedOrderCategory** | **string** | Orders belonging to this category must be loaded in the vehicle or in a compartment of the vehicle for this constraint to apply. The constraint will be ignored when no order belongs to this category. | 
**ForbiddenTaskCategory** | **string** | Tasks belonging to this category cannot be loaded or unloaded while the constraint applies. The constraint will be ignored when no task belongs to this category. | 
**VehicleCategory** | Pointer to **string** | The category of vehicles to which this constraint applies. If a vehicle has compartments, the constraint applies to each compartment individually within vehicles of this category. When omitted the constraint applies to all vehicles. The constraint will be ignored when no vehicle belongs to this category. | [optional] 

## Methods

### NewTaskLoadingIncompatibilityConstraint

`func NewTaskLoadingIncompatibilityConstraint(loadedOrderCategory string, forbiddenTaskCategory string, ) *TaskLoadingIncompatibilityConstraint`

NewTaskLoadingIncompatibilityConstraint instantiates a new TaskLoadingIncompatibilityConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskLoadingIncompatibilityConstraintWithDefaults

`func NewTaskLoadingIncompatibilityConstraintWithDefaults() *TaskLoadingIncompatibilityConstraint`

NewTaskLoadingIncompatibilityConstraintWithDefaults instantiates a new TaskLoadingIncompatibilityConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadedOrderCategory

`func (o *TaskLoadingIncompatibilityConstraint) GetLoadedOrderCategory() string`

GetLoadedOrderCategory returns the LoadedOrderCategory field if non-nil, zero value otherwise.

### GetLoadedOrderCategoryOk

`func (o *TaskLoadingIncompatibilityConstraint) GetLoadedOrderCategoryOk() (*string, bool)`

GetLoadedOrderCategoryOk returns a tuple with the LoadedOrderCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedOrderCategory

`func (o *TaskLoadingIncompatibilityConstraint) SetLoadedOrderCategory(v string)`

SetLoadedOrderCategory sets LoadedOrderCategory field to given value.


### GetForbiddenTaskCategory

`func (o *TaskLoadingIncompatibilityConstraint) GetForbiddenTaskCategory() string`

GetForbiddenTaskCategory returns the ForbiddenTaskCategory field if non-nil, zero value otherwise.

### GetForbiddenTaskCategoryOk

`func (o *TaskLoadingIncompatibilityConstraint) GetForbiddenTaskCategoryOk() (*string, bool)`

GetForbiddenTaskCategoryOk returns a tuple with the ForbiddenTaskCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenTaskCategory

`func (o *TaskLoadingIncompatibilityConstraint) SetForbiddenTaskCategory(v string)`

SetForbiddenTaskCategory sets ForbiddenTaskCategory field to given value.


### GetVehicleCategory

`func (o *TaskLoadingIncompatibilityConstraint) GetVehicleCategory() string`

GetVehicleCategory returns the VehicleCategory field if non-nil, zero value otherwise.

### GetVehicleCategoryOk

`func (o *TaskLoadingIncompatibilityConstraint) GetVehicleCategoryOk() (*string, bool)`

GetVehicleCategoryOk returns a tuple with the VehicleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategory

`func (o *TaskLoadingIncompatibilityConstraint) SetVehicleCategory(v string)`

SetVehicleCategory sets VehicleCategory field to given value.

### HasVehicleCategory

`func (o *TaskLoadingIncompatibilityConstraint) HasVehicleCategory() bool`

HasVehicleCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


