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

// InlineResponse20089Data struct for InlineResponse20089Data
type InlineResponse20089Data struct {
	AccountProfits []InlineResponse20089DataAccountProfits `json:"accountProfits"`
	// Total Rows
	TotalNum int64 `json:"totalNum"`
	// Rows per page
	PageSize int64 `json:"pageSize"`
}

// NewInlineResponse20089Data instantiates a new InlineResponse20089Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089Data(accountProfits []InlineResponse20089DataAccountProfits, totalNum int64, pageSize int64) *InlineResponse20089Data {
	this := InlineResponse20089Data{}
	this.AccountProfits = accountProfits
	this.TotalNum = totalNum
	this.PageSize = pageSize
	return &this
}

// NewInlineResponse20089DataWithDefaults instantiates a new InlineResponse20089Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089DataWithDefaults() *InlineResponse20089Data {
	this := InlineResponse20089Data{}
	return &this
}

// GetAccountProfits returns the AccountProfits field value
func (o *InlineResponse20089Data) GetAccountProfits() []InlineResponse20089DataAccountProfits {
	if o == nil {
		var ret []InlineResponse20089DataAccountProfits
		return ret
	}

	return o.AccountProfits
}

// GetAccountProfitsOk returns a tuple with the AccountProfits field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089Data) GetAccountProfitsOk() (*[]InlineResponse20089DataAccountProfits, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountProfits, true
}

// SetAccountProfits sets field value
func (o *InlineResponse20089Data) SetAccountProfits(v []InlineResponse20089DataAccountProfits) {
	o.AccountProfits = v
}

// GetTotalNum returns the TotalNum field value
func (o *InlineResponse20089Data) GetTotalNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089Data) GetTotalNumOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalNum, true
}

// SetTotalNum sets field value
func (o *InlineResponse20089Data) SetTotalNum(v int64) {
	o.TotalNum = v
}

// GetPageSize returns the PageSize field value
func (o *InlineResponse20089Data) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20089Data) GetPageSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *InlineResponse20089Data) SetPageSize(v int64) {
	o.PageSize = v
}

func (o InlineResponse20089Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountProfits"] = o.AccountProfits
	}
	if true {
		toSerialize["totalNum"] = o.TotalNum
	}
	if true {
		toSerialize["pageSize"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20089Data struct {
	value *InlineResponse20089Data
	isSet bool
}

func (v NullableInlineResponse20089Data) Get() *InlineResponse20089Data {
	return v.value
}

func (v *NullableInlineResponse20089Data) Set(val *InlineResponse20089Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089Data(val *InlineResponse20089Data) *NullableInlineResponse20089Data {
	return &NullableInlineResponse20089Data{value: val, isSet: true}
}

func (v NullableInlineResponse20089Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


