# CapitalConfigGetallNetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRegex** | **string** |  | 
**Coin** | **string** |  | 
**DepositDesc** | **string** | shown only when \&quot;depositEnable\&quot; is false. | 
**DepositEnable** | **bool** |  | 
**IsDefault** | **bool** |  | 
**MemoRegex** | **string** |  | 
**MinConfirm** | **int64** | min number for balance confirmation. | 
**Name** | **string** |  | 
**ResetAddressStatus** | **bool** |  | 
**SpecialTips** | **string** |  | 
**UnLockConfirm** | **int64** | confirmation number for balance unlock. | 
**WithdrawDesc** | **string** | shown only when \&quot;withdrawEnable\&quot; is false | 
**WithdrawEnable** | **bool** |  | 
**WithdrawFee** | **string** |  | 
**WithdrawIntegerMultiple** | **string** |  | 
**WithdrawMax** | **string** |  | 
**WithdrawMin** | **string** |  | 
**SameAddress** | **bool** |  | 

## Methods

### NewCapitalConfigGetallNetworkList

`func NewCapitalConfigGetallNetworkList(addressRegex string, coin string, depositDesc string, depositEnable bool, isDefault bool, memoRegex string, minConfirm int64, name string, resetAddressStatus bool, specialTips string, unLockConfirm int64, withdrawDesc string, withdrawEnable bool, withdrawFee string, withdrawIntegerMultiple string, withdrawMax string, withdrawMin string, sameAddress bool, ) *CapitalConfigGetallNetworkList`

NewCapitalConfigGetallNetworkList instantiates a new CapitalConfigGetallNetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapitalConfigGetallNetworkListWithDefaults

`func NewCapitalConfigGetallNetworkListWithDefaults() *CapitalConfigGetallNetworkList`

NewCapitalConfigGetallNetworkListWithDefaults instantiates a new CapitalConfigGetallNetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRegex

`func (o *CapitalConfigGetallNetworkList) GetAddressRegex() string`

GetAddressRegex returns the AddressRegex field if non-nil, zero value otherwise.

### GetAddressRegexOk

`func (o *CapitalConfigGetallNetworkList) GetAddressRegexOk() (*string, bool)`

GetAddressRegexOk returns a tuple with the AddressRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRegex

`func (o *CapitalConfigGetallNetworkList) SetAddressRegex(v string)`

SetAddressRegex sets AddressRegex field to given value.


### GetCoin

`func (o *CapitalConfigGetallNetworkList) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *CapitalConfigGetallNetworkList) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *CapitalConfigGetallNetworkList) SetCoin(v string)`

SetCoin sets Coin field to given value.


### GetDepositDesc

`func (o *CapitalConfigGetallNetworkList) GetDepositDesc() string`

GetDepositDesc returns the DepositDesc field if non-nil, zero value otherwise.

### GetDepositDescOk

`func (o *CapitalConfigGetallNetworkList) GetDepositDescOk() (*string, bool)`

GetDepositDescOk returns a tuple with the DepositDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositDesc

`func (o *CapitalConfigGetallNetworkList) SetDepositDesc(v string)`

SetDepositDesc sets DepositDesc field to given value.


### GetDepositEnable

`func (o *CapitalConfigGetallNetworkList) GetDepositEnable() bool`

GetDepositEnable returns the DepositEnable field if non-nil, zero value otherwise.

### GetDepositEnableOk

`func (o *CapitalConfigGetallNetworkList) GetDepositEnableOk() (*bool, bool)`

GetDepositEnableOk returns a tuple with the DepositEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositEnable

`func (o *CapitalConfigGetallNetworkList) SetDepositEnable(v bool)`

SetDepositEnable sets DepositEnable field to given value.


### GetIsDefault

`func (o *CapitalConfigGetallNetworkList) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CapitalConfigGetallNetworkList) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CapitalConfigGetallNetworkList) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetMemoRegex

`func (o *CapitalConfigGetallNetworkList) GetMemoRegex() string`

GetMemoRegex returns the MemoRegex field if non-nil, zero value otherwise.

### GetMemoRegexOk

`func (o *CapitalConfigGetallNetworkList) GetMemoRegexOk() (*string, bool)`

GetMemoRegexOk returns a tuple with the MemoRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoRegex

`func (o *CapitalConfigGetallNetworkList) SetMemoRegex(v string)`

SetMemoRegex sets MemoRegex field to given value.


### GetMinConfirm

`func (o *CapitalConfigGetallNetworkList) GetMinConfirm() int64`

GetMinConfirm returns the MinConfirm field if non-nil, zero value otherwise.

### GetMinConfirmOk

`func (o *CapitalConfigGetallNetworkList) GetMinConfirmOk() (*int64, bool)`

GetMinConfirmOk returns a tuple with the MinConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinConfirm

`func (o *CapitalConfigGetallNetworkList) SetMinConfirm(v int64)`

SetMinConfirm sets MinConfirm field to given value.


### GetName

`func (o *CapitalConfigGetallNetworkList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapitalConfigGetallNetworkList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapitalConfigGetallNetworkList) SetName(v string)`

SetName sets Name field to given value.


### GetResetAddressStatus

`func (o *CapitalConfigGetallNetworkList) GetResetAddressStatus() bool`

GetResetAddressStatus returns the ResetAddressStatus field if non-nil, zero value otherwise.

### GetResetAddressStatusOk

`func (o *CapitalConfigGetallNetworkList) GetResetAddressStatusOk() (*bool, bool)`

GetResetAddressStatusOk returns a tuple with the ResetAddressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAddressStatus

`func (o *CapitalConfigGetallNetworkList) SetResetAddressStatus(v bool)`

SetResetAddressStatus sets ResetAddressStatus field to given value.


### GetSpecialTips

`func (o *CapitalConfigGetallNetworkList) GetSpecialTips() string`

GetSpecialTips returns the SpecialTips field if non-nil, zero value otherwise.

### GetSpecialTipsOk

`func (o *CapitalConfigGetallNetworkList) GetSpecialTipsOk() (*string, bool)`

GetSpecialTipsOk returns a tuple with the SpecialTips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialTips

`func (o *CapitalConfigGetallNetworkList) SetSpecialTips(v string)`

SetSpecialTips sets SpecialTips field to given value.


### GetUnLockConfirm

`func (o *CapitalConfigGetallNetworkList) GetUnLockConfirm() int64`

GetUnLockConfirm returns the UnLockConfirm field if non-nil, zero value otherwise.

### GetUnLockConfirmOk

`func (o *CapitalConfigGetallNetworkList) GetUnLockConfirmOk() (*int64, bool)`

GetUnLockConfirmOk returns a tuple with the UnLockConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnLockConfirm

`func (o *CapitalConfigGetallNetworkList) SetUnLockConfirm(v int64)`

SetUnLockConfirm sets UnLockConfirm field to given value.


### GetWithdrawDesc

`func (o *CapitalConfigGetallNetworkList) GetWithdrawDesc() string`

GetWithdrawDesc returns the WithdrawDesc field if non-nil, zero value otherwise.

### GetWithdrawDescOk

`func (o *CapitalConfigGetallNetworkList) GetWithdrawDescOk() (*string, bool)`

GetWithdrawDescOk returns a tuple with the WithdrawDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawDesc

`func (o *CapitalConfigGetallNetworkList) SetWithdrawDesc(v string)`

SetWithdrawDesc sets WithdrawDesc field to given value.


### GetWithdrawEnable

`func (o *CapitalConfigGetallNetworkList) GetWithdrawEnable() bool`

GetWithdrawEnable returns the WithdrawEnable field if non-nil, zero value otherwise.

### GetWithdrawEnableOk

`func (o *CapitalConfigGetallNetworkList) GetWithdrawEnableOk() (*bool, bool)`

GetWithdrawEnableOk returns a tuple with the WithdrawEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawEnable

`func (o *CapitalConfigGetallNetworkList) SetWithdrawEnable(v bool)`

SetWithdrawEnable sets WithdrawEnable field to given value.


### GetWithdrawFee

`func (o *CapitalConfigGetallNetworkList) GetWithdrawFee() string`

GetWithdrawFee returns the WithdrawFee field if non-nil, zero value otherwise.

### GetWithdrawFeeOk

`func (o *CapitalConfigGetallNetworkList) GetWithdrawFeeOk() (*string, bool)`

GetWithdrawFeeOk returns a tuple with the WithdrawFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawFee

`func (o *CapitalConfigGetallNetworkList) SetWithdrawFee(v string)`

SetWithdrawFee sets WithdrawFee field to given value.


### GetWithdrawIntegerMultiple

`func (o *CapitalConfigGetallNetworkList) GetWithdrawIntegerMultiple() string`

GetWithdrawIntegerMultiple returns the WithdrawIntegerMultiple field if non-nil, zero value otherwise.

### GetWithdrawIntegerMultipleOk

`func (o *CapitalConfigGetallNetworkList) GetWithdrawIntegerMultipleOk() (*string, bool)`

GetWithdrawIntegerMultipleOk returns a tuple with the WithdrawIntegerMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawIntegerMultiple

`func (o *CapitalConfigGetallNetworkList) SetWithdrawIntegerMultiple(v string)`

SetWithdrawIntegerMultiple sets WithdrawIntegerMultiple field to given value.


### GetWithdrawMax

`func (o *CapitalConfigGetallNetworkList) GetWithdrawMax() string`

GetWithdrawMax returns the WithdrawMax field if non-nil, zero value otherwise.

### GetWithdrawMaxOk

`func (o *CapitalConfigGetallNetworkList) GetWithdrawMaxOk() (*string, bool)`

GetWithdrawMaxOk returns a tuple with the WithdrawMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawMax

`func (o *CapitalConfigGetallNetworkList) SetWithdrawMax(v string)`

SetWithdrawMax sets WithdrawMax field to given value.


### GetWithdrawMin

`func (o *CapitalConfigGetallNetworkList) GetWithdrawMin() string`

GetWithdrawMin returns the WithdrawMin field if non-nil, zero value otherwise.

### GetWithdrawMinOk

`func (o *CapitalConfigGetallNetworkList) GetWithdrawMinOk() (*string, bool)`

GetWithdrawMinOk returns a tuple with the WithdrawMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawMin

`func (o *CapitalConfigGetallNetworkList) SetWithdrawMin(v string)`

SetWithdrawMin sets WithdrawMin field to given value.


### GetSameAddress

`func (o *CapitalConfigGetallNetworkList) GetSameAddress() bool`

GetSameAddress returns the SameAddress field if non-nil, zero value otherwise.

### GetSameAddressOk

`func (o *CapitalConfigGetallNetworkList) GetSameAddressOk() (*bool, bool)`

GetSameAddressOk returns a tuple with the SameAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameAddress

`func (o *CapitalConfigGetallNetworkList) SetSameAddress(v bool)`

SetSameAddress sets SameAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


