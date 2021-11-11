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

// InlineResponse20037UserAssetDribbletDetails struct for InlineResponse20037UserAssetDribbletDetails
type InlineResponse20037UserAssetDribbletDetails struct {
	TransId int64 `json:"transId"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	Amount string `json:"amount"`
	OperateTime int64 `json:"operateTime"`
	TransferedAmount string `json:"transferedAmount"`
	FromAsset string `json:"fromAsset"`
}

// NewInlineResponse20037UserAssetDribbletDetails instantiates a new InlineResponse20037UserAssetDribbletDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20037UserAssetDribbletDetails(transId int64, serviceChargeAmount string, amount string, operateTime int64, transferedAmount string, fromAsset string) *InlineResponse20037UserAssetDribbletDetails {
	this := InlineResponse20037UserAssetDribbletDetails{}
	this.TransId = transId
	this.ServiceChargeAmount = serviceChargeAmount
	this.Amount = amount
	this.OperateTime = operateTime
	this.TransferedAmount = transferedAmount
	this.FromAsset = fromAsset
	return &this
}

// NewInlineResponse20037UserAssetDribbletDetailsWithDefaults instantiates a new InlineResponse20037UserAssetDribbletDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20037UserAssetDribbletDetailsWithDefaults() *InlineResponse20037UserAssetDribbletDetails {
	this := InlineResponse20037UserAssetDribbletDetails{}
	return &this
}

// GetTransId returns the TransId field value
func (o *InlineResponse20037UserAssetDribbletDetails) GetTransId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TransId
}

// GetTransIdOk returns a tuple with the TransId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20037UserAssetDribbletDetails) GetTransIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransId, true
}

// SetTransId sets field value
func (o *InlineResponse20037UserAssetDribbletDetails) SetTransId(v int64) {
	o.TransId = v
}

// GetServiceChargeAmount returns the ServiceChargeAmount field value
func (o *InlineResponse20037UserAssetDribbletDetails) GetServiceChargeAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceChargeAmount
}

// GetServiceChargeAmountOk returns a tuple with the ServiceChargeAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20037UserAssetDribbletDetails) GetServiceChargeAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceChargeAmount, true
}

// SetServiceChargeAmount sets field value
func (o *InlineResponse20037UserAssetDribbletDetails) SetServiceChargeAmount(v string) {
	o.ServiceChargeAmount = v
}

// GetAmount returns the Amount field value
func (o *InlineResponse20037UserAssetDribbletDetails) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20037UserAssetDribbletDetails) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse20037UserAssetDribbletDetails) SetAmount(v string) {
	o.Amount = v
}

// GetOperateTime returns the OperateTime field value
func (o *InlineResponse20037UserAssetDribbletDetails) GetOperateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OperateTime
}

// GetOperateTimeOk returns a tuple with the OperateTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20037UserAssetDribbletDetails) GetOperateTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OperateTime, true
}

// SetOperateTime sets field value
func (o *InlineResponse20037UserAssetDribbletDetails) SetOperateTime(v int64) {
	o.OperateTime = v
}

// GetTransferedAmount returns the TransferedAmount field value
func (o *InlineResponse20037UserAssetDribbletDetails) GetTransferedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferedAmount
}

// GetTransferedAmountOk returns a tuple with the TransferedAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20037UserAssetDribbletDetails) GetTransferedAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferedAmount, true
}

// SetTransferedAmount sets field value
func (o *InlineResponse20037UserAssetDribbletDetails) SetTransferedAmount(v string) {
	o.TransferedAmount = v
}

// GetFromAsset returns the FromAsset field value
func (o *InlineResponse20037UserAssetDribbletDetails) GetFromAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20037UserAssetDribbletDetails) GetFromAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FromAsset, true
}

// SetFromAsset sets field value
func (o *InlineResponse20037UserAssetDribbletDetails) SetFromAsset(v string) {
	o.FromAsset = v
}

func (o InlineResponse20037UserAssetDribbletDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transId"] = o.TransId
	}
	if true {
		toSerialize["serviceChargeAmount"] = o.ServiceChargeAmount
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["operateTime"] = o.OperateTime
	}
	if true {
		toSerialize["transferedAmount"] = o.TransferedAmount
	}
	if true {
		toSerialize["fromAsset"] = o.FromAsset
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20037UserAssetDribbletDetails struct {
	value *InlineResponse20037UserAssetDribbletDetails
	isSet bool
}

func (v NullableInlineResponse20037UserAssetDribbletDetails) Get() *InlineResponse20037UserAssetDribbletDetails {
	return v.value
}

func (v *NullableInlineResponse20037UserAssetDribbletDetails) Set(val *InlineResponse20037UserAssetDribbletDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20037UserAssetDribbletDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20037UserAssetDribbletDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20037UserAssetDribbletDetails(val *InlineResponse20037UserAssetDribbletDetails) *NullableInlineResponse20037UserAssetDribbletDetails {
	return &NullableInlineResponse20037UserAssetDribbletDetails{value: val, isSet: true}
}

func (v NullableInlineResponse20037UserAssetDribbletDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20037UserAssetDribbletDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

