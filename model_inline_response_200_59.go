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

// InlineResponse20059 struct for InlineResponse20059
type InlineResponse20059 struct {
	Email string `json:"email"`
	IsFuturesEnabled bool `json:"isFuturesEnabled"`
}

// NewInlineResponse20059 instantiates a new InlineResponse20059 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20059(email string, isFuturesEnabled bool) *InlineResponse20059 {
	this := InlineResponse20059{}
	this.Email = email
	this.IsFuturesEnabled = isFuturesEnabled
	return &this
}

// NewInlineResponse20059WithDefaults instantiates a new InlineResponse20059 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20059WithDefaults() *InlineResponse20059 {
	this := InlineResponse20059{}
	return &this
}

// GetEmail returns the Email field value
func (o *InlineResponse20059) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20059) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineResponse20059) SetEmail(v string) {
	o.Email = v
}

// GetIsFuturesEnabled returns the IsFuturesEnabled field value
func (o *InlineResponse20059) GetIsFuturesEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFuturesEnabled
}

// GetIsFuturesEnabledOk returns a tuple with the IsFuturesEnabled field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20059) GetIsFuturesEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsFuturesEnabled, true
}

// SetIsFuturesEnabled sets field value
func (o *InlineResponse20059) SetIsFuturesEnabled(v bool) {
	o.IsFuturesEnabled = v
}

func (o InlineResponse20059) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["isFuturesEnabled"] = o.IsFuturesEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20059 struct {
	value *InlineResponse20059
	isSet bool
}

func (v NullableInlineResponse20059) Get() *InlineResponse20059 {
	return v.value
}

func (v *NullableInlineResponse20059) Set(val *InlineResponse20059) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20059) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20059) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20059(val *InlineResponse20059) *NullableInlineResponse20059 {
	return &NullableInlineResponse20059{value: val, isSet: true}
}

func (v NullableInlineResponse20059) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20059) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


