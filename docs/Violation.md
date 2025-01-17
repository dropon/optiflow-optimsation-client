# Violation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ViolationType**](ViolationType.md) |  | 
**TimeExceedance** | Pointer to **int32** | Specifies a violation that involves a time limit: Time by which the limit is exceeded. Is only filled for the following violation types: _PLANNING_HORIZON_, _ROUTE_START_INTERVAL_, _DRIVER_AVAILABILITY_, _OPENING_INTERVAL_, _MAXIMUM_TRAVEL_TIME_OF_DRIVER_, _MAXIMUM_DRIVING_TIME_OF_DRIVER_. | [optional] 
**DistanceExceedance** | Pointer to **int32** | Specifies a violation that involves a distance limit: Distance by which the limit is exceeded. Is only filled for the following violation type: _MAXIMUM_DISTANCE_. | [optional] 
**NumberOfStopsExceedance** | Pointer to **int32** | Specifies a violation that involves a stop limit: Number of stops by which the limit is exceeded. Is only filled for the following violation type: _MAXIMUM_NUMBER_OF_CUSTOMER_STOPS_. | [optional] 
**CapacityExceedance** | Pointer to **[]int32** | Specifies a violation of the capacities of the used vehicle: Exceedance of the capacity in each quantity dimension of goods. Is only filled for the following violation type: _VEHICLE_CAPACITY_. | [optional] 
**MissingEquipment** | Pointer to **[]string** | Specifies a violation for missing equipment (e.g. missing equipment of the used vehicle) that is required by the transports of the route. Is only filled for the following violation type: _VEHICLE_EQUIPMENT_. | [optional] 
**MixedLoadingProhibitions** | Pointer to [**[]MixedLoadingProhibition**](MixedLoadingProhibition.md) | Is only filled for the following violation type: _MIXED_LOADING_PROHIBITION_. | [optional] 

## Methods

### NewViolation

`func NewViolation(type_ ViolationType, ) *Violation`

NewViolation instantiates a new Violation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViolationWithDefaults

`func NewViolationWithDefaults() *Violation`

NewViolationWithDefaults instantiates a new Violation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Violation) GetType() ViolationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Violation) GetTypeOk() (*ViolationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Violation) SetType(v ViolationType)`

SetType sets Type field to given value.


### GetTimeExceedance

`func (o *Violation) GetTimeExceedance() int32`

GetTimeExceedance returns the TimeExceedance field if non-nil, zero value otherwise.

### GetTimeExceedanceOk

`func (o *Violation) GetTimeExceedanceOk() (*int32, bool)`

GetTimeExceedanceOk returns a tuple with the TimeExceedance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeExceedance

`func (o *Violation) SetTimeExceedance(v int32)`

SetTimeExceedance sets TimeExceedance field to given value.

### HasTimeExceedance

`func (o *Violation) HasTimeExceedance() bool`

HasTimeExceedance returns a boolean if a field has been set.

### GetDistanceExceedance

`func (o *Violation) GetDistanceExceedance() int32`

GetDistanceExceedance returns the DistanceExceedance field if non-nil, zero value otherwise.

### GetDistanceExceedanceOk

`func (o *Violation) GetDistanceExceedanceOk() (*int32, bool)`

GetDistanceExceedanceOk returns a tuple with the DistanceExceedance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceExceedance

`func (o *Violation) SetDistanceExceedance(v int32)`

SetDistanceExceedance sets DistanceExceedance field to given value.

### HasDistanceExceedance

`func (o *Violation) HasDistanceExceedance() bool`

HasDistanceExceedance returns a boolean if a field has been set.

### GetNumberOfStopsExceedance

`func (o *Violation) GetNumberOfStopsExceedance() int32`

GetNumberOfStopsExceedance returns the NumberOfStopsExceedance field if non-nil, zero value otherwise.

### GetNumberOfStopsExceedanceOk

`func (o *Violation) GetNumberOfStopsExceedanceOk() (*int32, bool)`

GetNumberOfStopsExceedanceOk returns a tuple with the NumberOfStopsExceedance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfStopsExceedance

`func (o *Violation) SetNumberOfStopsExceedance(v int32)`

SetNumberOfStopsExceedance sets NumberOfStopsExceedance field to given value.

### HasNumberOfStopsExceedance

`func (o *Violation) HasNumberOfStopsExceedance() bool`

HasNumberOfStopsExceedance returns a boolean if a field has been set.

### GetCapacityExceedance

`func (o *Violation) GetCapacityExceedance() []int32`

GetCapacityExceedance returns the CapacityExceedance field if non-nil, zero value otherwise.

### GetCapacityExceedanceOk

`func (o *Violation) GetCapacityExceedanceOk() (*[]int32, bool)`

GetCapacityExceedanceOk returns a tuple with the CapacityExceedance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityExceedance

`func (o *Violation) SetCapacityExceedance(v []int32)`

SetCapacityExceedance sets CapacityExceedance field to given value.

### HasCapacityExceedance

`func (o *Violation) HasCapacityExceedance() bool`

HasCapacityExceedance returns a boolean if a field has been set.

### GetMissingEquipment

`func (o *Violation) GetMissingEquipment() []string`

GetMissingEquipment returns the MissingEquipment field if non-nil, zero value otherwise.

### GetMissingEquipmentOk

`func (o *Violation) GetMissingEquipmentOk() (*[]string, bool)`

GetMissingEquipmentOk returns a tuple with the MissingEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingEquipment

`func (o *Violation) SetMissingEquipment(v []string)`

SetMissingEquipment sets MissingEquipment field to given value.

### HasMissingEquipment

`func (o *Violation) HasMissingEquipment() bool`

HasMissingEquipment returns a boolean if a field has been set.

### GetMixedLoadingProhibitions

`func (o *Violation) GetMixedLoadingProhibitions() []MixedLoadingProhibition`

GetMixedLoadingProhibitions returns the MixedLoadingProhibitions field if non-nil, zero value otherwise.

### GetMixedLoadingProhibitionsOk

`func (o *Violation) GetMixedLoadingProhibitionsOk() (*[]MixedLoadingProhibition, bool)`

GetMixedLoadingProhibitionsOk returns a tuple with the MixedLoadingProhibitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixedLoadingProhibitions

`func (o *Violation) SetMixedLoadingProhibitions(v []MixedLoadingProhibition)`

SetMixedLoadingProhibitions sets MixedLoadingProhibitions field to given value.

### HasMixedLoadingProhibitions

`func (o *Violation) HasMixedLoadingProhibitions() bool`

HasMixedLoadingProhibitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


