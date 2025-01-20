# VehicleCosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerHour** | **float64** | Specifies the cost for every hour the vehicle is used. | 
**PerKilometer** | **float64** | Specifies the cost for every kilometer driven with the vehicle. | 
**Fixed** | Pointer to **float64** | Specifies the fixed cost for assigning a route to the vehicle. | [optional] [default to 0]
**Overtimes** | Pointer to [**[]OvertimeCost**](OvertimeCost.md) | A list of overtime costs that describe an increasing cost if the route duration goes above a threshold. | [optional] [default to []]

## Methods

### NewVehicleCosts

`func NewVehicleCosts(perHour float64, perKilometer float64, ) *VehicleCosts`

NewVehicleCosts instantiates a new VehicleCosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleCostsWithDefaults

`func NewVehicleCostsWithDefaults() *VehicleCosts`

NewVehicleCostsWithDefaults instantiates a new VehicleCosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerHour

`func (o *VehicleCosts) GetPerHour() float64`

GetPerHour returns the PerHour field if non-nil, zero value otherwise.

### GetPerHourOk

`func (o *VehicleCosts) GetPerHourOk() (*float64, bool)`

GetPerHourOk returns a tuple with the PerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerHour

`func (o *VehicleCosts) SetPerHour(v float64)`

SetPerHour sets PerHour field to given value.


### GetPerKilometer

`func (o *VehicleCosts) GetPerKilometer() float64`

GetPerKilometer returns the PerKilometer field if non-nil, zero value otherwise.

### GetPerKilometerOk

`func (o *VehicleCosts) GetPerKilometerOk() (*float64, bool)`

GetPerKilometerOk returns a tuple with the PerKilometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerKilometer

`func (o *VehicleCosts) SetPerKilometer(v float64)`

SetPerKilometer sets PerKilometer field to given value.


### GetFixed

`func (o *VehicleCosts) GetFixed() float64`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *VehicleCosts) GetFixedOk() (*float64, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *VehicleCosts) SetFixed(v float64)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *VehicleCosts) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetOvertimes

`func (o *VehicleCosts) GetOvertimes() []OvertimeCost`

GetOvertimes returns the Overtimes field if non-nil, zero value otherwise.

### GetOvertimesOk

`func (o *VehicleCosts) GetOvertimesOk() (*[]OvertimeCost, bool)`

GetOvertimesOk returns a tuple with the Overtimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvertimes

`func (o *VehicleCosts) SetOvertimes(v []OvertimeCost)`

SetOvertimes sets Overtimes field to given value.

### HasOvertimes

`func (o *VehicleCosts) HasOvertimes() bool`

HasOvertimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


