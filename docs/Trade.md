# Trade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | trade id | 
**Price** | **string** | price | 
**Qty** | **string** | amount of base asset | 
**QuoteQty** | **string** | amount of quote asset | 
**Time** | **int64** | Trade executed timestamp, as same as &#x60;T&#x60; in the stream | 
**IsBuyerMaker** | **bool** |  | 
**IsBestMatch** | **bool** |  | 

## Methods

### NewTrade

`func NewTrade(id int64, price string, qty string, quoteQty string, time int64, isBuyerMaker bool, isBestMatch bool, ) *Trade`

NewTrade instantiates a new Trade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeWithDefaults

`func NewTradeWithDefaults() *Trade`

NewTradeWithDefaults instantiates a new Trade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Trade) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trade) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trade) SetId(v int64)`

SetId sets Id field to given value.


### GetPrice

`func (o *Trade) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Trade) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Trade) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetQty

`func (o *Trade) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *Trade) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *Trade) SetQty(v string)`

SetQty sets Qty field to given value.


### GetQuoteQty

`func (o *Trade) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *Trade) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *Trade) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.


### GetTime

`func (o *Trade) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Trade) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Trade) SetTime(v int64)`

SetTime sets Time field to given value.


### GetIsBuyerMaker

`func (o *Trade) GetIsBuyerMaker() bool`

GetIsBuyerMaker returns the IsBuyerMaker field if non-nil, zero value otherwise.

### GetIsBuyerMakerOk

`func (o *Trade) GetIsBuyerMakerOk() (*bool, bool)`

GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyerMaker

`func (o *Trade) SetIsBuyerMaker(v bool)`

SetIsBuyerMaker sets IsBuyerMaker field to given value.


### GetIsBestMatch

`func (o *Trade) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *Trade) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *Trade) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


