# InlineResponse2001Symbols

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Status** | **string** |  | 
**BaseAsset** | **string** |  | 
**BaseAssetPrecision** | **int32** |  | 
**QuoteAsset** | **string** |  | 
**QuoteAssetPrecision** | **int32** |  | 
**BaseCommissionPrecision** | **int32** |  | 
**QuoteCommissionPrecision** | **int32** |  | 
**OrderTypes** | **[]string** |  | 
**IcebergAllowed** | **bool** |  | 
**OcoAllowed** | **bool** |  | 
**QuoteOrderQtyMarketAllowed** | **bool** |  | 
**IsSpotTradingAllowed** | **bool** |  | 
**IsMarginTradingAllowed** | **bool** |  | 
**Filters** | [**[]InlineResponse2001Filters**](InlineResponse2001Filters.md) |  | 
**Permissions** | **[]string** |  | 

## Methods

### NewInlineResponse2001Symbols

`func NewInlineResponse2001Symbols(symbol string, status string, baseAsset string, baseAssetPrecision int32, quoteAsset string, quoteAssetPrecision int32, baseCommissionPrecision int32, quoteCommissionPrecision int32, orderTypes []string, icebergAllowed bool, ocoAllowed bool, quoteOrderQtyMarketAllowed bool, isSpotTradingAllowed bool, isMarginTradingAllowed bool, filters []InlineResponse2001Filters, permissions []string, ) *InlineResponse2001Symbols`

NewInlineResponse2001Symbols instantiates a new InlineResponse2001Symbols object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001SymbolsWithDefaults

`func NewInlineResponse2001SymbolsWithDefaults() *InlineResponse2001Symbols`

NewInlineResponse2001SymbolsWithDefaults instantiates a new InlineResponse2001Symbols object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InlineResponse2001Symbols) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineResponse2001Symbols) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineResponse2001Symbols) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetStatus

`func (o *InlineResponse2001Symbols) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2001Symbols) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2001Symbols) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBaseAsset

`func (o *InlineResponse2001Symbols) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *InlineResponse2001Symbols) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *InlineResponse2001Symbols) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.


### GetBaseAssetPrecision

`func (o *InlineResponse2001Symbols) GetBaseAssetPrecision() int32`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *InlineResponse2001Symbols) GetBaseAssetPrecisionOk() (*int32, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *InlineResponse2001Symbols) SetBaseAssetPrecision(v int32)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.


### GetQuoteAsset

`func (o *InlineResponse2001Symbols) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *InlineResponse2001Symbols) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *InlineResponse2001Symbols) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.


### GetQuoteAssetPrecision

`func (o *InlineResponse2001Symbols) GetQuoteAssetPrecision() int32`

GetQuoteAssetPrecision returns the QuoteAssetPrecision field if non-nil, zero value otherwise.

### GetQuoteAssetPrecisionOk

`func (o *InlineResponse2001Symbols) GetQuoteAssetPrecisionOk() (*int32, bool)`

GetQuoteAssetPrecisionOk returns a tuple with the QuoteAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAssetPrecision

`func (o *InlineResponse2001Symbols) SetQuoteAssetPrecision(v int32)`

SetQuoteAssetPrecision sets QuoteAssetPrecision field to given value.


### GetBaseCommissionPrecision

`func (o *InlineResponse2001Symbols) GetBaseCommissionPrecision() int32`

GetBaseCommissionPrecision returns the BaseCommissionPrecision field if non-nil, zero value otherwise.

### GetBaseCommissionPrecisionOk

`func (o *InlineResponse2001Symbols) GetBaseCommissionPrecisionOk() (*int32, bool)`

GetBaseCommissionPrecisionOk returns a tuple with the BaseCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCommissionPrecision

`func (o *InlineResponse2001Symbols) SetBaseCommissionPrecision(v int32)`

SetBaseCommissionPrecision sets BaseCommissionPrecision field to given value.


### GetQuoteCommissionPrecision

`func (o *InlineResponse2001Symbols) GetQuoteCommissionPrecision() int32`

GetQuoteCommissionPrecision returns the QuoteCommissionPrecision field if non-nil, zero value otherwise.

### GetQuoteCommissionPrecisionOk

`func (o *InlineResponse2001Symbols) GetQuoteCommissionPrecisionOk() (*int32, bool)`

GetQuoteCommissionPrecisionOk returns a tuple with the QuoteCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCommissionPrecision

`func (o *InlineResponse2001Symbols) SetQuoteCommissionPrecision(v int32)`

SetQuoteCommissionPrecision sets QuoteCommissionPrecision field to given value.


### GetOrderTypes

`func (o *InlineResponse2001Symbols) GetOrderTypes() []string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *InlineResponse2001Symbols) GetOrderTypesOk() (*[]string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *InlineResponse2001Symbols) SetOrderTypes(v []string)`

SetOrderTypes sets OrderTypes field to given value.


### GetIcebergAllowed

`func (o *InlineResponse2001Symbols) GetIcebergAllowed() bool`

GetIcebergAllowed returns the IcebergAllowed field if non-nil, zero value otherwise.

### GetIcebergAllowedOk

`func (o *InlineResponse2001Symbols) GetIcebergAllowedOk() (*bool, bool)`

GetIcebergAllowedOk returns a tuple with the IcebergAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergAllowed

`func (o *InlineResponse2001Symbols) SetIcebergAllowed(v bool)`

SetIcebergAllowed sets IcebergAllowed field to given value.


### GetOcoAllowed

`func (o *InlineResponse2001Symbols) GetOcoAllowed() bool`

GetOcoAllowed returns the OcoAllowed field if non-nil, zero value otherwise.

### GetOcoAllowedOk

`func (o *InlineResponse2001Symbols) GetOcoAllowedOk() (*bool, bool)`

GetOcoAllowedOk returns a tuple with the OcoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcoAllowed

`func (o *InlineResponse2001Symbols) SetOcoAllowed(v bool)`

SetOcoAllowed sets OcoAllowed field to given value.


### GetQuoteOrderQtyMarketAllowed

`func (o *InlineResponse2001Symbols) GetQuoteOrderQtyMarketAllowed() bool`

GetQuoteOrderQtyMarketAllowed returns the QuoteOrderQtyMarketAllowed field if non-nil, zero value otherwise.

### GetQuoteOrderQtyMarketAllowedOk

`func (o *InlineResponse2001Symbols) GetQuoteOrderQtyMarketAllowedOk() (*bool, bool)`

GetQuoteOrderQtyMarketAllowedOk returns a tuple with the QuoteOrderQtyMarketAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteOrderQtyMarketAllowed

`func (o *InlineResponse2001Symbols) SetQuoteOrderQtyMarketAllowed(v bool)`

SetQuoteOrderQtyMarketAllowed sets QuoteOrderQtyMarketAllowed field to given value.


### GetIsSpotTradingAllowed

`func (o *InlineResponse2001Symbols) GetIsSpotTradingAllowed() bool`

GetIsSpotTradingAllowed returns the IsSpotTradingAllowed field if non-nil, zero value otherwise.

### GetIsSpotTradingAllowedOk

`func (o *InlineResponse2001Symbols) GetIsSpotTradingAllowedOk() (*bool, bool)`

GetIsSpotTradingAllowedOk returns a tuple with the IsSpotTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpotTradingAllowed

`func (o *InlineResponse2001Symbols) SetIsSpotTradingAllowed(v bool)`

SetIsSpotTradingAllowed sets IsSpotTradingAllowed field to given value.


### GetIsMarginTradingAllowed

`func (o *InlineResponse2001Symbols) GetIsMarginTradingAllowed() bool`

GetIsMarginTradingAllowed returns the IsMarginTradingAllowed field if non-nil, zero value otherwise.

### GetIsMarginTradingAllowedOk

`func (o *InlineResponse2001Symbols) GetIsMarginTradingAllowedOk() (*bool, bool)`

GetIsMarginTradingAllowedOk returns a tuple with the IsMarginTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarginTradingAllowed

`func (o *InlineResponse2001Symbols) SetIsMarginTradingAllowed(v bool)`

SetIsMarginTradingAllowed sets IsMarginTradingAllowed field to given value.


### GetFilters

`func (o *InlineResponse2001Symbols) GetFilters() []InlineResponse2001Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InlineResponse2001Symbols) GetFiltersOk() (*[]InlineResponse2001Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InlineResponse2001Symbols) SetFilters(v []InlineResponse2001Filters)`

SetFilters sets Filters field to given value.


### GetPermissions

`func (o *InlineResponse2001Symbols) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse2001Symbols) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse2001Symbols) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


