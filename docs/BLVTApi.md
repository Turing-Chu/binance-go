# \BLVTApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1BlvtRedeemPost**](BLVTApi.md#SapiV1BlvtRedeemPost) | **Post** /sapi/v1/blvt/redeem | Redeem BLVT (USER_DATA)
[**SapiV1BlvtRedeemRecordGet**](BLVTApi.md#SapiV1BlvtRedeemRecordGet) | **Get** /sapi/v1/blvt/redeem/record | Redemption Record (USER_DATA)
[**SapiV1BlvtSubscribePost**](BLVTApi.md#SapiV1BlvtSubscribePost) | **Post** /sapi/v1/blvt/subscribe | Subscribe BLVT (USER_DATA)
[**SapiV1BlvtSubscribeRecordGet**](BLVTApi.md#SapiV1BlvtSubscribeRecordGet) | **Get** /sapi/v1/blvt/subscribe/record | Query Subscription Record (USER_DATA)
[**SapiV1BlvtTokenInfoGet**](BLVTApi.md#SapiV1BlvtTokenInfoGet) | **Get** /sapi/v1/blvt/tokenInfo | BLVT Info (MARKET_DATA)
[**SapiV1BlvtUserLimitGet**](BLVTApi.md#SapiV1BlvtUserLimitGet) | **Get** /sapi/v1/blvt/userLimit | BLVT User Limit Info (USER_DATA)



## SapiV1BlvtRedeemPost

> InlineResponse200100 SapiV1BlvtRedeemPost(ctx).TokenName(tokenName).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Redeem BLVT (USER_DATA)



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
    tokenName := "tokenName_example" // string | BTCDOWN, BTCUP
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BLVTApi.SapiV1BlvtRedeemPost(context.Background()).TokenName(tokenName).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.SapiV1BlvtRedeemPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BlvtRedeemPost`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.SapiV1BlvtRedeemPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BlvtRedeemPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenName** | **string** | BTCDOWN, BTCUP | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1BlvtRedeemRecordGet

> []InlineResponse200101 SapiV1BlvtRedeemRecordGet(ctx).Timestamp(timestamp).Signature(signature).TokenName(tokenName).Id(id).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Redemption Record (USER_DATA)



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
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    tokenName := "tokenName_example" // string | BTCDOWN, BTCUP (optional)
    id := int64(789) // int64 |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(56) // int32 | default 1000, max 1000 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BLVTApi.SapiV1BlvtRedeemRecordGet(context.Background()).Timestamp(timestamp).Signature(signature).TokenName(tokenName).Id(id).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.SapiV1BlvtRedeemRecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BlvtRedeemRecordGet`: []InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.SapiV1BlvtRedeemRecordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BlvtRedeemRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **tokenName** | **string** | BTCDOWN, BTCUP | 
 **id** | **int64** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | default 1000, max 1000 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse200101**](InlineResponse200101.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1BlvtSubscribePost

> InlineResponse20098 SapiV1BlvtSubscribePost(ctx).TokenName(tokenName).Cost(cost).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Subscribe BLVT (USER_DATA)



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
    tokenName := "tokenName_example" // string | BTCDOWN, BTCUP
    cost := float64(1.2) // float64 | Spot balance
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BLVTApi.SapiV1BlvtSubscribePost(context.Background()).TokenName(tokenName).Cost(cost).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.SapiV1BlvtSubscribePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BlvtSubscribePost`: InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.SapiV1BlvtSubscribePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BlvtSubscribePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenName** | **string** | BTCDOWN, BTCUP | 
 **cost** | **float64** | Spot balance | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20098**](InlineResponse20098.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1BlvtSubscribeRecordGet

> InlineResponse20099 SapiV1BlvtSubscribeRecordGet(ctx).Timestamp(timestamp).Signature(signature).TokenName(tokenName).Id(id).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Subscription Record (USER_DATA)



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
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    tokenName := "tokenName_example" // string | BTCDOWN, BTCUP (optional)
    id := int64(789) // int64 |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BLVTApi.SapiV1BlvtSubscribeRecordGet(context.Background()).Timestamp(timestamp).Signature(signature).TokenName(tokenName).Id(id).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.SapiV1BlvtSubscribeRecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BlvtSubscribeRecordGet`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.SapiV1BlvtSubscribeRecordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BlvtSubscribeRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **tokenName** | **string** | BTCDOWN, BTCUP | 
 **id** | **int64** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1BlvtTokenInfoGet

> []InlineResponse20097 SapiV1BlvtTokenInfoGet(ctx).TokenName(tokenName).Execute()

BLVT Info (MARKET_DATA)



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
    tokenName := "tokenName_example" // string | BTCDOWN, BTCUP (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BLVTApi.SapiV1BlvtTokenInfoGet(context.Background()).TokenName(tokenName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.SapiV1BlvtTokenInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BlvtTokenInfoGet`: []InlineResponse20097
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.SapiV1BlvtTokenInfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BlvtTokenInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenName** | **string** | BTCDOWN, BTCUP | 

### Return type

[**[]InlineResponse20097**](InlineResponse20097.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1BlvtUserLimitGet

> []InlineResponse200102 SapiV1BlvtUserLimitGet(ctx).Timestamp(timestamp).Signature(signature).TokenName(tokenName).RecvWindow(recvWindow).Execute()

BLVT User Limit Info (USER_DATA)



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
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    tokenName := "tokenName_example" // string | BTCDOWN, BTCUP (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BLVTApi.SapiV1BlvtUserLimitGet(context.Background()).Timestamp(timestamp).Signature(signature).TokenName(tokenName).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.SapiV1BlvtUserLimitGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BlvtUserLimitGet`: []InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.SapiV1BlvtUserLimitGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BlvtUserLimitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **tokenName** | **string** | BTCDOWN, BTCUP | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse200102**](InlineResponse200102.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

