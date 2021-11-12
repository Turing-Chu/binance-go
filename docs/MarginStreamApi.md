# \MarginStreamApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseMarginListenKey**](MarginStreamApi.md#CloseMarginListenKey) | **Delete** /sapi/v1/userDataStream | Close a ListenKey (USER_STREAM)
[**CreateMarginListenKey**](MarginStreamApi.md#CreateMarginListenKey) | **Post** /sapi/v1/userDataStream | Create a ListenKey (USER_STREAM)
[**DelayMarginListenKey**](MarginStreamApi.md#DelayMarginListenKey) | **Put** /sapi/v1/userDataStream | Ping/Keep-alive a ListenKey (USER_STREAM)



## CloseMarginListenKey

> map[string]interface{} CloseMarginListenKey(ctx).ListenKey(listenKey).Execute()

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
    resp, r, err := api_client.MarginStreamApi.CloseMarginListenKey(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginStreamApi.CloseMarginListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseMarginListenKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginStreamApi.CloseMarginListenKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseMarginListenKeyRequest struct via the builder pattern


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


## CreateMarginListenKey

> InlineResponse20070 CreateMarginListenKey(ctx).Execute()

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
    resp, r, err := api_client.MarginStreamApi.CreateMarginListenKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginStreamApi.CreateMarginListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMarginListenKey`: InlineResponse20070
    fmt.Fprintf(os.Stdout, "Response from `MarginStreamApi.CreateMarginListenKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginListenKeyRequest struct via the builder pattern


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


## DelayMarginListenKey

> map[string]interface{} DelayMarginListenKey(ctx).ListenKey(listenKey).Execute()

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
    resp, r, err := api_client.MarginStreamApi.DelayMarginListenKey(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginStreamApi.DelayMarginListenKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DelayMarginListenKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginStreamApi.DelayMarginListenKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDelayMarginListenKeyRequest struct via the builder pattern


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

