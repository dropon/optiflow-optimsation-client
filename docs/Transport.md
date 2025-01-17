# Transport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID across all transports. The ID does not influence the result. | 
**Quantities** | Pointer to **[]int32** | A list of quantity dimensions of goods that have to be transported from pickup location to delivery location. The maximum length of this list is 100. That is, up to 100 different quantity dimensions (e.g. number of pallets, weight, volume, etc.) can be distinguished. Transports can only be executed by a vehicle with a higher (or an equal) maximum capacity in every quantity dimension. The length of this list has to be the same for all transports and all capacities of all vehicles. If and only if this list of quantities is empty for all transports, the capacities of each vehicle must be empty. | [optional] 
**PickupLocationId** | **string** | The ID of the location where goods have to be picked up. | 
**PickupServiceTime** | Pointer to **int32** | The transport-dependent service time [s] that is required to pick up the transport goods at the pickup location. | [optional] [default to 0]
**DeliveryLocationId** | **string** | The ID of the location where goods have to be delivered to. | 
**DeliveryServiceTime** | Pointer to **int32** | The transport-dependent service time [s] that is required to deliver the transport goods at the delivery location. | [optional] [default to 0]
**RequiredEquipment** | Pointer to **[]string** | A list of required vehicle equipment. If empty, no equipment is required. A vehicle can only be assigned to the transport if this list is a subset of (or equal to) the vehicle&#39;s equipment. | [optional] 
**Priority** | Pointer to **NullableInt32** | The priority of this transport. 0 is the lowest priority, 9 the highest. This field is only considered during the optimization if considerTransportPriorities is set. In this case every transport must have a priority set otherwise an exception is thrown. If considerTransportPriorities is set to false either every or no transport may have a priority set.  See [here](./concepts/transport-priorities) for more information. | [optional] 
**LoadCategory** | Pointer to **string** | The load category of this transport. To specify a load category is useful if there are transports with one load category that shall not be mixed with transports with another load category on one trip. If two load categories are mutually exclusive on a trip, this can be specified as a mixed loading prohibition. Transports with conflicting load categories will not be planned together on one trip.  See [here](./concepts/mixed-loading-prohibition) for more information. | [optional] 

## Methods

### NewTransport

`func NewTransport(id string, pickupLocationId string, deliveryLocationId string, ) *Transport`

NewTransport instantiates a new Transport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportWithDefaults

`func NewTransportWithDefaults() *Transport`

NewTransportWithDefaults instantiates a new Transport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transport) SetId(v string)`

SetId sets Id field to given value.


### GetQuantities

`func (o *Transport) GetQuantities() []int32`

GetQuantities returns the Quantities field if non-nil, zero value otherwise.

### GetQuantitiesOk

`func (o *Transport) GetQuantitiesOk() (*[]int32, bool)`

GetQuantitiesOk returns a tuple with the Quantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantities

`func (o *Transport) SetQuantities(v []int32)`

SetQuantities sets Quantities field to given value.

### HasQuantities

`func (o *Transport) HasQuantities() bool`

HasQuantities returns a boolean if a field has been set.

### GetPickupLocationId

`func (o *Transport) GetPickupLocationId() string`

GetPickupLocationId returns the PickupLocationId field if non-nil, zero value otherwise.

### GetPickupLocationIdOk

`func (o *Transport) GetPickupLocationIdOk() (*string, bool)`

GetPickupLocationIdOk returns a tuple with the PickupLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupLocationId

`func (o *Transport) SetPickupLocationId(v string)`

SetPickupLocationId sets PickupLocationId field to given value.


### GetPickupServiceTime

`func (o *Transport) GetPickupServiceTime() int32`

GetPickupServiceTime returns the PickupServiceTime field if non-nil, zero value otherwise.

### GetPickupServiceTimeOk

`func (o *Transport) GetPickupServiceTimeOk() (*int32, bool)`

GetPickupServiceTimeOk returns a tuple with the PickupServiceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupServiceTime

`func (o *Transport) SetPickupServiceTime(v int32)`

SetPickupServiceTime sets PickupServiceTime field to given value.

### HasPickupServiceTime

`func (o *Transport) HasPickupServiceTime() bool`

HasPickupServiceTime returns a boolean if a field has been set.

### GetDeliveryLocationId

`func (o *Transport) GetDeliveryLocationId() string`

GetDeliveryLocationId returns the DeliveryLocationId field if non-nil, zero value otherwise.

### GetDeliveryLocationIdOk

`func (o *Transport) GetDeliveryLocationIdOk() (*string, bool)`

GetDeliveryLocationIdOk returns a tuple with the DeliveryLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLocationId

`func (o *Transport) SetDeliveryLocationId(v string)`

SetDeliveryLocationId sets DeliveryLocationId field to given value.


### GetDeliveryServiceTime

`func (o *Transport) GetDeliveryServiceTime() int32`

GetDeliveryServiceTime returns the DeliveryServiceTime field if non-nil, zero value otherwise.

### GetDeliveryServiceTimeOk

`func (o *Transport) GetDeliveryServiceTimeOk() (*int32, bool)`

GetDeliveryServiceTimeOk returns a tuple with the DeliveryServiceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryServiceTime

`func (o *Transport) SetDeliveryServiceTime(v int32)`

SetDeliveryServiceTime sets DeliveryServiceTime field to given value.

### HasDeliveryServiceTime

`func (o *Transport) HasDeliveryServiceTime() bool`

HasDeliveryServiceTime returns a boolean if a field has been set.

### GetRequiredEquipment

`func (o *Transport) GetRequiredEquipment() []string`

GetRequiredEquipment returns the RequiredEquipment field if non-nil, zero value otherwise.

### GetRequiredEquipmentOk

`func (o *Transport) GetRequiredEquipmentOk() (*[]string, bool)`

GetRequiredEquipmentOk returns a tuple with the RequiredEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredEquipment

`func (o *Transport) SetRequiredEquipment(v []string)`

SetRequiredEquipment sets RequiredEquipment field to given value.

### HasRequiredEquipment

`func (o *Transport) HasRequiredEquipment() bool`

HasRequiredEquipment returns a boolean if a field has been set.

### GetPriority

`func (o *Transport) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Transport) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Transport) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Transport) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *Transport) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *Transport) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetLoadCategory

`func (o *Transport) GetLoadCategory() string`

GetLoadCategory returns the LoadCategory field if non-nil, zero value otherwise.

### GetLoadCategoryOk

`func (o *Transport) GetLoadCategoryOk() (*string, bool)`

GetLoadCategoryOk returns a tuple with the LoadCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadCategory

`func (o *Transport) SetLoadCategory(v string)`

SetLoadCategory sets LoadCategory field to given value.

### HasLoadCategory

`func (o *Transport) HasLoadCategory() bool`

HasLoadCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


