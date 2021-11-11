# InlineResponse20040CTR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinWithdrawAmount** | **string** |  | 
**DepositStatus** | **bool** | deposit status (false if ALL of networks&#39; are false) | 
**WithdrawFee** | **int64** |  | 
**WithdrawStatus** | **bool** | withdrawStatus status (false if ALL of networks&#39; are false) | 
**DepositTip** | **string** |  | 

## Methods

### NewInlineResponse20040CTR

`func NewInlineResponse20040CTR(minWithdrawAmount string, depositStatus bool, withdrawFee int64, withdrawStatus bool, depositTip string, ) *InlineResponse20040CTR`

NewInlineResponse20040CTR instantiates a new InlineResponse20040CTR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20040CTRWithDefaults

`func NewInlineResponse20040CTRWithDefaults() *InlineResponse20040CTR`

NewInlineResponse20040CTRWithDefaults instantiates a new InlineResponse20040CTR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinWithdrawAmount

`func (o *InlineResponse20040CTR) GetMinWithdrawAmount() string`

GetMinWithdrawAmount returns the MinWithdrawAmount field if non-nil, zero value otherwise.

### GetMinWithdrawAmountOk

`func (o *InlineResponse20040CTR) GetMinWithdrawAmountOk() (*string, bool)`

GetMinWithdrawAmountOk returns a tuple with the MinWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWithdrawAmount

`func (o *InlineResponse20040CTR) SetMinWithdrawAmount(v string)`

SetMinWithdrawAmount sets MinWithdrawAmount field to given value.


### GetDepositStatus

`func (o *InlineResponse20040CTR) GetDepositStatus() bool`

GetDepositStatus returns the DepositStatus field if non-nil, zero value otherwise.

### GetDepositStatusOk

`func (o *InlineResponse20040CTR) GetDepositStatusOk() (*bool, bool)`

GetDepositStatusOk returns a tuple with the DepositStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositStatus

`func (o *InlineResponse20040CTR) SetDepositStatus(v bool)`

SetDepositStatus sets DepositStatus field to given value.


### GetWithdrawFee

`func (o *InlineResponse20040CTR) GetWithdrawFee() int64`

GetWithdrawFee returns the WithdrawFee field if non-nil, zero value otherwise.

### GetWithdrawFeeOk

`func (o *InlineResponse20040CTR) GetWithdrawFeeOk() (*int64, bool)`

GetWithdrawFeeOk returns a tuple with the WithdrawFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawFee

`func (o *InlineResponse20040CTR) SetWithdrawFee(v int64)`

SetWithdrawFee sets WithdrawFee field to given value.


### GetWithdrawStatus

`func (o *InlineResponse20040CTR) GetWithdrawStatus() bool`

GetWithdrawStatus returns the WithdrawStatus field if non-nil, zero value otherwise.

### GetWithdrawStatusOk

`func (o *InlineResponse20040CTR) GetWithdrawStatusOk() (*bool, bool)`

GetWithdrawStatusOk returns a tuple with the WithdrawStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawStatus

`func (o *InlineResponse20040CTR) SetWithdrawStatus(v bool)`

SetWithdrawStatus sets WithdrawStatus field to given value.


### GetDepositTip

`func (o *InlineResponse20040CTR) GetDepositTip() string`

GetDepositTip returns the DepositTip field if non-nil, zero value otherwise.

### GetDepositTipOk

`func (o *InlineResponse20040CTR) GetDepositTipOk() (*string, bool)`

GetDepositTipOk returns a tuple with the DepositTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositTip

`func (o *InlineResponse20040CTR) SetDepositTip(v string)`

SetDepositTip sets DepositTip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


