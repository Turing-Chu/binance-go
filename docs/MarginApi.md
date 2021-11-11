# \MarginApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1BnbBurnGet**](MarginApi.md#SapiV1BnbBurnGet) | **Get** /sapi/v1/bnbBurn | Get All Isolated Margin Symbol(USER_DATA)
[**SapiV1BnbBurnPost**](MarginApi.md#SapiV1BnbBurnPost) | **Post** /sapi/v1/bnbBurn | Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
[**SapiV1MarginAccountGet**](MarginApi.md#SapiV1MarginAccountGet) | **Get** /sapi/v1/margin/account | Query Cross Margin Account Details (USER_DATA)
[**SapiV1MarginAllAssetsGet**](MarginApi.md#SapiV1MarginAllAssetsGet) | **Get** /sapi/v1/margin/allAssets | Get All Margin Assets (MARKET_DATA)
[**SapiV1MarginAllOrderListGet**](MarginApi.md#SapiV1MarginAllOrderListGet) | **Get** /sapi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**SapiV1MarginAllOrdersGet**](MarginApi.md#SapiV1MarginAllOrdersGet) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#39;s All Orders (USER_DATA)
[**SapiV1MarginAllPairsGet**](MarginApi.md#SapiV1MarginAllPairsGet) | **Get** /sapi/v1/margin/allPairs | Get All Cross Margin Pairs (MARKET_DATA)
[**SapiV1MarginAssetGet**](MarginApi.md#SapiV1MarginAssetGet) | **Get** /sapi/v1/margin/asset | Query Margin Asset (MARKET_DATA)
[**SapiV1MarginForceLiquidationRecGet**](MarginApi.md#SapiV1MarginForceLiquidationRecGet) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**SapiV1MarginInterestHistoryGet**](MarginApi.md#SapiV1MarginInterestHistoryGet) | **Get** /sapi/v1/margin/interestHistory | Get Interest History (USER_DATA)
[**SapiV1MarginInterestRateHistoryGet**](MarginApi.md#SapiV1MarginInterestRateHistoryGet) | **Get** /sapi/v1/margin/interestRateHistory | Margin Interest Rate History (USER_DATA)
[**SapiV1MarginIsolatedAccountDelete**](MarginApi.md#SapiV1MarginIsolatedAccountDelete) | **Delete** /sapi/v1/margin/isolated/account | Disable Isolated Margin Account (TRADE)
[**SapiV1MarginIsolatedAccountGet**](MarginApi.md#SapiV1MarginIsolatedAccountGet) | **Get** /sapi/v1/margin/isolated/account | Query Isolated Margin Account Info (USER_DATA)
[**SapiV1MarginIsolatedAccountLimitGet**](MarginApi.md#SapiV1MarginIsolatedAccountLimitGet) | **Get** /sapi/v1/margin/isolated/accountLimit | Query Enabled Isolated Margin Account Limit (USER_DATA)
[**SapiV1MarginIsolatedAccountPost**](MarginApi.md#SapiV1MarginIsolatedAccountPost) | **Post** /sapi/v1/margin/isolated/account | Enable Isolated Margin Account (TRADE)
[**SapiV1MarginIsolatedAllPairsGet**](MarginApi.md#SapiV1MarginIsolatedAllPairsGet) | **Get** /sapi/v1/margin/isolated/allPairs | Get All Isolated Margin Symbol(USER_DATA)
[**SapiV1MarginIsolatedPairGet**](MarginApi.md#SapiV1MarginIsolatedPairGet) | **Get** /sapi/v1/margin/isolated/pair | Query Isolated Margin Symbol (USER_DATA)
[**SapiV1MarginIsolatedTransferGet**](MarginApi.md#SapiV1MarginIsolatedTransferGet) | **Get** /sapi/v1/margin/isolated/transfer | Get Isolated Margin Transfer History (USER_DATA)
[**SapiV1MarginIsolatedTransferPost**](MarginApi.md#SapiV1MarginIsolatedTransferPost) | **Post** /sapi/v1/margin/isolated/transfer | Isolated Margin Account Transfer (MARGIN)
[**SapiV1MarginLoanGet**](MarginApi.md#SapiV1MarginLoanGet) | **Get** /sapi/v1/margin/loan | Query Loan Record (USER_DATA)
[**SapiV1MarginLoanPost**](MarginApi.md#SapiV1MarginLoanPost) | **Post** /sapi/v1/margin/loan | Margin Account Borrow (MARGIN)
[**SapiV1MarginMaxBorrowableGet**](MarginApi.md#SapiV1MarginMaxBorrowableGet) | **Get** /sapi/v1/margin/maxBorrowable | Query Max Borrow (USER_DATA)
[**SapiV1MarginMaxTransferableGet**](MarginApi.md#SapiV1MarginMaxTransferableGet) | **Get** /sapi/v1/margin/maxTransferable | Query Max Transfer-Out Amount (USER_DATA)
[**SapiV1MarginMyTradesGet**](MarginApi.md#SapiV1MarginMyTradesGet) | **Get** /sapi/v1/margin/myTrades | Query Margin Account&#39;s Trade List (USER_DATA)
[**SapiV1MarginOpenOrderListGet**](MarginApi.md#SapiV1MarginOpenOrderListGet) | **Get** /sapi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**SapiV1MarginOpenOrdersDelete**](MarginApi.md#SapiV1MarginOpenOrdersDelete) | **Delete** /sapi/v1/margin/openOrders | Margin Account Cancel all Open Orders on a Symbol (TRADE)
[**SapiV1MarginOpenOrdersGet**](MarginApi.md#SapiV1MarginOpenOrdersGet) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Orders (USER_DATA)
[**SapiV1MarginOrderDelete**](MarginApi.md#SapiV1MarginOrderDelete) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**SapiV1MarginOrderGet**](MarginApi.md#SapiV1MarginOrderGet) | **Get** /sapi/v1/margin/order | Query Margin Account&#39;s Order (USER_DATA)
[**SapiV1MarginOrderListDelete**](MarginApi.md#SapiV1MarginOrderListDelete) | **Delete** /sapi/v1/margin/orderList | Margin Account Cancel OCO (TRADE)
[**SapiV1MarginOrderListGet**](MarginApi.md#SapiV1MarginOrderListGet) | **Get** /sapi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**SapiV1MarginOrderOcoPost**](MarginApi.md#SapiV1MarginOrderOcoPost) | **Post** /sapi/v1/margin/order/oco | Margin Account New OCO (TRADE)
[**SapiV1MarginOrderPost**](MarginApi.md#SapiV1MarginOrderPost) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**SapiV1MarginPairGet**](MarginApi.md#SapiV1MarginPairGet) | **Get** /sapi/v1/margin/pair | Query Cross Margin Pair (MARKET_DATA)
[**SapiV1MarginPriceIndexGet**](MarginApi.md#SapiV1MarginPriceIndexGet) | **Get** /sapi/v1/margin/priceIndex | Query Margin PriceIndex (MARKET_DATA)
[**SapiV1MarginRepayGet**](MarginApi.md#SapiV1MarginRepayGet) | **Get** /sapi/v1/margin/repay | Query Repay Record (USER_DATA)
[**SapiV1MarginRepayPost**](MarginApi.md#SapiV1MarginRepayPost) | **Post** /sapi/v1/margin/repay | Margin Account Repay (MARGIN)
[**SapiV1MarginTransferGet**](MarginApi.md#SapiV1MarginTransferGet) | **Get** /sapi/v1/margin/transfer | Get Cross Margin Transfer History (USER_DATA)
[**SapiV1MarginTransferPost**](MarginApi.md#SapiV1MarginTransferPost) | **Post** /sapi/v1/margin/transfer | Cross Margin Account Transfer (MARGIN)



## SapiV1BnbBurnGet

> BnbBurnStatus SapiV1BnbBurnGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Get All Isolated Margin Symbol(USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1BnbBurnGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1BnbBurnGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BnbBurnGet`: BnbBurnStatus
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1BnbBurnGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BnbBurnGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**BnbBurnStatus**](BnbBurnStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1BnbBurnPost

> BnbBurnStatus SapiV1BnbBurnPost(ctx).Timestamp(timestamp).Signature(signature).SpotBNBBurn(spotBNBBurn).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).Execute()

Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)



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
    spotBNBBurn := "true" // string | Determines whether to use BNB to pay for trading fees on SPOT (optional)
    interestBNBBurn := "false" // string | Determines whether to use BNB to pay for margin loan's interest (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1BnbBurnPost(context.Background()).Timestamp(timestamp).Signature(signature).SpotBNBBurn(spotBNBBurn).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1BnbBurnPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1BnbBurnPost`: BnbBurnStatus
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1BnbBurnPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1BnbBurnPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **spotBNBBurn** | **string** | Determines whether to use BNB to pay for trading fees on SPOT | 
 **interestBNBBurn** | **string** | Determines whether to use BNB to pay for margin loan&#39;s interest | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**BnbBurnStatus**](BnbBurnStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAccountGet

> InlineResponse20019 SapiV1MarginAccountGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Query Cross Margin Account Details (USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginAccountGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAccountGet`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAllAssetsGet

> []InlineResponse20014 SapiV1MarginAllAssetsGet(ctx).Execute()

Get All Margin Assets (MARKET_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginAllAssetsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAllAssetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAllAssetsGet`: []InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAllAssetsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAllAssetsGetRequest struct via the builder pattern


### Return type

[**[]InlineResponse20014**](InlineResponse20014.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAllOrderListGet

> []InlineResponse2006 SapiV1MarginAllOrderListGet(ctx).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)



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
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    symbol := "symbol_example" // string | Mandatory for isolated margin, not supported for cross margin (optional)
    fromId := "fromId_example" // string | If supplied, neither `startTime` or `endTime` can be provided (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(56) // int32 | Default Value: 500; Max Value: 1000 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginAllOrderListGet(context.Background()).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAllOrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAllOrderListGet`: []InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAllOrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAllOrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **symbol** | **string** | Mandatory for isolated margin, not supported for cross margin | 
 **fromId** | **string** | If supplied, neither &#x60;startTime&#x60; or &#x60;endTime&#x60; can be provided | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default Value: 500; Max Value: 1000 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse2006**](InlineResponse2006.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAllOrdersGet

> []MarginOrderDetail SapiV1MarginAllOrdersGet(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's All Orders (USER_DATA)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    orderId := int64(789) // int64 | Order id (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginAllOrdersGet(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAllOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAllOrdersGet`: []MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAllOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAllOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **orderId** | **int64** | Order id | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]MarginOrderDetail**](MarginOrderDetail.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAllPairsGet

> []InlineResponse20015 SapiV1MarginAllPairsGet(ctx).Execute()

Get All Cross Margin Pairs (MARKET_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginAllPairsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAllPairsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAllPairsGet`: []InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAllPairsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAllPairsGetRequest struct via the builder pattern


### Return type

[**[]InlineResponse20015**](InlineResponse20015.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginAssetGet

> InlineResponse20012 SapiV1MarginAssetGet(ctx).Asset(asset).Execute()

Query Margin Asset (MARKET_DATA)



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginAssetGet(context.Background()).Asset(asset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginAssetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginAssetGet`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginAssetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginAssetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginForceLiquidationRecGet

> InlineResponse20018 SapiV1MarginForceLiquidationRecGet(ctx).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Force Liquidation Record (USER_DATA)



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
    isolatedSymbol := "isolatedSymbol_example" // string | Isolated symbol (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginForceLiquidationRecGet(context.Background()).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginForceLiquidationRecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginForceLiquidationRecGet`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginForceLiquidationRecGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginForceLiquidationRecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **isolatedSymbol** | **string** | Isolated symbol | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginInterestHistoryGet

> InlineResponse20017 SapiV1MarginInterestHistoryGet(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

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
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    asset := "BNB" // string |  (optional)
    isolatedSymbol := "isolatedSymbol_example" // string | Isolated symbol (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    archived := "archived_example" // string | Default: false. Set to true for archived data from 6 months ago (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginInterestHistoryGet(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginInterestHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginInterestHistoryGet`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginInterestHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginInterestHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **isolatedSymbol** | **string** | Isolated symbol | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **archived** | **string** | Default: false. Set to true for archived data from 6 months ago | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginInterestRateHistoryGet

> []InlineResponse20028 SapiV1MarginInterestRateHistoryGet(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Margin Interest Rate History (USER_DATA)



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
    vipLevel := int32(1) // int32 | Defaults to user's vip level (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginInterestRateHistoryGet(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginInterestRateHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginInterestRateHistoryGet`: []InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginInterestRateHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginInterestRateHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **vipLevel** | **int32** | Defaults to user&#39;s vip level | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20028**](InlineResponse20028.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginIsolatedAccountDelete

> InlineResponse20025 SapiV1MarginIsolatedAccountDelete(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Disable Isolated Margin Account (TRADE)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginIsolatedAccountDelete(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginIsolatedAccountDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginIsolatedAccountDelete`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginIsolatedAccountDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginIsolatedAccountDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginIsolatedAccountGet

> IsolatedMarginAccountInfo SapiV1MarginIsolatedAccountGet(ctx).Timestamp(timestamp).Signature(signature).Symbols(symbols).RecvWindow(recvWindow).Execute()

Query Isolated Margin Account Info (USER_DATA)



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
    symbols := "BTCUSDT,BNBUSDT,ADAUSDT" // string | Max 5 symbols can be sent; separated by ',' (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginIsolatedAccountGet(context.Background()).Timestamp(timestamp).Signature(signature).Symbols(symbols).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginIsolatedAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginIsolatedAccountGet`: IsolatedMarginAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginIsolatedAccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginIsolatedAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **symbols** | **string** | Max 5 symbols can be sent; separated by &#39;,&#39; | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**IsolatedMarginAccountInfo**](IsolatedMarginAccountInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginIsolatedAccountLimitGet

> InlineResponse20026 SapiV1MarginIsolatedAccountLimitGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Query Enabled Isolated Margin Account Limit (USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginIsolatedAccountLimitGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginIsolatedAccountLimitGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginIsolatedAccountLimitGet`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginIsolatedAccountLimitGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginIsolatedAccountLimitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginIsolatedAccountPost

> InlineResponse20025 SapiV1MarginIsolatedAccountPost(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Enable Isolated Margin Account (TRADE)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginIsolatedAccountPost(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginIsolatedAccountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginIsolatedAccountPost`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginIsolatedAccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginIsolatedAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginIsolatedAllPairsGet

> []InlineResponse20027 SapiV1MarginIsolatedAllPairsGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Get All Isolated Margin Symbol(USER_DATA)



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
    resp, r, err := api_client.MarginApi.SapiV1MarginIsolatedAllPairsGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginIsolatedAllPairsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginIsolatedAllPairsGet`: []InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginIsolatedAllPairsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginIsolatedAllPairsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20027**](InlineResponse20027.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginIsolatedPairGet

> InlineResponse20027 SapiV1MarginIsolatedPairGet(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Query Isolated Margin Symbol (USER_DATA)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginIsolatedPairGet(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginIsolatedPairGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginIsolatedPairGet`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginIsolatedPairGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginIsolatedPairGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginIsolatedTransferGet

> MarginTransferDetails SapiV1MarginIsolatedTransferGet(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).Asset(asset).TransFrom(transFrom).TransTo(transTo).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Isolated Margin Transfer History (USER_DATA)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    asset := "BNB" // string |  (optional)
    transFrom := "SPOT" // string |  (optional)
    transTo := "ISOLATED_MARGIN" // string |  (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginIsolatedTransferGet(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).Asset(asset).TransFrom(transFrom).TransTo(transTo).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginIsolatedTransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginIsolatedTransferGet`: MarginTransferDetails
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginIsolatedTransferGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginIsolatedTransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **transFrom** | **string** |  | 
 **transTo** | **string** |  | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**MarginTransferDetails**](MarginTransferDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginIsolatedTransferPost

> map[string]interface{} SapiV1MarginIsolatedTransferPost(ctx).Asset(asset).Symbol(symbol).Amount(amount).Timestamp(timestamp).Signature(signature).TransFrom(transFrom).TransTo(transTo).RecvWindow(recvWindow).Execute()

Isolated Margin Account Transfer (MARGIN)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    amount := float64(1.01) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    transFrom := "SPOT" // string |  (optional)
    transTo := "ISOLATED_MARGIN" // string |  (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginIsolatedTransferPost(context.Background()).Asset(asset).Symbol(symbol).Amount(amount).Timestamp(timestamp).Signature(signature).TransFrom(transFrom).TransTo(transTo).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginIsolatedTransferPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginIsolatedTransferPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginIsolatedTransferPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginIsolatedTransferPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **transFrom** | **string** |  | 
 **transTo** | **string** |  | 
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


## SapiV1MarginLoanGet

> InlineResponse20010 SapiV1MarginLoanGet(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Query Loan Record (USER_DATA)



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
    isolatedSymbol := "isolatedSymbol_example" // string | Isolated symbol (optional)
    txId := int64(123456789) // int64 | the tranId in  `POST /sapi/v1/margin/loan` (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    archived := "archived_example" // string | Default: false. Set to true for archived data from 6 months ago (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginLoanGet(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginLoanGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginLoanGet`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginLoanGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginLoanGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isolatedSymbol** | **string** | Isolated symbol | 
 **txId** | **int64** | the tranId in  &#x60;POST /sapi/v1/margin/loan&#x60; | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **archived** | **string** | Default: false. Set to true for archived data from 6 months ago | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginLoanPost

> Transaction SapiV1MarginLoanPost(ctx).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Margin Account Borrow (MARGIN)



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
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginLoanPost(context.Background()).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginLoanPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginLoanPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginLoanPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginLoanPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginMaxBorrowableGet

> InlineResponse20023 SapiV1MarginMaxBorrowableGet(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Borrow (USER_DATA)



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
    isolatedSymbol := "isolatedSymbol_example" // string | Isolated symbol (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginMaxBorrowableGet(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginMaxBorrowableGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginMaxBorrowableGet`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginMaxBorrowableGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginMaxBorrowableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isolatedSymbol** | **string** | Isolated symbol | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginMaxTransferableGet

> InlineResponse20024 SapiV1MarginMaxTransferableGet(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Transfer-Out Amount (USER_DATA)



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
    isolatedSymbol := "isolatedSymbol_example" // string | Isolated symbol (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginMaxTransferableGet(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginMaxTransferableGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginMaxTransferableGet`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginMaxTransferableGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginMaxTransferableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isolatedSymbol** | **string** | Isolated symbol | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginMyTradesGet

> []MarginTrade SapiV1MarginMyTradesGet(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's Trade List (USER_DATA)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginMyTradesGet(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginMyTradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginMyTradesGet`: []MarginTrade
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginMyTradesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginMyTradesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]MarginTrade**](MarginTrade.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOpenOrderListGet

> []InlineResponse20022 SapiV1MarginOpenOrderListGet(ctx).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)



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
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    symbol := "symbol_example" // string | Mandatory for isolated margin, not supported for cross margin (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOpenOrderListGet(context.Background()).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOpenOrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOpenOrderListGet`: []InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOpenOrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOpenOrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **symbol** | **string** | Mandatory for isolated margin, not supported for cross margin | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse20022**](InlineResponse20022.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOpenOrdersDelete

> []AnyOfcanceledMarginOrderDetailmarginOcoOrder SapiV1MarginOpenOrdersDelete(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Margin Account Cancel all Open Orders on a Symbol (TRADE)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOpenOrdersDelete(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOpenOrdersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOpenOrdersDelete`: []AnyOfcanceledMarginOrderDetailmarginOcoOrder
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOpenOrdersDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOpenOrdersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]AnyOfcanceledMarginOrderDetailmarginOcoOrder**](anyOf&lt;canceledMarginOrderDetail,marginOcoOrder&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOpenOrdersGet

> []MarginOrderDetail SapiV1MarginOpenOrdersGet(ctx).Timestamp(timestamp).Signature(signature).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Query Margin Account's Open Orders (USER_DATA)



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
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOpenOrdersGet(context.Background()).Timestamp(timestamp).Signature(signature).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOpenOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOpenOrdersGet`: []MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOpenOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOpenOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]MarginOrderDetail**](MarginOrderDetail.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderDelete

> MarginOrder SapiV1MarginOrderDelete(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel Order (TRADE)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    orderId := int64(789) // int64 | Order id (optional)
    origClientOrderId := "origClientOrderId_example" // string | Order id from client (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderDelete(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderDelete`: MarginOrder
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **orderId** | **int64** | Order id | 
 **origClientOrderId** | **string** | Order id from client | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**MarginOrder**](MarginOrder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderGet

> MarginOrderDetail SapiV1MarginOrderGet(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's Order (USER_DATA)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    orderId := int64(789) // int64 | Order id (optional)
    origClientOrderId := "origClientOrderId_example" // string | Order id from client (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderGet(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderGet`: MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **orderId** | **int64** | Order id | 
 **origClientOrderId** | **string** | Order id from client | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**MarginOrderDetail**](MarginOrderDetail.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderListDelete

> MarginOcoOrder SapiV1MarginOrderListDelete(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel OCO (TRADE)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    orderListId := int64(789) // int64 | Order list id (optional)
    listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderListDelete(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderListDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderListDelete`: MarginOcoOrder
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderListDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderListDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **orderListId** | **int64** | Order list id | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**MarginOcoOrder**](MarginOcoOrder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderListGet

> InlineResponse20021 SapiV1MarginOrderListGet(ctx).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)



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
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    symbol := "symbol_example" // string | Mandatory for isolated margin, not supported for cross margin (optional)
    orderListId := int64(789) // int64 | Order list id (optional)
    origClientOrderId := "origClientOrderId_example" // string | Order id from client (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderListGet(context.Background()).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderListGet`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **symbol** | **string** | Mandatory for isolated margin, not supported for cross margin | 
 **orderListId** | **int64** | Order list id | 
 **origClientOrderId** | **string** | Order id from client | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderOcoPost

> InlineResponse20020 SapiV1MarginOrderOcoPost(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).RecvWindow(recvWindow).Execute()

Margin Account New OCO (TRADE)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    side := "SELL" // string | SELL or BUY
    quantity := float64(1.2) // float64 | 
    price := float64(1.2) // float64 | Order price
    stopPrice := float64(1.2) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
    limitClientOrderId := "limitClientOrderId_example" // string | A unique Id for the limit order (optional)
    limitIcebergQty := float64(1.2) // float64 |  (optional)
    stopClientOrderId := "stopClientOrderId_example" // string | A unique Id for the stop loss/stop loss limit leg (optional)
    stopLimitPrice := float64(1.2) // float64 | If provided, stopLimitTimeInForce is required. (optional)
    stopIcebergQty := float64(1.2) // float64 |  (optional)
    stopLimitTimeInForce := "stopLimitTimeInForce_example" // string |  (optional)
    newOrderRespType := "newOrderRespType_example" // string | Set the response JSON. (optional)
    sideEffectType := "sideEffectType_example" // string | Default NO_SIDE_EFFECT (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderOcoPost(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderOcoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderOcoPost`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderOcoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderOcoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **quantity** | **float64** |  | 
 **price** | **float64** | Order price | 
 **stopPrice** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **limitClientOrderId** | **string** | A unique Id for the limit order | 
 **limitIcebergQty** | **float64** |  | 
 **stopClientOrderId** | **string** | A unique Id for the stop loss/stop loss limit leg | 
 **stopLimitPrice** | **float64** | If provided, stopLimitTimeInForce is required. | 
 **stopIcebergQty** | **float64** |  | 
 **stopLimitTimeInForce** | **string** |  | 
 **newOrderRespType** | **string** | Set the response JSON. | 
 **sideEffectType** | **string** | Default NO_SIDE_EFFECT | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginOrderPost

> OneOfmarginOrderResponseAckmarginOrderResponseResultmarginOrderResponseFull SapiV1MarginOrderPost(ctx).Symbol(symbol).Side(side).Type_(type_).Quantity(quantity).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).RecvWindow(recvWindow).Execute()

Margin Account New Order (TRADE)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT
    side := "SELL" // string | SELL or BUY
    type_ := "type__example" // string | Order type
    quantity := float64(1.2) // float64 | 
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    quoteOrderQty := float64(1.2) // float64 | Quote quantity (optional)
    price := float64(1.2) // float64 | Order price (optional)
    stopPrice := float64(20.01) // float64 | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    icebergQty := float64(1.2) // float64 | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. (optional)
    newOrderRespType := "newOrderRespType_example" // string | Set the response JSON. (optional)
    sideEffectType := "sideEffectType_example" // string | Default NO_SIDE_EFFECT (optional)
    timeInForce := "timeInForce_example" // string | Order time in force (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginOrderPost(context.Background()).Symbol(symbol).Side(side).Type_(type_).Quantity(quantity).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginOrderPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginOrderPost`: OneOfmarginOrderResponseAckmarginOrderResponseResultmarginOrderResponseFull
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginOrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginOrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **type_** | **string** | Order type | 
 **quantity** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **quoteOrderQty** | **float64** | Quote quantity | 
 **price** | **float64** | Order price | 
 **stopPrice** | **float64** | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **icebergQty** | **float64** | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. | 
 **newOrderRespType** | **string** | Set the response JSON. | 
 **sideEffectType** | **string** | Default NO_SIDE_EFFECT | 
 **timeInForce** | **string** | Order time in force | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOfmarginOrderResponseAckmarginOrderResponseResultmarginOrderResponseFull**](oneOf&lt;marginOrderResponseAck,marginOrderResponseResult,marginOrderResponseFull&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginPairGet

> InlineResponse20013 SapiV1MarginPairGet(ctx).Symbol(symbol).Execute()

Query Cross Margin Pair (MARKET_DATA)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginPairGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginPairGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginPairGet`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginPairGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginPairGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginPriceIndexGet

> InlineResponse20016 SapiV1MarginPriceIndexGet(ctx).Symbol(symbol).Execute()

Query Margin PriceIndex (MARKET_DATA)



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginPriceIndexGet(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginPriceIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginPriceIndexGet`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginPriceIndexGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginPriceIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 

### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginRepayGet

> InlineResponse20011 SapiV1MarginRepayGet(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Query Repay Record (USER_DATA)



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
    isolatedSymbol := "isolatedSymbol_example" // string | Isolated symbol (optional)
    txId := int64(2970933056) // int64 | the tranId in  `POST /sapi/v1/margin/repay` (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    archived := "archived_example" // string | Default: false. Set to true for archived data from 6 months ago (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginRepayGet(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginRepayGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginRepayGet`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginRepayGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginRepayGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isolatedSymbol** | **string** | Isolated symbol | 
 **txId** | **int64** | the tranId in  &#x60;POST /sapi/v1/margin/repay&#x60; | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **archived** | **string** | Default: false. Set to true for archived data from 6 months ago | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginRepayPost

> Transaction SapiV1MarginRepayPost(ctx).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Margin Account Repay (MARGIN)



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
    isIsolated := "isIsolated_example" // string | For isolated margin or not, 'TRUE', 'FALSE', default 'FALSE' (optional)
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginRepayPost(context.Background()).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginRepayPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginRepayPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginRepayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginRepayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **isIsolated** | **string** | For isolated margin or not, &#39;TRUE&#39;, &#39;FALSE&#39;, default &#39;FALSE&#39; | 
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginTransferGet

> InlineResponse2009 SapiV1MarginTransferGet(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Get Cross Margin Transfer History (USER_DATA)



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
    type_ := "type__example" // string | Tranfer Type (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    current := int32(1) // int32 | Current querying page. Start from 1. Default:1 (optional)
    size := int32(100) // int32 | Default:10 Max:100 (optional)
    archived := "archived_example" // string | Default: false. Set to true for archived data from 6 months ago (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginTransferGet(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginTransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginTransferGet`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginTransferGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginTransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **asset** | **string** |  | 
 **type_** | **string** | Tranfer Type | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **current** | **int32** | Current querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10 Max:100 | 
 **archived** | **string** | Default: false. Set to true for archived data from 6 months ago | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SapiV1MarginTransferPost

> Transaction SapiV1MarginTransferPost(ctx).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).Type_(type_).RecvWindow(recvWindow).Execute()

Cross Margin Account Transfer (MARGIN)



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
    type_ := int32(56) // int32 | 1 -> transfer from main account to margin account \\ 2 -> transfer from margin account to main account (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarginApi.SapiV1MarginTransferPost(context.Background()).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).Type_(type_).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.SapiV1MarginTransferPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1MarginTransferPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.SapiV1MarginTransferPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1MarginTransferPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **type_** | **int32** | 1 -&gt; transfer from main account to margin account \\ 2 -&gt; transfer from margin account to main account | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

