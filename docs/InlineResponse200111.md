# InlineResponse200111

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Message** | **string** |  | 
**Data** | [**[]InlineResponse200111Data**](InlineResponse200111Data.md) |  | 
**Total** | **int32** |  | 
**Success** | **bool** |  | 

## Methods

### NewInlineResponse200111

`func NewInlineResponse200111(code string, message string, data []InlineResponse200111Data, total int32, success bool, ) *InlineResponse200111`

NewInlineResponse200111 instantiates a new InlineResponse200111 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200111WithDefaults

`func NewInlineResponse200111WithDefaults() *InlineResponse200111`

NewInlineResponse200111WithDefaults instantiates a new InlineResponse200111 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InlineResponse200111) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse200111) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse200111) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *InlineResponse200111) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse200111) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse200111) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *InlineResponse200111) GetData() []InlineResponse200111Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse200111) GetDataOk() (*[]InlineResponse200111Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse200111) SetData(v []InlineResponse200111Data)`

SetData sets Data field to given value.


### GetTotal

`func (o *InlineResponse200111) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200111) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200111) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetSuccess

`func (o *InlineResponse200111) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InlineResponse200111) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InlineResponse200111) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


