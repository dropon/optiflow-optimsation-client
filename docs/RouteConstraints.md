# RouteConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumDuration** | Pointer to **int32** | Restricts the maximum duration [s] of the route assigned to the vehicle. | [optional] 
**MaximumDistance** | Pointer to **int32** | Restricts the maximum distance [m] travelled on the route assigned to the vehicle. | [optional] 

## Methods

### NewRouteConstraints

`func NewRouteConstraints() *RouteConstraints`

NewRouteConstraints instantiates a new RouteConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteConstraintsWithDefaults

`func NewRouteConstraintsWithDefaults() *RouteConstraints`

NewRouteConstraintsWithDefaults instantiates a new RouteConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumDuration

`func (o *RouteConstraints) GetMaximumDuration() int32`

GetMaximumDuration returns the MaximumDuration field if non-nil, zero value otherwise.

### GetMaximumDurationOk

`func (o *RouteConstraints) GetMaximumDurationOk() (*int32, bool)`

GetMaximumDurationOk returns a tuple with the MaximumDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDuration

`func (o *RouteConstraints) SetMaximumDuration(v int32)`

SetMaximumDuration sets MaximumDuration field to given value.

### HasMaximumDuration

`func (o *RouteConstraints) HasMaximumDuration() bool`

HasMaximumDuration returns a boolean if a field has been set.

### GetMaximumDistance

`func (o *RouteConstraints) GetMaximumDistance() int32`

GetMaximumDistance returns the MaximumDistance field if non-nil, zero value otherwise.

### GetMaximumDistanceOk

`func (o *RouteConstraints) GetMaximumDistanceOk() (*int32, bool)`

GetMaximumDistanceOk returns a tuple with the MaximumDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDistance

`func (o *RouteConstraints) SetMaximumDistance(v int32)`

SetMaximumDistance sets MaximumDistance field to given value.

### HasMaximumDistance

`func (o *RouteConstraints) HasMaximumDistance() bool`

HasMaximumDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


