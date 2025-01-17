# RouteStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleId** | **string** | The unique identifier of the vehicle that is used to execute the route. Only a single route can be provided for each vehicle. | 
**Start** | **time.Time** | The point in time when the route should start. The start time will be respected as closely as possible. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). The date must not be before &#x60;1970-01-01T00:00:00+00:00&#x60; nor after &#x60;2037-12-31T23:59:59+00:00&#x60;. The date must provide an offset to UTC. | 
**Tasks** | Pointer to [**[]TaskStructure**](TaskStructure.md) | A sequence of tasks scheduled on the route. | [optional] [default to []]
**Breaks** | Pointer to [**[]BreakStructure**](BreakStructure.md) | A list of breaks scheduled on the route. When omitted, reconstruction will make a best effort to schedule breaks in order to satisfy the break settings. | [optional] [default to []]

## Methods

### NewRouteStructure

`func NewRouteStructure(vehicleId string, start time.Time, ) *RouteStructure`

NewRouteStructure instantiates a new RouteStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteStructureWithDefaults

`func NewRouteStructureWithDefaults() *RouteStructure`

NewRouteStructureWithDefaults instantiates a new RouteStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleId

`func (o *RouteStructure) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *RouteStructure) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *RouteStructure) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.


### GetStart

`func (o *RouteStructure) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *RouteStructure) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *RouteStructure) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetTasks

`func (o *RouteStructure) GetTasks() []TaskStructure`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *RouteStructure) GetTasksOk() (*[]TaskStructure, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *RouteStructure) SetTasks(v []TaskStructure)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *RouteStructure) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetBreaks

`func (o *RouteStructure) GetBreaks() []BreakStructure`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *RouteStructure) GetBreaksOk() (*[]BreakStructure, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *RouteStructure) SetBreaks(v []BreakStructure)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *RouteStructure) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


