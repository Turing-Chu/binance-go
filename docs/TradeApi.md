# \TradeApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3AccountGet**](TradeApi.md#ApiV3AccountGet) | **Get** /api/v3/account | Account Information (USER_DATA)
[**ApiV3AllOrderListGet**](TradeApi.md#ApiV3AllOrderListGet) | **Get** /api/v3/allOrderList | Query all OCO (USER_DATA)
[**ApiV3AllOrdersGet**](TradeApi.md#ApiV3AllOrdersGet) | **Get** /api/v3/allOrders | All Orders (USER_DATA)
[**ApiV3MyTradesGet**](TradeApi.md#ApiV3MyTradesGet) | **Get** /api/v3/myTrades | Account Trade List (USER_DATA)
[**ApiV3OpenOrderListGet**](TradeApi.md#ApiV3OpenOrderListGet) | **Get** /api/v3/openOrderList | Query Open OCO (USER_DATA)
[**ApiV3OpenOrdersDelete**](TradeApi.md#ApiV3OpenOrdersDelete) | **Delete** /api/v3/openOrders | Cancel all Open Orders on a Symbol (TRADE)
[**ApiV3OpenOrdersGet**](TradeApi.md#ApiV3OpenOrdersGet) | **Get** /api/v3/openOrders | Current Open Orders (USER_DATA)
[**ApiV3OrderDelete**](TradeApi.md#ApiV3OrderDelete) | **Delete** /api/v3/order | Cancel Order (TRADE)
[**ApiV3OrderGet**](TradeApi.md#ApiV3OrderGet) | **Get** /api/v3/order | Query Order (USER_DATA)
[**ApiV3OrderListDelete**](TradeApi.md#ApiV3OrderListDelete) | **Delete** /api/v3/orderList | Cancel OCO (TRADE)
[**ApiV3OrderListGet**](TradeApi.md#ApiV3OrderListGet) | **Get** /api/v3/orderList | Query OCO (USER_DATA)
[**ApiV3OrderOcoPost**](TradeApi.md#ApiV3OrderOcoPost) | **Post** /api/v3/order/oco | New OCO (TRADE)
[**ApiV3OrderPost**](TradeApi.md#ApiV3OrderPost) | **Post** /api/v3/order | New Order (TRADE)
[**ApiV3OrderTestPost**](TradeApi.md#ApiV3OrderTestPost) | **Post** /api/v3/order/test | Test New Order (TRADE)
[**ApiV3RateLimitOrderGet**](TradeApi.md#ApiV3RateLimitOrderGet) | **Get** /api/v3/rateLimit/order | Query Current Order Count Usage (TRADE)



## ApiV3AccountGet

> Account ApiV3AccountGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Account Information (USER_DATA)



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
    resp, r, err := api_client.TradeApi.ApiV3AccountGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3AccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3AccountGet`: Account
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3AccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3AccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**Account**](Account.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3AllOrderListGet

> []InlineResponse2006 ApiV3AllOrderListGet(ctx).Timestamp(timestamp).Signature(signature).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query all OCO (USER_DATA)



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
    fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3AllOrderListGet(context.Background()).Timestamp(timestamp).Signature(signature).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3AllOrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3AllOrderListGet`: []InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3AllOrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3AllOrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
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


## ApiV3AllOrdersGet

> []OrderDetails ApiV3AllOrdersGet(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

All Orders (USER_DATA)



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
    orderId := int64(789) // int64 | Order id (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3AllOrdersGet(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3AllOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3AllOrdersGet`: []OrderDetails
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3AllOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3AllOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **orderId** | **int64** | Order id | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]OrderDetails**](OrderDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3MyTradesGet

> []MyTrade ApiV3MyTradesGet(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)



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
    orderId := int64(789) // int64 | This can only be used in combination with symbol. (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3MyTradesGet(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3MyTradesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3MyTradesGet`: []MyTrade
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3MyTradesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3MyTradesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **orderId** | **int64** | This can only be used in combination with symbol. | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]MyTrade**](MyTrade.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OpenOrderListGet

> []InlineResponse2007 ApiV3OpenOrderListGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Query Open OCO (USER_DATA)



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
    resp, r, err := api_client.TradeApi.ApiV3OpenOrderListGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OpenOrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OpenOrderListGet`: []InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OpenOrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OpenOrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse2007**](InlineResponse2007.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OpenOrdersDelete

> []AnyOforderocoOrder ApiV3OpenOrdersDelete(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Cancel all Open Orders on a Symbol (TRADE)



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
    resp, r, err := api_client.TradeApi.ApiV3OpenOrdersDelete(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OpenOrdersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OpenOrdersDelete`: []AnyOforderocoOrder
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OpenOrdersDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OpenOrdersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]AnyOforderocoOrder**](anyOf&lt;order,ocoOrder&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OpenOrdersGet

> []OrderDetails ApiV3OpenOrdersGet(ctx).Timestamp(timestamp).Signature(signature).Symbol(symbol).RecvWindow(recvWindow).Execute()

Current Open Orders (USER_DATA)



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
    resp, r, err := api_client.TradeApi.ApiV3OpenOrdersGet(context.Background()).Timestamp(timestamp).Signature(signature).Symbol(symbol).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OpenOrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OpenOrdersGet`: []OrderDetails
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OpenOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OpenOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]OrderDetails**](OrderDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderDelete

> Order ApiV3OrderDelete(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Order (TRADE)



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
    orderId := int64(789) // int64 | Order id (optional)
    origClientOrderId := "origClientOrderId_example" // string | Order id from client (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderDelete(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderDelete`: Order
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **orderId** | **int64** | Order id | 
 **origClientOrderId** | **string** | Order id from client | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**Order**](Order.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderGet

> OrderDetails ApiV3OrderGet(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Order (USER_DATA)



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
    orderId := int64(789) // int64 | Order id (optional)
    origClientOrderId := "origClientOrderId_example" // string | Order id from client (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderGet(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderGet`: OrderDetails
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **orderId** | **int64** | Order id | 
 **origClientOrderId** | **string** | Order id from client | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OrderDetails**](OrderDetails.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderListDelete

> OcoOrder ApiV3OrderListDelete(ctx).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel OCO (TRADE)



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
    orderListId := int64(789) // int64 | Order list id (optional)
    listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderListDelete(context.Background()).Symbol(symbol).Timestamp(timestamp).Signature(signature).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderListDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderListDelete`: OcoOrder
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderListDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderListDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **orderListId** | **int64** | Order list id | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OcoOrder**](OcoOrder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderListGet

> InlineResponse2005 ApiV3OrderListGet(ctx).Timestamp(timestamp).Signature(signature).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query OCO (USER_DATA)



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
    orderListId := int64(789) // int64 | Order list id (optional)
    origClientOrderId := "origClientOrderId_example" // string | Order id from client (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderListGet(context.Background()).Timestamp(timestamp).Signature(signature).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderListGet`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **orderListId** | **int64** | Order list id | 
 **origClientOrderId** | **string** | Order id from client | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderOcoPost

> InlineResponse2004 ApiV3OrderOcoPost(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).Timestamp(timestamp).Signature(signature).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).RecvWindow(recvWindow).Execute()

New OCO (TRADE)



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
    resp, r, err := api_client.TradeApi.ApiV3OrderOcoPost(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).Timestamp(timestamp).Signature(signature).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderOcoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderOcoPost`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderOcoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderOcoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **quantity** | **float64** |  | 
 **price** | **float64** | Order price | 
 **stopPrice** | **float64** |  | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
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

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderPost

> OneOforderResponseAckorderResponseResultorderResponseFull ApiV3OrderPost(ctx).Symbol(symbol).Side(side).Type_(type_).Timestamp(timestamp).Signature(signature).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Execute()

New Order (TRADE)



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
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    timeInForce := "timeInForce_example" // string | Order time in force (optional)
    quantity := float64(1.2) // float64 | Order quantity (optional)
    quoteOrderQty := float64(1.2) // float64 | Quote quantity (optional)
    price := float64(1.2) // float64 | Order price (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    stopPrice := float64(20.01) // float64 | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. (optional)
    icebergQty := float64(1.2) // float64 | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. (optional)
    newOrderRespType := "newOrderRespType_example" // string | Set the response JSON. MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderPost(context.Background()).Symbol(symbol).Side(side).Type_(type_).Timestamp(timestamp).Signature(signature).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderPost`: OneOforderResponseAckorderResponseResultorderResponseFull
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **type_** | **string** | Order type | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **timeInForce** | **string** | Order time in force | 
 **quantity** | **float64** | Order quantity | 
 **quoteOrderQty** | **float64** | Quote quantity | 
 **price** | **float64** | Order price | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **stopPrice** | **float64** | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. | 
 **icebergQty** | **float64** | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. | 
 **newOrderRespType** | **string** | Set the response JSON. MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OneOforderResponseAckorderResponseResultorderResponseFull**](oneOf&lt;orderResponseAck,orderResponseResult,orderResponseFull&gt;.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3OrderTestPost

> map[string]interface{} ApiV3OrderTestPost(ctx).Symbol(symbol).Side(side).Type_(type_).Timestamp(timestamp).Signature(signature).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Execute()

Test New Order (TRADE)



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
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    timeInForce := "timeInForce_example" // string | Order time in force (optional)
    quantity := float64(1.2) // float64 | Order quantity (optional)
    quoteOrderQty := float64(1.2) // float64 | Quote quantity (optional)
    price := float64(1.2) // float64 | Order price (optional)
    newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
    stopPrice := float64(20.01) // float64 | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. (optional)
    icebergQty := float64(1.2) // float64 | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. (optional)
    newOrderRespType := "newOrderRespType_example" // string | Set the response JSON. MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TradeApi.ApiV3OrderTestPost(context.Background()).Symbol(symbol).Side(side).Type_(type_).Timestamp(timestamp).Signature(signature).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3OrderTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3OrderTestPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3OrderTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3OrderTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **side** | **string** | SELL or BUY | 
 **type_** | **string** | Order type | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **timeInForce** | **string** | Order time in force | 
 **quantity** | **float64** | Order quantity | 
 **quoteOrderQty** | **float64** | Quote quantity | 
 **price** | **float64** | Order price | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **stopPrice** | **float64** | Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. | 
 **icebergQty** | **float64** | Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order. | 
 **newOrderRespType** | **string** | Set the response JSON. MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
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


## ApiV3RateLimitOrderGet

> []InlineResponse2008 ApiV3RateLimitOrderGet(ctx).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()

Query Current Order Count Usage (TRADE)



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
    resp, r, err := api_client.TradeApi.ApiV3RateLimitOrderGet(context.Background()).Timestamp(timestamp).Signature(signature).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeApi.ApiV3RateLimitOrderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3RateLimitOrderGet`: []InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `TradeApi.ApiV3RateLimitOrderGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3RateLimitOrderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]InlineResponse2008**](InlineResponse2008.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

