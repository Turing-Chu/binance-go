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

// InlineResponse2001RateLimits struct for InlineResponse2001RateLimits
type InlineResponse2001RateLimits struct {
	RateLimitType string `json:"rateLimitType"`
	Interval string `json:"interval"`
	IntervalNum int32 `json:"intervalNum"`
	Limit int32 `json:"limit"`
}

// NewInlineResponse2001RateLimits instantiates a new InlineResponse2001RateLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001RateLimits(rateLimitType string, interval string, intervalNum int32, limit int32) *InlineResponse2001RateLimits {
	this := InlineResponse2001RateLimits{}
	this.RateLimitType = rateLimitType
	this.Interval = interval
	this.IntervalNum = intervalNum
	this.Limit = limit
	return &this
}

// NewInlineResponse2001RateLimitsWithDefaults instantiates a new InlineResponse2001RateLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001RateLimitsWithDefaults() *InlineResponse2001RateLimits {
	this := InlineResponse2001RateLimits{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value
func (o *InlineResponse2001RateLimits) GetRateLimitType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001RateLimits) GetRateLimitTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RateLimitType, true
}

// SetRateLimitType sets field value
func (o *InlineResponse2001RateLimits) SetRateLimitType(v string) {
	o.RateLimitType = v
}

// GetInterval returns the Interval field value
func (o *InlineResponse2001RateLimits) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001RateLimits) GetIntervalOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *InlineResponse2001RateLimits) SetInterval(v string) {
	o.Interval = v
}

// GetIntervalNum returns the IntervalNum field value
func (o *InlineResponse2001RateLimits) GetIntervalNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001RateLimits) GetIntervalNumOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IntervalNum, true
}

// SetIntervalNum sets field value
func (o *InlineResponse2001RateLimits) SetIntervalNum(v int32) {
	o.IntervalNum = v
}

// GetLimit returns the Limit field value
func (o *InlineResponse2001RateLimits) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001RateLimits) GetLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *InlineResponse2001RateLimits) SetLimit(v int32) {
	o.Limit = v
}

func (o InlineResponse2001RateLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rateLimitType"] = o.RateLimitType
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if true {
		toSerialize["intervalNum"] = o.IntervalNum
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001RateLimits struct {
	value *InlineResponse2001RateLimits
	isSet bool
}

func (v NullableInlineResponse2001RateLimits) Get() *InlineResponse2001RateLimits {
	return v.value
}

func (v *NullableInlineResponse2001RateLimits) Set(val *InlineResponse2001RateLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001RateLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001RateLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001RateLimits(val *InlineResponse2001RateLimits) *NullableInlineResponse2001RateLimits {
	return &NullableInlineResponse2001RateLimits{value: val, isSet: true}
}

func (v NullableInlineResponse2001RateLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001RateLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


