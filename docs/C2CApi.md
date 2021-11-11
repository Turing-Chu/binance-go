# \C2CApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SapiV1C2cOrderMatchListUserOrderHistoryGet**](C2CApi.md#SapiV1C2cOrderMatchListUserOrderHistoryGet) | **Get** /sapi/v1/c2c/orderMatch/listUserOrderHistory | Get C2C Trade History (USER_DATA)



## SapiV1C2cOrderMatchListUserOrderHistoryGet

> InlineResponse200111 SapiV1C2cOrderMatchListUserOrderHistoryGet(ctx).Timestamp(timestamp).Signature(signature).TradeType(tradeType).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()

Get C2C Trade History (USER_DATA)



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
    tradeType := "tradeType_example" // string |  (optional)
    startTimestamp := int64(789) // int64 | UTC timestamp in ms (optional)
    endTimestamp := int64(789) // int64 | UTC timestamp in ms (optional)
    page := int32(1) // int32 | Default 1 (optional)
    rows := int32(56) // int32 | default 100, max 100 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.C2CApi.SapiV1C2cOrderMatchListUserOrderHistoryGet(context.Background()).Timestamp(timestamp).Signature(signature).TradeType(tradeType).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `C2CApi.SapiV1C2cOrderMatchListUserOrderHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SapiV1C2cOrderMatchListUserOrderHistoryGet`: InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `C2CApi.SapiV1C2cOrderMatchListUserOrderHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSapiV1C2cOrderMatchListUserOrderHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **tradeType** | **string** |  | 
 **startTimestamp** | **int64** | UTC timestamp in ms | 
 **endTimestamp** | **int64** | UTC timestamp in ms | 
 **page** | **int32** | Default 1 | 
 **rows** | **int32** | default 100, max 100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

