# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**Stages** | [**[]PipelineStage**](PipelineStage.md) |  | 
**Status** | [**PipelineStatus**](PipelineStatus.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPipeline

`func NewPipeline(stages []PipelineStage, status PipelineStatus, createdAt time.Time, ) *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pipeline) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pipeline) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pipeline) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Pipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefId

`func (o *Pipeline) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Pipeline) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Pipeline) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Pipeline) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStages

`func (o *Pipeline) GetStages() []PipelineStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *Pipeline) GetStagesOk() (*[]PipelineStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *Pipeline) SetStages(v []PipelineStage)`

SetStages sets Stages field to given value.


### GetStatus

`func (o *Pipeline) GetStatus() PipelineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pipeline) GetStatusOk() (*PipelineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pipeline) SetStatus(v PipelineStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Pipeline) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pipeline) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pipeline) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


