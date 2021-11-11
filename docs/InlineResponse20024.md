# InlineResponse20024

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Account&#39;s currently max borrowable amount with sufficient system availability | 
**BorrowLimit** | **string** | Max borrowable amount limited by the account level | 

## Methods

### NewInlineResponse20024

`func NewInlineResponse20024(amount string, borrowLimit string, ) *InlineResponse20024`

NewInlineResponse20024 instantiates a new InlineResponse20024 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20024WithDefaults

`func NewInlineResponse20024WithDefaults() *InlineResponse20024`

NewInlineResponse20024WithDefaults instantiates a new InlineResponse20024 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InlineResponse20024) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InlineResponse20024) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InlineResponse20024) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBorrowLimit

`func (o *InlineResponse20024) GetBorrowLimit() string`

GetBorrowLimit returns the BorrowLimit field if non-nil, zero value otherwise.

### GetBorrowLimitOk

`func (o *InlineResponse20024) GetBorrowLimitOk() (*string, bool)`

GetBorrowLimitOk returns a tuple with the BorrowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowLimit

`func (o *InlineResponse20024) SetBorrowLimit(v string)`

SetBorrowLimit sets BorrowLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


