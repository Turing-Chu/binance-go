# InlineResponse200111Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNumber** | **string** |  | 
**AdvNo** | **string** |  | 
**TradeType** | **string** |  | 
**Asset** | **string** |  | 
**Fiat** | **string** |  | 
**FiatSymbol** | **string** |  | 
**Amount** | **string** | Quantity (in Crypto) | 
**TotalPrice** | **string** |  | 
**UnitPrice** | **string** | Unit Price (in Fiat) | 
**OrderStatus** | **string** | PENDING, TRADING, BUYER_PAYED, DISTRIBUTING, COMPLETED, IN_APPEAL, CANCELLED, CANCELLED_BY_SYSTEM | 
**CreateTime** | **int64** |  | 
**Commission** | **string** | Transaction Fee (in Crypto) | 
**CounterPartNickName** | **string** |  | 
**AdvertisementRole** | **string** |  | 

## Methods

### NewInlineResponse200111Data

`func NewInlineResponse200111Data(orderNumber string, advNo string, tradeType string, asset string, fiat string, fiatSymbol string, amount string, totalPrice string, unitPrice string, orderStatus string, createTime int64, commission string, counterPartNickName string, advertisementRole string, ) *InlineResponse200111Data`

NewInlineResponse200111Data instantiates a new InlineResponse200111Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200111DataWithDefaults

`func NewInlineResponse200111DataWithDefaults() *InlineResponse200111Data`

NewInlineResponse200111DataWithDefaults instantiates a new InlineResponse200111Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNumber

`func (o *InlineResponse200111Data) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *InlineResponse200111Data) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *InlineResponse200111Data) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.


### GetAdvNo

`func (o *InlineResponse200111Data) GetAdvNo() string`

GetAdvNo returns the AdvNo field if non-nil, zero value otherwise.

### GetAdvNoOk

`func (o *InlineResponse200111Data) GetAdvNoOk() (*string, bool)`

GetAdvNoOk returns a tuple with the AdvNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvNo

`func (o *InlineResponse200111Data) SetAdvNo(v string)`

SetAdvNo sets AdvNo field to given value.


### GetTradeType

`func (o *InlineResponse200111Data) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *InlineResponse200111Data) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *InlineResponse200111Data) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.


### GetAsset

`func (o *InlineResponse200111Data) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *InlineResponse200111Data) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *InlineResponse200111Data) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetFiat

`func (o *InlineResponse200111Data) GetFiat() string`

GetFiat returns the Fiat field if non-nil, zero value otherwise.

### GetFiatOk

`func (o *InlineResponse200111Data) GetFiatOk() (*string, bool)`

GetFiatOk returns a tuple with the Fiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiat

`func (o *InlineResponse200111Data) SetFiat(v string)`

SetFiat sets Fiat field to given value.


### GetFiatSymbol

`func (o *InlineResponse200111Data) GetFiatSymbol() string`

GetFiatSymbol returns the FiatSymbol field if non-nil, zero value otherwise.

### GetFiatSymbolOk

`func (o *InlineResponse200111Data) GetFiatSymbolOk() (*string, bool)`

GetFiatSymbolOk returns a tuple with the FiatSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatSymbol

`func (o *InlineResponse200111Data) SetFiatSymbol(v string)`

SetFiatSymbol sets FiatSymbol field to given value.


### GetAmount

`func (o *InlineResponse200111Data) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InlineResponse200111Data) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InlineResponse200111Data) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTotalPrice

`func (o *InlineResponse200111Data) GetTotalPrice() string`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *InlineResponse200111Data) GetTotalPriceOk() (*string, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *InlineResponse200111Data) SetTotalPrice(v string)`

SetTotalPrice sets TotalPrice field to given value.


### GetUnitPrice

`func (o *InlineResponse200111Data) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InlineResponse200111Data) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InlineResponse200111Data) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.


### GetOrderStatus

`func (o *InlineResponse200111Data) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *InlineResponse200111Data) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *InlineResponse200111Data) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.


### GetCreateTime

`func (o *InlineResponse200111Data) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *InlineResponse200111Data) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *InlineResponse200111Data) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.


### GetCommission

`func (o *InlineResponse200111Data) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *InlineResponse200111Data) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *InlineResponse200111Data) SetCommission(v string)`

SetCommission sets Commission field to given value.


### GetCounterPartNickName

`func (o *InlineResponse200111Data) GetCounterPartNickName() string`

GetCounterPartNickName returns the CounterPartNickName field if non-nil, zero value otherwise.

### GetCounterPartNickNameOk

`func (o *InlineResponse200111Data) GetCounterPartNickNameOk() (*string, bool)`

GetCounterPartNickNameOk returns a tuple with the CounterPartNickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterPartNickName

`func (o *InlineResponse200111Data) SetCounterPartNickName(v string)`

SetCounterPartNickName sets CounterPartNickName field to given value.


### GetAdvertisementRole

`func (o *InlineResponse200111Data) GetAdvertisementRole() string`

GetAdvertisementRole returns the AdvertisementRole field if non-nil, zero value otherwise.

### GetAdvertisementRoleOk

`func (o *InlineResponse200111Data) GetAdvertisementRoleOk() (*string, bool)`

GetAdvertisementRoleOk returns a tuple with the AdvertisementRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisementRole

`func (o *InlineResponse200111Data) SetAdvertisementRole(v string)`

SetAdvertisementRole sets AdvertisementRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


