# WebSocketMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DocumentId** | Pointer to **int64** |  | [optional] 
**Title** | **string** |  | 
**Summary** | **string** |  | 
**Tags** | **string** |  | 
**Sections** | [**[]NoteSection**](NoteSection.md) |  | 
**UserId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewWebSocketMessageData

`func NewWebSocketMessageData(title string, summary string, tags string, sections []NoteSection, userId string, createdAt time.Time, ) *WebSocketMessageData`

NewWebSocketMessageData instantiates a new WebSocketMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSocketMessageDataWithDefaults

`func NewWebSocketMessageDataWithDefaults() *WebSocketMessageData`

NewWebSocketMessageDataWithDefaults instantiates a new WebSocketMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebSocketMessageData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebSocketMessageData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebSocketMessageData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WebSocketMessageData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentId

`func (o *WebSocketMessageData) GetDocumentId() int64`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *WebSocketMessageData) GetDocumentIdOk() (*int64, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *WebSocketMessageData) SetDocumentId(v int64)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *WebSocketMessageData) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetTitle

`func (o *WebSocketMessageData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebSocketMessageData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebSocketMessageData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *WebSocketMessageData) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *WebSocketMessageData) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *WebSocketMessageData) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTags

`func (o *WebSocketMessageData) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WebSocketMessageData) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WebSocketMessageData) SetTags(v string)`

SetTags sets Tags field to given value.


### GetSections

`func (o *WebSocketMessageData) GetSections() []NoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *WebSocketMessageData) GetSectionsOk() (*[]NoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *WebSocketMessageData) SetSections(v []NoteSection)`

SetSections sets Sections field to given value.


### GetUserId

`func (o *WebSocketMessageData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebSocketMessageData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebSocketMessageData) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *WebSocketMessageData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebSocketMessageData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebSocketMessageData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


