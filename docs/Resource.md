# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the resource. | 
**MinimumBufferDuration** | Pointer to **int32** | The minimum duration [s] that must be scheduled between the end of one route and the start of another route using this resource. This allows for transition time, cleaning, or preparation between uses. | [optional] [default to 0]
**Categories** | **[]string** | A list of categories the shared resource belongs to that can be used to describe constraints or rules. | 
**Costs** | Pointer to [**ResourceCosts**](ResourceCosts.md) |  | [optional] 

## Methods

### NewResource

`func NewResource(id string, categories []string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.


### GetMinimumBufferDuration

`func (o *Resource) GetMinimumBufferDuration() int32`

GetMinimumBufferDuration returns the MinimumBufferDuration field if non-nil, zero value otherwise.

### GetMinimumBufferDurationOk

`func (o *Resource) GetMinimumBufferDurationOk() (*int32, bool)`

GetMinimumBufferDurationOk returns a tuple with the MinimumBufferDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBufferDuration

`func (o *Resource) SetMinimumBufferDuration(v int32)`

SetMinimumBufferDuration sets MinimumBufferDuration field to given value.

### HasMinimumBufferDuration

`func (o *Resource) HasMinimumBufferDuration() bool`

HasMinimumBufferDuration returns a boolean if a field has been set.

### GetCategories

`func (o *Resource) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Resource) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Resource) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetCosts

`func (o *Resource) GetCosts() ResourceCosts`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *Resource) GetCostsOk() (*ResourceCosts, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *Resource) SetCosts(v ResourceCosts)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *Resource) HasCosts() bool`

HasCosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


