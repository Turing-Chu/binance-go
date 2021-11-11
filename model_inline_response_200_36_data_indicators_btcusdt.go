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

// InlineResponse20036DataIndicatorsBTCUSDT struct for InlineResponse20036DataIndicatorsBTCUSDT
type InlineResponse20036DataIndicatorsBTCUSDT struct {
	// Unfilled Ratio (UFR)
	I string `json:"i"`
	// Count of all orders
	C int64 `json:"c"`
	// Current UFR value
	V float32 `json:"v"`
	// Trigger UFR value
	T float32 `json:"t"`
}

// NewInlineResponse20036DataIndicatorsBTCUSDT instantiates a new InlineResponse20036DataIndicatorsBTCUSDT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036DataIndicatorsBTCUSDT(i string, c int64, v float32, t float32) *InlineResponse20036DataIndicatorsBTCUSDT {
	this := InlineResponse20036DataIndicatorsBTCUSDT{}
	this.I = i
	this.C = c
	this.V = v
	this.T = t
	return &this
}

// NewInlineResponse20036DataIndicatorsBTCUSDTWithDefaults instantiates a new InlineResponse20036DataIndicatorsBTCUSDT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036DataIndicatorsBTCUSDTWithDefaults() *InlineResponse20036DataIndicatorsBTCUSDT {
	this := InlineResponse20036DataIndicatorsBTCUSDT{}
	return &this
}

// GetI returns the I field value
func (o *InlineResponse20036DataIndicatorsBTCUSDT) GetI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.I
}

// GetIOk returns a tuple with the I field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20036DataIndicatorsBTCUSDT) GetIOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.I, true
}

// SetI sets field value
func (o *InlineResponse20036DataIndicatorsBTCUSDT) SetI(v string) {
	o.I = v
}

// GetC returns the C field value
func (o *InlineResponse20036DataIndicatorsBTCUSDT) GetC() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.C
}

// GetCOk returns a tuple with the C field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20036DataIndicatorsBTCUSDT) GetCOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.C, true
}

// SetC sets field value
func (o *InlineResponse20036DataIndicatorsBTCUSDT) SetC(v int64) {
	o.C = v
}

// GetV returns the V field value
func (o *InlineResponse20036DataIndicatorsBTCUSDT) GetV() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.V
}

// GetVOk returns a tuple with the V field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20036DataIndicatorsBTCUSDT) GetVOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.V, true
}

// SetV sets field value
func (o *InlineResponse20036DataIndicatorsBTCUSDT) SetV(v float32) {
	o.V = v
}

// GetT returns the T field value
func (o *InlineResponse20036DataIndicatorsBTCUSDT) GetT() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20036DataIndicatorsBTCUSDT) GetTOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *InlineResponse20036DataIndicatorsBTCUSDT) SetT(v float32) {
	o.T = v
}

func (o InlineResponse20036DataIndicatorsBTCUSDT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["i"] = o.I
	}
	if true {
		toSerialize["c"] = o.C
	}
	if true {
		toSerialize["v"] = o.V
	}
	if true {
		toSerialize["t"] = o.T
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20036DataIndicatorsBTCUSDT struct {
	value *InlineResponse20036DataIndicatorsBTCUSDT
	isSet bool
}

func (v NullableInlineResponse20036DataIndicatorsBTCUSDT) Get() *InlineResponse20036DataIndicatorsBTCUSDT {
	return v.value
}

func (v *NullableInlineResponse20036DataIndicatorsBTCUSDT) Set(val *InlineResponse20036DataIndicatorsBTCUSDT) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20036DataIndicatorsBTCUSDT) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20036DataIndicatorsBTCUSDT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20036DataIndicatorsBTCUSDT(val *InlineResponse20036DataIndicatorsBTCUSDT) *NullableInlineResponse20036DataIndicatorsBTCUSDT {
	return &NullableInlineResponse20036DataIndicatorsBTCUSDT{value: val, isSet: true}
}

func (v NullableInlineResponse20036DataIndicatorsBTCUSDT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20036DataIndicatorsBTCUSDT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


