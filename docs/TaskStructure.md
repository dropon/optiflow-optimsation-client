# TaskStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | The unique identifier of the order whose pickup or delivery is described by this task. | 
**Type** | [**TaskType**](TaskType.md) |  | 
**TimeSlotId** | Pointer to **string** | The unique identifier of the time slot assigned to an appointment to execute the task. Required when time slots are defined at the task&#39;s location. | [optional] 
**DepotId** | Pointer to **string** | The unique identifier of the depot in case the task is a pickup or a delivery at a depot. Required for tasks scheduled at a depot, otherwise this must be omitted. | [optional] 
**CompartmentId** | Pointer to **string** | The unique identifier of the compartment that the order needs to be loaded in or unloaded from. Required when compartments are defined for the vehicle. | [optional] 

## Methods

### NewTaskStructure

`func NewTaskStructure(orderId string, type_ TaskType, ) *TaskStructure`

NewTaskStructure instantiates a new TaskStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStructureWithDefaults

`func NewTaskStructureWithDefaults() *TaskStructure`

NewTaskStructureWithDefaults instantiates a new TaskStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *TaskStructure) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TaskStructure) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TaskStructure) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetType

`func (o *TaskStructure) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskStructure) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskStructure) SetType(v TaskType)`

SetType sets Type field to given value.


### GetTimeSlotId

`func (o *TaskStructure) GetTimeSlotId() string`

GetTimeSlotId returns the TimeSlotId field if non-nil, zero value otherwise.

### GetTimeSlotIdOk

`func (o *TaskStructure) GetTimeSlotIdOk() (*string, bool)`

GetTimeSlotIdOk returns a tuple with the TimeSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotId

`func (o *TaskStructure) SetTimeSlotId(v string)`

SetTimeSlotId sets TimeSlotId field to given value.

### HasTimeSlotId

`func (o *TaskStructure) HasTimeSlotId() bool`

HasTimeSlotId returns a boolean if a field has been set.

### GetDepotId

`func (o *TaskStructure) GetDepotId() string`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *TaskStructure) GetDepotIdOk() (*string, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *TaskStructure) SetDepotId(v string)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *TaskStructure) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetCompartmentId

`func (o *TaskStructure) GetCompartmentId() string`

GetCompartmentId returns the CompartmentId field if non-nil, zero value otherwise.

### GetCompartmentIdOk

`func (o *TaskStructure) GetCompartmentIdOk() (*string, bool)`

GetCompartmentIdOk returns a tuple with the CompartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentId

`func (o *TaskStructure) SetCompartmentId(v string)`

SetCompartmentId sets CompartmentId field to given value.

### HasCompartmentId

`func (o *TaskStructure) HasCompartmentId() bool`

HasCompartmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


