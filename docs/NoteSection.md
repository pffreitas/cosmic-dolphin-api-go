# NoteSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NoteSectionType**](NoteSectionType.md) |  | 
**Content** | **map[string]string** |  | 

## Methods

### NewNoteSection

`func NewNoteSection(type_ NoteSectionType, content map[string]string, ) *NoteSection`

NewNoteSection instantiates a new NoteSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteSectionWithDefaults

`func NewNoteSectionWithDefaults() *NoteSection`

NewNoteSectionWithDefaults instantiates a new NoteSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NoteSection) GetType() NoteSectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoteSection) GetTypeOk() (*NoteSectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoteSection) SetType(v NoteSectionType)`

SetType sets Type field to given value.


### GetContent

`func (o *NoteSection) GetContent() map[string]string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteSection) GetContentOk() (*map[string]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteSection) SetContent(v map[string]string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


