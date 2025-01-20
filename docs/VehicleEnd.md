# VehicleEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** | The unique identifier of the location where a route assigned to the vehicle must end. | 
**LatestEndTime** | **time.Time** | The latest point in time a route assigned to the vehicle may end. This must not be earlier than the vehicle&#39;s earliest start time. When used in conjunction with a latest start time, the latest end time must not be earlier than the latest start time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before &#x60;1970-01-01T00:00:00+00:00&#x60; nor after &#x60;2037-12-31T23:59:59+00:00&#x60;. The date must provide an offset to UTC. | 
**Duration** | Pointer to **int32** | Describes how long [s] it takes for the vehicle to wrap up at its end location before the route ends. | [optional] [default to 0]

## Methods

### NewVehicleEnd

`func NewVehicleEnd(locationId string, latestEndTime time.Time, ) *VehicleEnd`

NewVehicleEnd instantiates a new VehicleEnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleEndWithDefaults

`func NewVehicleEndWithDefaults() *VehicleEnd`

NewVehicleEndWithDefaults instantiates a new VehicleEnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *VehicleEnd) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *VehicleEnd) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *VehicleEnd) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetLatestEndTime

`func (o *VehicleEnd) GetLatestEndTime() time.Time`

GetLatestEndTime returns the LatestEndTime field if non-nil, zero value otherwise.

### GetLatestEndTimeOk

`func (o *VehicleEnd) GetLatestEndTimeOk() (*time.Time, bool)`

GetLatestEndTimeOk returns a tuple with the LatestEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEndTime

`func (o *VehicleEnd) SetLatestEndTime(v time.Time)`

SetLatestEndTime sets LatestEndTime field to given value.


### GetDuration

`func (o *VehicleEnd) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VehicleEnd) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VehicleEnd) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VehicleEnd) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


