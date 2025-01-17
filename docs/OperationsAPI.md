# \OperationsAPI

All URIs are relative to *https://api.myptv.com/routeoptimization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOperation**](OperationsAPI.md#CancelOperation) | **Delete** /plans/{id}/operation | 
[**GetOperationStatus**](OperationsAPI.md#GetOperationStatus) | **Get** /plans/{id}/operation | 
[**StartEvaluation**](OperationsAPI.md#StartEvaluation) | **Post** /plans/{id}/operation/evaluation | 
[**StartOptimization**](OperationsAPI.md#StartOptimization) | **Post** /plans/{id}/operation/optimization | 



## CancelOperation

> CancelOperation(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the plan which corresponding operation should be cancelled.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OperationsAPI.CancelOperation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.CancelOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the plan which corresponding operation should be cancelled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperationStatus

> Operation GetOperationStatus(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the plan for which the corresponding operation status should be requested.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.GetOperationStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.GetOperationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperationStatus`: Operation
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.GetOperationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the plan for which the corresponding operation status should be requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Operation**](Operation.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartEvaluation

> StartEvaluation(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the plan to be evaluated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OperationsAPI.StartEvaluation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.StartEvaluation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the plan to be evaluated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOptimization

> StartOptimization(ctx, id).Quality(quality).TweaksToObjective(tweaksToObjective).ConsiderTransportPriorities(considerTransportPriorities).OptimizationPremium(optimizationPremium).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the plan to be optimized.
	quality := openapiclient.OptimizationQuality("STANDARD") // OptimizationQuality | The optimization quality defines the tradeoff between calculation time and quality. A high solution quality and a low calculation time are conflicting. A larger solution search space may lead to a better solution but will take more calculation time.   * `STANDARD` - This quality level represents a good trade-off between solution quality and calculation time.   * `HIGH` - On this quality level, the search space is larger than on the standard level. This may lead to a better solution but will take more calculation time. (optional) (default to "STANDARD")
	tweaksToObjective := []openapiclient.TweakToObjective{openapiclient.TweakToObjective("AVOID_INTERSECTIONS")} // []TweakToObjective | A tweak to objective defines additional optimization criteria. Without any tweak, the standard optimization goal is to plan as many transports as possible with as few vehicles as possible.   * `AVOID_INTERSECTIONS` - The focus of this approach is to avoid intersections of edges, where an edge is a straight line drawn between two consecutive stops in a route. The number of routes may be higher compared to the standard optimization goal.   * `IGNORE_MINIMIZATION_OF_NUMBER_OF_ROUTES` - With this tweak, the number of vehicles used in a plan, and thus the number of routes, is not minimized. (optional)
	considerTransportPriorities := true // bool | Set transport priority consideration to true or false.  See [here](./concepts/transport-priorities) for more information. (optional) (default to false)
	optimizationPremium := true // bool | Set Optimization Premium to true to boost quality of results for all **quality** modes by taking advantage of additional computational resources.  Transactions' packages for Optimization Premium are calculated differently.  See [here](./concepts/optimization-premium) for more information. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OperationsAPI.StartOptimization(context.Background(), id).Quality(quality).TweaksToObjective(tweaksToObjective).ConsiderTransportPriorities(considerTransportPriorities).OptimizationPremium(optimizationPremium).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.StartOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the plan to be optimized. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quality** | [**OptimizationQuality**](OptimizationQuality.md) | The optimization quality defines the tradeoff between calculation time and quality. A high solution quality and a low calculation time are conflicting. A larger solution search space may lead to a better solution but will take more calculation time.   * &#x60;STANDARD&#x60; - This quality level represents a good trade-off between solution quality and calculation time.   * &#x60;HIGH&#x60; - On this quality level, the search space is larger than on the standard level. This may lead to a better solution but will take more calculation time. | [default to &quot;STANDARD&quot;]
 **tweaksToObjective** | [**[]TweakToObjective**](TweakToObjective.md) | A tweak to objective defines additional optimization criteria. Without any tweak, the standard optimization goal is to plan as many transports as possible with as few vehicles as possible.   * &#x60;AVOID_INTERSECTIONS&#x60; - The focus of this approach is to avoid intersections of edges, where an edge is a straight line drawn between two consecutive stops in a route. The number of routes may be higher compared to the standard optimization goal.   * &#x60;IGNORE_MINIMIZATION_OF_NUMBER_OF_ROUTES&#x60; - With this tweak, the number of vehicles used in a plan, and thus the number of routes, is not minimized. | 
 **considerTransportPriorities** | **bool** | Set transport priority consideration to true or false.  See [here](./concepts/transport-priorities) for more information. | [default to false]
 **optimizationPremium** | **bool** | Set Optimization Premium to true to boost quality of results for all **quality** modes by taking advantage of additional computational resources.  Transactions&#39; packages for Optimization Premium are calculated differently.  See [here](./concepts/optimization-premium) for more information. | [default to false]

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

