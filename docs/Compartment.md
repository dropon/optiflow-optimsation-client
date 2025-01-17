# Compartment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the compartment, which must be distinct among all compartments within the same vehicle. | 
**LoadingStrategy** | Pointer to [**CompartmentLoadingStrategy**](CompartmentLoadingStrategy.md) |  | [optional] [default to NONE]
**Categories** | Pointer to **[]string** | A list of categories the compartment belongs to that can be used to describe constraints or rules. | [optional] [default to []]

## Methods

### NewCompartment

`func NewCompartment(id string, ) *Compartment`

NewCompartment instantiates a new Compartment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartmentWithDefaults

`func NewCompartmentWithDefaults() *Compartment`

NewCompartmentWithDefaults instantiates a new Compartment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Compartment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Compartment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Compartment) SetId(v string)`

SetId sets Id field to given value.


### GetLoadingStrategy

`func (o *Compartment) GetLoadingStrategy() CompartmentLoadingStrategy`

GetLoadingStrategy returns the LoadingStrategy field if non-nil, zero value otherwise.

### GetLoadingStrategyOk

`func (o *Compartment) GetLoadingStrategyOk() (*CompartmentLoadingStrategy, bool)`

GetLoadingStrategyOk returns a tuple with the LoadingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingStrategy

`func (o *Compartment) SetLoadingStrategy(v CompartmentLoadingStrategy)`

SetLoadingStrategy sets LoadingStrategy field to given value.

### HasLoadingStrategy

`func (o *Compartment) HasLoadingStrategy() bool`

HasLoadingStrategy returns a boolean if a field has been set.

### GetCategories

`func (o *Compartment) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Compartment) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Compartment) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Compartment) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


