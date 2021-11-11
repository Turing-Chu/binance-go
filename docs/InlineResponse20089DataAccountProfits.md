# InlineResponse20089DataAccountProfits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **int64** | Mining date | 
**Type** | **int64** | 0:Mining Wallet,5:Mining Address,7:Pool Savings,8:Transferred,31:Income Transfer ,32:Hashrate Resale-Mining Wallet 33:Hashrate Resale-Pool Savings | 
**HashTransfer** | **int32** | Transferred Hashrate | 
**TransferAmount** | **float32** | Transferred Income | 
**DayHashRate** | **int64** | Daily Hashrate | 
**ProfitAmount** | **float64** | Earnings Amount | 
**CoinName** | **string** | Coin Type | 
**Status** | **int32** | Status：0:Unpaid, 1:Paying  2：Paid | 

## Methods

### NewInlineResponse20089DataAccountProfits

`func NewInlineResponse20089DataAccountProfits(time int64, type_ int64, hashTransfer int32, transferAmount float32, dayHashRate int64, profitAmount float64, coinName string, status int32, ) *InlineResponse20089DataAccountProfits`

NewInlineResponse20089DataAccountProfits instantiates a new InlineResponse20089DataAccountProfits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20089DataAccountProfitsWithDefaults

`func NewInlineResponse20089DataAccountProfitsWithDefaults() *InlineResponse20089DataAccountProfits`

NewInlineResponse20089DataAccountProfitsWithDefaults instantiates a new InlineResponse20089DataAccountProfits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse20089DataAccountProfits) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse20089DataAccountProfits) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse20089DataAccountProfits) SetTime(v int64)`

SetTime sets Time field to given value.


### GetType

`func (o *InlineResponse20089DataAccountProfits) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20089DataAccountProfits) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20089DataAccountProfits) SetType(v int64)`

SetType sets Type field to given value.


### GetHashTransfer

`func (o *InlineResponse20089DataAccountProfits) GetHashTransfer() int32`

GetHashTransfer returns the HashTransfer field if non-nil, zero value otherwise.

### GetHashTransferOk

`func (o *InlineResponse20089DataAccountProfits) GetHashTransferOk() (*int32, bool)`

GetHashTransferOk returns a tuple with the HashTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTransfer

`func (o *InlineResponse20089DataAccountProfits) SetHashTransfer(v int32)`

SetHashTransfer sets HashTransfer field to given value.


### GetTransferAmount

`func (o *InlineResponse20089DataAccountProfits) GetTransferAmount() float32`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *InlineResponse20089DataAccountProfits) GetTransferAmountOk() (*float32, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *InlineResponse20089DataAccountProfits) SetTransferAmount(v float32)`

SetTransferAmount sets TransferAmount field to given value.


### GetDayHashRate

`func (o *InlineResponse20089DataAccountProfits) GetDayHashRate() int64`

GetDayHashRate returns the DayHashRate field if non-nil, zero value otherwise.

### GetDayHashRateOk

`func (o *InlineResponse20089DataAccountProfits) GetDayHashRateOk() (*int64, bool)`

GetDayHashRateOk returns a tuple with the DayHashRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayHashRate

`func (o *InlineResponse20089DataAccountProfits) SetDayHashRate(v int64)`

SetDayHashRate sets DayHashRate field to given value.


### GetProfitAmount

`func (o *InlineResponse20089DataAccountProfits) GetProfitAmount() float64`

GetProfitAmount returns the ProfitAmount field if non-nil, zero value otherwise.

### GetProfitAmountOk

`func (o *InlineResponse20089DataAccountProfits) GetProfitAmountOk() (*float64, bool)`

GetProfitAmountOk returns a tuple with the ProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitAmount

`func (o *InlineResponse20089DataAccountProfits) SetProfitAmount(v float64)`

SetProfitAmount sets ProfitAmount field to given value.


### GetCoinName

`func (o *InlineResponse20089DataAccountProfits) GetCoinName() string`

GetCoinName returns the CoinName field if non-nil, zero value otherwise.

### GetCoinNameOk

`func (o *InlineResponse20089DataAccountProfits) GetCoinNameOk() (*string, bool)`

GetCoinNameOk returns a tuple with the CoinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinName

`func (o *InlineResponse20089DataAccountProfits) SetCoinName(v string)`

SetCoinName sets CoinName field to given value.


### GetStatus

`func (o *InlineResponse20089DataAccountProfits) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20089DataAccountProfits) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20089DataAccountProfits) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


