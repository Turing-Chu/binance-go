# MarginOrderResponseAck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**OrderId** | **int64** |  | 
**ClientOrderId** | **string** |  | 
**IsIsolated** | **bool** |  | 
**TransactTime** | **int64** |  | 

## Methods

### NewMarginOrderResponseAck

`func NewMarginOrderResponseAck(symbol string, orderId int64, clientOrderId string, isIsolated bool, transactTime int64, ) *MarginOrderResponseAck`

NewMarginOrderResponseAck instantiates a new MarginOrderResponseAck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginOrderResponseAckWithDefaults

`func NewMarginOrderResponseAckWithDefaults() *MarginOrderResponseAck`

NewMarginOrderResponseAckWithDefaults instantiates a new MarginOrderResponseAck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MarginOrderResponseAck) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginOrderResponseAck) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginOrderResponseAck) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderId

`func (o *MarginOrderResponseAck) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginOrderResponseAck) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginOrderResponseAck) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetClientOrderId

`func (o *MarginOrderResponseAck) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *MarginOrderResponseAck) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *MarginOrderResponseAck) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetIsIsolated

`func (o *MarginOrderResponseAck) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginOrderResponseAck) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginOrderResponseAck) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.


### GetTransactTime

`func (o *MarginOrderResponseAck) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *MarginOrderResponseAck) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *MarginOrderResponseAck) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


