# OrderCompartmentCombinationConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OrderCompartmentCombinationConstraintType**](OrderCompartmentCombinationConstraintType.md) |  | 
**OrderCategory** | **string** | The category of orders to which the constraint applies. The constraint will be ignored when no order belongs to this category. | 
**CompartmentCategory** | **string** | The category of compartments to which the constraint applies. The constraint will be ignored when no compartment belongs to this category. | 

## Methods

### NewOrderCompartmentCombinationConstraint

`func NewOrderCompartmentCombinationConstraint(type_ OrderCompartmentCombinationConstraintType, orderCategory string, compartmentCategory string, ) *OrderCompartmentCombinationConstraint`

NewOrderCompartmentCombinationConstraint instantiates a new OrderCompartmentCombinationConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCompartmentCombinationConstraintWithDefaults

`func NewOrderCompartmentCombinationConstraintWithDefaults() *OrderCompartmentCombinationConstraint`

NewOrderCompartmentCombinationConstraintWithDefaults instantiates a new OrderCompartmentCombinationConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderCompartmentCombinationConstraint) GetType() OrderCompartmentCombinationConstraintType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderCompartmentCombinationConstraint) GetTypeOk() (*OrderCompartmentCombinationConstraintType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderCompartmentCombinationConstraint) SetType(v OrderCompartmentCombinationConstraintType)`

SetType sets Type field to given value.


### GetOrderCategory

`func (o *OrderCompartmentCombinationConstraint) GetOrderCategory() string`

GetOrderCategory returns the OrderCategory field if non-nil, zero value otherwise.

### GetOrderCategoryOk

`func (o *OrderCompartmentCombinationConstraint) GetOrderCategoryOk() (*string, bool)`

GetOrderCategoryOk returns a tuple with the OrderCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCategory

`func (o *OrderCompartmentCombinationConstraint) SetOrderCategory(v string)`

SetOrderCategory sets OrderCategory field to given value.


### GetCompartmentCategory

`func (o *OrderCompartmentCombinationConstraint) GetCompartmentCategory() string`

GetCompartmentCategory returns the CompartmentCategory field if non-nil, zero value otherwise.

### GetCompartmentCategoryOk

`func (o *OrderCompartmentCombinationConstraint) GetCompartmentCategoryOk() (*string, bool)`

GetCompartmentCategoryOk returns a tuple with the CompartmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentCategory

`func (o *OrderCompartmentCombinationConstraint) SetCompartmentCategory(v string)`

SetCompartmentCategory sets CompartmentCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


