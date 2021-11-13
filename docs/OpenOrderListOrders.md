# OpenOrderListOrders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**OrderId** | **int64** |  | 
**ClientOrderId** | **string** |  | 

## Methods

### NewOpenOrderListOrders

`func NewOpenOrderListOrders(symbol string, orderId int64, clientOrderId string, ) *OpenOrderListOrders`

NewOpenOrderListOrders instantiates a new OpenOrderListOrders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenOrderListOrdersWithDefaults

`func NewOpenOrderListOrdersWithDefaults() *OpenOrderListOrders`

NewOpenOrderListOrdersWithDefaults instantiates a new OpenOrderListOrders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OpenOrderListOrders) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OpenOrderListOrders) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OpenOrderListOrders) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderId

`func (o *OpenOrderListOrders) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OpenOrderListOrders) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OpenOrderListOrders) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetClientOrderId

`func (o *OpenOrderListOrders) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OpenOrderListOrders) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OpenOrderListOrders) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


