# ApiV3OpenOrderListOrders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**OrderId** | **int64** |  | 
**ClientOrderId** | **string** |  | 

## Methods

### NewApiV3OpenOrderListOrders

`func NewApiV3OpenOrderListOrders(symbol string, orderId int64, clientOrderId string, ) *ApiV3OpenOrderListOrders`

NewApiV3OpenOrderListOrders instantiates a new ApiV3OpenOrderListOrders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV3OpenOrderListOrdersWithDefaults

`func NewApiV3OpenOrderListOrdersWithDefaults() *ApiV3OpenOrderListOrders`

NewApiV3OpenOrderListOrdersWithDefaults instantiates a new ApiV3OpenOrderListOrders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ApiV3OpenOrderListOrders) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ApiV3OpenOrderListOrders) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ApiV3OpenOrderListOrders) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderId

`func (o *ApiV3OpenOrderListOrders) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ApiV3OpenOrderListOrders) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ApiV3OpenOrderListOrders) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetClientOrderId

`func (o *ApiV3OpenOrderListOrders) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *ApiV3OpenOrderListOrders) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *ApiV3OpenOrderListOrders) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


