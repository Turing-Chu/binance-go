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

// InlineResponse20039Rows struct for InlineResponse20039Rows
type InlineResponse20039Rows struct {
	Id int64 `json:"id"`
	Amount string `json:"amount"`
	Asset string `json:"asset"`
	DivTime int64 `json:"divTime"`
	EnInfo string `json:"enInfo"`
	TranId int64 `json:"tranId"`
}

// NewInlineResponse20039Rows instantiates a new InlineResponse20039Rows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20039Rows(id int64, amount string, asset string, divTime int64, enInfo string, tranId int64) *InlineResponse20039Rows {
	this := InlineResponse20039Rows{}
	this.Id = id
	this.Amount = amount
	this.Asset = asset
	this.DivTime = divTime
	this.EnInfo = enInfo
	this.TranId = tranId
	return &this
}

// NewInlineResponse20039RowsWithDefaults instantiates a new InlineResponse20039Rows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20039RowsWithDefaults() *InlineResponse20039Rows {
	this := InlineResponse20039Rows{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse20039Rows) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20039Rows) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse20039Rows) SetId(v int64) {
	o.Id = v
}

// GetAmount returns the Amount field value
func (o *InlineResponse20039Rows) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20039Rows) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse20039Rows) SetAmount(v string) {
	o.Amount = v
}

// GetAsset returns the Asset field value
func (o *InlineResponse20039Rows) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20039Rows) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20039Rows) SetAsset(v string) {
	o.Asset = v
}

// GetDivTime returns the DivTime field value
func (o *InlineResponse20039Rows) GetDivTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DivTime
}

// GetDivTimeOk returns a tuple with the DivTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20039Rows) GetDivTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DivTime, true
}

// SetDivTime sets field value
func (o *InlineResponse20039Rows) SetDivTime(v int64) {
	o.DivTime = v
}

// GetEnInfo returns the EnInfo field value
func (o *InlineResponse20039Rows) GetEnInfo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnInfo
}

// GetEnInfoOk returns a tuple with the EnInfo field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20039Rows) GetEnInfoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnInfo, true
}

// SetEnInfo sets field value
func (o *InlineResponse20039Rows) SetEnInfo(v string) {
	o.EnInfo = v
}

// GetTranId returns the TranId field value
func (o *InlineResponse20039Rows) GetTranId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20039Rows) GetTranIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TranId, true
}

// SetTranId sets field value
func (o *InlineResponse20039Rows) SetTranId(v int64) {
	o.TranId = v
}

func (o InlineResponse20039Rows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["divTime"] = o.DivTime
	}
	if true {
		toSerialize["enInfo"] = o.EnInfo
	}
	if true {
		toSerialize["tranId"] = o.TranId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20039Rows struct {
	value *InlineResponse20039Rows
	isSet bool
}

func (v NullableInlineResponse20039Rows) Get() *InlineResponse20039Rows {
	return v.value
}

func (v *NullableInlineResponse20039Rows) Set(val *InlineResponse20039Rows) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20039Rows) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20039Rows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20039Rows(val *InlineResponse20039Rows) *NullableInlineResponse20039Rows {
	return &NullableInlineResponse20039Rows{value: val, isSet: true}
}

func (v NullableInlineResponse20039Rows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20039Rows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


