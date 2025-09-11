# VehicleSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **float64** | The cost incurred when assigning at least one stop to this vehicle slot. | [optional] [default to 0]

## Methods

### NewVehicleSlot

`func NewVehicleSlot() *VehicleSlot`

NewVehicleSlot instantiates a new VehicleSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleSlotWithDefaults

`func NewVehicleSlotWithDefaults() *VehicleSlot`

NewVehicleSlotWithDefaults instantiates a new VehicleSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *VehicleSlot) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *VehicleSlot) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *VehicleSlot) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *VehicleSlot) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


