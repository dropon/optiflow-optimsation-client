# RouteDurationPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Importance** | Pointer to **float64** | A scale between 0 and 1 resembling a tradeoff between minimizing the total hour cost and steering towards the preferred duration of the vehicle. Higher values indicate a stronger preference for a route duration close to the preferred one. As the cost per hour increases, the influence of this tradeoff becomes more significant. | [optional] [default to 0]
**PreferredDuration** | Pointer to **int32** | The preferred duration [s] of the route assigned to the vehicle. If omitted, the preferred route duration of this vehicle is the average route duration. | [optional] 

## Methods

### NewRouteDurationPreference

`func NewRouteDurationPreference() *RouteDurationPreference`

NewRouteDurationPreference instantiates a new RouteDurationPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteDurationPreferenceWithDefaults

`func NewRouteDurationPreferenceWithDefaults() *RouteDurationPreference`

NewRouteDurationPreferenceWithDefaults instantiates a new RouteDurationPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportance

`func (o *RouteDurationPreference) GetImportance() float64`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *RouteDurationPreference) GetImportanceOk() (*float64, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *RouteDurationPreference) SetImportance(v float64)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *RouteDurationPreference) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### GetPreferredDuration

`func (o *RouteDurationPreference) GetPreferredDuration() int32`

GetPreferredDuration returns the PreferredDuration field if non-nil, zero value otherwise.

### GetPreferredDurationOk

`func (o *RouteDurationPreference) GetPreferredDurationOk() (*int32, bool)`

GetPreferredDurationOk returns a tuple with the PreferredDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDuration

`func (o *RouteDurationPreference) SetPreferredDuration(v int32)`

SetPreferredDuration sets PreferredDuration field to given value.

### HasPreferredDuration

`func (o *RouteDurationPreference) HasPreferredDuration() bool`

HasPreferredDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


