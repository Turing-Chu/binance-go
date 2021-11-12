# \StreamApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseSpotListenKey**](StreamApi.md#CloseSpotListenKey) | **Delete** /api/v3/userDataStream | Close a ListenKey (USER_STREAM)
[**CreateSpotListenKey**](StreamApi.md#CreateSpotListenKey) | **Post** /api/v3/userDataStream | Create a ListenKey (USER_STREAM)
[**DelaySpotListenKey**](StreamApi.md#DelaySpotListenKey) | **Put** /api/v3/userDataStream | Ping/Keep-alive a ListenKey (USER_STREAM)



## CloseSpotListenKey

> map[string]interface{} CloseSpotListenKey(ctx).ListenKey(listenKey).Execute()

Close a ListenKey (USER_STREAM)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    listenKey := "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1" // string | User websocket listen key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamApi.CloseSpotListenKey(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamApi.CloseSpotListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseSpotListenKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StreamApi.CloseSpotListenKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseSpotListenKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** | User websocket listen key | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpotListenKey

> InlineResponse20070 CreateSpotListenKey(ctx).Execute()

Create a ListenKey (USER_STREAM)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamApi.CreateSpotListenKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamApi.CreateSpotListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpotListenKey`: InlineResponse20070
    fmt.Fprintf(os.Stdout, "Response from `StreamApi.CreateSpotListenKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpotListenKeyRequest struct via the builder pattern


### Return type

[**InlineResponse20070**](InlineResponse20070.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelaySpotListenKey

> map[string]interface{} DelaySpotListenKey(ctx).ListenKey(listenKey).Execute()

Ping/Keep-alive a ListenKey (USER_STREAM)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    listenKey := "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1" // string | User websocket listen key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StreamApi.DelaySpotListenKey(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamApi.DelaySpotListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DelaySpotListenKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StreamApi.DelaySpotListenKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDelaySpotListenKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** | User websocket listen key | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

