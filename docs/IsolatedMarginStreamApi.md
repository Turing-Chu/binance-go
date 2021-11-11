# \IsolatedMarginStreamApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1UserDataStreamIsolatedDelete**](IsolatedMarginStreamApi.md#SapiV1UserDataStreamIsolatedDelete) | **Delete** /sapi/v1/userDataStream/isolated | Close a ListenKey (USER_STREAM)
[**SapiV1UserDataStreamIsolatedPost**](IsolatedMarginStreamApi.md#SapiV1UserDataStreamIsolatedPost) | **Post** /sapi/v1/userDataStream/isolated | Generate a Listen Key (USER_STREAM)
[**SapiV1UserDataStreamIsolatedPut**](IsolatedMarginStreamApi.md#SapiV1UserDataStreamIsolatedPut) | **Put** /sapi/v1/userDataStream/isolated | Ping/Keep-alive a Listen Key (USER_STREAM)



## SapiV1UserDataStreamIsolatedDelete

> map[string]interface{} SapiV1UserDataStreamIsolatedDelete(ctx).ListenKey(listenKey).Execute()

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
    resp, r, err := api_client.IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedDelete(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1UserDataStreamIsolatedDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1UserDataStreamIsolatedDeleteRequest struct via the builder pattern


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


## SapiV1UserDataStreamIsolatedPost

> InlineResponse20071 SapiV1UserDataStreamIsolatedPost(ctx).Execute()

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
    resp, r, err := api_client.IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1UserDataStreamIsolatedPost`: InlineResponse20071
    fmt.Fprintf(os.Stdout, "Response from `IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1UserDataStreamIsolatedPostRequest struct via the builder pattern


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


## SapiV1UserDataStreamIsolatedPut

> map[string]interface{} SapiV1UserDataStreamIsolatedPut(ctx).ListenKey(listenKey).Execute()

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
    resp, r, err := api_client.IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedPut(context.Background()).ListenKey(listenKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1UserDataStreamIsolatedPut`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IsolatedMarginStreamApi.SapiV1UserDataStreamIsolatedPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1UserDataStreamIsolatedPutRequest struct via the builder pattern


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

