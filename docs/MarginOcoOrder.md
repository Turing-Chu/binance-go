# MarginOcoOrder

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
**Orders** | [**[]MarginOcoOrderOrders**](MarginOcoOrderOrders.md) |  | 
**OrderReports** | [**[]MarginOcoOrderOrderReports**](MarginOcoOrderOrderReports.md) |  | 

## Methods

### NewMarginOcoOrder

`func NewMarginOcoOrder(orderListId int64, contingencyType string, listStatusType string, listOrderStatus string, listClientOrderId string, transactionTime int64, symbol string, isIsolated bool, orders []MarginOcoOrderOrders, orderReports []MarginOcoOrderOrderReports, ) *MarginOcoOrder`

NewMarginOcoOrder instantiates a new MarginOcoOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginOcoOrderWithDefaults

`func NewMarginOcoOrderWithDefaults() *MarginOcoOrder`

NewMarginOcoOrderWithDefaults instantiates a new MarginOcoOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *MarginOcoOrder) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginOcoOrder) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginOcoOrder) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetContingencyType

`func (o *MarginOcoOrder) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginOcoOrder) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginOcoOrder) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.


### GetListStatusType

`func (o *MarginOcoOrder) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginOcoOrder) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginOcoOrder) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.


### GetListOrderStatus

`func (o *MarginOcoOrder) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginOcoOrder) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginOcoOrder) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.


### GetListClientOrderId

`func (o *MarginOcoOrder) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginOcoOrder) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginOcoOrder) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.


### GetTransactionTime

`func (o *MarginOcoOrder) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginOcoOrder) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginOcoOrder) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.


### GetSymbol

`func (o *MarginOcoOrder) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginOcoOrder) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginOcoOrder) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetIsIsolated

`func (o *MarginOcoOrder) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginOcoOrder) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginOcoOrder) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.


### GetOrders

`func (o *MarginOcoOrder) GetOrders() []MarginOcoOrderOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginOcoOrder) GetOrdersOk() (*[]MarginOcoOrderOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginOcoOrder) SetOrders(v []MarginOcoOrderOrders)`

SetOrders sets Orders field to given value.


### GetOrderReports

`func (o *MarginOcoOrder) GetOrderReports() []MarginOcoOrderOrderReports`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *MarginOcoOrder) GetOrderReportsOk() (*[]MarginOcoOrderOrderReports, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *MarginOcoOrder) SetOrderReports(v []MarginOcoOrderOrderReports)`

SetOrderReports sets OrderReports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


