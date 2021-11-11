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

// InlineResponse20074 struct for InlineResponse20074
type InlineResponse20074 struct {
	Asset string `json:"asset"`
	AvgAnnualInterestRate string `json:"avgAnnualInterestRate"`
	CanPurchase bool `json:"canPurchase"`
	CanRedeem bool `json:"canRedeem"`
	DailyInterestPerThousand string `json:"dailyInterestPerThousand"`
	Featured bool `json:"featured"`
	MinPurchaseAmount string `json:"minPurchaseAmount"`
	ProductId string `json:"productId"`
	PurchasedAmount string `json:"purchasedAmount"`
	Status string `json:"status"`
	UpLimit string `json:"upLimit"`
	UpLimitPerUser string `json:"upLimitPerUser"`
}

// NewInlineResponse20074 instantiates a new InlineResponse20074 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20074(asset string, avgAnnualInterestRate string, canPurchase bool, canRedeem bool, dailyInterestPerThousand string, featured bool, minPurchaseAmount string, productId string, purchasedAmount string, status string, upLimit string, upLimitPerUser string) *InlineResponse20074 {
	this := InlineResponse20074{}
	this.Asset = asset
	this.AvgAnnualInterestRate = avgAnnualInterestRate
	this.CanPurchase = canPurchase
	this.CanRedeem = canRedeem
	this.DailyInterestPerThousand = dailyInterestPerThousand
	this.Featured = featured
	this.MinPurchaseAmount = minPurchaseAmount
	this.ProductId = productId
	this.PurchasedAmount = purchasedAmount
	this.Status = status
	this.UpLimit = upLimit
	this.UpLimitPerUser = upLimitPerUser
	return &this
}

// NewInlineResponse20074WithDefaults instantiates a new InlineResponse20074 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20074WithDefaults() *InlineResponse20074 {
	this := InlineResponse20074{}
	return &this
}

// GetAsset returns the Asset field value
func (o *InlineResponse20074) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20074) SetAsset(v string) {
	o.Asset = v
}

// GetAvgAnnualInterestRate returns the AvgAnnualInterestRate field value
func (o *InlineResponse20074) GetAvgAnnualInterestRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvgAnnualInterestRate
}

// GetAvgAnnualInterestRateOk returns a tuple with the AvgAnnualInterestRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetAvgAnnualInterestRateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvgAnnualInterestRate, true
}

// SetAvgAnnualInterestRate sets field value
func (o *InlineResponse20074) SetAvgAnnualInterestRate(v string) {
	o.AvgAnnualInterestRate = v
}

// GetCanPurchase returns the CanPurchase field value
func (o *InlineResponse20074) GetCanPurchase() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanPurchase
}

// GetCanPurchaseOk returns a tuple with the CanPurchase field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetCanPurchaseOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CanPurchase, true
}

// SetCanPurchase sets field value
func (o *InlineResponse20074) SetCanPurchase(v bool) {
	o.CanPurchase = v
}

// GetCanRedeem returns the CanRedeem field value
func (o *InlineResponse20074) GetCanRedeem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanRedeem
}

// GetCanRedeemOk returns a tuple with the CanRedeem field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetCanRedeemOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CanRedeem, true
}

// SetCanRedeem sets field value
func (o *InlineResponse20074) SetCanRedeem(v bool) {
	o.CanRedeem = v
}

// GetDailyInterestPerThousand returns the DailyInterestPerThousand field value
func (o *InlineResponse20074) GetDailyInterestPerThousand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DailyInterestPerThousand
}

// GetDailyInterestPerThousandOk returns a tuple with the DailyInterestPerThousand field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetDailyInterestPerThousandOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DailyInterestPerThousand, true
}

// SetDailyInterestPerThousand sets field value
func (o *InlineResponse20074) SetDailyInterestPerThousand(v string) {
	o.DailyInterestPerThousand = v
}

// GetFeatured returns the Featured field value
func (o *InlineResponse20074) GetFeatured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetFeaturedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Featured, true
}

// SetFeatured sets field value
func (o *InlineResponse20074) SetFeatured(v bool) {
	o.Featured = v
}

// GetMinPurchaseAmount returns the MinPurchaseAmount field value
func (o *InlineResponse20074) GetMinPurchaseAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinPurchaseAmount
}

// GetMinPurchaseAmountOk returns a tuple with the MinPurchaseAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetMinPurchaseAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinPurchaseAmount, true
}

// SetMinPurchaseAmount sets field value
func (o *InlineResponse20074) SetMinPurchaseAmount(v string) {
	o.MinPurchaseAmount = v
}

// GetProductId returns the ProductId field value
func (o *InlineResponse20074) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetProductIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *InlineResponse20074) SetProductId(v string) {
	o.ProductId = v
}

// GetPurchasedAmount returns the PurchasedAmount field value
func (o *InlineResponse20074) GetPurchasedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchasedAmount
}

// GetPurchasedAmountOk returns a tuple with the PurchasedAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetPurchasedAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PurchasedAmount, true
}

// SetPurchasedAmount sets field value
func (o *InlineResponse20074) SetPurchasedAmount(v string) {
	o.PurchasedAmount = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20074) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20074) SetStatus(v string) {
	o.Status = v
}

// GetUpLimit returns the UpLimit field value
func (o *InlineResponse20074) GetUpLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpLimit
}

// GetUpLimitOk returns a tuple with the UpLimit field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetUpLimitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpLimit, true
}

// SetUpLimit sets field value
func (o *InlineResponse20074) SetUpLimit(v string) {
	o.UpLimit = v
}

// GetUpLimitPerUser returns the UpLimitPerUser field value
func (o *InlineResponse20074) GetUpLimitPerUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpLimitPerUser
}

// GetUpLimitPerUserOk returns a tuple with the UpLimitPerUser field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20074) GetUpLimitPerUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpLimitPerUser, true
}

// SetUpLimitPerUser sets field value
func (o *InlineResponse20074) SetUpLimitPerUser(v string) {
	o.UpLimitPerUser = v
}

func (o InlineResponse20074) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["avgAnnualInterestRate"] = o.AvgAnnualInterestRate
	}
	if true {
		toSerialize["canPurchase"] = o.CanPurchase
	}
	if true {
		toSerialize["canRedeem"] = o.CanRedeem
	}
	if true {
		toSerialize["dailyInterestPerThousand"] = o.DailyInterestPerThousand
	}
	if true {
		toSerialize["featured"] = o.Featured
	}
	if true {
		toSerialize["minPurchaseAmount"] = o.MinPurchaseAmount
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["purchasedAmount"] = o.PurchasedAmount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["upLimit"] = o.UpLimit
	}
	if true {
		toSerialize["upLimitPerUser"] = o.UpLimitPerUser
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20074 struct {
	value *InlineResponse20074
	isSet bool
}

func (v NullableInlineResponse20074) Get() *InlineResponse20074 {
	return v.value
}

func (v *NullableInlineResponse20074) Set(val *InlineResponse20074) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20074) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20074) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20074(val *InlineResponse20074) *NullableInlineResponse20074 {
	return &NullableInlineResponse20074{value: val, isSet: true}
}

func (v NullableInlineResponse20074) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20074) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

