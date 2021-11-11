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

// InlineResponse20035 struct for InlineResponse20035
type InlineResponse20035 struct {
	Data string `json:"data"`
}

// NewInlineResponse20035 instantiates a new InlineResponse20035 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20035(data string) *InlineResponse20035 {
	this := InlineResponse20035{}
	this.Data = data
	return &this
}

// NewInlineResponse20035WithDefaults instantiates a new InlineResponse20035 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20035WithDefaults() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// GetData returns the Data field value
func (o *InlineResponse20035) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InlineResponse20035) SetData(v string) {
	o.Data = v
}

func (o InlineResponse20035) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20035 struct {
	value *InlineResponse20035
	isSet bool
}

func (v NullableInlineResponse20035) Get() *InlineResponse20035 {
	return v.value
}

func (v *NullableInlineResponse20035) Set(val *InlineResponse20035) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20035) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20035) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20035(val *InlineResponse20035) *NullableInlineResponse20035 {
	return &NullableInlineResponse20035{value: val, isSet: true}
}

func (v NullableInlineResponse20035) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20035) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


