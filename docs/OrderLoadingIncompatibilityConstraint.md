# OrderLoadingIncompatibilityConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadedOrderCategory** | **string** | Orders belonging to this category must be loaded in the vehicle for this constraint to apply. The constraint will be ignored when no order belongs to this category. | 
**ForbiddenOrderCategory** | **string** | Orders belonging to this category cannot be loaded or unloaded while the constraint applies. The constraint will be ignored when no order belongs to this category. | 
**VehicleCategory** | Pointer to **string** | The category of vehicles to which this constraint applies. When omitted the constraint applies to all vehicles. The constraint will be ignored when no vehicle belongs to this category. | [optional] 

## Methods

### NewOrderLoadingIncompatibilityConstraint

`func NewOrderLoadingIncompatibilityConstraint(loadedOrderCategory string, forbiddenOrderCategory string, ) *OrderLoadingIncompatibilityConstraint`

NewOrderLoadingIncompatibilityConstraint instantiates a new OrderLoadingIncompatibilityConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLoadingIncompatibilityConstraintWithDefaults

`func NewOrderLoadingIncompatibilityConstraintWithDefaults() *OrderLoadingIncompatibilityConstraint`

NewOrderLoadingIncompatibilityConstraintWithDefaults instantiates a new OrderLoadingIncompatibilityConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadedOrderCategory

`func (o *OrderLoadingIncompatibilityConstraint) GetLoadedOrderCategory() string`

GetLoadedOrderCategory returns the LoadedOrderCategory field if non-nil, zero value otherwise.

### GetLoadedOrderCategoryOk

`func (o *OrderLoadingIncompatibilityConstraint) GetLoadedOrderCategoryOk() (*string, bool)`

GetLoadedOrderCategoryOk returns a tuple with the LoadedOrderCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedOrderCategory

`func (o *OrderLoadingIncompatibilityConstraint) SetLoadedOrderCategory(v string)`

SetLoadedOrderCategory sets LoadedOrderCategory field to given value.


### GetForbiddenOrderCategory

`func (o *OrderLoadingIncompatibilityConstraint) GetForbiddenOrderCategory() string`

GetForbiddenOrderCategory returns the ForbiddenOrderCategory field if non-nil, zero value otherwise.

### GetForbiddenOrderCategoryOk

`func (o *OrderLoadingIncompatibilityConstraint) GetForbiddenOrderCategoryOk() (*string, bool)`

GetForbiddenOrderCategoryOk returns a tuple with the ForbiddenOrderCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenOrderCategory

`func (o *OrderLoadingIncompatibilityConstraint) SetForbiddenOrderCategory(v string)`

SetForbiddenOrderCategory sets ForbiddenOrderCategory field to given value.


### GetVehicleCategory

`func (o *OrderLoadingIncompatibilityConstraint) GetVehicleCategory() string`

GetVehicleCategory returns the VehicleCategory field if non-nil, zero value otherwise.

### GetVehicleCategoryOk

`func (o *OrderLoadingIncompatibilityConstraint) GetVehicleCategoryOk() (*string, bool)`

GetVehicleCategoryOk returns a tuple with the VehicleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategory

`func (o *OrderLoadingIncompatibilityConstraint) SetVehicleCategory(v string)`

SetVehicleCategory sets VehicleCategory field to given value.

### HasVehicleCategory

`func (o *OrderLoadingIncompatibilityConstraint) HasVehicleCategory() bool`

HasVehicleCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


