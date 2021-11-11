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

// SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos struct for SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos
type SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos struct {
	EntryPrice string `json:"entryPrice"`
	MarkPrice string `json:"markPrice"`
	Leverage string `json:"leverage"`
	Isolated string `json:"isolated"`
	IsolatedWallet string `json:"isolatedWallet"`
	IsolatedMargin string `json:"isolatedMargin"`
	IsAutoAddMargin string `json:"isAutoAddMargin"`
	PositionSide string `json:"positionSide"`
	PositionAmount string `json:"positionAmount"`
	Symbol string `json:"symbol"`
	UnrealizedProfit string `json:"unrealizedProfit"`
}

// NewSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos instantiates a new SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos(entryPrice string, markPrice string, leverage string, isolated string, isolatedWallet string, isolatedMargin string, isAutoAddMargin string, positionSide string, positionAmount string, symbol string, unrealizedProfit string) *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos {
	this := SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos{}
	this.EntryPrice = entryPrice
	this.MarkPrice = markPrice
	this.Leverage = leverage
	this.Isolated = isolated
	this.IsolatedWallet = isolatedWallet
	this.IsolatedMargin = isolatedMargin
	this.IsAutoAddMargin = isAutoAddMargin
	this.PositionSide = positionSide
	this.PositionAmount = positionAmount
	this.Symbol = symbol
	this.UnrealizedProfit = unrealizedProfit
	return &this
}

// NewSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVosWithDefaults instantiates a new SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVosWithDefaults() *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos {
	this := SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetEntryPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetEntryPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntryPrice, true
}

// SetEntryPrice sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetEntryPrice(v string) {
	o.EntryPrice = v
}

// GetMarkPrice returns the MarkPrice field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetMarkPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetMarkPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarkPrice, true
}

// SetMarkPrice sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetMarkPrice(v string) {
	o.MarkPrice = v
}

// GetLeverage returns the Leverage field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetLeverage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetLeverageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Leverage, true
}

// SetLeverage sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetLeverage(v string) {
	o.Leverage = v
}

// GetIsolated returns the Isolated field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetIsolated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Isolated
}

// GetIsolatedOk returns a tuple with the Isolated field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetIsolatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Isolated, true
}

// SetIsolated sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetIsolated(v string) {
	o.Isolated = v
}

// GetIsolatedWallet returns the IsolatedWallet field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetIsolatedWallet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsolatedWallet
}

// GetIsolatedWalletOk returns a tuple with the IsolatedWallet field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetIsolatedWalletOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsolatedWallet, true
}

// SetIsolatedWallet sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetIsolatedWallet(v string) {
	o.IsolatedWallet = v
}

// GetIsolatedMargin returns the IsolatedMargin field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetIsolatedMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsolatedMargin
}

// GetIsolatedMarginOk returns a tuple with the IsolatedMargin field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetIsolatedMarginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsolatedMargin, true
}

// SetIsolatedMargin sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetIsolatedMargin(v string) {
	o.IsolatedMargin = v
}

// GetIsAutoAddMargin returns the IsAutoAddMargin field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetIsAutoAddMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsAutoAddMargin
}

// GetIsAutoAddMarginOk returns a tuple with the IsAutoAddMargin field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetIsAutoAddMarginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsAutoAddMargin, true
}

// SetIsAutoAddMargin sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetIsAutoAddMargin(v string) {
	o.IsAutoAddMargin = v
}

// GetPositionSide returns the PositionSide field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetPositionSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetPositionSideOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PositionSide, true
}

// SetPositionSide sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetPositionSide(v string) {
	o.PositionSide = v
}

// GetPositionAmount returns the PositionAmount field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetPositionAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PositionAmount
}

// GetPositionAmountOk returns a tuple with the PositionAmount field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetPositionAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PositionAmount, true
}

// SetPositionAmount sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetPositionAmount(v string) {
	o.PositionAmount = v
}

// GetSymbol returns the Symbol field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetSymbol(v string) {
	o.Symbol = v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetUnrealizedProfit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnrealizedProfit, true
}

// SetUnrealizedProfit sets field value
func (o *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = v
}

func (o SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if true {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if true {
		toSerialize["leverage"] = o.Leverage
	}
	if true {
		toSerialize["isolated"] = o.Isolated
	}
	if true {
		toSerialize["isolatedWallet"] = o.IsolatedWallet
	}
	if true {
		toSerialize["isolatedMargin"] = o.IsolatedMargin
	}
	if true {
		toSerialize["isAutoAddMargin"] = o.IsAutoAddMargin
	}
	if true {
		toSerialize["positionSide"] = o.PositionSide
	}
	if true {
		toSerialize["positionAmount"] = o.PositionAmount
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	return json.Marshal(toSerialize)
}

type NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos struct {
	value *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos
	isSet bool
}

func (v NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) Get() *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos {
	return v.value
}

func (v *NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) Set(val *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos(val *SubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) *NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos {
	return &NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos{value: val, isSet: true}
}

func (v NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountCOINFuturesPositionRiskDeliveryPositionRiskVos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

