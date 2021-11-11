# \BSwapApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1BswapAddLiquidityPreviewGet**](BSwapApi.md#SapiV1BswapAddLiquidityPreviewGet) | **Get** /sapi/v1/bswap/addLiquidityPreview | Add Liquidity Preview (USER_DATA)
[**SapiV1BswapLiquidityAddPost**](BSwapApi.md#SapiV1BswapLiquidityAddPost) | **Post** /sapi/v1/bswap/liquidityAdd | Add Liquidity (TRADE)
[**SapiV1BswapLiquidityGet**](BSwapApi.md#SapiV1BswapLiquidityGet) | **Get** /sapi/v1/bswap/liquidity | Liquidity information of a pool (USER_DATA)
[**SapiV1BswapLiquidityOpsGet**](BSwapApi.md#SapiV1BswapLiquidityOpsGet) | **Get** /sapi/v1/bswap/liquidityOps | Liquidity Operation Record (USER_DATA)
[**SapiV1BswapLiquidityRemovePost**](BSwapApi.md#SapiV1BswapLiquidityRemovePost) | **Post** /sapi/v1/bswap/liquidityRemove | Remove Liquidity (TRADE)
[**SapiV1BswapPoolConfigureGet**](BSwapApi.md#SapiV1BswapPoolConfigureGet) | **Get** /sapi/v1/bswap/poolConfigure | Pool Configure (USER_DATA)
[**SapiV1BswapPoolsGet**](BSwapApi.md#SapiV1BswapPoolsGet) | **Get** /sapi/v1/bswap/pools | List All Swap Pools (MARKET_DATA)
[**SapiV1BswapQuoteGet**](BSwapApi.md#SapiV1BswapQuoteGet) | **Get** /sapi/v1/bswap/quote | Request Quote (USER_DATA)
[**SapiV1BswapRemoveLiquidityPreviewGet**](BSwapApi.md#SapiV1BswapRemoveLiquidityPreviewGet) | **Get** /sapi/v1/bswap/removeLiquidityPreview | Remove Liquidity Preview (USER_DATA)
[**SapiV1BswapSwapGet**](BSwapApi.md#SapiV1BswapSwapGet) | **Get** /sapi/v1/bswap/swap | Swap History (USER_DATA)
[**SapiV1BswapSwapPost**](BSwapApi.md#SapiV1BswapSwapPost) | **Post** /sapi/v1/bswap/swap | Swap (TRADE)



## SapiV1BswapAddLiquidityPreviewGet

> OneOfbswapAddLiquidityPreviewCombinationbswapAddLiquidityPreviewSingle SapiV1BswapAddLiquidityPreviewGet(ctx).PoolId(poolId).Type_(type_).QuoteAsset(quoteAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapAddLiquidityPreviewGet(context.Background()).PoolId(poolId).Type_(type_).QuoteAsset(quoteAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapAddLiquidityPreviewGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapAddLiquidityPreviewGet`: OneOfbswapAddLiquidityPreviewCombinationbswapAddLiquidityPreviewSingle
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapAddLiquidityPreviewGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapAddLiquidityPreviewGetRequest struct via the builder pattern


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


## SapiV1BswapLiquidityAddPost

> InlineResponse200105 SapiV1BswapLiquidityAddPost(ctx).PoolId(poolId).Asset(asset).Quantity(quantity).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapLiquidityAddPost(context.Background()).PoolId(poolId).Asset(asset).Quantity(quantity).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapLiquidityAddPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapLiquidityAddPost`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapLiquidityAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapLiquidityAddPostRequest struct via the builder pattern


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


## SapiV1BswapLiquidityGet

> []InlineResponse200104 SapiV1BswapLiquidityGet(ctx).Timestamp(timestamp).Signature(signature).PoolId(poolId).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapLiquidityGet(context.Background()).Timestamp(timestamp).Signature(signature).PoolId(poolId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapLiquidityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapLiquidityGet`: []InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapLiquidityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapLiquidityGetRequest struct via the builder pattern


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


## SapiV1BswapLiquidityOpsGet

> []InlineResponse200106 SapiV1BswapLiquidityOpsGet(ctx).Timestamp(timestamp).Signature(signature).OperationId(operationId).PoolId(poolId).Operation(operation).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapLiquidityOpsGet(context.Background()).Timestamp(timestamp).Signature(signature).OperationId(operationId).PoolId(poolId).Operation(operation).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapLiquidityOpsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapLiquidityOpsGet`: []InlineResponse200106
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapLiquidityOpsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapLiquidityOpsGetRequest struct via the builder pattern


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


## SapiV1BswapLiquidityRemovePost

> InlineResponse200105 SapiV1BswapLiquidityRemovePost(ctx).PoolId(poolId).Type_(type_).ShareAmount(shareAmount).Timestamp(timestamp).Signature(signature).Asset(asset).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapLiquidityRemovePost(context.Background()).PoolId(poolId).Type_(type_).ShareAmount(shareAmount).Timestamp(timestamp).Signature(signature).Asset(asset).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapLiquidityRemovePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapLiquidityRemovePost`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapLiquidityRemovePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapLiquidityRemovePostRequest struct via the builder pattern


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


## SapiV1BswapPoolConfigureGet

> []InlineResponse200110 SapiV1BswapPoolConfigureGet(ctx).Timestamp(timestamp).Signature(signature).PoolId(poolId).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapPoolConfigureGet(context.Background()).Timestamp(timestamp).Signature(signature).PoolId(poolId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapPoolConfigureGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapPoolConfigureGet`: []InlineResponse200110
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapPoolConfigureGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapPoolConfigureGetRequest struct via the builder pattern


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


## SapiV1BswapPoolsGet

> []InlineResponse200103 SapiV1BswapPoolsGet(ctx).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapPoolsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapPoolsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapPoolsGet`: []InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapPoolsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapPoolsGetRequest struct via the builder pattern


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


## SapiV1BswapQuoteGet

> InlineResponse200107 SapiV1BswapQuoteGet(ctx).QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapQuoteGet(context.Background()).QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapQuoteGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapQuoteGet`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapQuoteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapQuoteGetRequest struct via the builder pattern


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


## SapiV1BswapRemoveLiquidityPreviewGet

> OneOfbswapRmvLiquidityPreviewCombinationbswapRmvLiquidityPreviewSingle SapiV1BswapRemoveLiquidityPreviewGet(ctx).PoolId(poolId).Type_(type_).QuoteAsset(quoteAsset).ShareAmount(shareAmount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapRemoveLiquidityPreviewGet(context.Background()).PoolId(poolId).Type_(type_).QuoteAsset(quoteAsset).ShareAmount(shareAmount).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapRemoveLiquidityPreviewGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapRemoveLiquidityPreviewGet`: OneOfbswapRmvLiquidityPreviewCombinationbswapRmvLiquidityPreviewSingle
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapRemoveLiquidityPreviewGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapRemoveLiquidityPreviewGetRequest struct via the builder pattern


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


## SapiV1BswapSwapGet

> []InlineResponse200108 SapiV1BswapSwapGet(ctx).Timestamp(timestamp).Signature(signature).SwapId(swapId).StartTime(startTime).EndTime(endTime).Status(status).QuoteAsset(quoteAsset).BaseAsset(baseAsset).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapSwapGet(context.Background()).Timestamp(timestamp).Signature(signature).SwapId(swapId).StartTime(startTime).EndTime(endTime).Status(status).QuoteAsset(quoteAsset).BaseAsset(baseAsset).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapSwapGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapSwapGet`: []InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapSwapGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapSwapGetRequest struct via the builder pattern


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


## SapiV1BswapSwapPost

> InlineResponse200109 SapiV1BswapSwapPost(ctx).QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.BSwapApi.SapiV1BswapSwapPost(context.Background()).QuoteAsset(quoteAsset).BaseAsset(baseAsset).QuoteQty(quoteQty).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSwapApi.SapiV1BswapSwapPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BswapSwapPost`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `BSwapApi.SapiV1BswapSwapPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BswapSwapPostRequest struct via the builder pattern


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

