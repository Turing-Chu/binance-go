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

// InlineResponse20032 struct for InlineResponse20032
type InlineResponse20032 struct {
	Amount string `json:"amount"`
	Coin string `json:"coin"`
	Network string `json:"network"`
	Status int32 `json:"status"`
	Address string `json:"address"`
	AddressTag string `json:"addressTag"`
	TxId string `json:"txId"`
	InsertTime int64 `json:"insertTime"`
	TransferType int32 `json:"transferType"`
	// confirm times for unlocking
	UnlockConfirm string `json:"unlockConfirm"`
	ConfirmTimes string `json:"confirmTimes"`
}

// NewInlineResponse20032 instantiates a new InlineResponse20032 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20032(amount string, coin string, network string, status int32, address string, addressTag string, txId string, insertTime int64, transferType int32, unlockConfirm string, confirmTimes string) *InlineResponse20032 {
	this := InlineResponse20032{}
	this.Amount = amount
	this.Coin = coin
	this.Network = network
	this.Status = status
	this.Address = address
	this.AddressTag = addressTag
	this.TxId = txId
	this.InsertTime = insertTime
	this.TransferType = transferType
	this.UnlockConfirm = unlockConfirm
	this.ConfirmTimes = confirmTimes
	return &this
}

// NewInlineResponse20032WithDefaults instantiates a new InlineResponse20032 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20032WithDefaults() *InlineResponse20032 {
	this := InlineResponse20032{}
	return &this
}

// GetAmount returns the Amount field value
func (o *InlineResponse20032) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse20032) SetAmount(v string) {
	o.Amount = v
}

// GetCoin returns the Coin field value
func (o *InlineResponse20032) GetCoin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coin
}

// GetCoinOk returns a tuple with the Coin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetCoinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Coin, true
}

// SetCoin sets field value
func (o *InlineResponse20032) SetCoin(v string) {
	o.Coin = v
}

// GetNetwork returns the Network field value
func (o *InlineResponse20032) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *InlineResponse20032) SetNetwork(v string) {
	o.Network = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20032) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetStatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20032) SetStatus(v int32) {
	o.Status = v
}

// GetAddress returns the Address field value
func (o *InlineResponse20032) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InlineResponse20032) SetAddress(v string) {
	o.Address = v
}

// GetAddressTag returns the AddressTag field value
func (o *InlineResponse20032) GetAddressTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressTag
}

// GetAddressTagOk returns a tuple with the AddressTag field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetAddressTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddressTag, true
}

// SetAddressTag sets field value
func (o *InlineResponse20032) SetAddressTag(v string) {
	o.AddressTag = v
}

// GetTxId returns the TxId field value
func (o *InlineResponse20032) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetTxIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *InlineResponse20032) SetTxId(v string) {
	o.TxId = v
}

// GetInsertTime returns the InsertTime field value
func (o *InlineResponse20032) GetInsertTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InsertTime
}

// GetInsertTimeOk returns a tuple with the InsertTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetInsertTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InsertTime, true
}

// SetInsertTime sets field value
func (o *InlineResponse20032) SetInsertTime(v int64) {
	o.InsertTime = v
}

// GetTransferType returns the TransferType field value
func (o *InlineResponse20032) GetTransferType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetTransferTypeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferType, true
}

// SetTransferType sets field value
func (o *InlineResponse20032) SetTransferType(v int32) {
	o.TransferType = v
}

// GetUnlockConfirm returns the UnlockConfirm field value
func (o *InlineResponse20032) GetUnlockConfirm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnlockConfirm
}

// GetUnlockConfirmOk returns a tuple with the UnlockConfirm field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetUnlockConfirmOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnlockConfirm, true
}

// SetUnlockConfirm sets field value
func (o *InlineResponse20032) SetUnlockConfirm(v string) {
	o.UnlockConfirm = v
}

// GetConfirmTimes returns the ConfirmTimes field value
func (o *InlineResponse20032) GetConfirmTimes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmTimes
}

// GetConfirmTimesOk returns a tuple with the ConfirmTimes field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20032) GetConfirmTimesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmTimes, true
}

// SetConfirmTimes sets field value
func (o *InlineResponse20032) SetConfirmTimes(v string) {
	o.ConfirmTimes = v
}

func (o InlineResponse20032) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["coin"] = o.Coin
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["addressTag"] = o.AddressTag
	}
	if true {
		toSerialize["txId"] = o.TxId
	}
	if true {
		toSerialize["insertTime"] = o.InsertTime
	}
	if true {
		toSerialize["transferType"] = o.TransferType
	}
	if true {
		toSerialize["unlockConfirm"] = o.UnlockConfirm
	}
	if true {
		toSerialize["confirmTimes"] = o.ConfirmTimes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20032 struct {
	value *InlineResponse20032
	isSet bool
}

func (v NullableInlineResponse20032) Get() *InlineResponse20032 {
	return v.value
}

func (v *NullableInlineResponse20032) Set(val *InlineResponse20032) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20032) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20032) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20032(val *InlineResponse20032) *NullableInlineResponse20032 {
	return &NullableInlineResponse20032{value: val, isSet: true}
}

func (v NullableInlineResponse20032) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20032) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


