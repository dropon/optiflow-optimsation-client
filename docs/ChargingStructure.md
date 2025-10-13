# ChargingStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** | The point in time when charging starts. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before &#x60;1970-01-01T00:00:00+00:00&#x60; nor after &#x60;2037-12-31T23:59:59+00:00&#x60;. The date must provide an offset to UTC. | 
**Duration** | **int32** | Charging duration [s]. | 
**LocationId** | **string** | The unique identifier of the location where charging is taking place. | 
**ChargingStationId** | **string** | The unique identifier of the charging station where charging is taking place. | 

## Methods

### NewChargingStructure

`func NewChargingStructure(start time.Time, duration int32, locationId string, chargingStationId string, ) *ChargingStructure`

NewChargingStructure instantiates a new ChargingStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingStructureWithDefaults

`func NewChargingStructureWithDefaults() *ChargingStructure`

NewChargingStructureWithDefaults instantiates a new ChargingStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *ChargingStructure) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ChargingStructure) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ChargingStructure) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetDuration

`func (o *ChargingStructure) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ChargingStructure) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ChargingStructure) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetLocationId

`func (o *ChargingStructure) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ChargingStructure) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ChargingStructure) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetChargingStationId

`func (o *ChargingStructure) GetChargingStationId() string`

GetChargingStationId returns the ChargingStationId field if non-nil, zero value otherwise.

### GetChargingStationIdOk

`func (o *ChargingStructure) GetChargingStationIdOk() (*string, bool)`

GetChargingStationIdOk returns a tuple with the ChargingStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingStationId

`func (o *ChargingStructure) SetChargingStationId(v string)`

SetChargingStationId sets ChargingStationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


