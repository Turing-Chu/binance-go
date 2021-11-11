# OrderResponseAck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**OrderId** | **int64** |  | 
**OrderListId** | **int64** |  | 
**ClientOrderId** | **string** |  | 
**TransactTime** | **int64** |  | 

## Methods

### NewOrderResponseAck

`func NewOrderResponseAck(symbol string, orderId int64, orderListId int64, clientOrderId string, transactTime int64, ) *OrderResponseAck`

NewOrderResponseAck instantiates a new OrderResponseAck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseAckWithDefaults

`func NewOrderResponseAckWithDefaults() *OrderResponseAck`

NewOrderResponseAckWithDefaults instantiates a new OrderResponseAck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderResponseAck) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderResponseAck) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderResponseAck) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderId

`func (o *OrderResponseAck) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderResponseAck) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderResponseAck) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetOrderListId

`func (o *OrderResponseAck) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderResponseAck) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderResponseAck) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetClientOrderId

`func (o *OrderResponseAck) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderResponseAck) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderResponseAck) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetTransactTime

`func (o *OrderResponseAck) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OrderResponseAck) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OrderResponseAck) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


