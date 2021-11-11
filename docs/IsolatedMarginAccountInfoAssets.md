# IsolatedMarginAccountInfoAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseAsset** | [**IsolatedMarginAccountInfoBaseAsset**](IsolatedMarginAccountInfoBaseAsset.md) |  | 
**QuoteAsset** | [**IsolatedMarginAccountInfoQuoteAsset**](IsolatedMarginAccountInfoQuoteAsset.md) |  | 
**Symbol** | **string** |  | 
**IsolatedCreated** | **bool** |  | 
**Enabled** | **bool** | true-enabled, false-disabled | 
**MarginLevel** | **string** |  | 
**MarginLevelStatus** | **string** | \&quot;EXCESSIVE\&quot;, \&quot;NORMAL\&quot;, \&quot;MARGIN_CALL\&quot;, \&quot;PRE_LIQUIDATION\&quot;, \&quot;FORCE_LIQUIDATION\&quot; | 
**MarginRatio** | **string** |  | 
**IndexPrice** | **string** |  | 
**LiquidatePrice** | **string** |  | 
**LiquidateRate** | **string** |  | 
**TradeEnabled** | **bool** |  | 

## Methods

### NewIsolatedMarginAccountInfoAssets

`func NewIsolatedMarginAccountInfoAssets(baseAsset IsolatedMarginAccountInfoBaseAsset, quoteAsset IsolatedMarginAccountInfoQuoteAsset, symbol string, isolatedCreated bool, enabled bool, marginLevel string, marginLevelStatus string, marginRatio string, indexPrice string, liquidatePrice string, liquidateRate string, tradeEnabled bool, ) *IsolatedMarginAccountInfoAssets`

NewIsolatedMarginAccountInfoAssets instantiates a new IsolatedMarginAccountInfoAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsolatedMarginAccountInfoAssetsWithDefaults

`func NewIsolatedMarginAccountInfoAssetsWithDefaults() *IsolatedMarginAccountInfoAssets`

NewIsolatedMarginAccountInfoAssetsWithDefaults instantiates a new IsolatedMarginAccountInfoAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAsset

`func (o *IsolatedMarginAccountInfoAssets) GetBaseAsset() IsolatedMarginAccountInfoBaseAsset`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *IsolatedMarginAccountInfoAssets) GetBaseAssetOk() (*IsolatedMarginAccountInfoBaseAsset, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *IsolatedMarginAccountInfoAssets) SetBaseAsset(v IsolatedMarginAccountInfoBaseAsset)`

SetBaseAsset sets BaseAsset field to given value.


### GetQuoteAsset

`func (o *IsolatedMarginAccountInfoAssets) GetQuoteAsset() IsolatedMarginAccountInfoQuoteAsset`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *IsolatedMarginAccountInfoAssets) GetQuoteAssetOk() (*IsolatedMarginAccountInfoQuoteAsset, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *IsolatedMarginAccountInfoAssets) SetQuoteAsset(v IsolatedMarginAccountInfoQuoteAsset)`

SetQuoteAsset sets QuoteAsset field to given value.


### GetSymbol

`func (o *IsolatedMarginAccountInfoAssets) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *IsolatedMarginAccountInfoAssets) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *IsolatedMarginAccountInfoAssets) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetIsolatedCreated

`func (o *IsolatedMarginAccountInfoAssets) GetIsolatedCreated() bool`

GetIsolatedCreated returns the IsolatedCreated field if non-nil, zero value otherwise.

### GetIsolatedCreatedOk

`func (o *IsolatedMarginAccountInfoAssets) GetIsolatedCreatedOk() (*bool, bool)`

GetIsolatedCreatedOk returns a tuple with the IsolatedCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedCreated

`func (o *IsolatedMarginAccountInfoAssets) SetIsolatedCreated(v bool)`

SetIsolatedCreated sets IsolatedCreated field to given value.


### GetEnabled

`func (o *IsolatedMarginAccountInfoAssets) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IsolatedMarginAccountInfoAssets) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IsolatedMarginAccountInfoAssets) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMarginLevel

`func (o *IsolatedMarginAccountInfoAssets) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *IsolatedMarginAccountInfoAssets) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *IsolatedMarginAccountInfoAssets) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.


### GetMarginLevelStatus

`func (o *IsolatedMarginAccountInfoAssets) GetMarginLevelStatus() string`

GetMarginLevelStatus returns the MarginLevelStatus field if non-nil, zero value otherwise.

### GetMarginLevelStatusOk

`func (o *IsolatedMarginAccountInfoAssets) GetMarginLevelStatusOk() (*string, bool)`

GetMarginLevelStatusOk returns a tuple with the MarginLevelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevelStatus

`func (o *IsolatedMarginAccountInfoAssets) SetMarginLevelStatus(v string)`

SetMarginLevelStatus sets MarginLevelStatus field to given value.


### GetMarginRatio

`func (o *IsolatedMarginAccountInfoAssets) GetMarginRatio() string`

GetMarginRatio returns the MarginRatio field if non-nil, zero value otherwise.

### GetMarginRatioOk

`func (o *IsolatedMarginAccountInfoAssets) GetMarginRatioOk() (*string, bool)`

GetMarginRatioOk returns a tuple with the MarginRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRatio

`func (o *IsolatedMarginAccountInfoAssets) SetMarginRatio(v string)`

SetMarginRatio sets MarginRatio field to given value.


### GetIndexPrice

`func (o *IsolatedMarginAccountInfoAssets) GetIndexPrice() string`

GetIndexPrice returns the IndexPrice field if non-nil, zero value otherwise.

### GetIndexPriceOk

`func (o *IsolatedMarginAccountInfoAssets) GetIndexPriceOk() (*string, bool)`

GetIndexPriceOk returns a tuple with the IndexPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPrice

`func (o *IsolatedMarginAccountInfoAssets) SetIndexPrice(v string)`

SetIndexPrice sets IndexPrice field to given value.


### GetLiquidatePrice

`func (o *IsolatedMarginAccountInfoAssets) GetLiquidatePrice() string`

GetLiquidatePrice returns the LiquidatePrice field if non-nil, zero value otherwise.

### GetLiquidatePriceOk

`func (o *IsolatedMarginAccountInfoAssets) GetLiquidatePriceOk() (*string, bool)`

GetLiquidatePriceOk returns a tuple with the LiquidatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidatePrice

`func (o *IsolatedMarginAccountInfoAssets) SetLiquidatePrice(v string)`

SetLiquidatePrice sets LiquidatePrice field to given value.


### GetLiquidateRate

`func (o *IsolatedMarginAccountInfoAssets) GetLiquidateRate() string`

GetLiquidateRate returns the LiquidateRate field if non-nil, zero value otherwise.

### GetLiquidateRateOk

`func (o *IsolatedMarginAccountInfoAssets) GetLiquidateRateOk() (*string, bool)`

GetLiquidateRateOk returns a tuple with the LiquidateRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidateRate

`func (o *IsolatedMarginAccountInfoAssets) SetLiquidateRate(v string)`

SetLiquidateRate sets LiquidateRate field to given value.


### GetTradeEnabled

`func (o *IsolatedMarginAccountInfoAssets) GetTradeEnabled() bool`

GetTradeEnabled returns the TradeEnabled field if non-nil, zero value otherwise.

### GetTradeEnabledOk

`func (o *IsolatedMarginAccountInfoAssets) GetTradeEnabledOk() (*bool, bool)`

GetTradeEnabledOk returns a tuple with the TradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeEnabled

`func (o *IsolatedMarginAccountInfoAssets) SetTradeEnabled(v bool)`

SetTradeEnabled sets TradeEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


