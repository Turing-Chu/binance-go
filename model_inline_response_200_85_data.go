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

// InlineResponse20085Data struct for InlineResponse20085Data
type InlineResponse20085Data struct {
	AlgoName string `json:"algoName"`
	AlgoId int64 `json:"algoId"`
	PoolIndex int64 `json:"poolIndex"`
	Unit string `json:"unit"`
}

// NewInlineResponse20085Data instantiates a new InlineResponse20085Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20085Data(algoName string, algoId int64, poolIndex int64, unit string) *InlineResponse20085Data {
	this := InlineResponse20085Data{}
	this.AlgoName = algoName
	this.AlgoId = algoId
	this.PoolIndex = poolIndex
	this.Unit = unit
	return &this
}

// NewInlineResponse20085DataWithDefaults instantiates a new InlineResponse20085Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20085DataWithDefaults() *InlineResponse20085Data {
	this := InlineResponse20085Data{}
	return &this
}

// GetAlgoName returns the AlgoName field value
func (o *InlineResponse20085Data) GetAlgoName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlgoName
}

// GetAlgoNameOk returns a tuple with the AlgoName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Data) GetAlgoNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AlgoName, true
}

// SetAlgoName sets field value
func (o *InlineResponse20085Data) SetAlgoName(v string) {
	o.AlgoName = v
}

// GetAlgoId returns the AlgoId field value
func (o *InlineResponse20085Data) GetAlgoId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Data) GetAlgoIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AlgoId, true
}

// SetAlgoId sets field value
func (o *InlineResponse20085Data) SetAlgoId(v int64) {
	o.AlgoId = v
}

// GetPoolIndex returns the PoolIndex field value
func (o *InlineResponse20085Data) GetPoolIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PoolIndex
}

// GetPoolIndexOk returns a tuple with the PoolIndex field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Data) GetPoolIndexOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PoolIndex, true
}

// SetPoolIndex sets field value
func (o *InlineResponse20085Data) SetPoolIndex(v int64) {
	o.PoolIndex = v
}

// GetUnit returns the Unit field value
func (o *InlineResponse20085Data) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Data) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *InlineResponse20085Data) SetUnit(v string) {
	o.Unit = v
}

func (o InlineResponse20085Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["algoName"] = o.AlgoName
	}
	if true {
		toSerialize["algoId"] = o.AlgoId
	}
	if true {
		toSerialize["poolIndex"] = o.PoolIndex
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20085Data struct {
	value *InlineResponse20085Data
	isSet bool
}

func (v NullableInlineResponse20085Data) Get() *InlineResponse20085Data {
	return v.value
}

func (v *NullableInlineResponse20085Data) Set(val *InlineResponse20085Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20085Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20085Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20085Data(val *InlineResponse20085Data) *NullableInlineResponse20085Data {
	return &NullableInlineResponse20085Data{value: val, isSet: true}
}

func (v NullableInlineResponse20085Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20085Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


