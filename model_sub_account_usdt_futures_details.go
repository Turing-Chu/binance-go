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

// SubAccountUSDTFuturesDetails struct for SubAccountUSDTFuturesDetails
type SubAccountUSDTFuturesDetails struct {
	FutureAccountResp SubAccountUSDTFuturesDetailsFutureAccountResp `json:"futureAccountResp"`
}

// NewSubAccountUSDTFuturesDetails instantiates a new SubAccountUSDTFuturesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountUSDTFuturesDetails(futureAccountResp SubAccountUSDTFuturesDetailsFutureAccountResp) *SubAccountUSDTFuturesDetails {
	this := SubAccountUSDTFuturesDetails{}
	this.FutureAccountResp = futureAccountResp
	return &this
}

// NewSubAccountUSDTFuturesDetailsWithDefaults instantiates a new SubAccountUSDTFuturesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountUSDTFuturesDetailsWithDefaults() *SubAccountUSDTFuturesDetails {
	this := SubAccountUSDTFuturesDetails{}
	return &this
}

// GetFutureAccountResp returns the FutureAccountResp field value
func (o *SubAccountUSDTFuturesDetails) GetFutureAccountResp() SubAccountUSDTFuturesDetailsFutureAccountResp {
	if o == nil {
		var ret SubAccountUSDTFuturesDetailsFutureAccountResp
		return ret
	}

	return o.FutureAccountResp
}

// GetFutureAccountRespOk returns a tuple with the FutureAccountResp field value
// and a boolean to check if the value has been set.
func (o *SubAccountUSDTFuturesDetails) GetFutureAccountRespOk() (*SubAccountUSDTFuturesDetailsFutureAccountResp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FutureAccountResp, true
}

// SetFutureAccountResp sets field value
func (o *SubAccountUSDTFuturesDetails) SetFutureAccountResp(v SubAccountUSDTFuturesDetailsFutureAccountResp) {
	o.FutureAccountResp = v
}

func (o SubAccountUSDTFuturesDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["futureAccountResp"] = o.FutureAccountResp
	}
	return json.Marshal(toSerialize)
}

type NullableSubAccountUSDTFuturesDetails struct {
	value *SubAccountUSDTFuturesDetails
	isSet bool
}

func (v NullableSubAccountUSDTFuturesDetails) Get() *SubAccountUSDTFuturesDetails {
	return v.value
}

func (v *NullableSubAccountUSDTFuturesDetails) Set(val *SubAccountUSDTFuturesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountUSDTFuturesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountUSDTFuturesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccountUSDTFuturesDetails(val *SubAccountUSDTFuturesDetails) *NullableSubAccountUSDTFuturesDetails {
	return &NullableSubAccountUSDTFuturesDetails{value: val, isSet: true}
}

func (v NullableSubAccountUSDTFuturesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountUSDTFuturesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


