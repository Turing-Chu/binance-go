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

// Order struct for Order
type Order struct {
	Symbol string `json:"symbol"`
	OrigClientOrderId string `json:"origClientOrderId"`
	OrderId int64 `json:"orderId"`
	// Unless OCO, value will be -1
	OrderListId int64 `json:"orderListId"`
	ClientOrderId string `json:"clientOrderId"`
	Price string `json:"price"`
	OrigQty string `json:"origQty"`
	ExecutedQty string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status string `json:"status"`
	TimeInForce string `json:"timeInForce"`
	Type string `json:"type"`
	Side string `json:"side"`
}

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder(symbol string, origClientOrderId string, orderId int64, orderListId int64, clientOrderId string, price string, origQty string, executedQty string, cummulativeQuoteQty string, status string, timeInForce string, type_ string, side string) *Order {
	this := Order{}
	this.Symbol = symbol
	this.OrigClientOrderId = origClientOrderId
	this.OrderId = orderId
	this.OrderListId = orderListId
	this.ClientOrderId = clientOrderId
	this.Price = price
	this.OrigQty = origQty
	this.ExecutedQty = executedQty
	this.CummulativeQuoteQty = cummulativeQuoteQty
	this.Status = status
	this.TimeInForce = timeInForce
	this.Type = type_
	this.Side = side
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *Order) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Order) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Order) SetSymbol(v string) {
	o.Symbol = v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value
func (o *Order) GetOrigClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value
// and a boolean to check if the value has been set.
func (o *Order) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrigClientOrderId, true
}

// SetOrigClientOrderId sets field value
func (o *Order) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = v
}

// GetOrderId returns the OrderId field value
func (o *Order) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *Order) GetOrderIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *Order) SetOrderId(v int64) {
	o.OrderId = v
}

// GetOrderListId returns the OrderListId field value
func (o *Order) GetOrderListId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value
// and a boolean to check if the value has been set.
func (o *Order) GetOrderListIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderListId, true
}

// SetOrderListId sets field value
func (o *Order) SetOrderListId(v int64) {
	o.OrderListId = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *Order) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *Order) GetClientOrderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *Order) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

// GetPrice returns the Price field value
func (o *Order) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *Order) GetPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *Order) SetPrice(v string) {
	o.Price = v
}

// GetOrigQty returns the OrigQty field value
func (o *Order) GetOrigQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value
// and a boolean to check if the value has been set.
func (o *Order) GetOrigQtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrigQty, true
}

// SetOrigQty sets field value
func (o *Order) SetOrigQty(v string) {
	o.OrigQty = v
}

// GetExecutedQty returns the ExecutedQty field value
func (o *Order) GetExecutedQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value
// and a boolean to check if the value has been set.
func (o *Order) GetExecutedQtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExecutedQty, true
}

// SetExecutedQty sets field value
func (o *Order) SetExecutedQty(v string) {
	o.ExecutedQty = v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value
func (o *Order) GetCummulativeQuoteQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value
// and a boolean to check if the value has been set.
func (o *Order) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CummulativeQuoteQty, true
}

// SetCummulativeQuoteQty sets field value
func (o *Order) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = v
}

// GetStatus returns the Status field value
func (o *Order) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Order) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Order) SetStatus(v string) {
	o.Status = v
}

// GetTimeInForce returns the TimeInForce field value
func (o *Order) GetTimeInForce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value
// and a boolean to check if the value has been set.
func (o *Order) GetTimeInForceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeInForce, true
}

// SetTimeInForce sets field value
func (o *Order) SetTimeInForce(v string) {
	o.TimeInForce = v
}

// GetType returns the Type field value
func (o *Order) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Order) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Order) SetType(v string) {
	o.Type = v
}

// GetSide returns the Side field value
func (o *Order) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *Order) GetSideOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *Order) SetSide(v string) {
	o.Side = v
}

func (o Order) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["origClientOrderId"] = o.OrigClientOrderId
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["orderListId"] = o.OrderListId
	}
	if true {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["origQty"] = o.OrigQty
	}
	if true {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if true {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["side"] = o.Side
	}
	return json.Marshal(toSerialize)
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


