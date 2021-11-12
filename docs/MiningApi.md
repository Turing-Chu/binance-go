# \MiningApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlvtGetTokenIn**](MiningApi.md#BlvtGetTokenIn) | **Get** /sapi/v1/mining/statistics/user/status | Statistic List (USER_DATA)
[**MiningCancelHashrateResaleConfig**](MiningApi.md#MiningCancelHashrateResaleConfig) | **Post** /sapi/v1/mining/hash-transfer/config/cancel | Cancel Hashrate Resale configuration (USER_DATA)
[**MiningGetAccounts**](MiningApi.md#MiningGetAccounts) | **Get** /sapi/v1/mining/statistics/user/list | Account List (USER_DATA)
[**MiningGetAlgorithms**](MiningApi.md#MiningGetAlgorithms) | **Get** /sapi/v1/mining/pub/algoList | Acquiring Algorithm (MARKET_DATA)
[**MiningGetCoins**](MiningApi.md#MiningGetCoins) | **Get** /sapi/v1/mining/pub/coinList | Acquiring CoinName (MARKET_DATA)
[**MiningGetEarnings**](MiningApi.md#MiningGetEarnings) | **Get** /sapi/v1/mining/payment/list | Earnings List (USER_DATA)
[**MiningGetHashrateResaleDetai**](MiningApi.md#MiningGetHashrateResaleDetai) | **Get** /sapi/v1/mining/hash-transfer/profit/details | Hashrate Resale Details (USER_DATA)
[**MiningGetOtherPayment**](MiningApi.md#MiningGetOtherPayment) | **Get** /sapi/v1/mining/payment/other | Extra Bonus List (USER_DATA)
[**MiningGetWorkerDetail**](MiningApi.md#MiningGetWorkerDetail) | **Get** /sapi/v1/mining/worker/detail | Request for Detail Miner List (USER_DATA)
[**MiningGetWorkers**](MiningApi.md#MiningGetWorkers) | **Get** /sapi/v1/mining/worker/list | Request for Miner List (USER_DATA)
[**MiningListHashrateResale**](MiningApi.md#MiningListHashrateResale) | **Get** /sapi/v1/mining/hash-transfer/config/details/list | Hashrate Resale List (USER_DATA)
[**MiningRequestHashrateResale**](MiningApi.md#MiningRequestHashrateResale) | **Post** /sapi/v1/mining/hash-transfer/config | Hashrate Resale Request (USER_DATA)



## BlvtGetTokenIn

> InlineResponse20095 BlvtGetTokenIn(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.BlvtGetTokenIn(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.BlvtGetTokenIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlvtGetTokenIn`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.BlvtGetTokenIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlvtGetTokenInRequest struct via the builder pattern


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


## MiningCancelHashrateResaleConfig

> InlineResponse20094 MiningCancelHashrateResaleConfig(ctx).ConfigId(configId).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningCancelHashrateResaleConfig(context.Background()).ConfigId(configId).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningCancelHashrateResaleConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningCancelHashrateResaleConfig`: InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningCancelHashrateResaleConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningCancelHashrateResaleConfigRequest struct via the builder pattern


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


## MiningGetAccounts

> InlineResponse20096 MiningGetAccounts(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningGetAccounts(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningGetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningGetAccounts`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningGetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningGetAccountsRequest struct via the builder pattern


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


## MiningGetAlgorithms

> InlineResponse20085 MiningGetAlgorithms(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningGetAlgorithms(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningGetAlgorithms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningGetAlgorithms`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningGetAlgorithms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningGetAlgorithmsRequest struct via the builder pattern


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


## MiningGetCoins

> InlineResponse20086 MiningGetCoins(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningGetCoins(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningGetCoins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningGetCoins`: InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningGetCoins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningGetCoinsRequest struct via the builder pattern


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


## MiningGetEarnings

> InlineResponse20089 MiningGetEarnings(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningGetEarnings(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningGetEarnings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningGetEarnings`: InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningGetEarnings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningGetEarningsRequest struct via the builder pattern


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


## MiningGetHashrateResaleDetai

> InlineResponse20092 MiningGetHashrateResaleDetai(ctx).ConfigId(configId).UserName(userName).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningGetHashrateResaleDetai(context.Background()).ConfigId(configId).UserName(userName).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningGetHashrateResaleDetai``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningGetHashrateResaleDetai`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningGetHashrateResaleDetai`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningGetHashrateResaleDetaiRequest struct via the builder pattern


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


## MiningGetOtherPayment

> InlineResponse20090 MiningGetOtherPayment(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningGetOtherPayment(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningGetOtherPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningGetOtherPayment`: InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningGetOtherPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningGetOtherPaymentRequest struct via the builder pattern


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


## MiningGetWorkerDetail

> InlineResponse20087 MiningGetWorkerDetail(ctx).Algo(algo).UserName(userName).WorkerName(workerName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningGetWorkerDetail(context.Background()).Algo(algo).UserName(userName).WorkerName(workerName).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningGetWorkerDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningGetWorkerDetail`: InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningGetWorkerDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningGetWorkerDetailRequest struct via the builder pattern


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


## MiningGetWorkers

> InlineResponse20088 MiningGetWorkers(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningGetWorkers(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningGetWorkers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningGetWorkers`: InlineResponse20088
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningGetWorkers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningGetWorkersRequest struct via the builder pattern


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


## MiningListHashrateResale

> InlineResponse20091 MiningListHashrateResale(ctx).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningListHashrateResale(context.Background()).Timestamp(timestamp).Signature(signature).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningListHashrateResale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningListHashrateResale`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningListHashrateResale`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningListHashrateResaleRequest struct via the builder pattern


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


## MiningRequestHashrateResale

> InlineResponse20093 MiningRequestHashrateResale(ctx).UserName(userName).Algo(algo).ToPoolUser(toPoolUser).HashRate(hashRate).Timestamp(timestamp).Signature(signature).StartDate(startDate).EndDate(endDate).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MiningApi.MiningRequestHashrateResale(context.Background()).UserName(userName).Algo(algo).ToPoolUser(toPoolUser).HashRate(hashRate).Timestamp(timestamp).Signature(signature).StartDate(startDate).EndDate(endDate).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiningApi.MiningRequestHashrateResale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MiningRequestHashrateResale`: InlineResponse20093
    fmt.Fprintf(os.Stdout, "Response from `MiningApi.MiningRequestHashrateResale`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiningRequestHashrateResaleRequest struct via the builder pattern


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

