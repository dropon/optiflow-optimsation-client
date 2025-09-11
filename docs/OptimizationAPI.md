# \OptimizationAPI

All URIs are relative to *https://api.myptv.com/routeoptimization/optiflow/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptimizationResult**](OptimizationAPI.md#GetOptimizationResult) | **Get** /optimizations/{id} | 
[**StartOptimization**](OptimizationAPI.md#StartOptimization) | **Post** /optimizations | 
[**StopOptimization**](OptimizationAPI.md#StopOptimization) | **Post** /optimizations/{id}/stop | 



## GetOptimizationResult

> OptimizationResult GetOptimizationResult(ctx, id).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A unique identifier of the optimization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptimizationAPI.GetOptimizationResult(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.GetOptimizationResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptimizationResult`: OptimizationResult
	fmt.Fprintf(os.Stdout, "Response from `OptimizationAPI.GetOptimizationResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique identifier of the optimization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptimizationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OptimizationResult**](OptimizationResult.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOptimization

> OptimizationIdentifier StartOptimization(ctx).OptimizationRequest(optimizationRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	optimizationRequest := *openapiclient.NewOptimizationRequest(*openapiclient.NewOptimizationSettings(int32(7200)), []openapiclient.Location{*openapiclient.NewLocation("LOCATION-123", float64(50.95136607213874), float64(3.8068358460359364))}, *openapiclient.NewOrders(), []openapiclient.Vehicle{*openapiclient.NewVehicle("VEHICLE-123", *openapiclient.NewVehicleStart(time.Now()), *openapiclient.NewVehicleEnd(time.Now()), *openapiclient.NewVehicleRouting("EUR_VAN"), *openapiclient.NewVehicleCosts(float64(40), float64(0.5)))}) // OptimizationRequest | The data needed to optimize the routes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptimizationAPI.StartOptimization(context.Background()).OptimizationRequest(optimizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.StartOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOptimization`: OptimizationIdentifier
	fmt.Fprintf(os.Stdout, "Response from `OptimizationAPI.StartOptimization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optimizationRequest** | [**OptimizationRequest**](OptimizationRequest.md) | The data needed to optimize the routes. | 

### Return type

[**OptimizationIdentifier**](OptimizationIdentifier.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopOptimization

> StopOptimization(ctx, id).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A unique identifier of the optimization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OptimizationAPI.StopOptimization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.StopOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique identifier of the optimization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopOptimizationRequest struct via the builder pattern


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

