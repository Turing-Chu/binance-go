# \MarketApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Depth**](MarketApi.md#Depth) | **Get** /api/v3/depth | Order Book
[**GetAggregateTrades**](MarketApi.md#GetAggregateTrades) | **Get** /api/v3/aggTrades | Compressed/Aggregate Trades List
[**GetAvgPrice**](MarketApi.md#GetAvgPrice) | **Get** /api/v3/avgPrice | Current Average Price
[**GetBookTicker**](MarketApi.md#GetBookTicker) | **Get** /api/v3/ticker/bookTicker | Symbol Order Book Ticker
[**GetCurrentServerTime**](MarketApi.md#GetCurrentServerTime) | **Get** /api/v3/time | Check Server Time
[**GetExchangeInfo**](MarketApi.md#GetExchangeInfo) | **Get** /api/v3/exchangeInfo | Exchange Information
[**GetHistoricalTrades**](MarketApi.md#GetHistoricalTrades) | **Get** /api/v3/historicalTrades | Old Trade Lookup
[**GetKlines**](MarketApi.md#GetKlines) | **Get** /api/v3/klines | Kline/Candlestick Data
[**GetRecentTrades**](MarketApi.md#GetRecentTrades) | **Get** /api/v3/trades | Recent Trades List
[**GetTicker24hrPrice**](MarketApi.md#GetTicker24hrPrice) | **Get** /api/v3/ticker/24hr | 24hr Ticker Price Change Statistics
[**GetTickerPrice**](MarketApi.md#GetTickerPrice) | **Get** /api/v3/ticker/price | Symbol Price Ticker
[**Ping**](MarketApi.md#Ping) | **Get** /api/v3/ping | Test Connectivity



## Depth

> InlineResponse2002 Depth(ctx).Symbol(symbol).Limit(limit).Execute()

Order Book



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
    limit := int32(100) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.Depth(context.Background()).Symbol(symbol).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.Depth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Depth`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.Depth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAggregateTrades

> []AggTrade GetAggregateTrades(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Compressed/Aggregate Trades List



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
    fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetAggregateTrades(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetAggregateTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregateTrades`: []AggTrade
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetAggregateTrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregateTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 

### Return type

[**[]AggTrade**](AggTrade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvgPrice

> InlineResponse2003 GetAvgPrice(ctx).Symbol(symbol).Execute()

Current Average Price



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
    resp, r, err := api_client.MarketApi.GetAvgPrice(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetAvgPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvgPrice`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetAvgPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvgPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookTicker

> OneOfbookTickerbookTickerList GetBookTicker(ctx).Symbol(symbol).Execute()

Symbol Order Book Ticker



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetBookTicker(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetBookTicker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookTicker`: OneOfbookTickerbookTickerList
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetBookTicker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBookTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 

### Return type

[**OneOfbookTickerbookTickerList**](oneOf&lt;bookTicker,bookTickerList&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentServerTime

> InlineResponse200 GetCurrentServerTime(ctx).Execute()

Check Server Time



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
    resp, r, err := api_client.MarketApi.GetCurrentServerTime(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetCurrentServerTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentServerTime`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetCurrentServerTime`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentServerTimeRequest struct via the builder pattern


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeInfo

> InlineResponse2001 GetExchangeInfo(ctx).Symbol(symbol).ArraySymbols(arraySymbols).Execute()

Exchange Information



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT (optional)
    arraySymbols := "["BTCUSDT","BNBBTC"]" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetExchangeInfo(context.Background()).Symbol(symbol).ArraySymbols(arraySymbols).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetExchangeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeInfo`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetExchangeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **arraySymbols** | **string** |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricalTrades

> []Trade GetHistoricalTrades(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

Old Trade Lookup



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
    limit := int32(500) // int32 | Default 500; max 1000. (optional)
    fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetHistoricalTrades(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetHistoricalTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistoricalTrades`: []Trade
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetHistoricalTrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **limit** | **int32** | Default 500; max 1000. | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 

### Return type

[**[]Trade**](Trade.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKlines

> [][]OneOflongstring GetKlines(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Kline/Candlestick Data



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
    interval := "interval_example" // string | kline intervals
    startTime := int64(789) // int64 | UTC timestamp in ms (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    limit := int32(500) // int32 | Default 500; max 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetKlines(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetKlines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKlines`: [][]OneOflongstring
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetKlines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKlinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **interval** | **string** | kline intervals | 
 **startTime** | **int64** | UTC timestamp in ms | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **limit** | **int32** | Default 500; max 1000. | 

### Return type

[**[][]OneOflongstring**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecentTrades

> []Trade GetRecentTrades(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Trades List



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
    limit := int32(500) // int32 | Default 500; max 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetRecentTrades(context.Background()).Symbol(symbol).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetRecentTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecentTrades`: []Trade
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetRecentTrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 
 **limit** | **int32** | Default 500; max 1000. | 

### Return type

[**[]Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicker24hrPrice

> OneOftickertickerList GetTicker24hrPrice(ctx).Symbol(symbol).Execute()

24hr Ticker Price Change Statistics



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetTicker24hrPrice(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetTicker24hrPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTicker24hrPrice`: OneOftickertickerList
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetTicker24hrPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTicker24hrPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 

### Return type

[**OneOftickertickerList**](oneOf&lt;ticker,tickerList&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTickerPrice

> OneOfpriceTickerpriceTickerList GetTickerPrice(ctx).Symbol(symbol).Execute()

Symbol Price Ticker



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
    symbol := "BNBUSDT" // string | Trading symbol, e.g. BNBUSDT (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.GetTickerPrice(context.Background()).Symbol(symbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.GetTickerPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTickerPrice`: OneOfpriceTickerpriceTickerList
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.GetTickerPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTickerPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol, e.g. BNBUSDT | 

### Return type

[**OneOfpriceTickerpriceTickerList**](oneOf&lt;priceTicker,priceTickerList&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> map[string]interface{} Ping(ctx).Execute()

Test Connectivity



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
    resp, r, err := api_client.MarketApi.Ping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.Ping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Ping`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

