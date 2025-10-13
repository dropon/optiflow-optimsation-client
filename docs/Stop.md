# Stop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** | The unique identifier of the location where this stop is taking place. | [optional] 
**Approach** | Pointer to [**Leg**](Leg.md) |  | [optional] 
**Arrival** | Pointer to **time.Time** | The point in time when the vehicle arrives at the location. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**PreparationDuration** | Pointer to **int32** | A period of time [s] that is required between the arrival at the location and the start of the first appointment. The preparation duration is scheduled before tasks can be executed at the stop. | [optional] 
**Appointments** | Pointer to [**[]Appointment**](Appointment.md) | A list of appointments that describe the tasks that are scheduled for execution at this stop. Consecutive tasks are grouped into an appointment if they have been assigned to the same time slot. | [optional] 
**Charging** | Pointer to [**[]Charging**](Charging.md) | Indicates when the vehicle battery is recharged at the current stop. Charging can take place at the very beginning or at the very end of the stop. | [optional] 
**VehicleSlotIndex** | Pointer to **int32** | Describes the index of the vehicle slot this stop is assigned to. If there is no vehicle slot provided for this index, the stop is assigned to an extra vehicle slot. | [optional] 
**Departure** | Pointer to **time.Time** | The point in time when the vehicle departs at the location. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 

## Methods

### NewStop

`func NewStop() *Stop`

NewStop instantiates a new Stop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopWithDefaults

`func NewStopWithDefaults() *Stop`

NewStopWithDefaults instantiates a new Stop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *Stop) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Stop) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Stop) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Stop) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetApproach

`func (o *Stop) GetApproach() Leg`

GetApproach returns the Approach field if non-nil, zero value otherwise.

### GetApproachOk

`func (o *Stop) GetApproachOk() (*Leg, bool)`

GetApproachOk returns a tuple with the Approach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproach

`func (o *Stop) SetApproach(v Leg)`

SetApproach sets Approach field to given value.

### HasApproach

`func (o *Stop) HasApproach() bool`

HasApproach returns a boolean if a field has been set.

### GetArrival

`func (o *Stop) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *Stop) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *Stop) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.

### HasArrival

`func (o *Stop) HasArrival() bool`

HasArrival returns a boolean if a field has been set.

### GetPreparationDuration

`func (o *Stop) GetPreparationDuration() int32`

GetPreparationDuration returns the PreparationDuration field if non-nil, zero value otherwise.

### GetPreparationDurationOk

`func (o *Stop) GetPreparationDurationOk() (*int32, bool)`

GetPreparationDurationOk returns a tuple with the PreparationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationDuration

`func (o *Stop) SetPreparationDuration(v int32)`

SetPreparationDuration sets PreparationDuration field to given value.

### HasPreparationDuration

`func (o *Stop) HasPreparationDuration() bool`

HasPreparationDuration returns a boolean if a field has been set.

### GetAppointments

`func (o *Stop) GetAppointments() []Appointment`

GetAppointments returns the Appointments field if non-nil, zero value otherwise.

### GetAppointmentsOk

`func (o *Stop) GetAppointmentsOk() (*[]Appointment, bool)`

GetAppointmentsOk returns a tuple with the Appointments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppointments

`func (o *Stop) SetAppointments(v []Appointment)`

SetAppointments sets Appointments field to given value.

### HasAppointments

`func (o *Stop) HasAppointments() bool`

HasAppointments returns a boolean if a field has been set.

### GetCharging

`func (o *Stop) GetCharging() []Charging`

GetCharging returns the Charging field if non-nil, zero value otherwise.

### GetChargingOk

`func (o *Stop) GetChargingOk() (*[]Charging, bool)`

GetChargingOk returns a tuple with the Charging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharging

`func (o *Stop) SetCharging(v []Charging)`

SetCharging sets Charging field to given value.

### HasCharging

`func (o *Stop) HasCharging() bool`

HasCharging returns a boolean if a field has been set.

### GetVehicleSlotIndex

`func (o *Stop) GetVehicleSlotIndex() int32`

GetVehicleSlotIndex returns the VehicleSlotIndex field if non-nil, zero value otherwise.

### GetVehicleSlotIndexOk

`func (o *Stop) GetVehicleSlotIndexOk() (*int32, bool)`

GetVehicleSlotIndexOk returns a tuple with the VehicleSlotIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleSlotIndex

`func (o *Stop) SetVehicleSlotIndex(v int32)`

SetVehicleSlotIndex sets VehicleSlotIndex field to given value.

### HasVehicleSlotIndex

`func (o *Stop) HasVehicleSlotIndex() bool`

HasVehicleSlotIndex returns a boolean if a field has been set.

### GetDeparture

`func (o *Stop) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *Stop) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *Stop) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.

### HasDeparture

`func (o *Stop) HasDeparture() bool`

HasDeparture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


