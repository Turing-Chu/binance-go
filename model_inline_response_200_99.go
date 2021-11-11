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

// InlineResponse20099 struct for InlineResponse20099
type InlineResponse20099 struct {
	Id float64 `json:"id"`
	TokenName string `json:"tokenName"`
	// Subscription amount
	Amount string `json:"amount"`
	// NAV price of subscription
	Nav string `json:"nav"`
	// Subscription fee in usdt
	Fee string `json:"fee"`
	// Subscription cost in usdt
	TotalCharge string `json:"totalCharge"`
	Timestamp int64 `json:"timestamp"`
}

// NewInlineResponse20099 instantiates a new InlineResponse20099 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20099(id float64, tokenName string, amount string, nav string, fee string, totalCharge string, timestamp int64) *InlineResponse20099 {
	this := InlineResponse20099{}
	this.Id = id
	this.TokenName = tokenName
	this.Amount = amount
	this.Nav = nav
	this.Fee = fee
	this.TotalCharge = totalCharge
	this.Timestamp = timestamp
	return &this
}

// NewInlineResponse20099WithDefaults instantiates a new InlineResponse20099 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20099WithDefaults() *InlineResponse20099 {
	this := InlineResponse20099{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse20099) GetId() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20099) GetIdOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse20099) SetId(v float64) {
	o.Id = v
}

// GetTokenName returns the TokenName field value
func (o *InlineResponse20099) GetTokenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20099) GetTokenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenName, true
}

// SetTokenName sets field value
func (o *InlineResponse20099) SetTokenName(v string) {
	o.TokenName = v
}

// GetAmount returns the Amount field value
func (o *InlineResponse20099) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20099) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse20099) SetAmount(v string) {
	o.Amount = v
}

// GetNav returns the Nav field value
func (o *InlineResponse20099) GetNav() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nav
}

// GetNavOk returns a tuple with the Nav field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20099) GetNavOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nav, true
}

// SetNav sets field value
func (o *InlineResponse20099) SetNav(v string) {
	o.Nav = v
}

// GetFee returns the Fee field value
func (o *InlineResponse20099) GetFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20099) GetFeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *InlineResponse20099) SetFee(v string) {
	o.Fee = v
}

// GetTotalCharge returns the TotalCharge field value
func (o *InlineResponse20099) GetTotalCharge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalCharge
}

// GetTotalChargeOk returns a tuple with the TotalCharge field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20099) GetTotalChargeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCharge, true
}

// SetTotalCharge sets field value
func (o *InlineResponse20099) SetTotalCharge(v string) {
	o.TotalCharge = v
}

// GetTimestamp returns the Timestamp field value
func (o *InlineResponse20099) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20099) GetTimestampOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *InlineResponse20099) SetTimestamp(v int64) {
	o.Timestamp = v
}

func (o InlineResponse20099) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["tokenName"] = o.TokenName
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["nav"] = o.Nav
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["totalCharge"] = o.TotalCharge
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20099 struct {
	value *InlineResponse20099
	isSet bool
}

func (v NullableInlineResponse20099) Get() *InlineResponse20099 {
	return v.value
}

func (v *NullableInlineResponse20099) Set(val *InlineResponse20099) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20099) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20099) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20099(val *InlineResponse20099) *NullableInlineResponse20099 {
	return &NullableInlineResponse20099{value: val, isSet: true}
}

func (v NullableInlineResponse20099) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20099) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

