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

// InlineResponse20087Data struct for InlineResponse20087Data
type InlineResponse20087Data struct {
	// Mining Account name
	WorkerName string `json:"workerName"`
	// Type of hourly hashrate
	Type string `json:"type"`
	HashrateDatas []InlineResponse20087HashrateDatas `json:"hashrateDatas"`
}

// NewInlineResponse20087Data instantiates a new InlineResponse20087Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20087Data(workerName string, type_ string, hashrateDatas []InlineResponse20087HashrateDatas) *InlineResponse20087Data {
	this := InlineResponse20087Data{}
	this.WorkerName = workerName
	this.Type = type_
	this.HashrateDatas = hashrateDatas
	return &this
}

// NewInlineResponse20087DataWithDefaults instantiates a new InlineResponse20087Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20087DataWithDefaults() *InlineResponse20087Data {
	this := InlineResponse20087Data{}
	return &this
}

// GetWorkerName returns the WorkerName field value
func (o *InlineResponse20087Data) GetWorkerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkerName
}

// GetWorkerNameOk returns a tuple with the WorkerName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Data) GetWorkerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WorkerName, true
}

// SetWorkerName sets field value
func (o *InlineResponse20087Data) SetWorkerName(v string) {
	o.WorkerName = v
}

// GetType returns the Type field value
func (o *InlineResponse20087Data) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Data) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20087Data) SetType(v string) {
	o.Type = v
}

// GetHashrateDatas returns the HashrateDatas field value
func (o *InlineResponse20087Data) GetHashrateDatas() []InlineResponse20087HashrateDatas {
	if o == nil {
		var ret []InlineResponse20087HashrateDatas
		return ret
	}

	return o.HashrateDatas
}

// GetHashrateDatasOk returns a tuple with the HashrateDatas field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20087Data) GetHashrateDatasOk() (*[]InlineResponse20087HashrateDatas, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HashrateDatas, true
}

// SetHashrateDatas sets field value
func (o *InlineResponse20087Data) SetHashrateDatas(v []InlineResponse20087HashrateDatas) {
	o.HashrateDatas = v
}

func (o InlineResponse20087Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workerName"] = o.WorkerName
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["hashrateDatas"] = o.HashrateDatas
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20087Data struct {
	value *InlineResponse20087Data
	isSet bool
}

func (v NullableInlineResponse20087Data) Get() *InlineResponse20087Data {
	return v.value
}

func (v *NullableInlineResponse20087Data) Set(val *InlineResponse20087Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20087Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20087Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20087Data(val *InlineResponse20087Data) *NullableInlineResponse20087Data {
	return &NullableInlineResponse20087Data{value: val, isSet: true}
}

func (v NullableInlineResponse20087Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20087Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


