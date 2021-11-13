# BswapLiquidityShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShareAmount** | **float64** |  | 
**SharePercentage** | **float64** |  | 
**Asset** | [**BswapLiquidityShareAsset**](BswapLiquidityShareAsset.md) |  | 

## Methods

### NewBswapLiquidityShare

`func NewBswapLiquidityShare(shareAmount float64, sharePercentage float64, asset BswapLiquidityShareAsset, ) *BswapLiquidityShare`

NewBswapLiquidityShare instantiates a new BswapLiquidityShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBswapLiquidityShareWithDefaults

`func NewBswapLiquidityShareWithDefaults() *BswapLiquidityShare`

NewBswapLiquidityShareWithDefaults instantiates a new BswapLiquidityShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShareAmount

`func (o *BswapLiquidityShare) GetShareAmount() float64`

GetShareAmount returns the ShareAmount field if non-nil, zero value otherwise.

### GetShareAmountOk

`func (o *BswapLiquidityShare) GetShareAmountOk() (*float64, bool)`

GetShareAmountOk returns a tuple with the ShareAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareAmount

`func (o *BswapLiquidityShare) SetShareAmount(v float64)`

SetShareAmount sets ShareAmount field to given value.


### GetSharePercentage

`func (o *BswapLiquidityShare) GetSharePercentage() float64`

GetSharePercentage returns the SharePercentage field if non-nil, zero value otherwise.

### GetSharePercentageOk

`func (o *BswapLiquidityShare) GetSharePercentageOk() (*float64, bool)`

GetSharePercentageOk returns a tuple with the SharePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePercentage

`func (o *BswapLiquidityShare) SetSharePercentage(v float64)`

SetSharePercentage sets SharePercentage field to given value.


### GetAsset

`func (o *BswapLiquidityShare) GetAsset() BswapLiquidityShareAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *BswapLiquidityShare) GetAssetOk() (*BswapLiquidityShareAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *BswapLiquidityShare) SetAsset(v BswapLiquidityShareAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


