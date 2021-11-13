# BswapPoolConfigureLiquidity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstantA** | **int64** | \&quot;NA\&quot; if pool is an innovation pool | 
**MinRedeemShare** | **float64** |  | 
**SlippageTolerance** | **float64** | The swap proceeds only when the slippage is within the set range | 

## Methods

### NewBswapPoolConfigureLiquidity

`func NewBswapPoolConfigureLiquidity(constantA int64, minRedeemShare float64, slippageTolerance float64, ) *BswapPoolConfigureLiquidity`

NewBswapPoolConfigureLiquidity instantiates a new BswapPoolConfigureLiquidity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBswapPoolConfigureLiquidityWithDefaults

`func NewBswapPoolConfigureLiquidityWithDefaults() *BswapPoolConfigureLiquidity`

NewBswapPoolConfigureLiquidityWithDefaults instantiates a new BswapPoolConfigureLiquidity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstantA

`func (o *BswapPoolConfigureLiquidity) GetConstantA() int64`

GetConstantA returns the ConstantA field if non-nil, zero value otherwise.

### GetConstantAOk

`func (o *BswapPoolConfigureLiquidity) GetConstantAOk() (*int64, bool)`

GetConstantAOk returns a tuple with the ConstantA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstantA

`func (o *BswapPoolConfigureLiquidity) SetConstantA(v int64)`

SetConstantA sets ConstantA field to given value.


### GetMinRedeemShare

`func (o *BswapPoolConfigureLiquidity) GetMinRedeemShare() float64`

GetMinRedeemShare returns the MinRedeemShare field if non-nil, zero value otherwise.

### GetMinRedeemShareOk

`func (o *BswapPoolConfigureLiquidity) GetMinRedeemShareOk() (*float64, bool)`

GetMinRedeemShareOk returns a tuple with the MinRedeemShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRedeemShare

`func (o *BswapPoolConfigureLiquidity) SetMinRedeemShare(v float64)`

SetMinRedeemShare sets MinRedeemShare field to given value.


### GetSlippageTolerance

`func (o *BswapPoolConfigureLiquidity) GetSlippageTolerance() float64`

GetSlippageTolerance returns the SlippageTolerance field if non-nil, zero value otherwise.

### GetSlippageToleranceOk

`func (o *BswapPoolConfigureLiquidity) GetSlippageToleranceOk() (*float64, bool)`

GetSlippageToleranceOk returns a tuple with the SlippageTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageTolerance

`func (o *BswapPoolConfigureLiquidity) SetSlippageTolerance(v float64)`

SetSlippageTolerance sets SlippageTolerance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


