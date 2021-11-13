# InlineResponse2006

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
**Orders** | [**[]AllOrderListOrders**](AllOrderListOrders.md) |  | 

## Methods

### NewInlineResponse2006

`func NewInlineResponse2006(orderListId int64, contingencyType string, listStatusType string, listOrderStatus string, listClientOrderId string, transactionTime int64, symbol string, isIsolated bool, orders []AllOrderListOrders, ) *InlineResponse2006`

NewInlineResponse2006 instantiates a new InlineResponse2006 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2006WithDefaults

`func NewInlineResponse2006WithDefaults() *InlineResponse2006`

NewInlineResponse2006WithDefaults instantiates a new InlineResponse2006 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *InlineResponse2006) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *InlineResponse2006) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *InlineResponse2006) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetContingencyType

`func (o *InlineResponse2006) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *InlineResponse2006) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *InlineResponse2006) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.


### GetListStatusType

`func (o *InlineResponse2006) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *InlineResponse2006) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *InlineResponse2006) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.


### GetListOrderStatus

`func (o *InlineResponse2006) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *InlineResponse2006) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *InlineResponse2006) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.


### GetListClientOrderId

`func (o *InlineResponse2006) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *InlineResponse2006) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *InlineResponse2006) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.


### GetTransactionTime

`func (o *InlineResponse2006) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *InlineResponse2006) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *InlineResponse2006) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.


### GetSymbol

`func (o *InlineResponse2006) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineResponse2006) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineResponse2006) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetIsIsolated

`func (o *InlineResponse2006) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *InlineResponse2006) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *InlineResponse2006) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.


### GetOrders

`func (o *InlineResponse2006) GetOrders() []AllOrderListOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *InlineResponse2006) GetOrdersOk() (*[]AllOrderListOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *InlineResponse2006) SetOrders(v []AllOrderListOrders)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


