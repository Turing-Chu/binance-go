# MyTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Id** | **int64** | Trade id | 
**OrderId** | **int64** |  | 
**OrderListId** | **int64** |  | 
**Price** | **string** | Price | 
**Qty** | **string** | Amount of base asset | 
**QuoteQty** | **string** | Amount of quote asset | 
**Commission** | **string** |  | 
**CommissionAsset** | **string** |  | 
**Time** | **int64** | Trade timestamp | 
**IsBuyer** | **bool** |  | 
**IsMaker** | **bool** |  | 
**IsBestMatch** | **bool** |  | 

## Methods

### NewMyTrade

`func NewMyTrade(symbol string, id int64, orderId int64, orderListId int64, price string, qty string, quoteQty string, commission string, commissionAsset string, time int64, isBuyer bool, isMaker bool, isBestMatch bool, ) *MyTrade`

NewMyTrade instantiates a new MyTrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyTradeWithDefaults

`func NewMyTradeWithDefaults() *MyTrade`

NewMyTradeWithDefaults instantiates a new MyTrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MyTrade) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MyTrade) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MyTrade) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetId

`func (o *MyTrade) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyTrade) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyTrade) SetId(v int64)`

SetId sets Id field to given value.


### GetOrderId

`func (o *MyTrade) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MyTrade) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MyTrade) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetOrderListId

`func (o *MyTrade) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MyTrade) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MyTrade) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetPrice

`func (o *MyTrade) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MyTrade) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MyTrade) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetQty

`func (o *MyTrade) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *MyTrade) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *MyTrade) SetQty(v string)`

SetQty sets Qty field to given value.


### GetQuoteQty

`func (o *MyTrade) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *MyTrade) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *MyTrade) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.


### GetCommission

`func (o *MyTrade) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *MyTrade) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *MyTrade) SetCommission(v string)`

SetCommission sets Commission field to given value.


### GetCommissionAsset

`func (o *MyTrade) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *MyTrade) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *MyTrade) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.


### GetTime

`func (o *MyTrade) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MyTrade) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MyTrade) SetTime(v int64)`

SetTime sets Time field to given value.


### GetIsBuyer

`func (o *MyTrade) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *MyTrade) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *MyTrade) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.


### GetIsMaker

`func (o *MyTrade) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *MyTrade) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *MyTrade) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.


### GetIsBestMatch

`func (o *MyTrade) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *MyTrade) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *MyTrade) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


