# MarginTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commission** | **string** |  | 
**CommissionAsset** | **string** |  | 
**Id** | **int64** |  | 
**IsBestMatch** | **bool** |  | 
**IsBuyer** | **bool** |  | 
**IsMaker** | **bool** |  | 
**OrderId** | **int64** |  | 
**Price** | **string** |  | 
**Qty** | **string** |  | 
**Symbol** | **string** |  | 
**IsIsolated** | **bool** |  | 
**Time** | **int64** |  | 

## Methods

### NewMarginTrade

`func NewMarginTrade(commission string, commissionAsset string, id int64, isBestMatch bool, isBuyer bool, isMaker bool, orderId int64, price string, qty string, symbol string, isIsolated bool, time int64, ) *MarginTrade`

NewMarginTrade instantiates a new MarginTrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginTradeWithDefaults

`func NewMarginTradeWithDefaults() *MarginTrade`

NewMarginTradeWithDefaults instantiates a new MarginTrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommission

`func (o *MarginTrade) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *MarginTrade) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *MarginTrade) SetCommission(v string)`

SetCommission sets Commission field to given value.


### GetCommissionAsset

`func (o *MarginTrade) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *MarginTrade) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *MarginTrade) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.


### GetId

`func (o *MarginTrade) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarginTrade) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarginTrade) SetId(v int64)`

SetId sets Id field to given value.


### GetIsBestMatch

`func (o *MarginTrade) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *MarginTrade) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *MarginTrade) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.


### GetIsBuyer

`func (o *MarginTrade) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *MarginTrade) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *MarginTrade) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.


### GetIsMaker

`func (o *MarginTrade) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *MarginTrade) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *MarginTrade) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.


### GetOrderId

`func (o *MarginTrade) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginTrade) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginTrade) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetPrice

`func (o *MarginTrade) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginTrade) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginTrade) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetQty

`func (o *MarginTrade) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *MarginTrade) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *MarginTrade) SetQty(v string)`

SetQty sets Qty field to given value.


### GetSymbol

`func (o *MarginTrade) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginTrade) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginTrade) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetIsIsolated

`func (o *MarginTrade) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginTrade) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginTrade) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.


### GetTime

`func (o *MarginTrade) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MarginTrade) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MarginTrade) SetTime(v int64)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


