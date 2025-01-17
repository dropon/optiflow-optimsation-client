# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EventType**](EventType.md) |  | 
**StartTime** | **time.Time** | The start time of the event formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). | 
**Duration** | **int32** | The duration of the event [s]. The duration may be 0. | 
**TransportId** | Pointer to **string** | The ID of the corresponding transport if the event is a service event. Otherwise the ID is null. | [optional] 

## Methods

### NewEvent

`func NewEvent(type_ EventType, startTime time.Time, duration int32, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Event) GetType() EventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*EventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v EventType)`

SetType sets Type field to given value.


### GetStartTime

`func (o *Event) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Event) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Event) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetDuration

`func (o *Event) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Event) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Event) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTransportId

`func (o *Event) GetTransportId() string`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *Event) GetTransportIdOk() (*string, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *Event) SetTransportId(v string)`

SetTransportId sets TransportId field to given value.

### HasTransportId

`func (o *Event) HasTransportId() bool`

HasTransportId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


