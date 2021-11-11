# SnapshotSpotData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | [**[]SnapshotSpotDataBalances**](SnapshotSpotDataBalances.md) |  | 
**TotalAssetOfBtc** | **string** |  | 

## Methods

### NewSnapshotSpotData

`func NewSnapshotSpotData(balances []SnapshotSpotDataBalances, totalAssetOfBtc string, ) *SnapshotSpotData`

NewSnapshotSpotData instantiates a new SnapshotSpotData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSpotDataWithDefaults

`func NewSnapshotSpotDataWithDefaults() *SnapshotSpotData`

NewSnapshotSpotDataWithDefaults instantiates a new SnapshotSpotData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *SnapshotSpotData) GetBalances() []SnapshotSpotDataBalances`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *SnapshotSpotData) GetBalancesOk() (*[]SnapshotSpotDataBalances, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *SnapshotSpotData) SetBalances(v []SnapshotSpotDataBalances)`

SetBalances sets Balances field to given value.


### GetTotalAssetOfBtc

`func (o *SnapshotSpotData) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *SnapshotSpotData) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *SnapshotSpotData) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


