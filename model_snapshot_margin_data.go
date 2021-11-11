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

// SnapshotMarginData struct for SnapshotMarginData
type SnapshotMarginData struct {
	MarginLevel string `json:"marginLevel"`
	TotalAssetOfBtc string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc string `json:"totalNetAssetOfBtc"`
	UserAssets []SnapshotMarginDataUserAssets `json:"userAssets"`
}

// NewSnapshotMarginData instantiates a new SnapshotMarginData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotMarginData(marginLevel string, totalAssetOfBtc string, totalLiabilityOfBtc string, totalNetAssetOfBtc string, userAssets []SnapshotMarginDataUserAssets) *SnapshotMarginData {
	this := SnapshotMarginData{}
	this.MarginLevel = marginLevel
	this.TotalAssetOfBtc = totalAssetOfBtc
	this.TotalLiabilityOfBtc = totalLiabilityOfBtc
	this.TotalNetAssetOfBtc = totalNetAssetOfBtc
	this.UserAssets = userAssets
	return &this
}

// NewSnapshotMarginDataWithDefaults instantiates a new SnapshotMarginData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotMarginDataWithDefaults() *SnapshotMarginData {
	this := SnapshotMarginData{}
	return &this
}

// GetMarginLevel returns the MarginLevel field value
func (o *SnapshotMarginData) GetMarginLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarginLevel
}

// GetMarginLevelOk returns a tuple with the MarginLevel field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginData) GetMarginLevelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarginLevel, true
}

// SetMarginLevel sets field value
func (o *SnapshotMarginData) SetMarginLevel(v string) {
	o.MarginLevel = v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value
func (o *SnapshotMarginData) GetTotalAssetOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginData) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalAssetOfBtc, true
}

// SetTotalAssetOfBtc sets field value
func (o *SnapshotMarginData) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value
func (o *SnapshotMarginData) GetTotalLiabilityOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginData) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalLiabilityOfBtc, true
}

// SetTotalLiabilityOfBtc sets field value
func (o *SnapshotMarginData) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value
func (o *SnapshotMarginData) GetTotalNetAssetOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginData) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalNetAssetOfBtc, true
}

// SetTotalNetAssetOfBtc sets field value
func (o *SnapshotMarginData) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = v
}

// GetUserAssets returns the UserAssets field value
func (o *SnapshotMarginData) GetUserAssets() []SnapshotMarginDataUserAssets {
	if o == nil {
		var ret []SnapshotMarginDataUserAssets
		return ret
	}

	return o.UserAssets
}

// GetUserAssetsOk returns a tuple with the UserAssets field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginData) GetUserAssetsOk() (*[]SnapshotMarginDataUserAssets, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserAssets, true
}

// SetUserAssets sets field value
func (o *SnapshotMarginData) SetUserAssets(v []SnapshotMarginDataUserAssets) {
	o.UserAssets = v
}

func (o SnapshotMarginData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["userAssets"] = o.UserAssets
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotMarginData struct {
	value *SnapshotMarginData
	isSet bool
}

func (v NullableSnapshotMarginData) Get() *SnapshotMarginData {
	return v.value
}

func (v *NullableSnapshotMarginData) Set(val *SnapshotMarginData) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotMarginData) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotMarginData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotMarginData(val *SnapshotMarginData) *NullableSnapshotMarginData {
	return &NullableSnapshotMarginData{value: val, isSet: true}
}

func (v NullableSnapshotMarginData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotMarginData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


