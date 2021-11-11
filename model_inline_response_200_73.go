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

// InlineResponse20073 struct for InlineResponse20073
type InlineResponse20073 struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data []InlineResponse20073Data `json:"data"`
	Total int32 `json:"total"`
	Success bool `json:"success"`
}

// NewInlineResponse20073 instantiates a new InlineResponse20073 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20073(code string, message string, data []InlineResponse20073Data, total int32, success bool) *InlineResponse20073 {
	this := InlineResponse20073{}
	this.Code = code
	this.Message = message
	this.Data = data
	this.Total = total
	this.Success = success
	return &this
}

// NewInlineResponse20073WithDefaults instantiates a new InlineResponse20073 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20073WithDefaults() *InlineResponse20073 {
	this := InlineResponse20073{}
	return &this
}

// GetCode returns the Code field value
func (o *InlineResponse20073) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InlineResponse20073) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *InlineResponse20073) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InlineResponse20073) SetMessage(v string) {
	o.Message = v
}

// GetData returns the Data field value
func (o *InlineResponse20073) GetData() []InlineResponse20073Data {
	if o == nil {
		var ret []InlineResponse20073Data
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetDataOk() (*[]InlineResponse20073Data, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InlineResponse20073) SetData(v []InlineResponse20073Data) {
	o.Data = v
}

// GetTotal returns the Total field value
func (o *InlineResponse20073) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *InlineResponse20073) SetTotal(v int32) {
	o.Total = v
}

// GetSuccess returns the Success field value
func (o *InlineResponse20073) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *InlineResponse20073) SetSuccess(v bool) {
	o.Success = v
}

func (o InlineResponse20073) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20073 struct {
	value *InlineResponse20073
	isSet bool
}

func (v NullableInlineResponse20073) Get() *InlineResponse20073 {
	return v.value
}

func (v *NullableInlineResponse20073) Set(val *InlineResponse20073) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20073) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20073) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20073(val *InlineResponse20073) *NullableInlineResponse20073 {
	return &NullableInlineResponse20073{value: val, isSet: true}
}

func (v NullableInlineResponse20073) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20073) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


