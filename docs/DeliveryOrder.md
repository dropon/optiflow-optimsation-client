# DeliveryOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the order. This must be unique across all orders. | 
**Delivery** | [**TaskProperties**](TaskProperties.md) |  | 
**Properties** | Pointer to [**OrderProperties**](OrderProperties.md) |  | [optional] 

## Methods

### NewDeliveryOrder

`func NewDeliveryOrder(id string, delivery TaskProperties, ) *DeliveryOrder`

NewDeliveryOrder instantiates a new DeliveryOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOrderWithDefaults

`func NewDeliveryOrderWithDefaults() *DeliveryOrder`

NewDeliveryOrderWithDefaults instantiates a new DeliveryOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryOrder) SetId(v string)`

SetId sets Id field to given value.


### GetDelivery

`func (o *DeliveryOrder) GetDelivery() TaskProperties`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *DeliveryOrder) GetDeliveryOk() (*TaskProperties, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *DeliveryOrder) SetDelivery(v TaskProperties)`

SetDelivery sets Delivery field to given value.


### GetProperties

`func (o *DeliveryOrder) GetProperties() OrderProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeliveryOrder) GetPropertiesOk() (*OrderProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeliveryOrder) SetProperties(v OrderProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeliveryOrder) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


