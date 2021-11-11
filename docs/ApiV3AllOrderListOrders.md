# ApiV3AllOrderListOrders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**OrderId** | **int64** |  | 
**ClientOrderId** | **string** |  | 

## Methods

### NewApiV3AllOrderListOrders

`func NewApiV3AllOrderListOrders(symbol string, orderId int64, clientOrderId string, ) *ApiV3AllOrderListOrders`

NewApiV3AllOrderListOrders instantiates a new ApiV3AllOrderListOrders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV3AllOrderListOrdersWithDefaults

`func NewApiV3AllOrderListOrdersWithDefaults() *ApiV3AllOrderListOrders`

NewApiV3AllOrderListOrdersWithDefaults instantiates a new ApiV3AllOrderListOrders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ApiV3AllOrderListOrders) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ApiV3AllOrderListOrders) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ApiV3AllOrderListOrders) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderId

`func (o *ApiV3AllOrderListOrders) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ApiV3AllOrderListOrders) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ApiV3AllOrderListOrders) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetClientOrderId

`func (o *ApiV3AllOrderListOrders) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *ApiV3AllOrderListOrders) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *ApiV3AllOrderListOrders) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


