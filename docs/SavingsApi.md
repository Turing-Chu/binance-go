# \SavingsApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1LendingCustomizedFixedPurchasePost**](SavingsApi.md#SapiV1LendingCustomizedFixedPurchasePost) | **Post** /sapi/v1/lending/customizedFixed/purchase | Purchase Fixed/Activity Project (USER_DATA)
[**SapiV1LendingDailyProductListGet**](SavingsApi.md#SapiV1LendingDailyProductListGet) | **Get** /sapi/v1/lending/daily/product/list | Get Flexible Product List (USER_DATA)
[**SapiV1LendingDailyPurchasePost**](SavingsApi.md#SapiV1LendingDailyPurchasePost) | **Post** /sapi/v1/lending/daily/purchase | Purchase Flexible Product (USER_DATA)
[**SapiV1LendingDailyRedeemPost**](SavingsApi.md#SapiV1LendingDailyRedeemPost) | **Post** /sapi/v1/lending/daily/redeem | Redeem Flexible Product (USER_DATA)
[**SapiV1LendingDailyTokenPositionGet**](SavingsApi.md#SapiV1LendingDailyTokenPositionGet) | **Get** /sapi/v1/lending/daily/token/position | Get Flexible Product Position (USER_DATA)
[**SapiV1LendingDailyUserLeftQuotaGet**](SavingsApi.md#SapiV1LendingDailyUserLeftQuotaGet) | **Get** /sapi/v1/lending/daily/userLeftQuota | Get Left Daily Purchase Quota of Flexible Product (USER_DATA)
[**SapiV1LendingDailyUserRedemptionQuotaGet**](SavingsApi.md#SapiV1LendingDailyUserRedemptionQuotaGet) | **Get** /sapi/v1/lending/daily/userRedemptionQuota | Get Left Daily Redemption Quota of Flexible Product (USER_DATA)
[**SapiV1LendingPositionChangedPost**](SavingsApi.md#SapiV1LendingPositionChangedPost) | **Post** /sapi/v1/lending/positionChanged | Change Fixed/Activity Position to Daily Position (USER_DATA)
[**SapiV1LendingProjectListGet**](SavingsApi.md#SapiV1LendingProjectListGet) | **Get** /sapi/v1/lending/project/list | Get Fixed/Activity Project List(USER_DATA)
[**SapiV1LendingProjectPositionListGet**](SavingsApi.md#SapiV1LendingProjectPositionListGet) | **Get** /sapi/v1/lending/project/position/list | Get Fixed/Activity Project Position (USER_DATA)
[**SapiV1LendingUnionAccountGet**](SavingsApi.md#SapiV1LendingUnionAccountGet) | **Get** /sapi/v1/lending/union/account | Lending Account (USER_DATA)
[**SapiV1LendingUnionInterestHistoryGet**](SavingsApi.md#SapiV1LendingUnionInterestHistoryGet) | **Get** /sapi/v1/lending/union/interestHistory | Get Interest History (USER_DATA)
[**SapiV1LendingUnionPurchaseRecordGet**](SavingsApi.md#SapiV1LendingUnionPurchaseRecordGet) | **Get** /sapi/v1/lending/union/purchaseRecord | Get Purchase Record (USER_DATA)
[**SapiV1LendingUnionRedemptionRecordGet**](SavingsApi.md#SapiV1LendingUnionRedemptionRecordGet) | **Get** /sapi/v1/lending/union/redemptionRecord | Get Redemption Record (USER_DATA)



## SapiV1LendingCustomizedFixedPurchasePost

> InlineResponse20080 SapiV1LendingCustomizedFixedPurchasePost(ctx).ProjectId(projectId).Lot(lot).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Purchase Fixed/Activity Project (USER_DATA)



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
    projectId := "projectId_example" // string | 
    lot := "lot_example" // string | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingCustomizedFixedPurchasePost(context.Background()).ProjectId(projectId).Lot(lot).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingCustomizedFixedPurchasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingCustomizedFixedPurchasePost`: InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingCustomizedFixedPurchasePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingCustomizedFixedPurchasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | 
 **lot** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20080**](InlineResponse20080.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingDailyProductListGet

> []InlineResponse20074 SapiV1LendingDailyProductListGet(ctx).Status(status).Timestamp(timestamp).Signature(signature).Featured(featured).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Flexible Product List (USER_DATA)



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
    status := "status_example" // string | \"ALL\", \"SUBSCRIBABLE\", \"UNSUBSCRIBABLE\"; Default: 'ALL'
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    featured := "featured_example" // string | \"ALL\", \"TRUE\"; Default: \"ALL\" (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingDailyProductListGet(context.Background()).Status(status).Timestamp(timestamp).Signature(signature).Featured(featured).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingDailyProductListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingDailyProductListGet`: []InlineResponse20074
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingDailyProductListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingDailyProductListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | \&quot;ALL\&quot;, \&quot;SUBSCRIBABLE\&quot;, \&quot;UNSUBSCRIBABLE\&quot;; Default: &#39;ALL&#39; | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **featured** | **string** | \&quot;ALL\&quot;, \&quot;TRUE\&quot;; Default: \&quot;ALL\&quot; | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20074**](InlineResponse20074.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingDailyPurchasePost

> InlineResponse20076 SapiV1LendingDailyPurchasePost(ctx).ProductId(productId).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Purchase Flexible Product (USER_DATA)



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
    productId := "productId_example" // string | 
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingDailyPurchasePost(context.Background()).ProductId(productId).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingDailyPurchasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingDailyPurchasePost`: InlineResponse20076
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingDailyPurchasePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingDailyPurchasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20076**](InlineResponse20076.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingDailyRedeemPost

> map[string]interface{} SapiV1LendingDailyRedeemPost(ctx).ProductId(productId).Amount(amount).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Redeem Flexible Product (USER_DATA)



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
    productId := "productId_example" // string | 
    amount := float64(1.01) // float64 | 
    type_ := "type__example" // string | \"FAST\", \"NORMAL\"
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingDailyRedeemPost(context.Background()).ProductId(productId).Amount(amount).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingDailyRedeemPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingDailyRedeemPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingDailyRedeemPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingDailyRedeemPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **amount** | **float64** |  | 
 **type_** | **string** | \&quot;FAST\&quot;, \&quot;NORMAL\&quot; | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

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


## SapiV1LendingDailyTokenPositionGet

> []InlineResponse20078 SapiV1LendingDailyTokenPositionGet(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Get Flexible Product Position (USER_DATA)



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
    asset := "BTC" // string | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingDailyTokenPositionGet(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingDailyTokenPositionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingDailyTokenPositionGet`: []InlineResponse20078
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingDailyTokenPositionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingDailyTokenPositionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20078**](InlineResponse20078.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingDailyUserLeftQuotaGet

> InlineResponse20075 SapiV1LendingDailyUserLeftQuotaGet(ctx).ProductId(productId).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Get Left Daily Purchase Quota of Flexible Product (USER_DATA)



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
    productId := "productId_example" // string | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingDailyUserLeftQuotaGet(context.Background()).ProductId(productId).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingDailyUserLeftQuotaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingDailyUserLeftQuotaGet`: InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingDailyUserLeftQuotaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingDailyUserLeftQuotaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20075**](InlineResponse20075.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingDailyUserRedemptionQuotaGet

> InlineResponse20077 SapiV1LendingDailyUserRedemptionQuotaGet(ctx).ProductId(productId).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Get Left Daily Redemption Quota of Flexible Product (USER_DATA)



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
    productId := "productId_example" // string | 
    type_ := "type__example" // string | \"FAST\", \"NORMAL\"
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingDailyUserRedemptionQuotaGet(context.Background()).ProductId(productId).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingDailyUserRedemptionQuotaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingDailyUserRedemptionQuotaGet`: InlineResponse20077
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingDailyUserRedemptionQuotaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingDailyUserRedemptionQuotaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **type_** | **string** | \&quot;FAST\&quot;, \&quot;NORMAL\&quot; | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20077**](InlineResponse20077.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingPositionChangedPost

> InlineResponse20084 SapiV1LendingPositionChangedPost(ctx).ProjectId(projectId).Lot(lot).Timestamp(timestamp).Signature(signature).PositionId(positionId).RecvWindow(recvWindow).Execute()

Change Fixed/Activity Position to Daily Position (USER_DATA)



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
    projectId := "projectId_example" // string | 
    lot := "lot_example" // string | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    positionId := "positionId_example" // string |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingPositionChangedPost(context.Background()).ProjectId(projectId).Lot(lot).Timestamp(timestamp).Signature(signature).PositionId(positionId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingPositionChangedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingPositionChangedPost`: InlineResponse20084
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingPositionChangedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingPositionChangedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | 
 **lot** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **positionId** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20084**](InlineResponse20084.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingProjectListGet

> []InlineResponse20079 SapiV1LendingProjectListGet(ctx).Type_(type_).Status(status).IsSortAsc(isSortAsc).SortBy(sortBy).Timestamp(timestamp).Signature(signature).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Fixed/Activity Project List(USER_DATA)



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
    type_ := "type__example" // string | \"ACTIVITY\", \"CUSTOMIZED_FIXED\"
    status := "status_example" // string | \"ALL\", \"SUBSCRIBABLE\", \"UNSUBSCRIBABLE\"; Default: 'ALL'
    isSortAsc := true // bool | default \"true\"
    sortBy := "sortBy_example" // string | \"START_TIME\", \"LOT_SIZE\", \"INTEREST_RATE\", \"DURATION\"; default \"START_TIME
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    asset := "BNB" // string |  (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingProjectListGet(context.Background()).Type_(type_).Status(status).IsSortAsc(isSortAsc).SortBy(sortBy).Timestamp(timestamp).Signature(signature).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingProjectListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingProjectListGet`: []InlineResponse20079
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingProjectListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingProjectListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | \&quot;ACTIVITY\&quot;, \&quot;CUSTOMIZED_FIXED\&quot; | 
 **status** | **string** | \&quot;ALL\&quot;, \&quot;SUBSCRIBABLE\&quot;, \&quot;UNSUBSCRIBABLE\&quot;; Default: &#39;ALL&#39; | 
 **isSortAsc** | **bool** | default \&quot;true\&quot; | 
 **sortBy** | **string** | \&quot;START_TIME\&quot;, \&quot;LOT_SIZE\&quot;, \&quot;INTEREST_RATE\&quot;, \&quot;DURATION\&quot;; default \&quot;START_TIME | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20079**](InlineResponse20079.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingProjectPositionListGet

> []InlineResponse20081 SapiV1LendingProjectPositionListGet(ctx).Asset(asset).ProjectId(projectId).Status(status).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Get Fixed/Activity Project Position (USER_DATA)



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
    asset := "BTC" // string | 
    projectId := "projectId_example" // string | 
    status := "status_example" // string | \"ALL\", \"SUBSCRIBABLE\", \"UNSUBSCRIBABLE\"; Default: 'ALL'
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingProjectPositionListGet(context.Background()).Asset(asset).ProjectId(projectId).Status(status).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingProjectPositionListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingProjectPositionListGet`: []InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingProjectPositionListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingProjectPositionListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **projectId** | **string** |  | 
 **status** | **string** | \&quot;ALL\&quot;, \&quot;SUBSCRIBABLE\&quot;, \&quot;UNSUBSCRIBABLE\&quot;; Default: &#39;ALL&#39; | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20081**](InlineResponse20081.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingUnionAccountGet

> InlineResponse20082 SapiV1LendingUnionAccountGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Lending Account (USER_DATA)



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
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingUnionAccountGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingUnionAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingUnionAccountGet`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingUnionAccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingUnionAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20082**](InlineResponse20082.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingUnionInterestHistoryGet

> []InlineResponse20083 SapiV1LendingUnionInterestHistoryGet(ctx).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Interest History (USER_DATA)



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
    lendingType := "lendingType_example" // string | \"DAILY\" for flexible, \"ACTIVITY\" for activity, \"CUSTOMIZED_FIXED\" for fixed
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    asset := "BNB" // string |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingUnionInterestHistoryGet(context.Background()).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingUnionInterestHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingUnionInterestHistoryGet`: []InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingUnionInterestHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingUnionInterestHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lendingType** | **string** | \&quot;DAILY\&quot; for flexible, \&quot;ACTIVITY\&quot; for activity, \&quot;CUSTOMIZED_FIXED\&quot; for fixed | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20083**](InlineResponse20083.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingUnionPurchaseRecordGet

> OneOfarrayarray SapiV1LendingUnionPurchaseRecordGet(ctx).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Purchase Record (USER_DATA)



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
    lendingType := "lendingType_example" // string | \"DAILY\" for flexible, \"ACTIVITY\" for activity, \"CUSTOMIZED_FIXED\" for fixed
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    asset := "BNB" // string |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingUnionPurchaseRecordGet(context.Background()).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingUnionPurchaseRecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingUnionPurchaseRecordGet`: OneOfarrayarray
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingUnionPurchaseRecordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingUnionPurchaseRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lendingType** | **string** | \&quot;DAILY\&quot; for flexible, \&quot;ACTIVITY\&quot; for activity, \&quot;CUSTOMIZED_FIXED\&quot; for fixed | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfarrayarray**](oneOf&lt;array,array&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1LendingUnionRedemptionRecordGet

> OneOfarrayarray SapiV1LendingUnionRedemptionRecordGet(ctx).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Redemption Record (USER_DATA)



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
    lendingType := "lendingType_example" // string | \"DAILY\" for flexible, \"ACTIVITY\" for activity, \"CUSTOMIZED_FIXED\" for fixed
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    asset := "BNB" // string |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavingsApi.SapiV1LendingUnionRedemptionRecordGet(context.Background()).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.SapiV1LendingUnionRedemptionRecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1LendingUnionRedemptionRecordGet`: OneOfarrayarray
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.SapiV1LendingUnionRedemptionRecordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1LendingUnionRedemptionRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lendingType** | **string** | \&quot;DAILY\&quot; for flexible, \&quot;ACTIVITY\&quot; for activity, \&quot;CUSTOMIZED_FIXED\&quot; for fixed | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfarrayarray**](oneOf&lt;array,array&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

