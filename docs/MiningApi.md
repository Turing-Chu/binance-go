# \MiningApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1MiningHashTransferConfigCancelPost**](MiningApi.md#SapiV1MiningHashTransferConfigCancelPost) | **Post** /sapi/v1/mining/hash-transfer/config/cancel | Cancel Hashrate Resale configuration (USER_DATA)
[**SapiV1MiningHashTransferConfigDetailsListGet**](MiningApi.md#SapiV1MiningHashTransferConfigDetailsListGet) | **Get** /sapi/v1/mining/hash-transfer/config/details/list | Hashrate Resale List (USER_DATA)
[**SapiV1MiningHashTransferConfigPost**](MiningApi.md#SapiV1MiningHashTransferConfigPost) | **Post** /sapi/v1/mining/hash-transfer/config | Hashrate Resale Request (USER_DATA)
[**SapiV1MiningHashTransferProfitDetailsGet**](MiningApi.md#SapiV1MiningHashTransferProfitDetailsGet) | **Get** /sapi/v1/mining/hash-transfer/profit/details | Hashrate Resale Details (USER_DATA)
[**SapiV1MiningPaymentListGet**](MiningApi.md#SapiV1MiningPaymentListGet) | **Get** /sapi/v1/mining/payment/list | Earnings List (USER_DATA)
[**SapiV1MiningPaymentOtherGet**](MiningApi.md#SapiV1MiningPaymentOtherGet) | **Get** /sapi/v1/mining/payment/other | Extra Bonus List (USER_DATA)
[**SapiV1MiningPubAlgoListGet**](MiningApi.md#SapiV1MiningPubAlgoListGet) | **Get** /sapi/v1/mining/pub/algoList | Acquiring Algorithm (MARKET_DATA)
[**SapiV1MiningPubCoinListGet**](MiningApi.md#SapiV1MiningPubCoinListGet) | **Get** /sapi/v1/mining/pub/coinList | Acquiring CoinName (MARKET_DATA)
[**SapiV1MiningStatisticsUserListGet**](MiningApi.md#SapiV1MiningStatisticsUserListGet) | **Get** /sapi/v1/mining/statistics/user/list | Account List (USER_DATA)
[**SapiV1MiningStatisticsUserStatusGet**](MiningApi.md#SapiV1MiningStatisticsUserStatusGet) | **Get** /sapi/v1/mining/statistics/user/status | Statistic List (USER_DATA)
[**SapiV1MiningWorkerDetailGet**](MiningApi.md#SapiV1MiningWorkerDetailGet) | **Get** /sapi/v1/mining/worker/detail | Request for Detail Miner List (USER_DATA)
[**SapiV1MiningWorkerListGet**](MiningApi.md#SapiV1MiningWorkerListGet) | **Get** /sapi/v1/mining/worker/list | Request for Miner List (USER_DATA)



## SapiV1MiningHashTransferConfigCancelPost

> InlineResponse20094 SapiV1MiningHashTransferConfigCancelPost(ctx).ConfigId(configId).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Cancel Hashrate Resale configuration (USER_DATA)



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
    configId := "configId_example" // string | Mining ID
    userName := "userName_example" // string | Mining Account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningHashTransferConfigCancelPost(context.Background()).ConfigId(configId).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningHashTransferConfigCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningHashTransferConfigCancelPost`: InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningHashTransferConfigCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningHashTransferConfigCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configId** | **string** | Mining ID | 
 **userName** | **string** | Mining Account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20094**](InlineResponse20094.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningHashTransferConfigDetailsListGet

> InlineResponse20091 SapiV1MiningHashTransferConfigDetailsListGet(ctx).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Hashrate Resale List (USER_DATA)



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
    pageIndex := int32(56) // int32 | Page number, default is first page, start form 1 (optional)
    pageSize := "pageSize_example" // string | Number of pages, minimum 10, maximum 200 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningHashTransferConfigDetailsListGet(context.Background()).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningHashTransferConfigDetailsListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningHashTransferConfigDetailsListGet`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningHashTransferConfigDetailsListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningHashTransferConfigDetailsListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **pageIndex** | **int32** | Page number, default is first page, start form 1 | 
 **pageSize** | **string** | Number of pages, minimum 10, maximum 200 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningHashTransferConfigPost

> InlineResponse20093 SapiV1MiningHashTransferConfigPost(ctx).UserName(userName).Algo(algo).ToPoolUser(toPoolUser).HashRate(hashRate).Timestamp(timestamp).Signature(signature).StartDate(startDate).EndDate(endDate).RecvWindow(recvWindow).Execute()

Hashrate Resale Request (USER_DATA)



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
    userName := "userName_example" // string | Mining Account
    algo := "algo_example" // string | Algorithm(sha256)
    toPoolUser := "toPoolUser_example" // string | Mining Account
    hashRate := "hashRate_example" // string | Resale hashrate h/s must be transferred (BTC is greater than 500000000000 ETH is greater than 500000)
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    startDate := "startDate_example" // string | Search date, millisecond timestamp, while empty query all (optional)
    endDate := "endDate_example" // string | Search date, millisecond timestamp, while empty query all (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningHashTransferConfigPost(context.Background()).UserName(userName).Algo(algo).ToPoolUser(toPoolUser).HashRate(hashRate).Timestamp(timestamp).Signature(signature).StartDate(startDate).EndDate(endDate).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningHashTransferConfigPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningHashTransferConfigPost`: InlineResponse20093
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningHashTransferConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningHashTransferConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | Mining Account | 
 **algo** | **string** | Algorithm(sha256) | 
 **toPoolUser** | **string** | Mining Account | 
 **hashRate** | **string** | Resale hashrate h/s must be transferred (BTC is greater than 500000000000 ETH is greater than 500000) | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **startDate** | **string** | Search date, millisecond timestamp, while empty query all | 
 **endDate** | **string** | Search date, millisecond timestamp, while empty query all | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20093**](InlineResponse20093.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningHashTransferProfitDetailsGet

> InlineResponse20092 SapiV1MiningHashTransferProfitDetailsGet(ctx).ConfigId(configId).UserName(userName).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Hashrate Resale Details (USER_DATA)



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
    configId := "configId_example" // string | Mining ID
    userName := "userName_example" // string | Mining Account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    pageIndex := int32(56) // int32 | Page number, default is first page, start form 1 (optional)
    pageSize := "pageSize_example" // string | Number of pages, minimum 10, maximum 200 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningHashTransferProfitDetailsGet(context.Background()).ConfigId(configId).UserName(userName).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningHashTransferProfitDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningHashTransferProfitDetailsGet`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningHashTransferProfitDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningHashTransferProfitDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configId** | **string** | Mining ID | 
 **userName** | **string** | Mining Account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **pageIndex** | **int32** | Page number, default is first page, start form 1 | 
 **pageSize** | **string** | Number of pages, minimum 10, maximum 200 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningPaymentListGet

> InlineResponse20089 SapiV1MiningPaymentListGet(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Earnings List (USER_DATA)



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
    algo := "algo_example" // string | Algorithm(sha256)
    userName := "userName_example" // string | Mining Account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    coin := "BNB" // string | Coin name (optional)
    startDate := "startDate_example" // string | Search date, millisecond timestamp, while empty query all (optional)
    endDate := "endDate_example" // string | Search date, millisecond timestamp, while empty query all (optional)
    pageIndex := int32(56) // int32 | Page number, default is first page, start form 1 (optional)
    pageSize := "pageSize_example" // string | Number of pages, minimum 10, maximum 200 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningPaymentListGet(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningPaymentListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningPaymentListGet`: InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningPaymentListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningPaymentListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | 
 **userName** | **string** | Mining Account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **coin** | **string** | Coin name | 
 **startDate** | **string** | Search date, millisecond timestamp, while empty query all | 
 **endDate** | **string** | Search date, millisecond timestamp, while empty query all | 
 **pageIndex** | **int32** | Page number, default is first page, start form 1 | 
 **pageSize** | **string** | Number of pages, minimum 10, maximum 200 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20089**](InlineResponse20089.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningPaymentOtherGet

> InlineResponse20090 SapiV1MiningPaymentOtherGet(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Extra Bonus List (USER_DATA)



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
    algo := "algo_example" // string | Algorithm(sha256)
    userName := "userName_example" // string | Mining Account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    coin := "BNB" // string | Coin name (optional)
    startDate := "startDate_example" // string | Search date, millisecond timestamp, while empty query all (optional)
    endDate := "endDate_example" // string | Search date, millisecond timestamp, while empty query all (optional)
    pageIndex := int32(56) // int32 | Page number, default is first page, start form 1 (optional)
    pageSize := "pageSize_example" // string | Number of pages, minimum 10, maximum 200 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningPaymentOtherGet(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningPaymentOtherGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningPaymentOtherGet`: InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningPaymentOtherGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningPaymentOtherGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | 
 **userName** | **string** | Mining Account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **coin** | **string** | Coin name | 
 **startDate** | **string** | Search date, millisecond timestamp, while empty query all | 
 **endDate** | **string** | Search date, millisecond timestamp, while empty query all | 
 **pageIndex** | **int32** | Page number, default is first page, start form 1 | 
 **pageSize** | **string** | Number of pages, minimum 10, maximum 200 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20090**](InlineResponse20090.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningPubAlgoListGet

> InlineResponse20085 SapiV1MiningPubAlgoListGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Acquiring Algorithm (MARKET_DATA)



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
    resp, r, err := api_client.MiningApi.SapiV1MiningPubAlgoListGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningPubAlgoListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningPubAlgoListGet`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningPubAlgoListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningPubAlgoListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningPubCoinListGet

> InlineResponse20086 SapiV1MiningPubCoinListGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Acquiring CoinName (MARKET_DATA)



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
    resp, r, err := api_client.MiningApi.SapiV1MiningPubCoinListGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningPubCoinListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningPubCoinListGet`: InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningPubCoinListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningPubCoinListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20086**](InlineResponse20086.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningStatisticsUserListGet

> InlineResponse20096 SapiV1MiningStatisticsUserListGet(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Account List (USER_DATA)



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
    algo := "algo_example" // string | Algorithm(sha256)
    userName := "userName_example" // string | Mining Account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningStatisticsUserListGet(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningStatisticsUserListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningStatisticsUserListGet`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningStatisticsUserListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningStatisticsUserListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | 
 **userName** | **string** | Mining Account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20096**](InlineResponse20096.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningStatisticsUserStatusGet

> InlineResponse20095 SapiV1MiningStatisticsUserStatusGet(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Statistic List (USER_DATA)



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
    algo := "algo_example" // string | Algorithm(sha256)
    userName := "userName_example" // string | Mining Account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningStatisticsUserStatusGet(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningStatisticsUserStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningStatisticsUserStatusGet`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningStatisticsUserStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningStatisticsUserStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | 
 **userName** | **string** | Mining Account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningWorkerDetailGet

> InlineResponse20087 SapiV1MiningWorkerDetailGet(ctx).Algo(algo).UserName(userName).WorkerName(workerName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Request for Detail Miner List (USER_DATA)



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
    algo := "algo_example" // string | Algorithm(sha256)
    userName := "userName_example" // string | Mining Account
    workerName := "workerName_example" // string | Miner’s name
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningWorkerDetailGet(context.Background()).Algo(algo).UserName(userName).WorkerName(workerName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningWorkerDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningWorkerDetailGet`: InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningWorkerDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningWorkerDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | 
 **userName** | **string** | Mining Account | 
 **workerName** | **string** | Miner’s name | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20087**](InlineResponse20087.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MiningWorkerListGet

> InlineResponse20088 SapiV1MiningWorkerListGet(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()

Request for Miner List (USER_DATA)



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
    algo := "algo_example" // string | Algorithm(sha256)
    userName := "userName_example" // string | Mining Account
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    pageIndex := int32(56) // int32 | Page number, default is first page, start form 1 (optional)
    sort := int32(56) // int32 | sort sequence（default=0）0 positive sequence, 1 negative sequence (optional)
    sortColumn := int32(56) // int32 | Sort by( default 1): 1: miner name, 2: real-time computing power, 3: daily average computing power, 4: real-time rejection rate, 5: last submission time (optional)
    workerStatus := int32(56) // int32 | miners status（default=0）0 all, 1 valid, 2 invalid, 3 failure (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiningApi.SapiV1MiningWorkerListGet(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.SapiV1MiningWorkerListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MiningWorkerListGet`: InlineResponse20088
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.SapiV1MiningWorkerListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MiningWorkerListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | 
 **userName** | **string** | Mining Account | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **pageIndex** | **int32** | Page number, default is first page, start form 1 | 
 **sort** | **int32** | sort sequence（default&#x3D;0）0 positive sequence, 1 negative sequence | 
 **sortColumn** | **int32** | Sort by( default 1): 1: miner name, 2: real-time computing power, 3: daily average computing power, 4: real-time rejection rate, 5: last submission time | 
 **workerStatus** | **int32** | miners status（default&#x3D;0）0 all, 1 valid, 2 invalid, 3 failure | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20088**](InlineResponse20088.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

