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

// InlineResponse20029 struct for InlineResponse20029
type InlineResponse20029 struct {
	// 0: normal, 1：system maintenance
	Status int32 `json:"status"`
	// \"normal\", \"system_maintenance\"
	Msg string `json:"msg"`
}

// NewInlineResponse20029 instantiates a new InlineResponse20029 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20029(status int32, msg string) *InlineResponse20029 {
	this := InlineResponse20029{}
	this.Status = status
	this.Msg = msg
	return &this
}

// NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20029WithDefaults() *InlineResponse20029 {
	this := InlineResponse20029{}
	return &this
}

// GetStatus returns the Status field value
func (o *InlineResponse20029) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetStatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20029) SetStatus(v int32) {
	o.Status = v
}

// GetMsg returns the Msg field value
func (o *InlineResponse20029) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetMsgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *InlineResponse20029) SetMsg(v string) {
	o.Msg = v
}

func (o InlineResponse20029) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["msg"] = o.Msg
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20029 struct {
	value *InlineResponse20029
	isSet bool
}

func (v NullableInlineResponse20029) Get() *InlineResponse20029 {
	return v.value
}

func (v *NullableInlineResponse20029) Set(val *InlineResponse20029) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20029) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20029) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20029(val *InlineResponse20029) *NullableInlineResponse20029 {
	return &NullableInlineResponse20029{value: val, isSet: true}
}

func (v NullableInlineResponse20029) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20029) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


