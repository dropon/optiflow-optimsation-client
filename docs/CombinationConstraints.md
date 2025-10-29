# CombinationConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderCompartment** | Pointer to [**[]OrderCompartmentCombinationConstraint**](OrderCompartmentCombinationConstraint.md) | A list of constraints on combinations of orders and compartments. | [optional] [default to {}]
**OrderVehicle** | Pointer to [**[]OrderVehicleCombinationConstraint**](OrderVehicleCombinationConstraint.md) | A list of constraints on combinations of orders and vehicles. | [optional] [default to {}]
**OrderDepot** | Pointer to [**[]OrderDepotCombinationConstraint**](OrderDepotCombinationConstraint.md) | A list of constraints on combinations of orders and depots. | [optional] [default to {}]
**DepotVehicle** | Pointer to [**[]DepotVehicleCombinationConstraint**](DepotVehicleCombinationConstraint.md) | A list of constraints on combinations of depots and vehicles. | [optional] [default to {}]

## Methods

### NewCombinationConstraints

`func NewCombinationConstraints() *CombinationConstraints`

NewCombinationConstraints instantiates a new CombinationConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinationConstraintsWithDefaults

`func NewCombinationConstraintsWithDefaults() *CombinationConstraints`

NewCombinationConstraintsWithDefaults instantiates a new CombinationConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderCompartment

`func (o *CombinationConstraints) GetOrderCompartment() []OrderCompartmentCombinationConstraint`

GetOrderCompartment returns the OrderCompartment field if non-nil, zero value otherwise.

### GetOrderCompartmentOk

`func (o *CombinationConstraints) GetOrderCompartmentOk() (*[]OrderCompartmentCombinationConstraint, bool)`

GetOrderCompartmentOk returns a tuple with the OrderCompartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCompartment

`func (o *CombinationConstraints) SetOrderCompartment(v []OrderCompartmentCombinationConstraint)`

SetOrderCompartment sets OrderCompartment field to given value.

### HasOrderCompartment

`func (o *CombinationConstraints) HasOrderCompartment() bool`

HasOrderCompartment returns a boolean if a field has been set.

### GetOrderVehicle

`func (o *CombinationConstraints) GetOrderVehicle() []OrderVehicleCombinationConstraint`

GetOrderVehicle returns the OrderVehicle field if non-nil, zero value otherwise.

### GetOrderVehicleOk

`func (o *CombinationConstraints) GetOrderVehicleOk() (*[]OrderVehicleCombinationConstraint, bool)`

GetOrderVehicleOk returns a tuple with the OrderVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderVehicle

`func (o *CombinationConstraints) SetOrderVehicle(v []OrderVehicleCombinationConstraint)`

SetOrderVehicle sets OrderVehicle field to given value.

### HasOrderVehicle

`func (o *CombinationConstraints) HasOrderVehicle() bool`

HasOrderVehicle returns a boolean if a field has been set.

### GetOrderDepot

`func (o *CombinationConstraints) GetOrderDepot() []OrderDepotCombinationConstraint`

GetOrderDepot returns the OrderDepot field if non-nil, zero value otherwise.

### GetOrderDepotOk

`func (o *CombinationConstraints) GetOrderDepotOk() (*[]OrderDepotCombinationConstraint, bool)`

GetOrderDepotOk returns a tuple with the OrderDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDepot

`func (o *CombinationConstraints) SetOrderDepot(v []OrderDepotCombinationConstraint)`

SetOrderDepot sets OrderDepot field to given value.

### HasOrderDepot

`func (o *CombinationConstraints) HasOrderDepot() bool`

HasOrderDepot returns a boolean if a field has been set.

### GetDepotVehicle

`func (o *CombinationConstraints) GetDepotVehicle() []DepotVehicleCombinationConstraint`

GetDepotVehicle returns the DepotVehicle field if non-nil, zero value otherwise.

### GetDepotVehicleOk

`func (o *CombinationConstraints) GetDepotVehicleOk() (*[]DepotVehicleCombinationConstraint, bool)`

GetDepotVehicleOk returns a tuple with the DepotVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotVehicle

`func (o *CombinationConstraints) SetDepotVehicle(v []DepotVehicleCombinationConstraint)`

SetDepotVehicle sets DepotVehicle field to given value.

### HasDepotVehicle

`func (o *CombinationConstraints) HasDepotVehicle() bool`

HasDepotVehicle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


