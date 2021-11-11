# SnapshotSpot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int64** |  | 
**Msg** | **string** |  | 
**SnapshotVos** | [**[]SnapshotSpotSnapshotVos**](SnapshotSpotSnapshotVos.md) |  | 

## Methods

### NewSnapshotSpot

`func NewSnapshotSpot(code int64, msg string, snapshotVos []SnapshotSpotSnapshotVos, ) *SnapshotSpot`

NewSnapshotSpot instantiates a new SnapshotSpot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSpotWithDefaults

`func NewSnapshotSpotWithDefaults() *SnapshotSpot`

NewSnapshotSpotWithDefaults instantiates a new SnapshotSpot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SnapshotSpot) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SnapshotSpot) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SnapshotSpot) SetCode(v int64)`

SetCode sets Code field to given value.


### GetMsg

`func (o *SnapshotSpot) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SnapshotSpot) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SnapshotSpot) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetSnapshotVos

`func (o *SnapshotSpot) GetSnapshotVos() []SnapshotSpotSnapshotVos`

GetSnapshotVos returns the SnapshotVos field if non-nil, zero value otherwise.

### GetSnapshotVosOk

`func (o *SnapshotSpot) GetSnapshotVosOk() (*[]SnapshotSpotSnapshotVos, bool)`

GetSnapshotVosOk returns a tuple with the SnapshotVos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotVos

`func (o *SnapshotSpot) SetSnapshotVos(v []SnapshotSpotSnapshotVos)`

SetSnapshotVos sets SnapshotVos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


