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

// MarginOrderResponseFull struct for MarginOrderResponseFull
type MarginOrderResponseFull struct {
	Symbol string `json:"symbol"`
	OrderId int64 `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime int64 `json:"transactTime"`
	Price string `json:"price"`
	OrigQty string `json:"origQty"`
	ExecutedQty string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status string `json:"status"`
	TimeInForce string `json:"timeInForce"`
	Type string `json:"type"`
	Side string `json:"side"`
	// will not return if no margin trade happens
	MarginBuyBorrowAmount float64 `json:"marginBuyBorrowAmount"`
	// will not return if no margin trade happens
	MarginBuyBorrowAsset string `json:"marginBuyBorrowAsset"`
	IsIsolated bool `json:"isIsolated"`
	Fills []OrderResponseFullFills `json:"fills"`
}

// NewMarginOrderResponseFull instantiates a new MarginOrderResponseFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginOrderResponseFull(symbol string, orderId int64, clientOrderId string, transactTime int64, price string, origQty string, executedQty string, cummulativeQuoteQty string, status string, timeInForce string, type_ string, side string, marginBuyBorrowAmount float64, marginBuyBorrowAsset string, isIsolated bool, fills []OrderResponseFullFills) *MarginOrderResponseFull {
	this := MarginOrderResponseFull{}
	this.Symbol = symbol
	this.OrderId = orderId
	this.ClientOrderId = clientOrderId
	this.TransactTime = transactTime
	this.Price = price
	this.OrigQty = origQty
	this.ExecutedQty = executedQty
	this.CummulativeQuoteQty = cummulativeQuoteQty
	this.Status = status
	this.TimeInForce = timeInForce
	this.Type = type_
	this.Side = side
	this.MarginBuyBorrowAmount = marginBuyBorrowAmount
	this.MarginBuyBorrowAsset = marginBuyBorrowAsset
	this.IsIsolated = isIsolated
	this.Fills = fills
	return &this
}

// NewMarginOrderResponseFullWithDefaults instantiates a new MarginOrderResponseFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginOrderResponseFullWithDefaults() *MarginOrderResponseFull {
	this := MarginOrderResponseFull{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *MarginOrderResponseFull) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *MarginOrderResponseFull) SetSymbol(v string) {
	o.Symbol = v
}

// GetOrderId returns the OrderId field value
func (o *MarginOrderResponseFull) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetOrderIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *MarginOrderResponseFull) SetOrderId(v int64) {
	o.OrderId = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *MarginOrderResponseFull) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetClientOrderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *MarginOrderResponseFull) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

// GetTransactTime returns the TransactTime field value
func (o *MarginOrderResponseFull) GetTransactTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetTransactTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactTime, true
}

// SetTransactTime sets field value
func (o *MarginOrderResponseFull) SetTransactTime(v int64) {
	o.TransactTime = v
}

// GetPrice returns the Price field value
func (o *MarginOrderResponseFull) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *MarginOrderResponseFull) SetPrice(v string) {
	o.Price = v
}

// GetOrigQty returns the OrigQty field value
func (o *MarginOrderResponseFull) GetOrigQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetOrigQtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrigQty, true
}

// SetOrigQty sets field value
func (o *MarginOrderResponseFull) SetOrigQty(v string) {
	o.OrigQty = v
}

// GetExecutedQty returns the ExecutedQty field value
func (o *MarginOrderResponseFull) GetExecutedQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetExecutedQtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExecutedQty, true
}

// SetExecutedQty sets field value
func (o *MarginOrderResponseFull) SetExecutedQty(v string) {
	o.ExecutedQty = v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value
func (o *MarginOrderResponseFull) GetCummulativeQuoteQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CummulativeQuoteQty, true
}

// SetCummulativeQuoteQty sets field value
func (o *MarginOrderResponseFull) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = v
}

// GetStatus returns the Status field value
func (o *MarginOrderResponseFull) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MarginOrderResponseFull) SetStatus(v string) {
	o.Status = v
}

// GetTimeInForce returns the TimeInForce field value
func (o *MarginOrderResponseFull) GetTimeInForce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetTimeInForceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeInForce, true
}

// SetTimeInForce sets field value
func (o *MarginOrderResponseFull) SetTimeInForce(v string) {
	o.TimeInForce = v
}

// GetType returns the Type field value
func (o *MarginOrderResponseFull) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MarginOrderResponseFull) SetType(v string) {
	o.Type = v
}

// GetSide returns the Side field value
func (o *MarginOrderResponseFull) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetSideOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *MarginOrderResponseFull) SetSide(v string) {
	o.Side = v
}

// GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field value
func (o *MarginOrderResponseFull) GetMarginBuyBorrowAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MarginBuyBorrowAmount
}

// GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetMarginBuyBorrowAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarginBuyBorrowAmount, true
}

// SetMarginBuyBorrowAmount sets field value
func (o *MarginOrderResponseFull) SetMarginBuyBorrowAmount(v float64) {
	o.MarginBuyBorrowAmount = v
}

// GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field value
func (o *MarginOrderResponseFull) GetMarginBuyBorrowAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarginBuyBorrowAsset
}

// GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetMarginBuyBorrowAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MarginBuyBorrowAsset, true
}

// SetMarginBuyBorrowAsset sets field value
func (o *MarginOrderResponseFull) SetMarginBuyBorrowAsset(v string) {
	o.MarginBuyBorrowAsset = v
}

// GetIsIsolated returns the IsIsolated field value
func (o *MarginOrderResponseFull) GetIsIsolated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetIsIsolatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsIsolated, true
}

// SetIsIsolated sets field value
func (o *MarginOrderResponseFull) SetIsIsolated(v bool) {
	o.IsIsolated = v
}

// GetFills returns the Fills field value
func (o *MarginOrderResponseFull) GetFills() []OrderResponseFullFills {
	if o == nil {
		var ret []OrderResponseFullFills
		return ret
	}

	return o.Fills
}

// GetFillsOk returns a tuple with the Fills field value
// and a boolean to check if the value has been set.
func (o *MarginOrderResponseFull) GetFillsOk() (*[]OrderResponseFullFills, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fills, true
}

// SetFills sets field value
func (o *MarginOrderResponseFull) SetFills(v []OrderResponseFullFills) {
	o.Fills = v
}

func (o MarginOrderResponseFull) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["transactTime"] = o.TransactTime
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
	if true {
		toSerialize["marginBuyBorrowAmount"] = o.MarginBuyBorrowAmount
	}
	if true {
		toSerialize["marginBuyBorrowAsset"] = o.MarginBuyBorrowAsset
	}
	if true {
		toSerialize["isIsolated"] = o.IsIsolated
	}
	if true {
		toSerialize["fills"] = o.Fills
	}
	return json.Marshal(toSerialize)
}

type NullableMarginOrderResponseFull struct {
	value *MarginOrderResponseFull
	isSet bool
}

func (v NullableMarginOrderResponseFull) Get() *MarginOrderResponseFull {
	return v.value
}

func (v *NullableMarginOrderResponseFull) Set(val *MarginOrderResponseFull) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginOrderResponseFull) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginOrderResponseFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginOrderResponseFull(val *MarginOrderResponseFull) *NullableMarginOrderResponseFull {
	return &NullableMarginOrderResponseFull{value: val, isSet: true}
}

func (v NullableMarginOrderResponseFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginOrderResponseFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

