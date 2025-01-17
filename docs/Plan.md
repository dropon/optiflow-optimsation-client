# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the plan. It is generated when the plan is created. | [readonly] 
**Description** | Pointer to **string** | A description of the plan. | [optional] 
**Locations** | [**[]Location**](Location.md) | A list of depot or customer locations that may be referenced by vehicles, transports and stops. A location is either a depot location or a customer location. Depot locations act as trip delimiters. Each location must be referenced by another object. If a request contains a location not referenced by any other object, the request will be rejected. | 
**Vehicles** | [**[]Vehicle**](Vehicle.md) | A list of vehicles that can be used to transport goods. | 
**Drivers** | Pointer to [**[]Driver**](Driver.md) | A list of drivers. A driver is always assigned to a specific vehicle. In turn, a vehicle can but does not need to have a driver assigned to it. If a driver is assigned to a vehicle, the driver&#39;s restrictions apply, such as its limited availability. If no drivers are specified, the drivers of all vehicles are always available. | [optional] 
**Transports** | [**[]Transport**](Transport.md) | A list of transports, that is, orders to transport goods from one location to another location. Depending on your subscription, a more restrictive value for maximum number of transport may apply. Check request limits of your subscription. | 
**PlanningHorizon** | Pointer to [**TimeInterval**](TimeInterval.md) | The planning horizon for the plan, described by start and end date and time. All routes have to start and end within this planning horizon. All opening intervals outside of this planning horizon are not considered by the algorithm. If specified, the planning horizon is restricted to a maximum duration of two weeks. If not specified, the planning horizon is infinite. If no other time interval is specified within this plan, the planning horizon is required. | [optional] 
**Restrictions** | Pointer to [**PlanningRestrictions**](PlanningRestrictions.md) |  | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) | A list of routes. A route contains a sequence of stops. It specifies where and in which order goods are to be picked up or delivered. Each stop can be assigned to a trip. A route is subdivided into trips. Each route has at least one trip and a trip consists of at least two stops. At the beginning and end of each trip the vehicle does not carry any load. A trip starts at the vehicle start location or at a depot location, and ends at the vehicle end location or at a depot location. If routes are already given in input they are considered during planning. Transports which are already planned in an input route will also remain planned in the output routes. This might lead to violations of the routes.   See [here](./concepts/routes-and-trips) for more information. | [optional] 
**UnplannedVehicleIds** | Pointer to **[]string** | Returns the vehicle IDs that are not used in the response of an optimization operation. These vehicles are not assigned to any route. | [optional] [readonly] 
**UnplannedTransportIds** | Pointer to **[]string** | Returns the transport IDs that could not be planned in the response of an optimization operation. These transports are not part of the routes. | [optional] [readonly] 
**Warnings** | Pointer to [**[]Warning**](Warning.md) | A list of warnings concerning the validity of the result. | [optional] [readonly] 

## Methods

### NewPlan

`func NewPlan(id string, locations []Location, vehicles []Vehicle, transports []Transport, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *Plan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Plan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocations

`func (o *Plan) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Plan) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Plan) SetLocations(v []Location)`

SetLocations sets Locations field to given value.


### GetVehicles

`func (o *Plan) GetVehicles() []Vehicle`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *Plan) GetVehiclesOk() (*[]Vehicle, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *Plan) SetVehicles(v []Vehicle)`

SetVehicles sets Vehicles field to given value.


### GetDrivers

`func (o *Plan) GetDrivers() []Driver`

GetDrivers returns the Drivers field if non-nil, zero value otherwise.

### GetDriversOk

`func (o *Plan) GetDriversOk() (*[]Driver, bool)`

GetDriversOk returns a tuple with the Drivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivers

`func (o *Plan) SetDrivers(v []Driver)`

SetDrivers sets Drivers field to given value.

### HasDrivers

`func (o *Plan) HasDrivers() bool`

HasDrivers returns a boolean if a field has been set.

### GetTransports

`func (o *Plan) GetTransports() []Transport`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *Plan) GetTransportsOk() (*[]Transport, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *Plan) SetTransports(v []Transport)`

SetTransports sets Transports field to given value.


### GetPlanningHorizon

`func (o *Plan) GetPlanningHorizon() TimeInterval`

GetPlanningHorizon returns the PlanningHorizon field if non-nil, zero value otherwise.

### GetPlanningHorizonOk

`func (o *Plan) GetPlanningHorizonOk() (*TimeInterval, bool)`

GetPlanningHorizonOk returns a tuple with the PlanningHorizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningHorizon

`func (o *Plan) SetPlanningHorizon(v TimeInterval)`

SetPlanningHorizon sets PlanningHorizon field to given value.

### HasPlanningHorizon

`func (o *Plan) HasPlanningHorizon() bool`

HasPlanningHorizon returns a boolean if a field has been set.

### GetRestrictions

`func (o *Plan) GetRestrictions() PlanningRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *Plan) GetRestrictionsOk() (*PlanningRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *Plan) SetRestrictions(v PlanningRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *Plan) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetRoutes

`func (o *Plan) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Plan) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Plan) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *Plan) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetUnplannedVehicleIds

`func (o *Plan) GetUnplannedVehicleIds() []string`

GetUnplannedVehicleIds returns the UnplannedVehicleIds field if non-nil, zero value otherwise.

### GetUnplannedVehicleIdsOk

`func (o *Plan) GetUnplannedVehicleIdsOk() (*[]string, bool)`

GetUnplannedVehicleIdsOk returns a tuple with the UnplannedVehicleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplannedVehicleIds

`func (o *Plan) SetUnplannedVehicleIds(v []string)`

SetUnplannedVehicleIds sets UnplannedVehicleIds field to given value.

### HasUnplannedVehicleIds

`func (o *Plan) HasUnplannedVehicleIds() bool`

HasUnplannedVehicleIds returns a boolean if a field has been set.

### GetUnplannedTransportIds

`func (o *Plan) GetUnplannedTransportIds() []string`

GetUnplannedTransportIds returns the UnplannedTransportIds field if non-nil, zero value otherwise.

### GetUnplannedTransportIdsOk

`func (o *Plan) GetUnplannedTransportIdsOk() (*[]string, bool)`

GetUnplannedTransportIdsOk returns a tuple with the UnplannedTransportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplannedTransportIds

`func (o *Plan) SetUnplannedTransportIds(v []string)`

SetUnplannedTransportIds sets UnplannedTransportIds field to given value.

### HasUnplannedTransportIds

`func (o *Plan) HasUnplannedTransportIds() bool`

HasUnplannedTransportIds returns a boolean if a field has been set.

### GetWarnings

`func (o *Plan) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Plan) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Plan) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Plan) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


