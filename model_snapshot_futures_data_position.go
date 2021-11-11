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

// SnapshotFuturesDataPosition struct for SnapshotFuturesDataPosition
type SnapshotFuturesDataPosition struct {
	EntryPrice string `json:"entryPrice"`
	MarkPrice string `json:"markPrice"`
	PositionAmt string `json:"positionAmt"`
	Symbol string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
}

// NewSnapshotFuturesDataPosition instantiates a new SnapshotFuturesDataPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotFuturesDataPosition(entryPrice string, markPrice string, positionAmt string, symbol string, unRealizedProfit string) *SnapshotFuturesDataPosition {
	this := SnapshotFuturesDataPosition{}
	this.EntryPrice = entryPrice
	this.MarkPrice = markPrice
	this.PositionAmt = positionAmt
	this.Symbol = symbol
	this.UnRealizedProfit = unRealizedProfit
	return &this
}

// NewSnapshotFuturesDataPositionWithDefaults instantiates a new SnapshotFuturesDataPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotFuturesDataPositionWithDefaults() *SnapshotFuturesDataPosition {
	this := SnapshotFuturesDataPosition{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value
func (o *SnapshotFuturesDataPosition) GetEntryPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value
// and a boolean to check if the value has been set.
func (o *SnapshotFuturesDataPosition) GetEntryPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntryPrice, true
}

// SetEntryPrice sets field value
func (o *SnapshotFuturesDataPosition) SetEntryPrice(v string) {
	o.EntryPrice = v
}

// GetMarkPrice returns the MarkPrice field value
func (o *SnapshotFuturesDataPosition) GetMarkPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value
// and a boolean to check if the value has been set.
func (o *SnapshotFuturesDataPosition) GetMarkPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarkPrice, true
}

// SetMarkPrice sets field value
func (o *SnapshotFuturesDataPosition) SetMarkPrice(v string) {
	o.MarkPrice = v
}

// GetPositionAmt returns the PositionAmt field value
func (o *SnapshotFuturesDataPosition) GetPositionAmt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value
// and a boolean to check if the value has been set.
func (o *SnapshotFuturesDataPosition) GetPositionAmtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PositionAmt, true
}

// SetPositionAmt sets field value
func (o *SnapshotFuturesDataPosition) SetPositionAmt(v string) {
	o.PositionAmt = v
}

// GetSymbol returns the Symbol field value
func (o *SnapshotFuturesDataPosition) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SnapshotFuturesDataPosition) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SnapshotFuturesDataPosition) SetSymbol(v string) {
	o.Symbol = v
}

// GetUnRealizedProfit returns the UnRealizedProfit field value
func (o *SnapshotFuturesDataPosition) GetUnRealizedProfit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnRealizedProfit
}

// GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field value
// and a boolean to check if the value has been set.
func (o *SnapshotFuturesDataPosition) GetUnRealizedProfitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnRealizedProfit, true
}

// SetUnRealizedProfit sets field value
func (o *SnapshotFuturesDataPosition) SetUnRealizedProfit(v string) {
	o.UnRealizedProfit = v
}

func (o SnapshotFuturesDataPosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if true {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if true {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["unRealizedProfit"] = o.UnRealizedProfit
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotFuturesDataPosition struct {
	value *SnapshotFuturesDataPosition
	isSet bool
}

func (v NullableSnapshotFuturesDataPosition) Get() *SnapshotFuturesDataPosition {
	return v.value
}

func (v *NullableSnapshotFuturesDataPosition) Set(val *SnapshotFuturesDataPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotFuturesDataPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotFuturesDataPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotFuturesDataPosition(val *SnapshotFuturesDataPosition) *NullableSnapshotFuturesDataPosition {
	return &NullableSnapshotFuturesDataPosition{value: val, isSet: true}
}

func (v NullableSnapshotFuturesDataPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotFuturesDataPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


