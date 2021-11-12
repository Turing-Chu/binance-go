# \IsolatedMarginStreamApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseIsolatediMarginListenKey**](IsolatedMarginStreamApi.md#CloseIsolatediMarginListenKey) | **Delete** /sapi/v1/userDataStream/isolated | Close a ListenKey (USER_STREAM)
[**CreateIsolatediMarginListenKey**](IsolatedMarginStreamApi.md#CreateIsolatediMarginListenKey) | **Post** /sapi/v1/userDataStream/isolated | Generate a Listen Key (USER_STREAM)
[**DelayIsolatediMarginListenKey**](IsolatedMarginStreamApi.md#DelayIsolatediMarginListenKey) | **Put** /sapi/v1/userDataStream/isolated | Ping/Keep-alive a Listen Key (USER_STREAM)



## CloseIsolatediMarginListenKey

> map[string]interface{} CloseIsolatediMarginListenKey(ctx).ListenKey(listenKey).Execute()

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
    resp, r, err := api_client.IsolatedMarginStreamApi.CloseIsolatediMarginListenKey(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsolatedMarginStreamApi.CloseIsolatediMarginListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseIsolatediMarginListenKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IsolatedMarginStreamApi.CloseIsolatediMarginListenKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseIsolatediMarginListenKeyRequest struct via the builder pattern


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


## CreateIsolatediMarginListenKey

> InlineResponse20071 CreateIsolatediMarginListenKey(ctx).Execute()

Generate a Listen Key (USER_STREAM)



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
    resp, r, err := api_client.IsolatedMarginStreamApi.CreateIsolatediMarginListenKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsolatedMarginStreamApi.CreateIsolatediMarginListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIsolatediMarginListenKey`: InlineResponse20071
    fmt.Fprintf(os.Stdout, "Response from `IsolatedMarginStreamApi.CreateIsolatediMarginListenKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIsolatediMarginListenKeyRequest struct via the builder pattern


### Return type

[**InlineResponse20071**](InlineResponse20071.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelayIsolatediMarginListenKey

> map[string]interface{} DelayIsolatediMarginListenKey(ctx).ListenKey(listenKey).Execute()

Ping/Keep-alive a Listen Key (USER_STREAM)



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
    resp, r, err := api_client.IsolatedMarginStreamApi.DelayIsolatediMarginListenKey(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsolatedMarginStreamApi.DelayIsolatediMarginListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DelayIsolatediMarginListenKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IsolatedMarginStreamApi.DelayIsolatediMarginListenKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDelayIsolatediMarginListenKeyRequest struct via the builder pattern


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

