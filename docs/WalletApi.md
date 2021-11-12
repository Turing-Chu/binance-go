# \WalletApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyWithdraw**](WalletApi.md#ApplyWithdraw) | **Post** /sapi/v1/capital/withdraw/apply | Withdraw (USER_DATA)
[**DisableFastWithdraw**](WalletApi.md#DisableFastWithdraw) | **Post** /sapi/v1/account/disableFastWithdrawSwitch | Disable Fast Withdraw Switch (USER_DATA)
[**EnableFastWithdraw**](WalletApi.md#EnableFastWithdraw) | **Post** /sapi/v1/account/enableFastWithdrawSwitch | Enable Fast Withdraw Switch (USER_DATA)
[**GetAccountSnapshot**](WalletApi.md#GetAccountSnapshot) | **Get** /sapi/v1/accountSnapshot | Daily Account Snapshot (USER_DATA)
[**GetAccountStatus**](WalletApi.md#GetAccountStatus) | **Get** /sapi/v1/account/status | Account Status (USER_DATA)
[**GetAllTokenInfo**](WalletApi.md#GetAllTokenInfo) | **Get** /sapi/v1/capital/config/getall | All Coins&#39; Information (USER_DATA)
[**GetApiRestrictions**](WalletApi.md#GetApiRestrictions) | **Get** /sapi/v1/account/apiRestrictions | Get API Key Permission (USER_DATA)
[**GetApiTradingStatus**](WalletApi.md#GetApiTradingStatus) | **Get** /sapi/v1/account/apiTradingStatus | Account API Trading Status (USER_DATA)
[**GetAssetDetail**](WalletApi.md#GetAssetDetail) | **Get** /sapi/v1/asset/assetDetail | Asset Detail (USER_DATA)
[**GetAssetDividendHistory**](WalletApi.md#GetAssetDividendHistory) | **Get** /sapi/v1/asset/assetDividend | Asset Dividend Record (USER_DATA)
[**GetDepositAddress**](WalletApi.md#GetDepositAddress) | **Get** /sapi/v1/capital/deposit/address | Deposit Address (supporting network) (USER_DATA)
[**GetDepositRecords**](WalletApi.md#GetDepositRecords) | **Get** /sapi/v1/capital/deposit/hisrec | Deposit History（supporting network） (USER_DATA)
[**GetDustLog**](WalletApi.md#GetDustLog) | **Get** /sapi/v1/asset/dribblet | DustLog(USER_DATA)
[**GetFundingAsset**](WalletApi.md#GetFundingAsset) | **Post** /sapi/v1/asset/get-funding-asset | Funding Wallet (USER_DATA)
[**GetSystemStatus**](WalletApi.md#GetSystemStatus) | **Get** /sapi/v1/system/status | System Status (System)
[**GetTradeFee**](WalletApi.md#GetTradeFee) | **Get** /sapi/v1/asset/tradeFee | Trade Fee (USER_DATA)
[**GetTransferRecord**](WalletApi.md#GetTransferRecord) | **Get** /sapi/v1/asset/transfer | Query User Universal Transfer History (USER_DATA)
[**GetWithdrawRecords**](WalletApi.md#GetWithdrawRecords) | **Get** /sapi/v1/capital/withdraw/history | Withdraw History (supporting network) (USER_DATA)
[**Transfer**](WalletApi.md#Transfer) | **Post** /sapi/v1/asset/transfer | User Universal Transfer (USER_DATA)
[**TransferDustToBnb**](WalletApi.md#TransferDustToBnb) | **Post** /sapi/v1/asset/dust | Dust Transfer (USER_DATA)



## ApplyWithdraw

> InlineResponse20031 ApplyWithdraw(ctx).Coin(coin).Address(address).Amount(amount).Timestamp(timestamp).Signature(signature).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).WalletType(walletType).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.ApplyWithdraw(context.Background()).Coin(coin).Address(address).Amount(amount).Timestamp(timestamp).Signature(signature).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).WalletType(walletType).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.ApplyWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyWithdraw`: InlineResponse20031
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.ApplyWithdraw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyWithdrawRequest struct via the builder pattern


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


## DisableFastWithdraw

> map[string]interface{} DisableFastWithdraw(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.DisableFastWithdraw(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.DisableFastWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableFastWithdraw`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.DisableFastWithdraw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableFastWithdrawRequest struct via the builder pattern


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


## EnableFastWithdraw

> map[string]interface{} EnableFastWithdraw(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.EnableFastWithdraw(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.EnableFastWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableFastWithdraw`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.EnableFastWithdraw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableFastWithdrawRequest struct via the builder pattern


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


## GetAccountSnapshot

> OneOfsnapshotSpotsnapshotMarginsnapshotFutures GetAccountSnapshot(ctx).Type_(type_).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetAccountSnapshot(context.Background()).Type_(type_).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetAccountSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountSnapshot`: OneOfsnapshotSpotsnapshotMarginsnapshotFutures
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetAccountSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSnapshotRequest struct via the builder pattern


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


## GetAccountStatus

> InlineResponse20035 GetAccountStatus(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetAccountStatus(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetAccountStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountStatus`: InlineResponse20035
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetAccountStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountStatusRequest struct via the builder pattern


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


## GetAllTokenInfo

> []InlineResponse20030 GetAllTokenInfo(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetAllTokenInfo(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetAllTokenInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTokenInfo`: []InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetAllTokenInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTokenInfoRequest struct via the builder pattern


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


## GetApiRestrictions

> InlineResponse20045 GetApiRestrictions(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetApiRestrictions(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetApiRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiRestrictions`: InlineResponse20045
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetApiRestrictions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiRestrictionsRequest struct via the builder pattern


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


## GetApiTradingStatus

> InlineResponse20036 GetApiTradingStatus(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetApiTradingStatus(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetApiTradingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiTradingStatus`: InlineResponse20036
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetApiTradingStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTradingStatusRequest struct via the builder pattern


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


## GetAssetDetail

> InlineResponse20040 GetAssetDetail(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetAssetDetail(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetAssetDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDetail`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetAssetDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDetailRequest struct via the builder pattern


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


## GetAssetDividendHistory

> InlineResponse20039 GetAssetDividendHistory(ctx).Limit(limit).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetAssetDividendHistory(context.Background()).Limit(limit).Timestamp(timestamp).Signature(signature).Asset(asset).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetAssetDividendHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetDividendHistory`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetAssetDividendHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDividendHistoryRequest struct via the builder pattern


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


## GetDepositAddress

> InlineResponse20034 GetDepositAddress(ctx).Coin(coin).Timestamp(timestamp).Signature(signature).Network(network).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetDepositAddress(context.Background()).Coin(coin).Timestamp(timestamp).Signature(signature).Network(network).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetDepositAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDepositAddress`: InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetDepositAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepositAddressRequest struct via the builder pattern


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


## GetDepositRecords

> []InlineResponse20032 GetDepositRecords(ctx).Coin(coin).Timestamp(timestamp).Signature(signature).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetDepositRecords(context.Background()).Coin(coin).Timestamp(timestamp).Signature(signature).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetDepositRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDepositRecords`: []InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetDepositRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepositRecordsRequest struct via the builder pattern


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


## GetDustLog

> InlineResponse20037 GetDustLog(ctx).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetDustLog(context.Background()).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetDustLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDustLog`: InlineResponse20037
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetDustLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDustLogRequest struct via the builder pattern


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


## GetFundingAsset

> []InlineResponse20044 GetFundingAsset(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetFundingAsset(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetFundingAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingAsset`: []InlineResponse20044
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetFundingAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingAssetRequest struct via the builder pattern


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


## GetSystemStatus

> InlineResponse20029 GetSystemStatus(ctx).Execute()

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
    resp, r, err := api_client.WalletApi.GetSystemStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetSystemStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemStatus`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetSystemStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemStatusRequest struct via the builder pattern


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


## GetTradeFee

> []InlineResponse20041 GetTradeFee(ctx).Timestamp(timestamp).Signature(signature).Symbol(symbol).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetTradeFee(context.Background()).Timestamp(timestamp).Signature(signature).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetTradeFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeFee`: []InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetTradeFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeFeeRequest struct via the builder pattern


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


## GetTransferRecord

> InlineResponse20042 GetTransferRecord(ctx).Type_(type_).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetTransferRecord(context.Background()).Type_(type_).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetTransferRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransferRecord`: InlineResponse20042
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetTransferRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferRecordRequest struct via the builder pattern


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


## GetWithdrawRecords

> []InlineResponse20033 GetWithdrawRecords(ctx).Coin(coin).Timestamp(timestamp).Signature(signature).WithdrawOrderId(withdrawOrderId).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.GetWithdrawRecords(context.Background()).Coin(coin).Timestamp(timestamp).Signature(signature).WithdrawOrderId(withdrawOrderId).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.GetWithdrawRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWithdrawRecords`: []InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.GetWithdrawRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawRecordsRequest struct via the builder pattern


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


## Transfer

> InlineResponse20043 Transfer(ctx).Type_(type_).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.Transfer(context.Background()).Type_(type_).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.Transfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Transfer`: InlineResponse20043
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.Transfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferRequest struct via the builder pattern


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


## TransferDustToBnb

> InlineResponse20038 TransferDustToBnb(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.WalletApi.TransferDustToBnb(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.TransferDustToBnb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferDustToBnb`: InlineResponse20038
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.TransferDustToBnb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferDustToBnbRequest struct via the builder pattern


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

