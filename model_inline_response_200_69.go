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

// InlineResponse20069 struct for InlineResponse20069
type InlineResponse20069 struct {
	Coin string `json:"coin"`
	Name string `json:"name"`
	TotalBalance string `json:"totalBalance"`
	AvailableBalance string `json:"availableBalance"`
	InOrder string `json:"inOrder"`
	BtcValue string `json:"btcValue"`
}

// NewInlineResponse20069 instantiates a new InlineResponse20069 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20069(coin string, name string, totalBalance string, availableBalance string, inOrder string, btcValue string) *InlineResponse20069 {
	this := InlineResponse20069{}
	this.Coin = coin
	this.Name = name
	this.TotalBalance = totalBalance
	this.AvailableBalance = availableBalance
	this.InOrder = inOrder
	this.BtcValue = btcValue
	return &this
}

// NewInlineResponse20069WithDefaults instantiates a new InlineResponse20069 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20069WithDefaults() *InlineResponse20069 {
	this := InlineResponse20069{}
	return &this
}

// GetCoin returns the Coin field value
func (o *InlineResponse20069) GetCoin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coin
}

// GetCoinOk returns a tuple with the Coin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20069) GetCoinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Coin, true
}

// SetCoin sets field value
func (o *InlineResponse20069) SetCoin(v string) {
	o.Coin = v
}

// GetName returns the Name field value
func (o *InlineResponse20069) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20069) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse20069) SetName(v string) {
	o.Name = v
}

// GetTotalBalance returns the TotalBalance field value
func (o *InlineResponse20069) GetTotalBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalBalance
}

// GetTotalBalanceOk returns a tuple with the TotalBalance field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20069) GetTotalBalanceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalBalance, true
}

// SetTotalBalance sets field value
func (o *InlineResponse20069) SetTotalBalance(v string) {
	o.TotalBalance = v
}

// GetAvailableBalance returns the AvailableBalance field value
func (o *InlineResponse20069) GetAvailableBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20069) GetAvailableBalanceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableBalance, true
}

// SetAvailableBalance sets field value
func (o *InlineResponse20069) SetAvailableBalance(v string) {
	o.AvailableBalance = v
}

// GetInOrder returns the InOrder field value
func (o *InlineResponse20069) GetInOrder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InOrder
}

// GetInOrderOk returns a tuple with the InOrder field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20069) GetInOrderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InOrder, true
}

// SetInOrder sets field value
func (o *InlineResponse20069) SetInOrder(v string) {
	o.InOrder = v
}

// GetBtcValue returns the BtcValue field value
func (o *InlineResponse20069) GetBtcValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtcValue
}

// GetBtcValueOk returns a tuple with the BtcValue field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20069) GetBtcValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BtcValue, true
}

// SetBtcValue sets field value
func (o *InlineResponse20069) SetBtcValue(v string) {
	o.BtcValue = v
}

func (o InlineResponse20069) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["coin"] = o.Coin
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["totalBalance"] = o.TotalBalance
	}
	if true {
		toSerialize["availableBalance"] = o.AvailableBalance
	}
	if true {
		toSerialize["inOrder"] = o.InOrder
	}
	if true {
		toSerialize["btcValue"] = o.BtcValue
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20069 struct {
	value *InlineResponse20069
	isSet bool
}

func (v NullableInlineResponse20069) Get() *InlineResponse20069 {
	return v.value
}

func (v *NullableInlineResponse20069) Set(val *InlineResponse20069) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20069) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20069) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20069(val *InlineResponse20069) *NullableInlineResponse20069 {
	return &NullableInlineResponse20069{value: val, isSet: true}
}

func (v NullableInlineResponse20069) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20069) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

