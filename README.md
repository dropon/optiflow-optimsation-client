# Go API client for openapi

With the Route Optimization service you can schedule and optimize the routes of your fleet.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.11
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://developer.myptv.com/](https://developer.myptv.com/)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.myptv.com/routeoptimization/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OperationsAPI* | [**CancelOperation**](docs/OperationsAPI.md#canceloperation) | **Delete** /plans/{id}/operation | 
*OperationsAPI* | [**GetOperationStatus**](docs/OperationsAPI.md#getoperationstatus) | **Get** /plans/{id}/operation | 
*OperationsAPI* | [**StartEvaluation**](docs/OperationsAPI.md#startevaluation) | **Post** /plans/{id}/operation/evaluation | 
*OperationsAPI* | [**StartOptimization**](docs/OperationsAPI.md#startoptimization) | **Post** /plans/{id}/operation/optimization | 
*PlansAPI* | [**CreatePlan**](docs/PlansAPI.md#createplan) | **Post** /plans | 
*PlansAPI* | [**DeletePlan**](docs/PlansAPI.md#deleteplan) | **Delete** /plans/{id} | 
*PlansAPI* | [**GetPlan**](docs/PlansAPI.md#getplan) | **Get** /plans/{id} | 
*PlansAPI* | [**GetPlanSummaries**](docs/PlansAPI.md#getplansummaries) | **Get** /plans/summaries | 


## Documentation For Models

 - [BreakRule](docs/BreakRule.md)
 - [CapacitiesChangePosition](docs/CapacitiesChangePosition.md)
 - [CausingError](docs/CausingError.md)
 - [CustomerLocationAttributes](docs/CustomerLocationAttributes.md)
 - [DailyRestPosition](docs/DailyRestPosition.md)
 - [DailyRestRule](docs/DailyRestRule.md)
 - [DepotLocationAttributes](docs/DepotLocationAttributes.md)
 - [Driver](docs/Driver.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [Event](docs/Event.md)
 - [EventType](docs/EventType.md)
 - [Location](docs/Location.md)
 - [LocationType](docs/LocationType.md)
 - [MixedLoadingProhibition](docs/MixedLoadingProhibition.md)
 - [Operation](docs/Operation.md)
 - [OperationStatus](docs/OperationStatus.md)
 - [OptimizationQuality](docs/OptimizationQuality.md)
 - [Plan](docs/Plan.md)
 - [PlanSummaries](docs/PlanSummaries.md)
 - [PlanSummary](docs/PlanSummary.md)
 - [PlanningRestrictions](docs/PlanningRestrictions.md)
 - [PositionInTrip](docs/PositionInTrip.md)
 - [RoadAccess](docs/RoadAccess.md)
 - [Route](docs/Route.md)
 - [RouteReport](docs/RouteReport.md)
 - [Stop](docs/Stop.md)
 - [StopReport](docs/StopReport.md)
 - [TimeInterval](docs/TimeInterval.md)
 - [Transport](docs/Transport.md)
 - [TweakToObjective](docs/TweakToObjective.md)
 - [Vehicle](docs/Vehicle.md)
 - [Violation](docs/Violation.md)
 - [ViolationType](docs/ViolationType.md)
 - [Warning](docs/Warning.md)
 - [WayReport](docs/WayReport.md)
 - [WorkLogbook](docs/WorkLogbook.md)
 - [WorkingHoursPreset](docs/WorkingHoursPreset.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### apiKeyAuth

- **Type**: API key
- **API key parameter name**: apiKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apiKeyAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"apiKeyAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


