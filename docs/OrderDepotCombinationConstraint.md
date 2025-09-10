# OrderDepotCombinationConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OrderDepotCombinationConstraintType**](OrderDepotCombinationConstraintType.md) |  | 
**OrderCategory** | **string** | The category of orders to which the constraint applies. The constraint will be ignored if no order belongs to this category. | 
**DepotCategory** | **string** | The category of depots to which the constraint applies. The constraint will be ignored if no depot belongs to this category. | 

## Methods

### NewOrderDepotCombinationConstraint

`func NewOrderDepotCombinationConstraint(type_ OrderDepotCombinationConstraintType, orderCategory string, depotCategory string, ) *OrderDepotCombinationConstraint`

NewOrderDepotCombinationConstraint instantiates a new OrderDepotCombinationConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDepotCombinationConstraintWithDefaults

`func NewOrderDepotCombinationConstraintWithDefaults() *OrderDepotCombinationConstraint`

NewOrderDepotCombinationConstraintWithDefaults instantiates a new OrderDepotCombinationConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderDepotCombinationConstraint) GetType() OrderDepotCombinationConstraintType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderDepotCombinationConstraint) GetTypeOk() (*OrderDepotCombinationConstraintType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderDepotCombinationConstraint) SetType(v OrderDepotCombinationConstraintType)`

SetType sets Type field to given value.


### GetOrderCategory

`func (o *OrderDepotCombinationConstraint) GetOrderCategory() string`

GetOrderCategory returns the OrderCategory field if non-nil, zero value otherwise.

### GetOrderCategoryOk

`func (o *OrderDepotCombinationConstraint) GetOrderCategoryOk() (*string, bool)`

GetOrderCategoryOk returns a tuple with the OrderCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCategory

`func (o *OrderDepotCombinationConstraint) SetOrderCategory(v string)`

SetOrderCategory sets OrderCategory field to given value.


### GetDepotCategory

`func (o *OrderDepotCombinationConstraint) GetDepotCategory() string`

GetDepotCategory returns the DepotCategory field if non-nil, zero value otherwise.

### GetDepotCategoryOk

`func (o *OrderDepotCombinationConstraint) GetDepotCategoryOk() (*string, bool)`

GetDepotCategoryOk returns a tuple with the DepotCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotCategory

`func (o *OrderDepotCombinationConstraint) SetDepotCategory(v string)`

SetDepotCategory sets DepotCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


