# RouteStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** | The unique identifier of the start location of the route. This is the start location of the assigned vehicle. | [optional] 
**Start** | Pointer to **time.Time** | The point in time when the route starts. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**Duration** | Pointer to **int32** | The duration [s] between the start of the route and the departure at the start location. | [optional] 
**Departure** | Pointer to **time.Time** | The point in time when the vehicle departs at its start location. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 

## Methods

### NewRouteStart

`func NewRouteStart() *RouteStart`

NewRouteStart instantiates a new RouteStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteStartWithDefaults

`func NewRouteStartWithDefaults() *RouteStart`

NewRouteStartWithDefaults instantiates a new RouteStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *RouteStart) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *RouteStart) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *RouteStart) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *RouteStart) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetStart

`func (o *RouteStart) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *RouteStart) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *RouteStart) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *RouteStart) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetDuration

`func (o *RouteStart) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RouteStart) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RouteStart) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RouteStart) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDeparture

`func (o *RouteStart) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *RouteStart) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *RouteStart) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.

### HasDeparture

`func (o *RouteStart) HasDeparture() bool`

HasDeparture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


