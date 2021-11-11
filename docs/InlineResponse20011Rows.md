# InlineResponse20011Rows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsolatedSymbol** | **string** | Isolated symbol, will not be returned for crossed margin | 
**Amount** | **string** | Total amount repaid | 
**Asset** | **string** |  | 
**Interest** | **string** | Interest repaid | 
**Principal** | **string** | Principal repaid | 
**Status** | **string** | One of PENDING (pending execution), CONFIRMED (successfully execution), FAILED (execution failed, nothing happened to your account) | 
**Timestamp** | **int64** |  | 
**TxId** | **int64** |  | 

## Methods

### NewInlineResponse20011Rows

`func NewInlineResponse20011Rows(isolatedSymbol string, amount string, asset string, interest string, principal string, status string, timestamp int64, txId int64, ) *InlineResponse20011Rows`

NewInlineResponse20011Rows instantiates a new InlineResponse20011Rows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20011RowsWithDefaults

`func NewInlineResponse20011RowsWithDefaults() *InlineResponse20011Rows`

NewInlineResponse20011RowsWithDefaults instantiates a new InlineResponse20011Rows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsolatedSymbol

`func (o *InlineResponse20011Rows) GetIsolatedSymbol() string`

GetIsolatedSymbol returns the IsolatedSymbol field if non-nil, zero value otherwise.

### GetIsolatedSymbolOk

`func (o *InlineResponse20011Rows) GetIsolatedSymbolOk() (*string, bool)`

GetIsolatedSymbolOk returns a tuple with the IsolatedSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedSymbol

`func (o *InlineResponse20011Rows) SetIsolatedSymbol(v string)`

SetIsolatedSymbol sets IsolatedSymbol field to given value.


### GetAmount

`func (o *InlineResponse20011Rows) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InlineResponse20011Rows) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InlineResponse20011Rows) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAsset

`func (o *InlineResponse20011Rows) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *InlineResponse20011Rows) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *InlineResponse20011Rows) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetInterest

`func (o *InlineResponse20011Rows) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *InlineResponse20011Rows) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *InlineResponse20011Rows) SetInterest(v string)`

SetInterest sets Interest field to given value.


### GetPrincipal

`func (o *InlineResponse20011Rows) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *InlineResponse20011Rows) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *InlineResponse20011Rows) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetStatus

`func (o *InlineResponse20011Rows) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20011Rows) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20011Rows) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *InlineResponse20011Rows) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse20011Rows) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse20011Rows) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetTxId

`func (o *InlineResponse20011Rows) GetTxId() int64`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *InlineResponse20011Rows) GetTxIdOk() (*int64, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *InlineResponse20011Rows) SetTxId(v int64)`

SetTxId sets TxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


