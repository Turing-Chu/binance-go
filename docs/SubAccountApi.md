# \SubAccountApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountCreateVirtualSubAccount**](SubAccountApi.md#SubaccountCreateVirtualSubAccount) | **Post** /sapi/v1/sub-account/virtualSubAccount | Create a Virtual Sub-account(For Master Account)
[**SubaccountDeposit**](SubAccountApi.md#SubaccountDeposit) | **Post** /sapi/v1/managed-subaccount/deposit | Deposit assets into the managed sub-account（For Investor Master Account）
[**SubaccountEnableFutures**](SubAccountApi.md#SubaccountEnableFutures) | **Post** /sapi/v1/sub-account/futures/enable | Enable Futures for Sub-account (For Master Account)
[**SubaccountEnableLeverageToken**](SubAccountApi.md#SubaccountEnableLeverageToken) | **Post** /sapi/v1/sub-account/blvt/enable | Enable Leverage Token for Sub-account (For Master Account)
[**SubaccountEnableMargin**](SubAccountApi.md#SubaccountEnableMargin) | **Post** /sapi/v1/sub-account/margin/enable | Enable Margin for Sub-account (For Master Account)
[**SubaccountFuturesInternalTransfer**](SubAccountApi.md#SubaccountFuturesInternalTransfer) | **Post** /sapi/v1/sub-account/futures/internalTransfer | Sub-account Futures Asset Transfer (For Master Account)
[**SubaccountFuturesTransfer**](SubAccountApi.md#SubaccountFuturesTransfer) | **Post** /sapi/v1/sub-account/futures/transfer | Transfer for Sub-account (For Master Account)
[**SubaccountGetAssetDetail**](SubAccountApi.md#SubaccountGetAssetDetail) | **Get** /sapi/v1/managed-subaccount/asset | Managed sub-account asset details（For Investor Master Account)
[**SubaccountGetAssets**](SubAccountApi.md#SubaccountGetAssets) | **Get** /sapi/v3/sub-account/assets | Sub-account Assets (For Master Account)
[**SubaccountGetDepositAddress**](SubAccountApi.md#SubaccountGetDepositAddress) | **Get** /sapi/v1/capital/deposit/subAddress | Sub-account Spot Assets Summary (For Master Account)
[**SubaccountGetDepositRecords**](SubAccountApi.md#SubaccountGetDepositRecords) | **Get** /sapi/v1/capital/deposit/subHisrec | Sub-account Deposit History (For Master Account)
[**SubaccountGetDetailFuturesAccount**](SubAccountApi.md#SubaccountGetDetailFuturesAccount) | **Get** /sapi/v2/sub-account/futures/account | Detail on Sub-account&#39;s Futures Account V2 (For Master Account)
[**SubaccountGetFuturesAccountDetail**](SubAccountApi.md#SubaccountGetFuturesAccountDetail) | **Get** /sapi/v1/sub-account/futures/account | Detail on Sub-account&#39;s Futures Account (For Master Account)
[**SubaccountGetFuturesAccountSummary**](SubAccountApi.md#SubaccountGetFuturesAccountSummary) | **Get** /sapi/v1/sub-account/futures/accountSummary | Summary of Sub-account&#39;s Futures Account (For Master Account)
[**SubaccountGetFuturesAccountSummaryV2**](SubAccountApi.md#SubaccountGetFuturesAccountSummaryV2) | **Get** /sapi/v2/sub-account/futures/accountSummary | Summary of Sub-account&#39;s Futures Account V2 (For Master Account)
[**SubaccountGetFuturesPositionRisk**](SubAccountApi.md#SubaccountGetFuturesPositionRisk) | **Get** /sapi/v1/sub-account/futures/positionRisk | Futures Position-Risk of Sub-account (For Master Account)
[**SubaccountGetFuturesPositionRiskV2**](SubAccountApi.md#SubaccountGetFuturesPositionRiskV2) | **Get** /sapi/v2/sub-account/futures/positionRisk | Futures Position-Risk of Sub-account V2 (For Master Account)
[**SubaccountGetFuturesTransferRecord**](SubAccountApi.md#SubaccountGetFuturesTransferRecord) | **Get** /sapi/v1/sub-account/futures/internalTransfer | Sub-account Futures Asset Transfer History (For Master Account)
[**SubaccountGetMarginAccountDetail**](SubAccountApi.md#SubaccountGetMarginAccountDetail) | **Get** /sapi/v1/sub-account/margin/account | Detail on Sub-account&#39;s Margin Account (For Master Account)
[**SubaccountGetMarginAccountSummary**](SubAccountApi.md#SubaccountGetMarginAccountSummary) | **Get** /sapi/v1/sub-account/margin/accountSummary | Summary of Sub-account&#39;s Margin Account (For Master Account)
[**SubaccountGetSpotSummary**](SubAccountApi.md#SubaccountGetSpotSummary) | **Get** /sapi/v1/sub-account/spotSummary | Sub-account Spot Assets Summary (For Master Account)
[**SubaccountGetSpotTransferHistory**](SubAccountApi.md#SubaccountGetSpotTransferHistory) | **Get** /sapi/v1/sub-account/sub/transfer/history | Sub-account Spot Asset Transfer History (For Master Account)
[**SubaccountGetTransferHistory**](SubAccountApi.md#SubaccountGetTransferHistory) | **Get** /sapi/v1/sub-account/transfer/subUserHistory | Sub-account Transfer History (For Sub-account)
[**SubaccountGetUniversalTransferRecord**](SubAccountApi.md#SubaccountGetUniversalTransferRecord) | **Get** /sapi/v1/sub-account/universalTransfer | Universal Transfer History (For Master Account)
[**SubaccountList**](SubAccountApi.md#SubaccountList) | **Get** /sapi/v1/sub-account/list | Query Sub-account List (For Master Account)
[**SubaccountMarginTransfer**](SubAccountApi.md#SubaccountMarginTransfer) | **Post** /sapi/v1/sub-account/margin/transfer | Margin Transfer for Sub-account (For Master Account)
[**SubaccountStatus**](SubAccountApi.md#SubaccountStatus) | **Get** /sapi/v1/sub-account/status | Sub-account&#39;s Status on Margin/Futures (For Master Account)
[**SubaccountTransferBetweenSub**](SubAccountApi.md#SubaccountTransferBetweenSub) | **Post** /sapi/v1/sub-account/transfer/subToSub | Transfer to Sub-account of Same Master (For Sub-account)
[**SubaccountTransferToMaster**](SubAccountApi.md#SubaccountTransferToMaster) | **Post** /sapi/v1/sub-account/transfer/subToMaster | Transfer to Master (For Sub-account)
[**SubaccountUniversalTransfer**](SubAccountApi.md#SubaccountUniversalTransfer) | **Post** /sapi/v1/sub-account/universalTransfer | Universal Transfer (For Master Account)
[**SubaccountWithdrawlAssets**](SubAccountApi.md#SubaccountWithdrawlAssets) | **Post** /sapi/v1/managed-subaccount/withdraw | Withdrawl assets from the managed sub-account（For Investor Master Account)



## SubaccountCreateVirtualSubAccount

> InlineResponse20046 SubaccountCreateVirtualSubAccount(ctx).SubAccountString(subAccountString).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Create a Virtual Sub-account(For Master Account)



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
    subAccountString := "subAccountString_example" // string | Please input a string. We will create a virtual email using that string for you to register
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountCreateVirtualSubAccount(context.Background()).SubAccountString(subAccountString).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountCreateVirtualSubAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountCreateVirtualSubAccount`: InlineResponse20046
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountCreateVirtualSubAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateVirtualSubAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subAccountString** | **string** | Please input a string. We will create a virtual email using that string for you to register | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20046**](InlineResponse20046.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountDeposit

> InlineResponse20068 SubaccountDeposit(ctx).ToEmail(toEmail).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Deposit assets into the managed sub-account（For Investor Master Account）



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
    toEmail := "toEmail_example" // string | Recipient email
    asset := "BTC" // string | 
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountDeposit(context.Background()).ToEmail(toEmail).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountDeposit`: InlineResponse20068
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toEmail** | **string** | Recipient email | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20068**](InlineResponse20068.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountEnableFutures

> InlineResponse20059 SubaccountEnableFutures(ctx).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Enable Futures for Sub-account (For Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountEnableFutures(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountEnableFutures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountEnableFutures`: InlineResponse20059
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountEnableFutures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountEnableFuturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20059**](InlineResponse20059.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountEnableLeverageToken

> InlineResponse20067 SubaccountEnableLeverageToken(ctx).Email(email).EnableBlvt(enableBlvt).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Enable Leverage Token for Sub-account (For Master Account)



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
    email := "email_example" // string | Sub-account email
    enableBlvt := true // bool | Only true for now
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountEnableLeverageToken(context.Background()).Email(email).EnableBlvt(enableBlvt).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountEnableLeverageToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountEnableLeverageToken`: InlineResponse20067
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountEnableLeverageToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountEnableLeverageTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **enableBlvt** | **bool** | Only true for now | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20067**](InlineResponse20067.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountEnableMargin

> InlineResponse20056 SubaccountEnableMargin(ctx).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Enable Margin for Sub-account (For Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountEnableMargin(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountEnableMargin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountEnableMargin`: InlineResponse20056
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountEnableMargin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountEnableMarginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20056**](InlineResponse20056.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountFuturesInternalTransfer

> InlineResponse20050 SubaccountFuturesInternalTransfer(ctx).FromEmail(fromEmail).ToEmail(toEmail).FuturesType(futuresType).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Sub-account Futures Asset Transfer (For Master Account)



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
    fromEmail := "fromEmail_example" // string | Sender email
    toEmail := "toEmail_example" // string | Recipient email
    futuresType := int32(2) // int32 | 1:USDT-margined Futures,2: Coin-margined Futures
    asset := "BTC" // string | 
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountFuturesInternalTransfer(context.Background()).FromEmail(fromEmail).ToEmail(toEmail).FuturesType(futuresType).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountFuturesInternalTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountFuturesInternalTransfer`: InlineResponse20050
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountFuturesInternalTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountFuturesInternalTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromEmail** | **string** | Sender email | 
 **toEmail** | **string** | Recipient email | 
 **futuresType** | **int32** | 1:USDT-margined Futures,2: Coin-margined Futures | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20050**](InlineResponse20050.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountFuturesTransfer

> InlineResponse20063 SubaccountFuturesTransfer(ctx).Email(email).Asset(asset).Amount(amount).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Transfer for Sub-account (For Master Account)



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
    email := "email_example" // string | Sub-account email
    asset := "BTC" // string | 
    amount := float64(1.01) // float64 | 
    type_ := int32(56) // int32 | 1: transfer from subaccount's spot account to its USDT-margined futures account  2: transfer from subaccount's USDT-margined futures account to its spot account  3: transfer from subaccount's spot account to its COIN-margined futures account  4:transfer from subaccount's COIN-margined futures account to its spot account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountFuturesTransfer(context.Background()).Email(email).Asset(asset).Amount(amount).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountFuturesTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountFuturesTransfer`: InlineResponse20063
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountFuturesTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountFuturesTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **type_** | **int32** | 1: transfer from subaccount&#39;s spot account to its USDT-margined futures account  2: transfer from subaccount&#39;s USDT-margined futures account to its spot account  3: transfer from subaccount&#39;s spot account to its COIN-margined futures account  4:transfer from subaccount&#39;s COIN-margined futures account to its spot account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20063**](InlineResponse20063.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetAssetDetail

> []InlineResponse20069 SubaccountGetAssetDetail(ctx).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Managed sub-account asset details（For Investor Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetAssetDetail(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetAssetDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetAssetDetail`: []InlineResponse20069
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetAssetDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetAssetDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20069**](InlineResponse20069.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetAssets

> InlineResponse20051 SubaccountGetAssets(ctx).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Sub-account Assets (For Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetAssets(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetAssets`: InlineResponse20051
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20051**](InlineResponse20051.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetDepositAddress

> InlineResponse20053 SubaccountGetDepositAddress(ctx).Email(email).Coin(coin).Timestamp(timestamp).Signature(signature).Network(network).RecvWindow(recvWindow).Execute()

Sub-account Spot Assets Summary (For Master Account)



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
    email := "email_example" // string | Sub-account email
    coin := "BNB" // string | Coin name
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    network := "network_example" // string |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetDepositAddress(context.Background()).Email(email).Coin(coin).Timestamp(timestamp).Signature(signature).Network(network).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetDepositAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetDepositAddress`: InlineResponse20053
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetDepositAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetDepositAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **coin** | **string** | Coin name | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **network** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20053**](InlineResponse20053.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetDepositRecords

> []InlineResponse20054 SubaccountGetDepositRecords(ctx).Email(email).Timestamp(timestamp).Signature(signature).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).Execute()

Sub-account Deposit History (For Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    coin := "BNB" // string | Coin name (optional)
    status := int32(56) // int32 | 0(0:pending,6: credited but cannot withdraw, 1:success) (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int64(789) // int64 |  (optional)
    offset := int32(56) // int32 |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetDepositRecords(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetDepositRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetDepositRecords`: []InlineResponse20054
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetDepositRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetDepositRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **coin** | **string** | Coin name | 
 **status** | **int32** | 0(0:pending,6: credited but cannot withdraw, 1:success) | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int64** |  | 
 **offset** | **int32** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20054**](InlineResponse20054.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetDetailFuturesAccount

> OneOfsubAccountUSDTFuturesDetailssubAccountCOINFuturesDetails SubaccountGetDetailFuturesAccount(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Detail on Sub-account's Futures Account V2 (For Master Account)



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
    email := "email_example" // string | Sub-account email
    futuresType := int32(56) // int32 | 1:USDT Margined Futures  2:COIN Margined Futures
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetDetailFuturesAccount(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetDetailFuturesAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetDetailFuturesAccount`: OneOfsubAccountUSDTFuturesDetailssubAccountCOINFuturesDetails
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetDetailFuturesAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetDetailFuturesAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **futuresType** | **int32** | 1:USDT Margined Futures  2:COIN Margined Futures | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfsubAccountUSDTFuturesDetailssubAccountCOINFuturesDetails**](oneOf&lt;subAccountUSDTFuturesDetails,subAccountCOINFuturesDetails&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetFuturesAccountDetail

> InlineResponse20060 SubaccountGetFuturesAccountDetail(ctx).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Detail on Sub-account's Futures Account (For Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetFuturesAccountDetail(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetFuturesAccountDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetFuturesAccountDetail`: InlineResponse20060
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetFuturesAccountDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetFuturesAccountDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20060**](InlineResponse20060.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetFuturesAccountSummary

> InlineResponse20061 SubaccountGetFuturesAccountSummary(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Summary of Sub-account's Futures Account (For Master Account)



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
    resp, r, err := api_client.SubAccountApi.SubaccountGetFuturesAccountSummary(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetFuturesAccountSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetFuturesAccountSummary`: InlineResponse20061
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetFuturesAccountSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetFuturesAccountSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20061**](InlineResponse20061.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetFuturesAccountSummaryV2

> OneOfsubAccountUSDTFuturesSummarysubAccountCOINFuturesSummary SubaccountGetFuturesAccountSummaryV2(ctx).FuturesType(futuresType).Timestamp(timestamp).Signature(signature).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Summary of Sub-account's Futures Account V2 (For Master Account)



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
    futuresType := int32(56) // int32 | 1:USDT Margined Futures  2:COIN Margined Futures
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    page := int32(1) // int32 | Default 1 (optional)
    limit := "limit_example" // string | Default 10, Max 20 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetFuturesAccountSummaryV2(context.Background()).FuturesType(futuresType).Timestamp(timestamp).Signature(signature).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetFuturesAccountSummaryV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetFuturesAccountSummaryV2`: OneOfsubAccountUSDTFuturesSummarysubAccountCOINFuturesSummary
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetFuturesAccountSummaryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetFuturesAccountSummaryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **futuresType** | **int32** | 1:USDT Margined Futures  2:COIN Margined Futures | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **page** | **int32** | Default 1 | 
 **limit** | **string** | Default 10, Max 20 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfsubAccountUSDTFuturesSummarysubAccountCOINFuturesSummary**](oneOf&lt;subAccountUSDTFuturesSummary,subAccountCOINFuturesSummary&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetFuturesPositionRisk

> []InlineResponse20062 SubaccountGetFuturesPositionRisk(ctx).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Futures Position-Risk of Sub-account (For Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetFuturesPositionRisk(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetFuturesPositionRisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetFuturesPositionRisk`: []InlineResponse20062
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetFuturesPositionRisk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetFuturesPositionRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20062**](InlineResponse20062.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetFuturesPositionRiskV2

> OneOfsubAccountUSDTFuturesPositionRisksubAccountCOINFuturesPositionRisk SubaccountGetFuturesPositionRiskV2(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Futures Position-Risk of Sub-account V2 (For Master Account)



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
    email := "email_example" // string | Sub-account email
    futuresType := int32(56) // int32 | 1:USDT Margined Futures  2:COIN Margined Futures
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetFuturesPositionRiskV2(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetFuturesPositionRiskV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetFuturesPositionRiskV2`: OneOfsubAccountUSDTFuturesPositionRisksubAccountCOINFuturesPositionRisk
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetFuturesPositionRiskV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetFuturesPositionRiskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **futuresType** | **int32** | 1:USDT Margined Futures  2:COIN Margined Futures | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfsubAccountUSDTFuturesPositionRisksubAccountCOINFuturesPositionRisk**](oneOf&lt;subAccountUSDTFuturesPositionRisk,subAccountCOINFuturesPositionRisk&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetFuturesTransferRecord

> InlineResponse20049 SubaccountGetFuturesTransferRecord(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Sub-account Futures Asset Transfer History (For Master Account)



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
    email := "email_example" // string | Sub-account email
    futuresType := int32(2) // int32 | 1:USDT-margined Futures, 2: Coin-margined Futures
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    page := int32(1) // int32 | Default 1 (optional)
    limit := int32(56) // int32 | Default value: 50, Max value: 500 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetFuturesTransferRecord(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetFuturesTransferRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetFuturesTransferRecord`: InlineResponse20049
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetFuturesTransferRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetFuturesTransferRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **futuresType** | **int32** | 1:USDT-margined Futures, 2: Coin-margined Futures | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **page** | **int32** | Default 1 | 
 **limit** | **int32** | Default value: 50, Max value: 500 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20049**](InlineResponse20049.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetMarginAccountDetail

> InlineResponse20057 SubaccountGetMarginAccountDetail(ctx).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Detail on Sub-account's Margin Account (For Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetMarginAccountDetail(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetMarginAccountDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetMarginAccountDetail`: InlineResponse20057
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetMarginAccountDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetMarginAccountDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20057**](InlineResponse20057.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetMarginAccountSummary

> InlineResponse20058 SubaccountGetMarginAccountSummary(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Summary of Sub-account's Margin Account (For Master Account)



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
    resp, r, err := api_client.SubAccountApi.SubaccountGetMarginAccountSummary(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetMarginAccountSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetMarginAccountSummary`: InlineResponse20058
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetMarginAccountSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetMarginAccountSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20058**](InlineResponse20058.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSpotSummary

> InlineResponse20052 SubaccountGetSpotSummary(ctx).Email(email).Timestamp(timestamp).Signature(signature).Page(page).Size(size).RecvWindow(recvWindow).Execute()

Sub-account Spot Assets Summary (For Master Account)



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
    email := "email_example" // string | Sub-account email
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    page := int32(1) // int32 | Default 1 (optional)
    size := int32(56) // int32 | Default:10 Max:20 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetSpotSummary(context.Background()).Email(email).Timestamp(timestamp).Signature(signature).Page(page).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetSpotSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetSpotSummary`: InlineResponse20052
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetSpotSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSpotSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **page** | **int32** | Default 1 | 
 **size** | **int32** | Default:10 Max:20 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20052**](InlineResponse20052.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSpotTransferHistory

> []InlineResponse20048 SubaccountGetSpotTransferHistory(ctx).Timestamp(timestamp).Signature(signature).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Sub-account Spot Asset Transfer History (For Master Account)



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
    fromEmail := "fromEmail_example" // string | Sub-account email (optional)
    toEmail := "toEmail_example" // string | Sub-account email (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    page := int32(1) // int32 | Default 1 (optional)
    limit := int32(1) // int32 | Default 1 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetSpotTransferHistory(context.Background()).Timestamp(timestamp).Signature(signature).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetSpotTransferHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetSpotTransferHistory`: []InlineResponse20048
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetSpotTransferHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSpotTransferHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **fromEmail** | **string** | Sub-account email | 
 **toEmail** | **string** | Sub-account email | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **page** | **int32** | Default 1 | 
 **limit** | **int32** | Default 1 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20048**](InlineResponse20048.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetTransferHistory

> []InlineResponse20064 SubaccountGetTransferHistory(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Sub-account Transfer History (For Sub-account)



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
    type_ := int32(56) // int32 | 1: transfer in  2:  transfer out (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetTransferHistory(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetTransferHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetTransferHistory`: []InlineResponse20064
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetTransferHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetTransferHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **type_** | **int32** | 1: transfer in  2:  transfer out | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20064**](InlineResponse20064.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetUniversalTransferRecord

> []InlineResponse20065 SubaccountGetUniversalTransferRecord(ctx).Timestamp(timestamp).Signature(signature).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Universal Transfer History (For Master Account)



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
    fromEmail := "fromEmail_example" // string | Sub-account email (optional)
    toEmail := "toEmail_example" // string | Sub-account email (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    page := int32(1) // int32 | Default 1 (optional)
    limit := "limit_example" // string | Default 500, Max 500 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountGetUniversalTransferRecord(context.Background()).Timestamp(timestamp).Signature(signature).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountGetUniversalTransferRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountGetUniversalTransferRecord`: []InlineResponse20065
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountGetUniversalTransferRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetUniversalTransferRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **fromEmail** | **string** | Sub-account email | 
 **toEmail** | **string** | Sub-account email | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **page** | **int32** | Default 1 | 
 **limit** | **string** | Default 500, Max 500 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20065**](InlineResponse20065.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountList

> InlineResponse20047 SubaccountList(ctx).Timestamp(timestamp).Signature(signature).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account List (For Master Account)



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
    email := "email_example" // string | Sub-account email (optional)
    isFreeze := "isFreeze_example" // string |  (optional)
    page := int32(1) // int32 | Default 1 (optional)
    limit := int32(1) // int32 | Default 1; max 200 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountList(context.Background()).Timestamp(timestamp).Signature(signature).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountList`: InlineResponse20047
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **email** | **string** | Sub-account email | 
 **isFreeze** | **string** |  | 
 **page** | **int32** | Default 1 | 
 **limit** | **int32** | Default 1; max 200 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20047**](InlineResponse20047.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountMarginTransfer

> InlineResponse20063 SubaccountMarginTransfer(ctx).Email(email).Asset(asset).Amount(amount).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Margin Transfer for Sub-account (For Master Account)



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
    email := "email_example" // string | Sub-account email
    asset := "BTC" // string | 
    amount := float64(1.01) // float64 | 
    type_ := int32(56) // int32 | 1: transfer from subaccount's spot account to margin account  2: transfer from subaccount's margin account to its spot account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountMarginTransfer(context.Background()).Email(email).Asset(asset).Amount(amount).Type_(type_).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountMarginTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountMarginTransfer`: InlineResponse20063
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountMarginTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountMarginTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **type_** | **int32** | 1: transfer from subaccount&#39;s spot account to margin account  2: transfer from subaccount&#39;s margin account to its spot account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20063**](InlineResponse20063.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountStatus

> []InlineResponse20055 SubaccountStatus(ctx).Timestamp(timestamp).Signature(signature).Email(email).RecvWindow(recvWindow).Execute()

Sub-account's Status on Margin/Futures (For Master Account)



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
    email := "email_example" // string | Sub-account email (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountStatus(context.Background()).Timestamp(timestamp).Signature(signature).Email(email).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountStatus`: []InlineResponse20055
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **email** | **string** | Sub-account email | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20055**](InlineResponse20055.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountTransferBetweenSub

> InlineResponse20063 SubaccountTransferBetweenSub(ctx).ToEmail(toEmail).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Transfer to Sub-account of Same Master (For Sub-account)



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
    toEmail := "toEmail_example" // string | Recipient email
    asset := "BTC" // string | 
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountTransferBetweenSub(context.Background()).ToEmail(toEmail).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountTransferBetweenSub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountTransferBetweenSub`: InlineResponse20063
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountTransferBetweenSub`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountTransferBetweenSubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toEmail** | **string** | Recipient email | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20063**](InlineResponse20063.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountTransferToMaster

> InlineResponse20063 SubaccountTransferToMaster(ctx).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Transfer to Master (For Sub-account)



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
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountTransferToMaster(context.Background()).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountTransferToMaster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountTransferToMaster`: InlineResponse20063
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountTransferToMaster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountTransferToMasterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20063**](InlineResponse20063.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountUniversalTransfer

> InlineResponse20066 SubaccountUniversalTransfer(ctx).FromAccountType(fromAccountType).ToAccountType(toAccountType).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).FromEmail(fromEmail).ToEmail(toEmail).RecvWindow(recvWindow).Execute()

Universal Transfer (For Master Account)



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
    fromAccountType := "fromAccountType_example" // string | 
    toAccountType := "toAccountType_example" // string | 
    asset := "BTC" // string | 
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    fromEmail := "fromEmail_example" // string | Sub-account email (optional)
    toEmail := "toEmail_example" // string | Sub-account email (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountUniversalTransfer(context.Background()).FromAccountType(fromAccountType).ToAccountType(toAccountType).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).FromEmail(fromEmail).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountUniversalTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountUniversalTransfer`: InlineResponse20066
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountUniversalTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountUniversalTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAccountType** | **string** |  | 
 **toAccountType** | **string** |  | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **fromEmail** | **string** | Sub-account email | 
 **toEmail** | **string** | Sub-account email | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20066**](InlineResponse20066.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountWithdrawlAssets

> InlineResponse20068 SubaccountWithdrawlAssets(ctx).FromEmail(fromEmail).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).TransferDate(transferDate).RecvWindow(recvWindow).Execute()

Withdrawl assets from the managed sub-account（For Investor Master Account)



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
    fromEmail := "fromEmail_example" // string | Sender email
    asset := "BTC" // string | 
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    transferDate := int64(789) // int64 | Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubAccountApi.SubaccountWithdrawlAssets(context.Background()).FromEmail(fromEmail).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).TransferDate(transferDate).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubAccountApi.SubaccountWithdrawlAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubaccountWithdrawlAssets`: InlineResponse20068
    fmt.Fprintf(os.Stdout, "Response from `SubAccountApi.SubaccountWithdrawlAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountWithdrawlAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromEmail** | **string** | Sender email | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **transferDate** | **int64** | Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20068**](InlineResponse20068.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

