# ChargingStation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the charging station. This must be unique across all charging stations of one location. | 
**MaximumPower** | **float64** | The maximum power output [kW] of the charging station. | 
**PreparationDuration** | Pointer to **int32** | A period of time [s] required before charging can begin after the vehicle arrives at the charging station. This duration accounts for activities such as positioning the vehicle and connecting the charger. | [optional] 
**CompletionDuration** | Pointer to **int32** | A period of time [s] required after charging is completed before the vehicle can depart. This duration accounts for activities such as disconnecting the charger and preparing the vehicle for departure. | [optional] 

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


### GetPreparationDuration

`func (o *ChargingStation) GetPreparationDuration() int32`

GetPreparationDuration returns the PreparationDuration field if non-nil, zero value otherwise.

### GetPreparationDurationOk

`func (o *ChargingStation) GetPreparationDurationOk() (*int32, bool)`

GetPreparationDurationOk returns a tuple with the PreparationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationDuration

`func (o *ChargingStation) SetPreparationDuration(v int32)`

SetPreparationDuration sets PreparationDuration field to given value.

### HasPreparationDuration

`func (o *ChargingStation) HasPreparationDuration() bool`

HasPreparationDuration returns a boolean if a field has been set.

### GetCompletionDuration

`func (o *ChargingStation) GetCompletionDuration() int32`

GetCompletionDuration returns the CompletionDuration field if non-nil, zero value otherwise.

### GetCompletionDurationOk

`func (o *ChargingStation) GetCompletionDurationOk() (*int32, bool)`

GetCompletionDurationOk returns a tuple with the CompletionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDuration

`func (o *ChargingStation) SetCompletionDuration(v int32)`

SetCompletionDuration sets CompletionDuration field to given value.

### HasCompletionDuration

`func (o *ChargingStation) HasCompletionDuration() bool`

HasCompletionDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


