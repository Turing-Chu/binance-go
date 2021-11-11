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

// InlineResponse20075 struct for InlineResponse20075
type InlineResponse20075 struct {
	Asset string `json:"asset"`
	LeftQuota string `json:"leftQuota"`
}

// NewInlineResponse20075 instantiates a new InlineResponse20075 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20075(asset string, leftQuota string) *InlineResponse20075 {
	this := InlineResponse20075{}
	this.Asset = asset
	this.LeftQuota = leftQuota
	return &this
}

// NewInlineResponse20075WithDefaults instantiates a new InlineResponse20075 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20075WithDefaults() *InlineResponse20075 {
	this := InlineResponse20075{}
	return &this
}

// GetAsset returns the Asset field value
func (o *InlineResponse20075) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20075) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20075) SetAsset(v string) {
	o.Asset = v
}

// GetLeftQuota returns the LeftQuota field value
func (o *InlineResponse20075) GetLeftQuota() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeftQuota
}

// GetLeftQuotaOk returns a tuple with the LeftQuota field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20075) GetLeftQuotaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LeftQuota, true
}

// SetLeftQuota sets field value
func (o *InlineResponse20075) SetLeftQuota(v string) {
	o.LeftQuota = v
}

func (o InlineResponse20075) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["leftQuota"] = o.LeftQuota
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20075 struct {
	value *InlineResponse20075
	isSet bool
}

func (v NullableInlineResponse20075) Get() *InlineResponse20075 {
	return v.value
}

func (v *NullableInlineResponse20075) Set(val *InlineResponse20075) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20075) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20075) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20075(val *InlineResponse20075) *NullableInlineResponse20075 {
	return &NullableInlineResponse20075{value: val, isSet: true}
}

func (v NullableInlineResponse20075) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20075) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


