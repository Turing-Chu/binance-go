# InlineResponse20022

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderListId** | **int64** |  | 
**ContingencyType** | **string** |  | 
**ListStatusType** | **string** |  | 
**ListOrderStatus** | **string** |  | 
**ListClientOrderId** | **string** |  | 
**TransactionTime** | **int64** |  | 
**Symbol** | **string** |  | 
**IsIsolated** | **bool** |  | 
**Orders** | [**[]ApiV3OpenOrderListOrders**](ApiV3OpenOrderListOrders.md) |  | 

## Methods

### NewInlineResponse20022

`func NewInlineResponse20022(orderListId int64, contingencyType string, listStatusType string, listOrderStatus string, listClientOrderId string, transactionTime int64, symbol string, isIsolated bool, orders []ApiV3OpenOrderListOrders, ) *InlineResponse20022`

NewInlineResponse20022 instantiates a new InlineResponse20022 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20022WithDefaults

`func NewInlineResponse20022WithDefaults() *InlineResponse20022`

NewInlineResponse20022WithDefaults instantiates a new InlineResponse20022 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *InlineResponse20022) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *InlineResponse20022) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *InlineResponse20022) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetContingencyType

`func (o *InlineResponse20022) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *InlineResponse20022) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *InlineResponse20022) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.


### GetListStatusType

`func (o *InlineResponse20022) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *InlineResponse20022) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *InlineResponse20022) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.


### GetListOrderStatus

`func (o *InlineResponse20022) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *InlineResponse20022) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *InlineResponse20022) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.


### GetListClientOrderId

`func (o *InlineResponse20022) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *InlineResponse20022) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *InlineResponse20022) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.


### GetTransactionTime

`func (o *InlineResponse20022) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *InlineResponse20022) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *InlineResponse20022) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.


### GetSymbol

`func (o *InlineResponse20022) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineResponse20022) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineResponse20022) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetIsIsolated

`func (o *InlineResponse20022) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *InlineResponse20022) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *InlineResponse20022) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.


### GetOrders

`func (o *InlineResponse20022) GetOrders() []ApiV3OpenOrderListOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *InlineResponse20022) GetOrdersOk() (*[]ApiV3OpenOrderListOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *InlineResponse20022) SetOrders(v []ApiV3OpenOrderListOrders)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


