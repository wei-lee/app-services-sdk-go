# TopicsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** | The page offset | 
**Limit** | **int32** | number of entries per page | 
**Count** | **int32** | Total number of topics | 
**Items** | [**[]Topic**](Topic.md) | List of topics | 

## Methods

### NewTopicsList

`func NewTopicsList(offset int32, limit int32, count int32, items []Topic, ) *TopicsList`

NewTopicsList instantiates a new TopicsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicsListWithDefaults

`func NewTopicsListWithDefaults() *TopicsList`

NewTopicsListWithDefaults instantiates a new TopicsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *TopicsList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TopicsList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TopicsList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *TopicsList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TopicsList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TopicsList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetCount

`func (o *TopicsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopicsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopicsList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetItems

`func (o *TopicsList) GetItems() []Topic`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TopicsList) GetItemsOk() (*[]Topic, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TopicsList) SetItems(v []Topic)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


