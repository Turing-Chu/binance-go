# InlineResponse200110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolId** | **int64** |  | 
**PoolNmae** | **string** |  | 
**UpdateTime** | **int64** |  | 
**Liquidity** | [**SapiV1BswapPoolConfigureLiquidity**](SapiV1BswapPoolConfigureLiquidity.md) |  | 
**AssetConfigure** | [**SapiV1BswapPoolConfigureAssetConfigure**](SapiV1BswapPoolConfigureAssetConfigure.md) |  | 

## Methods

### NewInlineResponse200110

`func NewInlineResponse200110(poolId int64, poolNmae string, updateTime int64, liquidity SapiV1BswapPoolConfigureLiquidity, assetConfigure SapiV1BswapPoolConfigureAssetConfigure, ) *InlineResponse200110`

NewInlineResponse200110 instantiates a new InlineResponse200110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110WithDefaults

`func NewInlineResponse200110WithDefaults() *InlineResponse200110`

NewInlineResponse200110WithDefaults instantiates a new InlineResponse200110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolId

`func (o *InlineResponse200110) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *InlineResponse200110) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *InlineResponse200110) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetPoolNmae

`func (o *InlineResponse200110) GetPoolNmae() string`

GetPoolNmae returns the PoolNmae field if non-nil, zero value otherwise.

### GetPoolNmaeOk

`func (o *InlineResponse200110) GetPoolNmaeOk() (*string, bool)`

GetPoolNmaeOk returns a tuple with the PoolNmae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolNmae

`func (o *InlineResponse200110) SetPoolNmae(v string)`

SetPoolNmae sets PoolNmae field to given value.


### GetUpdateTime

`func (o *InlineResponse200110) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *InlineResponse200110) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *InlineResponse200110) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.


### GetLiquidity

`func (o *InlineResponse200110) GetLiquidity() SapiV1BswapPoolConfigureLiquidity`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *InlineResponse200110) GetLiquidityOk() (*SapiV1BswapPoolConfigureLiquidity, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *InlineResponse200110) SetLiquidity(v SapiV1BswapPoolConfigureLiquidity)`

SetLiquidity sets Liquidity field to given value.


### GetAssetConfigure

`func (o *InlineResponse200110) GetAssetConfigure() SapiV1BswapPoolConfigureAssetConfigure`

GetAssetConfigure returns the AssetConfigure field if non-nil, zero value otherwise.

### GetAssetConfigureOk

`func (o *InlineResponse200110) GetAssetConfigureOk() (*SapiV1BswapPoolConfigureAssetConfigure, bool)`

GetAssetConfigureOk returns a tuple with the AssetConfigure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetConfigure

`func (o *InlineResponse200110) SetAssetConfigure(v SapiV1BswapPoolConfigureAssetConfigure)`

SetAssetConfigure sets AssetConfigure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


