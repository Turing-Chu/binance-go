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

// InlineResponse2009Rows struct for InlineResponse2009Rows
type InlineResponse2009Rows struct {
	Amount string `json:"amount"`
	Asset string `json:"asset"`
	Status string `json:"status"`
	Timestamp int64 `json:"timestamp"`
	TxId int64 `json:"txId"`
	Type string `json:"type"`
}

// NewInlineResponse2009Rows instantiates a new InlineResponse2009Rows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009Rows(amount string, asset string, status string, timestamp int64, txId int64, type_ string) *InlineResponse2009Rows {
	this := InlineResponse2009Rows{}
	this.Amount = amount
	this.Asset = asset
	this.Status = status
	this.Timestamp = timestamp
	this.TxId = txId
	this.Type = type_
	return &this
}

// NewInlineResponse2009RowsWithDefaults instantiates a new InlineResponse2009Rows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009RowsWithDefaults() *InlineResponse2009Rows {
	this := InlineResponse2009Rows{}
	return &this
}

// GetAmount returns the Amount field value
func (o *InlineResponse2009Rows) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Rows) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse2009Rows) SetAmount(v string) {
	o.Amount = v
}

// GetAsset returns the Asset field value
func (o *InlineResponse2009Rows) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Rows) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse2009Rows) SetAsset(v string) {
	o.Asset = v
}

// GetStatus returns the Status field value
func (o *InlineResponse2009Rows) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Rows) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse2009Rows) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
func (o *InlineResponse2009Rows) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Rows) GetTimestampOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *InlineResponse2009Rows) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetTxId returns the TxId field value
func (o *InlineResponse2009Rows) GetTxId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Rows) GetTxIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *InlineResponse2009Rows) SetTxId(v int64) {
	o.TxId = v
}

// GetType returns the Type field value
func (o *InlineResponse2009Rows) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Rows) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse2009Rows) SetType(v string) {
	o.Type = v
}

func (o InlineResponse2009Rows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["txId"] = o.TxId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009Rows struct {
	value *InlineResponse2009Rows
	isSet bool
}

func (v NullableInlineResponse2009Rows) Get() *InlineResponse2009Rows {
	return v.value
}

func (v *NullableInlineResponse2009Rows) Set(val *InlineResponse2009Rows) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009Rows) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009Rows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009Rows(val *InlineResponse2009Rows) *NullableInlineResponse2009Rows {
	return &NullableInlineResponse2009Rows{value: val, isSet: true}
}

func (v NullableInlineResponse2009Rows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009Rows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

