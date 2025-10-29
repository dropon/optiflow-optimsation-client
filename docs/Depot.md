# Depot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the depot. | 
**LocationId** | **string** | The unique identifier of the location where the depot is situated. | 
**TimeSlotIds** | Pointer to **[]string** | A list of unique identifiers of the time slots of the depot location that can be used to execute tasks at this depot. When empty all time slots can be used. If more than 50 time slots are specified for the depot location, the list must not be empty. | [optional] [default to {}]
**Categories** | Pointer to **[]string** | A list of categories the depot belongs to that can be used to describe constraints or rules. | [optional] [default to {}]

## Methods

### NewDepot

`func NewDepot(id string, locationId string, ) *Depot`

NewDepot instantiates a new Depot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotWithDefaults

`func NewDepotWithDefaults() *Depot`

NewDepotWithDefaults instantiates a new Depot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Depot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Depot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Depot) SetId(v string)`

SetId sets Id field to given value.


### GetLocationId

`func (o *Depot) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Depot) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Depot) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetTimeSlotIds

`func (o *Depot) GetTimeSlotIds() []string`

GetTimeSlotIds returns the TimeSlotIds field if non-nil, zero value otherwise.

### GetTimeSlotIdsOk

`func (o *Depot) GetTimeSlotIdsOk() (*[]string, bool)`

GetTimeSlotIdsOk returns a tuple with the TimeSlotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotIds

`func (o *Depot) SetTimeSlotIds(v []string)`

SetTimeSlotIds sets TimeSlotIds field to given value.

### HasTimeSlotIds

`func (o *Depot) HasTimeSlotIds() bool`

HasTimeSlotIds returns a boolean if a field has been set.

### GetCategories

`func (o *Depot) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Depot) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Depot) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Depot) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


