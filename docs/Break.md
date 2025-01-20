# Break

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** | The point in time when the break starts. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 
**Duration** | Pointer to **int32** | The duration [s] of the break. | [optional] 
**End** | Pointer to **time.Time** | The point in time when the break ends. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | [optional] 

## Methods

### NewBreak

`func NewBreak() *Break`

NewBreak instantiates a new Break object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakWithDefaults

`func NewBreakWithDefaults() *Break`

NewBreakWithDefaults instantiates a new Break object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *Break) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Break) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Break) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Break) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetDuration

`func (o *Break) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Break) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Break) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Break) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnd

`func (o *Break) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Break) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Break) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Break) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


