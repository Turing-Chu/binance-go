# CoinInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coin** | **string** |  | 
**DepositAllEnable** | **bool** |  | 
**Free** | **string** |  | 
**Freeze** | **string** |  | 
**Ipoable** | **string** |  | 
**Ipoing** | **string** |  | 
**IsLegalMoney** | **bool** |  | 
**Locked** | **string** |  | 
**Name** | **string** |  | 
**NetworkList** | [**[]CapitalConfigGetallNetworkList**](CapitalConfigGetallNetworkList.md) |  | 
**Storage** | **string** |  | 
**Trading** | **bool** |  | 
**WithdrawAllEnable** | **bool** |  | 
**Withdrawing** | **string** |  | 

## Methods

### NewCoinInfo

`func NewCoinInfo(coin string, depositAllEnable bool, free string, freeze string, ipoable string, ipoing string, isLegalMoney bool, locked string, name string, networkList []CapitalConfigGetallNetworkList, storage string, trading bool, withdrawAllEnable bool, withdrawing string, ) *CoinInfo`

NewCoinInfo instantiates a new CoinInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoinInfoWithDefaults

`func NewCoinInfoWithDefaults() *CoinInfo`

NewCoinInfoWithDefaults instantiates a new CoinInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoin

`func (o *CoinInfo) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *CoinInfo) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *CoinInfo) SetCoin(v string)`

SetCoin sets Coin field to given value.


### GetDepositAllEnable

`func (o *CoinInfo) GetDepositAllEnable() bool`

GetDepositAllEnable returns the DepositAllEnable field if non-nil, zero value otherwise.

### GetDepositAllEnableOk

`func (o *CoinInfo) GetDepositAllEnableOk() (*bool, bool)`

GetDepositAllEnableOk returns a tuple with the DepositAllEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAllEnable

`func (o *CoinInfo) SetDepositAllEnable(v bool)`

SetDepositAllEnable sets DepositAllEnable field to given value.


### GetFree

`func (o *CoinInfo) GetFree() string`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *CoinInfo) GetFreeOk() (*string, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *CoinInfo) SetFree(v string)`

SetFree sets Free field to given value.


### GetFreeze

`func (o *CoinInfo) GetFreeze() string`

GetFreeze returns the Freeze field if non-nil, zero value otherwise.

### GetFreezeOk

`func (o *CoinInfo) GetFreezeOk() (*string, bool)`

GetFreezeOk returns a tuple with the Freeze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeze

`func (o *CoinInfo) SetFreeze(v string)`

SetFreeze sets Freeze field to given value.


### GetIpoable

`func (o *CoinInfo) GetIpoable() string`

GetIpoable returns the Ipoable field if non-nil, zero value otherwise.

### GetIpoableOk

`func (o *CoinInfo) GetIpoableOk() (*string, bool)`

GetIpoableOk returns a tuple with the Ipoable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpoable

`func (o *CoinInfo) SetIpoable(v string)`

SetIpoable sets Ipoable field to given value.


### GetIpoing

`func (o *CoinInfo) GetIpoing() string`

GetIpoing returns the Ipoing field if non-nil, zero value otherwise.

### GetIpoingOk

`func (o *CoinInfo) GetIpoingOk() (*string, bool)`

GetIpoingOk returns a tuple with the Ipoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpoing

`func (o *CoinInfo) SetIpoing(v string)`

SetIpoing sets Ipoing field to given value.


### GetIsLegalMoney

`func (o *CoinInfo) GetIsLegalMoney() bool`

GetIsLegalMoney returns the IsLegalMoney field if non-nil, zero value otherwise.

### GetIsLegalMoneyOk

`func (o *CoinInfo) GetIsLegalMoneyOk() (*bool, bool)`

GetIsLegalMoneyOk returns a tuple with the IsLegalMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegalMoney

`func (o *CoinInfo) SetIsLegalMoney(v bool)`

SetIsLegalMoney sets IsLegalMoney field to given value.


### GetLocked

`func (o *CoinInfo) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CoinInfo) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CoinInfo) SetLocked(v string)`

SetLocked sets Locked field to given value.


### GetName

`func (o *CoinInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoinInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoinInfo) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkList

`func (o *CoinInfo) GetNetworkList() []CapitalConfigGetallNetworkList`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *CoinInfo) GetNetworkListOk() (*[]CapitalConfigGetallNetworkList, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *CoinInfo) SetNetworkList(v []CapitalConfigGetallNetworkList)`

SetNetworkList sets NetworkList field to given value.


### GetStorage

`func (o *CoinInfo) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CoinInfo) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CoinInfo) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetTrading

`func (o *CoinInfo) GetTrading() bool`

GetTrading returns the Trading field if non-nil, zero value otherwise.

### GetTradingOk

`func (o *CoinInfo) GetTradingOk() (*bool, bool)`

GetTradingOk returns a tuple with the Trading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrading

`func (o *CoinInfo) SetTrading(v bool)`

SetTrading sets Trading field to given value.


### GetWithdrawAllEnable

`func (o *CoinInfo) GetWithdrawAllEnable() bool`

GetWithdrawAllEnable returns the WithdrawAllEnable field if non-nil, zero value otherwise.

### GetWithdrawAllEnableOk

`func (o *CoinInfo) GetWithdrawAllEnableOk() (*bool, bool)`

GetWithdrawAllEnableOk returns a tuple with the WithdrawAllEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawAllEnable

`func (o *CoinInfo) SetWithdrawAllEnable(v bool)`

SetWithdrawAllEnable sets WithdrawAllEnable field to given value.


### GetWithdrawing

`func (o *CoinInfo) GetWithdrawing() string`

GetWithdrawing returns the Withdrawing field if non-nil, zero value otherwise.

### GetWithdrawingOk

`func (o *CoinInfo) GetWithdrawingOk() (*string, bool)`

GetWithdrawingOk returns a tuple with the Withdrawing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawing

`func (o *CoinInfo) SetWithdrawing(v string)`

SetWithdrawing sets Withdrawing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


