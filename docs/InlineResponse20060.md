# InlineResponse20060

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Asset** | **string** |  | 
**Assets** | [**[]InlineResponse20060Assets**](InlineResponse20060Assets.md) |  | 
**CanDeposit** | **bool** |  | 
**CanTrade** | **bool** |  | 
**CanWithdraw** | **bool** |  | 
**FeeTier** | **int64** |  | 
**MaxWithdrawAmount** | **string** |  | 
**TotalInitialMargin** | **string** |  | 
**TotalMaintenanceMargin** | **string** |  | 
**TotalMarginBalance** | **string** |  | 
**TotalOpenOrderInitialMargin** | **string** |  | 
**TotalPositionInitialMargin** | **string** |  | 
**TotalUnrealizedProfit** | **string** |  | 
**TotalWalletBalance** | **string** |  | 
**UpdateTime** | **int64** |  | 

## Methods

### NewInlineResponse20060

`func NewInlineResponse20060(email string, asset string, assets []InlineResponse20060Assets, canDeposit bool, canTrade bool, canWithdraw bool, feeTier int64, maxWithdrawAmount string, totalInitialMargin string, totalMaintenanceMargin string, totalMarginBalance string, totalOpenOrderInitialMargin string, totalPositionInitialMargin string, totalUnrealizedProfit string, totalWalletBalance string, updateTime int64, ) *InlineResponse20060`

NewInlineResponse20060 instantiates a new InlineResponse20060 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20060WithDefaults

`func NewInlineResponse20060WithDefaults() *InlineResponse20060`

NewInlineResponse20060WithDefaults instantiates a new InlineResponse20060 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InlineResponse20060) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse20060) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse20060) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAsset

`func (o *InlineResponse20060) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *InlineResponse20060) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *InlineResponse20060) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetAssets

`func (o *InlineResponse20060) GetAssets() []InlineResponse20060Assets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *InlineResponse20060) GetAssetsOk() (*[]InlineResponse20060Assets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *InlineResponse20060) SetAssets(v []InlineResponse20060Assets)`

SetAssets sets Assets field to given value.


### GetCanDeposit

`func (o *InlineResponse20060) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *InlineResponse20060) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *InlineResponse20060) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.


### GetCanTrade

`func (o *InlineResponse20060) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *InlineResponse20060) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *InlineResponse20060) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.


### GetCanWithdraw

`func (o *InlineResponse20060) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *InlineResponse20060) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *InlineResponse20060) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.


### GetFeeTier

`func (o *InlineResponse20060) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *InlineResponse20060) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *InlineResponse20060) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.


### GetMaxWithdrawAmount

`func (o *InlineResponse20060) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *InlineResponse20060) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *InlineResponse20060) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.


### GetTotalInitialMargin

`func (o *InlineResponse20060) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *InlineResponse20060) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *InlineResponse20060) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.


### GetTotalMaintenanceMargin

`func (o *InlineResponse20060) GetTotalMaintenanceMargin() string`

GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field if non-nil, zero value otherwise.

### GetTotalMaintenanceMarginOk

`func (o *InlineResponse20060) GetTotalMaintenanceMarginOk() (*string, bool)`

GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintenanceMargin

`func (o *InlineResponse20060) SetTotalMaintenanceMargin(v string)`

SetTotalMaintenanceMargin sets TotalMaintenanceMargin field to given value.


### GetTotalMarginBalance

`func (o *InlineResponse20060) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *InlineResponse20060) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *InlineResponse20060) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.


### GetTotalOpenOrderInitialMargin

`func (o *InlineResponse20060) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *InlineResponse20060) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *InlineResponse20060) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.


### GetTotalPositionInitialMargin

`func (o *InlineResponse20060) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *InlineResponse20060) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *InlineResponse20060) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.


### GetTotalUnrealizedProfit

`func (o *InlineResponse20060) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *InlineResponse20060) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *InlineResponse20060) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.


### GetTotalWalletBalance

`func (o *InlineResponse20060) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *InlineResponse20060) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *InlineResponse20060) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.


### GetUpdateTime

`func (o *InlineResponse20060) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *InlineResponse20060) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *InlineResponse20060) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


