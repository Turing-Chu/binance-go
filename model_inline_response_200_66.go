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

// InlineResponse20066 struct for InlineResponse20066
type InlineResponse20066 struct {
	TranId int64 `json:"tranId"`
}

// NewInlineResponse20066 instantiates a new InlineResponse20066 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20066(tranId int64) *InlineResponse20066 {
	this := InlineResponse20066{}
	this.TranId = tranId
	return &this
}

// NewInlineResponse20066WithDefaults instantiates a new InlineResponse20066 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20066WithDefaults() *InlineResponse20066 {
	this := InlineResponse20066{}
	return &this
}

// GetTranId returns the TranId field value
func (o *InlineResponse20066) GetTranId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20066) GetTranIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TranId, true
}

// SetTranId sets field value
func (o *InlineResponse20066) SetTranId(v int64) {
	o.TranId = v
}

func (o InlineResponse20066) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tranId"] = o.TranId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20066 struct {
	value *InlineResponse20066
	isSet bool
}

func (v NullableInlineResponse20066) Get() *InlineResponse20066 {
	return v.value
}

func (v *NullableInlineResponse20066) Set(val *InlineResponse20066) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20066) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20066) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20066(val *InlineResponse20066) *NullableInlineResponse20066 {
	return &NullableInlineResponse20066{value: val, isSet: true}
}

func (v NullableInlineResponse20066) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20066) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


