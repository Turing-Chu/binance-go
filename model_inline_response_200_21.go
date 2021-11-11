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

// InlineResponse20021 struct for InlineResponse20021
type InlineResponse20021 struct {
	OrderListId int64 `json:"orderListId"`
	ContingencyType string `json:"contingencyType"`
	ListStatusType string `json:"listStatusType"`
	ListOrderStatus string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime int64 `json:"transactionTime"`
	Symbol string `json:"symbol"`
	IsIsolated bool `json:"isIsolated"`
	Orders []InlineResponse2005Orders `json:"orders"`
}

// NewInlineResponse20021 instantiates a new InlineResponse20021 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20021(orderListId int64, contingencyType string, listStatusType string, listOrderStatus string, listClientOrderId string, transactionTime int64, symbol string, isIsolated bool, orders []InlineResponse2005Orders) *InlineResponse20021 {
	this := InlineResponse20021{}
	this.OrderListId = orderListId
	this.ContingencyType = contingencyType
	this.ListStatusType = listStatusType
	this.ListOrderStatus = listOrderStatus
	this.ListClientOrderId = listClientOrderId
	this.TransactionTime = transactionTime
	this.Symbol = symbol
	this.IsIsolated = isIsolated
	this.Orders = orders
	return &this
}

// NewInlineResponse20021WithDefaults instantiates a new InlineResponse20021 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20021WithDefaults() *InlineResponse20021 {
	this := InlineResponse20021{}
	return &this
}

// GetOrderListId returns the OrderListId field value
func (o *InlineResponse20021) GetOrderListId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetOrderListIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderListId, true
}

// SetOrderListId sets field value
func (o *InlineResponse20021) SetOrderListId(v int64) {
	o.OrderListId = v
}

// GetContingencyType returns the ContingencyType field value
func (o *InlineResponse20021) GetContingencyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetContingencyTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContingencyType, true
}

// SetContingencyType sets field value
func (o *InlineResponse20021) SetContingencyType(v string) {
	o.ContingencyType = v
}

// GetListStatusType returns the ListStatusType field value
func (o *InlineResponse20021) GetListStatusType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetListStatusTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListStatusType, true
}

// SetListStatusType sets field value
func (o *InlineResponse20021) SetListStatusType(v string) {
	o.ListStatusType = v
}

// GetListOrderStatus returns the ListOrderStatus field value
func (o *InlineResponse20021) GetListOrderStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetListOrderStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListOrderStatus, true
}

// SetListOrderStatus sets field value
func (o *InlineResponse20021) SetListOrderStatus(v string) {
	o.ListOrderStatus = v
}

// GetListClientOrderId returns the ListClientOrderId field value
func (o *InlineResponse20021) GetListClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetListClientOrderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListClientOrderId, true
}

// SetListClientOrderId sets field value
func (o *InlineResponse20021) SetListClientOrderId(v string) {
	o.ListClientOrderId = v
}

// GetTransactionTime returns the TransactionTime field value
func (o *InlineResponse20021) GetTransactionTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetTransactionTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *InlineResponse20021) SetTransactionTime(v int64) {
	o.TransactionTime = v
}

// GetSymbol returns the Symbol field value
func (o *InlineResponse20021) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *InlineResponse20021) SetSymbol(v string) {
	o.Symbol = v
}

// GetIsIsolated returns the IsIsolated field value
func (o *InlineResponse20021) GetIsIsolated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetIsIsolatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsIsolated, true
}

// SetIsIsolated sets field value
func (o *InlineResponse20021) SetIsIsolated(v bool) {
	o.IsIsolated = v
}

// GetOrders returns the Orders field value
func (o *InlineResponse20021) GetOrders() []InlineResponse2005Orders {
	if o == nil {
		var ret []InlineResponse2005Orders
		return ret
	}

	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetOrdersOk() (*[]InlineResponse2005Orders, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Orders, true
}

// SetOrders sets field value
func (o *InlineResponse20021) SetOrders(v []InlineResponse2005Orders) {
	o.Orders = v
}

func (o InlineResponse20021) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orderListId"] = o.OrderListId
	}
	if true {
		toSerialize["contingencyType"] = o.ContingencyType
	}
	if true {
		toSerialize["listStatusType"] = o.ListStatusType
	}
	if true {
		toSerialize["listOrderStatus"] = o.ListOrderStatus
	}
	if true {
		toSerialize["listClientOrderId"] = o.ListClientOrderId
	}
	if true {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["isIsolated"] = o.IsIsolated
	}
	if true {
		toSerialize["orders"] = o.Orders
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20021 struct {
	value *InlineResponse20021
	isSet bool
}

func (v NullableInlineResponse20021) Get() *InlineResponse20021 {
	return v.value
}

func (v *NullableInlineResponse20021) Set(val *InlineResponse20021) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20021) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20021) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20021(val *InlineResponse20021) *NullableInlineResponse20021 {
	return &NullableInlineResponse20021{value: val, isSet: true}
}

func (v NullableInlineResponse20021) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20021) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


