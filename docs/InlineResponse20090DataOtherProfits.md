# InlineResponse20090DataOtherProfits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **int64** | Mining date | 
**CoinName** | **string** | Coin Name | 
**Type** | **int32** | 1: Merged Mining, 2: Activity Bonus, 3:Rebate 4:Smart Pool 6:Income Transfer 7:Pool Savings | 
**ProfitAmount** | **float64** |  | 
**Status** | **int32** | 0:Unpaid, 1:Paying  2：Paid | 

## Methods

### NewInlineResponse20090DataOtherProfits

`func NewInlineResponse20090DataOtherProfits(time int64, coinName string, type_ int32, profitAmount float64, status int32, ) *InlineResponse20090DataOtherProfits`

NewInlineResponse20090DataOtherProfits instantiates a new InlineResponse20090DataOtherProfits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20090DataOtherProfitsWithDefaults

`func NewInlineResponse20090DataOtherProfitsWithDefaults() *InlineResponse20090DataOtherProfits`

NewInlineResponse20090DataOtherProfitsWithDefaults instantiates a new InlineResponse20090DataOtherProfits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse20090DataOtherProfits) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse20090DataOtherProfits) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse20090DataOtherProfits) SetTime(v int64)`

SetTime sets Time field to given value.


### GetCoinName

`func (o *InlineResponse20090DataOtherProfits) GetCoinName() string`

GetCoinName returns the CoinName field if non-nil, zero value otherwise.

### GetCoinNameOk

`func (o *InlineResponse20090DataOtherProfits) GetCoinNameOk() (*string, bool)`

GetCoinNameOk returns a tuple with the CoinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinName

`func (o *InlineResponse20090DataOtherProfits) SetCoinName(v string)`

SetCoinName sets CoinName field to given value.


### GetType

`func (o *InlineResponse20090DataOtherProfits) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20090DataOtherProfits) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20090DataOtherProfits) SetType(v int32)`

SetType sets Type field to given value.


### GetProfitAmount

`func (o *InlineResponse20090DataOtherProfits) GetProfitAmount() float64`

GetProfitAmount returns the ProfitAmount field if non-nil, zero value otherwise.

### GetProfitAmountOk

`func (o *InlineResponse20090DataOtherProfits) GetProfitAmountOk() (*float64, bool)`

GetProfitAmountOk returns a tuple with the ProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitAmount

`func (o *InlineResponse20090DataOtherProfits) SetProfitAmount(v float64)`

SetProfitAmount sets ProfitAmount field to given value.


### GetStatus

`func (o *InlineResponse20090DataOtherProfits) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20090DataOtherProfits) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20090DataOtherProfits) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


