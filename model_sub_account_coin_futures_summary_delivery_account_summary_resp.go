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

// SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp struct for SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp
type SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp struct {
	TotalMarginBalanceOfBTC string `json:"totalMarginBalanceOfBTC"`
	TotalUnrealizedProfitOfBTC string `json:"totalUnrealizedProfitOfBTC"`
	TotalWalletBalanceOfBTC string `json:"totalWalletBalanceOfBTC"`
	Asset string `json:"asset"`
	SubAccountList []SubAccountCOINFuturesSummaryDeliveryAccountSummaryRespSubAccountList `json:"subAccountList"`
}

// NewSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp instantiates a new SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp(totalMarginBalanceOfBTC string, totalUnrealizedProfitOfBTC string, totalWalletBalanceOfBTC string, asset string, subAccountList []SubAccountCOINFuturesSummaryDeliveryAccountSummaryRespSubAccountList) *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp {
	this := SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp{}
	this.TotalMarginBalanceOfBTC = totalMarginBalanceOfBTC
	this.TotalUnrealizedProfitOfBTC = totalUnrealizedProfitOfBTC
	this.TotalWalletBalanceOfBTC = totalWalletBalanceOfBTC
	this.Asset = asset
	this.SubAccountList = subAccountList
	return &this
}

// NewSubAccountCOINFuturesSummaryDeliveryAccountSummaryRespWithDefaults instantiates a new SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountCOINFuturesSummaryDeliveryAccountSummaryRespWithDefaults() *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp {
	this := SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp{}
	return &this
}

// GetTotalMarginBalanceOfBTC returns the TotalMarginBalanceOfBTC field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetTotalMarginBalanceOfBTC() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalMarginBalanceOfBTC
}

// GetTotalMarginBalanceOfBTCOk returns a tuple with the TotalMarginBalanceOfBTC field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetTotalMarginBalanceOfBTCOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalMarginBalanceOfBTC, true
}

// SetTotalMarginBalanceOfBTC sets field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) SetTotalMarginBalanceOfBTC(v string) {
	o.TotalMarginBalanceOfBTC = v
}

// GetTotalUnrealizedProfitOfBTC returns the TotalUnrealizedProfitOfBTC field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetTotalUnrealizedProfitOfBTC() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalUnrealizedProfitOfBTC
}

// GetTotalUnrealizedProfitOfBTCOk returns a tuple with the TotalUnrealizedProfitOfBTC field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetTotalUnrealizedProfitOfBTCOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalUnrealizedProfitOfBTC, true
}

// SetTotalUnrealizedProfitOfBTC sets field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) SetTotalUnrealizedProfitOfBTC(v string) {
	o.TotalUnrealizedProfitOfBTC = v
}

// GetTotalWalletBalanceOfBTC returns the TotalWalletBalanceOfBTC field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetTotalWalletBalanceOfBTC() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalWalletBalanceOfBTC
}

// GetTotalWalletBalanceOfBTCOk returns a tuple with the TotalWalletBalanceOfBTC field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetTotalWalletBalanceOfBTCOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalWalletBalanceOfBTC, true
}

// SetTotalWalletBalanceOfBTC sets field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) SetTotalWalletBalanceOfBTC(v string) {
	o.TotalWalletBalanceOfBTC = v
}

// GetAsset returns the Asset field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) SetAsset(v string) {
	o.Asset = v
}

// GetSubAccountList returns the SubAccountList field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetSubAccountList() []SubAccountCOINFuturesSummaryDeliveryAccountSummaryRespSubAccountList {
	if o == nil {
		var ret []SubAccountCOINFuturesSummaryDeliveryAccountSummaryRespSubAccountList
		return ret
	}

	return o.SubAccountList
}

// GetSubAccountListOk returns a tuple with the SubAccountList field value
// and a boolean to check if the value has been set.
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) GetSubAccountListOk() (*[]SubAccountCOINFuturesSummaryDeliveryAccountSummaryRespSubAccountList, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubAccountList, true
}

// SetSubAccountList sets field value
func (o *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) SetSubAccountList(v []SubAccountCOINFuturesSummaryDeliveryAccountSummaryRespSubAccountList) {
	o.SubAccountList = v
}

func (o SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalMarginBalanceOfBTC"] = o.TotalMarginBalanceOfBTC
	}
	if true {
		toSerialize["totalUnrealizedProfitOfBTC"] = o.TotalUnrealizedProfitOfBTC
	}
	if true {
		toSerialize["totalWalletBalanceOfBTC"] = o.TotalWalletBalanceOfBTC
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["subAccountList"] = o.SubAccountList
	}
	return json.Marshal(toSerialize)
}

type NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp struct {
	value *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp
	isSet bool
}

func (v NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) Get() *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp {
	return v.value
}

func (v *NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) Set(val *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp(val *SubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) *NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp {
	return &NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp{value: val, isSet: true}
}

func (v NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountCOINFuturesSummaryDeliveryAccountSummaryResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

