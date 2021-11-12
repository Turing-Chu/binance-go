# \FiatApi

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FiatGetHistoryOrders**](FiatApi.md#FiatGetHistoryOrders) | **Get** /sapi/v1/fiat/orders | Fiat Deposit/Withdraw History (USER_DATA)
[**FiatGetPaymentsHistory**](FiatApi.md#FiatGetPaymentsHistory) | **Get** /sapi/v1/fiat/payments | Fiat Payments History (USER_DATA)



## FiatGetHistoryOrders

> InlineResponse20072 FiatGetHistoryOrders(ctx).TransactionType(transactionType).Timestamp(timestamp).Signature(signature).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()

Fiat Deposit/Withdraw History (USER_DATA)



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
    transactionType := "0" // string | 0-deposit, 1-withdraw
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    beginTime := int64(1626144956000) // int64 |  (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    page := int32(1) // int32 | Default 1 (optional)
    rows := int32(300) // int32 | Default 100, max 500 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FiatApi.FiatGetHistoryOrders(context.Background()).TransactionType(transactionType).Timestamp(timestamp).Signature(signature).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiatApi.FiatGetHistoryOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FiatGetHistoryOrders`: InlineResponse20072
    fmt.Fprintf(os.Stdout, "Response from `FiatApi.FiatGetHistoryOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFiatGetHistoryOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionType** | **string** | 0-deposit, 1-withdraw | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **beginTime** | **int64** |  | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **page** | **int32** | Default 1 | 
 **rows** | **int32** | Default 100, max 500 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20072**](InlineResponse20072.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FiatGetPaymentsHistory

> InlineResponse20073 FiatGetPaymentsHistory(ctx).TransactionType(transactionType).Timestamp(timestamp).Signature(signature).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()

Fiat Payments History (USER_DATA)



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
    transactionType := "0" // string | 0-deposit, 1-withdraw
    timestamp := int64(789) // int64 | UTC timestamp in ms
    signature := "signature_example" // string | Signature
    beginTime := int64(1626144956000) // int64 |  (optional)
    endTime := int64(789) // int64 | UTC timestamp in ms (optional)
    page := int32(1) // int32 | Default 1 (optional)
    rows := int32(300) // int32 | Default 100, max 500 (optional)
    recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FiatApi.FiatGetPaymentsHistory(context.Background()).TransactionType(transactionType).Timestamp(timestamp).Signature(signature).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiatApi.FiatGetPaymentsHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FiatGetPaymentsHistory`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `FiatApi.FiatGetPaymentsHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFiatGetPaymentsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionType** | **string** | 0-deposit, 1-withdraw | 
 **timestamp** | **int64** | UTC timestamp in ms | 
 **signature** | **string** | Signature | 
 **beginTime** | **int64** |  | 
 **endTime** | **int64** | UTC timestamp in ms | 
 **page** | **int32** | Default 1 | 
 **rows** | **int32** | Default 100, max 500 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

