# \WalletApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1AccountApiRestrictionsGet**](WalletApi.md#SapiV1AccountApiRestrictionsGet) | **Get** /sapi/v1/account/apiRestrictions | Get API Key Permission (USER_DATA)
[**SapiV1AccountApiTradingStatusGet**](WalletApi.md#SapiV1AccountApiTradingStatusGet) | **Get** /sapi/v1/account/apiTradingStatus | Account API Trading Status (USER_DATA)
[**SapiV1AccountDisableFastWithdrawSwitchPost**](WalletApi.md#SapiV1AccountDisableFastWithdrawSwitchPost) | **Post** /sapi/v1/account/disableFastWithdrawSwitch | Disable Fast Withdraw Switch (USER_DATA)
[**SapiV1AccountEnableFastWithdrawSwitchPost**](WalletApi.md#SapiV1AccountEnableFastWithdrawSwitchPost) | **Post** /sapi/v1/account/enableFastWithdrawSwitch | Enable Fast Withdraw Switch (USER_DATA)
[**SapiV1AccountSnapshotGet**](WalletApi.md#SapiV1AccountSnapshotGet) | **Get** /sapi/v1/accountSnapshot | Daily Account Snapshot (USER_DATA)
[**SapiV1AccountStatusGet**](WalletApi.md#SapiV1AccountStatusGet) | **Get** /sapi/v1/account/status | Account Status (USER_DATA)
[**SapiV1AssetAssetDetailGet**](WalletApi.md#SapiV1AssetAssetDetailGet) | **Get** /sapi/v1/asset/assetDetail | Asset Detail (USER_DATA)
[**SapiV1AssetAssetDividendGet**](WalletApi.md#SapiV1AssetAssetDividendGet) | **Get** /sapi/v1/asset/assetDividend | Asset Dividend Record (USER_DATA)
[**SapiV1AssetDribbletGet**](WalletApi.md#SapiV1AssetDribbletGet) | **Get** /sapi/v1/asset/dribblet | DustLog(USER_DATA)
[**SapiV1AssetDustPost**](WalletApi.md#SapiV1AssetDustPost) | **Post** /sapi/v1/asset/dust | Dust Transfer (USER_DATA)
[**SapiV1AssetGetFundingAssetPost**](WalletApi.md#SapiV1AssetGetFundingAssetPost) | **Post** /sapi/v1/asset/get-funding-asset | Funding Wallet (USER_DATA)
[**SapiV1AssetTradeFeeGet**](WalletApi.md#SapiV1AssetTradeFeeGet) | **Get** /sapi/v1/asset/tradeFee | Trade Fee (USER_DATA)
[**SapiV1AssetTransferGet**](WalletApi.md#SapiV1AssetTransferGet) | **Get** /sapi/v1/asset/transfer | Query User Universal Transfer History (USER_DATA)
[**SapiV1AssetTransferPost**](WalletApi.md#SapiV1AssetTransferPost) | **Post** /sapi/v1/asset/transfer | User Universal Transfer (USER_DATA)
[**SapiV1CapitalConfigGetallGet**](WalletApi.md#SapiV1CapitalConfigGetallGet) | **Get** /sapi/v1/capital/config/getall | All Coins&#39; Information (USER_DATA)
[**SapiV1CapitalDepositAddressGet**](WalletApi.md#SapiV1CapitalDepositAddressGet) | **Get** /sapi/v1/capital/deposit/address | Deposit Address (supporting network) (USER_DATA)
[**SapiV1CapitalDepositHisrecGet**](WalletApi.md#SapiV1CapitalDepositHisrecGet) | **Get** /sapi/v1/capital/deposit/hisrec | Deposit History（supporting network） (USER_DATA)
[**SapiV1CapitalWithdrawApplyPost**](WalletApi.md#SapiV1CapitalWithdrawApplyPost) | **Post** /sapi/v1/capital/withdraw/apply | Withdraw (USER_DATA)
[**SapiV1CapitalWithdrawHistoryGet**](WalletApi.md#SapiV1CapitalWithdrawHistoryGet) | **Get** /sapi/v1/capital/withdraw/history | Withdraw History (supporting network) (USER_DATA)
[**SapiV1SystemStatusGet**](WalletApi.md#SapiV1SystemStatusGet) | **Get** /sapi/v1/system/status | System Status (System)



## SapiV1AccountApiRestrictionsGet

> InlineResponse20045 SapiV1AccountApiRestrictionsGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Get API Key Permission (USER_DATA)



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
    resp, r, err := api_client.WalletApi.SapiV1AccountApiRestrictionsGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountApiRestrictionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AccountApiRestrictionsGet`: InlineResponse20045
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AccountApiRestrictionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountApiRestrictionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20045**](InlineResponse20045.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AccountApiTradingStatusGet

> InlineResponse20036 SapiV1AccountApiTradingStatusGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Account API Trading Status (USER_DATA)



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
    resp, r, err := api_client.WalletApi.SapiV1AccountApiTradingStatusGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountApiTradingStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AccountApiTradingStatusGet`: InlineResponse20036
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AccountApiTradingStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountApiTradingStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20036**](InlineResponse20036.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AccountDisableFastWithdrawSwitchPost

> map[string]interface{} SapiV1AccountDisableFastWithdrawSwitchPost(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Disable Fast Withdraw Switch (USER_DATA)



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
    resp, r, err := api_client.WalletApi.SapiV1AccountDisableFastWithdrawSwitchPost(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountDisableFastWithdrawSwitchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AccountDisableFastWithdrawSwitchPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AccountDisableFastWithdrawSwitchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountDisableFastWithdrawSwitchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## SapiV1AccountEnableFastWithdrawSwitchPost

> map[string]interface{} SapiV1AccountEnableFastWithdrawSwitchPost(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Enable Fast Withdraw Switch (USER_DATA)



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
    resp, r, err := api_client.WalletApi.SapiV1AccountEnableFastWithdrawSwitchPost(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountEnableFastWithdrawSwitchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AccountEnableFastWithdrawSwitchPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AccountEnableFastWithdrawSwitchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountEnableFastWithdrawSwitchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## SapiV1AccountSnapshotGet

> OneOfsnapshotSpotsnapshotMarginsnapshotFutures SapiV1AccountSnapshotGet(ctx).Type_(type_).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Daily Account Snapshot (USER_DATA)



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
    type_ := "type__example" // string | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(56) // int32 |  (optional) (default to 5)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AccountSnapshotGet(context.Background()).Type_(type_).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountSnapshotGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AccountSnapshotGet`: OneOfsnapshotSpotsnapshotMarginsnapshotFutures
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AccountSnapshotGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountSnapshotGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** |  | [default to 5]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfsnapshotSpotsnapshotMarginsnapshotFutures**](oneOf&lt;snapshotSpot,snapshotMargin,snapshotFutures&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AccountStatusGet

> InlineResponse20035 SapiV1AccountStatusGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Account Status (USER_DATA)



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
    resp, r, err := api_client.WalletApi.SapiV1AccountStatusGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AccountStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AccountStatusGet`: InlineResponse20035
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AccountStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AccountStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20035**](InlineResponse20035.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetAssetDetailGet

> InlineResponse20040 SapiV1AssetAssetDetailGet(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).RecvWindow(recvWindow).Execute()

Asset Detail (USER_DATA)



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
    asset := "BNB" // string |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetAssetDetailGet(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetAssetDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetAssetDetailGet`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetAssetDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetAssetDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetAssetDividendGet

> InlineResponse20039 SapiV1AssetAssetDividendGet(ctx).Limit(limit).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Asset Dividend Record (USER_DATA)



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
    limit := "limit_example" // string |  (default to "20")
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    asset := "BNB" // string |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetAssetDividendGet(context.Background()).Limit(limit).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetAssetDividendGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetAssetDividendGet`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetAssetDividendGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetAssetDividendGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** |  | [default to &quot;20&quot;]
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetDribbletGet

> InlineResponse20037 SapiV1AssetDribbletGet(ctx).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

DustLog(USER_DATA)



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
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetDribbletGet(context.Background()).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetDribbletGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetDribbletGet`: InlineResponse20037
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetDribbletGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetDribbletGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20037**](InlineResponse20037.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetDustPost

> InlineResponse20038 SapiV1AssetDustPost(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Dust Transfer (USER_DATA)



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
    asset := "asset=BTC&asset=USDT" // string | The asset being converted. For example, asset=BTC&asset=USDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetDustPost(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetDustPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetDustPost`: InlineResponse20038
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetDustPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetDustPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset being converted. For example, asset&#x3D;BTC&amp;asset&#x3D;USDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20038**](InlineResponse20038.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetGetFundingAssetPost

> []InlineResponse20044 SapiV1AssetGetFundingAssetPost(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

Funding Wallet (USER_DATA)



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
    asset := "BNB" // string |  (optional)
    needBtcValuation := "needBtcValuation_example" // string |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetGetFundingAssetPost(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetGetFundingAssetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetGetFundingAssetPost`: []InlineResponse20044
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetGetFundingAssetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetGetFundingAssetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **needBtcValuation** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20044**](InlineResponse20044.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetTradeFeeGet

> []InlineResponse20041 SapiV1AssetTradeFeeGet(ctx).Timestamp(timestamp).Signature(signature).Symbol(symbol).RecvWindow(recvWindow).Execute()

Trade Fee (USER_DATA)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetTradeFeeGet(context.Background()).Timestamp(timestamp).Signature(signature).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetTradeFeeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetTradeFeeGet`: []InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetTradeFeeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetTradeFeeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20041**](InlineResponse20041.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetTransferGet

> InlineResponse20042 SapiV1AssetTransferGet(ctx).Type_(type_).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()

Query User Universal Transfer History (USER_DATA)



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
    type_ := "MAIN_C2C" // string | Universal transfer type
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    fromSymbol := "BNBUSDT" // string | Must be sent when type are ISOLATEDMARGIN_MARGIN and ISOLATEDMARGIN_ISOLATEDMARGIN (optional)
    toSymbol := "BNBUSDT" // string | Must be sent when type are MARGIN_ISOLATEDMARGIN and ISOLATEDMARGIN_ISOLATEDMARGIN (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetTransferGet(context.Background()).Type_(type_).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetTransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetTransferGet`: InlineResponse20042
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetTransferGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetTransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Universal transfer type | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **fromSymbol** | **string** | Must be sent when type are ISOLATEDMARGIN_MARGIN and ISOLATEDMARGIN_ISOLATEDMARGIN | 
 **toSymbol** | **string** | Must be sent when type are MARGIN_ISOLATEDMARGIN and ISOLATEDMARGIN_ISOLATEDMARGIN | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20042**](InlineResponse20042.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1AssetTransferPost

> InlineResponse20043 SapiV1AssetTransferPost(ctx).Type_(type_).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()

User Universal Transfer (USER_DATA)



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
    type_ := "MAIN_C2C" // string | Universal transfer type
    asset := "BTC" // string | 
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    fromSymbol := "BNBUSDT" // string | Must be sent when type are ISOLATEDMARGIN_MARGIN and ISOLATEDMARGIN_ISOLATEDMARGIN (optional)
    toSymbol := "BNBUSDT" // string | Must be sent when type are MARGIN_ISOLATEDMARGIN and ISOLATEDMARGIN_ISOLATEDMARGIN (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1AssetTransferPost(context.Background()).Type_(type_).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1AssetTransferPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1AssetTransferPost`: InlineResponse20043
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1AssetTransferPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1AssetTransferPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Universal transfer type | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **fromSymbol** | **string** | Must be sent when type are ISOLATEDMARGIN_MARGIN and ISOLATEDMARGIN_ISOLATEDMARGIN | 
 **toSymbol** | **string** | Must be sent when type are MARGIN_ISOLATEDMARGIN and ISOLATEDMARGIN_ISOLATEDMARGIN | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20043**](InlineResponse20043.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1CapitalConfigGetallGet

> []InlineResponse20030 SapiV1CapitalConfigGetallGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

All Coins' Information (USER_DATA)



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
    resp, r, err := api_client.WalletApi.SapiV1CapitalConfigGetallGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalConfigGetallGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalConfigGetallGet`: []InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalConfigGetallGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalConfigGetallGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20030**](InlineResponse20030.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1CapitalDepositAddressGet

> InlineResponse20034 SapiV1CapitalDepositAddressGet(ctx).Coin(coin).Timestamp(timestamp).Signature(signature).Network(network).RecvWindow(recvWindow).Execute()

Deposit Address (supporting network) (USER_DATA)



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
    coin := "BNB" // string | Coin name
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    network := "ETH" // string |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalDepositAddressGet(context.Background()).Coin(coin).Timestamp(timestamp).Signature(signature).Network(network).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalDepositAddressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalDepositAddressGet`: InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalDepositAddressGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalDepositAddressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | Coin name | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **network** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20034**](InlineResponse20034.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1CapitalDepositHisrecGet

> []InlineResponse20032 SapiV1CapitalDepositHisrecGet(ctx).Coin(coin).Timestamp(timestamp).Signature(signature).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

Deposit History（supporting network） (USER_DATA)



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
    coin := "BNB" // string | Coin name
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    status := int32(56) // int32 | 0 -> pending\\ 6 -> credited but cannot withdraw\\ 1 -> success (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalDepositHisrecGet(context.Background()).Coin(coin).Timestamp(timestamp).Signature(signature).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalDepositHisrecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalDepositHisrecGet`: []InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalDepositHisrecGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalDepositHisrecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | Coin name | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **status** | **int32** | 0 -&gt; pending\\ 6 -&gt; credited but cannot withdraw\\ 1 -&gt; success | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **offset** | **int32** |  | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20032**](InlineResponse20032.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1CapitalWithdrawApplyPost

> InlineResponse20031 SapiV1CapitalWithdrawApplyPost(ctx).Coin(coin).Address(address).Amount(amount).Timestamp(timestamp).Signature(signature).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).WalletType(walletType).RecvWindow(recvWindow).Execute()

Withdraw (USER_DATA)



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
    coin := "BNB" // string | Coin name
    address := "address_example" // string | 
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    withdrawOrderId := "withdrawOrderId_example" // string | Client id for withdraw (optional)
    network := "network_example" // string | Get the value from `GET /sapi/v1/capital/config/getall` (optional)
    addressTag := "addressTag_example" // string | Secondary address identifier for coins like XRP,XMR etc. (optional)
    transactionFeeFlag := true // bool | When making internal transfer - `true` ->  returning the fee to the destination account; - `false` -> returning the fee back to the departure account. (optional) (default to false)
    name := "name_example" // string |  (optional)
    walletType := int32(56) // int32 | The wallet type for withdraw，0-Spot wallet, 1- Funding wallet. Default is Spot wallet (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalWithdrawApplyPost(context.Background()).Coin(coin).Address(address).Amount(amount).Timestamp(timestamp).Signature(signature).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).WalletType(walletType).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalWithdrawApplyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalWithdrawApplyPost`: InlineResponse20031
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalWithdrawApplyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalWithdrawApplyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | Coin name | 
 **address** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **withdrawOrderId** | **string** | Client id for withdraw | 
 **network** | **string** | Get the value from &#x60;GET /sapi/v1/capital/config/getall&#x60; | 
 **addressTag** | **string** | Secondary address identifier for coins like XRP,XMR etc. | 
 **transactionFeeFlag** | **bool** | When making internal transfer - &#x60;true&#x60; -&gt;  returning the fee to the destination account; - &#x60;false&#x60; -&gt; returning the fee back to the departure account. | [default to false]
 **name** | **string** |  | 
 **walletType** | **int32** | The wallet type for withdraw，0-Spot wallet, 1- Funding wallet. Default is Spot wallet | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20031**](InlineResponse20031.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1CapitalWithdrawHistoryGet

> []InlineResponse20033 SapiV1CapitalWithdrawHistoryGet(ctx).Coin(coin).Timestamp(timestamp).Signature(signature).WithdrawOrderId(withdrawOrderId).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

Withdraw History (supporting network) (USER_DATA)



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
    coin := "BNB" // string | Coin name
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    withdrawOrderId := "withdrawOrderId_example" // string |  (optional)
    status := int32(56) // int32 | 0:Email Sent 1:Cancelled 2:Awaiting Approval 3:Rejected 4:Processing 5:Failure 6:Completed (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    offset := int32(56) // int32 |  (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WalletApi.SapiV1CapitalWithdrawHistoryGet(context.Background()).Coin(coin).Timestamp(timestamp).Signature(signature).WithdrawOrderId(withdrawOrderId).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1CapitalWithdrawHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1CapitalWithdrawHistoryGet`: []InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1CapitalWithdrawHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1CapitalWithdrawHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | Coin name | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **withdrawOrderId** | **string** |  | 
 **status** | **int32** | 0:Email Sent 1:Cancelled 2:Awaiting Approval 3:Rejected 4:Processing 5:Failure 6:Completed | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **offset** | **int32** |  | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20033**](InlineResponse20033.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1SystemStatusGet

> InlineResponse20029 SapiV1SystemStatusGet(ctx).Execute()

System Status (System)



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
    resp, r, err := api_client.WalletApi.SapiV1SystemStatusGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.SapiV1SystemStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1SystemStatusGet`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.SapiV1SystemStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1SystemStatusGetRequest struct via the builder pattern


### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

