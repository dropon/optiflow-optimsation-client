# Orders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pickups** | Pointer to [**[]PickupOrder**](PickupOrder.md) | A list of orders that must be picked up from a specific location and transported to a depot. Please note that the upper bound on number of pickups is a technical limit. Check your individual price plan or contract to see which limits apply. | [optional] [default to []]
**Deliveries** | Pointer to [**[]DeliveryOrder**](DeliveryOrder.md) | A list of orders that must be delivered to a specific location and transported from a depot. Please note that the upper bound on number of deliveries is a technical limit. Check your individual price plan or contract to see which limits apply. | [optional] [default to []]
**PickupDeliveries** | Pointer to [**[]PickupDeliveryOrder**](PickupDeliveryOrder.md) | A list of orders that must be picked up at a specific location and delivered to a specific location. Please note that the upper bound on number of pickup-deliveries is a technical limit. Check your individual price plan or contract to see which limits apply. | [optional] [default to []]

## Methods

### NewOrders

`func NewOrders() *Orders`

NewOrders instantiates a new Orders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersWithDefaults

`func NewOrdersWithDefaults() *Orders`

NewOrdersWithDefaults instantiates a new Orders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickups

`func (o *Orders) GetPickups() []PickupOrder`

GetPickups returns the Pickups field if non-nil, zero value otherwise.

### GetPickupsOk

`func (o *Orders) GetPickupsOk() (*[]PickupOrder, bool)`

GetPickupsOk returns a tuple with the Pickups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickups

`func (o *Orders) SetPickups(v []PickupOrder)`

SetPickups sets Pickups field to given value.

### HasPickups

`func (o *Orders) HasPickups() bool`

HasPickups returns a boolean if a field has been set.

### GetDeliveries

`func (o *Orders) GetDeliveries() []DeliveryOrder`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *Orders) GetDeliveriesOk() (*[]DeliveryOrder, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *Orders) SetDeliveries(v []DeliveryOrder)`

SetDeliveries sets Deliveries field to given value.

### HasDeliveries

`func (o *Orders) HasDeliveries() bool`

HasDeliveries returns a boolean if a field has been set.

### GetPickupDeliveries

`func (o *Orders) GetPickupDeliveries() []PickupDeliveryOrder`

GetPickupDeliveries returns the PickupDeliveries field if non-nil, zero value otherwise.

### GetPickupDeliveriesOk

`func (o *Orders) GetPickupDeliveriesOk() (*[]PickupDeliveryOrder, bool)`

GetPickupDeliveriesOk returns a tuple with the PickupDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDeliveries

`func (o *Orders) SetPickupDeliveries(v []PickupDeliveryOrder)`

SetPickupDeliveries sets PickupDeliveries field to given value.

### HasPickupDeliveries

`func (o *Orders) HasPickupDeliveries() bool`

HasPickupDeliveries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


