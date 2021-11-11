# OcoOrder

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
**Orders** | [**[]InlineResponse20020Orders**](InlineResponse20020Orders.md) |  | 
**OrderReports** | [**[]OcoOrderOrderReports**](OcoOrderOrderReports.md) |  | 

## Methods

### NewOcoOrder

`func NewOcoOrder(orderListId int64, contingencyType string, listStatusType string, listOrderStatus string, listClientOrderId string, transactionTime int64, symbol string, orders []InlineResponse20020Orders, orderReports []OcoOrderOrderReports, ) *OcoOrder`

NewOcoOrder instantiates a new OcoOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOcoOrderWithDefaults

`func NewOcoOrderWithDefaults() *OcoOrder`

NewOcoOrderWithDefaults instantiates a new OcoOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderListId

`func (o *OcoOrder) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OcoOrder) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OcoOrder) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.


### GetContingencyType

`func (o *OcoOrder) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *OcoOrder) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *OcoOrder) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.


### GetListStatusType

`func (o *OcoOrder) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *OcoOrder) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *OcoOrder) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.


### GetListOrderStatus

`func (o *OcoOrder) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *OcoOrder) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *OcoOrder) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.


### GetListClientOrderId

`func (o *OcoOrder) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *OcoOrder) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *OcoOrder) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.


### GetTransactionTime

`func (o *OcoOrder) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *OcoOrder) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *OcoOrder) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.


### GetSymbol

`func (o *OcoOrder) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OcoOrder) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OcoOrder) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrders

`func (o *OcoOrder) GetOrders() []InlineResponse20020Orders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OcoOrder) GetOrdersOk() (*[]InlineResponse20020Orders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OcoOrder) SetOrders(v []InlineResponse20020Orders)`

SetOrders sets Orders field to given value.


### GetOrderReports

`func (o *OcoOrder) GetOrderReports() []OcoOrderOrderReports`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *OcoOrder) GetOrderReportsOk() (*[]OcoOrderOrderReports, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *OcoOrder) SetOrderReports(v []OcoOrderOrderReports)`

SetOrderReports sets OrderReports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


