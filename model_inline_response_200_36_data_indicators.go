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

// InlineResponse20036DataIndicators The indicators updated every 30 seconds
type InlineResponse20036DataIndicators struct {
	BTCUSDT []InlineResponse20036DataIndicatorsBTCUSDT `json:"BTCUSDT"`
}

// NewInlineResponse20036DataIndicators instantiates a new InlineResponse20036DataIndicators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036DataIndicators(bTCUSDT []InlineResponse20036DataIndicatorsBTCUSDT) *InlineResponse20036DataIndicators {
	this := InlineResponse20036DataIndicators{}
	this.BTCUSDT = bTCUSDT
	return &this
}

// NewInlineResponse20036DataIndicatorsWithDefaults instantiates a new InlineResponse20036DataIndicators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036DataIndicatorsWithDefaults() *InlineResponse20036DataIndicators {
	this := InlineResponse20036DataIndicators{}
	return &this
}

// GetBTCUSDT returns the BTCUSDT field value
func (o *InlineResponse20036DataIndicators) GetBTCUSDT() []InlineResponse20036DataIndicatorsBTCUSDT {
	if o == nil {
		var ret []InlineResponse20036DataIndicatorsBTCUSDT
		return ret
	}

	return o.BTCUSDT
}

// GetBTCUSDTOk returns a tuple with the BTCUSDT field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20036DataIndicators) GetBTCUSDTOk() (*[]InlineResponse20036DataIndicatorsBTCUSDT, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BTCUSDT, true
}

// SetBTCUSDT sets field value
func (o *InlineResponse20036DataIndicators) SetBTCUSDT(v []InlineResponse20036DataIndicatorsBTCUSDT) {
	o.BTCUSDT = v
}

func (o InlineResponse20036DataIndicators) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["BTCUSDT"] = o.BTCUSDT
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20036DataIndicators struct {
	value *InlineResponse20036DataIndicators
	isSet bool
}

func (v NullableInlineResponse20036DataIndicators) Get() *InlineResponse20036DataIndicators {
	return v.value
}

func (v *NullableInlineResponse20036DataIndicators) Set(val *InlineResponse20036DataIndicators) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20036DataIndicators) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20036DataIndicators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20036DataIndicators(val *InlineResponse20036DataIndicators) *NullableInlineResponse20036DataIndicators {
	return &NullableInlineResponse20036DataIndicators{value: val, isSet: true}
}

func (v NullableInlineResponse20036DataIndicators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20036DataIndicators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

