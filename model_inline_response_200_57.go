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

// InlineResponse20057 struct for InlineResponse20057
type InlineResponse20057 struct {
	Email string `json:"email"`
	MarginLevel string `json:"marginLevel"`
	TotalAssetOfBtc string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc string `json:"totalNetAssetOfBtc"`
	MarginTradeCoeffVo InlineResponse20057MarginTradeCoeffVo `json:"marginTradeCoeffVo"`
	MarginUserAssetVoList []InlineResponse20019UserAssets `json:"marginUserAssetVoList"`
}

// NewInlineResponse20057 instantiates a new InlineResponse20057 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20057(email string, marginLevel string, totalAssetOfBtc string, totalLiabilityOfBtc string, totalNetAssetOfBtc string, marginTradeCoeffVo InlineResponse20057MarginTradeCoeffVo, marginUserAssetVoList []InlineResponse20019UserAssets) *InlineResponse20057 {
	this := InlineResponse20057{}
	this.Email = email
	this.MarginLevel = marginLevel
	this.TotalAssetOfBtc = totalAssetOfBtc
	this.TotalLiabilityOfBtc = totalLiabilityOfBtc
	this.TotalNetAssetOfBtc = totalNetAssetOfBtc
	this.MarginTradeCoeffVo = marginTradeCoeffVo
	this.MarginUserAssetVoList = marginUserAssetVoList
	return &this
}

// NewInlineResponse20057WithDefaults instantiates a new InlineResponse20057 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20057WithDefaults() *InlineResponse20057 {
	this := InlineResponse20057{}
	return &this
}

// GetEmail returns the Email field value
func (o *InlineResponse20057) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineResponse20057) SetEmail(v string) {
	o.Email = v
}

// GetMarginLevel returns the MarginLevel field value
func (o *InlineResponse20057) GetMarginLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetMarginLevelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarginLevel, true
}

// SetMarginLevel sets field value
func (o *InlineResponse20057) SetMarginLevel(v string) {
	o.MarginLevel = v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value
func (o *InlineResponse20057) GetTotalAssetOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalAssetOfBtc, true
}

// SetTotalAssetOfBtc sets field value
func (o *InlineResponse20057) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value
func (o *InlineResponse20057) GetTotalLiabilityOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalLiabilityOfBtc, true
}

// SetTotalLiabilityOfBtc sets field value
func (o *InlineResponse20057) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value
func (o *InlineResponse20057) GetTotalNetAssetOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalNetAssetOfBtc, true
}

// SetTotalNetAssetOfBtc sets field value
func (o *InlineResponse20057) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = v
}

// GetMarginTradeCoeffVo returns the MarginTradeCoeffVo field value
func (o *InlineResponse20057) GetMarginTradeCoeffVo() InlineResponse20057MarginTradeCoeffVo {
	if o == nil {
		var ret InlineResponse20057MarginTradeCoeffVo
		return ret
	}

	return o.MarginTradeCoeffVo
}

// GetMarginTradeCoeffVoOk returns a tuple with the MarginTradeCoeffVo field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetMarginTradeCoeffVoOk() (*InlineResponse20057MarginTradeCoeffVo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarginTradeCoeffVo, true
}

// SetMarginTradeCoeffVo sets field value
func (o *InlineResponse20057) SetMarginTradeCoeffVo(v InlineResponse20057MarginTradeCoeffVo) {
	o.MarginTradeCoeffVo = v
}

// GetMarginUserAssetVoList returns the MarginUserAssetVoList field value
func (o *InlineResponse20057) GetMarginUserAssetVoList() []InlineResponse20019UserAssets {
	if o == nil {
		var ret []InlineResponse20019UserAssets
		return ret
	}

	return o.MarginUserAssetVoList
}

// GetMarginUserAssetVoListOk returns a tuple with the MarginUserAssetVoList field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetMarginUserAssetVoListOk() (*[]InlineResponse20019UserAssets, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarginUserAssetVoList, true
}

// SetMarginUserAssetVoList sets field value
func (o *InlineResponse20057) SetMarginUserAssetVoList(v []InlineResponse20019UserAssets) {
	o.MarginUserAssetVoList = v
}

func (o InlineResponse20057) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["marginLevel"] = o.MarginLevel
	}
	if true {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if true {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if true {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}
	if true {
		toSerialize["marginTradeCoeffVo"] = o.MarginTradeCoeffVo
	}
	if true {
		toSerialize["marginUserAssetVoList"] = o.MarginUserAssetVoList
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20057 struct {
	value *InlineResponse20057
	isSet bool
}

func (v NullableInlineResponse20057) Get() *InlineResponse20057 {
	return v.value
}

func (v *NullableInlineResponse20057) Set(val *InlineResponse20057) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20057) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20057) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20057(val *InlineResponse20057) *NullableInlineResponse20057 {
	return &NullableInlineResponse20057{value: val, isSet: true}
}

func (v NullableInlineResponse20057) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20057) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

