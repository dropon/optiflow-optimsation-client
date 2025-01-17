# PlanSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the plan. | 
**NumberOfLocations** | **int32** | Number of locations in the plan. | 
**NumberOfVehicles** | **int32** | Number of vehicles in the plan. | 
**NumberOfTransports** | **int32** | Number of transports in the plan. | 
**NumberOfRoutes** | **int32** | Number of routes in the plan. | 
**Description** | Pointer to **string** | The description of the plan. | [optional] 
**UpdateTime** | **time.Time** | Time of the latest update of the plan formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). | 
**CreationTime** | **time.Time** | Time of the creation of the plan formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). | 

## Methods

### NewPlanSummary

`func NewPlanSummary(id string, numberOfLocations int32, numberOfVehicles int32, numberOfTransports int32, numberOfRoutes int32, updateTime time.Time, creationTime time.Time, ) *PlanSummary`

NewPlanSummary instantiates a new PlanSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanSummaryWithDefaults

`func NewPlanSummaryWithDefaults() *PlanSummary`

NewPlanSummaryWithDefaults instantiates a new PlanSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanSummary) SetId(v string)`

SetId sets Id field to given value.


### GetNumberOfLocations

`func (o *PlanSummary) GetNumberOfLocations() int32`

GetNumberOfLocations returns the NumberOfLocations field if non-nil, zero value otherwise.

### GetNumberOfLocationsOk

`func (o *PlanSummary) GetNumberOfLocationsOk() (*int32, bool)`

GetNumberOfLocationsOk returns a tuple with the NumberOfLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLocations

`func (o *PlanSummary) SetNumberOfLocations(v int32)`

SetNumberOfLocations sets NumberOfLocations field to given value.


### GetNumberOfVehicles

`func (o *PlanSummary) GetNumberOfVehicles() int32`

GetNumberOfVehicles returns the NumberOfVehicles field if non-nil, zero value otherwise.

### GetNumberOfVehiclesOk

`func (o *PlanSummary) GetNumberOfVehiclesOk() (*int32, bool)`

GetNumberOfVehiclesOk returns a tuple with the NumberOfVehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVehicles

`func (o *PlanSummary) SetNumberOfVehicles(v int32)`

SetNumberOfVehicles sets NumberOfVehicles field to given value.


### GetNumberOfTransports

`func (o *PlanSummary) GetNumberOfTransports() int32`

GetNumberOfTransports returns the NumberOfTransports field if non-nil, zero value otherwise.

### GetNumberOfTransportsOk

`func (o *PlanSummary) GetNumberOfTransportsOk() (*int32, bool)`

GetNumberOfTransportsOk returns a tuple with the NumberOfTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTransports

`func (o *PlanSummary) SetNumberOfTransports(v int32)`

SetNumberOfTransports sets NumberOfTransports field to given value.


### GetNumberOfRoutes

`func (o *PlanSummary) GetNumberOfRoutes() int32`

GetNumberOfRoutes returns the NumberOfRoutes field if non-nil, zero value otherwise.

### GetNumberOfRoutesOk

`func (o *PlanSummary) GetNumberOfRoutesOk() (*int32, bool)`

GetNumberOfRoutesOk returns a tuple with the NumberOfRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRoutes

`func (o *PlanSummary) SetNumberOfRoutes(v int32)`

SetNumberOfRoutes sets NumberOfRoutes field to given value.


### GetDescription

`func (o *PlanSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PlanSummary) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PlanSummary) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PlanSummary) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.


### GetCreationTime

`func (o *PlanSummary) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PlanSummary) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PlanSummary) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


