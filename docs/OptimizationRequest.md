# OptimizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | [**OptimizationSettings**](OptimizationSettings.md) |  | 
**Locations** | [**[]Location**](Location.md) | A list of locations where goods have to be picked up or delivered, or where vehicles are located. Please note that the upper bound on the number of locations is a technical limit. Check your individual price plan or contract to see which limits apply. | 
**Orders** | [**Orders**](Orders.md) |  | 
**Vehicles** | [**[]Vehicle**](Vehicle.md) | A list of vehicles that can be used by optimization to schedule routes. Optimization can only assign a single route to a vehicle and must respect the properties and constraints of the vehicle. Please note that the upper bound on the number of vehicles is a technical limit. Check your individual price plan or contract to see which limits apply. | 
**Resources** | Pointer to [**[]Resource**](Resource.md) | A list of resources that can be assigned to vehicles. Resources are shared assets that cannot be used simultaneously by multiple vehicles. When vehicles share a resource, their routes must be scheduled such that they do not overlap in time. Please note that resources count as separate assets in your license allocation. The total asset count for any optimization includes the sum of routes and their distinct assigned resources. The upper bound on number of resources is a technical limit. Check your individual price plan or contract to see which limits apply. | [optional] 
**Depots** | Pointer to [**[]Depot**](Depot.md) | A list of depots where pickup orders can be delivered to or delivery orders can be picked up from. When providing pickup or delivery orders, at least one depot must be specified. When stopping at a depot, all pickup orders present in the vehicle are unloaded first. Afterwards, delivery orders can be loaded into the vehicle. Those must be delivered before stopping at the next depot. | [optional] [default to []]
**Routes** | Pointer to [**[]RouteStructure**](RouteStructure.md) | A list of routes that should be reconstructed prior to optimization. Reconstruction ensures all constraints are met and may involve removing orders, changing breaks, or adjusting the start time of the route. After the reconstruction, the optimization will try to improve the routes. The structure of the routes can be changed by the optimization as long as the constraints are satisfied. Check your individual price plan or contract to see whether or not the request may contain routes. | [optional] [default to []]
**Constraints** | Pointer to [**Constraints**](Constraints.md) |  | [optional] 
**Rules** | Pointer to [**Rules**](Rules.md) |  | [optional] 
**Metadata** | Pointer to [**OptimizationRequestMetadata**](OptimizationRequestMetadata.md) |  | [optional] 

## Methods

### NewOptimizationRequest

`func NewOptimizationRequest(settings OptimizationSettings, locations []Location, orders Orders, vehicles []Vehicle, ) *OptimizationRequest`

NewOptimizationRequest instantiates a new OptimizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationRequestWithDefaults

`func NewOptimizationRequestWithDefaults() *OptimizationRequest`

NewOptimizationRequestWithDefaults instantiates a new OptimizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *OptimizationRequest) GetSettings() OptimizationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *OptimizationRequest) GetSettingsOk() (*OptimizationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *OptimizationRequest) SetSettings(v OptimizationSettings)`

SetSettings sets Settings field to given value.


### GetLocations

`func (o *OptimizationRequest) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *OptimizationRequest) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *OptimizationRequest) SetLocations(v []Location)`

SetLocations sets Locations field to given value.


### GetOrders

`func (o *OptimizationRequest) GetOrders() Orders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OptimizationRequest) GetOrdersOk() (*Orders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OptimizationRequest) SetOrders(v Orders)`

SetOrders sets Orders field to given value.


### GetVehicles

`func (o *OptimizationRequest) GetVehicles() []Vehicle`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *OptimizationRequest) GetVehiclesOk() (*[]Vehicle, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *OptimizationRequest) SetVehicles(v []Vehicle)`

SetVehicles sets Vehicles field to given value.


### GetResources

`func (o *OptimizationRequest) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *OptimizationRequest) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *OptimizationRequest) SetResources(v []Resource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *OptimizationRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDepots

`func (o *OptimizationRequest) GetDepots() []Depot`

GetDepots returns the Depots field if non-nil, zero value otherwise.

### GetDepotsOk

`func (o *OptimizationRequest) GetDepotsOk() (*[]Depot, bool)`

GetDepotsOk returns a tuple with the Depots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepots

`func (o *OptimizationRequest) SetDepots(v []Depot)`

SetDepots sets Depots field to given value.

### HasDepots

`func (o *OptimizationRequest) HasDepots() bool`

HasDepots returns a boolean if a field has been set.

### GetRoutes

`func (o *OptimizationRequest) GetRoutes() []RouteStructure`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *OptimizationRequest) GetRoutesOk() (*[]RouteStructure, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *OptimizationRequest) SetRoutes(v []RouteStructure)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *OptimizationRequest) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetConstraints

`func (o *OptimizationRequest) GetConstraints() Constraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *OptimizationRequest) GetConstraintsOk() (*Constraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *OptimizationRequest) SetConstraints(v Constraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *OptimizationRequest) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetRules

`func (o *OptimizationRequest) GetRules() Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *OptimizationRequest) GetRulesOk() (*Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *OptimizationRequest) SetRules(v Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *OptimizationRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetMetadata

`func (o *OptimizationRequest) GetMetadata() OptimizationRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OptimizationRequest) GetMetadataOk() (*OptimizationRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OptimizationRequest) SetMetadata(v OptimizationRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OptimizationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


