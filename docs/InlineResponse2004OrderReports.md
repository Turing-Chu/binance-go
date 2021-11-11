# InlineResponse2004OrderReports

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
**StopPrice** | **string** |  | 

## Methods

### NewInlineResponse2004OrderReports

`func NewInlineResponse2004OrderReports(symbol string, orderId int64, orderListId int64, clientOrderId string, transactTime int64, price string, origQty string, executedQty string, cummulativeQuoteQty string, status string, timeInForce string, type_ string, side string, stopPrice string, ) *InlineResponse2004OrderReports`

NewInlineResponse2004OrderReports instantiates a new InlineResponse2004OrderReports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2004OrderReportsWithDefaults

`func NewInlineResponse2004OrderReportsWithDefaults() *InlineResponse2004OrderReports`

NewInlineResponse2004OrderReportsWithDefaults instantiates a new InlineResponse2004OrderReports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InlineResponse2004OrderReports) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineResponse2004OrderReports) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineResponse2004OrderReports) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderId

`func (o *InlineResponse2004OrderReports) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *InlineResponse2004OrderReports) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *InlineResponse2004OrderReports) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetOrderListId

`func (o *InlineResponse2004OrderReports) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *InlineResponse2004OrderReports) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *InlineResponse2004OrderReports) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetClientOrderId

`func (o *InlineResponse2004OrderReports) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *InlineResponse2004OrderReports) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *InlineResponse2004OrderReports) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetTransactTime

`func (o *InlineResponse2004OrderReports) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *InlineResponse2004OrderReports) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *InlineResponse2004OrderReports) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.


### GetPrice

`func (o *InlineResponse2004OrderReports) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse2004OrderReports) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse2004OrderReports) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetOrigQty

`func (o *InlineResponse2004OrderReports) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *InlineResponse2004OrderReports) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *InlineResponse2004OrderReports) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.


### GetExecutedQty

`func (o *InlineResponse2004OrderReports) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *InlineResponse2004OrderReports) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *InlineResponse2004OrderReports) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.


### GetCummulativeQuoteQty

`func (o *InlineResponse2004OrderReports) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *InlineResponse2004OrderReports) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *InlineResponse2004OrderReports) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.


### GetStatus

`func (o *InlineResponse2004OrderReports) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2004OrderReports) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2004OrderReports) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeInForce

`func (o *InlineResponse2004OrderReports) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *InlineResponse2004OrderReports) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *InlineResponse2004OrderReports) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.


### GetType

`func (o *InlineResponse2004OrderReports) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse2004OrderReports) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse2004OrderReports) SetType(v string)`

SetType sets Type field to given value.


### GetSide

`func (o *InlineResponse2004OrderReports) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *InlineResponse2004OrderReports) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *InlineResponse2004OrderReports) SetSide(v string)`

SetSide sets Side field to given value.


### GetStopPrice

`func (o *InlineResponse2004OrderReports) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *InlineResponse2004OrderReports) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *InlineResponse2004OrderReports) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


