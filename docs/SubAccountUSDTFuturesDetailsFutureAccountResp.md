# SubAccountUSDTFuturesDetailsFutureAccountResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
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

### NewSubAccountUSDTFuturesDetailsFutureAccountResp

`func NewSubAccountUSDTFuturesDetailsFutureAccountResp(email string, assets []InlineResponse20060Assets, canDeposit bool, canTrade bool, canWithdraw bool, feeTier int64, maxWithdrawAmount string, totalInitialMargin string, totalMaintenanceMargin string, totalMarginBalance string, totalOpenOrderInitialMargin string, totalPositionInitialMargin string, totalUnrealizedProfit string, totalWalletBalance string, updateTime int64, ) *SubAccountUSDTFuturesDetailsFutureAccountResp`

NewSubAccountUSDTFuturesDetailsFutureAccountResp instantiates a new SubAccountUSDTFuturesDetailsFutureAccountResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubAccountUSDTFuturesDetailsFutureAccountRespWithDefaults

`func NewSubAccountUSDTFuturesDetailsFutureAccountRespWithDefaults() *SubAccountUSDTFuturesDetailsFutureAccountResp`

NewSubAccountUSDTFuturesDetailsFutureAccountRespWithDefaults instantiates a new SubAccountUSDTFuturesDetailsFutureAccountResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAssets

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetAssets() []InlineResponse20060Assets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetAssetsOk() (*[]InlineResponse20060Assets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetAssets(v []InlineResponse20060Assets)`

SetAssets sets Assets field to given value.


### GetCanDeposit

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.


### GetCanTrade

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.


### GetCanWithdraw

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.


### GetFeeTier

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.


### GetMaxWithdrawAmount

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.


### GetTotalInitialMargin

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.


### GetTotalMaintenanceMargin

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalMaintenanceMargin() string`

GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field if non-nil, zero value otherwise.

### GetTotalMaintenanceMarginOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalMaintenanceMarginOk() (*string, bool)`

GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintenanceMargin

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetTotalMaintenanceMargin(v string)`

SetTotalMaintenanceMargin sets TotalMaintenanceMargin field to given value.


### GetTotalMarginBalance

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.


### GetTotalOpenOrderInitialMargin

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.


### GetTotalPositionInitialMargin

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.


### GetTotalUnrealizedProfit

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.


### GetTotalWalletBalance

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.


### GetUpdateTime

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *SubAccountUSDTFuturesDetailsFutureAccountResp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


