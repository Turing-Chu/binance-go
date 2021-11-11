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

// InlineResponse20079 struct for InlineResponse20079
type InlineResponse20079 struct {
	Asset string `json:"asset"`
	DisplayPriority int64 `json:"displayPriority"`
	Duration int64 `json:"duration"`
	InterestPerLot string `json:"interestPerLot"`
	InterestRate string `json:"interestRate"`
	LotSize string `json:"lotSize"`
	LotsLowLimit int64 `json:"lotsLowLimit"`
	LotsPurchased int64 `json:"lotsPurchased"`
	LotsUpLimit int64 `json:"lotsUpLimit"`
	MaxLotsPerUser int64 `json:"maxLotsPerUser"`
	NeedKyc bool `json:"needKyc"`
	ProjectId string `json:"projectId"`
	ProjectName string `json:"projectName"`
	Status string `json:"status"`
	Type string `json:"type"`
	WithAreaLimitation bool `json:"withAreaLimitation"`
}

// NewInlineResponse20079 instantiates a new InlineResponse20079 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079(asset string, displayPriority int64, duration int64, interestPerLot string, interestRate string, lotSize string, lotsLowLimit int64, lotsPurchased int64, lotsUpLimit int64, maxLotsPerUser int64, needKyc bool, projectId string, projectName string, status string, type_ string, withAreaLimitation bool) *InlineResponse20079 {
	this := InlineResponse20079{}
	this.Asset = asset
	this.DisplayPriority = displayPriority
	this.Duration = duration
	this.InterestPerLot = interestPerLot
	this.InterestRate = interestRate
	this.LotSize = lotSize
	this.LotsLowLimit = lotsLowLimit
	this.LotsPurchased = lotsPurchased
	this.LotsUpLimit = lotsUpLimit
	this.MaxLotsPerUser = maxLotsPerUser
	this.NeedKyc = needKyc
	this.ProjectId = projectId
	this.ProjectName = projectName
	this.Status = status
	this.Type = type_
	this.WithAreaLimitation = withAreaLimitation
	return &this
}

// NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079WithDefaults() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// GetAsset returns the Asset field value
func (o *InlineResponse20079) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20079) SetAsset(v string) {
	o.Asset = v
}

// GetDisplayPriority returns the DisplayPriority field value
func (o *InlineResponse20079) GetDisplayPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DisplayPriority
}

// GetDisplayPriorityOk returns a tuple with the DisplayPriority field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetDisplayPriorityOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayPriority, true
}

// SetDisplayPriority sets field value
func (o *InlineResponse20079) SetDisplayPriority(v int64) {
	o.DisplayPriority = v
}

// GetDuration returns the Duration field value
func (o *InlineResponse20079) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetDurationOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *InlineResponse20079) SetDuration(v int64) {
	o.Duration = v
}

// GetInterestPerLot returns the InterestPerLot field value
func (o *InlineResponse20079) GetInterestPerLot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterestPerLot
}

// GetInterestPerLotOk returns a tuple with the InterestPerLot field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetInterestPerLotOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterestPerLot, true
}

// SetInterestPerLot sets field value
func (o *InlineResponse20079) SetInterestPerLot(v string) {
	o.InterestPerLot = v
}

// GetInterestRate returns the InterestRate field value
func (o *InlineResponse20079) GetInterestRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetInterestRateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterestRate, true
}

// SetInterestRate sets field value
func (o *InlineResponse20079) SetInterestRate(v string) {
	o.InterestRate = v
}

// GetLotSize returns the LotSize field value
func (o *InlineResponse20079) GetLotSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LotSize
}

// GetLotSizeOk returns a tuple with the LotSize field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetLotSizeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LotSize, true
}

// SetLotSize sets field value
func (o *InlineResponse20079) SetLotSize(v string) {
	o.LotSize = v
}

// GetLotsLowLimit returns the LotsLowLimit field value
func (o *InlineResponse20079) GetLotsLowLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LotsLowLimit
}

// GetLotsLowLimitOk returns a tuple with the LotsLowLimit field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetLotsLowLimitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LotsLowLimit, true
}

// SetLotsLowLimit sets field value
func (o *InlineResponse20079) SetLotsLowLimit(v int64) {
	o.LotsLowLimit = v
}

// GetLotsPurchased returns the LotsPurchased field value
func (o *InlineResponse20079) GetLotsPurchased() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LotsPurchased
}

// GetLotsPurchasedOk returns a tuple with the LotsPurchased field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetLotsPurchasedOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LotsPurchased, true
}

// SetLotsPurchased sets field value
func (o *InlineResponse20079) SetLotsPurchased(v int64) {
	o.LotsPurchased = v
}

// GetLotsUpLimit returns the LotsUpLimit field value
func (o *InlineResponse20079) GetLotsUpLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LotsUpLimit
}

// GetLotsUpLimitOk returns a tuple with the LotsUpLimit field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetLotsUpLimitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LotsUpLimit, true
}

// SetLotsUpLimit sets field value
func (o *InlineResponse20079) SetLotsUpLimit(v int64) {
	o.LotsUpLimit = v
}

// GetMaxLotsPerUser returns the MaxLotsPerUser field value
func (o *InlineResponse20079) GetMaxLotsPerUser() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxLotsPerUser
}

// GetMaxLotsPerUserOk returns a tuple with the MaxLotsPerUser field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetMaxLotsPerUserOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxLotsPerUser, true
}

// SetMaxLotsPerUser sets field value
func (o *InlineResponse20079) SetMaxLotsPerUser(v int64) {
	o.MaxLotsPerUser = v
}

// GetNeedKyc returns the NeedKyc field value
func (o *InlineResponse20079) GetNeedKyc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeedKyc
}

// GetNeedKycOk returns a tuple with the NeedKyc field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetNeedKycOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NeedKyc, true
}

// SetNeedKyc sets field value
func (o *InlineResponse20079) SetNeedKyc(v bool) {
	o.NeedKyc = v
}

// GetProjectId returns the ProjectId field value
func (o *InlineResponse20079) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *InlineResponse20079) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProjectName returns the ProjectName field value
func (o *InlineResponse20079) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetProjectNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *InlineResponse20079) SetProjectName(v string) {
	o.ProjectName = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20079) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20079) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *InlineResponse20079) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20079) SetType(v string) {
	o.Type = v
}

// GetWithAreaLimitation returns the WithAreaLimitation field value
func (o *InlineResponse20079) GetWithAreaLimitation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WithAreaLimitation
}

// GetWithAreaLimitationOk returns a tuple with the WithAreaLimitation field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetWithAreaLimitationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WithAreaLimitation, true
}

// SetWithAreaLimitation sets field value
func (o *InlineResponse20079) SetWithAreaLimitation(v bool) {
	o.WithAreaLimitation = v
}

func (o InlineResponse20079) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["displayPriority"] = o.DisplayPriority
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["interestPerLot"] = o.InterestPerLot
	}
	if true {
		toSerialize["interestRate"] = o.InterestRate
	}
	if true {
		toSerialize["lotSize"] = o.LotSize
	}
	if true {
		toSerialize["lotsLowLimit"] = o.LotsLowLimit
	}
	if true {
		toSerialize["lotsPurchased"] = o.LotsPurchased
	}
	if true {
		toSerialize["lotsUpLimit"] = o.LotsUpLimit
	}
	if true {
		toSerialize["maxLotsPerUser"] = o.MaxLotsPerUser
	}
	if true {
		toSerialize["needKyc"] = o.NeedKyc
	}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	if true {
		toSerialize["projectName"] = o.ProjectName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["withAreaLimitation"] = o.WithAreaLimitation
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20079 struct {
	value *InlineResponse20079
	isSet bool
}

func (v NullableInlineResponse20079) Get() *InlineResponse20079 {
	return v.value
}

func (v *NullableInlineResponse20079) Set(val *InlineResponse20079) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079(val *InlineResponse20079) *NullableInlineResponse20079 {
	return &NullableInlineResponse20079{value: val, isSet: true}
}

func (v NullableInlineResponse20079) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


