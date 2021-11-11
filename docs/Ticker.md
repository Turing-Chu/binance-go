# Ticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**PriceChange** | **string** |  | 
**PriceChangePercent** | **string** |  | 
**PrevClosePrice** | **string** |  | 
**LastPrice** | **string** |  | 
**BidPrice** | **string** |  | 
**BidQty** | **string** |  | 
**AskPrice** | **string** |  | 
**AskQty** | **string** |  | 
**OpenPrice** | **string** |  | 
**HighPrice** | **string** |  | 
**LowPrice** | **string** |  | 
**Volume** | **string** |  | 
**QuoteVolume** | **string** |  | 
**OpenTime** | **int64** |  | 
**CloseTime** | **int64** |  | 
**FirstId** | **int64** |  | 
**LastId** | **int64** |  | 
**Count** | **int64** |  | 

## Methods

### NewTicker

`func NewTicker(symbol string, priceChange string, priceChangePercent string, prevClosePrice string, lastPrice string, bidPrice string, bidQty string, askPrice string, askQty string, openPrice string, highPrice string, lowPrice string, volume string, quoteVolume string, openTime int64, closeTime int64, firstId int64, lastId int64, count int64, ) *Ticker`

NewTicker instantiates a new Ticker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerWithDefaults

`func NewTickerWithDefaults() *Ticker`

NewTickerWithDefaults instantiates a new Ticker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Ticker) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Ticker) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Ticker) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetPriceChange

`func (o *Ticker) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *Ticker) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *Ticker) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.


### GetPriceChangePercent

`func (o *Ticker) GetPriceChangePercent() string`

GetPriceChangePercent returns the PriceChangePercent field if non-nil, zero value otherwise.

### GetPriceChangePercentOk

`func (o *Ticker) GetPriceChangePercentOk() (*string, bool)`

GetPriceChangePercentOk returns a tuple with the PriceChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePercent

`func (o *Ticker) SetPriceChangePercent(v string)`

SetPriceChangePercent sets PriceChangePercent field to given value.


### GetPrevClosePrice

`func (o *Ticker) GetPrevClosePrice() string`

GetPrevClosePrice returns the PrevClosePrice field if non-nil, zero value otherwise.

### GetPrevClosePriceOk

`func (o *Ticker) GetPrevClosePriceOk() (*string, bool)`

GetPrevClosePriceOk returns a tuple with the PrevClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevClosePrice

`func (o *Ticker) SetPrevClosePrice(v string)`

SetPrevClosePrice sets PrevClosePrice field to given value.


### GetLastPrice

`func (o *Ticker) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *Ticker) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *Ticker) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.


### GetBidPrice

`func (o *Ticker) GetBidPrice() string`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *Ticker) GetBidPriceOk() (*string, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *Ticker) SetBidPrice(v string)`

SetBidPrice sets BidPrice field to given value.


### GetBidQty

`func (o *Ticker) GetBidQty() string`

GetBidQty returns the BidQty field if non-nil, zero value otherwise.

### GetBidQtyOk

`func (o *Ticker) GetBidQtyOk() (*string, bool)`

GetBidQtyOk returns a tuple with the BidQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidQty

`func (o *Ticker) SetBidQty(v string)`

SetBidQty sets BidQty field to given value.


### GetAskPrice

`func (o *Ticker) GetAskPrice() string`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *Ticker) GetAskPriceOk() (*string, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *Ticker) SetAskPrice(v string)`

SetAskPrice sets AskPrice field to given value.


### GetAskQty

`func (o *Ticker) GetAskQty() string`

GetAskQty returns the AskQty field if non-nil, zero value otherwise.

### GetAskQtyOk

`func (o *Ticker) GetAskQtyOk() (*string, bool)`

GetAskQtyOk returns a tuple with the AskQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskQty

`func (o *Ticker) SetAskQty(v string)`

SetAskQty sets AskQty field to given value.


### GetOpenPrice

`func (o *Ticker) GetOpenPrice() string`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *Ticker) GetOpenPriceOk() (*string, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *Ticker) SetOpenPrice(v string)`

SetOpenPrice sets OpenPrice field to given value.


### GetHighPrice

`func (o *Ticker) GetHighPrice() string`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *Ticker) GetHighPriceOk() (*string, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *Ticker) SetHighPrice(v string)`

SetHighPrice sets HighPrice field to given value.


### GetLowPrice

`func (o *Ticker) GetLowPrice() string`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *Ticker) GetLowPriceOk() (*string, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *Ticker) SetLowPrice(v string)`

SetLowPrice sets LowPrice field to given value.


### GetVolume

`func (o *Ticker) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Ticker) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Ticker) SetVolume(v string)`

SetVolume sets Volume field to given value.


### GetQuoteVolume

`func (o *Ticker) GetQuoteVolume() string`

GetQuoteVolume returns the QuoteVolume field if non-nil, zero value otherwise.

### GetQuoteVolumeOk

`func (o *Ticker) GetQuoteVolumeOk() (*string, bool)`

GetQuoteVolumeOk returns a tuple with the QuoteVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVolume

`func (o *Ticker) SetQuoteVolume(v string)`

SetQuoteVolume sets QuoteVolume field to given value.


### GetOpenTime

`func (o *Ticker) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *Ticker) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *Ticker) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.


### GetCloseTime

`func (o *Ticker) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *Ticker) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *Ticker) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.


### GetFirstId

`func (o *Ticker) GetFirstId() int64`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *Ticker) GetFirstIdOk() (*int64, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *Ticker) SetFirstId(v int64)`

SetFirstId sets FirstId field to given value.


### GetLastId

`func (o *Ticker) GetLastId() int64`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *Ticker) GetLastIdOk() (*int64, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *Ticker) SetLastId(v int64)`

SetLastId sets LastId field to given value.


### GetCount

`func (o *Ticker) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Ticker) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Ticker) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


