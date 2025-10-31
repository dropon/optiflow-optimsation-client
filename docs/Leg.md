# Leg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartLocationId** | Pointer to **string** | The unique identifier of the location where the vehicle departs. | [optional] 
**Departure** | Pointer to **time.Time** | The point in time when the vehicle departs from the start location. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**EndLocationId** | Pointer to **string** | The unique identifier of the location where the vehicle arrives. | [optional] 
**Arrival** | Pointer to **time.Time** | The point in time when the vehicle arrives at the end location. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**Breaks** | Pointer to [**[]Break**](Break.md) | A list of breaks taken while travelling from the start location to the end location. | [optional] [default to []]
**Distance** | Pointer to **int32** | The distance [m] travelled between the start location and end location. | [optional] 
**DrivingDuration** | Pointer to **int32** | The duration [s] it takes to drive from the start location to the end location. | [optional] 

## Methods

### NewLeg

`func NewLeg() *Leg`

NewLeg instantiates a new Leg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegWithDefaults

`func NewLegWithDefaults() *Leg`

NewLegWithDefaults instantiates a new Leg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartLocationId

`func (o *Leg) GetStartLocationId() string`

GetStartLocationId returns the StartLocationId field if non-nil, zero value otherwise.

### GetStartLocationIdOk

`func (o *Leg) GetStartLocationIdOk() (*string, bool)`

GetStartLocationIdOk returns a tuple with the StartLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocationId

`func (o *Leg) SetStartLocationId(v string)`

SetStartLocationId sets StartLocationId field to given value.

### HasStartLocationId

`func (o *Leg) HasStartLocationId() bool`

HasStartLocationId returns a boolean if a field has been set.

### GetDeparture

`func (o *Leg) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *Leg) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *Leg) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.

### HasDeparture

`func (o *Leg) HasDeparture() bool`

HasDeparture returns a boolean if a field has been set.

### GetEndLocationId

`func (o *Leg) GetEndLocationId() string`

GetEndLocationId returns the EndLocationId field if non-nil, zero value otherwise.

### GetEndLocationIdOk

`func (o *Leg) GetEndLocationIdOk() (*string, bool)`

GetEndLocationIdOk returns a tuple with the EndLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocationId

`func (o *Leg) SetEndLocationId(v string)`

SetEndLocationId sets EndLocationId field to given value.

### HasEndLocationId

`func (o *Leg) HasEndLocationId() bool`

HasEndLocationId returns a boolean if a field has been set.

### GetArrival

`func (o *Leg) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *Leg) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *Leg) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.

### HasArrival

`func (o *Leg) HasArrival() bool`

HasArrival returns a boolean if a field has been set.

### GetBreaks

`func (o *Leg) GetBreaks() []Break`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *Leg) GetBreaksOk() (*[]Break, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *Leg) SetBreaks(v []Break)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *Leg) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.

### GetDistance

`func (o *Leg) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Leg) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Leg) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *Leg) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetDrivingDuration

`func (o *Leg) GetDrivingDuration() int32`

GetDrivingDuration returns the DrivingDuration field if non-nil, zero value otherwise.

### GetDrivingDurationOk

`func (o *Leg) GetDrivingDurationOk() (*int32, bool)`

GetDrivingDurationOk returns a tuple with the DrivingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingDuration

`func (o *Leg) SetDrivingDuration(v int32)`

SetDrivingDuration sets DrivingDuration field to given value.

### HasDrivingDuration

`func (o *Leg) HasDrivingDuration() bool`

HasDrivingDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


