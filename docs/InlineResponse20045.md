# InlineResponse20045

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRestrict** | **bool** |  | 
**CreateTime** | **int64** |  | 
**EnableWithdrawals** | **bool** | This option allows you to withdraw via API. You must apply the IP Access Restriction filter in order to enable withdrawals | 
**EnableInternalTransfer** | **bool** | This option authorizes this key to transfer funds between your master account and your sub account instantly | 
**PermitsUniversalTransfer** | **bool** | Authorizes this key to be used for a dedicated universal transfer API to transfer multiple supported currencies. Each business&#39;s own transfer API rights are not affected by this authorization | 
**EnableVanillaOptions** | **bool** | Authorizes this key to Vanilla options trading | 
**EnableReading** | **bool** |  | 
**EnableFutures** | **bool** | API Key created before your futures account opened does not support futures API service | 
**EnableMargin** | **bool** | This option can be adjusted after the Cross Margin account transfer is completed | 
**EnableSpotAndMarginTrading** | **bool** |  | 
**TradingAuthorityExpirationTime** | **int64** | Expiration time for spot and margin trading permission | 

## Methods

### NewInlineResponse20045

`func NewInlineResponse20045(ipRestrict bool, createTime int64, enableWithdrawals bool, enableInternalTransfer bool, permitsUniversalTransfer bool, enableVanillaOptions bool, enableReading bool, enableFutures bool, enableMargin bool, enableSpotAndMarginTrading bool, tradingAuthorityExpirationTime int64, ) *InlineResponse20045`

NewInlineResponse20045 instantiates a new InlineResponse20045 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20045WithDefaults

`func NewInlineResponse20045WithDefaults() *InlineResponse20045`

NewInlineResponse20045WithDefaults instantiates a new InlineResponse20045 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRestrict

`func (o *InlineResponse20045) GetIpRestrict() bool`

GetIpRestrict returns the IpRestrict field if non-nil, zero value otherwise.

### GetIpRestrictOk

`func (o *InlineResponse20045) GetIpRestrictOk() (*bool, bool)`

GetIpRestrictOk returns a tuple with the IpRestrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestrict

`func (o *InlineResponse20045) SetIpRestrict(v bool)`

SetIpRestrict sets IpRestrict field to given value.


### GetCreateTime

`func (o *InlineResponse20045) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *InlineResponse20045) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *InlineResponse20045) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.


### GetEnableWithdrawals

`func (o *InlineResponse20045) GetEnableWithdrawals() bool`

GetEnableWithdrawals returns the EnableWithdrawals field if non-nil, zero value otherwise.

### GetEnableWithdrawalsOk

`func (o *InlineResponse20045) GetEnableWithdrawalsOk() (*bool, bool)`

GetEnableWithdrawalsOk returns a tuple with the EnableWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWithdrawals

`func (o *InlineResponse20045) SetEnableWithdrawals(v bool)`

SetEnableWithdrawals sets EnableWithdrawals field to given value.


### GetEnableInternalTransfer

`func (o *InlineResponse20045) GetEnableInternalTransfer() bool`

GetEnableInternalTransfer returns the EnableInternalTransfer field if non-nil, zero value otherwise.

### GetEnableInternalTransferOk

`func (o *InlineResponse20045) GetEnableInternalTransferOk() (*bool, bool)`

GetEnableInternalTransferOk returns a tuple with the EnableInternalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternalTransfer

`func (o *InlineResponse20045) SetEnableInternalTransfer(v bool)`

SetEnableInternalTransfer sets EnableInternalTransfer field to given value.


### GetPermitsUniversalTransfer

`func (o *InlineResponse20045) GetPermitsUniversalTransfer() bool`

GetPermitsUniversalTransfer returns the PermitsUniversalTransfer field if non-nil, zero value otherwise.

### GetPermitsUniversalTransferOk

`func (o *InlineResponse20045) GetPermitsUniversalTransferOk() (*bool, bool)`

GetPermitsUniversalTransferOk returns a tuple with the PermitsUniversalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitsUniversalTransfer

`func (o *InlineResponse20045) SetPermitsUniversalTransfer(v bool)`

SetPermitsUniversalTransfer sets PermitsUniversalTransfer field to given value.


### GetEnableVanillaOptions

`func (o *InlineResponse20045) GetEnableVanillaOptions() bool`

GetEnableVanillaOptions returns the EnableVanillaOptions field if non-nil, zero value otherwise.

### GetEnableVanillaOptionsOk

`func (o *InlineResponse20045) GetEnableVanillaOptionsOk() (*bool, bool)`

GetEnableVanillaOptionsOk returns a tuple with the EnableVanillaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVanillaOptions

`func (o *InlineResponse20045) SetEnableVanillaOptions(v bool)`

SetEnableVanillaOptions sets EnableVanillaOptions field to given value.


### GetEnableReading

`func (o *InlineResponse20045) GetEnableReading() bool`

GetEnableReading returns the EnableReading field if non-nil, zero value otherwise.

### GetEnableReadingOk

`func (o *InlineResponse20045) GetEnableReadingOk() (*bool, bool)`

GetEnableReadingOk returns a tuple with the EnableReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReading

`func (o *InlineResponse20045) SetEnableReading(v bool)`

SetEnableReading sets EnableReading field to given value.


### GetEnableFutures

`func (o *InlineResponse20045) GetEnableFutures() bool`

GetEnableFutures returns the EnableFutures field if non-nil, zero value otherwise.

### GetEnableFuturesOk

`func (o *InlineResponse20045) GetEnableFuturesOk() (*bool, bool)`

GetEnableFuturesOk returns a tuple with the EnableFutures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFutures

`func (o *InlineResponse20045) SetEnableFutures(v bool)`

SetEnableFutures sets EnableFutures field to given value.


### GetEnableMargin

`func (o *InlineResponse20045) GetEnableMargin() bool`

GetEnableMargin returns the EnableMargin field if non-nil, zero value otherwise.

### GetEnableMarginOk

`func (o *InlineResponse20045) GetEnableMarginOk() (*bool, bool)`

GetEnableMarginOk returns a tuple with the EnableMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMargin

`func (o *InlineResponse20045) SetEnableMargin(v bool)`

SetEnableMargin sets EnableMargin field to given value.


### GetEnableSpotAndMarginTrading

`func (o *InlineResponse20045) GetEnableSpotAndMarginTrading() bool`

GetEnableSpotAndMarginTrading returns the EnableSpotAndMarginTrading field if non-nil, zero value otherwise.

### GetEnableSpotAndMarginTradingOk

`func (o *InlineResponse20045) GetEnableSpotAndMarginTradingOk() (*bool, bool)`

GetEnableSpotAndMarginTradingOk returns a tuple with the EnableSpotAndMarginTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpotAndMarginTrading

`func (o *InlineResponse20045) SetEnableSpotAndMarginTrading(v bool)`

SetEnableSpotAndMarginTrading sets EnableSpotAndMarginTrading field to given value.


### GetTradingAuthorityExpirationTime

`func (o *InlineResponse20045) GetTradingAuthorityExpirationTime() int64`

GetTradingAuthorityExpirationTime returns the TradingAuthorityExpirationTime field if non-nil, zero value otherwise.

### GetTradingAuthorityExpirationTimeOk

`func (o *InlineResponse20045) GetTradingAuthorityExpirationTimeOk() (*int64, bool)`

GetTradingAuthorityExpirationTimeOk returns a tuple with the TradingAuthorityExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAuthorityExpirationTime

`func (o *InlineResponse20045) SetTradingAuthorityExpirationTime(v int64)`

SetTradingAuthorityExpirationTime sets TradingAuthorityExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


