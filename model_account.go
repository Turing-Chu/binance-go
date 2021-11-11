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

// Account struct for Account
type Account struct {
	MakerCommission int64 `json:"makerCommission"`
	TakerCommission int64 `json:"takerCommission"`
	BuyerCommission int64 `json:"buyerCommission"`
	SellerCommission int64 `json:"sellerCommission"`
	CanTrade bool `json:"canTrade"`
	CanWithdraw bool `json:"canWithdraw"`
	CanDeposit bool `json:"canDeposit"`
	UpdateTime int64 `json:"updateTime"`
	AccountType string `json:"accountType"`
	Balances []AccountBalances `json:"balances"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(makerCommission int64, takerCommission int64, buyerCommission int64, sellerCommission int64, canTrade bool, canWithdraw bool, canDeposit bool, updateTime int64, accountType string, balances []AccountBalances) *Account {
	this := Account{}
	this.MakerCommission = makerCommission
	this.TakerCommission = takerCommission
	this.BuyerCommission = buyerCommission
	this.SellerCommission = sellerCommission
	this.CanTrade = canTrade
	this.CanWithdraw = canWithdraw
	this.CanDeposit = canDeposit
	this.UpdateTime = updateTime
	this.AccountType = accountType
	this.Balances = balances
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetMakerCommission returns the MakerCommission field value
func (o *Account) GetMakerCommission() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MakerCommission
}

// GetMakerCommissionOk returns a tuple with the MakerCommission field value
// and a boolean to check if the value has been set.
func (o *Account) GetMakerCommissionOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MakerCommission, true
}

// SetMakerCommission sets field value
func (o *Account) SetMakerCommission(v int64) {
	o.MakerCommission = v
}

// GetTakerCommission returns the TakerCommission field value
func (o *Account) GetTakerCommission() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value
// and a boolean to check if the value has been set.
func (o *Account) GetTakerCommissionOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TakerCommission, true
}

// SetTakerCommission sets field value
func (o *Account) SetTakerCommission(v int64) {
	o.TakerCommission = v
}

// GetBuyerCommission returns the BuyerCommission field value
func (o *Account) GetBuyerCommission() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BuyerCommission
}

// GetBuyerCommissionOk returns a tuple with the BuyerCommission field value
// and a boolean to check if the value has been set.
func (o *Account) GetBuyerCommissionOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BuyerCommission, true
}

// SetBuyerCommission sets field value
func (o *Account) SetBuyerCommission(v int64) {
	o.BuyerCommission = v
}

// GetSellerCommission returns the SellerCommission field value
func (o *Account) GetSellerCommission() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SellerCommission
}

// GetSellerCommissionOk returns a tuple with the SellerCommission field value
// and a boolean to check if the value has been set.
func (o *Account) GetSellerCommissionOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SellerCommission, true
}

// SetSellerCommission sets field value
func (o *Account) SetSellerCommission(v int64) {
	o.SellerCommission = v
}

// GetCanTrade returns the CanTrade field value
func (o *Account) GetCanTrade() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value
// and a boolean to check if the value has been set.
func (o *Account) GetCanTradeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CanTrade, true
}

// SetCanTrade sets field value
func (o *Account) SetCanTrade(v bool) {
	o.CanTrade = v
}

// GetCanWithdraw returns the CanWithdraw field value
func (o *Account) GetCanWithdraw() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value
// and a boolean to check if the value has been set.
func (o *Account) GetCanWithdrawOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CanWithdraw, true
}

// SetCanWithdraw sets field value
func (o *Account) SetCanWithdraw(v bool) {
	o.CanWithdraw = v
}

// GetCanDeposit returns the CanDeposit field value
func (o *Account) GetCanDeposit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value
// and a boolean to check if the value has been set.
func (o *Account) GetCanDepositOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CanDeposit, true
}

// SetCanDeposit sets field value
func (o *Account) SetCanDeposit(v bool) {
	o.CanDeposit = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *Account) GetUpdateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *Account) GetUpdateTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *Account) SetUpdateTime(v int64) {
	o.UpdateTime = v
}

// GetAccountType returns the AccountType field value
func (o *Account) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *Account) SetAccountType(v string) {
	o.AccountType = v
}

// GetBalances returns the Balances field value
func (o *Account) GetBalances() []AccountBalances {
	if o == nil {
		var ret []AccountBalances
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *Account) GetBalancesOk() (*[]AccountBalances, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *Account) SetBalances(v []AccountBalances) {
	o.Balances = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["makerCommission"] = o.MakerCommission
	}
	if true {
		toSerialize["takerCommission"] = o.TakerCommission
	}
	if true {
		toSerialize["buyerCommission"] = o.BuyerCommission
	}
	if true {
		toSerialize["sellerCommission"] = o.SellerCommission
	}
	if true {
		toSerialize["canTrade"] = o.CanTrade
	}
	if true {
		toSerialize["canWithdraw"] = o.CanWithdraw
	}
	if true {
		toSerialize["canDeposit"] = o.CanDeposit
	}
	if true {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if true {
		toSerialize["accountType"] = o.AccountType
	}
	if true {
		toSerialize["balances"] = o.Balances
	}
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

