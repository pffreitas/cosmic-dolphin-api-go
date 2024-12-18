# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DocumentId** | Pointer to **int64** |  | [optional] 
**Title** | **string** |  | 
**Summary** | **string** |  | 
**Tags** | **string** |  | 
**Type** | [**NoteType**](NoteType.md) |  | 
**Sections** | [**[]NoteSection**](NoteSection.md) |  | 
**UserId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewNote

`func NewNote(title string, summary string, tags string, type_ NoteType, sections []NoteSection, userId string, createdAt time.Time, ) *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Note) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Note) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Note) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Note) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentId

`func (o *Note) GetDocumentId() int64`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *Note) GetDocumentIdOk() (*int64, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *Note) SetDocumentId(v int64)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *Note) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetTitle

`func (o *Note) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Note) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Note) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *Note) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Note) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Note) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTags

`func (o *Note) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Note) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Note) SetTags(v string)`

SetTags sets Tags field to given value.


### GetType

`func (o *Note) GetType() NoteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Note) GetTypeOk() (*NoteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Note) SetType(v NoteType)`

SetType sets Type field to given value.


### GetSections

`func (o *Note) GetSections() []NoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Note) GetSectionsOk() (*[]NoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *Note) SetSections(v []NoteSection)`

SetSections sets Sections field to given value.


### GetUserId

`func (o *Note) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Note) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Note) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *Note) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Note) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Note) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


