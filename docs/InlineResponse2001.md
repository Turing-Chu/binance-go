# InlineResponse2001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | **string** |  | 
**ServerTime** | **int64** |  | 
**RateLimits** | [**[]InlineResponse2001RateLimits**](InlineResponse2001RateLimits.md) |  | 
**ExchangeFilters** | **[]map[string]interface{}** |  | 
**Symbols** | [**[]InlineResponse2001Symbols**](InlineResponse2001Symbols.md) |  | 

## Methods

### NewInlineResponse2001

`func NewInlineResponse2001(timezone string, serverTime int64, rateLimits []InlineResponse2001RateLimits, exchangeFilters []map[string]interface{}, symbols []InlineResponse2001Symbols, ) *InlineResponse2001`

NewInlineResponse2001 instantiates a new InlineResponse2001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001WithDefaults

`func NewInlineResponse2001WithDefaults() *InlineResponse2001`

NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *InlineResponse2001) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InlineResponse2001) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InlineResponse2001) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetServerTime

`func (o *InlineResponse2001) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *InlineResponse2001) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *InlineResponse2001) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.


### GetRateLimits

`func (o *InlineResponse2001) GetRateLimits() []InlineResponse2001RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *InlineResponse2001) GetRateLimitsOk() (*[]InlineResponse2001RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *InlineResponse2001) SetRateLimits(v []InlineResponse2001RateLimits)`

SetRateLimits sets RateLimits field to given value.


### GetExchangeFilters

`func (o *InlineResponse2001) GetExchangeFilters() []map[string]interface{}`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *InlineResponse2001) GetExchangeFiltersOk() (*[]map[string]interface{}, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *InlineResponse2001) SetExchangeFilters(v []map[string]interface{})`

SetExchangeFilters sets ExchangeFilters field to given value.


### GetSymbols

`func (o *InlineResponse2001) GetSymbols() []InlineResponse2001Symbols`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *InlineResponse2001) GetSymbolsOk() (*[]InlineResponse2001Symbols, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *InlineResponse2001) SetSymbols(v []InlineResponse2001Symbols)`

SetSymbols sets Symbols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


