# Charging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** | The point in time when the charging starts. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**Duration** | Pointer to **int32** | The duration [s] of charging. | [optional] 
**ChargingStationId** | Pointer to **string** | The unique identifier of the charging station where charging is taking place. | [optional] 
**End** | Pointer to **time.Time** | The point in time when charging ends. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 

## Methods

### NewCharging

`func NewCharging() *Charging`

NewCharging instantiates a new Charging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingWithDefaults

`func NewChargingWithDefaults() *Charging`

NewChargingWithDefaults instantiates a new Charging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *Charging) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Charging) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Charging) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Charging) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetDuration

`func (o *Charging) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Charging) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Charging) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Charging) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetChargingStationId

`func (o *Charging) GetChargingStationId() string`

GetChargingStationId returns the ChargingStationId field if non-nil, zero value otherwise.

### GetChargingStationIdOk

`func (o *Charging) GetChargingStationIdOk() (*string, bool)`

GetChargingStationIdOk returns a tuple with the ChargingStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingStationId

`func (o *Charging) SetChargingStationId(v string)`

SetChargingStationId sets ChargingStationId field to given value.

### HasChargingStationId

`func (o *Charging) HasChargingStationId() bool`

HasChargingStationId returns a boolean if a field has been set.

### GetEnd

`func (o *Charging) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Charging) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Charging) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Charging) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


