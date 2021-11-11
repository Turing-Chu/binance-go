# InlineResponse20087Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerName** | **string** | Mining Account name | 
**Type** | **string** | Type of hourly hashrate | 
**HashrateDatas** | [**[]InlineResponse20087HashrateDatas**](InlineResponse20087HashrateDatas.md) |  | 

## Methods

### NewInlineResponse20087Data

`func NewInlineResponse20087Data(workerName string, type_ string, hashrateDatas []InlineResponse20087HashrateDatas, ) *InlineResponse20087Data`

NewInlineResponse20087Data instantiates a new InlineResponse20087Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20087DataWithDefaults

`func NewInlineResponse20087DataWithDefaults() *InlineResponse20087Data`

NewInlineResponse20087DataWithDefaults instantiates a new InlineResponse20087Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerName

`func (o *InlineResponse20087Data) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *InlineResponse20087Data) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *InlineResponse20087Data) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.


### GetType

`func (o *InlineResponse20087Data) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20087Data) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20087Data) SetType(v string)`

SetType sets Type field to given value.


### GetHashrateDatas

`func (o *InlineResponse20087Data) GetHashrateDatas() []InlineResponse20087HashrateDatas`

GetHashrateDatas returns the HashrateDatas field if non-nil, zero value otherwise.

### GetHashrateDatasOk

`func (o *InlineResponse20087Data) GetHashrateDatasOk() (*[]InlineResponse20087HashrateDatas, bool)`

GetHashrateDatasOk returns a tuple with the HashrateDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashrateDatas

`func (o *InlineResponse20087Data) SetHashrateDatas(v []InlineResponse20087HashrateDatas)`

SetHashrateDatas sets HashrateDatas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


