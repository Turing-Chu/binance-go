# \BSwapApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BswapAddLiquidityPreview**](BSwapApi.md#BswapAddLiquidityPreview) | **Get** /sapi/v1/bswap/addLiquidityPreview | Add Liquidity Preview (USER_DATA)
[**BswapAddLiquidityToPool**](BSwapApi.md#BswapAddLiquidityToPool) | **Post** /sapi/v1/bswap/liquidityAdd | Add Liquidity (TRADE)
[**BswapGetLiquidityOperationHistory**](BSwapApi.md#BswapGetLiquidityOperationHistory) | **Get** /sapi/v1/bswap/liquidityOps | Liquidity Operation Record (USER_DATA)
[**BswapGetPoolConfigure**](BSwapApi.md#BswapGetPoolConfigure) | **Get** /sapi/v1/bswap/poolConfigure | Pool Configure (USER_DATA)
[**BswapGetPoolLiquidityInfo**](BSwapApi.md#BswapGetPoolLiquidityInfo) | **Get** /sapi/v1/bswap/liquidity | Liquidity information of a pool (USER_DATA)
[**BswapGetPools**](BSwapApi.md#BswapGetPools) | **Get** /sapi/v1/bswap/pools | List All Swap Pools (MARKET_DATA)
[**BswapGetQuoteBasePrice**](BSwapApi.md#BswapGetQuoteBasePrice) | **Get** /sapi/v1/bswap/quote | Request Quote (USER_DATA)
[**BswapRemoveLiquidityFromPool**](BSwapApi.md#BswapRemoveLiquidityFromPool) | **Post** /sapi/v1/bswap/liquidityRemove | Remove Liquidity (TRADE)
[**BswapRemoveLiquidityPreview**](BSwapApi.md#BswapRemoveLiquidityPreview) | **Get** /sapi/v1/bswap/removeLiquidityPreview | Remove Liquidity Preview (USER_DATA)
[**SwapGetHistory**](BSwapApi.md#SwapGetHistory) | **Get** /sapi/v1/bswap/swap | Swap History (USER_DATA)
[**SwapQoteAssetToBaseAsset**](BSwapApi.md#SwapQoteAssetToBaseAsset) | **Post** /sapi/v1/bswap/swap | Swap (TRADE)



## BswapAddLiquidityPreview

> OneOfbswapAddLiquidityPreviewCombinationbswapAddLiquidityPreviewSingle BswapAddLiquidityPreview(ctx).PoolId(poolId).Type_(type_).QuoteAsset(quoteAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Add Liquidity Preview (USER_DATA)



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
    poolId := int64(2) // int64 | 
    type_ := "SINGLE" // string | \"SINGLE\" for adding a single token;\"COMBINATION\" for adding dual tokens
    quoteAsset := "USDT" // string | 
    quoteQty := float64(1.2) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.BswapAddLiquidityPreview(context.Background()).PoolId(poolId).Type_(type_).QuoteAsset(quoteAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapAddLiquidityPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapAddLiquidityPreview`: OneOfbswapAddLiquidityPreviewCombinationbswapAddLiquidityPreviewSingle
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapAddLiquidityPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBswapAddLiquidityPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **int64** |  | 
 **type_** | **string** | \&quot;SINGLE\&quot; for adding a single token;\&quot;COMBINATION\&quot; for adding dual tokens | 
 **quoteAsset** | **string** |  | 
 **quoteQty** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfbswapAddLiquidityPreviewCombinationbswapAddLiquidityPreviewSingle**](oneOf&lt;bswapAddLiquidityPreviewCombination,bswapAddLiquidityPreviewSingle&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BswapAddLiquidityToPool

> InlineResponse200105 BswapAddLiquidityToPool(ctx).PoolId(poolId).Asset(asset).Quantity(quantity).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Add Liquidity (TRADE)



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
    poolId := int64(789) // int64 | 
    asset := "BTC" // string | 
    quantity := float64(1.2) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.BswapAddLiquidityToPool(context.Background()).PoolId(poolId).Asset(asset).Quantity(quantity).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapAddLiquidityToPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapAddLiquidityToPool`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapAddLiquidityToPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBswapAddLiquidityToPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **int64** |  | 
 **asset** | **string** |  | 
 **quantity** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse200105**](InlineResponse200105.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BswapGetLiquidityOperationHistory

> []InlineResponse200106 BswapGetLiquidityOperationHistory(ctx).Timestamp(timestamp).Signature(signature).OperationId(operationId).PoolId(poolId).Operation(operation).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Liquidity Operation Record (USER_DATA)



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
    operationId := int64(789) // int64 |  (optional)
    poolId := int64(789) // int64 |  (optional)
    operation := "operation_example" // string |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.BswapGetLiquidityOperationHistory(context.Background()).Timestamp(timestamp).Signature(signature).OperationId(operationId).PoolId(poolId).Operation(operation).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapGetLiquidityOperationHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapGetLiquidityOperationHistory`: []InlineResponse200106
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapGetLiquidityOperationHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBswapGetLiquidityOperationHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **operationId** | **int64** |  | 
 **poolId** | **int64** |  | 
 **operation** | **string** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse200106**](InlineResponse200106.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BswapGetPoolConfigure

> []InlineResponse200110 BswapGetPoolConfigure(ctx).Timestamp(timestamp).Signature(signature).PoolId(poolId).RecvWindow(recvWindow).Execute()

Pool Configure (USER_DATA)



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
    poolId := int64(2) // int64 |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.BswapGetPoolConfigure(context.Background()).Timestamp(timestamp).Signature(signature).PoolId(poolId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapGetPoolConfigure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapGetPoolConfigure`: []InlineResponse200110
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapGetPoolConfigure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBswapGetPoolConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **poolId** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse200110**](InlineResponse200110.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BswapGetPoolLiquidityInfo

> []InlineResponse200104 BswapGetPoolLiquidityInfo(ctx).Timestamp(timestamp).Signature(signature).PoolId(poolId).RecvWindow(recvWindow).Execute()

Liquidity information of a pool (USER_DATA)



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
    poolId := int64(789) // int64 |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.BswapGetPoolLiquidityInfo(context.Background()).Timestamp(timestamp).Signature(signature).PoolId(poolId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapGetPoolLiquidityInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapGetPoolLiquidityInfo`: []InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapGetPoolLiquidityInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBswapGetPoolLiquidityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **poolId** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse200104**](InlineResponse200104.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BswapGetPools

> []InlineResponse200103 BswapGetPools(ctx).Execute()

List All Swap Pools (MARKET_DATA)



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
    resp, r, err := api_client.BSwapApi.BswapGetPools(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapGetPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapGetPools`: []InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapGetPools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBswapGetPoolsRequest struct via the builder pattern


### Return type

[**[]InlineResponse200103**](InlineResponse200103.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BswapGetQuoteBasePrice

> InlineResponse200107 BswapGetQuoteBasePrice(ctx).QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Request Quote (USER_DATA)



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
    quoteAsset := "USDT" // string | 
    baseAsset := "BUSD" // string | 
    quoteQty := float64(1.2) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.BswapGetQuoteBasePrice(context.Background()).QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapGetQuoteBasePrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapGetQuoteBasePrice`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapGetQuoteBasePrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBswapGetQuoteBasePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteAsset** | **string** |  | 
 **baseAsset** | **string** |  | 
 **quoteQty** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BswapRemoveLiquidityFromPool

> InlineResponse200105 BswapRemoveLiquidityFromPool(ctx).PoolId(poolId).Type_(type_).ShareAmount(shareAmount).Timestamp(timestamp).Signature(signature).Asset(asset).RecvWindow(recvWindow).Execute()

Remove Liquidity (TRADE)



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
    poolId := int64(789) // int64 | 
    type_ := "SINGLE" // string | Can be `SINGLE` for single asset removal, `COMBINATION` for combination of all coins removal
    shareAmount := float64(1.2) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    asset := "BNB" // string | Mandatory for single asset removal (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.BswapRemoveLiquidityFromPool(context.Background()).PoolId(poolId).Type_(type_).ShareAmount(shareAmount).Timestamp(timestamp).Signature(signature).Asset(asset).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapRemoveLiquidityFromPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapRemoveLiquidityFromPool`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapRemoveLiquidityFromPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBswapRemoveLiquidityFromPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **int64** |  | 
 **type_** | **string** | Can be &#x60;SINGLE&#x60; for single asset removal, &#x60;COMBINATION&#x60; for combination of all coins removal | 
 **shareAmount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** | Mandatory for single asset removal | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse200105**](InlineResponse200105.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BswapRemoveLiquidityPreview

> OneOfbswapRmvLiquidityPreviewCombinationbswapRmvLiquidityPreviewSingle BswapRemoveLiquidityPreview(ctx).PoolId(poolId).Type_(type_).QuoteAsset(quoteAsset).ShareAmount(shareAmount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Remove Liquidity Preview (USER_DATA)



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
    poolId := int64(2) // int64 | 
    type_ := "SINGLE" // string | Type is \"SINGLE\", remove and obtain a single token;Type is \"COMBINATION\", remove and obtain dual token.
    quoteAsset := "USDT" // string | 
    shareAmount := float64(1.2) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.BswapRemoveLiquidityPreview(context.Background()).PoolId(poolId).Type_(type_).QuoteAsset(quoteAsset).ShareAmount(shareAmount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.BswapRemoveLiquidityPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BswapRemoveLiquidityPreview`: OneOfbswapRmvLiquidityPreviewCombinationbswapRmvLiquidityPreviewSingle
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.BswapRemoveLiquidityPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBswapRemoveLiquidityPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **int64** |  | 
 **type_** | **string** | Type is \&quot;SINGLE\&quot;, remove and obtain a single token;Type is \&quot;COMBINATION\&quot;, remove and obtain dual token. | 
 **quoteAsset** | **string** |  | 
 **shareAmount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfbswapRmvLiquidityPreviewCombinationbswapRmvLiquidityPreviewSingle**](oneOf&lt;bswapRmvLiquidityPreviewCombination,bswapRmvLiquidityPreviewSingle&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapGetHistory

> []InlineResponse200108 SwapGetHistory(ctx).Timestamp(timestamp).Signature(signature).SwapId(swapId).StartTime(startTime).EndTime(endTime).Status(status).QuoteAsset(quoteAsset).BaseAsset(baseAsset).Limit(limit).RecvWindow(recvWindow).Execute()

Swap History (USER_DATA)



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
    swapId := int64(789) // int64 |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    status := int32(56) // int32 | 0: pending for swap, 1: success, 2: failed (optional)
    quoteAsset := "USDT" // string |  (optional)
    baseAsset := "BUSD" // string |  (optional)
    limit := int32(56) // int32 | default 3, max 100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.SwapGetHistory(context.Background()).Timestamp(timestamp).Signature(signature).SwapId(swapId).StartTime(startTime).EndTime(endTime).Status(status).QuoteAsset(quoteAsset).BaseAsset(baseAsset).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SwapGetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapGetHistory`: []InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SwapGetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwapGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **swapId** | **int64** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **status** | **int32** | 0: pending for swap, 1: success, 2: failed | 
 **quoteAsset** | **string** |  | 
 **baseAsset** | **string** |  | 
 **limit** | **int32** | default 3, max 100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse200108**](InlineResponse200108.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapQoteAssetToBaseAsset

> InlineResponse200109 SwapQoteAssetToBaseAsset(ctx).QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Swap (TRADE)



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
    quoteAsset := "USDT" // string | 
    baseAsset := "BUSD" // string | 
    quoteQty := float64(1.2) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSwapApi.SwapQoteAssetToBaseAsset(context.Background()).QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SwapQoteAssetToBaseAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapQoteAssetToBaseAsset`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SwapQoteAssetToBaseAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwapQoteAssetToBaseAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteAsset** | **string** |  | 
 **baseAsset** | **string** |  | 
 **quoteQty** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

