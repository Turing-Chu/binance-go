# \BLVTApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlvtGetRedeemRecord**](BLVTApi.md#BlvtGetRedeemRecord) | **Get** /sapi/v1/blvt/redeem/record | Redemption Record (USER_DATA)
[**BlvtGetSubscribeRecord**](BLVTApi.md#BlvtGetSubscribeRecord) | **Get** /sapi/v1/blvt/subscribe/record | Query Subscription Record (USER_DATA)
[**BlvtGetTokenInfo**](BLVTApi.md#BlvtGetTokenInfo) | **Get** /sapi/v1/blvt/tokenInfo | BLVT Info (MARKET_DATA)
[**BlvtGetUserLimitInfo**](BLVTApi.md#BlvtGetUserLimitInfo) | **Get** /sapi/v1/blvt/userLimit | BLVT User Limit Info (USER_DATA)
[**BlvtRedeem**](BLVTApi.md#BlvtRedeem) | **Post** /sapi/v1/blvt/redeem | Redeem BLVT (USER_DATA)
[**BlvtSubscribe**](BLVTApi.md#BlvtSubscribe) | **Post** /sapi/v1/blvt/subscribe | Subscribe BLVT (USER_DATA)



## BlvtGetRedeemRecord

> []InlineResponse200101 BlvtGetRedeemRecord(ctx).Timestamp(timestamp).Signature(signature).TokenName(tokenName).Id(id).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BLVTApi.BlvtGetRedeemRecord(context.Background()).Timestamp(timestamp).Signature(signature).TokenName(tokenName).Id(id).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.BlvtGetRedeemRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlvtGetRedeemRecord`: []InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.BlvtGetRedeemRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlvtGetRedeemRecordRequest struct via the builder pattern


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


## BlvtGetSubscribeRecord

> InlineResponse20099 BlvtGetSubscribeRecord(ctx).Timestamp(timestamp).Signature(signature).TokenName(tokenName).Id(id).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BLVTApi.BlvtGetSubscribeRecord(context.Background()).Timestamp(timestamp).Signature(signature).TokenName(tokenName).Id(id).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.BlvtGetSubscribeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlvtGetSubscribeRecord`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.BlvtGetSubscribeRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlvtGetSubscribeRecordRequest struct via the builder pattern


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


## BlvtGetTokenInfo

> []InlineResponse20097 BlvtGetTokenInfo(ctx).TokenName(tokenName).Execute()

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
    resp, r, err := api_client.BLVTApi.BlvtGetTokenInfo(context.Background()).TokenName(tokenName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.BlvtGetTokenInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlvtGetTokenInfo`: []InlineResponse20097
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.BlvtGetTokenInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlvtGetTokenInfoRequest struct via the builder pattern


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


## BlvtGetUserLimitInfo

> []InlineResponse200102 BlvtGetUserLimitInfo(ctx).Timestamp(timestamp).Signature(signature).TokenName(tokenName).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BLVTApi.BlvtGetUserLimitInfo(context.Background()).Timestamp(timestamp).Signature(signature).TokenName(tokenName).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.BlvtGetUserLimitInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlvtGetUserLimitInfo`: []InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.BlvtGetUserLimitInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlvtGetUserLimitInfoRequest struct via the builder pattern


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


## BlvtRedeem

> InlineResponse200100 BlvtRedeem(ctx).TokenName(tokenName).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BLVTApi.BlvtRedeem(context.Background()).TokenName(tokenName).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.BlvtRedeem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlvtRedeem`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.BlvtRedeem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlvtRedeemRequest struct via the builder pattern


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


## BlvtSubscribe

> InlineResponse20098 BlvtSubscribe(ctx).TokenName(tokenName).Cost(cost).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BLVTApi.BlvtSubscribe(context.Background()).TokenName(tokenName).Cost(cost).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BLVTApi.BlvtSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlvtSubscribe`: InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `BLVTApi.BlvtSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlvtSubscribeRequest struct via the builder pattern


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

