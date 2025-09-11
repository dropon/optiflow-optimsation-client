# RouteEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** | The unique identifier of the location where the route ends. | [optional] 
**Approach** | Pointer to [**Leg**](Leg.md) |  | [optional] 
**Arrival** | Pointer to **time.Time** | The point in time when the vehicle arrives at its end location. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**Duration** | Pointer to **int32** | The duration [s] between the arrival at the end location and the end of the route. | [optional] 
**End** | Pointer to **time.Time** | The point in time when the route ends. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 

## Methods

### NewRouteEnd

`func NewRouteEnd() *RouteEnd`

NewRouteEnd instantiates a new RouteEnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteEndWithDefaults

`func NewRouteEndWithDefaults() *RouteEnd`

NewRouteEndWithDefaults instantiates a new RouteEnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *RouteEnd) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *RouteEnd) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *RouteEnd) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *RouteEnd) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetApproach

`func (o *RouteEnd) GetApproach() Leg`

GetApproach returns the Approach field if non-nil, zero value otherwise.

### GetApproachOk

`func (o *RouteEnd) GetApproachOk() (*Leg, bool)`

GetApproachOk returns a tuple with the Approach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproach

`func (o *RouteEnd) SetApproach(v Leg)`

SetApproach sets Approach field to given value.

### HasApproach

`func (o *RouteEnd) HasApproach() bool`

HasApproach returns a boolean if a field has been set.

### GetArrival

`func (o *RouteEnd) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *RouteEnd) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *RouteEnd) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.

### HasArrival

`func (o *RouteEnd) HasArrival() bool`

HasArrival returns a boolean if a field has been set.

### GetDuration

`func (o *RouteEnd) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RouteEnd) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RouteEnd) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RouteEnd) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnd

`func (o *RouteEnd) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *RouteEnd) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *RouteEnd) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *RouteEnd) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


