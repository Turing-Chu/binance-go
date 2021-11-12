# \MarginApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableBnbBurn**](MarginApi.md#EnableBnbBurn) | **Post** /sapi/v1/bnbBurn | Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
[**GetBnbBurnStatus**](MarginApi.md#GetBnbBurnStatus) | **Get** /sapi/v1/bnbBurn | Get All Isolated Margin Symbol(USER_DATA)
[**MarginAccountRepay**](MarginApi.md#MarginAccountRepay) | **Post** /sapi/v1/margin/repay | Margin Account Repay (MARGIN)
[**MarginApplyLoan**](MarginApi.md#MarginApplyLoan) | **Post** /sapi/v1/margin/loan | Margin Account Borrow (MARGIN)
[**MarginCancelAllSymbolOrders**](MarginApi.md#MarginCancelAllSymbolOrders) | **Delete** /sapi/v1/margin/openOrders | Margin Account Cancel all Open Orders on a Symbol (TRADE)
[**MarginCancelOcoOrders**](MarginApi.md#MarginCancelOcoOrders) | **Delete** /sapi/v1/margin/orderList | Margin Account Cancel OCO (TRADE)
[**MarginCancelOrder**](MarginApi.md#MarginCancelOrder) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**MarginCreateOcoOrder**](MarginApi.md#MarginCreateOcoOrder) | **Post** /sapi/v1/margin/order/oco | Margin Account New OCO (TRADE)
[**MarginCreateOrder**](MarginApi.md#MarginCreateOrder) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**MarginDisableAccount**](MarginApi.md#MarginDisableAccount) | **Delete** /sapi/v1/margin/isolated/account | Disable Isolated Margin Account (TRADE)
[**MarginEnableAccount**](MarginApi.md#MarginEnableAccount) | **Post** /sapi/v1/margin/isolated/account | Enable Isolated Margin Account (TRADE)
[**MarginExecuteTransfer**](MarginApi.md#MarginExecuteTransfer) | **Post** /sapi/v1/margin/transfer | Cross Margin Account Transfer (MARGIN)
[**MarginGetAccount**](MarginApi.md#MarginGetAccount) | **Get** /sapi/v1/margin/account | Query Cross Margin Account Details (USER_DATA)
[**MarginGetAccountInfo**](MarginApi.md#MarginGetAccountInfo) | **Get** /sapi/v1/margin/isolated/account | Query Isolated Margin Account Info (USER_DATA)
[**MarginGetAccountLimit**](MarginApi.md#MarginGetAccountLimit) | **Get** /sapi/v1/margin/isolated/accountLimit | Query Enabled Isolated Margin Account Limit (USER_DATA)
[**MarginGetAllAssets**](MarginApi.md#MarginGetAllAssets) | **Get** /sapi/v1/margin/allAssets | Get All Margin Assets (MARKET_DATA)
[**MarginGetAllCrossPairs**](MarginApi.md#MarginGetAllCrossPairs) | **Get** /sapi/v1/margin/allPairs | Get All Cross Margin Pairs (MARKET_DATA)
[**MarginGetAllOcoOrders**](MarginApi.md#MarginGetAllOcoOrders) | **Get** /sapi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**MarginGetAllOrders**](MarginApi.md#MarginGetAllOrders) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#39;s All Orders (USER_DATA)
[**MarginGetAllPairs**](MarginApi.md#MarginGetAllPairs) | **Get** /sapi/v1/margin/isolated/allPairs | Get All Isolated Margin Symbol(USER_DATA)
[**MarginGetAsset**](MarginApi.md#MarginGetAsset) | **Get** /sapi/v1/margin/asset | Query Margin Asset (MARKET_DATA)
[**MarginGetCrossTransferRecord**](MarginApi.md#MarginGetCrossTransferRecord) | **Get** /sapi/v1/margin/transfer | Get Cross Margin Transfer History (USER_DATA)
[**MarginGetForceLiquidationRecord**](MarginApi.md#MarginGetForceLiquidationRecord) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**MarginGetInterestRateRecord**](MarginApi.md#MarginGetInterestRateRecord) | **Get** /sapi/v1/margin/interestRateHistory | Margin Interest Rate History (USER_DATA)
[**MarginGetInterestRecord**](MarginApi.md#MarginGetInterestRecord) | **Get** /sapi/v1/margin/interestHistory | Get Interest History (USER_DATA)
[**MarginGetLoan**](MarginApi.md#MarginGetLoan) | **Get** /sapi/v1/margin/loan | Query Loan Record (USER_DATA)
[**MarginGetMaxBorrowable**](MarginApi.md#MarginGetMaxBorrowable) | **Get** /sapi/v1/margin/maxBorrowable | Query Max Borrow (USER_DATA)
[**MarginGetMaxTransferable**](MarginApi.md#MarginGetMaxTransferable) | **Get** /sapi/v1/margin/maxTransferable | Query Max Transfer-Out Amount (USER_DATA)
[**MarginGetOcoOrders**](MarginApi.md#MarginGetOcoOrders) | **Get** /sapi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**MarginGetOrder**](MarginApi.md#MarginGetOrder) | **Get** /sapi/v1/margin/order | Query Margin Account&#39;s Order (USER_DATA)
[**MarginGetOrders**](MarginApi.md#MarginGetOrders) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Orders (USER_DATA)
[**MarginGetPair**](MarginApi.md#MarginGetPair) | **Get** /sapi/v1/margin/isolated/pair | Query Isolated Margin Symbol (USER_DATA)
[**MarginGetPairs**](MarginApi.md#MarginGetPairs) | **Get** /sapi/v1/margin/pair | Query Cross Margin Pair (MARKET_DATA)
[**MarginGetPriceIndex**](MarginApi.md#MarginGetPriceIndex) | **Get** /sapi/v1/margin/priceIndex | Query Margin PriceIndex (MARKET_DATA)
[**MarginGetRepayRecord**](MarginApi.md#MarginGetRepayRecord) | **Get** /sapi/v1/margin/repay | Query Repay Record (USER_DATA)
[**MarginGetSpecificOcoOrders**](MarginApi.md#MarginGetSpecificOcoOrders) | **Get** /sapi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**MarginGetTrades**](MarginApi.md#MarginGetTrades) | **Get** /sapi/v1/margin/myTrades | Query Margin Account&#39;s Trade List (USER_DATA)
[**MarginGetTransferRecord**](MarginApi.md#MarginGetTransferRecord) | **Get** /sapi/v1/margin/isolated/transfer | Get Isolated Margin Transfer History (USER_DATA)
[**MarginTransfer**](MarginApi.md#MarginTransfer) | **Post** /sapi/v1/margin/isolated/transfer | Isolated Margin Account Transfer (MARGIN)



## EnableBnbBurn

> BnbBurnStatus EnableBnbBurn(ctx).Timestamp(timestamp).Signature(signature).SpotBNBBurn(spotBNBBurn).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.EnableBnbBurn(context.Background()).Timestamp(timestamp).Signature(signature).SpotBNBBurn(spotBNBBurn).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.EnableBnbBurn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableBnbBurn`: BnbBurnStatus
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.EnableBnbBurn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableBnbBurnRequest struct via the builder pattern


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


## GetBnbBurnStatus

> BnbBurnStatus GetBnbBurnStatus(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.GetBnbBurnStatus(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.GetBnbBurnStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBnbBurnStatus`: BnbBurnStatus
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.GetBnbBurnStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBnbBurnStatusRequest struct via the builder pattern


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


## MarginAccountRepay

> Transaction MarginAccountRepay(ctx).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginAccountRepay(context.Background()).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginAccountRepay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginAccountRepay`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginAccountRepay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginAccountRepayRequest struct via the builder pattern


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


## MarginApplyLoan

> Transaction MarginApplyLoan(ctx).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginApplyLoan(context.Background()).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginApplyLoan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginApplyLoan`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginApplyLoan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginApplyLoanRequest struct via the builder pattern


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


## MarginCancelAllSymbolOrders

> []AnyOfcanceledMarginOrderDetailmarginOcoOrder MarginCancelAllSymbolOrders(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginCancelAllSymbolOrders(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginCancelAllSymbolOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginCancelAllSymbolOrders`: []AnyOfcanceledMarginOrderDetailmarginOcoOrder
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginCancelAllSymbolOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCancelAllSymbolOrdersRequest struct via the builder pattern


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


## MarginCancelOcoOrders

> MarginOcoOrder MarginCancelOcoOrders(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginCancelOcoOrders(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginCancelOcoOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginCancelOcoOrders`: MarginOcoOrder
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginCancelOcoOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCancelOcoOrdersRequest struct via the builder pattern


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


## MarginCancelOrder

> MarginOrder MarginCancelOrder(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginCancelOrder(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginCancelOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginCancelOrder`: MarginOrder
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginCancelOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCancelOrderRequest struct via the builder pattern


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


## MarginCreateOcoOrder

> InlineResponse20020 MarginCreateOcoOrder(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginCreateOcoOrder(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginCreateOcoOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginCreateOcoOrder`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginCreateOcoOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateOcoOrderRequest struct via the builder pattern


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


## MarginCreateOrder

> OneOfmarginOrderResponseAckmarginOrderResponseResultmarginOrderResponseFull MarginCreateOrder(ctx).Symbol(symbol).Side(side).Type_(type_).Quantity(quantity).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginCreateOrder(context.Background()).Symbol(symbol).Side(side).Type_(type_).Quantity(quantity).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginCreateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginCreateOrder`: OneOfmarginOrderResponseAckmarginOrderResponseResultmarginOrderResponseFull
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginCreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateOrderRequest struct via the builder pattern


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


## MarginDisableAccount

> InlineResponse20025 MarginDisableAccount(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginDisableAccount(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginDisableAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginDisableAccount`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginDisableAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDisableAccountRequest struct via the builder pattern


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


## MarginEnableAccount

> InlineResponse20025 MarginEnableAccount(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginEnableAccount(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginEnableAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginEnableAccount`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginEnableAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginEnableAccountRequest struct via the builder pattern


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


## MarginExecuteTransfer

> Transaction MarginExecuteTransfer(ctx).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).Type_(type_).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginExecuteTransfer(context.Background()).Asset(asset).Amount(amount).Timestamp(timestamp).Signature(signature).Type_(type_).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginExecuteTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginExecuteTransfer`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginExecuteTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginExecuteTransferRequest struct via the builder pattern


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


## MarginGetAccount

> InlineResponse20019 MarginGetAccount(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAccount(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAccount`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAccountRequest struct via the builder pattern


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


## MarginGetAccountInfo

> IsolatedMarginAccountInfo MarginGetAccountInfo(ctx).Timestamp(timestamp).Signature(signature).Symbols(symbols).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAccountInfo(context.Background()).Timestamp(timestamp).Signature(signature).Symbols(symbols).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAccountInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAccountInfo`: IsolatedMarginAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAccountInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAccountInfoRequest struct via the builder pattern


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


## MarginGetAccountLimit

> InlineResponse20026 MarginGetAccountLimit(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAccountLimit(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAccountLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAccountLimit`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAccountLimit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAccountLimitRequest struct via the builder pattern


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


## MarginGetAllAssets

> []InlineResponse20014 MarginGetAllAssets(ctx).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAllAssets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAllAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAllAssets`: []InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAllAssets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAllAssetsRequest struct via the builder pattern


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


## MarginGetAllCrossPairs

> []InlineResponse20015 MarginGetAllCrossPairs(ctx).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAllCrossPairs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAllCrossPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAllCrossPairs`: []InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAllCrossPairs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAllCrossPairsRequest struct via the builder pattern


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


## MarginGetAllOcoOrders

> []InlineResponse2006 MarginGetAllOcoOrders(ctx).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAllOcoOrders(context.Background()).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAllOcoOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAllOcoOrders`: []InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAllOcoOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAllOcoOrdersRequest struct via the builder pattern


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


## MarginGetAllOrders

> []MarginOrderDetail MarginGetAllOrders(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAllOrders(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAllOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAllOrders`: []MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAllOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAllOrdersRequest struct via the builder pattern


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


## MarginGetAllPairs

> []InlineResponse20027 MarginGetAllPairs(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAllPairs(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAllPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAllPairs`: []InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAllPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAllPairsRequest struct via the builder pattern


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


## MarginGetAsset

> InlineResponse20012 MarginGetAsset(ctx).Asset(asset).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetAsset(context.Background()).Asset(asset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetAsset`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetAssetRequest struct via the builder pattern


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


## MarginGetCrossTransferRecord

> InlineResponse2009 MarginGetCrossTransferRecord(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetCrossTransferRecord(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetCrossTransferRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetCrossTransferRecord`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetCrossTransferRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetCrossTransferRecordRequest struct via the builder pattern


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


## MarginGetForceLiquidationRecord

> InlineResponse20018 MarginGetForceLiquidationRecord(ctx).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetForceLiquidationRecord(context.Background()).Timestamp(timestamp).Signature(signature).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetForceLiquidationRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetForceLiquidationRecord`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetForceLiquidationRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetForceLiquidationRecordRequest struct via the builder pattern


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


## MarginGetInterestRateRecord

> []InlineResponse20028 MarginGetInterestRateRecord(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetInterestRateRecord(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetInterestRateRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetInterestRateRecord`: []InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetInterestRateRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetInterestRateRecordRequest struct via the builder pattern


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


## MarginGetInterestRecord

> InlineResponse20017 MarginGetInterestRecord(ctx).Timestamp(timestamp).Signature(signature).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetInterestRecord(context.Background()).Timestamp(timestamp).Signature(signature).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetInterestRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetInterestRecord`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetInterestRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetInterestRecordRequest struct via the builder pattern


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


## MarginGetLoan

> InlineResponse20010 MarginGetLoan(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetLoan(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetLoan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetLoan`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetLoan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetLoanRequest struct via the builder pattern


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


## MarginGetMaxBorrowable

> InlineResponse20023 MarginGetMaxBorrowable(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetMaxBorrowable(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetMaxBorrowable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetMaxBorrowable`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetMaxBorrowable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMaxBorrowableRequest struct via the builder pattern


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


## MarginGetMaxTransferable

> InlineResponse20024 MarginGetMaxTransferable(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetMaxTransferable(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetMaxTransferable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetMaxTransferable`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetMaxTransferable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMaxTransferableRequest struct via the builder pattern


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


## MarginGetOcoOrders

> []InlineResponse20022 MarginGetOcoOrders(ctx).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetOcoOrders(context.Background()).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetOcoOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetOcoOrders`: []InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetOcoOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetOcoOrdersRequest struct via the builder pattern


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


## MarginGetOrder

> MarginOrderDetail MarginGetOrder(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetOrder(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetOrder`: MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetOrderRequest struct via the builder pattern


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


## MarginGetOrders

> []MarginOrderDetail MarginGetOrders(ctx).Timestamp(timestamp).Signature(signature).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetOrders(context.Background()).Timestamp(timestamp).Signature(signature).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetOrders`: []MarginOrderDetail
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetOrdersRequest struct via the builder pattern


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


## MarginGetPair

> InlineResponse20027 MarginGetPair(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetPair(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetPair`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetPairRequest struct via the builder pattern


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


## MarginGetPairs

> InlineResponse20013 MarginGetPairs(ctx).Symbol(symbol).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetPairs(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetPairs`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetPairsRequest struct via the builder pattern


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


## MarginGetPriceIndex

> InlineResponse20016 MarginGetPriceIndex(ctx).Symbol(symbol).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetPriceIndex(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetPriceIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetPriceIndex`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetPriceIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetPriceIndexRequest struct via the builder pattern


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


## MarginGetRepayRecord

> InlineResponse20011 MarginGetRepayRecord(ctx).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetRepayRecord(context.Background()).Asset(asset).Timestamp(timestamp).Signature(signature).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetRepayRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetRepayRecord`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetRepayRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetRepayRecordRequest struct via the builder pattern


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


## MarginGetSpecificOcoOrders

> InlineResponse20021 MarginGetSpecificOcoOrders(ctx).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetSpecificOcoOrders(context.Background()).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetSpecificOcoOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetSpecificOcoOrders`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetSpecificOcoOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetSpecificOcoOrdersRequest struct via the builder pattern


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


## MarginGetTrades

> []MarginTrade MarginGetTrades(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetTrades(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).IsIsolated(isIsolated).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetTrades`: []MarginTrade
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetTrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetTradesRequest struct via the builder pattern


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


## MarginGetTransferRecord

> MarginTransferDetails MarginGetTransferRecord(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).Asset(asset).TransFrom(transFrom).TransTo(transTo).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginGetTransferRecord(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).Asset(asset).TransFrom(transFrom).TransTo(transTo).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginGetTransferRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginGetTransferRecord`: MarginTransferDetails
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginGetTransferRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetTransferRecordRequest struct via the builder pattern


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


## MarginTransfer

> map[string]interface{} MarginTransfer(ctx).Asset(asset).Symbol(symbol).Amount(amount).Timestamp(timestamp).Signature(signature).TransFrom(transFrom).TransTo(transTo).RecvWindow(recvWindow).Execute()

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
    resp, r, err := api_client.MarginApi.MarginTransfer(context.Background()).Asset(asset).Symbol(symbol).Amount(amount).Timestamp(timestamp).Signature(signature).TransFrom(transFrom).TransTo(transTo).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarginApi.MarginTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarginTransfer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarginApi.MarginTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginTransferRequest struct via the builder pattern


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

