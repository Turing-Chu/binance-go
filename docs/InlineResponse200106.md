# InlineResponse200106

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **int64** |  | 
**PoolId** | **int64** |  | 
**PoolName** | **string** |  | 
**Operation** | **string** | \&quot;ADD\&quot; or \&quot;REMOVE\&quot; | 
**Status** | **int32** | 0: pending, 1: success, 2: failed | 
**UpdateTime** | **int64** |  | 
**ShareAmount** | **string** |  | 

## Methods

### NewInlineResponse200106

`func NewInlineResponse200106(operationId int64, poolId int64, poolName string, operation string, status int32, updateTime int64, shareAmount string, ) *InlineResponse200106`

NewInlineResponse200106 instantiates a new InlineResponse200106 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200106WithDefaults

`func NewInlineResponse200106WithDefaults() *InlineResponse200106`

NewInlineResponse200106WithDefaults instantiates a new InlineResponse200106 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *InlineResponse200106) GetOperationId() int64`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *InlineResponse200106) GetOperationIdOk() (*int64, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *InlineResponse200106) SetOperationId(v int64)`

SetOperationId sets OperationId field to given value.


### GetPoolId

`func (o *InlineResponse200106) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *InlineResponse200106) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *InlineResponse200106) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetPoolName

`func (o *InlineResponse200106) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *InlineResponse200106) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *InlineResponse200106) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### GetOperation

`func (o *InlineResponse200106) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse200106) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse200106) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetStatus

`func (o *InlineResponse200106) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200106) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200106) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetUpdateTime

`func (o *InlineResponse200106) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *InlineResponse200106) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *InlineResponse200106) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.


### GetShareAmount

`func (o *InlineResponse200106) GetShareAmount() string`

GetShareAmount returns the ShareAmount field if non-nil, zero value otherwise.

### GetShareAmountOk

`func (o *InlineResponse200106) GetShareAmountOk() (*string, bool)`

GetShareAmountOk returns a tuple with the ShareAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareAmount

`func (o *InlineResponse200106) SetShareAmount(v string)`

SetShareAmount sets ShareAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


