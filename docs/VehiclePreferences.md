# VehiclePreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compactness** | Pointer to **float64** | A scale between 0 and 1 resembling a tradeoff between minimizing distance cost and maximizing compactness, where higher values indicate a stronger preference for compact routes. As the cost per kilometer increases, the influence of this tradeoff becomes more significant. A route is considered compact if all stops for executing non-depot tasks are close to each other. | [optional] [default to 0]

## Methods

### NewVehiclePreferences

`func NewVehiclePreferences() *VehiclePreferences`

NewVehiclePreferences instantiates a new VehiclePreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehiclePreferencesWithDefaults

`func NewVehiclePreferencesWithDefaults() *VehiclePreferences`

NewVehiclePreferencesWithDefaults instantiates a new VehiclePreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompactness

`func (o *VehiclePreferences) GetCompactness() float64`

GetCompactness returns the Compactness field if non-nil, zero value otherwise.

### GetCompactnessOk

`func (o *VehiclePreferences) GetCompactnessOk() (*float64, bool)`

GetCompactnessOk returns a tuple with the Compactness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactness

`func (o *VehiclePreferences) SetCompactness(v float64)`

SetCompactness sets Compactness field to given value.

### HasCompactness

`func (o *VehiclePreferences) HasCompactness() bool`

HasCompactness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


