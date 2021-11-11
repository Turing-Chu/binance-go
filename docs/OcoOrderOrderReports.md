# OcoOrderOrderReports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**OrigClientOrderId** | **string** |  | 
**OrderId** | **int64** |  | 
**OrderListId** | **int64** |  | 
**ClientOrderId** | **string** |  | 
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

### NewOcoOrderOrderReports

`func NewOcoOrderOrderReports(symbol string, origClientOrderId string, orderId int64, orderListId int64, clientOrderId string, price string, origQty string, executedQty string, cummulativeQuoteQty string, status string, timeInForce string, type_ string, side string, stopPrice string, ) *OcoOrderOrderReports`

NewOcoOrderOrderReports instantiates a new OcoOrderOrderReports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOcoOrderOrderReportsWithDefaults

`func NewOcoOrderOrderReportsWithDefaults() *OcoOrderOrderReports`

NewOcoOrderOrderReportsWithDefaults instantiates a new OcoOrderOrderReports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OcoOrderOrderReports) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OcoOrderOrderReports) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OcoOrderOrderReports) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrigClientOrderId

`func (o *OcoOrderOrderReports) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *OcoOrderOrderReports) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *OcoOrderOrderReports) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.


### GetOrderId

`func (o *OcoOrderOrderReports) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OcoOrderOrderReports) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OcoOrderOrderReports) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetOrderListId

`func (o *OcoOrderOrderReports) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OcoOrderOrderReports) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OcoOrderOrderReports) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetClientOrderId

`func (o *OcoOrderOrderReports) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OcoOrderOrderReports) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OcoOrderOrderReports) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetPrice

`func (o *OcoOrderOrderReports) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OcoOrderOrderReports) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OcoOrderOrderReports) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetOrigQty

`func (o *OcoOrderOrderReports) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OcoOrderOrderReports) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OcoOrderOrderReports) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.


### GetExecutedQty

`func (o *OcoOrderOrderReports) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OcoOrderOrderReports) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OcoOrderOrderReports) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.


### GetCummulativeQuoteQty

`func (o *OcoOrderOrderReports) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OcoOrderOrderReports) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OcoOrderOrderReports) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.


### GetStatus

`func (o *OcoOrderOrderReports) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OcoOrderOrderReports) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OcoOrderOrderReports) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeInForce

`func (o *OcoOrderOrderReports) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OcoOrderOrderReports) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OcoOrderOrderReports) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.


### GetType

`func (o *OcoOrderOrderReports) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OcoOrderOrderReports) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OcoOrderOrderReports) SetType(v string)`

SetType sets Type field to given value.


### GetSide

`func (o *OcoOrderOrderReports) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OcoOrderOrderReports) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OcoOrderOrderReports) SetSide(v string)`

SetSide sets Side field to given value.


### GetStopPrice

`func (o *OcoOrderOrderReports) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OcoOrderOrderReports) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OcoOrderOrderReports) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


