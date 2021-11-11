# InlineResponse20098

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** |  | 
**Status** | **string** | S, P, and F for \&quot;success\&quot;, \&quot;pending\&quot;, and \&quot;failure\&quot; | 
**TokenName** | **string** |  | 
**Amount** | **string** | subscribed token amount | 
**Cost** | **string** | subscription cost in usdt | 
**Timestamp** | **int64** |  | 

## Methods

### NewInlineResponse20098

`func NewInlineResponse20098(id float64, status string, tokenName string, amount string, cost string, timestamp int64, ) *InlineResponse20098`

NewInlineResponse20098 instantiates a new InlineResponse20098 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20098WithDefaults

`func NewInlineResponse20098WithDefaults() *InlineResponse20098`

NewInlineResponse20098WithDefaults instantiates a new InlineResponse20098 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20098) GetId() float64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20098) GetIdOk() (*float64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20098) SetId(v float64)`

SetId sets Id field to given value.


### GetStatus

`func (o *InlineResponse20098) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20098) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20098) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenName

`func (o *InlineResponse20098) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *InlineResponse20098) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *InlineResponse20098) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetAmount

`func (o *InlineResponse20098) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InlineResponse20098) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InlineResponse20098) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCost

`func (o *InlineResponse20098) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *InlineResponse20098) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *InlineResponse20098) SetCost(v string)`

SetCost sets Cost field to given value.


### GetTimestamp

`func (o *InlineResponse20098) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse20098) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse20098) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


