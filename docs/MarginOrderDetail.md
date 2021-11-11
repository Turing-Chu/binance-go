# MarginOrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | **string** |  | 
**CummulativeQuoteQty** | **string** |  | 
**ExecutedQty** | **string** |  | 
**IcebergQty** | **string** |  | 
**IsWorking** | **bool** |  | 
**OrderId** | **int64** |  | 
**OrigQty** | **string** |  | 
**Price** | **string** |  | 
**Side** | **string** |  | 
**Status** | **string** |  | 
**StopPrice** | **string** |  | 
**Symbol** | **string** |  | 
**IsIsolated** | **bool** |  | 
**Time** | **int64** |  | 
**TimeInForce** | **string** |  | 
**Type** | **string** |  | 
**UpdateTime** | **int64** |  | 

## Methods

### NewMarginOrderDetail

`func NewMarginOrderDetail(clientOrderId string, cummulativeQuoteQty string, executedQty string, icebergQty string, isWorking bool, orderId int64, origQty string, price string, side string, status string, stopPrice string, symbol string, isIsolated bool, time int64, timeInForce string, type_ string, updateTime int64, ) *MarginOrderDetail`

NewMarginOrderDetail instantiates a new MarginOrderDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginOrderDetailWithDefaults

`func NewMarginOrderDetailWithDefaults() *MarginOrderDetail`

NewMarginOrderDetailWithDefaults instantiates a new MarginOrderDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *MarginOrderDetail) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *MarginOrderDetail) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *MarginOrderDetail) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetCummulativeQuoteQty

`func (o *MarginOrderDetail) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *MarginOrderDetail) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *MarginOrderDetail) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.


### GetExecutedQty

`func (o *MarginOrderDetail) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *MarginOrderDetail) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *MarginOrderDetail) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.


### GetIcebergQty

`func (o *MarginOrderDetail) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *MarginOrderDetail) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *MarginOrderDetail) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.


### GetIsWorking

`func (o *MarginOrderDetail) GetIsWorking() bool`

GetIsWorking returns the IsWorking field if non-nil, zero value otherwise.

### GetIsWorkingOk

`func (o *MarginOrderDetail) GetIsWorkingOk() (*bool, bool)`

GetIsWorkingOk returns a tuple with the IsWorking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorking

`func (o *MarginOrderDetail) SetIsWorking(v bool)`

SetIsWorking sets IsWorking field to given value.


### GetOrderId

`func (o *MarginOrderDetail) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginOrderDetail) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginOrderDetail) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetOrigQty

`func (o *MarginOrderDetail) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *MarginOrderDetail) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *MarginOrderDetail) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.


### GetPrice

`func (o *MarginOrderDetail) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginOrderDetail) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginOrderDetail) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetSide

`func (o *MarginOrderDetail) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *MarginOrderDetail) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *MarginOrderDetail) SetSide(v string)`

SetSide sets Side field to given value.


### GetStatus

`func (o *MarginOrderDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarginOrderDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarginOrderDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStopPrice

`func (o *MarginOrderDetail) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *MarginOrderDetail) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *MarginOrderDetail) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.


### GetSymbol

`func (o *MarginOrderDetail) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginOrderDetail) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginOrderDetail) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetIsIsolated

`func (o *MarginOrderDetail) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginOrderDetail) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginOrderDetail) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.


### GetTime

`func (o *MarginOrderDetail) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MarginOrderDetail) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MarginOrderDetail) SetTime(v int64)`

SetTime sets Time field to given value.


### GetTimeInForce

`func (o *MarginOrderDetail) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *MarginOrderDetail) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *MarginOrderDetail) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.


### GetType

`func (o *MarginOrderDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarginOrderDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarginOrderDetail) SetType(v string)`

SetType sets Type field to given value.


### GetUpdateTime

`func (o *MarginOrderDetail) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *MarginOrderDetail) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *MarginOrderDetail) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


