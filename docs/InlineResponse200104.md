# InlineResponse200104

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolId** | **int64** |  | 
**PoolNmae** | **string** |  | 
**UpdateTime** | **int64** |  | 
**Liquidity** | [**SapiV1BswapLiquidityLiquidity**](SapiV1BswapLiquidityLiquidity.md) |  | 
**Share** | [**SapiV1BswapLiquidityShare**](SapiV1BswapLiquidityShare.md) |  | 

## Methods

### NewInlineResponse200104

`func NewInlineResponse200104(poolId int64, poolNmae string, updateTime int64, liquidity SapiV1BswapLiquidityLiquidity, share SapiV1BswapLiquidityShare, ) *InlineResponse200104`

NewInlineResponse200104 instantiates a new InlineResponse200104 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200104WithDefaults

`func NewInlineResponse200104WithDefaults() *InlineResponse200104`

NewInlineResponse200104WithDefaults instantiates a new InlineResponse200104 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolId

`func (o *InlineResponse200104) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *InlineResponse200104) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *InlineResponse200104) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetPoolNmae

`func (o *InlineResponse200104) GetPoolNmae() string`

GetPoolNmae returns the PoolNmae field if non-nil, zero value otherwise.

### GetPoolNmaeOk

`func (o *InlineResponse200104) GetPoolNmaeOk() (*string, bool)`

GetPoolNmaeOk returns a tuple with the PoolNmae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolNmae

`func (o *InlineResponse200104) SetPoolNmae(v string)`

SetPoolNmae sets PoolNmae field to given value.


### GetUpdateTime

`func (o *InlineResponse200104) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *InlineResponse200104) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *InlineResponse200104) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.


### GetLiquidity

`func (o *InlineResponse200104) GetLiquidity() SapiV1BswapLiquidityLiquidity`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *InlineResponse200104) GetLiquidityOk() (*SapiV1BswapLiquidityLiquidity, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *InlineResponse200104) SetLiquidity(v SapiV1BswapLiquidityLiquidity)`

SetLiquidity sets Liquidity field to given value.


### GetShare

`func (o *InlineResponse200104) GetShare() SapiV1BswapLiquidityShare`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *InlineResponse200104) GetShareOk() (*SapiV1BswapLiquidityShare, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *InlineResponse200104) SetShare(v SapiV1BswapLiquidityShare)`

SetShare sets Share field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


