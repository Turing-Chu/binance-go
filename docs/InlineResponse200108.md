# InlineResponse200108

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwapId** | **int64** |  | 
**SwapTime** | **int64** |  | 
**Status** | **int32** | 0: pending, 1: success, 2: failed | 
**QuoteAsset** | **string** |  | 
**BaseAsset** | **string** |  | 
**QuoteQty** | **float64** |  | 
**BaseQty** | **float64** |  | 
**Price** | **float64** |  | 
**Fee** | **float64** |  | 

## Methods

### NewInlineResponse200108

`func NewInlineResponse200108(swapId int64, swapTime int64, status int32, quoteAsset string, baseAsset string, quoteQty float64, baseQty float64, price float64, fee float64, ) *InlineResponse200108`

NewInlineResponse200108 instantiates a new InlineResponse200108 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200108WithDefaults

`func NewInlineResponse200108WithDefaults() *InlineResponse200108`

NewInlineResponse200108WithDefaults instantiates a new InlineResponse200108 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwapId

`func (o *InlineResponse200108) GetSwapId() int64`

GetSwapId returns the SwapId field if non-nil, zero value otherwise.

### GetSwapIdOk

`func (o *InlineResponse200108) GetSwapIdOk() (*int64, bool)`

GetSwapIdOk returns a tuple with the SwapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapId

`func (o *InlineResponse200108) SetSwapId(v int64)`

SetSwapId sets SwapId field to given value.


### GetSwapTime

`func (o *InlineResponse200108) GetSwapTime() int64`

GetSwapTime returns the SwapTime field if non-nil, zero value otherwise.

### GetSwapTimeOk

`func (o *InlineResponse200108) GetSwapTimeOk() (*int64, bool)`

GetSwapTimeOk returns a tuple with the SwapTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapTime

`func (o *InlineResponse200108) SetSwapTime(v int64)`

SetSwapTime sets SwapTime field to given value.


### GetStatus

`func (o *InlineResponse200108) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200108) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200108) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetQuoteAsset

`func (o *InlineResponse200108) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *InlineResponse200108) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *InlineResponse200108) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.


### GetBaseAsset

`func (o *InlineResponse200108) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *InlineResponse200108) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *InlineResponse200108) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.


### GetQuoteQty

`func (o *InlineResponse200108) GetQuoteQty() float64`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *InlineResponse200108) GetQuoteQtyOk() (*float64, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *InlineResponse200108) SetQuoteQty(v float64)`

SetQuoteQty sets QuoteQty field to given value.


### GetBaseQty

`func (o *InlineResponse200108) GetBaseQty() float64`

GetBaseQty returns the BaseQty field if non-nil, zero value otherwise.

### GetBaseQtyOk

`func (o *InlineResponse200108) GetBaseQtyOk() (*float64, bool)`

GetBaseQtyOk returns a tuple with the BaseQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseQty

`func (o *InlineResponse200108) SetBaseQty(v float64)`

SetBaseQty sets BaseQty field to given value.


### GetPrice

`func (o *InlineResponse200108) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse200108) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse200108) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetFee

`func (o *InlineResponse200108) GetFee() float64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *InlineResponse200108) GetFeeOk() (*float64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *InlineResponse200108) SetFee(v float64)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


