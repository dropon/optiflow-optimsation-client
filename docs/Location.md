# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for this location. This ID can be used to reference the location from other elements, for example from transports or vehicles. The ID does not influence the result. | 
**Type** | Pointer to [**LocationType**](LocationType.md) |  | [optional] [default to CUSTOMER]
**Latitude** | **float64** | The latitude value of the location in degrees (WGS84/EPSG:4326) from south to north. | 
**Longitude** | **float64** | The longitude value of the location in degrees (WGS84/EPSG:4326) from west to east. | 
**RoadAccess** | Pointer to [**RoadAccess**](RoadAccess.md) |  | [optional] 
**IncludeLastMeters** | Pointer to **bool** | Include the air-line connection between given and matched coordinates in the relation distance and travel time. Will be ignored when **roadAccess** is specified. We will refer to this type of location as an _off-road location_. | [optional] [default to true]
**MatchSideOfStreet** | Pointer to **bool** | Specifies that this waypoint will be reached at the side of street on which it is located. This is useful to prevent the driver from crossing the street to actually reach the location. Is disabled if an OSM profile is used. | [optional] [default to false]
**ApplyVehicleDependentServiceTimeFactor** | Pointer to **bool** | Indicates if vehicle-dependent service time factors are relevant for this location. The factors are typically relevant when the vehicles are (un)loaded by the drivers and may be irrelevant when the vehicles are (un)loaded by the location&#39;s staff. If this parameter is set to false, no vehicle-dependent service time factors are taken into account for this location when processing transports. | [optional] [default to true]
**OpeningIntervals** | Pointer to [**[]TimeInterval**](TimeInterval.md) | The opening intervals at this location, each specified by two points in time - the beginning and the end of the interval. Leaving this parameter empty means that the location is always open. Service (pickup or delivery) can only start within one of the opening intervals. If a planning horizon is defined for the request, all opening intervals outside of this horizon are not considered during the route optimization process. If no planning horizon is defined, it is required that the opening intervals of all locations do not span a horizon longer than two weeks. | [optional] 
**DepotLocationAttributes** | Pointer to [**DepotLocationAttributes**](DepotLocationAttributes.md) |  | [optional] 
**CustomerLocationAttributes** | Pointer to [**CustomerLocationAttributes**](CustomerLocationAttributes.md) |  | [optional] 

## Methods

### NewLocation

`func NewLocation(id string, latitude float64, longitude float64, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Location) GetType() LocationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Location) GetTypeOk() (*LocationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Location) SetType(v LocationType)`

SetType sets Type field to given value.

### HasType

`func (o *Location) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLatitude

`func (o *Location) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Location) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Location) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *Location) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Location) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Location) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetRoadAccess

`func (o *Location) GetRoadAccess() RoadAccess`

GetRoadAccess returns the RoadAccess field if non-nil, zero value otherwise.

### GetRoadAccessOk

`func (o *Location) GetRoadAccessOk() (*RoadAccess, bool)`

GetRoadAccessOk returns a tuple with the RoadAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadAccess

`func (o *Location) SetRoadAccess(v RoadAccess)`

SetRoadAccess sets RoadAccess field to given value.

### HasRoadAccess

`func (o *Location) HasRoadAccess() bool`

HasRoadAccess returns a boolean if a field has been set.

### GetIncludeLastMeters

`func (o *Location) GetIncludeLastMeters() bool`

GetIncludeLastMeters returns the IncludeLastMeters field if non-nil, zero value otherwise.

### GetIncludeLastMetersOk

`func (o *Location) GetIncludeLastMetersOk() (*bool, bool)`

GetIncludeLastMetersOk returns a tuple with the IncludeLastMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLastMeters

`func (o *Location) SetIncludeLastMeters(v bool)`

SetIncludeLastMeters sets IncludeLastMeters field to given value.

### HasIncludeLastMeters

`func (o *Location) HasIncludeLastMeters() bool`

HasIncludeLastMeters returns a boolean if a field has been set.

### GetMatchSideOfStreet

`func (o *Location) GetMatchSideOfStreet() bool`

GetMatchSideOfStreet returns the MatchSideOfStreet field if non-nil, zero value otherwise.

### GetMatchSideOfStreetOk

`func (o *Location) GetMatchSideOfStreetOk() (*bool, bool)`

GetMatchSideOfStreetOk returns a tuple with the MatchSideOfStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSideOfStreet

`func (o *Location) SetMatchSideOfStreet(v bool)`

SetMatchSideOfStreet sets MatchSideOfStreet field to given value.

### HasMatchSideOfStreet

`func (o *Location) HasMatchSideOfStreet() bool`

HasMatchSideOfStreet returns a boolean if a field has been set.

### GetApplyVehicleDependentServiceTimeFactor

`func (o *Location) GetApplyVehicleDependentServiceTimeFactor() bool`

GetApplyVehicleDependentServiceTimeFactor returns the ApplyVehicleDependentServiceTimeFactor field if non-nil, zero value otherwise.

### GetApplyVehicleDependentServiceTimeFactorOk

`func (o *Location) GetApplyVehicleDependentServiceTimeFactorOk() (*bool, bool)`

GetApplyVehicleDependentServiceTimeFactorOk returns a tuple with the ApplyVehicleDependentServiceTimeFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyVehicleDependentServiceTimeFactor

`func (o *Location) SetApplyVehicleDependentServiceTimeFactor(v bool)`

SetApplyVehicleDependentServiceTimeFactor sets ApplyVehicleDependentServiceTimeFactor field to given value.

### HasApplyVehicleDependentServiceTimeFactor

`func (o *Location) HasApplyVehicleDependentServiceTimeFactor() bool`

HasApplyVehicleDependentServiceTimeFactor returns a boolean if a field has been set.

### GetOpeningIntervals

`func (o *Location) GetOpeningIntervals() []TimeInterval`

GetOpeningIntervals returns the OpeningIntervals field if non-nil, zero value otherwise.

### GetOpeningIntervalsOk

`func (o *Location) GetOpeningIntervalsOk() (*[]TimeInterval, bool)`

GetOpeningIntervalsOk returns a tuple with the OpeningIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningIntervals

`func (o *Location) SetOpeningIntervals(v []TimeInterval)`

SetOpeningIntervals sets OpeningIntervals field to given value.

### HasOpeningIntervals

`func (o *Location) HasOpeningIntervals() bool`

HasOpeningIntervals returns a boolean if a field has been set.

### GetDepotLocationAttributes

`func (o *Location) GetDepotLocationAttributes() DepotLocationAttributes`

GetDepotLocationAttributes returns the DepotLocationAttributes field if non-nil, zero value otherwise.

### GetDepotLocationAttributesOk

`func (o *Location) GetDepotLocationAttributesOk() (*DepotLocationAttributes, bool)`

GetDepotLocationAttributesOk returns a tuple with the DepotLocationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotLocationAttributes

`func (o *Location) SetDepotLocationAttributes(v DepotLocationAttributes)`

SetDepotLocationAttributes sets DepotLocationAttributes field to given value.

### HasDepotLocationAttributes

`func (o *Location) HasDepotLocationAttributes() bool`

HasDepotLocationAttributes returns a boolean if a field has been set.

### GetCustomerLocationAttributes

`func (o *Location) GetCustomerLocationAttributes() CustomerLocationAttributes`

GetCustomerLocationAttributes returns the CustomerLocationAttributes field if non-nil, zero value otherwise.

### GetCustomerLocationAttributesOk

`func (o *Location) GetCustomerLocationAttributesOk() (*CustomerLocationAttributes, bool)`

GetCustomerLocationAttributesOk returns a tuple with the CustomerLocationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLocationAttributes

`func (o *Location) SetCustomerLocationAttributes(v CustomerLocationAttributes)`

SetCustomerLocationAttributes sets CustomerLocationAttributes field to given value.

### HasCustomerLocationAttributes

`func (o *Location) HasCustomerLocationAttributes() bool`

HasCustomerLocationAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


