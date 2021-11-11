# InlineResponse20091DataConfigDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | **int64** | Mining ID | 
**PoolUsername** | **string** | Transfer out of subaccount | 
**ToPoolUsername** | **string** | Transfer into subaccount | 
**AlgoName** | **string** | Transfer algorithm | 
**HashRate** | **int64** | Transferred Hashrate quantity | 
**StartDay** | **int64** | Start date | 
**EndDay** | **int64** | End date | 
**Status** | **int32** | 0 Processing, 1：Cancelled, 2：Terminated  | 

## Methods

### NewInlineResponse20091DataConfigDetails

`func NewInlineResponse20091DataConfigDetails(configId int64, poolUsername string, toPoolUsername string, algoName string, hashRate int64, startDay int64, endDay int64, status int32, ) *InlineResponse20091DataConfigDetails`

NewInlineResponse20091DataConfigDetails instantiates a new InlineResponse20091DataConfigDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20091DataConfigDetailsWithDefaults

`func NewInlineResponse20091DataConfigDetailsWithDefaults() *InlineResponse20091DataConfigDetails`

NewInlineResponse20091DataConfigDetailsWithDefaults instantiates a new InlineResponse20091DataConfigDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *InlineResponse20091DataConfigDetails) GetConfigId() int64`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *InlineResponse20091DataConfigDetails) GetConfigIdOk() (*int64, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *InlineResponse20091DataConfigDetails) SetConfigId(v int64)`

SetConfigId sets ConfigId field to given value.


### GetPoolUsername

`func (o *InlineResponse20091DataConfigDetails) GetPoolUsername() string`

GetPoolUsername returns the PoolUsername field if non-nil, zero value otherwise.

### GetPoolUsernameOk

`func (o *InlineResponse20091DataConfigDetails) GetPoolUsernameOk() (*string, bool)`

GetPoolUsernameOk returns a tuple with the PoolUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolUsername

`func (o *InlineResponse20091DataConfigDetails) SetPoolUsername(v string)`

SetPoolUsername sets PoolUsername field to given value.


### GetToPoolUsername

`func (o *InlineResponse20091DataConfigDetails) GetToPoolUsername() string`

GetToPoolUsername returns the ToPoolUsername field if non-nil, zero value otherwise.

### GetToPoolUsernameOk

`func (o *InlineResponse20091DataConfigDetails) GetToPoolUsernameOk() (*string, bool)`

GetToPoolUsernameOk returns a tuple with the ToPoolUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPoolUsername

`func (o *InlineResponse20091DataConfigDetails) SetToPoolUsername(v string)`

SetToPoolUsername sets ToPoolUsername field to given value.


### GetAlgoName

`func (o *InlineResponse20091DataConfigDetails) GetAlgoName() string`

GetAlgoName returns the AlgoName field if non-nil, zero value otherwise.

### GetAlgoNameOk

`func (o *InlineResponse20091DataConfigDetails) GetAlgoNameOk() (*string, bool)`

GetAlgoNameOk returns a tuple with the AlgoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoName

`func (o *InlineResponse20091DataConfigDetails) SetAlgoName(v string)`

SetAlgoName sets AlgoName field to given value.


### GetHashRate

`func (o *InlineResponse20091DataConfigDetails) GetHashRate() int64`

GetHashRate returns the HashRate field if non-nil, zero value otherwise.

### GetHashRateOk

`func (o *InlineResponse20091DataConfigDetails) GetHashRateOk() (*int64, bool)`

GetHashRateOk returns a tuple with the HashRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashRate

`func (o *InlineResponse20091DataConfigDetails) SetHashRate(v int64)`

SetHashRate sets HashRate field to given value.


### GetStartDay

`func (o *InlineResponse20091DataConfigDetails) GetStartDay() int64`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *InlineResponse20091DataConfigDetails) GetStartDayOk() (*int64, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *InlineResponse20091DataConfigDetails) SetStartDay(v int64)`

SetStartDay sets StartDay field to given value.


### GetEndDay

`func (o *InlineResponse20091DataConfigDetails) GetEndDay() int64`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *InlineResponse20091DataConfigDetails) GetEndDayOk() (*int64, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *InlineResponse20091DataConfigDetails) SetEndDay(v int64)`

SetEndDay sets EndDay field to given value.


### GetStatus

`func (o *InlineResponse20091DataConfigDetails) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20091DataConfigDetails) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20091DataConfigDetails) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


