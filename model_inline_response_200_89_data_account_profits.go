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

// InlineResponse20089DataAccountProfits struct for InlineResponse20089DataAccountProfits
type InlineResponse20089DataAccountProfits struct {
	// Mining date
	Time int64 `json:"time"`
	// 0:Mining Wallet,5:Mining Address,7:Pool Savings,8:Transferred,31:Income Transfer ,32:Hashrate Resale-Mining Wallet 33:Hashrate Resale-Pool Savings
	Type int64 `json:"type"`
	// Transferred Hashrate
	HashTransfer int32 `json:"hashTransfer"`
	// Transferred Income
	TransferAmount float32 `json:"transferAmount"`
	// Daily Hashrate
	DayHashRate int64 `json:"dayHashRate"`
	// Earnings Amount
	ProfitAmount float64 `json:"profitAmount"`
	// Coin Type
	CoinName string `json:"coinName"`
	// Status：0:Unpaid, 1:Paying  2：Paid
	Status int32 `json:"status"`
}

// NewInlineResponse20089DataAccountProfits instantiates a new InlineResponse20089DataAccountProfits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089DataAccountProfits(time int64, type_ int64, hashTransfer int32, transferAmount float32, dayHashRate int64, profitAmount float64, coinName string, status int32) *InlineResponse20089DataAccountProfits {
	this := InlineResponse20089DataAccountProfits{}
	this.Time = time
	this.Type = type_
	this.HashTransfer = hashTransfer
	this.TransferAmount = transferAmount
	this.DayHashRate = dayHashRate
	this.ProfitAmount = profitAmount
	this.CoinName = coinName
	this.Status = status
	return &this
}

// NewInlineResponse20089DataAccountProfitsWithDefaults instantiates a new InlineResponse20089DataAccountProfits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089DataAccountProfitsWithDefaults() *InlineResponse20089DataAccountProfits {
	this := InlineResponse20089DataAccountProfits{}
	return &this
}

// GetTime returns the Time field value
func (o *InlineResponse20089DataAccountProfits) GetTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089DataAccountProfits) GetTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *InlineResponse20089DataAccountProfits) SetTime(v int64) {
	o.Time = v
}

// GetType returns the Type field value
func (o *InlineResponse20089DataAccountProfits) GetType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089DataAccountProfits) GetTypeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20089DataAccountProfits) SetType(v int64) {
	o.Type = v
}

// GetHashTransfer returns the HashTransfer field value
func (o *InlineResponse20089DataAccountProfits) GetHashTransfer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HashTransfer
}

// GetHashTransferOk returns a tuple with the HashTransfer field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089DataAccountProfits) GetHashTransferOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HashTransfer, true
}

// SetHashTransfer sets field value
func (o *InlineResponse20089DataAccountProfits) SetHashTransfer(v int32) {
	o.HashTransfer = v
}

// GetTransferAmount returns the TransferAmount field value
func (o *InlineResponse20089DataAccountProfits) GetTransferAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TransferAmount
}

// GetTransferAmountOk returns a tuple with the TransferAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089DataAccountProfits) GetTransferAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferAmount, true
}

// SetTransferAmount sets field value
func (o *InlineResponse20089DataAccountProfits) SetTransferAmount(v float32) {
	o.TransferAmount = v
}

// GetDayHashRate returns the DayHashRate field value
func (o *InlineResponse20089DataAccountProfits) GetDayHashRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DayHashRate
}

// GetDayHashRateOk returns a tuple with the DayHashRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089DataAccountProfits) GetDayHashRateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DayHashRate, true
}

// SetDayHashRate sets field value
func (o *InlineResponse20089DataAccountProfits) SetDayHashRate(v int64) {
	o.DayHashRate = v
}

// GetProfitAmount returns the ProfitAmount field value
func (o *InlineResponse20089DataAccountProfits) GetProfitAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ProfitAmount
}

// GetProfitAmountOk returns a tuple with the ProfitAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089DataAccountProfits) GetProfitAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProfitAmount, true
}

// SetProfitAmount sets field value
func (o *InlineResponse20089DataAccountProfits) SetProfitAmount(v float64) {
	o.ProfitAmount = v
}

// GetCoinName returns the CoinName field value
func (o *InlineResponse20089DataAccountProfits) GetCoinName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CoinName
}

// GetCoinNameOk returns a tuple with the CoinName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089DataAccountProfits) GetCoinNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CoinName, true
}

// SetCoinName sets field value
func (o *InlineResponse20089DataAccountProfits) SetCoinName(v string) {
	o.CoinName = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20089DataAccountProfits) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089DataAccountProfits) GetStatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20089DataAccountProfits) SetStatus(v int32) {
	o.Status = v
}

func (o InlineResponse20089DataAccountProfits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["hashTransfer"] = o.HashTransfer
	}
	if true {
		toSerialize["transferAmount"] = o.TransferAmount
	}
	if true {
		toSerialize["dayHashRate"] = o.DayHashRate
	}
	if true {
		toSerialize["profitAmount"] = o.ProfitAmount
	}
	if true {
		toSerialize["coinName"] = o.CoinName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20089DataAccountProfits struct {
	value *InlineResponse20089DataAccountProfits
	isSet bool
}

func (v NullableInlineResponse20089DataAccountProfits) Get() *InlineResponse20089DataAccountProfits {
	return v.value
}

func (v *NullableInlineResponse20089DataAccountProfits) Set(val *InlineResponse20089DataAccountProfits) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089DataAccountProfits) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089DataAccountProfits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089DataAccountProfits(val *InlineResponse20089DataAccountProfits) *NullableInlineResponse20089DataAccountProfits {
	return &NullableInlineResponse20089DataAccountProfits{value: val, isSet: true}
}

func (v NullableInlineResponse20089DataAccountProfits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089DataAccountProfits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


