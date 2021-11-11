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

// ApiV3OpenOrderListOrders struct for ApiV3OpenOrderListOrders
type ApiV3OpenOrderListOrders struct {
	Symbol string `json:"symbol"`
	OrderId int64 `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

// NewApiV3OpenOrderListOrders instantiates a new ApiV3OpenOrderListOrders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV3OpenOrderListOrders(symbol string, orderId int64, clientOrderId string) *ApiV3OpenOrderListOrders {
	this := ApiV3OpenOrderListOrders{}
	this.Symbol = symbol
	this.OrderId = orderId
	this.ClientOrderId = clientOrderId
	return &this
}

// NewApiV3OpenOrderListOrdersWithDefaults instantiates a new ApiV3OpenOrderListOrders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV3OpenOrderListOrdersWithDefaults() *ApiV3OpenOrderListOrders {
	this := ApiV3OpenOrderListOrders{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ApiV3OpenOrderListOrders) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ApiV3OpenOrderListOrders) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ApiV3OpenOrderListOrders) SetSymbol(v string) {
	o.Symbol = v
}

// GetOrderId returns the OrderId field value
func (o *ApiV3OpenOrderListOrders) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *ApiV3OpenOrderListOrders) GetOrderIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *ApiV3OpenOrderListOrders) SetOrderId(v int64) {
	o.OrderId = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *ApiV3OpenOrderListOrders) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *ApiV3OpenOrderListOrders) GetClientOrderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *ApiV3OpenOrderListOrders) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

func (o ApiV3OpenOrderListOrders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	return json.Marshal(toSerialize)
}

type NullableApiV3OpenOrderListOrders struct {
	value *ApiV3OpenOrderListOrders
	isSet bool
}

func (v NullableApiV3OpenOrderListOrders) Get() *ApiV3OpenOrderListOrders {
	return v.value
}

func (v *NullableApiV3OpenOrderListOrders) Set(val *ApiV3OpenOrderListOrders) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV3OpenOrderListOrders) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV3OpenOrderListOrders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV3OpenOrderListOrders(val *ApiV3OpenOrderListOrders) *NullableApiV3OpenOrderListOrders {
	return &NullableApiV3OpenOrderListOrders{value: val, isSet: true}
}

func (v NullableApiV3OpenOrderListOrders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV3OpenOrderListOrders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


