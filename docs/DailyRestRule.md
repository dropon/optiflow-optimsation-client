# DailyRestRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyRestTime** | **int32** | The minimum duration of a daily rest [s]. Minimum is 3 hours, maximum is 22 hours. | 
**MaximumDrivingTimeBetweenDailyRests** | Pointer to **NullableInt32** | Maximum duration that the driver is allowed to drive [s] before taking a daily rest. The maximum driving time is considered as infinite if it is not set. | [optional] 
**MaximumTravelTimeBetweenDailyRests** | Pointer to **NullableInt32** | Maximum duration that the driver is allowed to travel [s] before taking a daily. The maximum travel time is considered as infinite if it is not set. | [optional] 
**DailyRestPosition** | Pointer to [**DailyRestPosition**](DailyRestPosition.md) |  | [optional] [default to BETWEEN_TRIPS]

## Methods

### NewDailyRestRule

`func NewDailyRestRule(dailyRestTime int32, ) *DailyRestRule`

NewDailyRestRule instantiates a new DailyRestRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyRestRuleWithDefaults

`func NewDailyRestRuleWithDefaults() *DailyRestRule`

NewDailyRestRuleWithDefaults instantiates a new DailyRestRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyRestTime

`func (o *DailyRestRule) GetDailyRestTime() int32`

GetDailyRestTime returns the DailyRestTime field if non-nil, zero value otherwise.

### GetDailyRestTimeOk

`func (o *DailyRestRule) GetDailyRestTimeOk() (*int32, bool)`

GetDailyRestTimeOk returns a tuple with the DailyRestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyRestTime

`func (o *DailyRestRule) SetDailyRestTime(v int32)`

SetDailyRestTime sets DailyRestTime field to given value.


### GetMaximumDrivingTimeBetweenDailyRests

`func (o *DailyRestRule) GetMaximumDrivingTimeBetweenDailyRests() int32`

GetMaximumDrivingTimeBetweenDailyRests returns the MaximumDrivingTimeBetweenDailyRests field if non-nil, zero value otherwise.

### GetMaximumDrivingTimeBetweenDailyRestsOk

`func (o *DailyRestRule) GetMaximumDrivingTimeBetweenDailyRestsOk() (*int32, bool)`

GetMaximumDrivingTimeBetweenDailyRestsOk returns a tuple with the MaximumDrivingTimeBetweenDailyRests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDrivingTimeBetweenDailyRests

`func (o *DailyRestRule) SetMaximumDrivingTimeBetweenDailyRests(v int32)`

SetMaximumDrivingTimeBetweenDailyRests sets MaximumDrivingTimeBetweenDailyRests field to given value.

### HasMaximumDrivingTimeBetweenDailyRests

`func (o *DailyRestRule) HasMaximumDrivingTimeBetweenDailyRests() bool`

HasMaximumDrivingTimeBetweenDailyRests returns a boolean if a field has been set.

### SetMaximumDrivingTimeBetweenDailyRestsNil

`func (o *DailyRestRule) SetMaximumDrivingTimeBetweenDailyRestsNil(b bool)`

 SetMaximumDrivingTimeBetweenDailyRestsNil sets the value for MaximumDrivingTimeBetweenDailyRests to be an explicit nil

### UnsetMaximumDrivingTimeBetweenDailyRests
`func (o *DailyRestRule) UnsetMaximumDrivingTimeBetweenDailyRests()`

UnsetMaximumDrivingTimeBetweenDailyRests ensures that no value is present for MaximumDrivingTimeBetweenDailyRests, not even an explicit nil
### GetMaximumTravelTimeBetweenDailyRests

`func (o *DailyRestRule) GetMaximumTravelTimeBetweenDailyRests() int32`

GetMaximumTravelTimeBetweenDailyRests returns the MaximumTravelTimeBetweenDailyRests field if non-nil, zero value otherwise.

### GetMaximumTravelTimeBetweenDailyRestsOk

`func (o *DailyRestRule) GetMaximumTravelTimeBetweenDailyRestsOk() (*int32, bool)`

GetMaximumTravelTimeBetweenDailyRestsOk returns a tuple with the MaximumTravelTimeBetweenDailyRests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumTravelTimeBetweenDailyRests

`func (o *DailyRestRule) SetMaximumTravelTimeBetweenDailyRests(v int32)`

SetMaximumTravelTimeBetweenDailyRests sets MaximumTravelTimeBetweenDailyRests field to given value.

### HasMaximumTravelTimeBetweenDailyRests

`func (o *DailyRestRule) HasMaximumTravelTimeBetweenDailyRests() bool`

HasMaximumTravelTimeBetweenDailyRests returns a boolean if a field has been set.

### SetMaximumTravelTimeBetweenDailyRestsNil

`func (o *DailyRestRule) SetMaximumTravelTimeBetweenDailyRestsNil(b bool)`

 SetMaximumTravelTimeBetweenDailyRestsNil sets the value for MaximumTravelTimeBetweenDailyRests to be an explicit nil

### UnsetMaximumTravelTimeBetweenDailyRests
`func (o *DailyRestRule) UnsetMaximumTravelTimeBetweenDailyRests()`

UnsetMaximumTravelTimeBetweenDailyRests ensures that no value is present for MaximumTravelTimeBetweenDailyRests, not even an explicit nil
### GetDailyRestPosition

`func (o *DailyRestRule) GetDailyRestPosition() DailyRestPosition`

GetDailyRestPosition returns the DailyRestPosition field if non-nil, zero value otherwise.

### GetDailyRestPositionOk

`func (o *DailyRestRule) GetDailyRestPositionOk() (*DailyRestPosition, bool)`

GetDailyRestPositionOk returns a tuple with the DailyRestPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyRestPosition

`func (o *DailyRestRule) SetDailyRestPosition(v DailyRestPosition)`

SetDailyRestPosition sets DailyRestPosition field to given value.

### HasDailyRestPosition

`func (o *DailyRestRule) HasDailyRestPosition() bool`

HasDailyRestPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


