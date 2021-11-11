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

// SnapshotFuturesSnapshotVos struct for SnapshotFuturesSnapshotVos
type SnapshotFuturesSnapshotVos struct {
	Data SnapshotFuturesData `json:"data"`
	Type string `json:"type"`
	UpdateTime int64 `json:"updateTime"`
}

// NewSnapshotFuturesSnapshotVos instantiates a new SnapshotFuturesSnapshotVos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotFuturesSnapshotVos(data SnapshotFuturesData, type_ string, updateTime int64) *SnapshotFuturesSnapshotVos {
	this := SnapshotFuturesSnapshotVos{}
	this.Data = data
	this.Type = type_
	this.UpdateTime = updateTime
	return &this
}

// NewSnapshotFuturesSnapshotVosWithDefaults instantiates a new SnapshotFuturesSnapshotVos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotFuturesSnapshotVosWithDefaults() *SnapshotFuturesSnapshotVos {
	this := SnapshotFuturesSnapshotVos{}
	return &this
}

// GetData returns the Data field value
func (o *SnapshotFuturesSnapshotVos) GetData() SnapshotFuturesData {
	if o == nil {
		var ret SnapshotFuturesData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SnapshotFuturesSnapshotVos) GetDataOk() (*SnapshotFuturesData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SnapshotFuturesSnapshotVos) SetData(v SnapshotFuturesData) {
	o.Data = v
}

// GetType returns the Type field value
func (o *SnapshotFuturesSnapshotVos) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SnapshotFuturesSnapshotVos) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SnapshotFuturesSnapshotVos) SetType(v string) {
	o.Type = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *SnapshotFuturesSnapshotVos) GetUpdateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *SnapshotFuturesSnapshotVos) GetUpdateTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *SnapshotFuturesSnapshotVos) SetUpdateTime(v int64) {
	o.UpdateTime = v
}

func (o SnapshotFuturesSnapshotVos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotFuturesSnapshotVos struct {
	value *SnapshotFuturesSnapshotVos
	isSet bool
}

func (v NullableSnapshotFuturesSnapshotVos) Get() *SnapshotFuturesSnapshotVos {
	return v.value
}

func (v *NullableSnapshotFuturesSnapshotVos) Set(val *SnapshotFuturesSnapshotVos) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotFuturesSnapshotVos) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotFuturesSnapshotVos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotFuturesSnapshotVos(val *SnapshotFuturesSnapshotVos) *NullableSnapshotFuturesSnapshotVos {
	return &NullableSnapshotFuturesSnapshotVos{value: val, isSet: true}
}

func (v NullableSnapshotFuturesSnapshotVos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotFuturesSnapshotVos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


