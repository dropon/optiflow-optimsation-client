# OrderProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loads** | Pointer to [**[]Load**](Load.md) | A list of load definitions that describe the necessary vehicle capacity to transport the order. For each dimension, the sum of the values of orders loaded into the vehicle must be lower than or equal to the value of the vehicle. | [optional] [default to []]
**OutsourcingCost** | Pointer to **float64** | Defines the cost for not scheduling the order on a route. This cost is weighed against the cost of scheduling the order on a route. When omitted the optimization will try to schedule the order regardless of the added cost. | [optional] 
**Categories** | Pointer to **[]string** | A list of categories the order belongs to that can be used to describe constraints or rules. | [optional] [default to []]

## Methods

### NewOrderProperties

`func NewOrderProperties() *OrderProperties`

NewOrderProperties instantiates a new OrderProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPropertiesWithDefaults

`func NewOrderPropertiesWithDefaults() *OrderProperties`

NewOrderPropertiesWithDefaults instantiates a new OrderProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoads

`func (o *OrderProperties) GetLoads() []Load`

GetLoads returns the Loads field if non-nil, zero value otherwise.

### GetLoadsOk

`func (o *OrderProperties) GetLoadsOk() (*[]Load, bool)`

GetLoadsOk returns a tuple with the Loads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoads

`func (o *OrderProperties) SetLoads(v []Load)`

SetLoads sets Loads field to given value.

### HasLoads

`func (o *OrderProperties) HasLoads() bool`

HasLoads returns a boolean if a field has been set.

### GetOutsourcingCost

`func (o *OrderProperties) GetOutsourcingCost() float64`

GetOutsourcingCost returns the OutsourcingCost field if non-nil, zero value otherwise.

### GetOutsourcingCostOk

`func (o *OrderProperties) GetOutsourcingCostOk() (*float64, bool)`

GetOutsourcingCostOk returns a tuple with the OutsourcingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsourcingCost

`func (o *OrderProperties) SetOutsourcingCost(v float64)`

SetOutsourcingCost sets OutsourcingCost field to given value.

### HasOutsourcingCost

`func (o *OrderProperties) HasOutsourcingCost() bool`

HasOutsourcingCost returns a boolean if a field has been set.

### GetCategories

`func (o *OrderProperties) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *OrderProperties) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *OrderProperties) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *OrderProperties) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


