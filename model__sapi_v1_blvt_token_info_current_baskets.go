/*
Binance Public Spot API

OpenAPI Specifications for the Binance Public Spot API  API documents:   - [https://github.com/binance/binance-spot-api-docs](https://github.com/binance/binance-spot-api-docs)   - [https://binance-docs.github.io/apidocs/spot/en](https://binance-docs.github.io/apidocs/spot/en)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binanceapi

import (
	"encoding/json"
)

// SapiV1BlvtTokenInfoCurrentBaskets struct for SapiV1BlvtTokenInfoCurrentBaskets
type SapiV1BlvtTokenInfoCurrentBaskets struct {
	Symbol string `json:"symbol"`
	Amount string `json:"amount"`
	NotionalValue string `json:"notionalValue"`
}

// NewSapiV1BlvtTokenInfoCurrentBaskets instantiates a new SapiV1BlvtTokenInfoCurrentBaskets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSapiV1BlvtTokenInfoCurrentBaskets(symbol string, amount string, notionalValue string) *SapiV1BlvtTokenInfoCurrentBaskets {
	this := SapiV1BlvtTokenInfoCurrentBaskets{}
	this.Symbol = symbol
	this.Amount = amount
	this.NotionalValue = notionalValue
	return &this
}

// NewSapiV1BlvtTokenInfoCurrentBasketsWithDefaults instantiates a new SapiV1BlvtTokenInfoCurrentBaskets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSapiV1BlvtTokenInfoCurrentBasketsWithDefaults() *SapiV1BlvtTokenInfoCurrentBaskets {
	this := SapiV1BlvtTokenInfoCurrentBaskets{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *SapiV1BlvtTokenInfoCurrentBaskets) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SapiV1BlvtTokenInfoCurrentBaskets) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SapiV1BlvtTokenInfoCurrentBaskets) SetSymbol(v string) {
	o.Symbol = v
}

// GetAmount returns the Amount field value
func (o *SapiV1BlvtTokenInfoCurrentBaskets) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *SapiV1BlvtTokenInfoCurrentBaskets) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *SapiV1BlvtTokenInfoCurrentBaskets) SetAmount(v string) {
	o.Amount = v
}

// GetNotionalValue returns the NotionalValue field value
func (o *SapiV1BlvtTokenInfoCurrentBaskets) GetNotionalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotionalValue
}

// GetNotionalValueOk returns a tuple with the NotionalValue field value
// and a boolean to check if the value has been set.
func (o *SapiV1BlvtTokenInfoCurrentBaskets) GetNotionalValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotionalValue, true
}

// SetNotionalValue sets field value
func (o *SapiV1BlvtTokenInfoCurrentBaskets) SetNotionalValue(v string) {
	o.NotionalValue = v
}

func (o SapiV1BlvtTokenInfoCurrentBaskets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["notionalValue"] = o.NotionalValue
	}
	return json.Marshal(toSerialize)
}

type NullableSapiV1BlvtTokenInfoCurrentBaskets struct {
	value *SapiV1BlvtTokenInfoCurrentBaskets
	isSet bool
}

func (v NullableSapiV1BlvtTokenInfoCurrentBaskets) Get() *SapiV1BlvtTokenInfoCurrentBaskets {
	return v.value
}

func (v *NullableSapiV1BlvtTokenInfoCurrentBaskets) Set(val *SapiV1BlvtTokenInfoCurrentBaskets) {
	v.value = val
	v.isSet = true
}

func (v NullableSapiV1BlvtTokenInfoCurrentBaskets) IsSet() bool {
	return v.isSet
}

func (v *NullableSapiV1BlvtTokenInfoCurrentBaskets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSapiV1BlvtTokenInfoCurrentBaskets(val *SapiV1BlvtTokenInfoCurrentBaskets) *NullableSapiV1BlvtTokenInfoCurrentBaskets {
	return &NullableSapiV1BlvtTokenInfoCurrentBaskets{value: val, isSet: true}
}

func (v NullableSapiV1BlvtTokenInfoCurrentBaskets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSapiV1BlvtTokenInfoCurrentBaskets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


