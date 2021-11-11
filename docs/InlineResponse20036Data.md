# InlineResponse20036Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLocked** | **bool** | API trading function is locked or not | 
**PlannedRecoverTime** | **int64** | If API trading function is locked, this is the planned recover time | 
**TriggerCondition** | [**InlineResponse20036DataTriggerCondition**](InlineResponse20036DataTriggerCondition.md) |  | 
**Indicators** | [**InlineResponse20036DataIndicators**](InlineResponse20036DataIndicators.md) |  | 
**UpdateTime** | **int64** |  | 

## Methods

### NewInlineResponse20036Data

`func NewInlineResponse20036Data(isLocked bool, plannedRecoverTime int64, triggerCondition InlineResponse20036DataTriggerCondition, indicators InlineResponse20036DataIndicators, updateTime int64, ) *InlineResponse20036Data`

NewInlineResponse20036Data instantiates a new InlineResponse20036Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20036DataWithDefaults

`func NewInlineResponse20036DataWithDefaults() *InlineResponse20036Data`

NewInlineResponse20036DataWithDefaults instantiates a new InlineResponse20036Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLocked

`func (o *InlineResponse20036Data) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *InlineResponse20036Data) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *InlineResponse20036Data) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetPlannedRecoverTime

`func (o *InlineResponse20036Data) GetPlannedRecoverTime() int64`

GetPlannedRecoverTime returns the PlannedRecoverTime field if non-nil, zero value otherwise.

### GetPlannedRecoverTimeOk

`func (o *InlineResponse20036Data) GetPlannedRecoverTimeOk() (*int64, bool)`

GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRecoverTime

`func (o *InlineResponse20036Data) SetPlannedRecoverTime(v int64)`

SetPlannedRecoverTime sets PlannedRecoverTime field to given value.


### GetTriggerCondition

`func (o *InlineResponse20036Data) GetTriggerCondition() InlineResponse20036DataTriggerCondition`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *InlineResponse20036Data) GetTriggerConditionOk() (*InlineResponse20036DataTriggerCondition, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *InlineResponse20036Data) SetTriggerCondition(v InlineResponse20036DataTriggerCondition)`

SetTriggerCondition sets TriggerCondition field to given value.


### GetIndicators

`func (o *InlineResponse20036Data) GetIndicators() InlineResponse20036DataIndicators`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *InlineResponse20036Data) GetIndicatorsOk() (*InlineResponse20036DataIndicators, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *InlineResponse20036Data) SetIndicators(v InlineResponse20036DataIndicators)`

SetIndicators sets Indicators field to given value.


### GetUpdateTime

`func (o *InlineResponse20036Data) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *InlineResponse20036Data) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *InlineResponse20036Data) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


