# InlineResponse20033

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Amount** | **string** |  | 
**ApplyTime** | **string** |  | 
**Coin** | **string** |  | 
**Id** | **string** |  | 
**WithdrawOrderId** | **string** | will not be returned if there&#39;s no withdrawOrderId for this withdraw. | 
**Network** | **string** |  | 
**TransferType** | **int32** | 1 for internal transfer, 0 for external transfer | 
**Status** | **int32** |  | 
**TransactionFee** | **string** |  | 
**ConfirmNo** | **int32** |  | 
**TxId** | **string** |  | 

## Methods

### NewInlineResponse20033

`func NewInlineResponse20033(address string, amount string, applyTime string, coin string, id string, withdrawOrderId string, network string, transferType int32, status int32, transactionFee string, confirmNo int32, txId string, ) *InlineResponse20033`

NewInlineResponse20033 instantiates a new InlineResponse20033 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20033WithDefaults

`func NewInlineResponse20033WithDefaults() *InlineResponse20033`

NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *InlineResponse20033) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse20033) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse20033) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *InlineResponse20033) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InlineResponse20033) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InlineResponse20033) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetApplyTime

`func (o *InlineResponse20033) GetApplyTime() string`

GetApplyTime returns the ApplyTime field if non-nil, zero value otherwise.

### GetApplyTimeOk

`func (o *InlineResponse20033) GetApplyTimeOk() (*string, bool)`

GetApplyTimeOk returns a tuple with the ApplyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTime

`func (o *InlineResponse20033) SetApplyTime(v string)`

SetApplyTime sets ApplyTime field to given value.


### GetCoin

`func (o *InlineResponse20033) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *InlineResponse20033) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *InlineResponse20033) SetCoin(v string)`

SetCoin sets Coin field to given value.


### GetId

`func (o *InlineResponse20033) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20033) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20033) SetId(v string)`

SetId sets Id field to given value.


### GetWithdrawOrderId

`func (o *InlineResponse20033) GetWithdrawOrderId() string`

GetWithdrawOrderId returns the WithdrawOrderId field if non-nil, zero value otherwise.

### GetWithdrawOrderIdOk

`func (o *InlineResponse20033) GetWithdrawOrderIdOk() (*string, bool)`

GetWithdrawOrderIdOk returns a tuple with the WithdrawOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawOrderId

`func (o *InlineResponse20033) SetWithdrawOrderId(v string)`

SetWithdrawOrderId sets WithdrawOrderId field to given value.


### GetNetwork

`func (o *InlineResponse20033) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20033) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20033) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTransferType

`func (o *InlineResponse20033) GetTransferType() int32`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *InlineResponse20033) GetTransferTypeOk() (*int32, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *InlineResponse20033) SetTransferType(v int32)`

SetTransferType sets TransferType field to given value.


### GetStatus

`func (o *InlineResponse20033) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20033) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20033) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTransactionFee

`func (o *InlineResponse20033) GetTransactionFee() string`

GetTransactionFee returns the TransactionFee field if non-nil, zero value otherwise.

### GetTransactionFeeOk

`func (o *InlineResponse20033) GetTransactionFeeOk() (*string, bool)`

GetTransactionFeeOk returns a tuple with the TransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionFee

`func (o *InlineResponse20033) SetTransactionFee(v string)`

SetTransactionFee sets TransactionFee field to given value.


### GetConfirmNo

`func (o *InlineResponse20033) GetConfirmNo() int32`

GetConfirmNo returns the ConfirmNo field if non-nil, zero value otherwise.

### GetConfirmNoOk

`func (o *InlineResponse20033) GetConfirmNoOk() (*int32, bool)`

GetConfirmNoOk returns a tuple with the ConfirmNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmNo

`func (o *InlineResponse20033) SetConfirmNo(v int32)`

SetConfirmNo sets ConfirmNo field to given value.


### GetTxId

`func (o *InlineResponse20033) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *InlineResponse20033) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *InlineResponse20033) SetTxId(v string)`

SetTxId sets TxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


