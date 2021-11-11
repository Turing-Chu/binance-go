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

// InlineResponse20061 struct for InlineResponse20061
type InlineResponse20061 struct {
	TotalInitialMargin string `json:"totalInitialMargin"`
	TotalMaintenanceMargin string `json:"totalMaintenanceMargin"`
	TotalMarginBalance string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit string `json:"totalUnrealizedProfit"`
	TotalWalletBalance string `json:"totalWalletBalance"`
	Asset string `json:"asset"`
	SubAccountList []InlineResponse20061SubAccountList `json:"subAccountList"`
}

// NewInlineResponse20061 instantiates a new InlineResponse20061 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20061(totalInitialMargin string, totalMaintenanceMargin string, totalMarginBalance string, totalOpenOrderInitialMargin string, totalPositionInitialMargin string, totalUnrealizedProfit string, totalWalletBalance string, asset string, subAccountList []InlineResponse20061SubAccountList) *InlineResponse20061 {
	this := InlineResponse20061{}
	this.TotalInitialMargin = totalInitialMargin
	this.TotalMaintenanceMargin = totalMaintenanceMargin
	this.TotalMarginBalance = totalMarginBalance
	this.TotalOpenOrderInitialMargin = totalOpenOrderInitialMargin
	this.TotalPositionInitialMargin = totalPositionInitialMargin
	this.TotalUnrealizedProfit = totalUnrealizedProfit
	this.TotalWalletBalance = totalWalletBalance
	this.Asset = asset
	this.SubAccountList = subAccountList
	return &this
}

// NewInlineResponse20061WithDefaults instantiates a new InlineResponse20061 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20061WithDefaults() *InlineResponse20061 {
	this := InlineResponse20061{}
	return &this
}

// GetTotalInitialMargin returns the TotalInitialMargin field value
func (o *InlineResponse20061) GetTotalInitialMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalInitialMargin, true
}

// SetTotalInitialMargin sets field value
func (o *InlineResponse20061) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = v
}

// GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field value
func (o *InlineResponse20061) GetTotalMaintenanceMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalMaintenanceMargin
}

// GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetTotalMaintenanceMarginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalMaintenanceMargin, true
}

// SetTotalMaintenanceMargin sets field value
func (o *InlineResponse20061) SetTotalMaintenanceMargin(v string) {
	o.TotalMaintenanceMargin = v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value
func (o *InlineResponse20061) GetTotalMarginBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalMarginBalance, true
}

// SetTotalMarginBalance sets field value
func (o *InlineResponse20061) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value
func (o *InlineResponse20061) GetTotalOpenOrderInitialMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalOpenOrderInitialMargin, true
}

// SetTotalOpenOrderInitialMargin sets field value
func (o *InlineResponse20061) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value
func (o *InlineResponse20061) GetTotalPositionInitialMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalPositionInitialMargin, true
}

// SetTotalPositionInitialMargin sets field value
func (o *InlineResponse20061) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value
func (o *InlineResponse20061) GetTotalUnrealizedProfit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalUnrealizedProfit, true
}

// SetTotalUnrealizedProfit sets field value
func (o *InlineResponse20061) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value
func (o *InlineResponse20061) GetTotalWalletBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalWalletBalance, true
}

// SetTotalWalletBalance sets field value
func (o *InlineResponse20061) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = v
}

// GetAsset returns the Asset field value
func (o *InlineResponse20061) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20061) SetAsset(v string) {
	o.Asset = v
}

// GetSubAccountList returns the SubAccountList field value
func (o *InlineResponse20061) GetSubAccountList() []InlineResponse20061SubAccountList {
	if o == nil {
		var ret []InlineResponse20061SubAccountList
		return ret
	}

	return o.SubAccountList
}

// GetSubAccountListOk returns a tuple with the SubAccountList field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetSubAccountListOk() (*[]InlineResponse20061SubAccountList, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubAccountList, true
}

// SetSubAccountList sets field value
func (o *InlineResponse20061) SetSubAccountList(v []InlineResponse20061SubAccountList) {
	o.SubAccountList = v
}

func (o InlineResponse20061) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalInitialMargin"] = o.TotalInitialMargin
	}
	if true {
		toSerialize["totalMaintenanceMargin"] = o.TotalMaintenanceMargin
	}
	if true {
		toSerialize["totalMarginBalance"] = o.TotalMarginBalance
	}
	if true {
		toSerialize["totalOpenOrderInitialMargin"] = o.TotalOpenOrderInitialMargin
	}
	if true {
		toSerialize["totalPositionInitialMargin"] = o.TotalPositionInitialMargin
	}
	if true {
		toSerialize["totalUnrealizedProfit"] = o.TotalUnrealizedProfit
	}
	if true {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["subAccountList"] = o.SubAccountList
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20061 struct {
	value *InlineResponse20061
	isSet bool
}

func (v NullableInlineResponse20061) Get() *InlineResponse20061 {
	return v.value
}

func (v *NullableInlineResponse20061) Set(val *InlineResponse20061) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20061) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20061) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20061(val *InlineResponse20061) *NullableInlineResponse20061 {
	return &NullableInlineResponse20061{value: val, isSet: true}
}

func (v NullableInlineResponse20061) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20061) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


