# TaskProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** | The unique identifier of the location where the task must be scheduled. | 
**TimeSlotIds** | Pointer to **[]string** | A list of unique identifiers of the time slots of the location that can be used to execute this task. When empty all time slots can be used. If more than 50 time slots are specified for the location, the list must not be empty. | [optional] [default to {}]
**Duration** | Pointer to **int32** | The duration [s] it takes to execute this task. | [optional] [default to 0]
**Categories** | Pointer to **[]string** | A list of categories the task belongs to that can be used to describe constraints or rules. | [optional] [default to {}]

## Methods

### NewTaskProperties

`func NewTaskProperties(locationId string, ) *TaskProperties`

NewTaskProperties instantiates a new TaskProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPropertiesWithDefaults

`func NewTaskPropertiesWithDefaults() *TaskProperties`

NewTaskPropertiesWithDefaults instantiates a new TaskProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *TaskProperties) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *TaskProperties) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *TaskProperties) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetTimeSlotIds

`func (o *TaskProperties) GetTimeSlotIds() []string`

GetTimeSlotIds returns the TimeSlotIds field if non-nil, zero value otherwise.

### GetTimeSlotIdsOk

`func (o *TaskProperties) GetTimeSlotIdsOk() (*[]string, bool)`

GetTimeSlotIdsOk returns a tuple with the TimeSlotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotIds

`func (o *TaskProperties) SetTimeSlotIds(v []string)`

SetTimeSlotIds sets TimeSlotIds field to given value.

### HasTimeSlotIds

`func (o *TaskProperties) HasTimeSlotIds() bool`

HasTimeSlotIds returns a boolean if a field has been set.

### GetDuration

`func (o *TaskProperties) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TaskProperties) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TaskProperties) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TaskProperties) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetCategories

`func (o *TaskProperties) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TaskProperties) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TaskProperties) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *TaskProperties) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


