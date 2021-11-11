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

// InlineResponse20088DataWorkerDatas struct for InlineResponse20088DataWorkerDatas
type InlineResponse20088DataWorkerDatas struct {
	WorkerId string `json:"workerId"`
	WorkerName string `json:"workerName"`
	// Status：1 valid, 2 invalid, 3 no longer valid
	Status int64 `json:"status"`
	// Real-time rate
	HashRate int64 `json:"hashRate"`
	// 24H Hashrate
	DayHashRate int64 `json:"dayHashRate"`
	// Real-time Rejection Rate
	RejectRate int64 `json:"rejectRate"`
	// Last submission time
	LastShareTime int64 `json:"lastShareTime"`
}

// NewInlineResponse20088DataWorkerDatas instantiates a new InlineResponse20088DataWorkerDatas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20088DataWorkerDatas(workerId string, workerName string, status int64, hashRate int64, dayHashRate int64, rejectRate int64, lastShareTime int64) *InlineResponse20088DataWorkerDatas {
	this := InlineResponse20088DataWorkerDatas{}
	this.WorkerId = workerId
	this.WorkerName = workerName
	this.Status = status
	this.HashRate = hashRate
	this.DayHashRate = dayHashRate
	this.RejectRate = rejectRate
	this.LastShareTime = lastShareTime
	return &this
}

// NewInlineResponse20088DataWorkerDatasWithDefaults instantiates a new InlineResponse20088DataWorkerDatas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20088DataWorkerDatasWithDefaults() *InlineResponse20088DataWorkerDatas {
	this := InlineResponse20088DataWorkerDatas{}
	return &this
}

// GetWorkerId returns the WorkerId field value
func (o *InlineResponse20088DataWorkerDatas) GetWorkerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20088DataWorkerDatas) GetWorkerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *InlineResponse20088DataWorkerDatas) SetWorkerId(v string) {
	o.WorkerId = v
}

// GetWorkerName returns the WorkerName field value
func (o *InlineResponse20088DataWorkerDatas) GetWorkerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkerName
}

// GetWorkerNameOk returns a tuple with the WorkerName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20088DataWorkerDatas) GetWorkerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WorkerName, true
}

// SetWorkerName sets field value
func (o *InlineResponse20088DataWorkerDatas) SetWorkerName(v string) {
	o.WorkerName = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20088DataWorkerDatas) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20088DataWorkerDatas) GetStatusOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20088DataWorkerDatas) SetStatus(v int64) {
	o.Status = v
}

// GetHashRate returns the HashRate field value
func (o *InlineResponse20088DataWorkerDatas) GetHashRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HashRate
}

// GetHashRateOk returns a tuple with the HashRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20088DataWorkerDatas) GetHashRateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HashRate, true
}

// SetHashRate sets field value
func (o *InlineResponse20088DataWorkerDatas) SetHashRate(v int64) {
	o.HashRate = v
}

// GetDayHashRate returns the DayHashRate field value
func (o *InlineResponse20088DataWorkerDatas) GetDayHashRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DayHashRate
}

// GetDayHashRateOk returns a tuple with the DayHashRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20088DataWorkerDatas) GetDayHashRateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DayHashRate, true
}

// SetDayHashRate sets field value
func (o *InlineResponse20088DataWorkerDatas) SetDayHashRate(v int64) {
	o.DayHashRate = v
}

// GetRejectRate returns the RejectRate field value
func (o *InlineResponse20088DataWorkerDatas) GetRejectRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RejectRate
}

// GetRejectRateOk returns a tuple with the RejectRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20088DataWorkerDatas) GetRejectRateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RejectRate, true
}

// SetRejectRate sets field value
func (o *InlineResponse20088DataWorkerDatas) SetRejectRate(v int64) {
	o.RejectRate = v
}

// GetLastShareTime returns the LastShareTime field value
func (o *InlineResponse20088DataWorkerDatas) GetLastShareTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastShareTime
}

// GetLastShareTimeOk returns a tuple with the LastShareTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20088DataWorkerDatas) GetLastShareTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastShareTime, true
}

// SetLastShareTime sets field value
func (o *InlineResponse20088DataWorkerDatas) SetLastShareTime(v int64) {
	o.LastShareTime = v
}

func (o InlineResponse20088DataWorkerDatas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workerId"] = o.WorkerId
	}
	if true {
		toSerialize["workerName"] = o.WorkerName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["hashRate"] = o.HashRate
	}
	if true {
		toSerialize["dayHashRate"] = o.DayHashRate
	}
	if true {
		toSerialize["rejectRate"] = o.RejectRate
	}
	if true {
		toSerialize["lastShareTime"] = o.LastShareTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20088DataWorkerDatas struct {
	value *InlineResponse20088DataWorkerDatas
	isSet bool
}

func (v NullableInlineResponse20088DataWorkerDatas) Get() *InlineResponse20088DataWorkerDatas {
	return v.value
}

func (v *NullableInlineResponse20088DataWorkerDatas) Set(val *InlineResponse20088DataWorkerDatas) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20088DataWorkerDatas) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20088DataWorkerDatas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20088DataWorkerDatas(val *InlineResponse20088DataWorkerDatas) *NullableInlineResponse20088DataWorkerDatas {
	return &NullableInlineResponse20088DataWorkerDatas{value: val, isSet: true}
}

func (v NullableInlineResponse20088DataWorkerDatas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20088DataWorkerDatas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


