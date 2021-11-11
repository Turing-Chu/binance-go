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

// InlineResponse2007 struct for InlineResponse2007
type InlineResponse2007 struct {
	OrderListId int64 `json:"orderListId"`
	ContingencyType string `json:"contingencyType"`
	ListStatusType string `json:"listStatusType"`
	ListOrderStatus string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime int64 `json:"transactionTime"`
	Symbol string `json:"symbol"`
	Orders []ApiV3OpenOrderListOrders `json:"orders"`
}

// NewInlineResponse2007 instantiates a new InlineResponse2007 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007(orderListId int64, contingencyType string, listStatusType string, listOrderStatus string, listClientOrderId string, transactionTime int64, symbol string, orders []ApiV3OpenOrderListOrders) *InlineResponse2007 {
	this := InlineResponse2007{}
	this.OrderListId = orderListId
	this.ContingencyType = contingencyType
	this.ListStatusType = listStatusType
	this.ListOrderStatus = listOrderStatus
	this.ListClientOrderId = listClientOrderId
	this.TransactionTime = transactionTime
	this.Symbol = symbol
	this.Orders = orders
	return &this
}

// NewInlineResponse2007WithDefaults instantiates a new InlineResponse2007 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007WithDefaults() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// GetOrderListId returns the OrderListId field value
func (o *InlineResponse2007) GetOrderListId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetOrderListIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderListId, true
}

// SetOrderListId sets field value
func (o *InlineResponse2007) SetOrderListId(v int64) {
	o.OrderListId = v
}

// GetContingencyType returns the ContingencyType field value
func (o *InlineResponse2007) GetContingencyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetContingencyTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContingencyType, true
}

// SetContingencyType sets field value
func (o *InlineResponse2007) SetContingencyType(v string) {
	o.ContingencyType = v
}

// GetListStatusType returns the ListStatusType field value
func (o *InlineResponse2007) GetListStatusType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetListStatusTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListStatusType, true
}

// SetListStatusType sets field value
func (o *InlineResponse2007) SetListStatusType(v string) {
	o.ListStatusType = v
}

// GetListOrderStatus returns the ListOrderStatus field value
func (o *InlineResponse2007) GetListOrderStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetListOrderStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListOrderStatus, true
}

// SetListOrderStatus sets field value
func (o *InlineResponse2007) SetListOrderStatus(v string) {
	o.ListOrderStatus = v
}

// GetListClientOrderId returns the ListClientOrderId field value
func (o *InlineResponse2007) GetListClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetListClientOrderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListClientOrderId, true
}

// SetListClientOrderId sets field value
func (o *InlineResponse2007) SetListClientOrderId(v string) {
	o.ListClientOrderId = v
}

// GetTransactionTime returns the TransactionTime field value
func (o *InlineResponse2007) GetTransactionTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetTransactionTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *InlineResponse2007) SetTransactionTime(v int64) {
	o.TransactionTime = v
}

// GetSymbol returns the Symbol field value
func (o *InlineResponse2007) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *InlineResponse2007) SetSymbol(v string) {
	o.Symbol = v
}

// GetOrders returns the Orders field value
func (o *InlineResponse2007) GetOrders() []ApiV3OpenOrderListOrders {
	if o == nil {
		var ret []ApiV3OpenOrderListOrders
		return ret
	}

	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetOrdersOk() (*[]ApiV3OpenOrderListOrders, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Orders, true
}

// SetOrders sets field value
func (o *InlineResponse2007) SetOrders(v []ApiV3OpenOrderListOrders) {
	o.Orders = v
}

func (o InlineResponse2007) MarshalJSON() ([]byte, error) {
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
		toSerialize["orders"] = o.Orders
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2007 struct {
	value *InlineResponse2007
	isSet bool
}

func (v NullableInlineResponse2007) Get() *InlineResponse2007 {
	return v.value
}

func (v *NullableInlineResponse2007) Set(val *InlineResponse2007) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007(val *InlineResponse2007) *NullableInlineResponse2007 {
	return &NullableInlineResponse2007{value: val, isSet: true}
}

func (v NullableInlineResponse2007) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


