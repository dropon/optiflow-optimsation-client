# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier of the location. | 
**Latitude** | **float64** | The latitude value of the location in degrees (WGS84/EPSG:4326) from south to north. | 
**Longitude** | **float64** | The longitude value of the location in degrees (WGS84/EPSG:4326) from west to east. | 
**MatchSideOfStreet** | Pointer to **bool** | Requires the vehicle to approach from the side of the street where the location is situated. This prevents the driver from needing to cross the street. | [optional] [default to false]
**StopProperties** | Pointer to [**StopProperties**](StopProperties.md) |  | [optional] 
**ChargingStations** | Pointer to [**[]ChargingStation**](ChargingStation.md) | Specifies the charging capabilities that can be used at the location for electric vehicles. | [optional] 
**Layover** | Pointer to [**LayoverCondition**](LayoverCondition.md) |  | [optional] [default to DISALLOW]
**Categories** | Pointer to **[]string** | A list of categories the location belongs to that can be used to describe constraints or rules. | [optional] [default to []]

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

### GetStopProperties

`func (o *Location) GetStopProperties() StopProperties`

GetStopProperties returns the StopProperties field if non-nil, zero value otherwise.

### GetStopPropertiesOk

`func (o *Location) GetStopPropertiesOk() (*StopProperties, bool)`

GetStopPropertiesOk returns a tuple with the StopProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProperties

`func (o *Location) SetStopProperties(v StopProperties)`

SetStopProperties sets StopProperties field to given value.

### HasStopProperties

`func (o *Location) HasStopProperties() bool`

HasStopProperties returns a boolean if a field has been set.

### GetChargingStations

`func (o *Location) GetChargingStations() []ChargingStation`

GetChargingStations returns the ChargingStations field if non-nil, zero value otherwise.

### GetChargingStationsOk

`func (o *Location) GetChargingStationsOk() (*[]ChargingStation, bool)`

GetChargingStationsOk returns a tuple with the ChargingStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingStations

`func (o *Location) SetChargingStations(v []ChargingStation)`

SetChargingStations sets ChargingStations field to given value.

### HasChargingStations

`func (o *Location) HasChargingStations() bool`

HasChargingStations returns a boolean if a field has been set.

### GetLayover

`func (o *Location) GetLayover() LayoverCondition`

GetLayover returns the Layover field if non-nil, zero value otherwise.

### GetLayoverOk

`func (o *Location) GetLayoverOk() (*LayoverCondition, bool)`

GetLayoverOk returns a tuple with the Layover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayover

`func (o *Location) SetLayover(v LayoverCondition)`

SetLayover sets Layover field to given value.

### HasLayover

`func (o *Location) HasLayover() bool`

HasLayover returns a boolean if a field has been set.

### GetCategories

`func (o *Location) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Location) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Location) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Location) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


