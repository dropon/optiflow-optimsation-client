# StopConcurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleSlots** | Pointer to [**[]VehicleSlot**](VehicleSlot.md) | A list of available vehicle slots. | [optional] [default to []]
**ViolationCostPerExtraSlot** | **float64** | The additional cost incurred for using an extra vehicle slot beyond the available ones. This must be greater than or equal to the cost of each vehicle slot. | 
**MinimumBufferDuration** | Pointer to **int32** | The minimum duration [s] between the end of the last appointment of a stop and the start of the first appointment of another stop assigned to the same vehicle slot. | [optional] [default to 0]

## Methods

### NewStopConcurrency

`func NewStopConcurrency(violationCostPerExtraSlot float64, ) *StopConcurrency`

NewStopConcurrency instantiates a new StopConcurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopConcurrencyWithDefaults

`func NewStopConcurrencyWithDefaults() *StopConcurrency`

NewStopConcurrencyWithDefaults instantiates a new StopConcurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleSlots

`func (o *StopConcurrency) GetVehicleSlots() []VehicleSlot`

GetVehicleSlots returns the VehicleSlots field if non-nil, zero value otherwise.

### GetVehicleSlotsOk

`func (o *StopConcurrency) GetVehicleSlotsOk() (*[]VehicleSlot, bool)`

GetVehicleSlotsOk returns a tuple with the VehicleSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleSlots

`func (o *StopConcurrency) SetVehicleSlots(v []VehicleSlot)`

SetVehicleSlots sets VehicleSlots field to given value.

### HasVehicleSlots

`func (o *StopConcurrency) HasVehicleSlots() bool`

HasVehicleSlots returns a boolean if a field has been set.

### GetViolationCostPerExtraSlot

`func (o *StopConcurrency) GetViolationCostPerExtraSlot() float64`

GetViolationCostPerExtraSlot returns the ViolationCostPerExtraSlot field if non-nil, zero value otherwise.

### GetViolationCostPerExtraSlotOk

`func (o *StopConcurrency) GetViolationCostPerExtraSlotOk() (*float64, bool)`

GetViolationCostPerExtraSlotOk returns a tuple with the ViolationCostPerExtraSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCostPerExtraSlot

`func (o *StopConcurrency) SetViolationCostPerExtraSlot(v float64)`

SetViolationCostPerExtraSlot sets ViolationCostPerExtraSlot field to given value.


### GetMinimumBufferDuration

`func (o *StopConcurrency) GetMinimumBufferDuration() int32`

GetMinimumBufferDuration returns the MinimumBufferDuration field if non-nil, zero value otherwise.

### GetMinimumBufferDurationOk

`func (o *StopConcurrency) GetMinimumBufferDurationOk() (*int32, bool)`

GetMinimumBufferDurationOk returns a tuple with the MinimumBufferDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBufferDuration

`func (o *StopConcurrency) SetMinimumBufferDuration(v int32)`

SetMinimumBufferDuration sets MinimumBufferDuration field to given value.

### HasMinimumBufferDuration

`func (o *StopConcurrency) HasMinimumBufferDuration() bool`

HasMinimumBufferDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


