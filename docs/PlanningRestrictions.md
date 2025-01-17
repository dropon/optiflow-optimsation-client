# PlanningRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleTripPerRoute** | Pointer to **bool** | Restricts the number of trips per route to one.  See [here](./concepts/route-structure-control) for more information. | [optional] [default to false]
**SingleDepotPerRoute** | Pointer to **bool** | Restricts the number of different depot locations in the route of a vehicle to at most one. The vehicle start and end locations do not count here.  See [here](./concepts/route-structure-control) for more information. | [optional] [default to false]
**MixedLoadingProhibitions** | Pointer to [**[]MixedLoadingProhibition**](MixedLoadingProhibition.md) | Defines restrictions regarding which load categories of transports are disallowed to be mixed on the same trip.  | [optional] 

## Methods

### NewPlanningRestrictions

`func NewPlanningRestrictions() *PlanningRestrictions`

NewPlanningRestrictions instantiates a new PlanningRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanningRestrictionsWithDefaults

`func NewPlanningRestrictionsWithDefaults() *PlanningRestrictions`

NewPlanningRestrictionsWithDefaults instantiates a new PlanningRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleTripPerRoute

`func (o *PlanningRestrictions) GetSingleTripPerRoute() bool`

GetSingleTripPerRoute returns the SingleTripPerRoute field if non-nil, zero value otherwise.

### GetSingleTripPerRouteOk

`func (o *PlanningRestrictions) GetSingleTripPerRouteOk() (*bool, bool)`

GetSingleTripPerRouteOk returns a tuple with the SingleTripPerRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTripPerRoute

`func (o *PlanningRestrictions) SetSingleTripPerRoute(v bool)`

SetSingleTripPerRoute sets SingleTripPerRoute field to given value.

### HasSingleTripPerRoute

`func (o *PlanningRestrictions) HasSingleTripPerRoute() bool`

HasSingleTripPerRoute returns a boolean if a field has been set.

### GetSingleDepotPerRoute

`func (o *PlanningRestrictions) GetSingleDepotPerRoute() bool`

GetSingleDepotPerRoute returns the SingleDepotPerRoute field if non-nil, zero value otherwise.

### GetSingleDepotPerRouteOk

`func (o *PlanningRestrictions) GetSingleDepotPerRouteOk() (*bool, bool)`

GetSingleDepotPerRouteOk returns a tuple with the SingleDepotPerRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleDepotPerRoute

`func (o *PlanningRestrictions) SetSingleDepotPerRoute(v bool)`

SetSingleDepotPerRoute sets SingleDepotPerRoute field to given value.

### HasSingleDepotPerRoute

`func (o *PlanningRestrictions) HasSingleDepotPerRoute() bool`

HasSingleDepotPerRoute returns a boolean if a field has been set.

### GetMixedLoadingProhibitions

`func (o *PlanningRestrictions) GetMixedLoadingProhibitions() []MixedLoadingProhibition`

GetMixedLoadingProhibitions returns the MixedLoadingProhibitions field if non-nil, zero value otherwise.

### GetMixedLoadingProhibitionsOk

`func (o *PlanningRestrictions) GetMixedLoadingProhibitionsOk() (*[]MixedLoadingProhibition, bool)`

GetMixedLoadingProhibitionsOk returns a tuple with the MixedLoadingProhibitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixedLoadingProhibitions

`func (o *PlanningRestrictions) SetMixedLoadingProhibitions(v []MixedLoadingProhibition)`

SetMixedLoadingProhibitions sets MixedLoadingProhibitions field to given value.

### HasMixedLoadingProhibitions

`func (o *PlanningRestrictions) HasMixedLoadingProhibitions() bool`

HasMixedLoadingProhibitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


