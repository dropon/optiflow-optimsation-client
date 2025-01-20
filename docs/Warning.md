# Warning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable message that describes the warning. | 
**WarningCode** | **string** | A constant string that can be used to identify this warning class programmatically. A warningCode can have **details** to provide information in additional properties which are described with the code they apply to. They are of type string unless otherwise specified. Note that additional warningCodes as well as the **details** of existing warningCodes may be added at any time. Furthermore, the **description** may change at any time. * &#x60;ROUTE_OPTIMIZATION_NO_VEHICLE_FOR_ORDER&#x60; - An order cannot be scheduled on any route of the provided vehicles.   * &#x60;orderId&#x60; - The unique identifier of the order. * &#x60;ROUTE_OPTIMIZATION_NO_ORDER_FOR_VEHICLE&#x60; - No order can be scheduled on the route of a vehicle.   * &#x60;vehicleId&#x60; - The unique identifier of the vehicle. * &#x60;ROUTE_OPTIMIZATION_NO_DEPOTS_FOR_VEHICLE&#x60; - No depot can be scheduled on the route of a vehicle.   * &#x60;vehicleId&#x60; - The ID of the vehicle. * &#x60;ROUTE_OPTIMIZATION_LOCATION_UNREACHABLE&#x60; - Travel times and distances cannot be calculated to or from the provided location.   * &#x60;locationId&#x60; - The unique identifier of the location. * &#x60;ROUTE_OPTIMIZATION_DEPOT_INVALID_FOR_BREAKS&#x60; - The total preparation duration of the depot exceeds the maximum work time between breaks for a vehicle.   * &#x60;depotId&#x60; - The unique identifier of the depot.   * &#x60;vehicleId&#x60; - The unique identifier of the vehicle. * &#x60;ROUTE_OPTIMIZATION_TRAFFIC_DATA_MISSING&#x60; - Traffic data is missing for routes to and from some locations. Optimistic travel times will be used instead.   * &#x60;locationIds&#x60; - The unique identifiers of the locations to and from which traffic data is missing.   * &#x60;profile&#x60; - The vehicle routing profile for which traffic data is missing.   * &#x60;trafficMode&#x60; - The vehicle traffic mode for which traffic data is missing. | 
**Details** | Pointer to **map[string]interface{}** | Additional properties specific to this class of warnings. | [optional] 

## Methods

### NewWarning

`func NewWarning(description string, warningCode string, ) *Warning`

NewWarning instantiates a new Warning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningWithDefaults

`func NewWarningWithDefaults() *Warning`

NewWarningWithDefaults instantiates a new Warning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Warning) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Warning) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Warning) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetWarningCode

`func (o *Warning) GetWarningCode() string`

GetWarningCode returns the WarningCode field if non-nil, zero value otherwise.

### GetWarningCodeOk

`func (o *Warning) GetWarningCodeOk() (*string, bool)`

GetWarningCodeOk returns a tuple with the WarningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningCode

`func (o *Warning) SetWarningCode(v string)`

SetWarningCode sets WarningCode field to given value.


### GetDetails

`func (o *Warning) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Warning) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Warning) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Warning) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


