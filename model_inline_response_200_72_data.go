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

// InlineResponse20072Data struct for InlineResponse20072Data
type InlineResponse20072Data struct {
	OrderNo string `json:"orderNo"`
	FiatCurrency string `json:"fiatCurrency"`
	IndicatedAmount string `json:"indicatedAmount"`
	Amount string `json:"amount"`
	TotalFee string `json:"totalFee"`
	Method string `json:"method"`
	// Processing, Failed, Successful, Finished, Refunding, Refunded, Refund Failed, Order Partial credit Stopped
	Status string `json:"status"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
}

// NewInlineResponse20072Data instantiates a new InlineResponse20072Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20072Data(orderNo string, fiatCurrency string, indicatedAmount string, amount string, totalFee string, method string, status string, createTime int64, updateTime int64) *InlineResponse20072Data {
	this := InlineResponse20072Data{}
	this.OrderNo = orderNo
	this.FiatCurrency = fiatCurrency
	this.IndicatedAmount = indicatedAmount
	this.Amount = amount
	this.TotalFee = totalFee
	this.Method = method
	this.Status = status
	this.CreateTime = createTime
	this.UpdateTime = updateTime
	return &this
}

// NewInlineResponse20072DataWithDefaults instantiates a new InlineResponse20072Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20072DataWithDefaults() *InlineResponse20072Data {
	this := InlineResponse20072Data{}
	return &this
}

// GetOrderNo returns the OrderNo field value
func (o *InlineResponse20072Data) GetOrderNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetOrderNoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderNo, true
}

// SetOrderNo sets field value
func (o *InlineResponse20072Data) SetOrderNo(v string) {
	o.OrderNo = v
}

// GetFiatCurrency returns the FiatCurrency field value
func (o *InlineResponse20072Data) GetFiatCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FiatCurrency
}

// GetFiatCurrencyOk returns a tuple with the FiatCurrency field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetFiatCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FiatCurrency, true
}

// SetFiatCurrency sets field value
func (o *InlineResponse20072Data) SetFiatCurrency(v string) {
	o.FiatCurrency = v
}

// GetIndicatedAmount returns the IndicatedAmount field value
func (o *InlineResponse20072Data) GetIndicatedAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndicatedAmount
}

// GetIndicatedAmountOk returns a tuple with the IndicatedAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetIndicatedAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IndicatedAmount, true
}

// SetIndicatedAmount sets field value
func (o *InlineResponse20072Data) SetIndicatedAmount(v string) {
	o.IndicatedAmount = v
}

// GetAmount returns the Amount field value
func (o *InlineResponse20072Data) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse20072Data) SetAmount(v string) {
	o.Amount = v
}

// GetTotalFee returns the TotalFee field value
func (o *InlineResponse20072Data) GetTotalFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalFee
}

// GetTotalFeeOk returns a tuple with the TotalFee field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetTotalFeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalFee, true
}

// SetTotalFee sets field value
func (o *InlineResponse20072Data) SetTotalFee(v string) {
	o.TotalFee = v
}

// GetMethod returns the Method field value
func (o *InlineResponse20072Data) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *InlineResponse20072Data) SetMethod(v string) {
	o.Method = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20072Data) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20072Data) SetStatus(v string) {
	o.Status = v
}

// GetCreateTime returns the CreateTime field value
func (o *InlineResponse20072Data) GetCreateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetCreateTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *InlineResponse20072Data) SetCreateTime(v int64) {
	o.CreateTime = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *InlineResponse20072Data) GetUpdateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20072Data) GetUpdateTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *InlineResponse20072Data) SetUpdateTime(v int64) {
	o.UpdateTime = v
}

func (o InlineResponse20072Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orderNo"] = o.OrderNo
	}
	if true {
		toSerialize["fiatCurrency"] = o.FiatCurrency
	}
	if true {
		toSerialize["indicatedAmount"] = o.IndicatedAmount
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["totalFee"] = o.TotalFee
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["createTime"] = o.CreateTime
	}
	if true {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20072Data struct {
	value *InlineResponse20072Data
	isSet bool
}

func (v NullableInlineResponse20072Data) Get() *InlineResponse20072Data {
	return v.value
}

func (v *NullableInlineResponse20072Data) Set(val *InlineResponse20072Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20072Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20072Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20072Data(val *InlineResponse20072Data) *NullableInlineResponse20072Data {
	return &NullableInlineResponse20072Data{value: val, isSet: true}
}

func (v NullableInlineResponse20072Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20072Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

