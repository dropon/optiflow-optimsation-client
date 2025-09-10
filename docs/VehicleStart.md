# VehicleStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** | The unique identifier of the location where a route assigned to the vehicle must start. When omitted, the route starts at the location of the first stop. | [optional] 
**EarliestStartTime** | **time.Time** | The earliest point in time a route assigned to the vehicle may start. This must be not be later than the vehicle&#39;s latest end time. When used in conjunction with a latest start time, the earliest start time must not be later than the latest start time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before &#x60;1970-01-01T00:00:00+00:00&#x60; nor after &#x60;2037-12-31T23:59:59+00:00&#x60;. The date must provide an offset to UTC. | 
**LatestStartTime** | Pointer to **time.Time** | The latest point in time a route assigned to the vehicle may start. This must not be earlier than the vehicle&#39;s earliest start time nor later than the vehicle&#39;s earliest end time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before &#x60;1970-01-01T00:00:00+00:00&#x60; nor after &#x60;2037-12-31T23:59:59+00:00&#x60;. The date must provide an offset to UTC. | [optional] 
**Duration** | Pointer to **int32** | Describes how long [s] it takes for the vehicle to depart at its start location after the route starts. | [optional] [default to 0]

## Methods

### NewVehicleStart

`func NewVehicleStart(earliestStartTime time.Time, ) *VehicleStart`

NewVehicleStart instantiates a new VehicleStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStartWithDefaults

`func NewVehicleStartWithDefaults() *VehicleStart`

NewVehicleStartWithDefaults instantiates a new VehicleStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *VehicleStart) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *VehicleStart) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *VehicleStart) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *VehicleStart) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetEarliestStartTime

`func (o *VehicleStart) GetEarliestStartTime() time.Time`

GetEarliestStartTime returns the EarliestStartTime field if non-nil, zero value otherwise.

### GetEarliestStartTimeOk

`func (o *VehicleStart) GetEarliestStartTimeOk() (*time.Time, bool)`

GetEarliestStartTimeOk returns a tuple with the EarliestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestStartTime

`func (o *VehicleStart) SetEarliestStartTime(v time.Time)`

SetEarliestStartTime sets EarliestStartTime field to given value.


### GetLatestStartTime

`func (o *VehicleStart) GetLatestStartTime() time.Time`

GetLatestStartTime returns the LatestStartTime field if non-nil, zero value otherwise.

### GetLatestStartTimeOk

`func (o *VehicleStart) GetLatestStartTimeOk() (*time.Time, bool)`

GetLatestStartTimeOk returns a tuple with the LatestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestStartTime

`func (o *VehicleStart) SetLatestStartTime(v time.Time)`

SetLatestStartTime sets LatestStartTime field to given value.

### HasLatestStartTime

`func (o *VehicleStart) HasLatestStartTime() bool`

HasLatestStartTime returns a boolean if a field has been set.

### GetDuration

`func (o *VehicleStart) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VehicleStart) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VehicleStart) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VehicleStart) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


