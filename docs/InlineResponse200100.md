# InlineResponse200100

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** |  | 
**Status** | **string** | S, P, and F for \&quot;success\&quot;, \&quot;pending\&quot;, and \&quot;failure\&quot; | 
**TokenName** | **string** |  | 
**RedeemAmount** | **string** | Redemption token amount | 
**Amount** | **string** | Redemption value in usdt | 
**Timestamp** | **int64** |  | 

## Methods

### NewInlineResponse200100

`func NewInlineResponse200100(id float64, status string, tokenName string, redeemAmount string, amount string, timestamp int64, ) *InlineResponse200100`

NewInlineResponse200100 instantiates a new InlineResponse200100 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200100WithDefaults

`func NewInlineResponse200100WithDefaults() *InlineResponse200100`

NewInlineResponse200100WithDefaults instantiates a new InlineResponse200100 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200100) GetId() float64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200100) GetIdOk() (*float64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200100) SetId(v float64)`

SetId sets Id field to given value.


### GetStatus

`func (o *InlineResponse200100) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200100) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200100) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenName

`func (o *InlineResponse200100) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *InlineResponse200100) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *InlineResponse200100) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetRedeemAmount

`func (o *InlineResponse200100) GetRedeemAmount() string`

GetRedeemAmount returns the RedeemAmount field if non-nil, zero value otherwise.

### GetRedeemAmountOk

`func (o *InlineResponse200100) GetRedeemAmountOk() (*string, bool)`

GetRedeemAmountOk returns a tuple with the RedeemAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemAmount

`func (o *InlineResponse200100) SetRedeemAmount(v string)`

SetRedeemAmount sets RedeemAmount field to given value.


### GetAmount

`func (o *InlineResponse200100) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InlineResponse200100) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InlineResponse200100) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTimestamp

`func (o *InlineResponse200100) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse200100) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse200100) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


