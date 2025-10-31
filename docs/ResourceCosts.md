# ResourceCosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fixed** | Pointer to **float64** | A one-time fixed cost incurred when this resource is used, regardless of how many vehicles use it. | [optional] [default to 0]
**PerVehicle** | Pointer to **float64** | The cost charged for each vehicle that uses this resource. | [optional] [default to 0]
**PerLayover** | Pointer to **float64** | The cost charged when a layover occurs between consecutive vehicles sharing this resource. This cost is weighed against the cost for these vehicles to end and start at their provided start and end locations. | [optional] [default to 0]
**Overvehicles** | Pointer to [**[]OvervehicleCost**](OvervehicleCost.md) | A list of overvehicle costs that describe an increasing cost if the number of vehicles that is assigned to this resource exceeds a threshold. For each exceeded threshold, the extra fixed cost and the additional cost for the extra vehicle contribute to the total cost of the route. | [optional] [default to []]

## Methods

### NewResourceCosts

`func NewResourceCosts() *ResourceCosts`

NewResourceCosts instantiates a new ResourceCosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCostsWithDefaults

`func NewResourceCostsWithDefaults() *ResourceCosts`

NewResourceCostsWithDefaults instantiates a new ResourceCosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixed

`func (o *ResourceCosts) GetFixed() float64`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *ResourceCosts) GetFixedOk() (*float64, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *ResourceCosts) SetFixed(v float64)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *ResourceCosts) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetPerVehicle

`func (o *ResourceCosts) GetPerVehicle() float64`

GetPerVehicle returns the PerVehicle field if non-nil, zero value otherwise.

### GetPerVehicleOk

`func (o *ResourceCosts) GetPerVehicleOk() (*float64, bool)`

GetPerVehicleOk returns a tuple with the PerVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerVehicle

`func (o *ResourceCosts) SetPerVehicle(v float64)`

SetPerVehicle sets PerVehicle field to given value.

### HasPerVehicle

`func (o *ResourceCosts) HasPerVehicle() bool`

HasPerVehicle returns a boolean if a field has been set.

### GetPerLayover

`func (o *ResourceCosts) GetPerLayover() float64`

GetPerLayover returns the PerLayover field if non-nil, zero value otherwise.

### GetPerLayoverOk

`func (o *ResourceCosts) GetPerLayoverOk() (*float64, bool)`

GetPerLayoverOk returns a tuple with the PerLayover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerLayover

`func (o *ResourceCosts) SetPerLayover(v float64)`

SetPerLayover sets PerLayover field to given value.

### HasPerLayover

`func (o *ResourceCosts) HasPerLayover() bool`

HasPerLayover returns a boolean if a field has been set.

### GetOvervehicles

`func (o *ResourceCosts) GetOvervehicles() []OvervehicleCost`

GetOvervehicles returns the Overvehicles field if non-nil, zero value otherwise.

### GetOvervehiclesOk

`func (o *ResourceCosts) GetOvervehiclesOk() (*[]OvervehicleCost, bool)`

GetOvervehiclesOk returns a tuple with the Overvehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvervehicles

`func (o *ResourceCosts) SetOvervehicles(v []OvervehicleCost)`

SetOvervehicles sets Overvehicles field to given value.

### HasOvervehicles

`func (o *ResourceCosts) HasOvervehicles() bool`

HasOvervehicles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


