# CausingError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable message that describes the error. | 
**ErrorCode** | **string** | A constant string that can be used to identify this error class programmatically. An errorCode can have **details** to provide information in additional properties which are described with the code they apply to. They are of type string unless otherwise specified. Note that additional errorCodes as well as the **details** of existing errorCodes may be added at any time. Furthermore, the **description** may change at any time.  **Error codes for** &#x60;GENERAL_VALIDATION_ERROR&#x60; * &#x60;GENERAL_INVALID_VALUE&#x60; - A parameter is set to an invalid value.   * &#x60;value&#x60; - The invalid value. * &#x60;GENERAL_UNRECOGNIZED_PARAMETER&#x60; - A parameter is unknown. * &#x60;GENERAL_MISSING_PARAMETER&#x60; - A required parameter is missing. * &#x60;GENERAL_INVALID_INTERVAL&#x60; - A time interval is invalid, i.e. start is greater than end. * &#x60;GENERAL_MINIMUM_LENGTH_VIOLATED&#x60; - The minimum length is violated.   * &#x60;minimumLength&#x60; - The minimum length (integer). * &#x60;GENERAL_MAXIMUM_LENGTH_VIOLATED&#x60; - The maximum length is violated.   * &#x60;maximumLength&#x60; - The maximum length (integer). * &#x60;GENERAL_DUPLICATE_ID&#x60; - Two or more objects of the same type have the same ID.   * &#x60;value&#x60; - The duplicated value.   * &#x60;indexes&#x60; - The list indexes of the objects with the same ID. * &#x60;GENERAL_MINIMUM_VALUE_VIOLATED&#x60; - The minimum value restriction is violated.   * &#x60;minimumValue&#x60; - The minimum value (integer or double). * &#x60;GENERAL_MAXIMUM_VALUE_VIOLATED&#x60; - The maximum value restriction is violated.   * &#x60;maximumValue&#x60; - The maximum value (integer or double). * &#x60;GENERAL_DUPLICATE_PARAMETER&#x60; - A parameter is duplicated. * &#x60;ROUTEOPTIMIZATION_EMPTY_ID&#x60; - The ID is empty or contains only whitespace. * &#x60;ROUTEOPTIMIZATION_LOCATION_CANNOT_BE_MATCHED&#x60; - The location cannot be matched to a street and no airline distance fallback is possible. * &#x60;ROUTEOPTIMIZATION_DUPLICATE_TRIP_ID&#x60; - Two or more trips have the same trip ID.   * &#x60;duplicatedTripId&#x60; - The duplicated trip ID.   * &#x60;routeIndexes&#x60; - The list indexes of the routes with the same trip ID. * &#x60;ROUTEOPTIMIZATION_DUPLICATE_TRANSPORT&#x60; - The transport with ID &#39;&#39; is picked up and/or delivered multiple times.   * &#x60;transportId&#x60; - The duplicated transport ID. * &#x60;ROUTEOPTIMIZATION_SINGLE_TRIP_PER_ROUTE_VIOLATED&#x60; - There is more than one trip in the given route although the restriction &#39;singleTripPerRoute&#39; is activated. * &#x60;ROUTEOPTIMIZATION_SINGLE_DEPOT_PER_ROUTE_VIOLATED&#x60; - There is more than one depot in the given route although the restriction &#39;singleDepotPerRoute&#39; is activated. * &#x60;ROUTEOPTIMIZATION_ILLEGAL_EMPTY_STOP&#x60; - Empty stops without any pickup or delivery are only allowed for the vehicle start/end location or for trip starts/ends at depots. * &#x60;ROUTEOPTIMIZATION_DELIVERY_EXECUTED_BEFORE_PICKUP&#x60; - Delivery of transport with ID &#39;&#39; is executed before the corresponding pickup.   * &#x60;transportId&#x60; - The transport ID. * &#x60;ROUTEOPTIMIZATION_NO_TRANSPORT_IN_ROUTE&#x60; - There is no transport in this route. Please add a transport or remove the whole route. * &#x60;ROUTEOPTIMIZATION_DELIVERY_MISSING_IN_ROUTE&#x60; - Delivery of transport with ID &#39;&#39; is missing in one of the trips in the given route. Pickup and delivery of a transport need to be in the same trip.   * &#x60;transportId&#x60; - The transport ID. * &#x60;ROUTEOPTIMIZATION_PICKUP_MISSING_IN_ROUTE&#x60; - Pickup of transport with ID &#39;&#39; is missing in one of the trips in the given route. Pickup and delivery of a transport need to be in the same trip.   * &#x60;transportId&#x60; - The transport ID. * &#x60;ROUTEOPTIMIZATION_PICKUP_EQUAL_TO_DELIVERY_LOCATION&#x60; - Pickup and delivery of transport with ID &#39;&#39; are at the same location. Pickup and delivery of a transport need to be at different locations.   * &#x60;transportId&#x60; - The transport ID.   * &#x60;locationId&#x60; - The location ID of the transport pickup and delivery. * &#x60;ROUTEOPTIMIZATION_TRIP_STRUCTURE_VIOLATED&#x60; - Trip structure violated. * &#x60;ROUTEOPTIMIZATION_INVALID_OR_INCONSISTENT_TRIP_ID&#x60; - Trip ID is invalid or not all stops of a trip have the same ID. * &#x60;ROUTEOPTIMIZATION_NO_TIME_INTERVAL&#x60; - The planning horizon is required if there is no other time interval given in the plan. * &#x60;ROUTEOPTIMIZATION_CONTIGUOUS_DRIVER_AVAILABILITIES&#x60; - The driver availability given in the parameter field must end at least two seconds before the driver availabilty with list index  starts or the contiguous intervals have to be merged.   * &#x60;contiguousAvailabilityIndex&#x60; - The list index of the driver availability where the start must have a gap of more than one second to the end of the availability given in the parameter field. * &#x60;ROUTEOPTIMIZATION_PICKUP_MISSING_IN_LOCATIONS&#x60; - The locations list does not contain the pickup location with ID &#39;&#39;.   * &#x60;locationId&#x60; - The location ID. * &#x60;ROUTEOPTIMIZATION_DELIVERY_MISSING_IN_LOCATIONS&#x60; - The locations list does not contain the delivery location with ID &#39;&#39;.   * &#x60;locationId&#x60; - The location ID. * &#x60;ROUTEOPTIMIZATION_UNKNOWN_VEHICLE_ID&#x60; - A vehicle with ID &#39;&#39; does not exist in the vehicles list.   * &#x60;vehicleId&#x60; - The vehicle ID. * &#x60;ROUTEOPTIMIZATION_VEHICLE_REFERENCED_BY_MULTIPLE_DRIVERS&#x60; - A vehicle can only be referenced by one driver. Vehicle with ID &#39;&#39; is referenced by multiple drivers.   * &#x60;vehicleId&#x60; - The vehicle ID.   * &#x60;driverIndexes&#x60; - The indexes of the drivers with the same vehicle ID. * &#x60;ROUTEOPTIMIZATION_VEHICLE_REFERENCED_BY_MULTIPLE_ROUTES&#x60; - A vehicle can only be referenced by one route. Vehicle with ID &#39;&#39; is referenced by multiple routes.   * &#x60;vehicleId&#x60; - The vehicle ID.   * &#x60;routeIndexes&#x60; - The indexes of the routes with the same vehicle ID. * &#x60;ROUTEOPTIMIZATION_VEHICLE_START_LOCATION_MISSING_IN_ROUTE&#x60; - Vehicle start location must be the first stop of the vehicle route.   * &#x60;routeIndex&#x60; - The list index of the route with missing start location.   * &#x60;expectedLocationIdOfFirstStop&#x60; - The expected location ID of the first stop of the route which is the given vehicle start location.   * &#x60;actualLocationIdOfFirstStop&#x60; - The actual location ID of the first stop of the route. * &#x60;ROUTEOPTIMIZATION_VEHICLE_END_LOCATION_MISSING_IN_ROUTE&#x60; - Vehicle end location must be the last stop of the vehicle route.   * &#x60;routeIndex&#x60; - The list index of the route with missing end location.   * &#x60;expectedLocationIdOfLastStop&#x60; - The expected location ID of the last stop of the route which is the given vehicle end location.   * &#x60;actualLocationIdOfLastStop&#x60; - The actual location ID of the last stop of the route. * &#x60;ROUTEOPTIMIZATION_INCONSISTENT_NUMBER_OF_QUANTITIES_AND_CAPACITIES&#x60; - The list of the capacities of all vehicles and the list of the quantities of all transports must have the same length. The reference list is given in the parameter field.   * &#x60;expectedLength&#x60; - The expected length of the list.   * &#x60;vehicleIndexes&#x60; - The list indexes of the vehicles with the wrong number of capacities.   * &#x60;transportIndexes&#x60; - The list indexes of the transports with the wrong number of quantities. * &#x60;ROUTEOPTIMIZATION_UNREFERENCED_LOCATIONS&#x60; - Locations with IDs &#39;&#39; are not referenced or used anywhere. These and all other unreferenced locations must be left out.   * &#x60;locationIds&#x60; - The IDs of some unreferenced locations. * &#x60;ROUTEOPTIMIZATION_UNKNOWN_LOCATION_ID&#x60; - A location with ID &#39;&#39; does not exist in the locations list.   * &#x60;locationId&#x60; - The location ID. * &#x60;ROUTEOPTIMIZATION_UNKNOWN_TRANSPORT_ID&#x60; - A transport with ID &#39;&#39; does not exist in the transports list.   * &#x60;transportId&#x60; - The transport ID. * &#x60;ROUTEOPTIMIZATION_LOCATIONS_TOO_FAR_AWAY&#x60; - Locations or their road access coordinates are too far away from each other when using vehicle profile &#39;&#39;. All locations and road access coordinates have to be inside a rectangle with edges of at most &#39;&#39; km length.   * &#x60;profile&#x60; - The profile for which the locations are too far away from each other.  * &#x60;distance&#x60; - The maximum allowed distance in km. * &#x60;ROUTEOPTIMIZATION_PICKUP_AT_UNEXPECTED_LOCATION&#x60; - Transport with ID &#39;&#39; should not be picked up at this location.   * &#x60;transportId&#x60; - The ID of the transport which is picked up at an unexpected location.   * &#x60;expectedPickupLocationId&#x60; - The location ID where transport should be picked up.   * &#x60;actualPickupLocationId&#x60; - The location ID of the stop where transport is actually picked up. * &#x60;ROUTEOPTIMIZATION_DELIVERY_AT_UNEXPECTED_LOCATION&#x60; - Transport with ID &#39;&#39; should not be delivered at this location.   * &#x60;transportId&#x60; - The ID of the transport which is delivered at an unexpected location.   * &#x60;expectedDeliveryLocationId&#x60; - The location ID where transport should be delivered.   * &#x60;actualDeliveryLocationId&#x60; - The location ID of the stop where transport is actually delivered. * &#x60;ROUTEOPTIMIZATION_ILLEGAL_SPLITTED_STOP&#x60; - Subsequent stops at the same location are only allowed if they are in different trips. The given stops must be merged to one stop.   * &#x60;stopIndexes&#x60; - The list indexes of the stops which should be merged. * &#x60;ROUTEOPTIMIZATION_INVALID_LOADING_SEQUENCE&#x60; - A location in a route is left with unfinished tasks: Either a location is left although there is still load FOR it on the vehicle or a location is visited again although there are still pickups FROM it on the vehicle.   * &#x60;locationId&#x60; - The location ID with unfinished tasks. * &#x60;ROUTEOPTIMIZATION_STOP_SERVICE_TIME_TOO_LONG&#x60; - The service time of the stop exceeds the travel time between daily rests or the working time between breaks defined by the rules for the driver. * &#x60;ROUTEOPTIMIZATION_PLANNING_HORIZON_TOO_LONG&#x60; - Planning horizon must not be longer than 14 days. * &#x60;ROUTEOPTIMIZATION_PLAN_IN_OPTIMIZATION&#x60; - Plan with ID &#39;&#39; is already in optimization. Please wait or terminate current optimization process.   * &#x60;planId&#x60; - The plan ID. * &#x60;ROUTEOPTIMIZATION_EMPTY_VALUE&#x60; - A parameter is empty or contains only whitespace.   **Error codes for** &#x60;ROUTEOPTIMIZATION_RESTRICTION_EXCEEDED&#x60; * &#x60;ROUTEOPTIMIZATION_NUMBER_OF_TRANSPORTS_EXCEEDED&#x60; - The request contains too many transports.   * &#x60;transports&#x60; - The number of transports.   * &#x60;limit&#x60; - The maximum number of transports. * &#x60;ROUTEOPTIMIZATION_NUMBER_OF_LOAD_CATEGORIES_EXCEEDED&#x60; - The request contains too many load categories.   * &#x60;load categories&#x60; - The number of load categories   * &#x60;limit&#x60; - The maximum number of load categories.  **Error codes for** &#x60;GENERAL_RESOURCE_NOT_FOUND&#x60; * &#x60;GENERAL_INVALID_ID&#x60; - The ID does not exist.   * &#x60;value&#x60; - The invalid ID. | 
**Parameter** | Pointer to **string** | The name of the affected query or path parameter or a JSONPath to the affected property of the request. | [optional] 
**Details** | Pointer to **map[string]interface{}** | Additional properties specific to this error class. | [optional] 

## Methods

### NewCausingError

`func NewCausingError(description string, errorCode string, ) *CausingError`

NewCausingError instantiates a new CausingError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCausingErrorWithDefaults

`func NewCausingErrorWithDefaults() *CausingError`

NewCausingErrorWithDefaults instantiates a new CausingError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CausingError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CausingError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CausingError) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetErrorCode

`func (o *CausingError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CausingError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CausingError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetParameter

`func (o *CausingError) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *CausingError) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *CausingError) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *CausingError) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetDetails

`func (o *CausingError) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CausingError) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CausingError) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CausingError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


