# OrderVehicleCombinationConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OrderVehicleCombinationConstraintType**](OrderVehicleCombinationConstraintType.md) |  | 
**OrderCategory** | **string** | The category of orders to which the constraint applies. The constraint will be ignored if no order belongs to this category. | 
**VehicleCategory** | **string** | The category of vehicles to which the constraint applies. The constraint will be ignored if no vehicle belongs to this category. | 
**ViolationCost** | Pointer to **float64** | The cost incurred when an order-vehicle combination does not meet this constraint. When omitted, all order-vehicle combinations must satisfy this constraint. | [optional] 

## Methods

### NewOrderVehicleCombinationConstraint

`func NewOrderVehicleCombinationConstraint(type_ OrderVehicleCombinationConstraintType, orderCategory string, vehicleCategory string, ) *OrderVehicleCombinationConstraint`

NewOrderVehicleCombinationConstraint instantiates a new OrderVehicleCombinationConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderVehicleCombinationConstraintWithDefaults

`func NewOrderVehicleCombinationConstraintWithDefaults() *OrderVehicleCombinationConstraint`

NewOrderVehicleCombinationConstraintWithDefaults instantiates a new OrderVehicleCombinationConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderVehicleCombinationConstraint) GetType() OrderVehicleCombinationConstraintType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderVehicleCombinationConstraint) GetTypeOk() (*OrderVehicleCombinationConstraintType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderVehicleCombinationConstraint) SetType(v OrderVehicleCombinationConstraintType)`

SetType sets Type field to given value.


### GetOrderCategory

`func (o *OrderVehicleCombinationConstraint) GetOrderCategory() string`

GetOrderCategory returns the OrderCategory field if non-nil, zero value otherwise.

### GetOrderCategoryOk

`func (o *OrderVehicleCombinationConstraint) GetOrderCategoryOk() (*string, bool)`

GetOrderCategoryOk returns a tuple with the OrderCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCategory

`func (o *OrderVehicleCombinationConstraint) SetOrderCategory(v string)`

SetOrderCategory sets OrderCategory field to given value.


### GetVehicleCategory

`func (o *OrderVehicleCombinationConstraint) GetVehicleCategory() string`

GetVehicleCategory returns the VehicleCategory field if non-nil, zero value otherwise.

### GetVehicleCategoryOk

`func (o *OrderVehicleCombinationConstraint) GetVehicleCategoryOk() (*string, bool)`

GetVehicleCategoryOk returns a tuple with the VehicleCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCategory

`func (o *OrderVehicleCombinationConstraint) SetVehicleCategory(v string)`

SetVehicleCategory sets VehicleCategory field to given value.


### GetViolationCost

`func (o *OrderVehicleCombinationConstraint) GetViolationCost() float64`

GetViolationCost returns the ViolationCost field if non-nil, zero value otherwise.

### GetViolationCostOk

`func (o *OrderVehicleCombinationConstraint) GetViolationCostOk() (*float64, bool)`

GetViolationCostOk returns a tuple with the ViolationCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCost

`func (o *OrderVehicleCombinationConstraint) SetViolationCost(v float64)`

SetViolationCost sets ViolationCost field to given value.

### HasViolationCost

`func (o *OrderVehicleCombinationConstraint) HasViolationCost() bool`

HasViolationCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


