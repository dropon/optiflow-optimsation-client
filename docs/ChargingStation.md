# ChargingStation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the charging station. This must be unique across all charging stations of one location. | 
**MaximumPower** | **float64** | The maximum power output [kW] of the charging station. | 

## Methods

### NewChargingStation

`func NewChargingStation(id string, maximumPower float64, ) *ChargingStation`

NewChargingStation instantiates a new ChargingStation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingStationWithDefaults

`func NewChargingStationWithDefaults() *ChargingStation`

NewChargingStationWithDefaults instantiates a new ChargingStation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChargingStation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargingStation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargingStation) SetId(v string)`

SetId sets Id field to given value.


### GetMaximumPower

`func (o *ChargingStation) GetMaximumPower() float64`

GetMaximumPower returns the MaximumPower field if non-nil, zero value otherwise.

### GetMaximumPowerOk

`func (o *ChargingStation) GetMaximumPowerOk() (*float64, bool)`

GetMaximumPowerOk returns a tuple with the MaximumPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPower

`func (o *ChargingStation) SetMaximumPower(v float64)`

SetMaximumPower sets MaximumPower field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


