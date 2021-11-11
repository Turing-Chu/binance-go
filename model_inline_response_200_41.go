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

// InlineResponse20041 struct for InlineResponse20041
type InlineResponse20041 struct {
	Symbol string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

// NewInlineResponse20041 instantiates a new InlineResponse20041 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20041(symbol string, makerCommission string, takerCommission string) *InlineResponse20041 {
	this := InlineResponse20041{}
	this.Symbol = symbol
	this.MakerCommission = makerCommission
	this.TakerCommission = takerCommission
	return &this
}

// NewInlineResponse20041WithDefaults instantiates a new InlineResponse20041 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20041WithDefaults() *InlineResponse20041 {
	this := InlineResponse20041{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *InlineResponse20041) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20041) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *InlineResponse20041) SetSymbol(v string) {
	o.Symbol = v
}

// GetMakerCommission returns the MakerCommission field value
func (o *InlineResponse20041) GetMakerCommission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MakerCommission
}

// GetMakerCommissionOk returns a tuple with the MakerCommission field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20041) GetMakerCommissionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MakerCommission, true
}

// SetMakerCommission sets field value
func (o *InlineResponse20041) SetMakerCommission(v string) {
	o.MakerCommission = v
}

// GetTakerCommission returns the TakerCommission field value
func (o *InlineResponse20041) GetTakerCommission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20041) GetTakerCommissionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TakerCommission, true
}

// SetTakerCommission sets field value
func (o *InlineResponse20041) SetTakerCommission(v string) {
	o.TakerCommission = v
}

func (o InlineResponse20041) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["makerCommission"] = o.MakerCommission
	}
	if true {
		toSerialize["takerCommission"] = o.TakerCommission
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20041 struct {
	value *InlineResponse20041
	isSet bool
}

func (v NullableInlineResponse20041) Get() *InlineResponse20041 {
	return v.value
}

func (v *NullableInlineResponse20041) Set(val *InlineResponse20041) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20041) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20041) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20041(val *InlineResponse20041) *NullableInlineResponse20041 {
	return &NullableInlineResponse20041{value: val, isSet: true}
}

func (v NullableInlineResponse20041) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20041) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


