/*
Binance Public Spot API

OpenAPI Specifications for the Binance Public Spot API  API documents:   - [https://github.com/binance/binance-spot-api-docs](https://github.com/binance/binance-spot-api-docs)   - [https://binance-docs.github.io/apidocs/spot/en](https://binance-docs.github.io/apidocs/spot/en)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BswapRmvLiquidityPreviewSingle struct for BswapRmvLiquidityPreviewSingle
type BswapRmvLiquidityPreviewSingle struct {
	QuoteAsset string `json:"quoteAsset"`
	QuoteAmt int64 `json:"quoteAmt"`
	Price float64 `json:"price"`
	Slippage float64 `json:"slippage"`
	Fee float64 `json:"fee"`
}

// NewBswapRmvLiquidityPreviewSingle instantiates a new BswapRmvLiquidityPreviewSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBswapRmvLiquidityPreviewSingle(quoteAsset string, quoteAmt int64, price float64, slippage float64, fee float64) *BswapRmvLiquidityPreviewSingle {
	this := BswapRmvLiquidityPreviewSingle{}
	this.QuoteAsset = quoteAsset
	this.QuoteAmt = quoteAmt
	this.Price = price
	this.Slippage = slippage
	this.Fee = fee
	return &this
}

// NewBswapRmvLiquidityPreviewSingleWithDefaults instantiates a new BswapRmvLiquidityPreviewSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBswapRmvLiquidityPreviewSingleWithDefaults() *BswapRmvLiquidityPreviewSingle {
	this := BswapRmvLiquidityPreviewSingle{}
	return &this
}

// GetQuoteAsset returns the QuoteAsset field value
func (o *BswapRmvLiquidityPreviewSingle) GetQuoteAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value
// and a boolean to check if the value has been set.
func (o *BswapRmvLiquidityPreviewSingle) GetQuoteAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QuoteAsset, true
}

// SetQuoteAsset sets field value
func (o *BswapRmvLiquidityPreviewSingle) SetQuoteAsset(v string) {
	o.QuoteAsset = v
}

// GetQuoteAmt returns the QuoteAmt field value
func (o *BswapRmvLiquidityPreviewSingle) GetQuoteAmt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.QuoteAmt
}

// GetQuoteAmtOk returns a tuple with the QuoteAmt field value
// and a boolean to check if the value has been set.
func (o *BswapRmvLiquidityPreviewSingle) GetQuoteAmtOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QuoteAmt, true
}

// SetQuoteAmt sets field value
func (o *BswapRmvLiquidityPreviewSingle) SetQuoteAmt(v int64) {
	o.QuoteAmt = v
}

// GetPrice returns the Price field value
func (o *BswapRmvLiquidityPreviewSingle) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *BswapRmvLiquidityPreviewSingle) GetPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *BswapRmvLiquidityPreviewSingle) SetPrice(v float64) {
	o.Price = v
}

// GetSlippage returns the Slippage field value
func (o *BswapRmvLiquidityPreviewSingle) GetSlippage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Slippage
}

// GetSlippageOk returns a tuple with the Slippage field value
// and a boolean to check if the value has been set.
func (o *BswapRmvLiquidityPreviewSingle) GetSlippageOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slippage, true
}

// SetSlippage sets field value
func (o *BswapRmvLiquidityPreviewSingle) SetSlippage(v float64) {
	o.Slippage = v
}

// GetFee returns the Fee field value
func (o *BswapRmvLiquidityPreviewSingle) GetFee() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *BswapRmvLiquidityPreviewSingle) GetFeeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *BswapRmvLiquidityPreviewSingle) SetFee(v float64) {
	o.Fee = v
}

func (o BswapRmvLiquidityPreviewSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if true {
		toSerialize["quoteAmt"] = o.QuoteAmt
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["slippage"] = o.Slippage
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	return json.Marshal(toSerialize)
}

type NullableBswapRmvLiquidityPreviewSingle struct {
	value *BswapRmvLiquidityPreviewSingle
	isSet bool
}

func (v NullableBswapRmvLiquidityPreviewSingle) Get() *BswapRmvLiquidityPreviewSingle {
	return v.value
}

func (v *NullableBswapRmvLiquidityPreviewSingle) Set(val *BswapRmvLiquidityPreviewSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableBswapRmvLiquidityPreviewSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableBswapRmvLiquidityPreviewSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBswapRmvLiquidityPreviewSingle(val *BswapRmvLiquidityPreviewSingle) *NullableBswapRmvLiquidityPreviewSingle {
	return &NullableBswapRmvLiquidityPreviewSingle{value: val, isSet: true}
}

func (v NullableBswapRmvLiquidityPreviewSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBswapRmvLiquidityPreviewSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


