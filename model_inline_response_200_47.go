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

// InlineResponse20047 struct for InlineResponse20047
type InlineResponse20047 struct {
	SubAccounts []InlineResponse20047SubAccounts `json:"subAccounts"`
}

// NewInlineResponse20047 instantiates a new InlineResponse20047 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20047(subAccounts []InlineResponse20047SubAccounts) *InlineResponse20047 {
	this := InlineResponse20047{}
	this.SubAccounts = subAccounts
	return &this
}

// NewInlineResponse20047WithDefaults instantiates a new InlineResponse20047 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20047WithDefaults() *InlineResponse20047 {
	this := InlineResponse20047{}
	return &this
}

// GetSubAccounts returns the SubAccounts field value
func (o *InlineResponse20047) GetSubAccounts() []InlineResponse20047SubAccounts {
	if o == nil {
		var ret []InlineResponse20047SubAccounts
		return ret
	}

	return o.SubAccounts
}

// GetSubAccountsOk returns a tuple with the SubAccounts field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20047) GetSubAccountsOk() (*[]InlineResponse20047SubAccounts, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubAccounts, true
}

// SetSubAccounts sets field value
func (o *InlineResponse20047) SetSubAccounts(v []InlineResponse20047SubAccounts) {
	o.SubAccounts = v
}

func (o InlineResponse20047) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subAccounts"] = o.SubAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20047 struct {
	value *InlineResponse20047
	isSet bool
}

func (v NullableInlineResponse20047) Get() *InlineResponse20047 {
	return v.value
}

func (v *NullableInlineResponse20047) Set(val *InlineResponse20047) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20047) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20047) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20047(val *InlineResponse20047) *NullableInlineResponse20047 {
	return &NullableInlineResponse20047{value: val, isSet: true}
}

func (v NullableInlineResponse20047) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20047) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


