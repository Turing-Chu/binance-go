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

// InlineResponse20070 struct for InlineResponse20070
type InlineResponse20070 struct {
	ListenKey string `json:"listenKey"`
}

// NewInlineResponse20070 instantiates a new InlineResponse20070 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20070(listenKey string) *InlineResponse20070 {
	this := InlineResponse20070{}
	this.ListenKey = listenKey
	return &this
}

// NewInlineResponse20070WithDefaults instantiates a new InlineResponse20070 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20070WithDefaults() *InlineResponse20070 {
	this := InlineResponse20070{}
	return &this
}

// GetListenKey returns the ListenKey field value
func (o *InlineResponse20070) GetListenKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetListenKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListenKey, true
}

// SetListenKey sets field value
func (o *InlineResponse20070) SetListenKey(v string) {
	o.ListenKey = v
}

func (o InlineResponse20070) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["listenKey"] = o.ListenKey
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20070 struct {
	value *InlineResponse20070
	isSet bool
}

func (v NullableInlineResponse20070) Get() *InlineResponse20070 {
	return v.value
}

func (v *NullableInlineResponse20070) Set(val *InlineResponse20070) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20070) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20070) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20070(val *InlineResponse20070) *NullableInlineResponse20070 {
	return &NullableInlineResponse20070{value: val, isSet: true}
}

func (v NullableInlineResponse20070) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20070) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


