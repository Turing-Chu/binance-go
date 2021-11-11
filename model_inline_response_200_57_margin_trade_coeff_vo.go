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

// InlineResponse20057MarginTradeCoeffVo struct for InlineResponse20057MarginTradeCoeffVo
type InlineResponse20057MarginTradeCoeffVo struct {
	// Liquidation margin ratio
	ForceLiquidationBar string `json:"forceLiquidationBar"`
	// Margin call margin ratio
	MarginCallBar string `json:"marginCallBar"`
	// Initial margin ratio
	NormalBar string `json:"normalBar"`
}

// NewInlineResponse20057MarginTradeCoeffVo instantiates a new InlineResponse20057MarginTradeCoeffVo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20057MarginTradeCoeffVo(forceLiquidationBar string, marginCallBar string, normalBar string) *InlineResponse20057MarginTradeCoeffVo {
	this := InlineResponse20057MarginTradeCoeffVo{}
	this.ForceLiquidationBar = forceLiquidationBar
	this.MarginCallBar = marginCallBar
	this.NormalBar = normalBar
	return &this
}

// NewInlineResponse20057MarginTradeCoeffVoWithDefaults instantiates a new InlineResponse20057MarginTradeCoeffVo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20057MarginTradeCoeffVoWithDefaults() *InlineResponse20057MarginTradeCoeffVo {
	this := InlineResponse20057MarginTradeCoeffVo{}
	return &this
}

// GetForceLiquidationBar returns the ForceLiquidationBar field value
func (o *InlineResponse20057MarginTradeCoeffVo) GetForceLiquidationBar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForceLiquidationBar
}

// GetForceLiquidationBarOk returns a tuple with the ForceLiquidationBar field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057MarginTradeCoeffVo) GetForceLiquidationBarOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ForceLiquidationBar, true
}

// SetForceLiquidationBar sets field value
func (o *InlineResponse20057MarginTradeCoeffVo) SetForceLiquidationBar(v string) {
	o.ForceLiquidationBar = v
}

// GetMarginCallBar returns the MarginCallBar field value
func (o *InlineResponse20057MarginTradeCoeffVo) GetMarginCallBar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarginCallBar
}

// GetMarginCallBarOk returns a tuple with the MarginCallBar field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057MarginTradeCoeffVo) GetMarginCallBarOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarginCallBar, true
}

// SetMarginCallBar sets field value
func (o *InlineResponse20057MarginTradeCoeffVo) SetMarginCallBar(v string) {
	o.MarginCallBar = v
}

// GetNormalBar returns the NormalBar field value
func (o *InlineResponse20057MarginTradeCoeffVo) GetNormalBar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NormalBar
}

// GetNormalBarOk returns a tuple with the NormalBar field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057MarginTradeCoeffVo) GetNormalBarOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NormalBar, true
}

// SetNormalBar sets field value
func (o *InlineResponse20057MarginTradeCoeffVo) SetNormalBar(v string) {
	o.NormalBar = v
}

func (o InlineResponse20057MarginTradeCoeffVo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["forceLiquidationBar"] = o.ForceLiquidationBar
	}
	if true {
		toSerialize["marginCallBar"] = o.MarginCallBar
	}
	if true {
		toSerialize["normalBar"] = o.NormalBar
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20057MarginTradeCoeffVo struct {
	value *InlineResponse20057MarginTradeCoeffVo
	isSet bool
}

func (v NullableInlineResponse20057MarginTradeCoeffVo) Get() *InlineResponse20057MarginTradeCoeffVo {
	return v.value
}

func (v *NullableInlineResponse20057MarginTradeCoeffVo) Set(val *InlineResponse20057MarginTradeCoeffVo) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20057MarginTradeCoeffVo) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20057MarginTradeCoeffVo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20057MarginTradeCoeffVo(val *InlineResponse20057MarginTradeCoeffVo) *NullableInlineResponse20057MarginTradeCoeffVo {
	return &NullableInlineResponse20057MarginTradeCoeffVo{value: val, isSet: true}
}

func (v NullableInlineResponse20057MarginTradeCoeffVo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20057MarginTradeCoeffVo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


