# Vehicle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the vehicle. | 
**Start** | [**VehicleStart**](VehicleStart.md) |  | 
**End** | [**VehicleEnd**](VehicleEnd.md) |  | 
**Routing** | [**VehicleRouting**](VehicleRouting.md) |  | 
**Costs** | [**VehicleCosts**](VehicleCosts.md) |  | 
**Preferences** | Pointer to [**VehiclePreferences**](VehiclePreferences.md) |  | [optional] 
**Breaks** | Pointer to [**BreakSettings**](BreakSettings.md) |  | [optional] 
**Compartments** | Pointer to [**[]Compartment**](Compartment.md) | A list of compartments available for loading orders. Orders placed into the same compartment must follow the compartment&#39;s loading strategy, which may impose restrictions on the unloading sequence. If orders are loaded into separate compartments, no restrictions apply to the unloading sequence. When unspecified or empty, the vehicle is assumed to have a single compartment with loading strategy &#x60;NONE&#x60; (resp. &#x60;LAST_IN_FIRST_OUT&#x60;) when optimizing pickup and/or delivery orders (resp. pickup-delivery orders). **Note that this default behaviour will change. The vehicle will be assumed to have a single compartment with loading strategy &#x60;NONE&#x60; regardless of the specified order types. When relying on the default &#x60;LAST_IN_FIRST_OUT&#x60; behaviour for pickup-delivery orders, please already specify a single compartment with loading strategy &#x60;LAST_IN_FIRST_OUT&#x60; in order to ensure this change will not impact your optimizations.** | [optional] [default to []]
**Constraints** | Pointer to [**VehicleConstraints**](VehicleConstraints.md) |  | [optional] 
**Categories** | Pointer to **[]string** | A list of categories the vehicle belongs to that can be used to describe constraints or rules. | [optional] [default to []]

## Methods

### NewVehicle

`func NewVehicle(id string, start VehicleStart, end VehicleEnd, routing VehicleRouting, costs VehicleCosts, ) *Vehicle`

NewVehicle instantiates a new Vehicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleWithDefaults

`func NewVehicleWithDefaults() *Vehicle`

NewVehicleWithDefaults instantiates a new Vehicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vehicle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vehicle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vehicle) SetId(v string)`

SetId sets Id field to given value.


### GetStart

`func (o *Vehicle) GetStart() VehicleStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Vehicle) GetStartOk() (*VehicleStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Vehicle) SetStart(v VehicleStart)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Vehicle) GetEnd() VehicleEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Vehicle) GetEndOk() (*VehicleEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Vehicle) SetEnd(v VehicleEnd)`

SetEnd sets End field to given value.


### GetRouting

`func (o *Vehicle) GetRouting() VehicleRouting`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *Vehicle) GetRoutingOk() (*VehicleRouting, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *Vehicle) SetRouting(v VehicleRouting)`

SetRouting sets Routing field to given value.


### GetCosts

`func (o *Vehicle) GetCosts() VehicleCosts`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *Vehicle) GetCostsOk() (*VehicleCosts, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *Vehicle) SetCosts(v VehicleCosts)`

SetCosts sets Costs field to given value.


### GetPreferences

`func (o *Vehicle) GetPreferences() VehiclePreferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *Vehicle) GetPreferencesOk() (*VehiclePreferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *Vehicle) SetPreferences(v VehiclePreferences)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *Vehicle) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetBreaks

`func (o *Vehicle) GetBreaks() BreakSettings`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *Vehicle) GetBreaksOk() (*BreakSettings, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *Vehicle) SetBreaks(v BreakSettings)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *Vehicle) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.

### GetCompartments

`func (o *Vehicle) GetCompartments() []Compartment`

GetCompartments returns the Compartments field if non-nil, zero value otherwise.

### GetCompartmentsOk

`func (o *Vehicle) GetCompartmentsOk() (*[]Compartment, bool)`

GetCompartmentsOk returns a tuple with the Compartments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartments

`func (o *Vehicle) SetCompartments(v []Compartment)`

SetCompartments sets Compartments field to given value.

### HasCompartments

`func (o *Vehicle) HasCompartments() bool`

HasCompartments returns a boolean if a field has been set.

### GetConstraints

`func (o *Vehicle) GetConstraints() VehicleConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Vehicle) GetConstraintsOk() (*VehicleConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Vehicle) SetConstraints(v VehicleConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *Vehicle) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetCategories

`func (o *Vehicle) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Vehicle) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Vehicle) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Vehicle) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


