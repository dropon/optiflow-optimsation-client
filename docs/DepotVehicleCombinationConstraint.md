# DepotVehicleCombinationConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DepotVehicleCombinationConstraintType**](DepotVehicleCombinationConstraintType.md) |  | 
**DepotCategory** | **string** | The category of depots to which the constraint applies. The constraint will be ignored when no depot belongs to this category. | 
**VehicleCategory** | **string** | The category of vehicles to which the constraint applies. The constraint will be ignored when no vehicle belongs to this category. | 

## Methods

### NewDepotVehicleCombinationConstraint

`func NewDepotVehicleCombinationConstraint(type_ DepotVehicleCombinationConstraintType, depotCategory string, vehicleCategory string, ) *DepotVehicleCombinationConstraint`

NewDepotVehicleCombinationConstraint instantiates a new DepotVehicleCombinationConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotVehicleCombinationConstraintWithDefaults

`func NewDepotVehicleCombinationConstraintWithDefaults() *DepotVehicleCombinationConstraint`

NewDepotVehicleCombinationConstraintWithDefaults instantiates a new DepotVehicleCombinationConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DepotVehicleCombinationConstraint) GetType() DepotVehicleCombinationConstraintType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DepotVehicleCombinationConstraint) GetTypeOk() (*DepotVehicleCombinationConstraintType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DepotVehicleCombinationConstraint) SetType(v DepotVehicleCombinationConstraintType)`

SetType sets Type field to given value.


### GetDepotCategory

`func (o *DepotVehicleCombinationConstraint) GetDepotCategory() string`

GetDepotCategory returns the DepotCategory field if non-nil, zero value otherwise.

### GetDepotCategoryOk

`func (o *DepotVehicleCombinationConstraint) GetDepotCategoryOk() (*string, bool)`

GetDepotCategoryOk returns a tuple with the DepotCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotCategory

`func (o *DepotVehicleCombinationConstraint) SetDepotCategory(v string)`

SetDepotCategory sets DepotCategory field to given value.


### GetVehicleCategory

`func (o *DepotVehicleCombinationConstraint) GetVehicleCategory() string`

GetVehicleCategory returns the VehicleCategory field if non-nil, zero value otherwise.

### GetVehicleCategoryOk

`func (o *DepotVehicleCombinationConstraint) GetVehicleCategoryOk() (*string, bool)`

GetVehicleCategoryOk returns a tuple with the VehicleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategory

`func (o *DepotVehicleCombinationConstraint) SetVehicleCategory(v string)`

SetVehicleCategory sets VehicleCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


