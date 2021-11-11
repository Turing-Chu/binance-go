# SubAccountCOINFuturesDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Assets** | [**[]SubAccountCOINFuturesDetailsAssets**](SubAccountCOINFuturesDetailsAssets.md) |  | 
**CanDeposit** | **bool** |  | 
**CanTrade** | **bool** |  | 
**CanWithdraw** | **bool** |  | 
**FeeTier** | **int64** |  | 
**UpdateTime** | **int64** |  | 

## Methods

### NewSubAccountCOINFuturesDetails

`func NewSubAccountCOINFuturesDetails(email string, assets []SubAccountCOINFuturesDetailsAssets, canDeposit bool, canTrade bool, canWithdraw bool, feeTier int64, updateTime int64, ) *SubAccountCOINFuturesDetails`

NewSubAccountCOINFuturesDetails instantiates a new SubAccountCOINFuturesDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubAccountCOINFuturesDetailsWithDefaults

`func NewSubAccountCOINFuturesDetailsWithDefaults() *SubAccountCOINFuturesDetails`

NewSubAccountCOINFuturesDetailsWithDefaults instantiates a new SubAccountCOINFuturesDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubAccountCOINFuturesDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubAccountCOINFuturesDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubAccountCOINFuturesDetails) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAssets

`func (o *SubAccountCOINFuturesDetails) GetAssets() []SubAccountCOINFuturesDetailsAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *SubAccountCOINFuturesDetails) GetAssetsOk() (*[]SubAccountCOINFuturesDetailsAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *SubAccountCOINFuturesDetails) SetAssets(v []SubAccountCOINFuturesDetailsAssets)`

SetAssets sets Assets field to given value.


### GetCanDeposit

`func (o *SubAccountCOINFuturesDetails) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *SubAccountCOINFuturesDetails) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *SubAccountCOINFuturesDetails) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.


### GetCanTrade

`func (o *SubAccountCOINFuturesDetails) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *SubAccountCOINFuturesDetails) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *SubAccountCOINFuturesDetails) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.


### GetCanWithdraw

`func (o *SubAccountCOINFuturesDetails) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *SubAccountCOINFuturesDetails) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *SubAccountCOINFuturesDetails) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.


### GetFeeTier

`func (o *SubAccountCOINFuturesDetails) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *SubAccountCOINFuturesDetails) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *SubAccountCOINFuturesDetails) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.


### GetUpdateTime

`func (o *SubAccountCOINFuturesDetails) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *SubAccountCOINFuturesDetails) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *SubAccountCOINFuturesDetails) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


