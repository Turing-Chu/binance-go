/*
Binance Public Spot API

OpenAPI Specifications for the Binance Public Spot API generated with [binance/binance-api-swagger/blob/master/spot_api.yaml](https://github.com/binance/binance-api-swagger/blob/master/spot_api.yaml) with commit [v1.2.0 release](https://github.com/binance/binance-api-swagger/commit/60d14be031c031600c853d5cdab86db5ab73603e)  API documents:   - [https://github.com/binance/binance-spot-api-docs](https://github.com/binance/binance-spot-api-docs)   - [https://binance-docs.github.io/apidocs/spot/en](https://binance-docs.github.io/apidocs/spot/en)

API version: 1.0
Contact: qishiwenjun@163.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binanceapi

import (
	"encoding/json"
)

// InlineResponse200110 struct for InlineResponse200110
type InlineResponse200110 struct {
	PoolId         int64                            `json:"poolId"`
	PoolNmae       string                           `json:"poolNmae"`
	UpdateTime     int64                            `json:"updateTime"`
	Liquidity      BswapPoolConfigureLiquidity      `json:"liquidity"`
	AssetConfigure BswapPoolConfigureAssetConfigure `json:"assetConfigure"`
}

// NewInlineResponse200110 instantiates a new InlineResponse200110 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200110(poolId int64, poolNmae string, updateTime int64, liquidity BswapPoolConfigureLiquidity, assetConfigure BswapPoolConfigureAssetConfigure) *InlineResponse200110 {
	this := InlineResponse200110{}
	this.PoolId = poolId
	this.PoolNmae = poolNmae
	this.UpdateTime = updateTime
	this.Liquidity = liquidity
	this.AssetConfigure = assetConfigure
	return &this
}

// NewInlineResponse200110WithDefaults instantiates a new InlineResponse200110 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200110WithDefaults() *InlineResponse200110 {
	this := InlineResponse200110{}
	return &this
}

// GetPoolId returns the PoolId field value
func (o *InlineResponse200110) GetPoolId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetPoolIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *InlineResponse200110) SetPoolId(v int64) {
	o.PoolId = v
}

// GetPoolNmae returns the PoolNmae field value
func (o *InlineResponse200110) GetPoolNmae() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolNmae
}

// GetPoolNmaeOk returns a tuple with the PoolNmae field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetPoolNmaeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolNmae, true
}

// SetPoolNmae sets field value
func (o *InlineResponse200110) SetPoolNmae(v string) {
	o.PoolNmae = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *InlineResponse200110) GetUpdateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetUpdateTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *InlineResponse200110) SetUpdateTime(v int64) {
	o.UpdateTime = v
}

// GetLiquidity returns the Liquidity field value
func (o *InlineResponse200110) GetLiquidity() BswapPoolConfigureLiquidity {
	if o == nil {
		var ret BswapPoolConfigureLiquidity
		return ret
	}

	return o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetLiquidityOk() (*BswapPoolConfigureLiquidity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Liquidity, true
}

// SetLiquidity sets field value
func (o *InlineResponse200110) SetLiquidity(v BswapPoolConfigureLiquidity) {
	o.Liquidity = v
}

// GetAssetConfigure returns the AssetConfigure field value
func (o *InlineResponse200110) GetAssetConfigure() BswapPoolConfigureAssetConfigure {
	if o == nil {
		var ret BswapPoolConfigureAssetConfigure
		return ret
	}

	return o.AssetConfigure
}

// GetAssetConfigureOk returns a tuple with the AssetConfigure field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetAssetConfigureOk() (*BswapPoolConfigureAssetConfigure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetConfigure, true
}

// SetAssetConfigure sets field value
func (o *InlineResponse200110) SetAssetConfigure(v BswapPoolConfigureAssetConfigure) {
	o.AssetConfigure = v
}

func (o InlineResponse200110) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["poolId"] = o.PoolId
	}
	if true {
		toSerialize["poolNmae"] = o.PoolNmae
	}
	if true {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if true {
		toSerialize["liquidity"] = o.Liquidity
	}
	if true {
		toSerialize["assetConfigure"] = o.AssetConfigure
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200110 struct {
	value *InlineResponse200110
	isSet bool
}

func (v NullableInlineResponse200110) Get() *InlineResponse200110 {
	return v.value
}

func (v *NullableInlineResponse200110) Set(val *InlineResponse200110) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200110) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200110) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200110(val *InlineResponse200110) *NullableInlineResponse200110 {
	return &NullableInlineResponse200110{value: val, isSet: true}
}

func (v NullableInlineResponse200110) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200110) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
