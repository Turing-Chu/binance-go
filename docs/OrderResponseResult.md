# OrderResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**OrderId** | **int64** |  | 
**OrderListId** | **int64** |  | 
**ClientOrderId** | **string** |  | 
**TransactTime** | **int64** |  | 
**Price** | **string** |  | 
**OrigQty** | **string** |  | 
**ExecutedQty** | **string** |  | 
**CummulativeQuoteQty** | **string** |  | 
**Status** | **string** |  | 
**TimeInForce** | **string** |  | 
**Type** | **string** |  | 
**Side** | **string** |  | 

## Methods

### NewOrderResponseResult

`func NewOrderResponseResult(symbol string, orderId int64, orderListId int64, clientOrderId string, transactTime int64, price string, origQty string, executedQty string, cummulativeQuoteQty string, status string, timeInForce string, type_ string, side string, ) *OrderResponseResult`

NewOrderResponseResult instantiates a new OrderResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseResultWithDefaults

`func NewOrderResponseResultWithDefaults() *OrderResponseResult`

NewOrderResponseResultWithDefaults instantiates a new OrderResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderResponseResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderResponseResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderResponseResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderId

`func (o *OrderResponseResult) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderResponseResult) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderResponseResult) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetOrderListId

`func (o *OrderResponseResult) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderResponseResult) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderResponseResult) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetClientOrderId

`func (o *OrderResponseResult) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderResponseResult) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderResponseResult) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetTransactTime

`func (o *OrderResponseResult) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OrderResponseResult) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OrderResponseResult) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.


### GetPrice

`func (o *OrderResponseResult) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderResponseResult) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderResponseResult) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetOrigQty

`func (o *OrderResponseResult) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OrderResponseResult) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OrderResponseResult) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.


### GetExecutedQty

`func (o *OrderResponseResult) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderResponseResult) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderResponseResult) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.


### GetCummulativeQuoteQty

`func (o *OrderResponseResult) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OrderResponseResult) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OrderResponseResult) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.


### GetStatus

`func (o *OrderResponseResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderResponseResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderResponseResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeInForce

`func (o *OrderResponseResult) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderResponseResult) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderResponseResult) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.


### GetType

`func (o *OrderResponseResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderResponseResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderResponseResult) SetType(v string)`

SetType sets Type field to given value.


### GetSide

`func (o *OrderResponseResult) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderResponseResult) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderResponseResult) SetSide(v string)`

SetSide sets Side field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


