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

// SnapshotMarginSnapshotVos struct for SnapshotMarginSnapshotVos
type SnapshotMarginSnapshotVos struct {
	Data SnapshotMarginData `json:"data"`
	Type string `json:"type"`
	UpdateTime int64 `json:"updateTime"`
}

// NewSnapshotMarginSnapshotVos instantiates a new SnapshotMarginSnapshotVos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotMarginSnapshotVos(data SnapshotMarginData, type_ string, updateTime int64) *SnapshotMarginSnapshotVos {
	this := SnapshotMarginSnapshotVos{}
	this.Data = data
	this.Type = type_
	this.UpdateTime = updateTime
	return &this
}

// NewSnapshotMarginSnapshotVosWithDefaults instantiates a new SnapshotMarginSnapshotVos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotMarginSnapshotVosWithDefaults() *SnapshotMarginSnapshotVos {
	this := SnapshotMarginSnapshotVos{}
	return &this
}

// GetData returns the Data field value
func (o *SnapshotMarginSnapshotVos) GetData() SnapshotMarginData {
	if o == nil {
		var ret SnapshotMarginData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginSnapshotVos) GetDataOk() (*SnapshotMarginData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SnapshotMarginSnapshotVos) SetData(v SnapshotMarginData) {
	o.Data = v
}

// GetType returns the Type field value
func (o *SnapshotMarginSnapshotVos) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginSnapshotVos) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SnapshotMarginSnapshotVos) SetType(v string) {
	o.Type = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *SnapshotMarginSnapshotVos) GetUpdateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginSnapshotVos) GetUpdateTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *SnapshotMarginSnapshotVos) SetUpdateTime(v int64) {
	o.UpdateTime = v
}

func (o SnapshotMarginSnapshotVos) MarshalJSON() ([]byte, error) {
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

type NullableSnapshotMarginSnapshotVos struct {
	value *SnapshotMarginSnapshotVos
	isSet bool
}

func (v NullableSnapshotMarginSnapshotVos) Get() *SnapshotMarginSnapshotVos {
	return v.value
}

func (v *NullableSnapshotMarginSnapshotVos) Set(val *SnapshotMarginSnapshotVos) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotMarginSnapshotVos) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotMarginSnapshotVos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotMarginSnapshotVos(val *SnapshotMarginSnapshotVos) *NullableSnapshotMarginSnapshotVos {
	return &NullableSnapshotMarginSnapshotVos{value: val, isSet: true}
}

func (v NullableSnapshotMarginSnapshotVos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotMarginSnapshotVos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


