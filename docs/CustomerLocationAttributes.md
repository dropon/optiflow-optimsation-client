# CustomerLocationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceTimePerTransportStop** | Pointer to **int32** | The service time [s] that is required each time this location is visited in order to pick goods up or to deliver them. The location-dependent service time represents, for example, the time to enter the premises or to register at a customer. Besides a location-dependent service time, the user may specify additional vehicle-dependent and transport-dependent service times at the vehicles and the transports respectively. | [optional] [default to 0]
**CustomerId** | Pointer to **string** | ID of the customer that can be set to link several customer locations with different opening intervals. Customer locations with the same customer ID may only differ in the location ID and in the opening intervals. If successive stops at customer locations share the same customer ID, the service time per transport stop of the location and of the vehicle are only considered at the first stop of the sequence. | [optional] 
**PositionInTrip** | Pointer to [**NullablePositionInTrip**](PositionInTrip.md) |  | [optional] 
**TripSectionNumber** | Pointer to **NullableInt32** | If tripSectionNumber is specified, a stop at this customer location will be visited after stops (within the same trip) at customer locations with specified lower tripSectionNumber and before stops (within the same trip) at customer locations with specified higher tripSectionNumber. Consequently, the trip section numbers of stops at customer locations must be non-decreasing within each trip.   If specified, positionInTrip must not be set for the same location.  See [here](./concepts/trip-sections-and-position-trips) for more information. | [optional] 

## Methods

### NewCustomerLocationAttributes

`func NewCustomerLocationAttributes() *CustomerLocationAttributes`

NewCustomerLocationAttributes instantiates a new CustomerLocationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerLocationAttributesWithDefaults

`func NewCustomerLocationAttributesWithDefaults() *CustomerLocationAttributes`

NewCustomerLocationAttributesWithDefaults instantiates a new CustomerLocationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceTimePerTransportStop

`func (o *CustomerLocationAttributes) GetServiceTimePerTransportStop() int32`

GetServiceTimePerTransportStop returns the ServiceTimePerTransportStop field if non-nil, zero value otherwise.

### GetServiceTimePerTransportStopOk

`func (o *CustomerLocationAttributes) GetServiceTimePerTransportStopOk() (*int32, bool)`

GetServiceTimePerTransportStopOk returns a tuple with the ServiceTimePerTransportStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTimePerTransportStop

`func (o *CustomerLocationAttributes) SetServiceTimePerTransportStop(v int32)`

SetServiceTimePerTransportStop sets ServiceTimePerTransportStop field to given value.

### HasServiceTimePerTransportStop

`func (o *CustomerLocationAttributes) HasServiceTimePerTransportStop() bool`

HasServiceTimePerTransportStop returns a boolean if a field has been set.

### GetCustomerId

`func (o *CustomerLocationAttributes) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerLocationAttributes) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerLocationAttributes) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CustomerLocationAttributes) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetPositionInTrip

`func (o *CustomerLocationAttributes) GetPositionInTrip() PositionInTrip`

GetPositionInTrip returns the PositionInTrip field if non-nil, zero value otherwise.

### GetPositionInTripOk

`func (o *CustomerLocationAttributes) GetPositionInTripOk() (*PositionInTrip, bool)`

GetPositionInTripOk returns a tuple with the PositionInTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionInTrip

`func (o *CustomerLocationAttributes) SetPositionInTrip(v PositionInTrip)`

SetPositionInTrip sets PositionInTrip field to given value.

### HasPositionInTrip

`func (o *CustomerLocationAttributes) HasPositionInTrip() bool`

HasPositionInTrip returns a boolean if a field has been set.

### SetPositionInTripNil

`func (o *CustomerLocationAttributes) SetPositionInTripNil(b bool)`

 SetPositionInTripNil sets the value for PositionInTrip to be an explicit nil

### UnsetPositionInTrip
`func (o *CustomerLocationAttributes) UnsetPositionInTrip()`

UnsetPositionInTrip ensures that no value is present for PositionInTrip, not even an explicit nil
### GetTripSectionNumber

`func (o *CustomerLocationAttributes) GetTripSectionNumber() int32`

GetTripSectionNumber returns the TripSectionNumber field if non-nil, zero value otherwise.

### GetTripSectionNumberOk

`func (o *CustomerLocationAttributes) GetTripSectionNumberOk() (*int32, bool)`

GetTripSectionNumberOk returns a tuple with the TripSectionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripSectionNumber

`func (o *CustomerLocationAttributes) SetTripSectionNumber(v int32)`

SetTripSectionNumber sets TripSectionNumber field to given value.

### HasTripSectionNumber

`func (o *CustomerLocationAttributes) HasTripSectionNumber() bool`

HasTripSectionNumber returns a boolean if a field has been set.

### SetTripSectionNumberNil

`func (o *CustomerLocationAttributes) SetTripSectionNumberNil(b bool)`

 SetTripSectionNumberNil sets the value for TripSectionNumber to be an explicit nil

### UnsetTripSectionNumber
`func (o *CustomerLocationAttributes) UnsetTripSectionNumber()`

UnsetTripSectionNumber ensures that no value is present for TripSectionNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


