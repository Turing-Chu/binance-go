# \SavingsApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LendingChangeFixedOrActivityToDailyPosition**](SavingsApi.md#LendingChangeFixedOrActivityToDailyPosition) | **Post** /sapi/v1/lending/positionChanged | Change Fixed/Activity Position to Daily Position (USER_DATA)
[**LendingGetAccount**](SavingsApi.md#LendingGetAccount) | **Get** /sapi/v1/lending/union/account | Lending Account (USER_DATA)
[**LendingGetFixedOrActivityProjectPosition**](SavingsApi.md#LendingGetFixedOrActivityProjectPosition) | **Get** /sapi/v1/lending/project/position/list | Get Fixed/Activity Project Position (USER_DATA)
[**LendingGetFixedOrActivityProjects**](SavingsApi.md#LendingGetFixedOrActivityProjects) | **Get** /sapi/v1/lending/project/list | Get Fixed/Activity Project List(USER_DATA)
[**LendingGetFlexibleProductPosition**](SavingsApi.md#LendingGetFlexibleProductPosition) | **Get** /sapi/v1/lending/daily/token/position | Get Flexible Product Position (USER_DATA)
[**LendingGetInterestHistory**](SavingsApi.md#LendingGetInterestHistory) | **Get** /sapi/v1/lending/union/interestHistory | Get Interest History (USER_DATA)
[**LendingGetLeftDailyPurchaseQuotaOfFlexibleProduct**](SavingsApi.md#LendingGetLeftDailyPurchaseQuotaOfFlexibleProduct) | **Get** /sapi/v1/lending/daily/userLeftQuota | Get Left Daily Purchase Quota of Flexible Product (USER_DATA)
[**LendingGetLeftDailyRedemptionQuotaOfFlexibleProduct**](SavingsApi.md#LendingGetLeftDailyRedemptionQuotaOfFlexibleProduct) | **Get** /sapi/v1/lending/daily/userRedemptionQuota | Get Left Daily Redemption Quota of Flexible Product (USER_DATA)
[**LendingGetPurchaseRecord**](SavingsApi.md#LendingGetPurchaseRecord) | **Get** /sapi/v1/lending/union/purchaseRecord | Get Purchase Record (USER_DATA)
[**LendingGetRedemptionRecord**](SavingsApi.md#LendingGetRedemptionRecord) | **Get** /sapi/v1/lending/union/redemptionRecord | Get Redemption Record (USER_DATA)
[**LendingListFlexibleProducts**](SavingsApi.md#LendingListFlexibleProducts) | **Get** /sapi/v1/lending/daily/product/list | Get Flexible Product List (USER_DATA)
[**LendingPurchaseFixedOrActivityProject**](SavingsApi.md#LendingPurchaseFixedOrActivityProject) | **Post** /sapi/v1/lending/customizedFixed/purchase | Purchase Fixed/Activity Project (USER_DATA)
[**LendingPurchaseFlexibleProduct**](SavingsApi.md#LendingPurchaseFlexibleProduct) | **Post** /sapi/v1/lending/daily/purchase | Purchase Flexible Product (USER_DATA)
[**LendingRedeemFlexibleProduct**](SavingsApi.md#LendingRedeemFlexibleProduct) | **Post** /sapi/v1/lending/daily/redeem | Redeem Flexible Product (USER_DATA)



## LendingChangeFixedOrActivityToDailyPosition

> InlineResponse20084 LendingChangeFixedOrActivityToDailyPosition(ctx).ProjectId(projectId).Lot(lot).Timestamp(timestamp).Signature(signature).PositionId(positionId).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingChangeFixedOrActivityToDailyPosition(context.Background()).ProjectId(projectId).Lot(lot).Timestamp(timestamp).Signature(signature).PositionId(positionId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingChangeFixedOrActivityToDailyPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingChangeFixedOrActivityToDailyPosition`: InlineResponse20084
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingChangeFixedOrActivityToDailyPosition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingChangeFixedOrActivityToDailyPositionRequest struct via the builder pattern


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


## LendingGetAccount

> InlineResponse20082 LendingGetAccount(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetAccount(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetAccount`: InlineResponse20082
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetAccountRequest struct via the builder pattern


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


## LendingGetFixedOrActivityProjectPosition

> []InlineResponse20081 LendingGetFixedOrActivityProjectPosition(ctx).Asset(asset).ProjectId(projectId).Status(status).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetFixedOrActivityProjectPosition(context.Background()).Asset(asset).ProjectId(projectId).Status(status).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetFixedOrActivityProjectPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetFixedOrActivityProjectPosition`: []InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetFixedOrActivityProjectPosition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetFixedOrActivityProjectPositionRequest struct via the builder pattern


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


## LendingGetFixedOrActivityProjects

> []InlineResponse20079 LendingGetFixedOrActivityProjects(ctx).Type_(type_).Status(status).IsSortAsc(isSortAsc).SortBy(sortBy).Timestamp(timestamp).Signature(signature).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetFixedOrActivityProjects(context.Background()).Type_(type_).Status(status).IsSortAsc(isSortAsc).SortBy(sortBy).Timestamp(timestamp).Signature(signature).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetFixedOrActivityProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetFixedOrActivityProjects`: []InlineResponse20079
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetFixedOrActivityProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetFixedOrActivityProjectsRequest struct via the builder pattern


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


## LendingGetFlexibleProductPosition

> []InlineResponse20078 LendingGetFlexibleProductPosition(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetFlexibleProductPosition(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetFlexibleProductPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetFlexibleProductPosition`: []InlineResponse20078
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetFlexibleProductPosition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetFlexibleProductPositionRequest struct via the builder pattern


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


## LendingGetInterestHistory

> []InlineResponse20083 LendingGetInterestHistory(ctx).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetInterestHistory(context.Background()).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetInterestHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetInterestHistory`: []InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetInterestHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetInterestHistoryRequest struct via the builder pattern


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


## LendingGetLeftDailyPurchaseQuotaOfFlexibleProduct

> InlineResponse20075 LendingGetLeftDailyPurchaseQuotaOfFlexibleProduct(ctx).ProductId(productId).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetLeftDailyPurchaseQuotaOfFlexibleProduct(context.Background()).ProductId(productId).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetLeftDailyPurchaseQuotaOfFlexibleProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetLeftDailyPurchaseQuotaOfFlexibleProduct`: InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetLeftDailyPurchaseQuotaOfFlexibleProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetLeftDailyPurchaseQuotaOfFlexibleProductRequest struct via the builder pattern


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


## LendingGetLeftDailyRedemptionQuotaOfFlexibleProduct

> InlineResponse20077 LendingGetLeftDailyRedemptionQuotaOfFlexibleProduct(ctx).ProductId(productId).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetLeftDailyRedemptionQuotaOfFlexibleProduct(context.Background()).ProductId(productId).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetLeftDailyRedemptionQuotaOfFlexibleProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetLeftDailyRedemptionQuotaOfFlexibleProduct`: InlineResponse20077
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetLeftDailyRedemptionQuotaOfFlexibleProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetLeftDailyRedemptionQuotaOfFlexibleProductRequest struct via the builder pattern


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


## LendingGetPurchaseRecord

> OneOfsavingsFlexiblePurchaseRecordsavingsFixedActivityPurchaseRecord LendingGetPurchaseRecord(ctx).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetPurchaseRecord(context.Background()).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetPurchaseRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetPurchaseRecord`: OneOfsavingsFlexiblePurchaseRecordsavingsFixedActivityPurchaseRecord
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetPurchaseRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetPurchaseRecordRequest struct via the builder pattern


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

[**OneOfsavingsFlexiblePurchaseRecordsavingsFixedActivityPurchaseRecord**](oneOf&lt;savingsFlexiblePurchaseRecord,savingsFixedActivityPurchaseRecord&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LendingGetRedemptionRecord

> OneOfsavingsFlexibleRedemptionRecordsavingsFixedActivityRedemptionRecord LendingGetRedemptionRecord(ctx).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingGetRedemptionRecord(context.Background()).LendingType(lendingType).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingGetRedemptionRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingGetRedemptionRecord`: OneOfsavingsFlexibleRedemptionRecordsavingsFixedActivityRedemptionRecord
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingGetRedemptionRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingGetRedemptionRecordRequest struct via the builder pattern


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

[**OneOfsavingsFlexibleRedemptionRecordsavingsFixedActivityRedemptionRecord**](oneOf&lt;savingsFlexibleRedemptionRecord,savingsFixedActivityRedemptionRecord&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LendingListFlexibleProducts

> []InlineResponse20074 LendingListFlexibleProducts(ctx).Status(status).Timestamp(timestamp).Signature(signature).Featured(featured).Current(current).Size(size).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingListFlexibleProducts(context.Background()).Status(status).Timestamp(timestamp).Signature(signature).Featured(featured).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingListFlexibleProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingListFlexibleProducts`: []InlineResponse20074
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingListFlexibleProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingListFlexibleProductsRequest struct via the builder pattern


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


## LendingPurchaseFixedOrActivityProject

> InlineResponse20080 LendingPurchaseFixedOrActivityProject(ctx).ProjectId(projectId).Lot(lot).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingPurchaseFixedOrActivityProject(context.Background()).ProjectId(projectId).Lot(lot).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingPurchaseFixedOrActivityProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingPurchaseFixedOrActivityProject`: InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingPurchaseFixedOrActivityProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingPurchaseFixedOrActivityProjectRequest struct via the builder pattern


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


## LendingPurchaseFlexibleProduct

> InlineResponse20076 LendingPurchaseFlexibleProduct(ctx).ProductId(productId).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingPurchaseFlexibleProduct(context.Background()).ProductId(productId).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingPurchaseFlexibleProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingPurchaseFlexibleProduct`: InlineResponse20076
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingPurchaseFlexibleProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingPurchaseFlexibleProductRequest struct via the builder pattern


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


## LendingRedeemFlexibleProduct

> map[string]interface{} LendingRedeemFlexibleProduct(ctx).ProductId(productId).Amount(amount).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.SavingsApi.LendingRedeemFlexibleProduct(context.Background()).ProductId(productId).Amount(amount).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavingsApi.LendingRedeemFlexibleProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LendingRedeemFlexibleProduct`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SavingsApi.LendingRedeemFlexibleProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLendingRedeemFlexibleProductRequest struct via the builder pattern


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

