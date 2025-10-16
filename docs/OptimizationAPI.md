# \OptimizationAPI

All URIs are relative to *https://api.myptv.com/routeoptimization/optiflow/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOptimization**](OptimizationAPI.md#DeleteOptimization) | **Delete** /optimizations/{id} | 
[**GetOptimizationProgress**](OptimizationAPI.md#GetOptimizationProgress) | **Get** /optimizations/{id}/progress | 
[**GetOptimizationRequest**](OptimizationAPI.md#GetOptimizationRequest) | **Get** /optimizations/{id}/request | 
[**GetOptimizationResult**](OptimizationAPI.md#GetOptimizationResult) | **Get** /optimizations/{id} | 
[**ListOptimizations**](OptimizationAPI.md#ListOptimizations) | **Get** /optimizations | 
[**StartOptimization**](OptimizationAPI.md#StartOptimization) | **Post** /optimizations | 
[**StopOptimization**](OptimizationAPI.md#StopOptimization) | **Post** /optimizations/{id}/stop | 



## DeleteOptimization

> DeleteOptimization(ctx, id).Execute()





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
	r, err := apiClient.OptimizationAPI.DeleteOptimization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.DeleteOptimization``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOptimizationRequest struct via the builder pattern


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


## GetOptimizationProgress

> OptimizationProgress GetOptimizationProgress(ctx, id).Execute()





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
	resp, r, err := apiClient.OptimizationAPI.GetOptimizationProgress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.GetOptimizationProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptimizationProgress`: OptimizationProgress
	fmt.Fprintf(os.Stdout, "Response from `OptimizationAPI.GetOptimizationProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique identifier of the optimization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptimizationProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OptimizationProgress**](OptimizationProgress.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptimizationRequest

> OptimizationRequest GetOptimizationRequest(ctx, id).Execute()





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
	resp, r, err := apiClient.OptimizationAPI.GetOptimizationRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.GetOptimizationRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptimizationRequest`: OptimizationRequest
	fmt.Fprintf(os.Stdout, "Response from `OptimizationAPI.GetOptimizationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A unique identifier of the optimization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptimizationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OptimizationRequest**](OptimizationRequest.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ListOptimizations

> OptimizationSummaries ListOptimizations(ctx).Cursor(cursor).Limit(limit).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Name(name).Tags(tags).Statuses(statuses).Execute()





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
	cursor := "cursor_example" // string | A pagination cursor for retrieving the next page of results. Use the `nextCursor` value from a previous response to fetch subsequent pages. Omit this parameter to retrieve the first page of optimizations. (optional)
	limit := int32(56) // int32 | Maximum number of optimizations to return per page. If more results exist, the response includes a `nextCursor` value for pagination. Use this cursor in subsequent requests to retrieve additional pages. (optional) (default to 50)
	createdBefore := time.Now() // time.Time | Filter optimizations created before this date and time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). (optional)
	createdAfter := time.Now() // time.Time | Filter optimizations created after this date and time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). (optional)
	name := "name_example" // string | Filter optimizations with this name. (optional)
	tags := []string{"Inner_example"} // []string | Filter optimizations containing all of these tags. (optional)
	statuses := []openapiclient.OptimizationStatus{openapiclient.OptimizationStatus("QUEUING")} // []OptimizationStatus | Filter optimizations having any of these statuses. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptimizationAPI.ListOptimizations(context.Background()).Cursor(cursor).Limit(limit).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Name(name).Tags(tags).Statuses(statuses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.ListOptimizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptimizations`: OptimizationSummaries
	fmt.Fprintf(os.Stdout, "Response from `OptimizationAPI.ListOptimizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptimizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | A pagination cursor for retrieving the next page of results. Use the &#x60;nextCursor&#x60; value from a previous response to fetch subsequent pages. Omit this parameter to retrieve the first page of optimizations. | 
 **limit** | **int32** | Maximum number of optimizations to return per page. If more results exist, the response includes a &#x60;nextCursor&#x60; value for pagination. Use this cursor in subsequent requests to retrieve additional pages. | [default to 50]
 **createdBefore** | **time.Time** | Filter optimizations created before this date and time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | 
 **createdAfter** | **time.Time** | Filter optimizations created after this date and time. Formatted according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339#section-5.6). | 
 **name** | **string** | Filter optimizations with this name. | 
 **tags** | **[]string** | Filter optimizations containing all of these tags. | 
 **statuses** | [**[]OptimizationStatus**](OptimizationStatus.md) | Filter optimizations having any of these statuses. | 

### Return type

[**OptimizationSummaries**](OptimizationSummaries.md)

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

