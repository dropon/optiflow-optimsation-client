# VehicleCosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerHour** | **float64** | Specifies the cost for every hour the vehicle is used. | 
**PerKilometer** | **float64** | Specifies the cost for every kilometer driven with the vehicle. | 
**PerStop** | Pointer to **float64** | Specifies the cost for every stop on the route of the vehicle. | [optional] 
**Fixed** | Pointer to **float64** | Specifies the fixed cost for assigning a route to the vehicle. | [optional] [default to 0]
**Overtimes** | Pointer to [**[]OvertimeCost**](OvertimeCost.md) | A list of overtime costs that describe an increasing cost if the route duration exceeds a threshold. For each exceeded threshold, the extra fixed cost and the additional cost for the extra hours contribute to the total cost of the route. | [optional] [default to []]
**Overdistances** | Pointer to [**[]OverdistanceCost**](OverdistanceCost.md) | A list of overdistance costs that describe an increasing cost if the total distance of a route exceeds a threshold. For each exceeded threshold, the extra fixed cost and the additional cost for the extra kilometers contribute to the total cost of the route. | [optional] [default to []]
**Overstops** | Pointer to [**[]OverstopCost**](OverstopCost.md) | A list of overstop costs that describe an increasing cost if the number of stops of a route exceeds a threshold. For each exceeded threshold, the extra fixed cost and the additional cost for the extra stops contribute to the total cost of the route. | [optional] [default to []]
**RepositioningEfforts** | Pointer to [**[]RepositioningEffortCost**](RepositioningEffortCost.md) | Specifies a list of repositioning effort costs that describe an increasing cost if the route&#39;s repositioning effort exceeds a threshold.  When two orders are loaded into the same compartment and delivered in the same order, we refer to them as a non-last-in-first-out (non-LIFO) pair. Any non-LIFO pair of orders requires repositioning in the vehicle, as the last picked-up order obstructs the first order that needs to be delivered. The effort involved in this repositioning is the minimum effort required for the two orders. The total repositioning effort for the route is the sum of the repositioning efforts for all non-LIFO pairs of orders. | [optional] [default to []]

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


### GetPerStop

`func (o *VehicleCosts) GetPerStop() float64`

GetPerStop returns the PerStop field if non-nil, zero value otherwise.

### GetPerStopOk

`func (o *VehicleCosts) GetPerStopOk() (*float64, bool)`

GetPerStopOk returns a tuple with the PerStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerStop

`func (o *VehicleCosts) SetPerStop(v float64)`

SetPerStop sets PerStop field to given value.

### HasPerStop

`func (o *VehicleCosts) HasPerStop() bool`

HasPerStop returns a boolean if a field has been set.

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

### GetOverdistances

`func (o *VehicleCosts) GetOverdistances() []OverdistanceCost`

GetOverdistances returns the Overdistances field if non-nil, zero value otherwise.

### GetOverdistancesOk

`func (o *VehicleCosts) GetOverdistancesOk() (*[]OverdistanceCost, bool)`

GetOverdistancesOk returns a tuple with the Overdistances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdistances

`func (o *VehicleCosts) SetOverdistances(v []OverdistanceCost)`

SetOverdistances sets Overdistances field to given value.

### HasOverdistances

`func (o *VehicleCosts) HasOverdistances() bool`

HasOverdistances returns a boolean if a field has been set.

### GetOverstops

`func (o *VehicleCosts) GetOverstops() []OverstopCost`

GetOverstops returns the Overstops field if non-nil, zero value otherwise.

### GetOverstopsOk

`func (o *VehicleCosts) GetOverstopsOk() (*[]OverstopCost, bool)`

GetOverstopsOk returns a tuple with the Overstops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverstops

`func (o *VehicleCosts) SetOverstops(v []OverstopCost)`

SetOverstops sets Overstops field to given value.

### HasOverstops

`func (o *VehicleCosts) HasOverstops() bool`

HasOverstops returns a boolean if a field has been set.

### GetRepositioningEfforts

`func (o *VehicleCosts) GetRepositioningEfforts() []RepositioningEffortCost`

GetRepositioningEfforts returns the RepositioningEfforts field if non-nil, zero value otherwise.

### GetRepositioningEffortsOk

`func (o *VehicleCosts) GetRepositioningEffortsOk() (*[]RepositioningEffortCost, bool)`

GetRepositioningEffortsOk returns a tuple with the RepositioningEfforts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositioningEfforts

`func (o *VehicleCosts) SetRepositioningEfforts(v []RepositioningEffortCost)`

SetRepositioningEfforts sets RepositioningEfforts field to given value.

### HasRepositioningEfforts

`func (o *VehicleCosts) HasRepositioningEfforts() bool`

HasRepositioningEfforts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


