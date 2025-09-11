# PickupDeliveryOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the order. This must be unique across all orders. | 
**Pickup** | [**TaskProperties**](TaskProperties.md) |  | 
**Delivery** | [**TaskProperties**](TaskProperties.md) |  | 
**Properties** | Pointer to [**OrderProperties**](OrderProperties.md) |  | [optional] 

## Methods

### NewPickupDeliveryOrder

`func NewPickupDeliveryOrder(id string, pickup TaskProperties, delivery TaskProperties, ) *PickupDeliveryOrder`

NewPickupDeliveryOrder instantiates a new PickupDeliveryOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickupDeliveryOrderWithDefaults

`func NewPickupDeliveryOrderWithDefaults() *PickupDeliveryOrder`

NewPickupDeliveryOrderWithDefaults instantiates a new PickupDeliveryOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PickupDeliveryOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PickupDeliveryOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PickupDeliveryOrder) SetId(v string)`

SetId sets Id field to given value.


### GetPickup

`func (o *PickupDeliveryOrder) GetPickup() TaskProperties`

GetPickup returns the Pickup field if non-nil, zero value otherwise.

### GetPickupOk

`func (o *PickupDeliveryOrder) GetPickupOk() (*TaskProperties, bool)`

GetPickupOk returns a tuple with the Pickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickup

`func (o *PickupDeliveryOrder) SetPickup(v TaskProperties)`

SetPickup sets Pickup field to given value.


### GetDelivery

`func (o *PickupDeliveryOrder) GetDelivery() TaskProperties`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *PickupDeliveryOrder) GetDeliveryOk() (*TaskProperties, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *PickupDeliveryOrder) SetDelivery(v TaskProperties)`

SetDelivery sets Delivery field to given value.


### GetProperties

`func (o *PickupDeliveryOrder) GetProperties() OrderProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PickupDeliveryOrder) GetPropertiesOk() (*OrderProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PickupDeliveryOrder) SetProperties(v OrderProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PickupDeliveryOrder) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


