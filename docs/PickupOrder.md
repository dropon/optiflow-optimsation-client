# PickupOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the order. This must be unique across all orders. | 
**Pickup** | [**TaskProperties**](TaskProperties.md) |  | 
**Properties** | Pointer to [**OrderProperties**](OrderProperties.md) |  | [optional] 

## Methods

### NewPickupOrder

`func NewPickupOrder(id string, pickup TaskProperties, ) *PickupOrder`

NewPickupOrder instantiates a new PickupOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickupOrderWithDefaults

`func NewPickupOrderWithDefaults() *PickupOrder`

NewPickupOrderWithDefaults instantiates a new PickupOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PickupOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PickupOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PickupOrder) SetId(v string)`

SetId sets Id field to given value.


### GetPickup

`func (o *PickupOrder) GetPickup() TaskProperties`

GetPickup returns the Pickup field if non-nil, zero value otherwise.

### GetPickupOk

`func (o *PickupOrder) GetPickupOk() (*TaskProperties, bool)`

GetPickupOk returns a tuple with the Pickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickup

`func (o *PickupOrder) SetPickup(v TaskProperties)`

SetPickup sets Pickup field to given value.


### GetProperties

`func (o *PickupOrder) GetProperties() OrderProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PickupOrder) GetPropertiesOk() (*OrderProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PickupOrder) SetProperties(v OrderProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PickupOrder) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


