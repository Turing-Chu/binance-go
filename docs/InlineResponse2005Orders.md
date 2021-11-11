# InlineResponse2005Orders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**OrderId** | **int64** |  | 
**ClientOrderId** | **string** |  | 

## Methods

### NewInlineResponse2005Orders

`func NewInlineResponse2005Orders(symbol string, orderId int64, clientOrderId string, ) *InlineResponse2005Orders`

NewInlineResponse2005Orders instantiates a new InlineResponse2005Orders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2005OrdersWithDefaults

`func NewInlineResponse2005OrdersWithDefaults() *InlineResponse2005Orders`

NewInlineResponse2005OrdersWithDefaults instantiates a new InlineResponse2005Orders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InlineResponse2005Orders) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineResponse2005Orders) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineResponse2005Orders) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderId

`func (o *InlineResponse2005Orders) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *InlineResponse2005Orders) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *InlineResponse2005Orders) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetClientOrderId

`func (o *InlineResponse2005Orders) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *InlineResponse2005Orders) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *InlineResponse2005Orders) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


