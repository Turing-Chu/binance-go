# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranId** | **int64** | transaction id | 

## Methods

### NewTransaction

`func NewTransaction(tranId int64, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranId

`func (o *Transaction) GetTranId() int64`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *Transaction) GetTranIdOk() (*int64, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *Transaction) SetTranId(v int64)`

SetTranId sets TranId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


