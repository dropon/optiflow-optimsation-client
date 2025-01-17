# DepotLocationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceTimePerPickupStop** | Pointer to **int32** | The service time [s] that is required each time this location is visited in order to pick up goods. The location-dependent service time represents, for example, the time to enter the premises or to register at a depot. Besides a location-dependent service time, the user may specify additional vehicle-dependent and transport-dependent service times at the vehicles and the transports respectively. | [optional] [default to 0]
**ServiceTimePerDeliveryStop** | Pointer to **int32** | The service time [s] that is required each time this location is visited in order to deliver goods. The location-dependent service time represents, for example, the time to enter the premises or to register at a depot. Besides a location-dependent service time, the user may specify additional vehicle-dependent and transport-dependent service times at the vehicles and the transports respectively. | [optional] [default to 0]

## Methods

### NewDepotLocationAttributes

`func NewDepotLocationAttributes() *DepotLocationAttributes`

NewDepotLocationAttributes instantiates a new DepotLocationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotLocationAttributesWithDefaults

`func NewDepotLocationAttributesWithDefaults() *DepotLocationAttributes`

NewDepotLocationAttributesWithDefaults instantiates a new DepotLocationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceTimePerPickupStop

`func (o *DepotLocationAttributes) GetServiceTimePerPickupStop() int32`

GetServiceTimePerPickupStop returns the ServiceTimePerPickupStop field if non-nil, zero value otherwise.

### GetServiceTimePerPickupStopOk

`func (o *DepotLocationAttributes) GetServiceTimePerPickupStopOk() (*int32, bool)`

GetServiceTimePerPickupStopOk returns a tuple with the ServiceTimePerPickupStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTimePerPickupStop

`func (o *DepotLocationAttributes) SetServiceTimePerPickupStop(v int32)`

SetServiceTimePerPickupStop sets ServiceTimePerPickupStop field to given value.

### HasServiceTimePerPickupStop

`func (o *DepotLocationAttributes) HasServiceTimePerPickupStop() bool`

HasServiceTimePerPickupStop returns a boolean if a field has been set.

### GetServiceTimePerDeliveryStop

`func (o *DepotLocationAttributes) GetServiceTimePerDeliveryStop() int32`

GetServiceTimePerDeliveryStop returns the ServiceTimePerDeliveryStop field if non-nil, zero value otherwise.

### GetServiceTimePerDeliveryStopOk

`func (o *DepotLocationAttributes) GetServiceTimePerDeliveryStopOk() (*int32, bool)`

GetServiceTimePerDeliveryStopOk returns a tuple with the ServiceTimePerDeliveryStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTimePerDeliveryStop

`func (o *DepotLocationAttributes) SetServiceTimePerDeliveryStop(v int32)`

SetServiceTimePerDeliveryStop sets ServiceTimePerDeliveryStop field to given value.

### HasServiceTimePerDeliveryStop

`func (o *DepotLocationAttributes) HasServiceTimePerDeliveryStop() bool`

HasServiceTimePerDeliveryStop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


