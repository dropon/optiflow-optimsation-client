# VehicleConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumLoads** | Pointer to [**[]Load**](Load.md) | A list of load dimensions that describe the capacity of the vehicle. For each dimension specified in this list, the sum of the values of the orders loaded in the vehicle must be lower than or equal to the value of the vehicle. For unspecified dimensions, the load of the vehicle is assumed to be unconstrained for this dimension. | [optional] [default to []]
**Route** | Pointer to [**RouteConstraints**](RouteConstraints.md) |  | [optional] 

## Methods

### NewVehicleConstraints

`func NewVehicleConstraints() *VehicleConstraints`

NewVehicleConstraints instantiates a new VehicleConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleConstraintsWithDefaults

`func NewVehicleConstraintsWithDefaults() *VehicleConstraints`

NewVehicleConstraintsWithDefaults instantiates a new VehicleConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumLoads

`func (o *VehicleConstraints) GetMaximumLoads() []Load`

GetMaximumLoads returns the MaximumLoads field if non-nil, zero value otherwise.

### GetMaximumLoadsOk

`func (o *VehicleConstraints) GetMaximumLoadsOk() (*[]Load, bool)`

GetMaximumLoadsOk returns a tuple with the MaximumLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLoads

`func (o *VehicleConstraints) SetMaximumLoads(v []Load)`

SetMaximumLoads sets MaximumLoads field to given value.

### HasMaximumLoads

`func (o *VehicleConstraints) HasMaximumLoads() bool`

HasMaximumLoads returns a boolean if a field has been set.

### GetRoute

`func (o *VehicleConstraints) GetRoute() RouteConstraints`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *VehicleConstraints) GetRouteOk() (*RouteConstraints, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *VehicleConstraints) SetRoute(v RouteConstraints)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *VehicleConstraints) HasRoute() bool`

HasRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


