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

// InlineResponse20064 struct for InlineResponse20064
type InlineResponse20064 struct {
	CounterParty string `json:"counterParty"`
	Email string `json:"email"`
	// 1 for transfer in, 2 for transfer out
	Type int32 `json:"type"`
	Asset string `json:"asset"`
	Qty string `json:"qty"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType string `json:"toAccountType"`
	Status string `json:"status"`
	TranId int64 `json:"tranId"`
	Time int64 `json:"time"`
}

// NewInlineResponse20064 instantiates a new InlineResponse20064 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20064(counterParty string, email string, type_ int32, asset string, qty string, fromAccountType string, toAccountType string, status string, tranId int64, time int64) *InlineResponse20064 {
	this := InlineResponse20064{}
	this.CounterParty = counterParty
	this.Email = email
	this.Type = type_
	this.Asset = asset
	this.Qty = qty
	this.FromAccountType = fromAccountType
	this.ToAccountType = toAccountType
	this.Status = status
	this.TranId = tranId
	this.Time = time
	return &this
}

// NewInlineResponse20064WithDefaults instantiates a new InlineResponse20064 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20064WithDefaults() *InlineResponse20064 {
	this := InlineResponse20064{}
	return &this
}

// GetCounterParty returns the CounterParty field value
func (o *InlineResponse20064) GetCounterParty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CounterParty
}

// GetCounterPartyOk returns a tuple with the CounterParty field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetCounterPartyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CounterParty, true
}

// SetCounterParty sets field value
func (o *InlineResponse20064) SetCounterParty(v string) {
	o.CounterParty = v
}

// GetEmail returns the Email field value
func (o *InlineResponse20064) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineResponse20064) SetEmail(v string) {
	o.Email = v
}

// GetType returns the Type field value
func (o *InlineResponse20064) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetTypeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20064) SetType(v int32) {
	o.Type = v
}

// GetAsset returns the Asset field value
func (o *InlineResponse20064) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20064) SetAsset(v string) {
	o.Asset = v
}

// GetQty returns the Qty field value
func (o *InlineResponse20064) GetQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetQtyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *InlineResponse20064) SetQty(v string) {
	o.Qty = v
}

// GetFromAccountType returns the FromAccountType field value
func (o *InlineResponse20064) GetFromAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAccountType
}

// GetFromAccountTypeOk returns a tuple with the FromAccountType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetFromAccountTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FromAccountType, true
}

// SetFromAccountType sets field value
func (o *InlineResponse20064) SetFromAccountType(v string) {
	o.FromAccountType = v
}

// GetToAccountType returns the ToAccountType field value
func (o *InlineResponse20064) GetToAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAccountType
}

// GetToAccountTypeOk returns a tuple with the ToAccountType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetToAccountTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ToAccountType, true
}

// SetToAccountType sets field value
func (o *InlineResponse20064) SetToAccountType(v string) {
	o.ToAccountType = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20064) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20064) SetStatus(v string) {
	o.Status = v
}

// GetTranId returns the TranId field value
func (o *InlineResponse20064) GetTranId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetTranIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TranId, true
}

// SetTranId sets field value
func (o *InlineResponse20064) SetTranId(v int64) {
	o.TranId = v
}

// GetTime returns the Time field value
func (o *InlineResponse20064) GetTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *InlineResponse20064) SetTime(v int64) {
	o.Time = v
}

func (o InlineResponse20064) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["counterParty"] = o.CounterParty
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["qty"] = o.Qty
	}
	if true {
		toSerialize["fromAccountType"] = o.FromAccountType
	}
	if true {
		toSerialize["toAccountType"] = o.ToAccountType
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["tranId"] = o.TranId
	}
	if true {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20064 struct {
	value *InlineResponse20064
	isSet bool
}

func (v NullableInlineResponse20064) Get() *InlineResponse20064 {
	return v.value
}

func (v *NullableInlineResponse20064) Set(val *InlineResponse20064) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20064) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20064) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20064(val *InlineResponse20064) *NullableInlineResponse20064 {
	return &NullableInlineResponse20064{value: val, isSet: true}
}

func (v NullableInlineResponse20064) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20064) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


