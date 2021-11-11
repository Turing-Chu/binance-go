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

// InlineResponse20096 struct for InlineResponse20096
type InlineResponse20096 struct {
	Code int64 `json:"code"`
	Msg string `json:"msg"`
	Data []InlineResponse20096Data `json:"data"`
}

// NewInlineResponse20096 instantiates a new InlineResponse20096 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20096(code int64, msg string, data []InlineResponse20096Data) *InlineResponse20096 {
	this := InlineResponse20096{}
	this.Code = code
	this.Msg = msg
	this.Data = data
	return &this
}

// NewInlineResponse20096WithDefaults instantiates a new InlineResponse20096 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20096WithDefaults() *InlineResponse20096 {
	this := InlineResponse20096{}
	return &this
}

// GetCode returns the Code field value
func (o *InlineResponse20096) GetCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20096) GetCodeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InlineResponse20096) SetCode(v int64) {
	o.Code = v
}

// GetMsg returns the Msg field value
func (o *InlineResponse20096) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20096) GetMsgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *InlineResponse20096) SetMsg(v string) {
	o.Msg = v
}

// GetData returns the Data field value
func (o *InlineResponse20096) GetData() []InlineResponse20096Data {
	if o == nil {
		var ret []InlineResponse20096Data
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20096) GetDataOk() (*[]InlineResponse20096Data, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InlineResponse20096) SetData(v []InlineResponse20096Data) {
	o.Data = v
}

func (o InlineResponse20096) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["msg"] = o.Msg
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20096 struct {
	value *InlineResponse20096
	isSet bool
}

func (v NullableInlineResponse20096) Get() *InlineResponse20096 {
	return v.value
}

func (v *NullableInlineResponse20096) Set(val *InlineResponse20096) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20096) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20096) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20096(val *InlineResponse20096) *NullableInlineResponse20096 {
	return &NullableInlineResponse20096{value: val, isSet: true}
}

func (v NullableInlineResponse20096) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20096) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


