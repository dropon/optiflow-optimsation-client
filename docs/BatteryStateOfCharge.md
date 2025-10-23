# BatteryStateOfCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initial** | Pointer to **float64** | The initial battery energy level as a fraction of total capacity (0.0 &#x3D; empty, 1.0 &#x3D; fully charged). | [optional] [default to 1]
**Minimum** | Pointer to **float64** | The minimum allowed battery energy level as a fraction of total capacity (0.0 &#x3D; empty, 1.0 &#x3D; fully charged). | [optional] [default to 0]

## Methods

### NewBatteryStateOfCharge

`func NewBatteryStateOfCharge() *BatteryStateOfCharge`

NewBatteryStateOfCharge instantiates a new BatteryStateOfCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatteryStateOfChargeWithDefaults

`func NewBatteryStateOfChargeWithDefaults() *BatteryStateOfCharge`

NewBatteryStateOfChargeWithDefaults instantiates a new BatteryStateOfCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitial

`func (o *BatteryStateOfCharge) GetInitial() float64`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *BatteryStateOfCharge) GetInitialOk() (*float64, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *BatteryStateOfCharge) SetInitial(v float64)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *BatteryStateOfCharge) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetMinimum

`func (o *BatteryStateOfCharge) GetMinimum() float64`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *BatteryStateOfCharge) GetMinimumOk() (*float64, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *BatteryStateOfCharge) SetMinimum(v float64)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *BatteryStateOfCharge) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


