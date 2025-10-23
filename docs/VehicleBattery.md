# VehicleBattery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | **float64** | The total battery capacity [kWh]. | 
**StateOfCharge** | Pointer to [**BatteryStateOfCharge**](BatteryStateOfCharge.md) |  | [optional] 
**Consumption** | [**BatteryConsumption**](BatteryConsumption.md) |  | 

## Methods

### NewVehicleBattery

`func NewVehicleBattery(capacity float64, consumption BatteryConsumption, ) *VehicleBattery`

NewVehicleBattery instantiates a new VehicleBattery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleBatteryWithDefaults

`func NewVehicleBatteryWithDefaults() *VehicleBattery`

NewVehicleBatteryWithDefaults instantiates a new VehicleBattery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *VehicleBattery) GetCapacity() float64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VehicleBattery) GetCapacityOk() (*float64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VehicleBattery) SetCapacity(v float64)`

SetCapacity sets Capacity field to given value.


### GetStateOfCharge

`func (o *VehicleBattery) GetStateOfCharge() BatteryStateOfCharge`

GetStateOfCharge returns the StateOfCharge field if non-nil, zero value otherwise.

### GetStateOfChargeOk

`func (o *VehicleBattery) GetStateOfChargeOk() (*BatteryStateOfCharge, bool)`

GetStateOfChargeOk returns a tuple with the StateOfCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOfCharge

`func (o *VehicleBattery) SetStateOfCharge(v BatteryStateOfCharge)`

SetStateOfCharge sets StateOfCharge field to given value.

### HasStateOfCharge

`func (o *VehicleBattery) HasStateOfCharge() bool`

HasStateOfCharge returns a boolean if a field has been set.

### GetConsumption

`func (o *VehicleBattery) GetConsumption() BatteryConsumption`

GetConsumption returns the Consumption field if non-nil, zero value otherwise.

### GetConsumptionOk

`func (o *VehicleBattery) GetConsumptionOk() (*BatteryConsumption, bool)`

GetConsumptionOk returns a tuple with the Consumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumption

`func (o *VehicleBattery) SetConsumption(v BatteryConsumption)`

SetConsumption sets Consumption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


