# AggTrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | **int64** | Aggregate tradeId | 
**P** | **string** | Price | 
**Q** | **string** | Quantity | 
**F** | **int64** | First tradeId | 
**L** | **int64** | Last tradeId | 
**T** | **bool** | Timestamp | 
**M** | **bool** | Was the buyer the maker? | 
**M** | **bool** | Was the trade the best price match? | 

## Methods

### NewAggTrade

`func NewAggTrade(a int64, p string, q string, f int64, l int64, t bool, m bool, m bool, ) *AggTrade`

NewAggTrade instantiates a new AggTrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggTradeWithDefaults

`func NewAggTradeWithDefaults() *AggTrade`

NewAggTradeWithDefaults instantiates a new AggTrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *AggTrade) GetA() int64`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *AggTrade) GetAOk() (*int64, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *AggTrade) SetA(v int64)`

SetA sets A field to given value.


### GetP

`func (o *AggTrade) GetP() string`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *AggTrade) GetPOk() (*string, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *AggTrade) SetP(v string)`

SetP sets P field to given value.


### GetQ

`func (o *AggTrade) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *AggTrade) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *AggTrade) SetQ(v string)`

SetQ sets Q field to given value.


### GetF

`func (o *AggTrade) GetF() int64`

GetF returns the F field if non-nil, zero value otherwise.

### GetFOk

`func (o *AggTrade) GetFOk() (*int64, bool)`

GetFOk returns a tuple with the F field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetF

`func (o *AggTrade) SetF(v int64)`

SetF sets F field to given value.


### GetL

`func (o *AggTrade) GetL() int64`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *AggTrade) GetLOk() (*int64, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *AggTrade) SetL(v int64)`

SetL sets L field to given value.


### GetT

`func (o *AggTrade) GetT() bool`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AggTrade) GetTOk() (*bool, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AggTrade) SetT(v bool)`

SetT sets T field to given value.


### GetM

`func (o *AggTrade) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *AggTrade) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *AggTrade) SetM(v bool)`

SetM sets M field to given value.


### GetM

`func (o *AggTrade) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *AggTrade) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *AggTrade) SetM(v bool)`

SetM sets M field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


