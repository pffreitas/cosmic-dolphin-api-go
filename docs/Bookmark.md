# Bookmark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SourceUrl** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**BookmarkMetadata**](BookmarkMetadata.md) |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 
**IsArchived** | Pointer to **bool** |  | [optional] 
**IsFavorite** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 

## Methods

### NewBookmark

`func NewBookmark(id string, sourceUrl string, userId string, ) *Bookmark`

NewBookmark instantiates a new Bookmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkWithDefaults

`func NewBookmarkWithDefaults() *Bookmark`

NewBookmarkWithDefaults instantiates a new Bookmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bookmark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bookmark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bookmark) SetId(v string)`

SetId sets Id field to given value.


### GetSourceUrl

`func (o *Bookmark) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *Bookmark) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *Bookmark) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetTitle

`func (o *Bookmark) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Bookmark) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Bookmark) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Bookmark) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Bookmark) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bookmark) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bookmark) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Bookmark) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Bookmark) GetMetadata() BookmarkMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Bookmark) GetMetadataOk() (*BookmarkMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Bookmark) SetMetadata(v BookmarkMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Bookmark) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCollectionId

`func (o *Bookmark) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Bookmark) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Bookmark) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *Bookmark) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetUserId

`func (o *Bookmark) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Bookmark) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Bookmark) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsArchived

`func (o *Bookmark) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Bookmark) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Bookmark) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *Bookmark) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetIsFavorite

`func (o *Bookmark) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *Bookmark) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *Bookmark) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *Bookmark) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetTags

`func (o *Bookmark) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Bookmark) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Bookmark) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Bookmark) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetContent

`func (o *Bookmark) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Bookmark) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Bookmark) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Bookmark) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSummary

`func (o *Bookmark) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Bookmark) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Bookmark) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Bookmark) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


