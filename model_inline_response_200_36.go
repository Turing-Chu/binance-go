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

// InlineResponse20036 struct for InlineResponse20036
type InlineResponse20036 struct {
	Data InlineResponse20036Data `json:"data"`
}

// NewInlineResponse20036 instantiates a new InlineResponse20036 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036(data InlineResponse20036Data) *InlineResponse20036 {
	this := InlineResponse20036{}
	this.Data = data
	return &this
}

// NewInlineResponse20036WithDefaults instantiates a new InlineResponse20036 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036WithDefaults() *InlineResponse20036 {
	this := InlineResponse20036{}
	return &this
}

// GetData returns the Data field value
func (o *InlineResponse20036) GetData() InlineResponse20036Data {
	if o == nil {
		var ret InlineResponse20036Data
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20036) GetDataOk() (*InlineResponse20036Data, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InlineResponse20036) SetData(v InlineResponse20036Data) {
	o.Data = v
}

func (o InlineResponse20036) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20036 struct {
	value *InlineResponse20036
	isSet bool
}

func (v NullableInlineResponse20036) Get() *InlineResponse20036 {
	return v.value
}

func (v *NullableInlineResponse20036) Set(val *InlineResponse20036) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20036) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20036) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20036(val *InlineResponse20036) *NullableInlineResponse20036 {
	return &NullableInlineResponse20036{value: val, isSet: true}
}

func (v NullableInlineResponse20036) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20036) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


